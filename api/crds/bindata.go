// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// installer.kubedb.com_kubedbcatalogs.yaml
// installer.kubedb.com_kubedboperators.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _installerKubedbCom_kubedbcatalogsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\x1b\x37\x0c\xbd\xfb\x57\x10\xe8\x21\x97\xda\xc6\xa2\x97\x62\x6e\xed\xa6\x05\x82\xf4\x0b\xbb\x69\xee\x1c\x89\x1e\xab\xab\x11\x15\x92\x72\x77\xfb\xeb\x0b\x69\x3c\xfe\xd8\x75\x9a\xae\x81\xe8\xa6\x47\x3e\x4a\x7c\x24\x35\x83\x39\x7c\x24\xd1\xc0\xa9\x03\xcc\x81\x1e\x8d\x52\xdd\xe9\xea\xe1\x7b\x5d\x05\x5e\xef\x6e\x7a\x32\xbc\x59\x3c\x84\xe4\x3b\xb8\x2d\x6a\x3c\xde\x91\x72\x11\x47\x6f\x69\x13\x52\xb0\xc0\x69\x31\x92\xa1\x47\xc3\x6e\x01\xe0\x84\xb0\x82\x1f\xc2\x48\x6a\x38\xe6\x0e\x52\x89\x71\x01\x10\xb1\xa7\xa8\xd5\x07\x00\x73\xee\xe0\xa1\xf4\xe4\xfb\x05\x40\xc2\x91\xe6\xad\x43\xc3\xc8\x83\xae\x42\x52\xc3\x18\x49\x56\x93\x61\xe5\x78\x5c\x68\x26\x57\x23\x0c\xc2\x25\x77\x70\xd1\x67\x8a\xb7\x3f\xc8\xa1\xd1\xc0\x12\xe6\xfd\xf2\x78\x6a\xdd\x60\xce\xea\xd8\x53\xdb\x4e\x59\xbe\x2f\x3d\xbd\xfd\xf1\x76\xba\x46\xc3\x63\x50\x7b\xff\xd2\xf6\x4b\x50\x6b\xf6\x1c\x8b\x60\x7c\x9e\x40\x33\x69\x48\x43\x89\x28\xcf\x8c\x0b\x80\x2c\xa4\x24\x3b\xfa\x33\x3d\x24\xfe\x3b\xfd\x1c\x28\x7a\xed\x60\x83\x51\xeb\x6d\xd4\x71\xa6\x0e\x7e\xab\x99\x64\x74\xe4\x17\x00\x3b\x8c\xc1\x37\x71\xa7\x5c\x38\x53\xfa\xe1\x8f\x77\x1f\xbf\xbb\x77\x5b\x1a\x71\x02\x6b\x64\xce\x24\x76\x48\x79\xd2\xfb\x50\xe9\x03\x06\xe0\x49\x9d\x84\xdc\x22\xc2\x9b\x1a\x6a\xf2\x01\x5f\x6b\x4b\x0a\xb6\x25\xd8\x4d\x18\x79\xd0\x76\x0c\xf0\x06\x6c\x1b\x14\x84\x5a\x0e\xc9\xda\x95\x4e\xc2\x42\x75\xc1\x04\xdc\xff\x45\xce\x56\x70\x5f\xf3\x14\x05\xdd\x72\x89\x1e\x1c\xa7\x1d\x89\x81\x90\xe3\x21\x85\x7f\x0e\x91\x15\x8c\xdb\x91\x11\x8d\xf6\xda\xce\x2b\x24\x23\x49\x18\xab\x08\x85\xbe\x05\x4c\x1e\x46\x7c\x02\xa1\x7a\x06\x94\x74\x12\xad\xb9\xe8\x0a\x7e\x65\x21\x08\x69\xc3\x1d\x6c\xcd\xb2\x76\xeb\xf5\x10\x6c\xee\x6d\xc7\xe3\x58\x52\xb0\xa7\xb5\xe3\x64\x12\xfa\x62\x2c\xba\xf6\xb4\xa3\xb8\xd6\x30\x2c\x51\xdc\x36\x18\x39\x2b\x42\x6b\xcc\x61\xd9\x2e\x9e\xac\x0d\xc8\xe8\xbf\x91\xfd\x20\xe8\x9b\x93\x9b\xda\x53\x2d\x9b\x9a\x84\x34\x1c\xe0\xd6\x58\x9f\xd5\xbd\xb6\x16\x04\x05\xdc\xd3\xa6\xfb\x1f\xe5\xad\x50\x55\xe5\xee\xa7\xfb\x0f\x30\x1f\xda\x4a\x70\xae\x79\x53\xfb\x48\xd3\xa3\xf0\x55\xa8\x90\x36\x24\x53\xe1\x36\xc2\x63\x8b\x48\xc9\x67\x0e\xc9\xda\xc6\xc5\x40\xe9\x5c\x74\x2d\xfd\x18\xac\x56\xfa\x53\x21\xb5\x5a\x9f\x15\xdc\x62\x4a\x6c\xd0\x13\x94\xec\xd1\xc8\xaf\xe0\x5d\x82\x5b\x1c\x29\xde\xa2\xd2\x57\x97\xbd\x2a\xac\xcb\x2a\xe9\x97\x85\x3f\x7d\x98\xce\x1d\x27\xb5\x0e\xf0\xfc\xae\x5c\xac\xd0\xd9\xd4\xdf\x67\x72\xb5\x5a\x55\xb2\xca\x82\x0d\x0b\x08\xf9\xa0\xf3\xa4\x9c\x84\xb9\x34\x8a\xfb\x57\xa9\xc6\x3a\x07\x3f\xef\x5e\x17\x45\x54\x0b\x4e\xa9\xea\xf3\xd2\x3c\xe7\xd5\x33\x47\xc2\xf4\x92\x6e\xce\xbf\x9e\x35\xd2\xe8\xd0\x6d\xe9\x1a\x2a\xa7\x81\xaf\xa0\x3d\xe9\xa7\xf8\x7a\x5a\x26\x71\x9c\xf0\xd1\x04\x7d\x7f\x05\x7d\xe8\xb9\x24\x47\x72\x05\x95\xd5\x06\xb9\x54\xb1\x2f\x32\x85\x1f\xaf\xcb\xb6\xb5\xdb\x6b\x69\x17\xfb\xbe\x2e\xcf\xee\x81\xe4\x8e\x86\xa0\x26\x4f\xcf\xc3\x5e\x9c\xab\xba\x36\x25\xc6\xfa\xa1\xfd\x7d\x47\x22\xc1\xd3\xff\x26\x5e\x41\xaa\x0f\x50\x90\xf3\x36\x5c\x82\x3b\xf9\x46\xcf\xd8\x79\x32\xff\x3d\xf6\xcf\xa0\xdd\xfc\x2b\xb4\xbb\xc1\x98\xb7\x78\x73\xc4\x9a\xd8\xcb\xfd\x7f\xca\x89\x19\xa0\x7d\xc3\x7d\x07\x26\x65\xfa\x89\x50\x63\xc1\x81\xf6\xc8\xbf\x01\x00\x00\xff\xff\x3b\xa9\x69\x86\x62\x09\x00\x00")

func installerKubedbCom_kubedbcatalogsYamlBytes() ([]byte, error) {
	return bindataRead(
		_installerKubedbCom_kubedbcatalogsYaml,
		"installer.kubedb.com_kubedbcatalogs.yaml",
	)
}

func installerKubedbCom_kubedbcatalogsYaml() (*asset, error) {
	bytes, err := installerKubedbCom_kubedbcatalogsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.kubedb.com_kubedbcatalogs.yaml", size: 2402, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _installerKubedbCom_kubedboperatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x6d\x73\xdc\x38\x72\xfe\xee\x5f\x81\x72\x52\x65\x29\xa5\x19\xad\x6f\x53\x57\x89\xee\xea\x52\x3a\xdb\xbb\xe5\x5a\xbf\xa8\x2c\xef\xee\x25\x97\xcb\x16\x86\xec\x99\xc1\x89\x04\xb8\x00\x28\x79\x36\x95\xff\x9e\x42\x03\xe0\x90\x43\x02\x04\x47\xd2\xd9\x7b\x47\x7e\xb1\x35\x24\x9b\x40\xa3\xd1\xe8\x7e\xfa\x21\x41\x2b\xf6\x03\x48\xc5\x04\xbf\x20\xb4\x62\xf0\x49\x03\x37\x7f\xa9\xe5\xcd\xbf\xa9\x25\x13\xe7\xb7\xcf\x57\xa0\xe9\xf3\x27\x37\x8c\xe7\x17\xe4\x45\xad\xb4\x28\x3f\x80\x12\xb5\xcc\xe0\x25\xac\x19\x67\x9a\x09\xfe\xa4\x04\x4d\x73\xaa\xe9\xc5\x13\x42\x32\x09\xd4\xfc\xf8\x91\x95\xa0\x34\x2d\xab\x0b\xc2\xeb\xa2\x78\x42\x48\x41\x57\x50\x28\x73\x0d\x21\xb4\xaa\x2e\xc8\x4d\xbd\x82\x7c\xf5\x84\x10\x4e\x4b\xf0\x7f\x8a\x0a\x24\xd5\x42\xaa\x25\xe3\x4a\xd3\xa2\x00\xb9\xb4\x67\x96\x99\x28\x9f\xa8\x0a\x32\x23\x62\x23\x45\x5d\x5d\x90\xc1\x6b\xac\x40\xf7\xa4\x8c\x6a\xd8\x08\xc9\xfc\xdf\x8b\xfd\x63\xcd\x1f\xb4\xaa\x54\x26\x72\xc0\x3f\x6d\x37\xbf\xab\x57\xf0\xf2\x8f\xef\x5d\x3b\xf0\x44\xc1\x94\xfe\x6e\xe0\xe4\x1b\xa6\x34\x5e\x50\x15\xb5\xa4\x45\xaf\x0f\x78\x4e\x31\xbe\xa9\x0b\x2a\x0f\xcf\x3e\x21\xa4\x92\xa0\x40\xde\xc2\xf7\xfc\x86\x8b\x3b\xfe\x0d\x83\x22\x57\x17\x64\x4d\x0b\x65\x5a\xa4\x32\x51\xc1\x05\x79\x67\x7a\x53\xd1\x0c\xf2\x27\x84\xdc\xd2\x82\xe5\xa8\x61\xdb\x1f\x51\x01\xbf\xbc\x7a\xfd\xc3\xd7\xd7\xd9\x16\x4a\x6a\x7f\x34\x92\xcd\x63\x74\xd3\x6d\xab\xf4\x66\xb8\x9b\xdf\x08\xc9\x41\x65\x92\x55\x28\x91\x3c\x33\xa2\xec\x35\x24\x37\x03\x0c\x8a\xe8\x2d\x90\x5b\xfb\x1b\xe4\x44\xe1\x63\x88\x58\x13\xbd\x65\x8a\x48\xc0\x3e\x70\x8d\x4d\x6a\x89\x25\xe6\x12\xca\x89\x58\xfd\x15\x32\xbd\x24\xd7\xa6\x9f\x52\x11\xb5\x15\x75\x91\x93\x4c\xf0\x5b\x90\x9a\x48\xc8\xc4\x86\xb3\x5f\x1a\xc9\x8a\x68\x81\x8f\x2c\xa8\x06\xa7\x5d\x7f\x30\xae\x41\x72\x5a\x18\x25\xd4\x70\x46\x28\xcf\x49\x49\x77\x44\x82\x79\x06\xa9\x79\x4b\x1a\x5e\xa2\x96\xe4\xad\x90\x40\x18\x5f\x8b\x0b\xb2\xd5\xba\x52\x17\xe7\xe7\x1b\xa6\xbd\x81\x67\xa2\x2c\x6b\xce\xf4\xee\x3c\x13\x5c\x4b\xb6\xaa\xcd\xb0\x9d\xe7\x70\x0b\xc5\xb9\x62\x9b\x05\x95\xd9\x96\x69\xc8\x74\x2d\xe1\x9c\x56\x6c\x81\x0d\xe7\x1a\x67\x49\x99\xff\x93\x74\xb3\x41\x3d\x6b\xb5\x54\xef\xcc\xb0\x29\x2d\x19\xdf\x34\x3f\xa3\x71\x05\xf5\x6e\xac\x8b\x30\x45\xa8\xbb\xcd\xb6\x7f\xaf\x5e\xf3\x93\xd1\xca\x87\x57\xd7\x1f\x89\x7f\x28\x0e\x41\x57\xe7\xa8\xed\xfd\x6d\x6a\xaf\x78\xa3\x28\xc6\xd7\x20\xed\xc0\xad\xa5\x28\x51\x22\xf0\xbc\x12\x8c\x6b\xfc\x23\x2b\x18\xf0\xae\xd2\x55\xbd\x2a\x99\x36\x23\xfd\x73\x0d\x4a\x9b\xf1\x59\x92\x17\x94\x73\xa1\xc9\x0a\x48\x5d\xe5\x54\x43\xbe\x24\xaf\x39\x79\x41\x4b\x28\x5e\x50\x05\x8f\xae\x76\xa3\x61\xb5\x30\x2a\x1d\x57\x7c\xdb\x3b\xf9\x63\x68\x7a\x98\x03\x3d\x51\xe7\x17\x42\x4a\xfa\xe9\x0d\xf0\x8d\xde\x5e\x90\xdf\x7e\x7d\x70\xae\xa2\xda\x98\xe4\x05\xf9\x9f\x3f\xd3\xc5\x2f\x7f\x39\xf9\xf3\x82\x2e\x7e\xf9\x6a\xf1\xef\x7f\xf9\x97\x3f\xbb\xff\x9c\xfe\xc7\x3f\x1f\xdc\x33\xd8\x48\xff\xb3\x1d\xc0\xe6\x67\xef\xee\x06\x8d\xa6\xeb\x8b\xae\x2b\xc8\x8c\x05\x99\x61\x34\xb7\x91\xb5\x90\x44\x42\xce\x94\x9f\xbd\x09\xfd\xa7\x79\x8e\x2e\x9d\x16\x57\x22\xbf\x86\xac\x96\x4c\xef\xae\x44\xc1\xb2\xde\xa5\x84\x30\x0d\x65\xef\xc7\x60\xff\xf6\xa7\xa8\x94\x74\xd7\x7d\xec\x1a\xd7\x92\xdd\xa1\xb0\x4e\x77\x5f\xaf\xb1\x5f\x6c\xcd\x20\x3f\xc3\x6e\x56\x22\x7f\xa6\xd0\x6f\xe4\x75\x61\x66\x48\x26\xb8\xd2\x92\x32\xae\xd5\xe1\x40\x05\x3a\x6c\x0e\x2e\x72\xb8\x0c\xb4\xa0\xd7\x8a\x97\xf8\xc7\x0a\x14\xde\xd6\xb4\xbc\xdd\x0a\x59\x17\xa0\x50\xfd\xae\x91\xcb\x01\xa1\xb1\x06\xd9\xf3\xb0\x06\x29\x21\x7f\x59\x1b\x45\x5e\x37\xe2\x5f\x6f\xb8\x68\x7e\x7e\xf5\x09\xb2\x5a\x1f\x78\xf4\x60\xdb\x3f\x1a\xd3\xb0\x82\x40\x92\x3b\x56\x14\xee\x31\xc6\xe7\xfa\x13\xa6\xc1\xe8\x84\x4d\xff\x0e\xd5\xb8\x3f\xf4\x96\x6a\xa2\xa8\x66\x6a\xbd\xc3\x7e\x36\x9a\x80\x4f\xc6\xf9\x60\x1c\xb1\x1f\x30\xb2\xda\x39\xbf\x63\xd6\xb8\xb3\xa0\xd8\x55\xad\x09\xd3\xe8\xac\xb2\xad\x10\x0a\x08\xb5\x8a\xc6\xe7\xdd\x32\x81\xcb\x02\x11\x1c\x88\x90\xa4\x34\x5e\x06\x97\x22\x08\x4a\x6c\x35\x67\x89\x1a\xd8\x8b\x63\x8a\x94\x42\xe9\xbd\xae\xfd\xfc\x31\xe2\xef\x98\xde\x46\x7a\x0f\x64\x63\x22\x1d\x50\x9a\xa8\xba\x34\x8d\xb8\x03\xb6\xd9\x6a\x75\x46\xd8\x12\x96\x38\xfc\x40\xb3\x6d\xeb\x71\x25\x40\xcf\x2e\xf7\x07\x2d\x0a\xd7\x95\x8e\x2d\xc1\xcf\x35\x93\x50\x1a\x5f\x4e\x4e\x1a\xc7\xef\x9c\xf1\x99\x3f\xdf\xb3\x92\xf0\x63\x06\x86\xe9\x8c\x80\xce\x96\xa7\x67\x24\x13\x65\x55\x6b\xa3\x73\xd3\xa7\xd5\xce\x4c\x71\x49\xdd\xe2\x23\x45\xbd\x89\x6b\x04\x0a\xd7\x50\x1f\x1d\xe0\x60\xe3\x32\x6d\x1c\x0b\xdf\x90\xa7\x56\x49\x4f\xfd\x22\xaf\xea\x32\x28\x91\x59\x65\xa0\xfe\x4a\xaa\xb3\xad\x8b\x45\x32\x21\x25\xa8\x4a\x70\x94\x88\x67\x5e\xed\xfb\xf2\xbb\xa8\x31\x18\x61\x27\xea\x14\x07\x17\x85\x6d\xd9\x66\xeb\xc7\x90\x4a\xc0\xdf\xba\x36\x31\x34\x79\x49\xd8\xfb\xf9\xa3\x33\xf1\x2e\x39\x81\xb2\xd2\xbb\x96\xa5\xb5\xc6\x58\x83\x2c\x9b\x1e\x52\x8c\x95\x43\x87\x5d\x1e\x94\x6d\x3f\x2b\x2b\xe3\x98\xb5\xb3\x3c\xf2\x15\x39\x41\xd3\x63\xfa\x99\xc2\x69\xb3\x10\xd5\xe9\x92\x5c\xfa\x00\x3c\x74\x8c\x37\x8a\x8b\xe6\xc9\xee\x11\xa6\xa1\x4a\x44\x84\x36\xcf\x0f\x5e\x33\xe6\x01\xdb\x8d\x03\x9e\xf5\xd6\xe5\xee\xd1\xd5\xb7\xb5\x1a\x05\x05\x64\xda\xf8\x61\x90\xe5\x19\xa1\x4a\x89\x8c\x99\x68\xa5\x19\xff\xa8\x48\x72\x60\x6a\x56\xcd\xe1\x0e\xa5\x77\x8a\x60\x58\xd1\x35\xdc\xb1\xeb\x7b\x5d\x34\x49\x89\x99\x69\xdd\xae\xb6\x1d\xc6\xa8\x44\x62\xe6\xb8\xb9\xff\x99\x72\xe9\x59\xbc\x77\x64\xdc\xee\x83\xcd\x0d\x36\xd3\x85\xbd\xee\x4c\x82\x60\xb7\xf8\x98\xd0\x91\x32\xae\x5c\xa8\x7f\x46\x28\xb9\x81\x9d\xcd\x0a\x4c\xe2\xe1\xe2\x22\xbc\x38\x49\xaa\x04\xbb\xb8\x18\x1f\x70\x03\x3b\x14\xe4\xd2\x88\x84\xfb\xd3\x47\xde\x1e\x37\x30\x18\x6c\x0c\x1d\xbd\x45\x1c\xc7\x0a\xdb\x88\x9a\x40\x4f\x3a\x45\x7f\xc4\xa6\xe0\x05\x03\x0c\xe7\x13\xef\x89\x04\x76\xc3\x87\x1f\x82\xa3\xfa\xf9\xa1\xc9\x61\xec\xc0\x3e\x53\x76\x80\xcc\x5c\xd9\xb2\x2a\xb9\x9f\x5a\xa0\x75\xe1\x54\xf1\x49\xe1\x0f\x26\x89\x6e\x9a\xa7\xd0\xf3\xbf\xe6\xe1\xa8\xe4\xf0\x78\x27\xf4\x6b\x7e\x46\x5e\x7d\x62\xca\x2c\xf8\x2f\x05\xa8\x77\x42\xe3\x9f\x4b\xf2\xad\xb6\x36\xf8\x66\xc4\x55\xb4\x9a\x38\x55\xb1\xb6\x1f\x47\xa9\xf5\x92\xdb\xf8\xdb\xa8\xa3\x9d\x6a\xaa\xa5\x09\xb0\xc7\x5d\xe2\xfe\x68\x26\x18\x53\x26\xf9\x13\xd2\xab\x05\x01\x03\x94\x39\x10\xea\xc7\x8e\xb2\x56\x98\x53\x72\xc1\x17\xb8\x5e\xfa\x36\x75\x9e\x65\xb5\x9e\xde\x4c\xd9\x19\x9f\x7e\xf3\xfc\x63\x93\x25\x86\x9b\xf6\xad\x36\x8f\x7b\xd3\x79\x48\xfa\x84\xdc\x37\x66\x4b\x6f\x31\x08\x63\x7c\x53\x34\x61\xd5\x19\xb9\xdb\xb2\x6c\x8b\x71\x7b\xb2\xd0\x15\x58\xd4\xa4\x92\x60\xd6\x3d\xaa\x8c\x6b\x34\xbf\x6c\x40\x9a\x70\x98\x79\x25\xb0\xf4\x86\x4a\xa8\x0a\x9a\x41\x4e\x72\x0c\x3a\x2d\x66\x41\x35\x6c\x58\x46\x4a\x90\x1b\x30\x69\x71\xb6\x4d\xb5\xfe\xe4\x05\xc5\x1e\x93\x27\x4b\x38\xed\x1c\x3e\x7c\x48\x9d\xd2\xa4\x85\xf1\x4c\x49\xd7\x89\x36\x9e\x98\xd2\xdc\x03\x24\x20\x7e\x71\x4a\xdf\x30\xe0\x70\x10\xe3\x67\x8e\x35\x30\x2f\x98\x63\x8d\x39\xd6\x08\x1e\x73\xac\xe1\x8f\x39\xd6\x98\x63\x8d\x39\xd6\x98\x63\x8d\x5f\x51\xac\x91\x28\xd4\xe2\x29\x13\x60\x9d\x1f\x2d\xce\x75\x88\xe3\x60\x60\xe3\x0b\x64\x1d\xc8\x66\xa4\x47\x26\x4c\xb8\x76\x6b\xd9\x47\x84\x88\x18\x47\x21\x92\xf2\x0d\x90\xe7\x8b\xe7\x5f\x7d\x15\xb7\xac\xb5\x90\x25\xd5\x17\xc6\xca\xbf\xfe\x4d\x82\x4e\xdc\x6c\x08\x5e\x39\x6e\x0f\x8b\x16\x22\x16\xb9\xc8\xea\x36\x8c\xd6\x8e\x8f\xd0\xd8\x60\x87\x90\xe7\x7b\xd4\x27\x9c\x97\x6b\x20\xea\x0e\xf8\xdd\x2b\x25\x04\x3b\xe7\x50\x67\x69\x9c\xbb\x26\x25\x68\x42\x75\x07\xda\x64\x25\x34\x05\x24\x5b\x06\xb1\xc5\xcc\xa0\x44\x5f\x1b\xc9\x89\xe0\x0e\xb9\x36\xb6\xb3\x4c\x6c\x71\xb8\xda\xd1\x2e\x8a\x90\x0c\xa8\x02\x13\x43\xac\xa0\x69\xb5\x28\x4d\x2b\x19\xd7\xde\x01\x9a\x26\x83\xd7\x6a\x50\xf0\x09\x2c\x37\x4b\x92\xd7\x28\x8e\x72\x57\xa5\x3d\xb5\xbd\x56\x3b\xa5\xa1\xc4\x1a\x8b\x90\xf8\x8f\xe9\xbe\x96\x3b\xa2\xc3\x88\x2e\xdc\x02\xd7\x35\x2d\x8a\x1d\x81\x5b\x96\xe9\x46\x7f\x58\x48\x66\xda\xd6\xc3\x42\xb3\x25\x25\x60\x3d\x9c\x8d\x51\x3f\x7d\x10\xbe\x59\x53\x5c\x06\x33\x15\x6d\xe4\x61\xf9\x27\x3e\x49\xcd\x65\x68\x39\xef\x3f\x84\x91\x7f\x92\xb6\x90\x1c\xe6\x24\x75\x51\x18\x7d\xdb\x42\x40\xbf\x79\x1e\x6c\x1f\xf5\x59\x1e\x8a\xb7\xd5\xac\x8e\xc5\xd9\xfa\x91\xad\x64\x5c\xbe\x7b\x69\x34\x32\xd6\x65\x42\x3e\x8a\x4a\x14\x62\xb3\x6b\xeb\x1e\x67\x3f\x16\x18\x9c\x64\x4a\x54\xbd\x72\x91\xed\x78\xe0\xf6\xee\x60\x28\x67\xcc\x7c\xce\x63\x87\x8e\x39\x8f\xed\x1d\x73\x1e\x9b\xd8\xc4\x39\x8f\xc5\x63\xce\x63\xe7\x3c\x76\xf4\x98\xf3\xd8\x81\x8b\x67\xcc\x7c\x8e\x35\x22\xc7\x1c\x6b\xf4\x8e\x39\xd6\x98\x63\x8d\x39\xd6\x98\x63\x8d\xe8\x31\xc7\x1a\x03\x17\x3f\x18\x66\x3e\x2e\x6e\x4c\x3d\x8b\x3e\xd0\x16\x45\x80\x83\x4d\x8a\x9e\xae\x44\x7e\x04\xa5\xbe\x12\x79\x84\x51\x6f\x41\xcd\x4c\x2c\x0a\x91\x51\x3d\xec\x0f\x10\x4e\x35\x62\x1c\x92\xaf\x68\x69\xb1\xda\x33\xf2\x8b\xe0\x60\x99\xce\x66\x9a\x21\xb2\x2a\xf4\x16\xa4\xb9\xfc\x44\x9d\x0e\x32\x55\x67\x96\xfe\xe0\x31\xb3\xf4\x67\x96\xbe\x3b\xda\x2c\xfd\x2d\x55\xd6\x2e\xed\x42\x18\x26\xed\xb7\xbc\x83\x71\x40\xbf\x8b\xb6\xf7\x33\x71\xf6\x8d\x11\x3a\x63\xc1\x57\x19\xf7\x03\x6f\xfb\x95\xbb\x72\x24\xe4\x57\xdd\xde\x44\xbc\xb7\xcd\xe1\xb0\xd1\x34\xcf\x21\x27\x15\xc8\x85\x35\x3d\x41\xd6\x8c\xe7\x03\x7d\xf1\xfd\x0f\x8a\x4d\xe4\xd1\x77\x1b\x39\xa1\x74\xd1\xae\xae\x74\x1c\xf4\x21\xab\x7e\x64\x2d\x6c\xc6\xef\x31\x59\xf5\x98\x79\xf9\xc5\x6d\x7a\xca\x8e\x79\xdb\xcf\x35\xc8\x1d\x11\xb7\x20\xf7\x99\x49\xf3\x9e\x67\x4a\x12\x82\x6b\x0f\x53\x24\xa3\xca\x3a\xea\xf1\x50\x6b\x5a\x76\x3a\xbd\x0e\xd2\xeb\xec\xa1\x08\x9b\xe5\x7b\xcc\x02\x15\x91\x18\xbd\x0d\x42\x1b\x03\xc5\x29\x2a\x53\x43\x78\x5b\xba\x4a\xba\x78\x52\x70\x3a\x38\xda\x01\xc8\x23\x3d\x2d\x68\x95\xf1\x46\x60\x8f\x74\x99\x07\xf0\xc8\x3d\xa1\x0f\x72\x04\xfc\x41\xa6\x41\x20\xe4\x50\xbd\xa6\x95\x6e\x9d\xee\xa3\x21\x13\x84\xb6\xec\x6b\x3a\x22\x42\x8e\xcb\x47\xa6\x23\x23\xe4\xb0\xfb\xcd\xf0\xc9\x1e\x4c\x32\xa9\xf3\x6d\x48\x25\x0c\x95\x4c\x12\xd9\x83\x55\xba\x70\x09\xda\x56\x07\x31\x79\x6c\x65\x4f\x43\x4b\xc8\xa1\xaa\x1d\x56\xc0\x30\x75\x3e\xc0\x4e\x26\x29\xa6\x8b\xb3\x04\xf1\x93\x49\x32\x43\x60\x46\x17\x43\x99\x2c\xb2\x8f\xb7\xf4\x70\x94\x87\x69\xa6\x6b\xe2\x1e\x88\x98\x24\xd6\x7e\x20\xe2\x21\xc1\x08\x32\x1d\x90\x20\xc7\xda\xe5\x54\x60\x82\x4c\x04\x27\xc8\x04\x80\x82\x4c\x05\x29\xc8\x54\xa0\x82\x4c\xee\x2f\x86\x10\x6f\x5a\x5f\x75\x19\x3b\x5a\x5f\x17\x98\xbc\x1a\x4d\x1e\xc1\x7e\xb4\x63\x9b\x6a\x03\x9d\x92\x56\xc6\x4b\xfc\xaf\x59\x9a\xd1\xf0\xff\x2f\x75\x1d\xa5\x4c\x2a\x13\x0a\x3b\xf0\xaf\x25\xc1\x63\x0e\xad\x87\x25\x0a\x35\xad\x61\x8a\x18\xdb\xb9\xa5\x85\x09\x40\x2c\x6d\xcb\xa5\x6a\xa6\xa5\x87\xf1\x5a\xea\xfc\xbe\xdb\x9a\xf4\xdc\x2c\xbe\x36\xcd\x63\x8a\x3c\xbd\x81\xdd\xd3\xb3\x9e\x1f\x79\xfa\x9a\x3f\x4d\x95\x4a\x5d\xaa\xd2\xf1\x19\x4d\xe4\x23\x78\xb1\x23\x4f\xf1\xdc\xd3\xd4\x89\x3d\x14\x2e\x4e\x09\x04\x8f\x00\xe5\x92\x2e\xe6\xfe\xe3\x3b\x53\x0b\x80\xfb\x1b\x1b\x7c\xc5\x27\xc6\xfb\x53\x29\x68\xa3\x8f\xa0\xae\xfb\x71\x10\x39\x69\x5e\x1b\xdf\x18\xcd\xeb\xd3\x70\x2a\xdd\xea\x52\x87\x89\x86\x21\x7f\x09\x94\x2b\xf2\xd4\xa3\x67\xcf\xd4\xbe\x8d\x4f\x1f\xae\xe2\x38\x69\x0e\xa7\xfb\x22\xed\x08\x6c\xdf\xa5\x84\xab\x07\x39\xbe\x43\x0b\xdd\x57\x89\x56\xb0\x87\x17\x73\x72\xe2\x33\xdd\x70\xee\xbd\x3f\x84\x44\x16\x65\xe7\x76\xae\xd9\xa2\x91\xb1\xcf\x7f\x4d\x46\x98\xea\x5e\x3d\xad\xb9\x6b\x01\x1e\xdc\x6c\x70\xbb\xbd\x45\xa5\xcc\xe0\xbb\x2d\xc8\x4e\x4f\x99\x72\x5f\x7b\xc2\x0a\x84\xac\x39\x37\xcf\x15\xdc\xc1\x7a\x49\x22\x8d\x9b\xb1\x1f\x2d\x72\x30\x89\x0d\xfb\xb1\xd7\x18\xfb\xef\x47\x29\x91\xea\x48\x3c\x80\x89\x5f\x92\x72\x9c\x49\xc1\xdd\x24\x32\xbf\x78\x24\x0e\xf5\x02\x79\xaa\x66\x59\xd3\xc7\x25\x79\x85\x93\xa0\xdd\x38\xa6\x70\x24\x69\x51\x88\xbb\x14\xef\x93\x6c\xd5\x69\xb1\xc1\xa2\xdd\x98\x87\x28\x19\x4c\xa6\xd9\xdf\x3d\x30\xcd\xfe\x00\x7a\xfa\x95\xb0\xec\x13\x41\xbd\x99\x6a\x3f\x53\xed\x5b\x54\x7b\xbc\xc9\x7a\xbe\x71\xce\x7d\xd8\x66\x90\x8b\x9f\xca\xb9\x27\x3f\x6e\x01\x67\x54\x04\x60\x33\x43\x54\xd6\x85\x66\xd5\xbe\x60\xad\x6c\xd3\x0a\x9b\x3e\x5a\xa2\x92\x3a\x40\x67\x63\x6f\x04\xd0\x6c\x7b\x38\x4d\xf0\x39\x58\xd0\x56\xe8\x91\x5d\x99\x85\x16\x85\xe3\xd6\x9b\xbc\x32\x3c\x46\xe0\x6a\x55\xec\x61\x20\xfc\x97\xee\x0b\x86\x0d\x68\x82\xc5\x89\x13\xb3\x58\x16\xc6\x1c\xcc\x92\xe5\xbd\x5a\xac\xe6\xda\x5b\x7f\x2d\x2a\x73\x0b\xbe\x40\xb2\x61\xb7\xc0\xf7\x8b\xf0\x89\x3a\x3d\x1d\xa3\x35\xe9\xc4\xd0\xa3\x1f\x58\x44\x84\x0e\x85\x1c\x67\x89\xcb\x7d\x44\x6c\x13\x08\x24\x2c\xf3\xbf\x6f\xad\x5e\x7f\x88\xc8\xdc\x17\x87\x82\x0b\x3c\xaa\xa7\x59\xe2\x9b\x01\x8c\x08\x65\xe3\xbd\x49\xc3\x41\x27\x94\x11\x8e\x28\x21\x10\x16\x76\x27\xf6\x98\x52\x3e\xf8\x9b\xbd\x3e\x91\x50\x32\x98\x42\x73\x1b\x2f\x17\xa4\xe6\x7f\xc7\x52\x1e\xa3\x05\x80\x99\xf3\x18\x3d\xd2\xc1\xfe\xbf\x3f\xea\x63\x04\xdc\xff\x42\x39\x90\x47\x83\xfa\x7f\x4b\xea\x63\x0c\xc8\x9f\x58\xed\x22\x63\x20\xfe\x3d\x09\x80\x63\x24\xc8\x64\x99\x01\xf0\x7e\x18\x90\x4f\x96\x3a\x04\xdc\x0f\x82\xf1\xc9\x12\x67\x06\xe1\xe8\x75\x9f\x9b\x41\x38\x11\x90\x3f\x16\x8c\x9f\x34\x3a\x53\x41\x78\x07\xaf\x27\x34\x23\x11\x80\xef\x43\xeb\x29\x5d\x1c\x05\xdf\x0f\x61\xf5\x34\xd0\x29\x06\xbc\x0f\x42\xea\x09\x62\x87\x41\xf7\x7b\x85\x53\xc9\xd6\x99\x78\x61\x2a\x84\x3e\x1d\x3e\x4f\xe0\x12\x4c\x80\xce\x3d\x30\x3e\x22\xf1\x21\x60\xf3\x24\x8f\x98\x3c\xd3\xd2\x3c\x44\x32\x4c\xfe\x18\x10\xf9\x44\x78\x3c\x25\x2d\x27\x83\xa9\x79\x0c\x1a\xb7\x99\xf0\x88\xc8\x74\x58\xbc\x9d\x0d\x8f\x75\x3f\x15\x12\x6f\xe7\xc3\x63\x95\xa9\x24\x38\xbc\x0f\x76\xa7\x57\x53\x26\x41\xe1\x49\xd6\x9a\x82\xbc\xa6\xc0\xdf\xf7\x06\x55\x47\xc9\xeb\x5c\xb3\x63\x09\xec\x6d\xbb\x0e\xb0\xd8\x07\xdb\x4c\x6f\x05\xcb\x49\x55\x6b\x47\xe5\x9d\xcc\x64\x1f\x94\xfa\x0f\xc5\x6e\xef\xa8\x3e\x4a\x71\x8f\x43\xda\x67\x47\x50\xdc\x83\x12\xdd\xb4\x3c\x82\xe2\x1e\x16\xe9\xa8\xef\x47\x51\xdc\x83\x52\x91\xfa\x7e\x1c\xc5\x7d\x74\xc6\x1f\x9a\x50\x78\xac\x3c\xcf\x3d\x28\x72\x9c\xff\x1e\xe1\xb9\x87\x11\xf2\x28\xff\x3d\xc2\x73\x0f\xab\x33\x99\xff\xde\xe3\xb9\x47\x4c\x7e\xe6\xbf\x1f\x1c\x33\xff\xbd\x75\xcc\xfc\xf7\xc4\xce\xce\xfc\xf7\x99\xff\x3e\x76\xcc\xfc\xf7\x99\xff\x3e\xf3\xdf\x67\xfe\xfb\xcc\x7f\x9f\xf9\xef\x03\xc7\xcc\x7f\x9f\xf9\xef\x33\xff\xbd\x75\xcc\xfc\xf7\x91\xae\xcc\xfc\xf7\x99\xff\x3e\xf3\xdf\x67\xfe\xfb\xcc\x7f\x1f\xb8\xe4\xb3\xf0\xdf\x3b\x20\x74\x90\x04\x1f\x81\x63\xf7\xdf\x4f\x99\x48\x82\x0f\xca\x5c\xc1\x38\x09\x3e\xd8\xec\xa0\xd4\xc0\x37\x7e\x92\x98\xf0\x61\xe8\xb5\xcd\x90\x9f\xc4\x84\x8f\x80\xe6\x03\x5f\xa5\xbf\xe7\xd7\xe7\x49\x8b\x21\x7f\x2c\x13\x3e\x6c\x02\x62\x66\xc2\xcf\x4c\xf8\x99\x09\x3f\x33\xe1\x67\x26\xbc\x3d\x66\x26\xfc\xcc\x84\x9f\x99\xf0\x33\x13\x7e\x66\xc2\xf7\x8e\x99\x09\x3f\xd8\xdc\x99\x09\x3f\x33\xe1\x67\x26\xfc\xfe\x98\x99\xf0\xdd\x63\x66\xc2\xcf\x4c\xf8\xc8\x31\x33\xe1\x1f\x87\x09\x1f\x3c\x45\x39\x17\xda\x06\xf7\x87\xed\x4f\x5b\x4c\x23\x2a\x0a\x3f\xb4\x62\x0a\xe4\x2d\xf4\x32\x95\x58\x6e\xb7\xda\x55\x54\x29\xcc\x20\x90\x21\xfc\x23\xac\xb6\x42\xdc\xfc\x49\xd2\xc1\xa9\x6f\x1f\xbe\x12\xa2\x00\xda\x47\x26\x32\x1a\xbe\x27\x30\xdc\xc0\xe9\xaa\x80\xb7\xb5\x6e\x3f\x7d\xfa\x93\xad\x98\x5e\x37\xa6\x0b\xda\x48\x51\x57\x57\x92\x09\xc9\xf4\xee\x2d\xe3\xac\xac\x07\xb9\xb0\x63\x25\x87\x78\xa1\x61\x0b\xb4\xd0\xdb\x6c\x0b\xd9\x60\x13\xc7\x92\x71\xdb\xdb\xe0\xd4\x88\xf7\x70\xf4\xdd\x0e\x39\x58\x0b\xba\x5f\x87\x8d\x61\x32\xbe\x79\x01\x52\x0f\xf6\x69\xac\xc7\x19\x7d\x31\xdc\x2c\x92\xe2\x4f\x36\xc0\x4d\x0c\x05\xc7\x2a\xcc\xb6\x1f\xe4\x7d\xda\x60\x25\x44\x56\xd4\x11\x09\x31\x7f\xb8\x68\x7a\x38\x75\xb4\x6b\x05\xdf\xd5\x2b\x68\x5c\xc7\x37\x3f\xe7\xfc\x1b\x21\x2f\x6f\x06\xc7\x21\xae\xa7\x5b\x90\x26\xe4\xf5\x93\xe7\xa1\x8d\x28\xa4\x80\x05\xc9\x0e\x93\xd3\xc5\xb0\x5b\x09\x5c\xd5\xf3\x1a\xbd\xeb\x86\x9c\x42\xef\xa2\xd6\x9c\xee\x9d\x33\x93\xaa\xf7\x63\x7b\x4e\xf4\x4e\x86\x07\xa6\x77\xe9\x81\xde\x53\x97\x8b\xcc\x0c\xe2\xb4\xc5\x42\xc2\x86\x29\x2d\x23\x2b\x43\xc0\x7c\x25\x54\x42\x31\x2d\x8e\xb8\x55\xd3\xcd\xc4\x7b\xc2\x86\xe2\xdb\x3f\x70\xc2\xb7\xaf\x77\x4a\xd3\xe4\x05\x38\x93\x4c\xb3\x8c\x16\x97\x79\xde\xaf\xb2\x86\xe7\x8e\xb5\xc2\x4b\x4e\x8b\x9d\x66\x59\x4f\xed\xb1\x1b\x71\x63\x24\xa6\x7a\x8e\x2d\x36\x88\x91\xc5\x23\x3e\xbf\x43\x2b\x43\xdc\xf3\x7f\x0e\x9b\x69\xea\x27\xa3\x6f\x26\xbe\x70\xaf\x4a\x7d\xf0\x77\x34\xd6\x63\xcb\xc0\x40\x14\xcb\x21\xa3\xd2\x67\xcf\x20\x8f\x79\x51\xb0\x60\x25\x1b\x5e\xf8\xc8\x64\x7c\x25\x21\x7a\xee\x74\xf1\xd9\x1b\x7c\xb8\xfb\x71\xe5\x2a\x04\x25\xfd\x64\xbc\x18\xa1\xa5\xa8\x2d\x72\xe1\xde\x1a\x8b\x04\xe4\x5e\x45\x3e\xc8\x27\x6f\x05\xd6\x70\xd7\xe2\x82\x6c\xb5\xae\xd4\xc5\xf9\xf9\x4d\xbd\x02\xc9\x41\x83\x5a\x32\x71\x9e\x8b\x4c\x9d\x67\x82\x67\x50\x69\xfc\xcf\x9a\x6d\x6a\x89\x81\xf1\x79\x49\x39\xdd\xc0\xc2\x3d\x76\xd1\x88\x5f\x34\x9a\x3e\x7f\x16\x5d\x2a\x23\x31\xbd\x7b\xeb\xee\x73\x69\xfc\x83\x7b\xfc\xa1\xce\xed\xca\x71\x94\xce\x65\xf3\x92\xd5\xeb\x35\x69\xe4\x33\x45\x44\xc9\xb4\xc9\xe0\xd6\x42\x12\xba\xb7\xd2\x30\xfa\xcf\xb4\x49\x55\x69\x5d\x68\x04\x37\x9c\x75\xe0\xeb\x7b\xf6\x4d\x4b\xf8\x54\x15\x2c\x63\xba\xd8\xed\x73\xe3\x33\xfb\x02\xed\x1d\x53\xe1\xc6\x5a\x2c\xac\xd9\x13\x1d\x47\x79\xe1\xb3\x62\x4c\x7d\xbf\x58\x8b\x89\x9e\x56\x90\xd5\x66\x61\x7d\x21\xb8\x86\x4f\x83\x2e\xb0\x33\xfc\xd7\xee\x7a\x22\xf0\x07\xd5\x70\x31\x1c\x2e\x22\x6b\x8e\x89\xfd\x31\x8e\x04\xa7\xde\x95\x64\xb7\xac\x80\x0d\xbc\x52\x19\xb5\x35\xa4\x24\x4e\xcf\xb3\xcb\xc0\xdd\x68\x36\x52\x14\x8a\xdc\x6d\x01\xb7\xfc\xa2\xa6\x25\x19\xa8\x70\x79\x22\xa3\x9c\x6c\x28\xe3\x76\xf7\xaa\xca\x0b\x45\x58\x82\x23\xe3\xa4\xa2\x12\xb8\xf6\x82\x5c\x81\xc1\x2c\x2e\x41\x99\x39\x93\x90\x19\xbb\x6b\xda\xd3\xbc\x55\xfa\x13\x87\xbb\x9f\xcc\x53\x14\x59\x17\x74\x63\x59\x42\x2b\x57\xed\x0f\xd7\xc8\x2d\x15\xcd\x59\xc7\xbe\x29\x41\x45\x30\x45\xb4\xac\x81\xd0\xe2\x8e\xee\xc2\x9d\xbf\x73\x74\x99\x96\x6c\xa6\x2e\xc8\xf3\x53\x1c\x5c\xaa\x48\x23\x3b\x27\xbf\x39\xc5\xf7\x61\x5f\x5c\x5e\xfd\x74\xfd\x9f\xd7\x3f\x5d\xbe\x7c\xfb\xfa\x5d\xdc\x4c\x63\x89\x48\x46\x2b\xba\x62\x05\x8b\x79\xac\xde\x6b\xaa\xed\x9b\x70\x9a\xe6\xf9\x79\x2e\x45\x65\xfb\xe1\xd1\xaa\xa6\x2f\x11\x54\xfd\x65\xcb\x73\x98\xfe\x3b\x4f\xe2\x2b\x92\x9d\x07\x6d\x24\xe5\xba\x59\x48\xc3\x86\xd4\xa8\x50\xd6\x5c\xb3\x32\xc8\x52\x4a\x29\x51\xd3\x3c\x5a\xaa\xe9\x56\xf5\xf1\x0d\xdb\x76\x93\x23\x77\x26\xc0\xb0\xdd\xc0\xc2\x8b\xdd\xed\x4b\xbf\xe4\xea\xfd\xf5\xeb\x3f\x1d\x8c\xc6\xae\x8a\x23\x82\x89\xd0\x6e\x0a\xb0\x6b\x86\x3c\x59\x3b\x1f\xa0\x14\xb7\xff\x48\xfa\x19\x0d\x2a\x1a\x1f\x17\x34\xb1\xae\x02\x6b\xde\x76\x0f\xbc\x75\x3f\x29\x91\xac\x78\x65\xdd\x11\xa8\x18\xcb\xa7\x75\xd7\x7e\x82\x62\x8d\xc6\xdc\xca\x35\xb3\xa4\xbf\xce\x9b\x21\x52\x88\x51\xaf\xb8\x15\x4a\x2f\x3b\xf3\x79\x4d\x0b\x15\x9c\x7c\xe3\x9e\xc9\x38\xd7\xb7\x26\xb0\x49\xd2\x4e\x73\x35\xc9\x81\x0b\xcf\x5b\x31\x4f\x41\xf2\x96\x14\x19\xb1\x51\x92\x16\x26\x17\x0e\x76\x65\x8d\x14\x19\x68\x3b\x2f\x74\x79\xde\x31\x31\xe5\xfb\x78\xd5\x3c\x31\xfe\x55\x82\x5a\x35\x9f\x24\x38\x70\x4c\xfb\xb8\x69\x8d\x8c\x0e\x9a\x63\x25\xad\xa2\x7a\xab\xa2\xaf\xef\x96\x54\xdd\x40\x6e\x2f\x74\xeb\xa0\x8b\xe7\xec\x93\x9a\xa6\x7d\x34\xfd\x5f\x03\xd5\xb5\x04\x5c\xe7\x62\xc1\xd6\x0a\x7c\x2e\x17\x1f\xb4\xc8\xdc\x30\x7d\x78\xcf\x8b\xdd\x07\x21\xf4\x37\xac\x00\xcb\x3d\x4d\x1a\xc0\x1f\x5d\xa4\x60\xf9\x67\x8d\xaa\xcc\x52\x47\x51\xee\x02\x95\x83\xa6\xb8\x6e\x44\x8f\xae\x2c\x66\xc0\xee\x69\x88\xb2\xe6\x97\xea\x5b\x29\xea\xa0\xb3\xeb\x2d\x90\xdf\xbe\x7e\x89\xf3\xa6\xb6\xab\x3a\x70\x2d\x77\x96\xe0\xeb\x0a\x25\x4d\x07\x23\xf3\xd4\xc5\x16\xdf\x1b\xfb\x39\xb0\x18\x13\xc7\xd4\x5c\x81\x5e\x92\xb7\x74\x47\x68\xa1\x84\x0f\x5e\x22\x53\xff\x4a\xe4\xd7\xdd\xd8\x73\x89\x6c\x15\x7b\x1b\x59\x09\xbd\x25\x07\x17\x60\x6d\xb8\x7f\x5f\x38\x1b\x68\xea\xc8\xad\x3a\x18\xe3\x3d\xb1\x9a\xde\x80\x22\x95\x84\x0c\x72\xe0\x59\x70\x74\x5a\x00\xdf\x6f\xff\x35\x3a\x82\x31\x0e\x3e\x8e\xe0\x3b\xc1\x8d\x59\xa6\xb1\xd5\x79\xce\x32\xc7\x7e\x73\x54\xb2\xbd\x49\x22\x0d\xc7\xc5\x65\x14\xc9\x38\xc6\x28\x63\xf3\x5f\x5a\xa2\x8e\xac\x1d\x2d\xfd\xbb\x7a\x05\x05\x68\x1b\x74\xde\x5a\xb8\xd0\x7e\x56\x84\x95\x74\x03\x84\x6a\x3f\xe0\xb1\xf9\x0a\x5c\xd5\xd2\x7f\xd0\x46\x93\x5c\x80\xad\x99\xb9\xa6\x7d\xff\xfa\x25\xf9\x8a\x9c\x98\xb6\x9d\xe2\x30\xae\x29\x2b\x62\xdf\x17\x57\x9a\xca\xc3\xbe\xb2\xb5\x17\x8d\x5d\x40\x9b\x23\x42\xda\x29\x75\x46\xb8\x20\xaa\x8e\xf8\x3e\xd7\x37\x13\x09\xfb\x00\xbb\x02\x69\x06\x15\xd3\xfd\x9e\xe9\x86\x4c\x34\xdc\xe6\xe9\xa6\xbb\x37\xd1\xb0\xd4\x87\x30\xdd\x44\xc7\xf2\xbd\xea\x63\xa6\xfe\xe8\xf9\x95\xef\x1f\xd0\xaf\xb4\x97\x6a\x63\xa3\xdd\x5e\x5b\x43\x2c\x41\xd3\x9c\x6a\xea\xfc\x8d\xbf\x20\xec\x75\xd3\x87\x34\x36\x74\xe1\x70\x7c\x6c\x48\xa3\x43\x17\x9e\x4c\x7f\x53\x6f\xa4\xe0\x0d\xe3\xf5\xa7\xf7\xd5\x60\x3d\xd7\x1f\xbd\xb1\xbf\x7e\x85\xb7\xe1\x10\xa3\x1d\xa2\x92\x2d\xad\x24\xf7\xf9\x53\x14\x55\xb4\xc7\xeb\xce\x50\x9e\x05\x42\x13\x9c\xae\xb4\xb0\x7c\x04\xb3\x02\x53\x9e\x8b\xf0\x5b\x24\x87\x8d\x6b\x3e\x82\xb5\x6f\xd0\x14\xe3\xf8\x35\xce\xf7\x94\x74\xb2\x80\x5b\x28\x92\x53\xa6\x37\xe6\x6a\x13\xc0\x78\xed\xe2\xed\x8e\xca\x81\x6e\x7f\x4f\x2a\x8a\xe7\x34\x69\x96\x91\xca\xa4\x10\x45\xb0\xf2\xd9\xeb\xc3\x07\x51\x80\xa5\xdc\xf9\x4e\x98\xdb\x3f\x7b\x1f\xf0\xa2\xd4\x3e\x60\x14\xdd\xe9\x03\xe6\x15\x9f\xbb\x0f\x75\x64\xe5\xe8\xf5\xc1\x2c\x33\xdd\x3e\xa0\xcf\xff\xbc\x7d\x18\x4d\x91\xef\x18\xcf\xc5\x9d\x9a\xea\x2a\x7f\xb4\xb7\xf9\x79\x9d\x19\xb7\xa1\x19\xdf\xa8\xb6\xbb\xa4\x45\x91\x04\x51\x0d\xf9\x4b\x8f\xc4\xe2\xdb\x70\x98\x71\xf5\xfc\x0e\x7a\xd0\xa0\xd0\x15\x18\xfd\x5b\xf4\xfd\x0b\x0e\xbf\x53\x7c\xda\xa6\x54\xf4\x85\x34\x72\x34\xa3\xc5\x75\x05\x59\xb2\x51\x7e\xfb\xf6\xfa\xb2\x7b\xab\x31\x51\xfb\xd2\x98\xe9\x89\x39\x4f\x68\x5e\x32\xe4\xbe\x46\xcd\xf2\xce\x96\xda\xc9\x89\x2f\x03\x6c\x98\xde\xd6\xab\x65\x26\xca\x56\x45\x60\xa1\xd8\x46\x9d\x3b\xab\x5a\x98\x96\xc7\xd9\x83\x8c\x17\xf8\x1a\x9f\x37\xfa\xfd\xe7\x0d\x5d\xe3\xb2\xa6\xf5\xa8\x70\xa4\xfc\xc5\x29\xb9\xae\x0c\xd8\xef\xfa\x3b\x5a\x82\xa5\xf4\xba\x94\xbe\xf9\x7e\x06\x2d\xaa\x2d\x5d\xa0\xf3\x8f\x8a\x36\xc6\xc2\x1c\x1d\x77\x2b\xf0\x55\x5d\xf3\x38\x5b\xf0\x77\xa9\x8c\xcd\xf0\xb1\x09\x6e\x96\x98\x96\x44\xc5\xb6\xf1\x83\x7b\x3b\xad\xbe\xb5\x98\x7e\xdf\xc3\x62\x50\x6d\xee\x0d\x20\xa3\xfd\xf6\xf0\x44\xbb\x75\x38\x74\x36\x0c\x8e\xe8\x7e\xf4\x8b\x6a\x5f\xba\xee\x9b\x7c\x63\x92\xca\x31\xef\x70\x37\x19\x5f\xe2\x9d\xeb\x60\x1e\x12\xed\xcc\x61\x8e\x32\x9c\x8b\x98\x4b\xba\xf9\xc8\xc8\x14\x1d\xc9\x55\x12\xc3\xce\xe8\x43\x1e\xd8\x4b\x93\x87\xf7\xd4\xf6\x88\xda\xae\xc9\xe4\xc3\x26\x3a\xd2\xd8\x41\xf3\xfd\xd0\x36\xa8\x87\xb4\xd5\xfb\x94\x57\x1f\x92\xe0\x33\xc8\xb3\x7a\x64\xd6\xcf\xba\x2e\x0a\xe3\xc8\xde\xdf\x82\x94\x2c\xef\x4d\xd4\x60\x47\x70\x1a\x5c\xd5\x45\x71\x25\x0a\x96\xf5\x58\x2e\xe3\xf7\x5d\x43\x26\xa1\x4f\x72\x08\x94\x62\x46\x29\xc5\xfd\xe2\x48\x21\x36\x6f\x86\x32\xa0\x18\x8d\x2f\x9c\x4f\x97\x82\x1b\x65\x33\xde\x1b\xef\x58\xd8\x42\x37\x30\x5c\x52\x48\xa0\x15\x1f\x43\x72\x92\xa2\x04\xbd\x85\xfa\x28\x9a\x6a\xf3\x06\xc1\x91\x24\xcf\x11\x1e\x82\xbc\x65\x19\xbc\xb5\x6a\x3c\xa6\x79\x45\xf4\x85\xaf\x07\x27\xc5\x1c\xef\x15\xc2\x33\x1c\xcd\xa1\x3f\xef\x9b\x51\xeb\x9d\xea\x6a\x2d\x75\x52\x1f\x35\xa1\xb9\xc8\x21\xf4\x25\x80\xc7\xa3\xfd\x87\xde\x4f\xfe\xd2\x88\x9c\x33\x29\x6f\x26\xe5\xcd\xa4\xbc\xc0\x31\x93\xf2\x26\x9e\x9e\x49\x79\x43\xc7\x4c\xca\x9b\x49\x79\x33\x29\x2f\x28\xfc\xcb\x23\x9d\xcd\xa4\xbc\x99\x94\xe7\xbb\x3a\x93\xf2\x66\x52\x1e\x99\x49\x79\x33\x29\x6f\x26\xe5\xf5\x8e\x99\x94\x37\x93\xf2\x66\x52\xde\x4c\xca\x6b\x8e\x99\x94\xf7\x77\x39\xdf\x67\x52\xde\x50\x1f\x66\x52\xde\xa3\xf5\x61\x26\xe5\x0d\x75\x74\x26\xe5\xcd\xa4\xbc\x99\x94\xf7\xc5\x12\xc3\x66\x52\xde\x4c\xca\x9b\x49\x79\x33\x29\xef\xd7\x41\xca\x7b\x64\xfe\x5d\xd5\xb3\x99\xc3\x26\x74\xeb\xad\x03\x51\xc6\x56\x14\x39\x7e\xba\xd6\x8d\xb3\x2f\x2e\x13\xaa\xb5\x64\xab\x5a\x0f\x54\x59\x8c\x2d\x64\xa2\x2c\x45\xbb\xa0\xe0\x43\xa4\x25\xb1\xd1\x16\x2d\x2e\x3a\xd3\xd2\x7d\x6a\x9d\x5c\x03\x0c\x17\x51\x5a\x4d\xc5\xf4\xcf\x43\x95\xee\x53\xcf\x62\x6d\x13\x42\xbb\xc2\x1d\x16\x2c\x63\x81\xc6\x3a\x0c\xbe\x76\xd4\xf3\xf4\xd2\x4e\x25\xe3\xd0\xeb\xca\xb3\x06\x0a\xfb\x09\xbc\xc3\x18\xf7\x20\xfc\x1b\xb4\x43\xc6\xed\xf6\x2c\x4b\x72\x2d\x4a\x20\xb7\xa2\xa8\x4b\xdb\x79\xc7\x59\xe9\x80\x79\x5a\x90\x6c\x8b\xdb\x7e\x61\x84\x78\x67\xc4\x86\xb6\x47\x10\x8e\x1d\xe1\x45\xa2\x6f\x32\xb7\x34\x34\xa1\x4a\xe4\x17\xe4\xbf\x39\x79\x6e\xcb\x0f\xe2\x0e\x4b\xaa\xdf\xbe\x7e\x19\x8e\x2b\x57\xf6\xc9\xdf\x5c\xa3\xba\xc8\x6f\xec\x9d\x0a\xf4\x86\xe5\x64\x65\x27\xbf\xf1\x62\x27\x1c\xee\x2c\x84\x6e\xd6\x40\xfb\x0d\xe0\xe1\xe0\x0a\x7d\x94\x6d\xa2\x87\xef\x9a\x46\xba\xc7\x9c\x92\xaf\xed\x73\x2a\x90\x2e\x4e\x33\xcf\x0a\x6f\x89\xfe\xfe\xc3\x33\xb7\xe9\x9a\xbc\x5b\xc8\xbb\xc5\x62\xb1\x30\xfd\xf4\xe0\xe2\x00\x40\x8a\x3b\x6f\x89\x9c\xad\xc3\x75\xdf\x46\xdb\x68\xdb\xfb\xa6\x28\xbf\xe9\x8e\xed\xc5\x72\xe8\x7b\xd5\x63\xa0\xce\xc8\x37\xe2\xa2\xc5\x81\x7b\x14\x06\x9a\xc5\x71\xb0\xc3\x93\x8b\x02\x87\x0b\x4b\x18\x66\x79\x78\x88\x65\xd2\x02\xe7\x2a\x70\xfb\x6d\x66\x86\xb3\xcc\x07\x18\xb5\x48\x41\xe0\x51\x8a\x01\x0f\x5f\x08\x38\xbe\x08\x60\xc1\xfe\xe0\xa4\x9f\x5a\x00\x68\x01\xfd\x81\x3c\x3e\x05\xfc\x0f\xc1\xc4\x21\xe7\x7c\x9c\x89\x8e\xc4\x95\x47\xc6\x60\x71\xb4\x3f\x8a\xf4\xdf\x03\xe5\x8f\x3b\x89\xfb\x20\xfc\x66\x7c\x06\x85\x26\x8e\xd9\x24\x68\xff\x11\x60\xfd\xcf\xe2\x56\xc6\x91\xfd\xe9\xa8\x7e\x0a\x4c\x75\x1f\x48\xdf\xb7\x60\x50\xf0\x54\x38\xff\xef\x6c\x91\x19\xa5\x61\xc7\xe0\xfc\xa3\xa1\xfc\x34\x76\xdb\xf1\x1c\x8d\x08\x7c\x7f\x34\x74\xff\xc8\x6d\x8e\xc1\xf5\x47\x43\xf5\x8f\xdc\xe6\x18\x3c\x7f\x34\x34\xff\xa8\x6d\x8e\xb3\x96\x5b\x29\x15\xc6\xbb\xe3\xfe\xed\xb2\xd9\x36\x11\x53\x30\x75\x58\xac\x5c\x33\xa9\x1a\x3e\x2f\x2e\x77\x81\x3c\xa4\xeb\x7a\x70\x93\x65\xcf\xb8\xef\x15\x3e\x9f\x99\x79\xce\x4a\x2a\x77\x26\xd8\x0e\x7b\xa0\x8e\xc3\xe4\xc2\x37\xd1\x47\x2a\x14\xc9\x9c\x48\x41\xdf\xc5\x15\x1b\x61\x2a\x8e\xd7\x8b\xc7\xaa\xc5\x31\x9a\xa1\xda\xa9\x4c\x0f\xbf\x13\xd5\xe5\x8f\xdb\xeb\x10\x2a\x68\x6d\x65\xd9\xbc\xf0\x95\x7b\x49\x58\x92\x30\x8d\x0e\xc6\x85\x98\x0c\x5f\x89\x5c\xd9\x0c\xae\xe6\xc6\x2a\x84\xd4\x2d\x19\x27\x2e\x81\xed\xad\x3f\xc3\x48\x7b\x89\xfb\x71\xfb\xa8\xb4\xa0\x35\x1f\xde\x71\x2d\xa2\xe5\x81\xce\xba\x1d\x6d\xec\x86\x85\x92\x43\x41\x2a\x2a\x69\x09\x1a\xa4\x5b\x61\x43\x61\xe5\x78\xc5\x83\x47\xc1\xd3\x4e\x63\xde\x39\x3c\x9a\x7a\xb1\xb8\xe1\x71\xe8\xd1\x24\x0d\x1b\xc3\x25\x30\xb1\x01\x3f\xf8\x3d\x6c\x1f\xb0\x05\xf1\xdd\x6d\x16\xa8\x9f\xc0\xa9\xf0\xea\x9d\x04\xf8\x0d\x4f\x83\xf1\x1a\xe1\xe3\xd5\x07\x63\xb5\x41\x33\x43\x10\x41\x6a\xbb\xa6\x94\xd0\xd5\xbb\xa0\x07\x2f\x10\x3e\x20\xe4\x3c\x36\x4d\xd2\x8b\x82\x8f\x51\x10\x7c\xf8\x62\xe0\x71\x85\xc0\xf8\xf6\xf2\x0f\x5f\x04\x7c\x84\x02\x60\x0a\xa8\x3f\xbe\x89\xcd\xa4\xa2\xdf\x63\x14\xfc\x8e\x2a\xf6\x79\x5d\x06\xa5\x4e\xd3\xf1\xc3\xe8\x32\xa9\x88\x77\x7c\x01\x8f\x88\x30\xf9\xea\xa8\xe2\x5d\x03\x35\x04\xc5\x3e\x56\xe1\xee\x38\xcf\x19\x8d\xb0\x8f\xf1\x9e\x68\x63\xe1\xf9\x75\x54\xb1\x2e\xbe\xb5\xf6\x43\x14\xea\x8e\x4f\x15\x82\xa7\x24\x54\x05\xcb\xe8\x8b\xa1\x77\x3a\x8e\xfb\x2c\x83\x7b\x41\xfe\x32\xcb\x86\x64\x46\x3f\xcd\x10\xde\xe5\x8e\x4c\x7a\x9f\xf7\x3e\x29\x95\xad\xbe\x4c\xff\xd0\x43\x28\x02\x3d\xaa\xc4\x69\x1b\x91\x3a\x86\x5a\x14\x20\x87\xf5\xd6\xc5\xcc\xd7\xe4\x20\x2c\xb2\x3b\x87\xb6\xee\x3f\xb4\xbc\x40\x84\xdf\x73\x65\x95\xc8\xed\x4b\x22\x1f\x1b\x59\x38\x7f\xb4\xa6\xd9\xd6\x25\x97\xf6\x8c\x89\xff\x07\x37\x98\x34\x1e\x4c\x5b\x57\xbd\xdf\x91\x12\x88\x96\xac\x2a\x80\xfc\xbe\xd9\x6b\xf7\x0c\xd6\x6b\xc8\xf4\x1f\x48\xad\x18\xdf\xb8\xf7\xdc\x83\x5b\x76\x36\xbb\xdd\xfe\xde\xff\xef\x0f\xfd\xd9\x15\x0f\x9c\xec\xf3\x12\xd2\x9c\x57\x78\x21\x61\xad\xd2\x04\xb8\x6e\x59\x19\x46\x0d\xd8\xd6\xf8\x7e\x98\x76\x6f\x57\xbc\x10\x03\xde\xb6\x08\xb5\x24\x3f\x6e\x81\xb7\x07\xd2\xbd\xa9\x1f\xdf\x53\x9c\x4a\x20\xef\xc4\xb5\x19\x8c\xba\x80\x33\x72\x25\x61\x0d\x72\xff\x0b\xba\xb7\x77\xe2\xd5\x27\xc8\x6a\x1d\x20\x34\x8c\x4c\xab\xe0\x16\xfe\x1d\x25\x7d\xb7\xdf\xb0\xdf\xf6\xab\xb3\x61\xff\xde\x14\x3d\xbe\x12\x72\x85\xc2\xe9\x30\xa0\xad\x1b\xd8\xa9\x66\x7b\xf4\x1b\xfb\x4c\xac\x96\x9f\x8d\xed\x96\xed\x77\x59\xb7\xfb\xa3\xff\xce\xbf\xfb\x54\xae\x18\xb7\x0d\xb3\x0f\xf4\x43\x89\xcf\xf4\x3b\x27\x07\xb9\x48\xe6\x22\x6c\xd2\x31\x8a\x0d\x7d\x91\x63\x40\xbb\xef\xbd\xb9\x37\xef\x5b\xda\x74\x7b\xf7\x4c\x11\x09\xf6\xfd\x67\x2c\xc9\xba\xb0\xc0\xbe\xb1\x1f\x68\x34\x6e\x64\xd7\x3c\xdd\xbe\x8f\xd8\xda\xd4\xff\xd5\xcf\x35\x2d\xba\x91\x86\xfb\xc9\x5e\x14\x90\xda\xd9\x12\xdb\xdc\x74\xc7\x8a\x3c\xa3\xd2\x7e\xdd\xc0\x4e\x71\xa2\x84\x03\xd9\xd0\xb3\x64\x94\x37\xee\x23\xa2\x60\x1c\x79\xe5\x32\x6b\x2a\x35\xcb\xea\x82\x4a\x62\xe6\xe2\x46\xc8\xdd\x51\xba\xdf\x1b\xe4\x35\x64\x82\xe7\x29\x70\xc7\xc7\xc3\x7b\xda\xa3\xa1\x6d\x59\x9e\x89\x1c\x63\x64\x56\x42\x24\xb8\x69\x4d\x87\x13\xbb\xd9\xae\xb7\x4e\xb1\xf6\x3e\xa5\x99\xb4\xad\xcf\x37\xa0\xd1\x06\x64\x36\xc1\x0d\xdb\x60\x04\x73\xda\xf2\xcc\xcd\xac\x5c\x92\x3f\xee\x7c\x09\xfb\xcc\x85\x3d\x3c\xf8\x3e\x15\x52\x04\x5c\xfb\xdc\xe4\xb0\x12\x5b\xd3\x7c\x2d\x24\xdc\x82\x24\x27\xb9\xc0\x5a\x28\xdc\xb2\x4c\x9f\x86\x4c\xef\xbf\x40\x0a\x34\x32\x0e\x1b\xaa\xd9\x6d\xb3\x39\xb9\xcf\xc3\xb5\x23\x49\x50\x45\xbe\x22\x27\x28\x8c\xb0\xb2\x84\x9c\x51\x0d\x45\x70\x53\x68\xff\xbd\x97\xc8\x7b\x8a\xf7\xc7\x07\x23\x60\xd0\x00\x10\xd4\x71\x86\x36\x96\x3d\xf0\x84\xcd\x72\x28\x42\xea\x72\x7e\xae\xbd\xb5\xbb\x9d\x83\x9d\x72\x5d\xb3\x8d\xb6\x77\x84\x23\x5f\xfb\xf8\xab\xb1\x35\x4a\x24\x6c\x70\x1e\xd9\x39\x72\xc4\x2c\x1a\x8d\x4b\x0f\x31\xa4\xa1\xc0\x68\xb1\xdf\x8f\xb8\xf3\xab\xdb\x76\xb2\xf3\xdb\x7e\x23\xc3\xce\xcf\x07\x1f\x48\xeb\x9c\xdb\x7f\x4c\xac\xf3\xf3\xc0\x5a\xb1\xe8\xc4\xcc\x9d\x13\xdd\xd0\xf7\x49\x54\x03\x07\x3f\xb9\x4d\x37\x2f\xc8\xed\x73\x4c\x3a\x9e\xef\x7f\x43\x97\x63\x21\xbc\xce\x69\xb7\x07\x6c\x7e\x81\xdc\x05\xfb\x83\x16\x92\x6e\xc0\xfd\xf2\xff\x01\x00\x00\xff\xff\xd6\x35\x4d\xf8\x6f\x0e\x01\x00")

func installerKubedbCom_kubedboperatorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_installerKubedbCom_kubedboperatorsYaml,
		"installer.kubedb.com_kubedboperators.yaml",
	)
}

func installerKubedbCom_kubedboperatorsYaml() (*asset, error) {
	bytes, err := installerKubedbCom_kubedboperatorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.kubedb.com_kubedboperators.yaml", size: 69231, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"installer.kubedb.com_kubedbcatalogs.yaml":  installerKubedbCom_kubedbcatalogsYaml,
	"installer.kubedb.com_kubedboperators.yaml": installerKubedbCom_kubedboperatorsYaml,
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
	"installer.kubedb.com_kubedbcatalogs.yaml":  {installerKubedbCom_kubedbcatalogsYaml, map[string]*bintree{}},
	"installer.kubedb.com_kubedboperators.yaml": {installerKubedbCom_kubedboperatorsYaml, map[string]*bintree{}},
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
