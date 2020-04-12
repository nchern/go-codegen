// Code generated by go-bindata.
// sources:
// pkg/generic/list/list.go
// pkg/generic/set/set.go
// pkg/generic/hashmap/hashmap.go
// pkg/generic/iterator/iterator.go
// pkg/generic/converter/converter.go
// DO NOT EDIT!

package generic

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

var _listListGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x82\xcd\x61\x21\x15\x8a\x92\x5e\x83\xe4\x10\x2c\xd6\x8b\x02\xe9\xa2\x88\xb3\xdd\x43\x11\x2c\x14\x99\x4a\xd8\xd0\xa4\x41\x52\xee\x7a\x03\xff\xf7\x82\x1f\x12\x29\x89\x94\x3f\xe2\x93\x2d\x72\x66\xde\x9b\xe1\x90\xf3\x56\x65\xf5\x5a\x3e\x23\x48\xb0\x90\x00\xe0\xe5\x8a\x71\x09\x53\x90\x9c\x09\xc6\xe5\x99\xfa\xdd\xd0\xea\x0c\x64\x00\xc8\xcd\x0a\xc1\x87\x4b\x28\x24\xc7\xf4\xb9\xfb\xfe\xca\xd1\x02\x57\xa5\x44\xb0\x6e\x68\x95\x3e\x5c\x66\xf0\x89\x31\xd2\xed\x7f\xc3\x02\x4b\xc6\xcd\x2e\x86\x98\xca\x1c\xae\x4b\x02\x47\x86\x77\x48\x88\x59\x43\x2b\x63\x59\x63\x2e\x64\x0e\x05\xaa\x18\x5d\x04\x8c\xb1\x90\x2a\x16\xe2\x75\x59\x21\xf8\x06\x92\x19\x26\x12\xf1\xb4\xf6\x29\x65\xd6\x12\x24\x9f\xc3\xdb\xd7\xe7\xd5\x4b\x49\xe1\xc3\x25\x48\x3e\xab\xed\xde\xca\x1f\x65\xf5\x92\xae\x2d\xfd\x2e\x11\x17\xf3\x4f\x24\x4d\x42\x99\x36\xbf\xa5\x9b\x61\x78\x4d\x39\xb9\x25\x24\xbc\x31\xc3\x74\x31\x53\x79\xf6\xb7\x73\xb8\x40\x75\xd9\x10\xf9\xcd\x96\x49\x45\x57\xb6\x77\xe5\x5e\xa6\x77\x88\xa6\x99\xe2\x05\x92\xf9\xcf\x72\x95\xe2\x1c\xfe\xd0\x34\x41\x32\x67\x5c\xa6\x4f\x1b\x5d\x66\x57\x71\x97\xd2\xef\x84\x51\x94\xba\xef\x79\x9b\x62\x77\x66\x20\xb9\x5d\xad\x10\x5d\xa4\x58\xa2\xa5\x80\x45\x51\xe8\xc5\xaf\x1c\x85\x56\xd9\x2a\x1d\x13\xdc\xda\x53\x7c\x2a\x05\xb2\x27\x29\x24\x6f\x2a\xa9\x8e\x51\xf5\x21\xfc\xfe\x68\xcc\x54\x27\xc0\xbf\xd0\x4f\x63\xd5\xf1\x52\x76\x1c\xc9\x86\x53\xb7\x39\xe3\x6c\x39\x27\xb8\x42\xa9\x72\x7e\xdb\x16\x45\x91\x8d\x43\x38\x2b\x9f\xa8\x17\x96\xc0\xab\x1b\xb8\x2c\x5f\x4d\x98\x1c\x12\x44\x8d\x69\x96\x81\xa4\x62\xab\x4d\x4a\x72\x68\x16\x3a\x0e\xbf\xb9\x44\xde\x14\xfd\x2b\x48\xb6\x3e\xf4\x7c\x43\x2b\xce\x28\xfe\x1f\x2d\x6c\x22\xa7\x00\xc7\x94\x22\xae\x1c\x82\xf0\x1d\x37\x75\x7d\xed\xa6\xf6\xb8\x82\xfa\xc7\x11\x4c\x09\xfc\xe8\x22\x64\x70\xea\x1e\x99\xc2\x8b\x86\xc8\x30\xb0\x29\xfd\x16\x24\x35\xe3\xf0\xdf\x1c\xae\x95\x19\x2f\xa9\x7a\x5f\x0a\x62\x03\x24\xb8\x86\x75\xba\xce\xf4\x7f\x1b\xce\x6c\xde\xc0\xd2\xf4\x96\xb7\x98\xc3\x75\x06\x92\x64\x0b\x12\x97\x94\xd9\x8e\x66\xb0\xeb\xae\xbb\x2c\x44\x57\x70\xbb\x95\x81\xe4\x99\x99\x07\xc8\xf0\x9b\x4c\xa4\x9f\x49\x17\xf3\xfa\x1c\xae\xd5\xc2\xd6\xf0\x4e\x2a\xc2\x04\xb2\x39\xa9\x93\xdb\xa6\x19\x18\xe4\x22\xe2\xc9\x0c\x9e\xa5\xd3\x91\x1f\xd0\x7d\x37\xd3\xe9\xe7\x52\x61\x2a\x3e\x78\xa2\x2b\x3e\x58\x6f\xf5\x62\xb9\xfe\xd0\xe0\x64\xd8\x03\x24\xca\xa3\xf7\x2c\x7b\x4f\x85\xc1\xfa\x8e\x1f\xa3\x9e\x91\x17\xbc\xa5\xbe\x6f\x43\x6b\x34\xc9\x1b\x34\xe4\x5c\x97\x44\xa0\x38\x7a\x78\x4c\xec\x83\xfe\x61\x04\x6f\xa0\x06\xf8\x9a\x53\xfc\xde\xef\x3b\x8c\x8e\xaa\xc7\x7a\x48\xc6\x85\x9d\xa4\xb4\xd7\xcc\x33\x8c\x1a\xba\x50\x74\xbc\xc0\x07\xd0\x34\xee\x37\x63\x9e\x7a\x43\x51\xbc\xb8\x80\x33\xc6\xa1\xd2\x45\x11\xbe\xdd\xe0\xf5\xbb\x0e\xd1\xd4\x60\x66\xd1\x3c\xfb\x53\x5a\x4f\x82\xb6\x57\xf3\xb6\x6d\x7f\x3c\xc2\x1b\xf7\x3f\xdf\xa3\x9b\xa7\x67\xbd\x42\x51\xa9\xd8\x19\x62\x16\xe7\x7a\xe1\x8d\xe4\xd0\xf8\x6d\x8d\x4d\xa1\x43\x19\x6b\xef\x41\x88\x1f\x5c\x5f\x46\x28\xa8\x8a\xad\x30\x5a\x04\xe6\x9b\xad\x4d\x3b\xe0\x8c\x5d\x9b\x9e\x06\x53\x4e\x81\x01\x6e\x2c\xf5\x98\x77\x4f\x94\x3d\xa7\x2f\x8d\x2c\x25\xe3\x22\x56\x99\x91\xaa\xe9\xd5\x5c\x75\xc1\x44\x57\x06\xf4\x8f\x73\x77\x33\x8c\xd8\xf1\xa5\xed\x7a\x6a\x64\x18\x30\xa4\x9d\x42\x11\xf5\x7e\x5b\x9a\xe9\x88\x21\xdd\xa5\x42\xe2\xda\x2f\x3a\xbc\x86\x9f\x74\xf7\x8f\xaf\x64\xb2\x05\xfe\xa8\xb7\xb5\xf1\x7c\xcf\x3f\x3d\x02\x8f\xa2\x35\xb8\x1a\x59\x8c\x46\xf6\xc5\x85\xa2\x77\x8f\x9f\x5f\x64\x4f\xcd\x9b\xde\xf3\x94\xa0\x97\x11\x48\x88\x6a\x5f\xea\xf5\xb1\xcb\x5d\xc0\x8f\x7e\x04\x75\x0d\x85\xf0\xae\x53\xfb\x88\x5a\x26\xa2\x30\xb1\x52\xe1\xee\x98\x68\xef\x55\xd6\x89\x53\xa7\x9b\x3c\x4a\x46\x73\xf9\xc4\x40\xb2\x6c\x24\xfa\x4f\x9b\x17\xf7\xff\x7c\x51\x1f\xbd\x63\x71\x71\x76\x4b\x2b\x52\xe8\x60\xc5\xfd\x1d\xab\x5e\xd3\xcc\xf6\x3e\x29\x34\x6c\xd1\x7a\x67\x9e\xe1\xdf\x94\xb4\xa6\xde\x0d\x08\x82\x9f\x52\x15\x8d\x88\x06\x1f\x5a\x43\xfb\x50\xb5\x14\xc8\xed\x70\x5d\xd2\x4b\xfc\x1d\x0a\xea\xb8\x44\x03\xca\xea\xe4\x59\xed\x56\x5b\x23\xee\x2d\x51\xdf\x75\xb2\x99\x48\x14\x7d\xa4\xb1\x46\x60\x6b\xbf\x75\xb5\xf9\x24\xd4\x3a\x0a\x35\x21\xca\xa6\xef\x8b\x76\x3c\xf2\xb2\x4c\x68\xb1\x1d\xa0\xca\xf1\x48\xd0\xc3\x14\xd8\xae\xc7\xa2\x8b\xe5\xfb\xbf\x83\xd9\xde\x42\x6c\x37\x31\x13\xea\x04\xbc\x7a\x82\x6b\x1a\x57\x9b\x1e\x87\x62\xe4\x99\xd1\x0b\xbe\x44\x33\x91\x86\xb7\xab\x13\x73\x1e\x5a\x07\x16\x85\xd8\xa9\xd5\xa2\x70\xce\x33\x04\xb8\xfb\x2a\x8f\xb5\xda\x74\x21\xad\xfd\x91\xa5\x8c\x28\xaf\x48\x6e\xca\x5a\x5b\x1e\x52\xcb\xa8\x3e\x0b\x83\xf8\xe6\x5a\x54\xed\x8f\x14\x17\x6e\x61\xa8\x9e\xfd\xa1\x58\x31\x49\x37\x04\xeb\x1d\x56\xdf\x6b\xaa\x41\xd4\x99\xfd\x0a\x00\x00\xff\xff\xfe\x58\xb8\xc1\x0c\x16\x00\x00")

func listListGoBytes() ([]byte, error) {
	return bindataRead(
		_listListGo,
		"list/list.go",
	)
}

func listListGo() (*asset, error) {
	bytes, err := listListGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "list/list.go", size: 5644, mode: os.FileMode(420), modTime: time.Unix(1573317568, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _setSetGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x4d\x8b\x1b\x31\x0c\x3d\x5b\xbf\x42\xa7\x62\x97\xac\xc9\x79\xd9\x14\x4a\x4f\x85\xd2\x43\x37\x3d\x85\x50\x6c\x47\x49\x87\x1d\xdb\xc1\xf6\xa4\x13\x86\xf9\xef\x45\xce\xe4\x83\x76\xa0\xf4\x66\x0b\xbd\xf7\xf4\xf4\x74\x34\xee\xcd\x1c\x08\x33\x15\x80\x72\x3e\x12\xae\x97\x98\x4b\x6a\xc2\xe1\xf6\x7f\xa5\xc2\xa5\xce\x15\x1c\x40\x78\xf4\xe6\xb8\x59\x2f\xb7\x36\xc6\x16\x46\x80\x7d\x17\x1c\x7e\xa5\x5f\xb5\x51\x9e\x4c\x9b\x51\x6b\xbd\x5e\x2a\x7c\x7f\xc1\x0e\x20\x12\x65\x7c\x5e\xe1\xbb\x5a\x18\xfc\xf3\x23\xc7\x30\x8e\xb5\x41\x7f\xdc\xed\x2a\x5c\x6b\xad\xb8\x52\xba\x14\x30\x51\xbe\x89\xc8\x3c\x51\x2a\xbc\xf6\x5e\xa5\x06\x10\xfb\x98\xf0\xc7\x02\x4f\x2c\x94\x4c\x38\x10\xd6\x86\x01\x84\xc8\xda\x6f\x4e\x5b\x5c\x61\x49\x1d\x81\x18\xe7\x18\xbf\x91\x8f\x27\x92\x3d\x4e\x74\x3b\x6a\xa9\x90\xcc\xda\x2f\xb0\x57\x73\x88\x2f\x14\xa4\xc2\x26\x4c\x0e\xeb\xb8\x2d\x05\x86\xcc\xf6\x7f\x8a\xa1\x98\x26\xe4\x49\x83\xbd\x3f\x20\x79\xc6\x7e\x3b\xeb\xb5\x6d\xa5\xc2\x97\x27\xf7\xd3\x04\x8e\x67\x00\xe1\xd8\xa4\x37\x6f\x24\xa7\xa2\x02\x71\x88\xc8\x48\x59\xa7\xaf\xdb\xe8\xef\xab\xc8\xda\xd7\xb2\x70\xf8\xf2\x84\x3d\x08\x31\x82\x10\xae\x8d\x99\xa4\x53\x20\x46\x79\x5f\xb9\x9b\x1b\xe2\x7b\x68\x62\x90\xf6\xf6\xbf\x67\xfb\x87\x92\x9d\x94\x2e\x79\xf6\xcc\x7d\xf7\x38\xc7\xfc\xda\xd9\x92\x8c\x2b\xff\x47\x7e\x8d\xeb\xdf\xfc\x9f\x43\xa1\x94\xc9\x15\x79\x9e\x11\x30\x4c\x9e\xb5\x07\x61\xf9\x75\xe6\x57\xb3\xaf\x39\x1a\x85\x1f\xea\xc3\x5e\x56\x6a\x16\x68\x71\x85\x76\x81\x66\x12\xcd\xb7\x18\x1e\xee\x59\xfd\x35\xb5\xa9\xf0\x66\x8f\x76\xd3\x6f\x2f\x31\x24\xca\xfc\xbe\x5e\x24\xd3\x8d\xc0\x77\x8a\xab\x7a\xf2\x8f\x8e\x7e\x07\x00\x00\xff\xff\xc8\x85\x69\x4d\xa4\x03\x00\x00")

func setSetGoBytes() ([]byte, error) {
	return bindataRead(
		_setSetGo,
		"set/set.go",
	)
}

func setSetGo() (*asset, error) {
	bytes, err := setSetGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "set/set.go", size: 932, mode: os.FileMode(420), modTime: time.Unix(1532984840, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _hashmapHashmapGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\xc1\x8e\x9b\x30\x10\x3d\xdb\x5f\x31\xdd\xc3\x0a\x2a\x84\x92\x6b\xa4\x3d\x55\x55\x2f\xdd\x1e\xb2\x6c\x7b\x58\xad\x2a\x2f\x99\x24\x88\xd8\x20\xe3\xd0\xd2\x28\xff\x5e\x8d\x63\x82\x21\x40\x8e\xe0\x37\xef\xbd\x99\x79\x53\x8a\x34\x17\x3b\x84\xbd\xa8\xf6\x52\x94\x9c\x67\xb2\x2c\xb4\x81\x87\xaa\x51\xe9\x03\xe7\xa6\x29\x11\x92\x05\x54\x46\x67\x6a\xe7\x3e\x97\x90\x29\x73\x7d\x4b\x96\xcf\xa2\xfc\x99\x55\x99\x29\x34\x6c\x8f\x2a\x0d\x92\x45\x04\xc9\x32\x84\x8f\xa2\x38\xf4\x61\x54\x88\x7a\x2b\x52\x84\x13\x67\x5f\x45\xba\x0f\x6a\x57\xd9\x27\x0a\x39\xfb\x86\x26\xc8\xb1\x81\x64\x11\x42\x50\x43\xb2\x8c\x60\x5b\x1c\xd5\xc6\xd2\x86\x9c\xbd\x5c\xdf\x23\xa8\xc5\x81\x14\x39\x7b\x2d\x37\xc2\x60\x50\xe9\x14\xa4\x28\xdf\x92\xc5\x3b\x19\x71\xdc\x9c\xad\x51\x16\x35\x5e\x69\xad\x41\xf6\xe5\x50\x28\x0c\x3a\xd4\xd9\x79\xfe\x10\x15\xb6\xbe\x2b\xa3\x8f\xa9\x21\xd3\xbf\xa5\x28\x3b\x6e\x02\x53\xcf\xf0\x03\xff\x38\x6c\xc7\x44\x70\x8d\x15\xac\x9e\xe0\xd1\x23\x3b\x71\x66\x59\x56\x1d\xcd\xe9\x1c\x71\x76\x26\xb4\x39\x6a\x05\x1a\xab\x11\xe2\x97\x46\xa5\xba\x50\xd9\x3f\xdc\x0c\x35\x6c\xd5\x23\xed\xcc\xd3\xc8\x94\x42\xbd\xea\x39\xb3\x2a\x2d\x73\x20\xe1\xb3\x67\x2b\x84\xf9\x89\x93\x52\xdd\xfe\x79\x02\x19\x53\x0f\x6f\x39\x36\xef\xad\x83\x69\xe6\xb9\x4d\x13\xef\xb6\xd0\x90\x47\x50\xd3\xa8\xb4\x50\x3b\x74\xf4\x60\xfb\xd8\xc2\x27\x57\x1b\x10\xc8\x56\xb0\x56\x93\xd1\xdc\x66\x9a\x1a\x89\x09\xd5\x7b\xf6\xe1\x89\x1e\xa6\x19\x66\x33\x35\xee\x9e\xb0\x64\xb2\x55\xb1\x1a\xfe\x86\xe5\xb4\xdc\x48\x46\x6d\xee\xda\xd1\xaf\x06\xb3\xdf\xe0\x01\x0d\x06\x97\x7f\x11\xe4\xd8\x84\xfc\xaa\x63\x4b\xa6\xb5\x06\xd1\xf7\x02\xeb\xa7\xe6\xce\x7e\x34\x56\xb1\x9d\x32\xed\x86\x9a\x1c\xe6\xd8\x5e\x93\x17\x4e\xef\x9a\x6c\x46\xaf\x97\xc7\x99\x3c\x1a\xfc\x6b\xb1\xf1\xfa\xd7\x33\x7d\xf4\xcc\x7b\x24\xf7\x43\x25\x63\xcb\x16\xaf\xbf\x17\x69\x4e\x5d\xc8\xd8\xca\xc5\x7e\x61\xe8\xe1\x5e\xd5\xe1\x82\x9c\x94\xbc\x7f\x21\x37\xa2\xbd\x93\xb9\xe8\x3b\x96\x51\xe9\xb1\x53\xea\x59\x98\xcc\xf3\x85\x69\xd8\xab\x83\x5b\xac\x27\x78\xbf\xd5\xbb\xa9\x9f\x52\xec\x0a\xc7\xf4\x26\x2e\xa0\x27\x3d\x71\x01\x43\x41\xef\x1c\x2e\xca\x5d\xdd\xac\xf2\xed\x4d\xf4\xd4\x47\x6e\xe2\x66\xa9\xda\x97\x75\x05\x33\xdb\x04\xcd\xcf\xfc\x7f\x00\x00\x00\xff\xff\x0f\x35\x0c\xd6\x6b\x07\x00\x00")

func hashmapHashmapGoBytes() ([]byte, error) {
	return bindataRead(
		_hashmapHashmapGo,
		"hashmap/hashmap.go",
	)
}

func hashmapHashmapGo() (*asset, error) {
	bytes, err := hashmapHashmapGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hashmap/hashmap.go", size: 1899, mode: os.FileMode(420), modTime: time.Unix(1586685795, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _iteratorIteratorGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x31\x6f\xdb\x30\x10\x85\x67\xf2\x57\xbc\xa9\x20\x8b\xb8\xf1\x1c\x38\x63\x5a\x64\xe9\x52\x6f\x45\x51\x10\xcc\x89\x21\x62\x9f\x8c\x13\x15\xb4\x08\xf4\xdf\x8b\x93\x68\xda\x55\x97\x22\x93\x40\xea\xdd\x77\xf7\x78\xef\x14\xe2\x4b\x48\x84\x5c\x48\x42\xe9\xc5\xda\xf2\xfb\x44\xd8\x6f\x31\x14\xc9\x9c\xda\xf9\x0b\xf1\xa2\xf8\x3c\x72\x44\x37\x72\x74\xe9\x7c\x85\xf8\x1c\x78\xb7\xc1\x7e\xeb\x41\x22\x4a\x51\x01\xf6\xdb\x6f\x87\x1c\xa9\x95\xba\x41\x22\xbe\xff\x50\xd9\x1a\xf8\x66\x8d\x50\x19\x85\x17\xb4\x8e\xf3\x0f\x55\x45\xa6\xeb\x05\x3f\x6f\xf0\x8a\xbb\x7b\x48\xe0\x44\x50\xa8\xfe\x31\x73\xd1\x6e\x83\x57\x6b\xcc\x64\xcd\x19\xc8\xf9\x60\xcd\x64\xa7\x66\xe5\x51\x75\x99\x0b\x49\x17\x22\x69\xed\xed\x2d\x1e\x44\xb0\x14\x0c\xb5\x5b\xee\x90\x0b\x9e\xc3\xe9\x44\x4c\x4f\x78\x1a\xf5\x3d\x50\x4d\xe7\x9e\xad\x79\x10\x71\xcd\xb1\x42\xbe\xd2\xaf\xd2\x28\xac\x07\x3a\xd0\x91\xb8\xa0\x93\xfe\x78\x79\x64\xa3\x42\xe7\xb1\xdb\xa8\x47\xec\xb7\x6d\xba\xd9\xc3\x50\x64\x8c\x45\x07\x23\x11\xd4\x06\xa6\xbd\x49\xd5\xbf\x6f\xe8\x79\x31\x2e\xe3\xa3\xe2\x3c\xae\x2c\x5c\xed\x20\x7f\x22\x91\xda\xe3\x7f\x3d\xad\xc0\x6b\x87\x7f\xd1\x55\x52\xf1\x35\x06\x84\x28\x14\x0a\x0d\x78\xac\xc0\x05\x7f\xc9\x98\xf2\x97\x26\xe7\x12\xd7\xad\x63\xe4\xcf\xdb\x7d\xb3\x26\x11\x6b\x46\x3e\x68\x2f\x8d\x87\x7e\xef\x70\x0c\x2f\xe4\xea\x48\xfe\x46\x73\x61\x4d\xea\x97\xd0\xf9\x39\x46\x89\x58\xdd\xe3\x1e\x9d\x26\x7c\x9e\xd5\x5b\x63\xe2\xa1\x1f\xe8\xfa\x66\x72\xde\x36\x4f\x89\xd8\x4e\xf6\x4f\x00\x00\x00\xff\xff\x3d\x35\xa2\x28\x4e\x03\x00\x00")

func iteratorIteratorGoBytes() ([]byte, error) {
	return bindataRead(
		_iteratorIteratorGo,
		"iterator/iterator.go",
	)
}

func iteratorIteratorGo() (*asset, error) {
	bytes, err := iteratorIteratorGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iterator/iterator.go", size: 846, mode: os.FileMode(420), modTime: time.Unix(1585770953, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _converterConverterGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8e\xb1\x8a\x86\x30\x10\x84\x6b\xf7\x29\xa6\x4c\xc0\xc2\xb4\x07\x3e\xc5\xa5\x13\x8b\x10\x56\x09\x6a\x94\x4d\x3c\x38\xc4\x77\x3f\xf6\xd0\xbf\xdb\x1d\xbe\x6f\x98\x23\xc4\x25\xcc\x8c\xb8\xe7\x1f\x96\xca\x42\x54\x7f\x0f\x86\xef\x50\xaa\xa4\x3c\x3f\xaf\x43\xca\x95\x65\x0a\x91\xaf\xfb\xc3\x7c\xaf\x29\x32\x86\xd1\x77\x44\xd3\x99\x23\x4c\x79\x53\x0b\xbf\x7b\xf7\x7f\x9a\xa7\x1c\x8a\x18\xdf\x59\x78\x67\xd5\x72\xb8\xa8\x11\x2e\xf8\xea\xb1\x85\x85\x8d\x66\x2d\x56\xce\xa6\x58\x4b\xcd\xb4\x0b\x52\x8b\x54\x79\x53\x44\x42\x9e\x19\x45\x25\xb5\x86\x34\xa2\x7f\x87\x1b\x85\x2c\x35\xb7\x16\xd6\x53\x32\x84\x0b\xdd\xf4\x17\x00\x00\xff\xff\xfe\x85\x26\x30\xe0\x00\x00\x00")

func converterConverterGoBytes() ([]byte, error) {
	return bindataRead(
		_converterConverterGo,
		"converter/converter.go",
	)
}

func converterConverterGo() (*asset, error) {
	bytes, err := converterConverterGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "converter/converter.go", size: 224, mode: os.FileMode(420), modTime: time.Unix(1577384986, 0)}
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
	"list/list.go": listListGo,
	"set/set.go": setSetGo,
	"hashmap/hashmap.go": hashmapHashmapGo,
	"iterator/iterator.go": iteratorIteratorGo,
	"converter/converter.go": converterConverterGo,
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
	"converter": &bintree{nil, map[string]*bintree{
		"converter.go": &bintree{converterConverterGo, map[string]*bintree{}},
	}},
	"hashmap": &bintree{nil, map[string]*bintree{
		"hashmap.go": &bintree{hashmapHashmapGo, map[string]*bintree{}},
	}},
	"iterator": &bintree{nil, map[string]*bintree{
		"iterator.go": &bintree{iteratorIteratorGo, map[string]*bintree{}},
	}},
	"list": &bintree{nil, map[string]*bintree{
		"list.go": &bintree{listListGo, map[string]*bintree{}},
	}},
	"set": &bintree{nil, map[string]*bintree{
		"set.go": &bintree{setSetGo, map[string]*bintree{}},
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

