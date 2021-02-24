package main

import (
	"fmt"
	"go/build"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"golang.org/x/tools/go/packages"

	ggkconf "github.com/ben1009/truss/gengokit"
	gengokit "github.com/ben1009/truss/gengokit/generator"
	"github.com/ben1009/truss/svcdef"
	"github.com/ben1009/truss/truss"
	"github.com/ben1009/truss/truss/execprotoc"
	"github.com/ben1009/truss/truss/getstarted"
	"github.com/ben1009/truss/truss/parsesvcname"
)

var (
	serviceNameFlag = flag.String("servicename", "", "The serviceName need created")
	createFlag      = flag.BoolP("create", "c", false, "Initialize a new service with the input serviceName")
	updateFlag      = flag.BoolP("update", "u", false, "Update the service created before, run the command at the <serviceName> folder")
	verboseFlag     = flag.BoolP("verbose", "v", false, "Verbose output")
	helpFlag        = flag.BoolP("help", "h", false, "Print usage")
)

var binName = filepath.Base(os.Args[0])

var (
	// version is compiled into truss with the flag
	// go install -ldflags "-X main.version=$SHA"
	version string
	// BuildDate is compiled into truss with the flag
	// go install -ldflags "-X main.date=$VERSION_DATE"
	date string
)

func init() {
	// If Version or VersionDate are not set, truss was not built with make
	if version == "" || date == "" {
		rebuild := promptNoMake()
		if !rebuild {
			os.Exit(1)
		}
		err := makeAndRunTruss(os.Args)
		if err != nil {
			log.Fatal(errors.Wrap(err, "please install truss with make manually"))
		}
		os.Exit(0)
	}

	var buildinfo string
	buildinfo = fmt.Sprintf("version: %s", version)
	buildinfo = fmt.Sprintf("%s version date: %s", buildinfo, date)

	flag.Usage = func() {
		if buildinfo != "" && (*verboseFlag || *helpFlag) {
			fmt.Fprintf(os.Stderr, "%s (%s)\n", binName, strings.TrimSpace(buildinfo))
		}
		fmt.Fprintf(os.Stderr, "\nUsage: %s [options] <serviceName>...\n", binName)
		fmt.Fprintf(os.Stderr, "\nGenerates go-kit services using proto3 and gRPC definitions.\n")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}
	log.SetLevel(log.InfoLevel)
	if *verboseFlag {
		log.SetLevel(log.DebugLevel)
	}
	if !*createFlag && !*updateFlag {
		_, _ = fmt.Fprintf(os.Stderr, "%s: please choose create or update option\n", binName)
		flag.Usage()
		os.Exit(1)
	}

	cfg := &truss.Config{}
	if *createFlag {
		if *serviceNameFlag == "" {
			_, _ = fmt.Fprintf(os.Stderr, "%s: missing serviceName\n", binName)
			flag.Usage()
			os.Exit(1)
		}

		err := parseCreateInput(cfg, *serviceNameFlag)
		if err != nil {
			log.Fatal(errors.Wrap(err, "cannot parse create input"))
			return
		}
	} else if *updateFlag {
		err := parseUpdateInput(cfg)
		if err != nil {
			log.Fatal(errors.Wrap(err, "cannot parse update input"))
			return
		}
	}

	sd, err := parseServiceDefinition(cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot parse input definition proto files"))
		return
	}

	genFiles, err := generateCode(cfg, sd)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot generate service"))
		return
	}

	for p, file := range genFiles {
		err := writeGenFile(file, filepath.Join(cfg.ServicePath, p))
		if err != nil {
			log.Fatal(errors.Wrap(err, "cannot to write output"))
			return
		}
	}
	runGoImports(cfg.ServicePath)
}

// runGoImports method to execute goimports on the given file / folder
func runGoImports(fileName string) {
	run(exec.Command("goimports", "-w", fileName))
}

// run runs a command with error handling. Fatal means the program will exit on error, so no error value is returned.
func run(cmd *exec.Cmd) {
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Fatal(fmt.Sprintf("Failed to run command: %s %s\n\nError: %v\n\nOutput: %s\n", cmd.Path, strings.Join(cmd.Args, " "), err, string(out)))
	}
}

func dirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

func genProto(cfg *truss.Config, svcName string) error {
	outPath, _ := os.Getwd()
	svcPath := path.Join(outPath, svcName)
	if dirExists(svcPath) {
		return errors.Errorf("dir:%s already exist, either change to a different serviceName or update instead", svcPath)
	}

	err := os.MkdirAll(svcPath, 0o755)
	if err != nil {
		return errors.Wrapf(err, "cannot create svcPath directory: %s", svcPath)
	}
	cfg.ServicePath = svcPath
	log.WithField("Service Path", cfg.ServicePath).Debug()

	pbPath := path.Join(svcPath, "pb")
	err = os.MkdirAll(pbPath, 0o755)
	if err != nil {
		return errors.Wrapf(err, "cannot create pbPath directory: %s", pbPath)
	}

	err = getstarted.Gen(svcName, pbPath)
	if err != nil {
		return errors.Wrap(err, "getstarted.Gen")
	}

	cfg.DefPaths = []string{fmt.Sprintf("%s/%s.proto", pbPath, svcName)}
	log.WithField("DefPaths", cfg.DefPaths).Debug()
	return nil
}

func getProto(cfg *truss.Config) error {
	svcPath, _ := os.Getwd()
	cfg.ServicePath = svcPath
	log.WithField("Service Path", cfg.ServicePath).Debug()

	pbPaths, err := findProtoFiles(cfg.ServicePath)
	if err != nil {
		return errors.Wrap(err, "getProto.findProtoFiles")
	}

	cfg.DefPaths = pbPaths
	log.WithField("DefPaths", cfg.DefPaths).Debug()
	return nil
}

func findProtoFiles(root string) ([]string, error) {
	includePaths := []string{path.Join(root, "pb")}
	protoFiles := []string{}
	for _, pbDir := range includePaths {
		fi, err := os.Stat(pbDir)
		if err != nil {
			return nil, errors.Wrap(err, "os.Stat")
		}
		mode := fi.Mode()
		if !mode.IsDir() {
			return nil, errors.New("pb directory not found")
		}

		f, err := os.Open(pbDir)
		if err != nil {
			return nil, errors.Wrap(err, "os.Open")
		}

		files, err := f.Readdirnames(0)
		if err != nil {
			return nil, errors.Wrap(err, "Readdirnames")
		}

		for _, f := range files {
			if strings.HasSuffix(f, ".proto") {
				protoFiles = append(protoFiles, path.Join(pbDir, f))
			}
		}
	}
	return protoFiles, nil
}

func parseUpdateInput(cfg *truss.Config) error {
	err := getProto(cfg)
	if err != nil {
		return errors.Wrap(err, "getProto")
	}

	err = setCfg(cfg)
	return errors.Wrap(err, "setCfg")
}

func parseCreateInput(cfg *truss.Config, svcName string) error {
	err := genProto(cfg, svcName)
	if err != nil {
		return errors.Wrap(err, "genProto")
	}

	err = setCfg(cfg)
	return errors.Wrap(err, "setCfg")
}

func setCfg(cfg *truss.Config) error {
	protoDir := filepath.Dir(cfg.DefPaths[0])
	p, err := packages.Load(nil, protoDir)
	if err != nil || len(p) == 0 {
		return errors.Wrap(err, "proto files not found in importable go package")
	}

	cfg.PBPackage = p[0].PkgPath
	cfg.PBPath = protoDir
	log.WithField("PB Package", cfg.PBPackage).Debug()
	log.WithField("PB Path", cfg.PBPath).Debug()

	// GOPATH
	cfg.GoPath = filepath.SplitList(os.Getenv("GOPATH"))
	if len(cfg.GoPath) == 0 {
		cfg.GoPath = filepath.SplitList(build.Default.GOPATH)
	}
	log.WithField("GOPATH", cfg.GoPath).Debug()

	if err := execprotoc.GeneratePBDotGo(cfg.DefPaths, cfg.GoPath, cfg.PBPath); err != nil {
		return errors.Wrap(err, "cannot create .pb.go files")
	}

	_, err = parsesvcname.FromPaths(cfg.GoPath, cfg.DefPaths)
	if err != nil {
		log.Warnf("No valid service is defined; exiting now: %v", err)
		log.Info(".pb.go generation with protoc-gen-go was successful.")
		return errors.Wrap(err, "parsesvcname.FromPaths")
	}

	p, err = packages.Load(nil, cfg.ServicePath)
	if err != nil || len(p) == 0 {
		return errors.Wrap(err, "generated service not found in importable go package")
	}

	log.WithField("Service Packages", p).Debug()
	cfg.ServicePackage = p[0].PkgPath
	// cfg.ServicePath = svcPath

	log.WithField("Service Package", cfg.ServicePackage).Debug()
	log.WithField("package name", p[0].Name).Debug()
	log.WithField("Service Path", cfg.ServicePath).Debug()

	// PrevGen
	cfg.PrevGen, err = readPreviousGeneration(cfg.ServicePath)
	return errors.Wrap(err, "cannot read previously generated files")
}

// parseServiceDefinition returns a svcdef which contains all necessary
// information for generating a truss service.
func parseServiceDefinition(cfg *truss.Config) (*svcdef.Svcdef, error) {
	protoDefPaths := cfg.DefPaths
	// Create the ServicePath so the .pb.go files may be place within it
	if cfg.PrevGen == nil {
		err := os.MkdirAll(cfg.ServicePath, 0o755)
		if err != nil {
			return nil, errors.Wrap(err, "cannot create service directory")
		}
	}

	// Get path names of .pb.go files
	pbgoPaths := []string{}
	for _, p := range protoDefPaths {
		base := filepath.Base(p)
		barename := strings.TrimSuffix(base, filepath.Ext(p))
		pbgp := filepath.Join(cfg.PBPath, barename+".pb.go")
		pbgoPaths = append(pbgoPaths, pbgp)
	}
	pbgoFiles, err := openFiles(pbgoPaths)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open all .pb.go files")
	}

	pbFiles, err := openFiles(protoDefPaths)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open all .proto files")
	}

	// Create the svcdef
	sd, err := svcdef.New(pbgoFiles, pbFiles)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create service definition; did you pass ALL the protobuf files to truss?")
	}

	return sd, nil
}

// generateCode returns a map[string]io.Reader that represents a gokit
// service
func generateCode(cfg *truss.Config, sd *svcdef.Svcdef) (map[string]io.Reader, error) {
	conf := ggkconf.Config{
		PBPackage:     cfg.PBPackage,
		GoPackage:     cfg.ServicePackage,
		PreviousFiles: cfg.PrevGen,
		Version:       version,
		VersionDate:   date,
	}

	genGokitFiles, err := gengokit.GenerateGokit(sd, conf)
	if err != nil {
		return nil, errors.Wrap(err, "cannot generate gokit service")
	}

	return genGokitFiles, nil
}

func openFiles(paths []string) (map[string]io.Reader, error) {
	rv := map[string]io.Reader{}
	for _, p := range paths {
		reader, err := os.Open(p)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot open file %q", p)
		}
		rv[p] = reader
	}
	return rv, nil
}

// writeGenFile writes a file at path to the filesystem
func writeGenFile(file io.Reader, path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return err
	}

	outFile, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "cannot create file %v", path)
	}

	_, err = io.Copy(outFile, file)
	if err != nil {
		return errors.Wrapf(err, "cannot write to %v", path)
	}
	return outFile.Close()
}

// readPreviousGeneration returns a map[string]io.Reader representing the files in serviceDir
func readPreviousGeneration(serviceDir string) (map[string]io.Reader, error) {
	if !fileExists(serviceDir) {
		return nil, nil
	}

	const handlersDirName = "handlers"
	// add svc dir to addFileToFiles walkable path, since move handlers dir inside svc
	const svcDirName = "svc"

	files := make(map[string]io.Reader)

	addFileToFiles := func(path string, info os.FileInfo, errF error) error {
		if info.IsDir() {
			switch info.Name() {
			// Only files within the handlers dir are used to
			// support regeneration.
			// See `gengokit/generator/gen.go:generateResponseFile`
			case filepath.Base(serviceDir), handlersDirName, svcDirName:
				return nil
			default:
				return filepath.SkipDir
			}
		}
		file, ioErr := os.Open(path)
		if ioErr != nil {
			return errors.Wrapf(ioErr, "cannot read file: %v", path)
		}

		// trim the prefix of the path to the proto files from the full path to the file
		relPath, err := filepath.Rel(serviceDir, path)
		if err != nil {
			return err
		}

		// ensure relPath is unix-style, so it matches what we look for later
		relPath = filepath.ToSlash(relPath)

		files[relPath] = file

		return nil
	}

	err := filepath.Walk(serviceDir, addFileToFiles)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot fully walk directory %v", serviceDir)
	}

	return files, nil
}

// fileExists checks if a file at the given path exists. Returns true if the
// file exists, and false if the file does not exist.
func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// promptNoMake prints that truss was not built with make and prompts the user
// asking if they would like for this process to be automated
// returns true if yes, false if not.
func promptNoMake() bool {
	const msg = `
truss was not built using Makefile.
Please run 'make' inside go import path %s.

Do you want to automatically run 'make' and rerun command:

	$ `
	fmt.Printf(msg, trussImportPath)
	for _, a := range os.Args {
		fmt.Print(a)
		fmt.Print(" ")
	}
	const q = `

? [Y/n] `
	fmt.Print(q)

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(strings.TrimSpace(response)) {
	case "y", "yes":
		return true
	}
	return false
}

const trussImportPath = "github.com/ben1009/truss"

// makeAndRunTruss installs truss by running make in trussImportPath.
// It then passes through args to newly installed truss.
func makeAndRunTruss(args []string) error {
	p, err := build.Default.Import(trussImportPath, "", build.FindOnly)
	if err != nil {
		return errors.Wrap(err, "could not find truss directory")
	}
	make := exec.Command("make")
	make.Dir = p.Dir
	err = make.Run()
	if err != nil {
		return errors.Wrap(err, "could not run make in truss directory")
	}
	truss := exec.Command("truss", args[1:]...)
	truss.Stdin, truss.Stdout, truss.Stderr = os.Stdin, os.Stdout, os.Stderr
	return truss.Run()
}
