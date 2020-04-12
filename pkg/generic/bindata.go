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

var _setSetGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\x8e\xe2\x38\x10\x3d\xdb\x5f\xf1\x4e\xab\x64\x45\xa7\x39\xb7\x9a\x95\x5a\x7b\x5a\xa9\xb5\x87\x69\xe6\xd4\x42\x23\xc7\xa9\x80\xd5\x89\x8d\x5c\x86\x06\x21\xfe\x7d\x64\xc7\x09\xf4\x0c\x1c\xe6\x44\xa5\xa8\x7a\xf5\xea\xd5\x4b\xb6\x4a\x7f\xa8\x35\x81\x29\x48\x19\x8e\x5b\xc2\x72\x0e\x0e\xde\xd8\xb5\x94\x8f\x8f\x58\xce\xdf\x28\xc0\xd3\xd6\x13\x93\x0d\x0c\x15\x4b\xe1\xda\x58\x47\x1d\xf5\x31\x39\x36\xc6\x52\x0e\x7e\xa7\x03\x4e\x52\xf4\xe8\xd5\xf6\x7d\x39\x5f\xd5\xce\x75\xf2\x9c\xe0\xfe\xa7\xcf\xa1\x4c\x7b\x52\x81\x18\xca\xc2\x58\x0e\xca\x32\x0d\xa0\xf1\xcf\x4f\x13\x36\x58\x9b\x3d\xd9\xcb\x88\x76\x67\xf5\xd4\x5e\xec\x55\xc7\xa8\xaa\x6a\x39\x2f\xf1\xf7\xd0\x74\x92\xc2\x13\xe3\x69\x81\xbf\x52\xe2\xd4\x3f\x5d\x13\x38\x9d\xcf\xa9\xa0\x7a\x69\x9a\xd4\x5e\x55\x55\x19\x33\x61\xe7\x2d\x3c\x71\x66\xf8\xd2\x34\x50\x4d\x93\x98\xe5\xe1\x08\x6e\x58\x7b\x20\x51\x70\x1e\x59\x62\xc4\x1a\xa9\x9c\xa4\x68\x9d\xc7\x8f\x19\xf6\x91\x88\x57\x76\x4d\x48\x05\x27\x29\x04\x57\xfd\xfb\x7e\x85\x05\x82\xdf\x91\x14\xe7\x3c\xf1\x1b\xf5\x6e\x4f\xf0\xe9\xe7\xcb\xdc\xd6\xbb\xfe\xce\xe4\xa1\xa9\x38\x20\x8f\x6d\xa8\xa3\x40\x05\x57\xfd\x0c\x87\x32\x23\xbf\x52\xdc\x2c\x2e\xc8\x93\x90\xb0\xbb\xbe\x26\x0f\x63\x11\x36\x74\x13\xfb\x95\x6c\x51\xc2\xd8\xac\x69\x12\xa8\x23\x1b\xc1\x47\xe4\x7f\x9d\x0d\xca\x58\x86\xde\x90\xfe\x60\x98\x16\xea\xeb\xc5\x50\x53\xe7\xec\x9a\xef\x8b\x37\x62\xe4\x25\xe2\x91\xae\x06\x46\xb1\x0e\xab\xf1\x28\x5d\x37\x2d\xa2\xa0\x37\xca\x5a\xea\x60\x02\x79\x15\x9c\x87\xdb\x93\x47\xd8\x18\xbe\x7d\xa4\xae\x2b\x4a\x3c\x3f\xc4\xb6\xe8\xda\x93\x14\x3a\x5e\xa7\x57\x1f\x54\xe4\x64\x29\xc5\xda\x21\x76\x16\x49\xce\x74\xc6\xc3\xe5\x86\x5c\xf5\x29\x2d\x34\x9e\x1f\x70\x90\x42\x9c\xa5\x10\xba\x73\x4c\x85\x2e\xa5\x38\x17\x17\x2f\xe9\x4c\xfa\xbb\x35\xce\x42\x6d\xb7\x9d\x21\xc6\x2e\x3d\xb9\x6d\xa4\x1c\xa3\xe0\x26\xc6\x50\xb6\x99\xe4\xbb\xb5\x41\x42\x2a\xea\xe9\xf9\xe2\xf8\x5f\x68\xd6\x99\xe6\xe0\xf2\x43\x24\x76\x11\x34\xd3\x7a\xdb\xd5\xc1\x2b\x1d\xc0\x39\x60\xd4\x83\xd3\xee\x2a\x38\xb6\xfc\x19\x85\xd1\xa2\x37\x59\xfc\x67\x03\x79\x26\x1d\xd0\x1a\xdb\x70\x74\xdb\xf0\x9c\x54\x6a\x71\x4c\xa2\xdc\x65\x34\xb5\x17\xc7\x1b\x94\x54\xa4\xc3\x55\x2f\x45\x1d\xa3\x63\x8c\x4c\x9b\x3c\xac\x4a\xfc\x93\x82\x7a\xb8\xb3\x9a\xa1\xc6\x02\xf5\x0c\x2a\xd3\xe4\xc9\x1b\x57\x5f\x8f\xf2\xb7\x3d\x55\x6a\x37\x2d\xea\xf7\xc3\x6a\xf0\x86\x27\x8e\xf1\xf8\x7e\x47\xb8\xb3\x8c\x6f\x3d\x16\xe9\x03\x73\xad\xc1\xcf\x00\x00\x00\xff\xff\xe5\x69\x87\x0b\x78\x05\x00\x00")

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

	info := bindataFileInfo{name: "set/set.go", size: 1400, mode: os.FileMode(420), modTime: time.Unix(1586691162, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _hashmapHashmapGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x41\x6f\xf2\x38\x10\x3d\xc7\xbf\x62\xf6\x3b\x7c\x0a\x2b\x14\xe0\x8a\xd4\xd3\x6a\xf5\xa9\xd2\x76\x0f\x6d\xba\x7b\xa8\xaa\xca\x98\x09\x58\x24\x76\x64\x3b\xb4\x6c\xc5\x7f\x5f\x8d\x9d\x10\x07\x08\x9c\xda\x84\x99\xf7\xde\xcc\xbc\x97\x9a\x8b\x1d\xdf\x20\x6c\xb9\xdd\x56\xbc\x66\x4c\x56\xb5\x36\x0e\x7e\xd8\x83\x12\x3f\x18\x73\x87\x1a\x21\x9f\x83\x75\x46\xaa\x4d\xfb\xb8\x00\xa9\x1c\x63\xb3\x19\xe4\xf3\x7c\xf1\xc4\xeb\x7f\xa4\x95\x4e\x1b\x90\x16\x38\xec\xdb\x87\xa2\x51\xc2\x49\xad\xc0\xe9\xf0\x0e\x2a\x5e\x43\xcd\xa5\xb1\x1d\xec\xa0\x99\xea\xd3\x7c\x3e\x85\x7c\x31\x81\x95\xd6\x65\xcc\x00\xf8\x55\x6b\x8b\x16\xdc\x16\x41\x68\xe5\x0c\x17\x0e\x74\xe1\x7f\x27\xe0\x01\x24\xe9\x43\x53\x70\x81\xf0\xcd\x92\xd9\x0c\xfe\xe4\x62\x1b\x44\x58\x40\xfa\x1f\x4b\xac\x50\x39\x90\xca\x23\x56\xbc\xce\xe0\xd1\x81\x75\xba\xb6\x20\x1d\x1a\x4e\xd2\x2d\xc8\x62\x30\x0f\x18\x74\x8d\x51\x16\x0a\x5e\x5a\x64\x09\xe1\xa6\x5d\xc1\x70\x9e\x09\xf3\xcc\xbf\xd0\x9d\x9a\x88\x69\xcf\xcb\x06\x49\x38\x87\x8d\xdc\xa3\x82\x1d\x1e\x32\x78\x2c\xfc\x8f\x3b\x3c\xc0\x27\xb7\xa0\xb4\x83\x42\x37\x6a\xed\xdf\x5a\x14\x5a\xad\x5b\x94\x16\xe0\x53\x96\x25\xac\xb0\xd3\xf1\x0b\x5d\x4a\xcd\xf9\x7c\x02\xe9\x1e\xf2\xc5\xb4\xed\xa7\x3d\xb6\x4a\x5e\xd0\x81\x45\x37\x2e\x83\x25\x2f\x27\x98\x29\x55\xd0\x25\x42\xef\x6b\xbd\xe6\x0e\xa1\xf1\x7f\xda\x23\x34\xc6\xd0\x06\xe9\xa8\x85\xd1\xd5\x09\x89\x8e\x91\x84\x86\xd4\x1a\x41\xcf\x6f\xf9\xfc\x9d\xae\xda\x6e\x28\x60\x3e\x63\xa5\xf7\x08\xc6\xff\xb1\xb1\x90\x80\xe7\xb6\xd2\x06\xb0\x50\x79\x1a\x30\x78\x83\x20\xfe\x28\xb5\x42\x10\x06\xbd\x2a\x0e\x42\xd7\x07\x1a\xaa\x6f\xf5\x15\x69\xcf\x7c\x6c\x2d\xbd\xe2\x16\x3b\xb3\x58\x67\x1a\xe1\xc8\x29\x1f\x34\xcc\x49\x2f\x15\xcf\x66\xf0\x37\x7e\x76\x95\x3d\xd3\x8a\x5b\x29\x40\x2a\xeb\xb8\x12\x18\x38\x4f\xf6\xf3\x56\x92\x16\x3e\x1a\x65\x79\x81\x1f\x50\x68\x43\xa6\xed\x56\xc6\x85\x40\x6b\x33\xe6\x1d\xd5\xc3\xf7\x32\x49\x8b\x41\x0b\xcb\x07\xf8\x19\x29\xfd\x66\x89\x97\xb8\xec\x35\x7e\x1f\xa7\x2c\x39\x52\xb5\x37\x87\x41\x7b\xa1\xfa\xe5\xa0\x84\xd1\x4a\xfe\x87\xeb\xc1\xaa\x4e\x72\x48\xe3\xd8\x28\xe7\x1a\x23\xb0\x73\xb9\x5e\xc0\x4f\xfa\x6e\x44\x72\xa5\x52\x68\x96\x83\x21\xbd\xe0\x23\x0b\xc8\x69\x05\xbf\x47\x13\x4e\xe0\xb6\x93\x89\x69\xdf\xbd\x79\x80\x2a\xa3\x75\xbc\xed\xf0\xf0\xde\x29\x18\x47\xbe\x95\x55\xc2\xa5\x23\xed\xa6\xb0\xa7\xad\x1b\xae\x36\xd8\xc2\x83\x9f\xa3\x80\xdf\xda\xde\x94\x8a\x7c\x47\xd2\x71\x26\x74\x82\x1b\x43\x5d\xc9\x15\xf5\x47\xf2\xe1\x81\x7e\x18\x47\xb8\x99\xa7\xeb\xea\xa9\x96\x44\x76\x2c\x9e\x23\x36\x4b\x35\x4e\x77\x25\x71\x3e\x1f\xdd\xea\x97\x67\xbb\x5f\x63\x89\x0e\xd3\xf0\x6e\x4a\x19\xa6\x0f\x47\xcb\xe3\x5b\xc6\xb9\xce\x22\x1a\x79\x3f\x76\xcd\x9d\xfb\x18\xb4\x99\xdf\x32\xdd\x86\x86\x3c\x8f\x84\x4f\x7d\x64\xce\x28\xf5\xde\xa3\xd1\xb7\xa9\x6a\x1c\x7e\xf9\xda\xec\xf9\xdf\x27\x7a\x18\x88\x8f\x40\xee\x9b\xaa\xca\x3c\x5a\xf6\xfc\x97\x16\x3b\x9a\xa2\xca\x3c\x5d\x16\x37\x4e\xa2\xba\x57\x55\x86\xca\x51\xca\xfb\x09\xb9\x20\x1d\x44\x26\xf0\xb7\x28\x57\xa9\xaf\x45\x69\x20\x61\xd4\xcf\x01\xe9\x7c\xd6\xb6\xdc\xd7\x46\x84\xf7\x47\xbd\xeb\xfa\x31\xc6\xbe\xf1\x1a\xdf\x48\x02\x06\xd4\x23\x09\x38\x27\x8c\xe2\x10\x98\xfb\xbe\x9b\xcc\x97\x99\x18\xb0\x5f\xc9\xc4\xc5\x51\x4d\x4c\xdb\x36\xdc\xb8\x26\x18\x76\x64\xff\x07\x00\x00\xff\xff\xdc\xa7\x95\x29\xef\x09\x00\x00")

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

	info := bindataFileInfo{name: "hashmap/hashmap.go", size: 2543, mode: os.FileMode(420), modTime: time.Unix(1586691671, 0)}
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

