package getstarted

import (
	"bytes"
	"os"
	"path"
	"strings"
	"text/template"

	gogen "github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type protoInfo struct {
	alias string
}

func (p protoInfo) FileName() string {
	return p.PackageName() + ".proto"
}

func (p protoInfo) PackageName() string {
	a := p.alias
	a = strings.Replace(a, "-", "", -1)
	a = strings.Replace(a, " ", "", -1)

	a = strings.ToLower(a)
	return a
}

func (p protoInfo) ServiceName() string {
	a := p.alias
	a = strings.Replace(a, "-", "_", -1)
	a = strings.Replace(a, " ", "_", -1)
	return gogen.CamelCase(a)
}

// Gen create a proto in the targetPath path with the given serviceName
func Gen(pkg, targetPath string) error {
	pkg = removeDotProtoSuffix(pkg)
	pinfo := protoInfo{
		alias: pkg,
	}
	file := path.Join(targetPath, pinfo.FileName())
	log.WithField("proto file Path", file).Debug()
	f, err := os.Create(file)
	if err != nil {
		log.Error(errors.Wrapf(err, "cannot create %q", pinfo.FileName()))
		return errors.Wrap(err, "Gen.Create")
	}

	code, err := renderTemplate(pinfo.FileName(), starterProto, pinfo)
	if err != nil {
		log.Error(err)
		return errors.Wrap(err, "Gen.renderTemplate")
	}

	_, err = f.Write(code)
	if err != nil {
		log.Error(errors.Wrapf(err, "cannot write default contents to %q", pinfo.FileName()))
		return errors.Wrap(err, "Gen.Write")
	}

	return nil
}

// removeDotProtoSuffix exists to preempt and warn a user who enters a name
// containing `.proto`. It will warn the user of their incorrect input and will
// demonstrate how their input can be corrected. Then, the program continues
// using the corrected input it warned about.
func removeDotProtoSuffix(pkg string) string {
	const dotProtoInName = `The name you provided has a suffix of '.proto' when it should not. Instead of
'{{.Got}}', you should provide '{{.Want}}'. Here's an example of the correct
command to enter next time:

	truss --getstarted {{.Want}}

For now this program is continuing as though you used '{{.Want}}'.
`
	want := strings.Replace(pkg, ".proto", "", -1)
	if strings.HasSuffix(pkg, ".proto") {
		executor := struct{ Got, Want string }{pkg, want}
		warn, err := renderTemplate("dotProtoInName", dotProtoInName, executor)
		if err != nil {
			log.Error(err)
		}
		log.Warn(string(warn))
	}
	return want
}

func renderTemplate(name string, tmpl string, executor interface{}) ([]byte, error) {
	codeTemplate := template.Must(template.New(name).Parse(tmpl))

	code := bytes.NewBuffer(nil)
	err := codeTemplate.Execute(code, executor)
	if err != nil {
		return nil, errors.Wrapf(err, "attempting to execute template %q", name)
	}
	return code.Bytes(), nil
}
