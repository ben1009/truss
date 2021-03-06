// Code generated by go-bindata.
// sources:
// NAMEservice/client/grpc/client.gotemplate
// NAMEservice/client/http/client.gotemplate
// NAMEservice/cmd/server/main.gotemplate
// NAMEservice/svc/endpoints.gotemplate
// NAMEservice/svc/handlers/handlers.gotemplate
// NAMEservice/svc/handlers/hooks.gotemplate
// NAMEservice/svc/serve.gotemplate
// NAMEservice/svc/transport_grpc.gotemplate
// NAMEservice/svc/transport_http.gotemplate
// DO NOT EDIT!

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clientGrpcClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xcf\x6f\xdb\x36\x14\x3e\x8b\x7f\xc5\x9b\x11\x0c\x52\xa0\x50\xf7\x0e\xb9\xd4\xc9\x8a\x0e\x6b\x1a\xa4\xc1\x76\x28\x8a\x82\xa6\x9e\x64\xc2\x32\xa9\x92\xb4\x13\x43\xd0\xff\x3e\x3c\x52\x72\xe4\xc4\x71\x73\x19\x76\x59\x81\x26\x22\xf9\xf1\xfd\xf8\xbe\x47\xf2\xa5\x28\x60\x6e\x4a\x84\x1a\x35\x5a\xe1\xb1\x84\xc5\x0e\xbc\xdd\x38\xc7\xe1\xea\x33\xdc\x7c\xbe\x87\xeb\xab\x8f\xf7\x9c\x15\x05\xdc\xa1\xdd\x68\xad\x74\x1d\x01\xf0\xa0\x9a\x06\xcc\x16\xed\x83\x55\x1e\xc1\x2f\x95\x83\x4a\x35\xc8\x19\xa1\x6f\x85\x5c\x89\x1a\xa1\xb6\xad\x84\xd6\x9a\xad\x2a\xd1\x81\x80\xfa\xee\x76\x0e\xb2\x51\xa8\x3d\x54\xc6\x82\x5f\x22\x74\x1d\xff\x82\x76\xab\x24\xf2\x1b\xb1\xc6\xbe\x07\x37\x0c\x59\x3b\x31\xc3\x98\x5a\xb7\xc6\x7a\x48\x59\x32\x93\x46\x7b\x7c\xf4\x33\x96\xcc\x6a\x63\xea\x06\x79\x6d\x1a\xa1\x6b\x6e\x6c\x5d\x10\xfa\xf5\x95\x62\x8d\x5e\x94\xc2\x8b\x00\x51\x7e\xb9\x59\x70\x69\xd6\x45\xbb\xaa\x0b\xb4\xd6\x58\x37\x63\x87\x2b\xb5\xb9\x58\x29\x5f\xd0\x7f\xd4\x65\x6b\x94\x26\xc7\x64\xcb\x5b\xa1\x5d\x08\xea\x15\xfc\x1e\xf0\x14\xd4\x14\x57\x9b\xa2\xb5\xc6\x9b\xc5\xa6\x2a\xfc\xae\xc5\xe0\xba\x28\xe0\x9e\xd8\x1c\x48\x61\xc9\xac\xeb\xf8\xc7\x90\xfb\xad\xf0\x4b\xb8\xe8\x7b\x28\xdc\x96\xac\xb5\x0b\xa0\xc5\xdb\xf7\x87\xcb\x33\x96\x05\x15\x6e\xf0\x01\x2c\xfa\x8d\xd5\x0e\x84\x1e\x69\x85\x85\x90\xab\xa8\xf5\xa1\x20\xd2\x68\x8d\xd2\x2b\xa3\x39\x7c\xf4\xa0\x1c\xc9\x43\x76\x2c\xba\xd6\x68\xa7\x16\xaa\x51\x7e\x07\xa6\x0a\xba\x49\xd1\x34\x68\xc1\x1b\x28\x95\x68\x72\x10\xba\x84\x46\x78\xb4\x20\x1b\xe3\x30\x8f\xa0\x27\x9b\xac\xda\x68\x49\x31\xa5\x34\x09\xe7\xc4\x08\x9f\x07\xd7\x73\xa3\x75\x0e\xa6\x25\x9c\x03\xce\x87\xe9\xcf\x61\x22\x83\xb4\x5d\xf0\x17\x55\x42\x23\xb4\x39\x04\xcd\x32\xe8\x58\xb2\x15\x16\xa4\x1c\xb2\x99\x1b\x5d\xa9\x9a\xb1\x84\xca\xec\x7b\x0e\x15\xbc\xbb\x04\x2b\x74\x8d\x7b\x3f\x1d\x4b\x12\xb4\x96\x16\xaa\xf4\x57\x29\x33\x96\x24\xaa\x22\x83\xf0\xcb\x25\x68\xd5\x04\x44\x12\x19\xa4\xf1\xe0\xcc\xf1\xbf\xad\x68\x53\xb4\x36\x87\x99\x14\x5a\x1b\x0f\xa2\x6d\x9b\xdd\x60\x79\x46\x86\x7a\x96\xf4\x8c\x25\x72\x92\x88\x23\x4f\x5f\xbf\x1d\x14\xce\x41\xa6\xe4\xee\xd8\xea\x7b\xac\x8c\xc5\x94\x82\x19\x0a\xff\x2f\xd1\x6c\xd0\xdd\x9b\x0f\x77\xb7\xf3\x4f\x43\x3d\xa7\x52\xf2\x25\x8a\x12\xad\xcb\xb2\x9c\xdc\x27\x5d\x77\x01\x0f\xca\x2f\xe1\xcc\x23\x39\xe7\x7d\xcf\x92\xc9\x6c\xbb\xaa\x89\x4c\x5a\x3a\xf3\xc8\x87\x53\x1b\xf9\x25\x6f\x84\x8c\x9c\x9d\xa9\x11\x34\xaa\xf0\x09\xfd\xd2\x94\xae\xef\x19\x00\x40\x92\xd0\xcf\xae\x53\x15\x68\x04\x7e\x17\x2b\x06\xef\x77\x6d\xd4\x0b\x66\xd7\xeb\xd6\xef\x66\x10\x0d\x07\xad\xba\xee\xde\xfc\x69\x1e\xd0\xc2\x99\x1a\x44\xbd\x1e\xce\x17\x8c\x07\x8d\x8f\x33\x61\x57\xd0\x83\xc2\x7a\x7d\xe3\x25\x1c\x32\x78\x83\x0f\x91\xc4\x34\xee\x25\x06\x75\x3e\x7c\xcf\xba\x6e\xe4\xa0\xef\x79\xd7\x4d\xf3\x8b\x93\xb3\x29\x54\x3d\x9f\xbc\xd6\xd2\x94\x48\x22\x4c\x56\xef\xf0\xc7\x06\x9d\x1f\x31\x57\x78\x14\x13\xf9\x19\x41\xa1\xc0\x3f\x98\xc0\xd4\x99\x7a\x49\x5f\xdf\x77\xfd\x88\x3d\x28\x29\xce\xf9\x30\x9f\xed\xa9\x4a\xb3\x30\x13\x89\xee\x3a\x6c\x1c\x0e\x2a\x4d\xff\xfd\xaf\xc0\x81\x02\xe1\xfa\xfd\x17\x45\xd0\xe5\x78\xa4\xc6\xcf\xfd\xd7\xf8\xc1\xc6\xab\xc6\x6d\xe5\xde\x8e\xeb\x08\x30\x3d\x84\xcf\x4f\x20\xdd\xf8\xc1\xdc\x0b\x29\xde\xc5\x33\xf9\xaa\x54\xf9\x93\xef\xa4\xcf\xe9\x86\x63\x7d\x78\x38\x88\x2b\x88\xa2\x41\x64\x8f\x9d\x8e\xa1\xef\x69\xdb\x49\xa2\xe9\x3d\x11\x70\xf8\x20\xf2\xb8\x63\x84\xfc\x4e\x0f\x84\x5f\x8a\xf0\x14\x6d\xd1\x7a\x07\x82\xec\x86\x47\xea\x48\x1e\x60\x91\x6e\x5d\x6f\x40\xc0\xc6\xa1\xbd\x28\xcd\x5a\x28\xfd\x0a\x34\xfa\xe0\x70\x6b\xd5\x5a\x58\xd5\xec\x68\x4f\xb5\x69\x40\x69\x10\xc3\xab\xc1\xd9\xdb\xaf\x31\x42\x86\x27\xed\x64\xda\xe9\x77\x18\xee\x6c\x3e\x8f\xbf\xf3\x70\x3e\xee\x42\xe8\x4a\x7b\xb4\x95\x90\xd8\xf5\x19\xa4\x93\xd1\xe4\x5d\x1b\x0f\x6c\x4c\xf6\xdd\xe5\xd3\x76\x9e\x9e\xff\xfc\xe2\xc8\x26\x06\x42\x75\x05\x3b\x51\x6d\x9a\xed\x87\x94\x27\xd7\xc4\x7f\x9e\xd5\x9b\x0e\xe3\x9b\x12\xdb\x17\xf8\xfe\x90\x3d\x2b\xef\x78\x81\xbc\xa9\xbc\x4f\xdd\x35\x47\xab\x3b\x6e\x18\x10\xaf\x15\xf7\xcf\x0b\x37\x3a\x08\x55\x7e\xe2\x24\x04\xd4\xcf\xaa\x3b\x28\x7b\x2a\x8f\x63\xc2\x8e\x11\xbc\x45\xd6\xc4\xe2\x8f\xd0\x69\x0d\xf1\x1c\xa9\xd0\xb0\x70\xa0\x63\xb2\x17\xf0\xc7\x78\x0b\xed\xe5\x62\x54\x0b\x07\x3d\x1d\x38\x6f\x37\xd2\x93\xb3\xa1\xdd\x81\xaf\xdf\x9c\xb7\x4a\xd7\xc3\xf5\x35\xed\xa9\xa2\x30\x94\x77\x18\x05\x01\xd6\xa6\x54\x95\x42\x17\x3b\xd4\x7d\xf3\x4b\xfd\x62\xf0\x76\xb0\x9f\xb6\xa6\xe7\xd3\x00\xb2\x98\x2e\x8b\x6c\xce\xfd\xe3\xd8\x8d\x7d\x41\x5d\xa6\x2b\xdc\x85\x16\x36\x46\x94\x1d\x1a\xeb\xf6\xb9\x06\xb3\x06\x8e\x19\x0e\x6d\xa7\x19\x7b\x39\xb8\x04\x32\xc9\xa6\x8d\x28\x35\x77\xfd\xe0\xff\x54\x47\x18\x62\x19\xc9\xc9\xe0\x58\x6f\x39\xad\xce\x67\xd1\x49\xff\xf8\xb2\x18\xd6\x25\x9c\x8f\x7f\x41\xf1\x4f\x57\xd9\x73\x44\x08\x9e\x7a\x8b\x56\xa8\xa9\x32\xc9\xd8\x88\xaf\x9e\x1a\xf1\x10\x5e\xe8\x28\x54\x05\xdb\x1c\x4c\x58\x93\xfe\x91\x87\x6c\xd2\x55\xc6\xd3\x21\xf6\xdf\x68\x31\x36\x1f\xd1\xf0\x25\xb5\xdc\xc4\x77\x18\xe6\xb0\xca\x61\x1b\x9e\xdc\x3e\x34\xdf\xb1\x95\x8f\xd0\x69\x33\x7f\xbe\x2e\xe1\x12\xf6\x09\xfc\x61\x94\x4e\xcf\xd7\x65\xfe\x34\x75\x4b\x7b\xa2\x55\xce\x79\x96\x8d\xe6\x06\x66\xa4\x7f\x8c\xec\xff\x13\x00\x00\xff\xff\x89\x3d\xda\xec\x3d\x0f\x00\x00")

func clientGrpcClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_clientGrpcClientGotemplate,
		"client/grpc/client.gotemplate",
	)
}

func clientGrpcClientGotemplate() (*asset, error) {
	bytes, err := clientGrpcClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/grpc/client.gotemplate", size: 3901, mode: os.FileMode(420), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientHttpClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x10\x34\x29\x91\xc8\x02\x56\xf8\x98\x88\xc3\x67\x9d\x3f\x95\xe5\xdd\x69\x18\xe0\x8d\xa1\xab\xdc\x01\x29\x7e\x61\x7c\x7b\xd7\x82\x5a\xfc\x7d\x52\x5f\x64\x63\xe4\xda\x9b\x07\x95\xf8\x34\xcb\x44\x2a\x2e\x4f\x0f\x39\xfc\x01\x59\x75\xce\x65\x8c\x23\x9b\x49\xda\xf6\xfd\xb6\xc1\x1a\x22\x5d\xed\x44\xe5\xfe\x27\x92\xe6\x5c\x7e\x01\x00\x00\xff\xff\x0b\x3c\x4c\x9e\x69\x00\x00\x00")

func clientHttpClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_clientHttpClientGotemplate,
		"client/http/client.gotemplate",
	)
}

func clientHttpClientGotemplate() (*asset, error) {
	bytes, err := clientHttpClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/http/client.gotemplate", size: 105, mode: os.FileMode(420), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdServerMainGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8f\x41\x6e\x83\x30\x10\x45\xd7\xcc\x29\xbe\x58\xc1\xa2\xf6\x21\x92\x2e\xb2\x69\xa2\x96\x1e\xc0\x81\x6f\xb0\x4a\x0c\xb2\x0d\x51\x84\xb8\x7b\x65\x2a\x75\x35\xb3\x78\x33\xff\x7d\xad\x71\x9a\x3a\xa2\xa7\x67\x30\x89\x1d\xee\x2f\xa4\xb0\xc4\xa8\x70\xbe\xe2\xe3\xda\xe0\xfd\x7c\x69\x94\x68\x8d\x4f\x86\xc5\x7b\xe7\xfb\x3f\x00\x4f\x37\x8e\x98\x56\x86\x67\x70\x89\x48\x83\x8b\xb0\x6e\xa4\x12\x99\x4d\xfb\x63\x7a\xe2\x61\x9c\x17\x71\x8f\x79\x0a\x09\x95\x14\xa5\x1d\x4d\x5f\x8a\x14\x5a\xa3\xc9\xfc\x17\xc3\xea\x5a\x4a\x51\x6e\x9b\xba\x1c\xdc\xcd\xa4\x01\x6f\xfb\x0e\x1d\xd7\xb6\x94\x5a\xc4\x2e\xbe\x3d\x5e\x55\x35\xb6\xe3\xf6\x7b\xee\x4c\x22\x4c\xd7\x05\xc6\xc8\x08\x67\x91\x06\xbe\x30\x98\x95\xb8\x93\xfe\x5f\x2c\xd1\xe7\x4e\x39\x38\x4a\x91\x87\xba\x99\x10\x59\xd5\x22\x45\x5c\x5b\x95\x15\x58\xe5\xed\x4c\x6b\x96\x31\x9d\x26\x6f\x5d\x5f\xcb\x2e\xbf\x01\x00\x00\xff\xff\xe2\x8f\xd5\x0c\x1e\x01\x00\x00")

func cmdServerMainGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_cmdServerMainGotemplate,
		"cmd/server/main.gotemplate",
	)
}

func cmdServerMainGotemplate() (*asset, error) {
	bytes, err := cmdServerMainGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/server/main.gotemplate", size: 286, mode: os.FileMode(420), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _svcEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x4b\x6f\xe4\xb8\x11\x3e\x4b\xbf\xa2\x56\x70\xe0\xee\x85\xcc\xbe\x7b\xe1\x43\x32\xeb\x24\x06\x32\x9e\xc1\xd8\x41\x0e\x8b\xc5\x80\x2d\x95\xba\x09\x53\x24\x87\xa4\xfa\x11\x41\xff\x3d\x28\x52\x52\xab\xdd\xf2\x3c\x02\xe4\x12\xec\xc1\x90\x9b\x8f\x8f\x5f\x7d\xf5\x60\x71\xb5\x82\x77\xba\x44\xd8\xa0\x42\xcb\x3d\x96\xb0\x3e\x82\xb7\x8d\x73\x0c\x7e\xfd\x00\x8f\x1f\x9e\xe1\xfe\xd7\x87\x67\x96\xae\x56\xf0\x09\x6d\xa3\x94\x50\x9b\xb8\x00\xf6\x42\x4a\xd0\x3b\xb4\x7b\x2b\x3c\x82\xdf\x0a\x07\x95\x90\xc8\xd2\xd4\xf0\xe2\x85\x6f\x10\xdc\xae\x48\x69\xeb\xf3\x30\x07\x85\x56\x9e\x0b\xe5\xa0\x46\xbf\xd5\xa5\x03\xaf\xa1\xe6\x2f\x08\x42\x95\x62\x27\xca\x86\x4b\x40\x55\x1a\x2d\x94\x77\x50\x59\x5d\x83\x43\xbb\x13\x05\xba\x9c\x90\x2c\x7e\x69\xd0\x79\xe0\xaa\x04\x8b\xce\x68\xe5\x10\xfc\xd1\x60\x40\xa2\xa5\xc4\x44\x3b\x3c\xa1\xe4\xc0\x1d\xec\x51\x4a\xfa\xa2\x2a\x74\x89\xd6\x11\x00\xe1\x95\xd8\xff\xae\xb4\xed\x37\x06\xb4\x3c\x0c\x70\xb2\xb0\x02\xdd\x58\x70\x8d\x31\xda\x92\x42\xde\x72\xe5\xe8\x7f\x3a\x4e\x70\x29\xfe\xcd\xbd\xd0\x8a\xd0\x2a\x6d\x6b\xee\x1d\x4b\x53\x51\x87\x15\x8b\x34\xc9\xaa\xda\x67\x69\x92\x91\xe5\x78\xf0\x59\x9a\x26\xd9\x46\xf8\x6d\xb3\x66\x85\xae\x57\x1b\x7d\xf3\x22\xfc\x8a\xfe\x06\xc6\xd9\xeb\x15\x1b\xbd\x32\x56\x7b\xbd\x6e\xaa\x55\x60\x47\x20\x66\x0d\x59\xdb\xb2\x8f\x7f\x79\x08\x47\x7d\xe4\x7e\x0b\x37\x5d\x97\xa5\xcb\x20\xf9\xfd\x28\x62\xa1\xa5\xc4\xc2\xbb\xc1\x1a\xbf\x9d\x88\x03\x7e\xcb\x3d\x14\xba\x36\x64\x3a\x57\xc0\xcb\x72\x50\x9c\xc1\x83\xbf\x76\x04\x56\x23\x57\x9e\x04\x5e\x23\x34\x0e\x4b\x52\x92\xc3\x16\xa5\x41\x0b\xce\xdb\xa6\xf0\x39\x4d\xf7\x47\xcd\x9f\x24\x94\xd7\xc0\x09\xce\x09\xb5\x91\x08\x86\x5b\x5e\xa3\x47\x4b\xe1\x45\xe3\x0f\x0a\x78\xf4\xa1\xcd\x41\xf8\x6b\x47\x87\x55\x8d\x0c\xbe\xa8\x1a\x55\x90\xce\x3d\x65\x85\xe4\x0a\x0d\xda\x84\xc0\x05\x4d\x7b\x0d\xda\x9b\xe1\x40\x02\x5c\x73\x27\x1c\x83\xbf\x6a\x0b\x78\xe0\xb5\x91\x98\xc3\x51\x37\x50\x8b\xcd\xd6\x83\xe1\x8e\xe2\x60\x22\x15\x11\x1c\x0f\x8a\xe7\x18\xab\xcb\xa6\xc0\x20\x03\x57\xb0\xf5\xde\xb0\xbf\x73\x55\x4a\xe2\xb8\x17\x7e\x0b\xc8\x8b\x6d\x1f\xce\xb0\x18\x4e\x5f\xc2\x5e\x58\x2c\xa1\x31\x11\xd4\x19\x2c\x44\x25\x0a\x30\xdc\x6f\x19\x2c\x1e\x02\x3f\xe1\x08\x7f\xcd\xd7\xf2\x08\x1c\x6a\xe1\x7c\x4c\x05\x28\xd1\x89\x8d\xa2\xad\x42\xed\xf4\x0b\x06\x29\x9f\xa2\x5b\xc6\xd4\x09\x14\xf1\xdc\xd9\xd1\x19\x04\x31\x28\xc9\x96\x53\x75\x0b\x29\x50\xf9\x73\x75\x27\x8e\x3b\x65\xa1\x3c\x52\xae\x46\x38\x2c\xbf\xe6\x46\x8a\xc8\xa8\x95\x20\x85\x6b\x8c\x61\x75\xe2\x2b\x94\x47\x5b\x71\x0a\xa8\x79\x4f\x10\xd8\x78\xd8\x7c\x25\x68\x5c\x2c\x3c\x7d\xea\xad\x82\x1f\x1e\x71\xff\xae\xb7\xa7\xd0\xf5\x5a\xa8\xa0\x53\xdd\x53\x9c\x38\x36\xef\xeb\x85\x6f\xac\x02\x11\x22\x99\x08\x16\x5c\x4a\xb4\x31\x98\x7b\xb2\x2c\x0d\xe6\x5c\x08\xda\xa6\x6d\x6b\xb9\xda\x20\x5c\x09\xb8\xbd\x03\x36\xac\x7f\x1f\x9d\xd1\x75\x69\xd2\xb6\x57\x82\x3d\xf2\x1a\xbb\x6e\xd8\x0f\x00\xa3\x11\x6c\x18\x4c\xdb\xf6\x86\x46\xbb\x2e\xed\xd2\x94\xc2\x0d\x1e\x71\x3f\x1e\xb9\xe8\xd3\x0f\xcc\x9a\xb5\xed\x78\x50\x04\x7e\x0a\x3e\x5d\x4e\x08\xb6\x69\xd2\x1b\x36\x8e\xb5\xc4\xe5\x6b\x6c\xa9\x50\xa4\xc9\x1c\xe1\x5b\x62\x4c\x05\x79\x66\x6e\x20\xb6\xcc\x09\x3f\x1a\x90\x04\x1b\xa6\x21\xf8\x6d\xa1\xe8\x88\xb6\x15\x15\x28\x04\xf6\xa9\x2f\xe2\xcf\x47\x13\x6d\x84\xec\xbe\x36\xfe\x98\x01\xa1\x07\x71\x16\x13\x7f\x2c\x61\x42\x6c\x51\xf8\x03\xf4\x75\x95\xbd\x8b\xdf\x9c\x62\xff\xe7\x20\xdd\xdf\x74\xc0\xbb\x12\xec\x53\xbc\x35\xc6\x33\xba\x6e\x09\x8b\xcb\x45\xaf\x98\x74\x5d\x0e\x68\xad\xb6\x4b\x12\x99\x58\x0f\x37\x4e\x18\x27\xeb\x90\xcd\xe9\x54\xf8\x03\xd1\x58\x06\x4b\x13\x51\x85\xd5\x3f\xdd\x81\x12\x12\xda\x38\x38\xf8\x4c\x09\x19\xc0\xe2\x68\xd4\x66\x98\x1b\x4e\x63\xdf\x43\x75\x99\x13\x56\x9a\x0c\xf2\xa2\x74\xf8\x3f\x57\x30\x5c\x46\x3f\x22\x62\xf2\xa3\x0a\x26\x17\xf2\x25\x97\xda\x25\x49\x97\x26\x33\xaa\x7d\x17\xbd\x0b\xe1\x42\x5c\x0f\x5f\x0a\xed\xf7\x54\x94\xa7\xf1\x1d\xca\xfe\x95\xc7\x10\xdf\x31\xf3\xa7\x21\x7f\xe5\x71\xae\x3c\x44\x57\xbc\x99\x5b\x31\xdd\xa7\x7b\xcf\x33\xfe\xa2\x8c\x9c\x69\x41\xd8\xf3\xce\x1c\x3a\xa6\xb1\x0a\xb7\xe4\xba\xb1\x77\x9a\x0c\x47\x9f\x4c\x9c\x45\xe8\x5f\xc8\xa2\x1e\x63\x2e\x10\x2f\xc2\x22\xec\xdb\x8d\xfe\x75\xec\x55\xb8\x05\x46\x71\xd5\x8c\x6b\xe7\x9c\x1b\xdd\x3b\xce\xec\x7a\x87\xc5\xe1\xee\x54\x8b\xa6\x3e\xfb\x97\xe5\xe6\xcf\x52\xde\x1f\x0a\x34\x1e\xf6\x96\x1b\x17\x2f\xea\x51\xbd\x4a\xa0\x2c\xa9\x4b\xe9\x2b\xfc\xa9\xa2\x06\xf7\x86\x1b\x6e\xa6\x39\x63\xef\x45\x59\x4a\xdc\x73\x8b\xa1\x2b\xfe\xa7\x1b\x5a\x5f\xea\x17\x8d\x91\x47\xba\xa8\xe8\xf2\xf5\x04\x5e\x8f\xab\x43\x77\x81\x3b\xb4\xc7\xd1\x95\x94\x68\x74\x0f\x0d\xfd\x16\xe1\x7d\x30\xd4\x7b\xd0\xfd\x9b\x4f\xae\xbf\x82\x2b\xea\xbd\xa8\x63\xc1\x92\xb6\xad\x8f\xa0\xc8\x07\xb1\x27\xc3\x43\x21\x9b\x12\xcb\xd8\x30\xaf\x91\x28\x90\xcd\x06\x4b\x76\xa1\xc6\xe2\xc4\x29\x87\xec\xc9\x73\xdf\xb8\x2c\x87\xec\xa3\x50\x9b\x6c\x99\x0e\x05\xe3\xe7\x49\xc5\x78\x6b\x3f\xcc\xa8\x92\x9f\xd8\x30\xc6\x9c\xb7\x42\x6d\x42\x38\x09\xd5\x0f\xdf\xde\x41\xcd\xcd\x6f\x71\xea\xf7\x28\x7f\xdb\xb5\xe1\x42\xba\x81\x6f\xdd\xb2\x49\x92\x4d\x22\x2a\xbb\x85\xb6\xcb\xfb\xad\xe3\x95\x94\x26\xe4\x8d\xcf\x44\x25\x84\x6f\x80\x1c\x69\xb5\xb1\xaa\x7c\xce\x41\xbf\xd0\xf4\x40\xec\x37\x3c\xfc\xfe\x0b\xfc\xa4\x5f\x62\x28\x1a\xae\x44\xb1\xa8\x6a\xcf\x9e\x8c\x15\xca\x57\x8b\xec\x7e\x80\x18\x3d\x78\xfd\x27\x77\x0d\xa5\x46\x07\x4a\x7b\xc0\x83\x70\xfe\x17\x70\x88\x53\xc7\x8f\xb1\xe3\xd8\x46\x67\x44\x6a\xb9\xec\x6b\x56\x89\x12\x3d\x2e\x06\x06\x61\xee\x64\x80\x50\xc5\x89\xfe\x28\xdf\xf7\x0b\x25\xaa\x00\x71\x77\x07\x67\x92\xf5\x99\x36\x5b\x79\xe1\x6e\xc2\x7c\x31\xbb\x64\x39\xa4\xde\x99\xe4\x31\xed\xfe\xc1\xd7\x28\xb1\x3c\x45\x43\x7c\x25\x6e\xd0\x0f\xb1\x3b\x7d\x13\xc4\x10\xde\x6f\x51\x8d\xb3\x7a\x12\xae\x3d\x58\x8c\xba\x3c\x66\x59\x9f\x08\x4d\x5c\x0c\xf1\xe9\xc9\xe3\xfb\x55\x14\xd4\x1a\x5b\x51\xc4\x37\xcb\x84\xc3\x56\x14\xdb\xb0\xd5\xa1\x9a\xa3\xd0\xf7\x83\xfd\xee\xa1\x1b\xd6\xb6\xef\x06\x2f\xad\x0a\xe5\x36\x06\x70\x7e\x59\x99\x67\x8a\x75\xfa\x96\x5d\xff\x75\x6d\xba\x20\x95\xf7\x76\x06\xc5\x2d\x16\x28\x76\xf1\xdd\x10\x4c\x7c\xf5\x1c\x63\xf0\x84\x38\xc2\x4c\x50\xc2\xc4\xf0\x9c\x19\xf3\x9e\x88\x52\x44\x96\xe8\xb9\x90\xe1\xe9\x31\xa4\x53\x78\xf7\xf6\x4f\x26\x2e\x85\x3f\xb2\xaf\x95\x90\x33\xdb\xa7\x95\xe4\x87\x15\xfd\xa3\xce\xfc\xff\xd4\x99\xb3\x6d\xf9\x7c\x4f\xf8\x56\xd9\xf9\x4f\x00\x00\x00\xff\xff\x75\x40\x80\x65\xc3\x12\x00\x00")

func svcEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcEndpointsGotemplate,
		"svc/endpoints.gotemplate",
	)
}

func svcEndpointsGotemplate() (*asset, error) {
	bytes, err := svcEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/endpoints.gotemplate", size: 4803, mode: os.FileMode(420), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _svcHandlersHandlersGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\x80\x20\x0c\x05\xd0\xbb\x53\xf4\x4c\xa2\x1d\xc6\x09\x88\x7c\xab\x11\xa9\x69\xeb\x89\xb0\x3b\xaf\x77\x4e\xb4\x03\x24\xba\x86\xfd\xee\x2c\x68\xa2\xcf\x1d\x7c\xe5\x56\x2a\xcc\x39\xf0\x7e\x35\x07\x7c\x13\xa5\x53\x8d\x0e\x2d\xa0\xc4\x63\x2c\x33\x00\x00\xff\xff\x52\xf7\x77\x02\x3f\x00\x00\x00")

func svcHandlersHandlersGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcHandlersHandlersGotemplate,
		"svc/handlers/handlers.gotemplate",
	)
}

func svcHandlersHandlersGotemplate() (*asset, error) {
	bytes, err := svcHandlersHandlersGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/handlers/handlers.gotemplate", size: 63, mode: os.FileMode(420), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _svcHandlersHooksGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\x80\x20\x0c\x05\xd0\xbb\x53\xf4\x4c\xa2\x1d\xc6\x09\x88\x7c\xab\x11\xa9\x69\xeb\x89\xb0\x3b\xaf\x77\x4e\xb4\x03\x24\xba\x86\xfd\xee\x2c\x68\xa2\xcf\x1d\x7c\xe5\x56\x2a\xcc\x39\xf0\x7e\x35\x07\x7c\x13\xa5\x53\x8d\x0e\x2d\xa0\xc4\x63\x2c\x33\x00\x00\xff\xff\x52\xf7\x77\x02\x3f\x00\x00\x00")

func svcHandlersHooksGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcHandlersHooksGotemplate,
		"svc/handlers/hooks.gotemplate",
	)
}

func svcHandlersHooksGotemplate() (*asset, error) {
	bytes, err := svcHandlersHooksGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/handlers/hooks.gotemplate", size: 63, mode: os.FileMode(420), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _svcServeGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\xdf\x6f\xdb\x36\x10\x7e\x16\xff\x8a\x2b\x81\x0d\x32\xe0\x52\x05\xb6\xee\xa1\x6b\x1e\xda\x38\x6d\x03\xac\x89\xe1\xb8\x7b\x1d\x68\xe9\x24\x13\x91\x48\xed\x48\xcb\x09\x0c\xff\xef\xc3\x51\x92\xe3\x74\x89\x97\xee\x45\xa2\xc8\xef\xbe\xfb\x78\xbf\x94\x65\x70\xee\x0a\x84\x0a\x2d\x92\x0e\x58\xc0\xea\x1e\x02\x6d\xbc\x57\x30\xbb\x86\xab\xeb\x25\x5c\xcc\x2e\x97\x4a\x64\x19\x2c\x90\x36\xd6\x1a\x5b\xf5\x00\xd8\x9a\xba\x06\xd7\x21\x6d\xc9\x04\x84\xb0\x36\x1e\x4a\x53\xa3\x12\xa2\xd5\xf9\xad\xae\x10\x7c\x97\x0b\x61\x9a\xd6\x51\x80\x54\x00\x00\xc8\xb2\xd6\x95\xec\x97\xce\x0f\x8b\xb2\x09\x52\x24\xb2\x76\x15\xbf\x2c\x86\xe1\x95\xad\x43\x68\x8f\xd7\x59\xdb\x92\x2b\xa5\x10\x49\x96\xc1\x2f\x05\xcc\x35\x85\x7b\x91\xc8\xca\xb9\xaa\x46\x55\xb9\x5a\xdb\x4a\x39\xaa\xb2\x8a\xda\x7c\xc0\x2d\x59\xd8\x0d\x52\x67\x72\x14\x49\xbb\x02\xb9\xdb\xa9\xf9\xc7\xcb\x28\x6b\xae\xc3\x1a\x5e\xef\xf7\xec\x65\xb7\x53\x8f\x37\x21\xf3\x5d\x9e\xad\xb5\x2d\x6a\x24\x2f\xc5\x44\x88\x4e\x13\xcc\xb0\xd4\x9b\x3a\x9c\x3b\x5b\x9a\x0a\xfa\x97\x10\xe5\xc6\xe6\x60\xac\x09\xe9\x04\x76\x22\xe1\x7b\xaa\x9b\x40\xc6\x56\x7f\x6a\x4a\x7f\x7e\x64\xa4\x66\xb8\xda\x54\x1f\x8a\x82\xa6\x20\x0b\x5e\x2b\x5d\x14\x24\xa7\x20\xdf\xbd\x7d\xf3\xdb\x1b\x5e\x44\x08\x68\x5b\x40\x83\x81\x4c\xee\xa1\x36\x3e\xa0\x05\x46\xa2\xf7\x72\xf2\x5f\x4e\xbe\x2c\x97\xf3\xc1\x07\x07\xef\xd8\xc5\xdb\xe8\x82\x01\x3f\xcc\xfa\x79\x31\x3f\x1f\x58\x39\xc8\xc7\xac\xbf\x46\xd6\x6a\x31\x3f\x87\x94\xb9\x27\xff\x26\x8f\x19\xf9\xe6\x11\xd0\x76\x86\x9c\x6d\xd0\x06\xe8\x34\x19\xbd\xaa\xd1\x4f\xc1\x94\xe0\x31\x28\xf8\x54\xeb\xca\xc3\x5a\x77\x08\x2d\x19\x47\x26\xdc\xc7\x5a\x83\x0b\xdb\x31\xde\x2b\x91\x98\x32\x12\xc3\xbb\x33\x70\x5e\x7d\xc6\x80\xb6\x4b\xe5\xec\xe2\xe3\xb7\xcf\x7f\x7d\x98\xcd\x16\x72\xf2\x7b\x0f\x78\x75\x06\x52\x72\x52\x92\x67\xb2\x00\x67\x11\x28\x92\x7d\x64\x8d\xf5\xfa\x98\x75\x7e\xbd\x58\x32\x5f\x3c\x7a\x8e\x6f\x0c\x38\x9c\x41\xd9\x04\x75\xd3\x92\xb1\xa1\x4c\xe5\xbb\x9f\xbc\x9c\x46\xd3\xc9\xe8\xe2\x09\xe1\x6c\xfd\x32\xdd\x47\x7e\x8e\x65\x3f\xc1\xc9\xc9\x7a\x19\xe7\x98\xd6\x23\xce\xbd\x10\x71\x3e\xc4\x32\xcf\x9d\x0d\xda\x58\x0f\x61\x8d\x40\xf8\xf7\xc6\x10\x16\x50\x1a\xac\x0b\x0f\xa5\x23\x18\x87\x83\x06\x8f\xd4\x21\x89\x70\xdf\xe2\x68\xed\x03\x6d\xf2\xc0\x6e\x0f\xda\x7d\x2c\x30\x91\x3c\x24\x61\xdc\x39\x48\x19\x36\x7a\x1d\xdc\xc1\x08\x3e\x68\x0a\x1e\x34\x58\xdc\x02\x57\xf5\xe0\x6d\x0a\xb1\xec\xc6\x0f\x6e\x1b\x0d\xb1\xb3\x86\x3d\xd8\x9a\xb0\x66\xf1\xcc\xd5\x6a\xef\xb1\xe0\x3b\xb1\x38\x06\xd7\xae\xaa\x90\xfa\x26\x8e\x9e\xd2\xbc\x1c\x7b\x3b\xf6\xb3\xef\x07\x08\xc7\x77\x1c\x08\xea\x0a\xb7\xc3\x5c\x49\x27\x22\x41\x5b\xb4\xce\xd8\xe0\x19\x73\x85\xdb\x8b\xf1\x3b\x1d\x6c\x87\xea\xff\x8a\xf9\x5a\x5b\x93\xeb\x1a\x0a\xd7\x68\x63\x95\x48\x90\x28\x67\xb3\x46\xdf\x62\xca\xc7\x80\x44\x8e\x06\x8b\x4b\x1b\x90\x68\xd3\x86\xd1\xb5\x12\x49\xe5\x1e\x74\x1c\xce\xbf\xf4\x3b\x29\xd3\x0d\xb6\xfd\x18\xe9\xfb\x70\x34\xe4\x5b\xf6\x53\x2a\xa9\x5d\xa5\xe6\x5c\xa6\xb5\x4d\x65\x20\x6d\x3d\x97\xa9\x1c\xc7\x12\x2f\x86\x06\xcf\xcb\xa3\x86\x61\xf2\xa4\x89\xc1\xe0\xc9\x32\x04\x02\xbf\x6e\xee\x38\x12\x49\xa3\x7a\x25\xa9\xcc\x22\x4d\x3f\xb7\x33\x39\xed\xe1\x83\xcc\x4f\x2c\x23\x9e\xa8\x4b\x5b\xe0\xdd\xe4\x84\x69\xde\x14\xb5\xb1\xf8\x3c\xc3\x79\x0f\x38\xc5\xc1\x0f\x53\x9f\xe0\x98\xf7\x80\x53\x1c\xfe\xbe\x59\xb9\xfa\x79\x8a\x9b\x78\x7e\x8a\x21\x90\xce\x4f\x68\x58\xf2\xf1\x24\xc6\x37\x16\xc5\xfb\xd7\x3d\xf2\x8f\x98\xc1\x0f\xb6\x38\x54\xe7\xf1\x4f\xa4\xe1\xc1\x92\x0e\x29\x8f\x63\xfd\x90\xcb\x1f\x48\x39\x1b\x7e\x97\xf1\xb1\x5d\xf9\x42\x6b\xce\xf7\x57\x7d\x8b\xbc\x79\x28\xb5\xb1\xca\x27\x2f\x50\xfc\xf0\x47\x5a\x1f\x0b\x8e\xad\xfb\x7f\x04\xb3\xa1\x9c\x1e\xeb\x1d\x67\x07\xab\xa9\xed\x94\xdb\x88\x55\x5b\x0c\x83\x9e\x54\x86\xbc\x7d\x02\x6c\xca\x88\x7d\x75\x06\xd6\xd4\xd1\xed\xe1\x36\x48\xc4\x9f\x84\x61\x43\x56\x24\xc9\x9e\x93\xe3\xa9\x1b\xc3\xc1\x34\xf1\x8e\xdf\x45\x23\x0e\x82\xf8\x8f\x1c\xfb\x83\x62\x77\xb4\x2b\xb5\xc0\x8a\xd5\xd0\x6e\xa7\x86\x09\xa2\xae\x74\x83\xfb\xfd\x00\xf3\x53\xf0\xd4\x3d\xaa\x02\xaf\xfa\x38\xd6\xf6\x38\x74\x8b\x8d\x7d\x25\x1e\x47\x08\xef\x0c\x07\xe7\xfd\xeb\x7e\x0a\xec\x85\xf8\x27\x00\x00\xff\xff\x6c\x11\xf8\x32\xe6\x09\x00\x00")

func svcServeGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcServeGotemplate,
		"svc/serve.gotemplate",
	)
}

func svcServeGotemplate() (*asset, error) {
	bytes, err := svcServeGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/serve.gotemplate", size: 2534, mode: os.FileMode(420), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _svcTransport_grpcGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x5b\x6f\x1b\x37\x13\x7d\x5e\xfe\x8a\xf9\x84\xe0\x83\x14\xc8\xdc\x3e\x07\xf0\x4b\x6c\x37\x09\xda\x24\x86\x6b\xb4\x0f\x41\x10\x50\xbb\xa3\x5d\x42\xbb\xe4\x9a\xe4\xca\x56\x89\xfd\xef\xc5\x70\x6f\xba\x59\x12\xfa\xd0\xa2\x02\x0c\xdb\xe4\xe1\x5c\xce\x99\x19\x52\x71\x0c\x37\x3a\x45\xc8\x50\xa1\x11\x0e\x53\x58\x6c\xc0\x99\xda\x5a\x0e\xb7\x5f\xe1\xcb\xd7\x47\xb8\xbb\xfd\xf4\xc8\x59\x1c\xc3\x03\x9a\x5a\x29\xa9\xb2\x16\x00\xcf\xb2\x28\x40\xaf\xd1\x3c\x1b\xe9\x10\x5c\x2e\x2d\x2c\x65\x81\x9c\xb1\x4a\x24\x2b\x91\x21\xd8\x75\xc2\xe8\xe8\x63\xbf\x07\x95\xd1\x6b\x99\xa2\x05\x8b\x66\x8d\xe6\xca\xca\x14\x61\x21\x55\x2a\x55\x66\x61\xa9\x0d\xb8\x1c\x21\x7b\xb8\xbf\x01\x67\x84\xb2\x95\x36\x2e\x78\xff\xe4\xa0\x76\xb2\x90\x7f\xa2\x0d\x90\x61\x37\xce\x4c\x95\xf0\xdf\x82\x39\xce\x98\x2c\x69\x11\xa6\x2c\x9a\x28\x74\x71\xee\x5c\x35\x61\xd1\x24\xd1\xca\xe1\x8b\x9b\x30\x16\x4d\x32\xad\xb3\x02\x79\xa6\x0b\xa1\x32\xae\x4d\x16\x4c\xc4\x25\x3a\x91\x0a\x27\x08\x43\x0b\x83\x07\x98\x64\xd2\xe5\xf5\x82\x27\xba\x8c\x33\x7d\xb5\x92\x2e\xa6\x9f\xdd\x10\xc8\xcb\x0e\x2e\xd3\x71\x65\xb4\xd3\x8b\x7a\x19\xbb\x4d\x85\x96\x0c\xf7\x64\x50\xbc\x32\x41\x16\x55\x0b\x98\x78\xcf\xef\xdf\x7f\x0a\x81\xdf\x0b\x97\xc3\x55\xd3\x4c\xd8\x2c\x30\xf7\x59\xac\xf0\xc3\xc3\xfd\x4d\x9b\x1f\x94\x62\x85\x16\x04\x58\x74\xa0\x97\x80\x2a\xad\xb4\x54\xce\x82\x58\x0b\x59\x88\x45\x81\x20\x68\x3f\x10\xe8\x3d\xef\xdc\xf0\x2f\xa2\xc4\xa6\xe9\x49\x5a\xd6\x2a\xd9\xb3\x3c\x1d\x4d\xdd\xf5\x7f\xcd\x41\x57\x4e\x6a\x65\x81\x73\xbe\xc3\x48\x47\xf7\xd7\xb0\x3d\x83\x6a\xc1\x5f\xf1\x05\x9e\x45\x76\x0b\x6b\xe1\xdd\x35\x7c\xfb\xfe\xba\x31\xcf\xa2\xe8\xd8\xee\x7b\x5c\x6a\x83\xd3\x5e\xa3\x47\x7d\xd3\x0a\x3a\x9b\xb3\xa8\xd9\xf7\x71\x0d\xa2\xaa\x50\xa5\xd3\x9d\xe5\x21\x1d\xce\xf9\x8c\x45\x06\x5d\x6d\x14\xfc\x9f\xbc\xb5\x3e\x7c\x90\xc7\x7b\x78\xd4\xbf\xea\x67\x34\xb0\x93\x12\x34\x0d\x8b\xbc\x37\x42\x65\x08\x6f\x24\x25\x32\xec\x7f\x46\x97\xeb\xd4\x12\x22\xf2\xbe\x3f\xfe\x46\x76\x5c\xbc\x83\xdd\x94\xbe\xe0\x73\xc7\x3a\x8b\xa2\x68\x60\x9e\x7b\x3f\x1c\xe9\x45\x98\x13\xe2\x16\x13\x9d\x06\xb1\xb6\x10\x0f\xf8\x54\xa3\x6d\x01\x77\xea\x28\xc0\x56\x5a\x59\x0c\x88\x1d\x26\x38\xe7\xb4\x48\xdc\x79\x7f\x45\x55\x44\x91\x37\xac\x09\x25\x37\x12\x02\xb2\xac\x0a\x2c\x91\xaa\x82\x7a\xce\xfb\x0f\x3a\x50\x71\x5c\x6b\xa9\x1c\x9a\xa5\x48\x90\x51\xb9\x6f\xdb\xb1\xce\xd4\x89\x03\xcf\xce\xf3\x77\x84\x3e\x80\x3d\xfe\x3e\x0a\x95\x16\x68\xd8\x18\x7c\x1b\x79\x67\x26\x8c\x91\x2d\xef\x4e\x8f\x89\x5c\x9e\xc3\xd9\x50\x01\x00\xbc\x97\x4b\x50\x08\xbc\x27\xfb\x71\x53\x75\xe5\x32\xb9\x2b\x2b\xb7\x99\x40\x87\x0c\x3d\x37\xb5\xf0\x76\x0c\x6c\x36\x06\x33\xe4\x3a\x4d\xdc\x0b\x74\xc3\x8a\x77\x35\x3e\x07\x83\x4f\xf0\x36\x74\xd9\x88\xef\xf4\x1f\x1c\x36\xcd\x0c\xa6\x87\xa0\xbd\xb0\x9a\x66\x0e\x68\x8c\x26\xe7\x21\x2e\xfa\xfc\x20\x0f\x55\xd8\xa0\x5c\xa9\x10\x0f\x44\x68\xfb\x90\x4a\x8c\x42\x0c\x21\xcd\x06\x03\x72\x19\xce\xfe\xef\x1a\x94\x2c\xb6\x0c\xd3\xa7\xeb\x33\x25\x8b\xe0\x60\xd8\x6b\xd8\x1e\xc2\x60\xc5\x2f\x49\x60\x36\x27\x63\x6c\xb4\xe1\x3d\x16\x16\xff\x59\x9e\xc3\x40\xff\x8f\x53\x7d\x51\x0e\xc7\xd8\x0e\x2d\xd7\xff\xa6\xc6\xeb\x9a\xa7\x9d\x53\xe7\x3b\x27\x8e\xe1\xd4\x48\x03\x49\x57\xd8\xde\x05\xdf\x1e\xe8\x10\x3f\x93\xc8\x2e\x17\x8e\x24\x5c\xa3\xa1\x0b\x30\x8c\xae\xf6\xda\x3b\x9c\x20\xa6\xb3\xec\x34\x08\xa8\x2d\x9a\xab\x54\x97\x42\xaa\x53\x60\x0e\xf7\x46\x96\xc2\xc8\x62\x43\x47\x96\x75\x01\x52\x85\xbb\x77\xeb\x16\x3d\x95\xc7\xf4\xc7\x61\x85\x51\x2e\x0f\xf8\x34\xce\x19\x4f\xe5\xb4\xf5\xdf\x56\xcd\x44\x54\x8e\xef\xae\xfb\x33\xc7\x9a\xe3\xa0\x34\x87\x7b\xcd\xe0\x53\xab\xdc\x51\xa5\xda\x0b\xe3\x22\xa5\x4e\xde\x2d\x47\xa5\x6a\x4f\xf4\x90\xd7\xb4\x3a\xaf\x42\xe7\x22\x68\x76\x42\xd9\xaa\xd8\x9c\x93\xea\x6f\x8c\xea\x93\x69\x1f\x53\x76\x88\xf7\x12\x69\xc7\x7e\xb4\x15\x51\xdf\x1f\xbe\x6c\x00\x1e\xb6\xb3\xad\xce\x0f\xc5\x7f\x2d\xa3\xcb\xe6\xcc\xb9\xa4\xa2\x63\x83\xe7\x23\x16\x15\x1a\xcb\xda\x66\x3c\x78\x1d\x1e\x9f\xf2\x65\x3a\x20\xf9\xe7\xdb\xd9\x3e\x80\x3a\x8f\x5e\x10\xab\x39\xac\x43\x26\xa1\x45\xca\x94\xd6\x23\xb9\x84\xf5\x38\x7c\xa3\xa8\x7d\xd0\x23\xac\x70\x13\x7a\x21\x4d\xe9\x3b\x94\x76\x39\x15\x60\xef\x85\x1e\x24\xa5\x70\x30\x5d\xcd\xe0\x39\x97\x49\x1e\xa0\x45\x01\x05\x15\x73\x67\x45\xa8\x34\x3c\xb2\xe8\x1b\x0b\xbf\x11\x4a\x2b\x99\x88\xe2\x23\x8a\x14\xcd\x2f\xb8\xa1\xe7\xbe\xeb\x1c\x59\xdd\x36\x94\x74\x90\x08\x05\x0b\xec\x4d\x24\x09\x5a\x8b\x29\xf9\x46\xe9\x72\x34\x9d\x67\xda\x27\x2a\xae\x87\x5c\xff\x90\x2e\xff\x5d\x14\x35\xb6\x57\x0c\xe5\xfa\xed\xa7\xef\xb3\xb3\xc0\x57\xa2\x9b\xae\x66\xa3\x85\xf0\x96\x1c\x06\x51\xe2\x5e\x58\xc3\xfe\x0a\x00\x00\xff\xff\xb3\x41\x42\x30\x69\x0e\x00\x00")

func svcTransport_grpcGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcTransport_grpcGotemplate,
		"svc/transport_grpc.gotemplate",
	)
}

func svcTransport_grpcGotemplate() (*asset, error) {
	bytes, err := svcTransport_grpcGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/transport_grpc.gotemplate", size: 3689, mode: os.FileMode(420), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _svcTransport_httpGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x1a\x29\x91\xe2\x05\xac\xf0\x31\x08\x93\xb3\xce\x1f\x1a\xeb\x76\xa7\x61\x80\x37\xa7\xae\xb2\x03\x52\xed\x42\xff\x8c\xa1\x15\x67\xb5\xd7\x93\xfa\x20\x3b\xbd\x9c\xa3\x9b\x53\x89\x77\x6f\x85\x18\xa9\x9a\xdc\xcd\xe5\xb0\x1b\x64\xd5\x88\x65\xce\xa3\xb4\x26\x69\xcb\xf9\xba\xa1\x75\x78\xda\xe1\x5f\x78\xfe\x1b\x49\x11\xcb\x2f\x00\x00\xff\xff\xdd\x3a\x4a\x8f\x6a\x00\x00\x00")

func svcTransport_httpGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcTransport_httpGotemplate,
		"svc/transport_http.gotemplate",
	)
}

func svcTransport_httpGotemplate() (*asset, error) {
	bytes, err := svcTransport_httpGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/transport_http.gotemplate", size: 106, mode: os.FileMode(420), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"client/grpc/client.gotemplate": clientGrpcClientGotemplate,
	"client/http/client.gotemplate": clientHttpClientGotemplate,
	"cmd/server/main.gotemplate": cmdServerMainGotemplate,
	"svc/endpoints.gotemplate": svcEndpointsGotemplate,
	"svc/handlers/handlers.gotemplate": svcHandlersHandlersGotemplate,
	"svc/handlers/hooks.gotemplate": svcHandlersHooksGotemplate,
	"svc/serve.gotemplate": svcServeGotemplate,
	"svc/transport_grpc.gotemplate": svcTransport_grpcGotemplate,
	"svc/transport_http.gotemplate": svcTransport_httpGotemplate,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"client": &bintree{nil, map[string]*bintree{
		"grpc": &bintree{nil, map[string]*bintree{
			"client.gotemplate": &bintree{clientGrpcClientGotemplate, map[string]*bintree{}},
		}},
		"http": &bintree{nil, map[string]*bintree{
			"client.gotemplate": &bintree{clientHttpClientGotemplate, map[string]*bintree{}},
		}},
	}},
	"cmd": &bintree{nil, map[string]*bintree{
		"server": &bintree{nil, map[string]*bintree{
			"main.gotemplate": &bintree{cmdServerMainGotemplate, map[string]*bintree{}},
		}},
	}},
	"svc": &bintree{nil, map[string]*bintree{
		"endpoints.gotemplate": &bintree{svcEndpointsGotemplate, map[string]*bintree{}},
		"handlers": &bintree{nil, map[string]*bintree{
			"handlers.gotemplate": &bintree{svcHandlersHandlersGotemplate, map[string]*bintree{}},
			"hooks.gotemplate": &bintree{svcHandlersHooksGotemplate, map[string]*bintree{}},
		}},
		"serve.gotemplate": &bintree{svcServeGotemplate, map[string]*bintree{}},
		"transport_grpc.gotemplate": &bintree{svcTransport_grpcGotemplate, map[string]*bintree{}},
		"transport_http.gotemplate": &bintree{svcTransport_httpGotemplate, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

