// Code generated by go-bindata.
// sources:
// js/event.js
// js/job.js
// js/mock8s.js
// js/run.js
// js/run_mock.js
// js/runner.js
// js/waitgroup.js
// DO NOT EDIT!

package lib

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

var _jsEventJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x41\x4a\xc6\x30\x10\x85\xf7\x39\xc5\xdb\x55\x41\xd2\x13\x74\xe1\x42\x70\xed\x0d\xa6\xc9\xd4\x04\x62\x52\x33\x53\x6b\x91\xde\x5d\x42\x94\x9f\xf2\xef\x86\x79\x8f\xef\x9b\x19\x47\xbc\x7c\x71\xd6\x57\xca\x3e\x71\x85\x67\x71\x35\xce\x2c\xd0\xc0\x48\x51\x14\x65\x01\xb7\x4a\x5b\x91\xe2\xd9\x45\x8f\x28\xa0\x9d\x2a\xa3\x2c\xd6\x2c\x5b\x76\x1a\x4b\xbe\x90\x1e\x1e\xf1\x63\x80\xce\xaf\x47\x47\x20\xfc\x69\xde\x59\xbb\x61\xa5\x4a\x1f\x18\x3c\x29\x0d\x4f\xd8\x43\x74\xa1\xc1\x5b\x34\x17\x7f\x34\x79\x9b\x2b\x7f\x6e\x2c\x6a\x0d\xa0\x21\x8a\x5d\x37\x09\x98\xf0\x6f\x6e\xb2\xf3\x96\xa5\xf4\xd6\xfb\x77\x95\xd3\x18\xfe\x5e\x4b\x55\xb1\x97\xb7\xa7\xcb\xed\xe6\x37\x00\x00\xff\xff\x38\xd9\xc5\x1d\x16\x01\x00\x00")

func jsEventJsBytes() ([]byte, error) {
	return bindataRead(
		_jsEventJs,
		"js/event.js",
	)
}

func jsEventJs() (*asset, error) {
	bytes, err := jsEventJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/event.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsJobJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x54\xc1\x6e\x1b\x37\x10\xbd\xef\x57\xbc\xea\x24\xb5\xca\xae\x83\xb4\x3d\xc4\xd0\xc1\xa8\x1d\x20\x29\x50\x04\xad\x8b\x1e\x8a\xa2\x98\xe5\xce\x6a\x19\x73\xc9\x2d\x39\x94\x2d\x04\xfe\xf7\x62\xb8\x96\x2c\x2b\xc9\x45\x02\x67\xf9\xde\xcc\xbc\x79\xc3\xa6\xc1\xed\xc0\xe8\xb8\xa7\xec\x04\x76\xa4\x2d\xc3\x26\x24\x09\xe6\x0e\xb9\xcd\x5e\x32\x5e\xff\x5c\x5f\xfc\x88\x1f\x30\xd2\x1d\x83\x7c\x87\xad\x95\xba\xda\x51\x04\x19\xdb\xbd\x2f\x98\x0d\x16\x7a\x78\x35\x43\xde\x3a\x12\x4e\xb2\xa8\xaa\xa6\xc1\xc7\x18\x24\xc8\x7e\x62\xf4\x21\xe2\x43\x68\xeb\xaa\xcf\xde\x88\x0d\x5e\x4f\x4b\x4f\x23\xaf\x21\x94\xee\xd2\x0a\x9f\x2b\x40\x99\xc7\x3d\x36\x90\xc1\xa6\xaa\x02\x6c\x8f\xe5\x77\xfc\x30\x85\x28\xa9\xfe\x97\x77\xec\x65\xbe\x09\xc8\x10\xc3\x3d\x16\x25\x06\x1f\x04\x7d\xc8\xbe\x5b\x54\xc0\xa3\x22\x9b\x06\xbf\xd1\xc8\xb8\xb7\xce\xa1\x65\x13\x46\x86\x0c\x8c\x29\x72\x6f\x1f\x4a\x41\xe5\x18\xba\xc6\x04\xdf\xdb\xed\x48\x13\xb4\xa0\x54\x57\x28\xf9\x6b\x3d\x61\x53\x82\x97\x33\xe3\xad\x96\xaa\x2a\x29\xd4\xd9\x24\x08\xfd\x5c\x3f\x24\x20\x66\x5f\xab\xa8\x7b\x50\x64\xf0\x03\x9b\x2c\xdc\xc1\x7a\x24\xfe\x2f\xb3\x37\x0c\xeb\x93\xed\x18\xa1\x9f\xf9\x08\x69\x60\xe7\xb0\x6c\x5a\xeb\x9b\x34\xac\x8e\xb9\x67\xd2\x0d\xfe\xfe\xe7\xf2\x49\x86\x13\x99\xce\xee\x94\xff\x93\xbe\xaf\x60\x82\x73\x3c\xeb\x1c\xfa\xd2\x40\xb3\x23\x97\x19\x13\xd9\x98\x34\xc6\x7e\x67\x63\xf0\xa3\x8a\xb7\xa3\x68\xa9\x75\x27\x9d\xb3\xdf\x61\x83\xcf\x8f\x97\x4f\x8c\xea\x94\xd9\x21\x6a\x02\xf2\x08\x93\x92\x93\x83\xd0\xf6\x88\xb2\x4f\x7e\x38\x7a\xe3\x00\x9f\x42\x57\x66\xa1\xf6\x62\x41\xbb\x57\xa9\x96\xab\x75\x61\x33\xc1\x0b\x59\x3f\x6b\x5a\x24\x57\x4d\xe7\xd1\xc0\x44\x26\xe1\xee\x98\xe2\x89\xe9\x89\x37\x66\x95\xd6\x77\xa9\x7c\xc4\xa7\xd0\xea\x18\x7e\xcd\x2d\x47\xcf\x72\xd2\x8f\x5e\xdc\xe0\xe0\xbd\xe5\x0b\x15\x5b\x32\x77\xdb\xa8\xde\x59\x9e\x19\xed\xf9\xce\x3d\x59\x59\xae\xaa\x12\x88\x2c\x39\xfa\xd9\xa1\xc0\x2c\xd1\x19\xd1\x37\x73\x1d\x84\xd8\x14\x01\x34\xb4\xc6\x59\xd2\xcb\x23\x69\xd3\x40\xf3\xfe\xe9\xc5\xba\xeb\xe0\x8b\x7c\x03\xc7\x79\x99\x34\xd9\x3d\xc5\x2e\xc1\x84\x71\x22\xb1\xad\x75\x56\xf6\x6b\xb4\x59\xd0\x05\x4e\xba\x12\x83\xf5\x65\x3a\x4d\x83\xeb\x9b\x8f\xbf\xdf\xfc\x72\x75\x7b\x73\xfd\x16\x7f\xcd\x2b\x81\xc8\x63\xd8\x71\x87\x2e\x47\xeb\xb7\xb8\x72\xd3\x40\xd5\x49\xc7\xcf\x99\xcf\xfa\xf9\x9a\x0e\x27\x15\x97\x9f\x84\xac\x70\x50\x99\xe3\xa0\x81\xc5\x1f\xd9\x18\xe6\x8e\xcb\x96\x36\xcd\x33\x62\xb9\x82\x21\xaf\x35\x19\x72\x8e\x3b\x04\x8f\x67\x39\xf5\xdc\x7e\x62\x23\xa9\xfe\x12\x67\x13\x28\x4b\x18\x49\xac\x62\xf7\x07\x86\x76\x7f\x1c\xfd\x29\xea\x7d\x5f\xea\x99\x8b\x4f\x58\xbc\x23\xeb\xb8\x5b\xac\x67\x07\x95\x37\x25\xa9\xc1\xf9\xc1\x70\xf1\x78\xfd\x12\x97\x7d\x2a\xfa\x8f\x21\xea\x7b\x42\x1e\xaf\x7f\xc2\x68\x7d\x16\x4e\x58\xbe\xb9\xb8\xc0\xf7\x78\xf3\x2a\xb1\x09\x5e\x37\x5f\x38\xee\xc8\xa5\xd5\xfa\x48\x0d\xb1\x23\x87\x2c\x2f\x33\x1c\x35\xff\x9a\x75\x34\xfe\x2e\x44\x7d\x31\xc7\xfd\x17\x86\x99\xfd\xf2\x58\x55\x87\xf8\x87\xd0\x62\xa3\xef\x6b\xf5\x7f\x00\x00\x00\xff\xff\x61\xed\x91\xe4\xe4\x05\x00\x00")

func jsJobJsBytes() ([]byte, error) {
	return bindataRead(
		_jsJobJs,
		"js/job.js",
	)
}

func jsJobJs() (*asset, error) {
	bytes, err := jsJobJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/job.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsMock8sJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcd\x8a\xdb\x30\x14\x85\xf7\x7e\x8a\x43\xe8\x62\x02\xc1\x43\xb7\x81\xac\x0a\xdd\xb5\x94\x0e\x74\x53\x4a\xb9\x91\x8e\x13\x75\x6c\xc9\xe8\x67\x3a\x50\xfc\xee\x45\xae\xa5\xd8\xd9\xb4\x8b\x6e\x94\x9b\xfb\xf3\x59\xe7\xea\x3c\x3e\xe2\x29\x3a\x4f\x0c\x4e\x3d\x63\x74\x1a\x9a\x9d\xb1\x26\x1a\x67\x03\xae\xf4\x84\x04\xec\xac\x0c\xdc\xe1\x88\x37\x9e\xc1\x25\xaf\xd8\xbc\x88\x9f\x67\x3e\x39\x1d\x70\xc2\xaf\x66\x6a\x9a\x0a\x53\xce\x76\xe6\x32\xc8\xf8\x4f\x88\x77\x73\xf7\x87\xdc\xbd\x80\x72\xe5\x39\x9d\xe9\x2d\x23\xff\x64\x81\x9f\x26\x5e\x3f\x3e\x1d\xd1\x25\xab\xf2\xf5\x1e\x6c\xd8\xcf\x05\xc0\x33\x26\x6f\x17\x98\x67\x03\x4c\x05\x53\x72\x0b\x44\x39\xcf\x2f\x6f\x8f\xcb\xdc\xe8\x74\x09\x81\x0b\xe3\x1a\x2e\x03\xf7\xb5\x56\x3f\xf1\xbd\xed\x8c\xd5\x0f\x45\xf9\xe1\x36\xc0\x7e\xd3\x5f\x27\xd8\xb3\x1d\x18\x45\x4b\x94\x36\x53\x71\x3a\x21\xff\xd6\xd6\x69\xbf\x84\xd3\x61\x09\x94\xa7\x44\xae\x6e\xa3\xd9\xad\xe1\x79\xd1\x49\x29\x52\x53\x1f\xf0\x39\x59\x6b\xec\xe5\x00\xb1\x1a\xef\xc5\xf4\xd4\x10\x4f\x04\x37\x10\x2f\xd2\x1b\x9d\xcf\xc4\xd0\x56\x80\x66\xd7\x86\x28\x31\xcd\xcb\xc5\x78\x95\xc0\x23\x76\x15\xba\xc3\x54\x7b\x8b\xd6\xaf\x79\xa8\x2a\xe9\xe5\xcc\x3e\xb4\x3f\xdc\x39\x6b\xf9\x86\x53\x66\x36\xf7\xeb\xca\xb9\x45\x5b\x53\xce\x59\x25\x5f\x23\x6d\xc8\x26\x2b\x0f\x50\x3d\x73\x7b\x91\xbf\xad\x61\x6b\x9e\xff\x76\xbf\x6c\x1d\xbe\x8e\xce\xc7\xd0\xae\x2c\x5e\xc2\x4d\x6d\xe3\xdd\x6d\xa2\xf6\x6d\x9c\x7c\xfb\x73\xc7\x99\x2d\x5a\x1d\xfc\x3b\x00\x00\xff\xff\xc6\x15\x83\xe4\x98\x03\x00\x00")

func jsMock8sJsBytes() ([]byte, error) {
	return bindataRead(
		_jsMock8sJs,
		"js/mock8s.js",
	)
}

func jsMock8sJs() (*asset, error) {
	bytes, err := jsMock8sJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/mock8s.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRunJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x7b\x6f\x1b\xb9\x11\xff\x5f\x9f\x62\x8e\x38\xe4\x24\x64\xbd\x4a\x70\x40\x71\x50\x2a\x14\x3e\xd9\xb9\x28\x89\x2d\x37\xf2\xa5\x2d\x0c\xc3\xa0\x76\x47\x5a\x46\x5c\x72\x41\x72\xa5\x53\x73\xfe\xee\xc5\x90\xdc\x87\xfc\x48\xd0\xa2\xff\xd8\xbb\xcb\x99\xdf\xbc\x1f\xd4\x78\x0c\xd7\x05\x82\x50\x0e\x8d\xe2\xd2\xc2\x5a\x1b\x30\xb5\x52\x42\x6d\xc0\x71\xbb\xb5\x29\x5c\x17\xc2\x42\x59\x5b\x07\x2b\x04\xa9\x79\x8e\x39\xac\x70\xad\x0d\x02\x57\x07\xd0\x6b\x70\x05\x0e\xc6\x63\xd0\xab\x2f\x98\x39\x0b\xae\xe0\x0e\x6a\x8b\x04\x34\x1c\xa5\x83\xf1\x98\x4e\x4f\xa5\x84\x0f\xf5\x0a\x8d\x42\x87\x16\x4e\xaf\xe6\x90\x71\x29\x2d\xd8\x42\xd7\x32\x0f\xe0\x19\x97\xe2\xdf\x98\x43\x81\x06\x53\x58\xb8\x02\x0d\x94\x3a\xaf\x25\xb6\x74\x4a\x3b\x82\x23\x5e\xf8\x69\xdb\x02\xfe\x04\xb9\x30\x98\x39\x79\x48\x07\x74\xbe\xe7\xc2\x81\xe3\x5b\xb4\xc0\xe1\x8b\x5e\x9d\x48\xb1\xc5\xa8\x22\x70\x95\x7b\x02\x0b\xb5\x72\x42\x82\x70\x60\xeb\x2c\x43\xcc\x2d\x68\x03\x6b\x2e\xa4\x4d\x07\xeb\x5a\x65\x4e\x68\xe5\x49\xdf\x6a\xf3\x5e\xaf\x86\x5f\xf4\x2a\x01\x1c\xc1\xd7\x01\xc0\x8e\x1b\x28\x0f\x30\x25\xf8\xc1\x00\xbc\xef\x86\xf4\x51\xc0\x14\x5e\xbd\x01\x01\x7f\x85\x9f\x5f\xd1\xc3\xcb\x97\x81\x03\x20\xd3\xca\x6a\x89\xa9\xd4\x9b\x21\xcb\x0a\xcc\xb6\xe4\x69\xeb\xb8\xab\x2d\xb9\x92\xc1\x4b\x28\x0f\x69\xa5\xf3\x4b\x5e\xe2\xc8\xf3\x10\xe4\x16\xa6\xd0\x19\x9b\xee\x85\x2b\x2e\x97\x43\x4c\x7b\xdf\x14\x2f\xd1\x56\x3c\xeb\x71\x95\x87\x4a\xe7\xc4\x99\x66\xda\xe0\xe7\xd7\x84\x9b\x6e\xd0\x0d\xfb\x22\x1e\xe9\xf5\x7e\xb9\xb8\x4c\xad\x33\x42\x6d\xc4\xfa\x30\xf4\x20\xa3\xd1\x63\xf5\xaf\x74\xfe\x40\x5f\x78\x09\x0c\x84\x05\xa1\xbc\x49\x18\x8f\x49\x6a\x30\x31\xad\x0a\x6e\x1b\x99\x62\x0d\xc3\xc7\x87\x30\x9d\x02\x7b\xcb\x85\xc4\x9c\x35\x5e\x03\x70\x85\xd1\x7b\x78\x4e\xe4\xda\x93\x83\xd3\x94\x72\xf4\x2f\xd3\x65\x25\x91\x82\xc7\xde\x78\x84\xfb\xef\x49\x5c\x86\xf8\x1f\x09\x35\xe8\x6a\xa3\xc0\x99\x1a\x7b\x20\xe3\x31\x2c\x25\x62\xe5\xe3\xcd\x21\xc7\xb5\x50\x98\x03\x2f\x75\xad\x9c\x2f\x07\x51\x62\xea\x49\x2d\xd1\x0d\x7f\x26\xdf\x11\x6f\x34\x82\xce\x73\xd0\xb5\xf3\x89\x45\xf1\x27\xa4\xea\x69\xd3\x82\x4d\x6c\x70\xef\xb3\x9a\xcc\x33\xb5\x8a\x39\x1d\x34\x50\x80\x3b\x54\x0e\x86\x38\x4a\x61\xee\x20\xd7\x68\xa9\x4a\x42\x09\x10\x89\x2b\xd0\x93\x77\x8e\x41\xaa\xca\x50\xd9\x82\xc0\xa4\xde\x9f\x48\xdc\xa1\x84\xca\x88\x52\x38\xb1\xc3\x5e\xfe\x53\x1d\xf7\x13\x7f\x3c\x86\x1f\xdf\x2f\x7e\x3d\xf9\xf1\x7a\x7e\x71\x7e\xf2\xe3\x6f\xf3\xeb\xe5\xbb\xd3\x58\x10\xdb\x5f\xac\xd7\xdd\x57\x85\xcf\x49\x32\xe3\x84\x2c\x3b\xe3\x0e\x53\xa5\xf7\xc3\x51\xfb\x09\xd3\x4c\x97\xa5\x70\xa9\xad\x57\x21\xe3\x86\xaf\x12\xf8\x65\xf4\x26\xa2\x65\x65\x04\x8b\xb0\xf1\x33\x75\x28\x34\x4f\x1e\x65\x25\x4c\x41\xe1\x7e\x76\x31\x0c\xcc\xa3\x23\x9e\x70\xf8\xc9\x3f\x5f\xe9\x7c\xd8\x21\x25\x5e\x61\x51\xf2\x4d\xc8\xcf\x70\x92\x96\xe8\x78\xce\x1d\x4f\x25\x5f\xa1\xb4\xe9\x17\xbd\x52\xc7\xf6\x3d\x4f\xbb\x42\xa9\xd5\xc6\x3a\x0d\x53\xc0\xd4\x60\xa5\x3d\x03\x3d\x49\x9e\xe1\x90\x8d\x59\x42\x8e\x18\x3d\x0f\x11\xdc\xe3\xf9\xc3\xe3\x20\xf8\xff\x34\xcf\x01\xd5\x8e\x2c\xb3\x69\xb4\x10\xd5\xee\x33\x37\x16\xa6\x70\x73\x4b\x64\x77\x29\xf2\xac\xa0\xc8\xa5\xa8\x76\x09\x34\xf1\x1c\xee\xb8\x4c\x60\x8b\x87\x26\xd3\x23\x63\x5a\xd5\xb6\x18\x7e\x25\x15\x27\x74\x9c\xc0\x8e\xcb\x1a\x27\xf4\xef\xde\x87\x84\xfe\x06\xf9\x67\x1a\xf6\x08\xd6\x09\x29\x61\xcf\x95\xa3\xdc\xe2\x79\x0e\x8e\x52\xca\x69\x9f\x72\xde\x97\x6d\x5b\xfe\x1b\xfc\xa3\x10\x12\xa9\xd7\xfa\x94\xb3\x98\xd5\x46\xb8\x43\xc0\x73\x85\x50\x9b\xc4\xa7\x2d\xcf\x73\xaa\x09\xe1\x60\xef\x1b\xbe\x41\x5b\x4b\x47\x3d\xa5\xb6\x68\x42\x6a\xaf\x90\x48\xf8\x4a\x22\x09\x23\xbd\x69\x16\x79\x90\x80\x57\x57\xd6\x19\xe4\x25\x8d\xb6\xa0\x0d\x55\xfd\x20\xb4\x80\x18\x0b\x6b\x8b\x0f\xcf\x39\x21\x76\x80\xe0\x0b\x76\x3a\x9b\x9f\xdd\x7d\x3a\xbf\x5a\xdc\x7d\x38\xff\x17\x4b\xe2\x61\xf4\xce\x11\x5a\xe8\x13\xa1\xe0\xbb\x48\x39\x5d\xc5\xf2\x6a\x63\x46\x73\xd7\x62\x98\xab\x7a\x87\xc6\x88\x3c\xcc\x53\xee\x1c\x96\x95\x77\xa8\x45\xe7\x55\xf7\x82\x6c\x74\x94\x06\xab\x4b\xf4\xa6\x02\x4a\xeb\x3b\xcd\xb1\xee\x8d\xd6\xb3\x8f\x8b\xcb\xf3\xbb\xdf\x3f\x7d\x64\xc9\x03\x5d\x33\xa9\x15\xfe\xfe\xe9\x63\xd0\xf4\x69\xee\xe5\xf2\xdd\x93\xbc\xd6\x16\xdf\xe1\xfc\x6d\x7e\xfd\x24\xe7\x46\xb8\xef\x70\xbe\x3b\x3f\x3d\xbb\x9b\x2d\x2e\x2e\xe6\xd7\x77\xf3\xb3\x3e\x40\xac\x83\x6f\xf0\xce\xe6\x1d\x3d\xa3\x96\xcd\x02\x75\x2c\x2c\x5b\x61\x96\x66\x5a\x39\x2e\x14\x1a\x7b\xf3\xea\x96\x8a\x82\x0a\x2b\xa0\x0d\x9a\x81\x4e\x5d\xfc\x8a\xbb\x22\x56\x78\xf7\xfe\xe7\x9f\xc0\xc6\xd6\x64\xac\x17\xd8\x4c\xab\xb5\xd8\x40\xc9\x2b\xd8\x69\x59\xf7\x9b\x81\x17\x18\x3e\xfa\x92\xf4\x99\xd1\x68\x1b\x5a\x53\x12\xf9\x2f\x78\x35\x81\xaf\xfd\x13\xb8\xbf\x4f\x8e\x18\xd8\x2e\xb3\x27\x56\xe4\x98\x71\xc3\x12\xa0\x0c\x39\x9c\x09\x33\x81\xaf\xf7\xed\x54\x4a\x3a\x6a\x91\x1b\xcb\x59\x42\x45\x66\xd0\x4d\xe0\x6b\x7c\xba\xf4\xc7\x16\xb3\x28\x64\x00\x70\xfb\xe6\x9b\x4e\x0a\x16\x5c\x90\x17\x9e\x37\xa3\x75\xd2\x04\xd8\xb8\xd0\x7a\xcb\xbe\xad\x7e\x8f\xbe\x7d\xfc\x96\x19\x8f\xf0\xc7\xd6\x16\x2c\x01\x83\x3c\x5f\x28\x79\x98\xf8\x19\x1d\xad\xe9\x95\x5d\x81\x10\x85\x36\x3d\x32\xbe\xc2\xb4\x79\x5a\x56\x98\x0d\x31\x89\xa1\x4d\xe0\x68\xa5\xda\x65\x76\x19\xc8\x46\x11\xf6\x7a\x71\xb6\x98\x50\xd8\x76\x68\x5c\xdb\xed\xb8\x02\xa1\x84\x83\xd6\x75\x40\x0b\x9a\x5f\x77\xe1\x75\xfa\x97\xc0\xda\xf7\x31\x51\xcf\x5a\x3f\x93\x5f\xa3\x3a\xb7\x4f\x8c\x02\xae\x94\x76\x9c\x3a\xb7\xbd\x61\xd4\xc6\x56\xe8\x78\x5f\x4d\xa1\xc7\x04\x78\xd2\x45\x8e\xdd\xc2\x14\xd8\x0d\x8d\xd8\x07\xdb\x5c\x94\xe3\x47\xf0\x6d\x93\xca\xef\xb5\x50\xde\x5b\x7e\xe5\x27\x8b\x4a\xbe\x45\xe0\x34\x2b\x69\x65\x28\xb9\xca\x27\x7d\x07\xf8\xcd\x21\x2e\xe4\x95\xd1\x2b\xbe\x92\x07\xd8\xa0\x42\x43\x5b\x1f\x07\x5b\xa0\x94\x60\x33\x23\x2a\x97\xd0\x2e\x68\xfc\x9e\x43\x5e\x09\x30\x5d\x0f\xf3\x7b\x39\x75\xba\x13\xd4\x50\x89\x0a\x69\x99\x4b\xfc\x8a\x5e\x5b\x22\x50\xb8\x97\x42\x21\x6d\x95\xd6\x21\xcf\x69\xc5\x7a\xf1\xa2\x09\x28\x8d\xfa\x92\x76\x5c\xe6\x8d\xa1\xf6\x4e\x65\xeb\x0d\x69\x5a\x7b\x4b\xd3\x9e\xa4\x5f\xb4\x50\x43\x06\x2f\x5e\x00\x6b\x9b\x75\x56\xa6\xe4\xef\x1b\x56\x72\xa1\x52\x5b\x78\x27\x06\xde\xa6\x3b\xfc\x77\x6b\xf8\xe0\xc1\xc2\x3c\x33\xc8\xbd\x1f\x42\xd9\x53\xd7\xa0\x08\x65\x65\x17\x6b\x15\x37\x96\x6d\x8a\x7f\x38\x54\x96\x82\x9e\xb6\xe4\x69\x46\x08\x38\xcc\xca\xd1\xb3\xd8\xcd\x1e\xf9\x7c\x49\x77\x32\x7a\xb7\x82\x88\x1c\xb8\x1e\xa1\xc7\x2b\x61\x9a\xa6\x2c\x2c\x47\x61\x33\xee\xb6\xa7\x37\xb4\xa0\xb6\x0b\xe3\x83\xda\xf2\xf7\xba\x24\xec\x03\xdd\xd5\xc9\xbf\x5e\xf3\x0d\x4c\xc3\x63\xb3\xa7\x61\xa5\x69\x4c\x4c\x1f\x0e\xab\x26\xbc\x3f\x34\x8c\x4d\x78\x7b\x40\x8c\x67\x22\x1f\xf7\x9a\xcd\x44\x72\x87\xd6\xb1\x26\xc6\x8f\xc6\x3f\xfc\x40\xa9\xd3\x40\x3d\x12\x1e\xa6\x5d\xc3\xed\x3b\x48\x85\x19\x4c\x9b\xcc\x0a\x7d\x8a\xa4\x9e\x1c\xb5\xb8\x66\xa5\x98\xc4\x8e\xd9\x6b\x85\x9f\x67\x4b\xbf\x4b\x74\xe3\xaa\x91\x7a\x9f\x3c\x45\xfb\x71\x31\x3b\xfd\x78\x77\x75\x7a\xfd\xae\xe3\xf0\x1e\x7d\x86\xfe\xd3\xf9\xe7\xf9\x72\xbe\xb8\x7c\x6a\x7c\x06\x86\xdb\xa4\xf3\xdb\xa4\x75\x5f\x12\x6f\x78\xa1\xda\xe1\x86\xf5\xfd\xc8\xfa\x3c\x57\xb5\x94\x57\x5a\x8a\xec\x40\xbb\x91\xdc\xf3\x83\x8d\x16\xf7\xa7\xc5\x13\xa6\x3f\x3b\x05\xbc\x3d\x8d\x76\xcf\xc6\x2a\x46\xc9\xa7\x34\xaa\xdd\xff\x77\x55\x8b\x29\x4d\xe0\x47\xb9\x7c\x74\x5d\xa8\x74\xae\xfc\xc4\xa3\x88\xcf\xbb\x7c\x8e\xcc\x41\x17\xb6\x15\x2a\x67\x13\x7f\x5b\x8d\x1a\x30\x5e\x89\xcf\x68\xa8\x9c\xe9\x60\xf7\xba\xf9\xde\x14\x3e\x9b\xb4\xb7\x4e\x46\x12\xd8\x04\x1a\x59\xcd\xe7\x70\x23\xe8\x11\x02\xb0\x02\x8d\x70\x7c\x43\xe4\xec\xef\xb5\xde\x6e\x79\x6b\x32\x81\x73\xc5\x37\x98\xff\x7a\x60\x31\x47\x59\x3c\x6b\x13\x87\xf5\x66\x0b\x21\x87\x21\x1c\x8f\x19\xf9\xa2\xaf\x98\x41\xdf\xc9\x43\xe8\x09\xf3\x12\x77\x68\x5a\x89\xac\x37\x7f\xba\xe0\x43\x4f\xdf\xce\x38\xaf\x0e\x5d\x6c\x93\xfe\xa1\xcf\x2e\x36\xe9\xbc\x7b\x74\x1a\x53\xf3\x08\xdb\x1f\x8c\x57\x42\x8d\x69\x21\x78\xf0\xd9\x6f\x0a\x4d\x33\xef\x9d\xdd\xf6\x09\xc7\x63\x78\x3b\xff\xe7\xc5\xf9\x04\x66\x05\x57\x1b\x7f\xa1\x60\xf3\xf5\xa5\x76\x57\x06\x2d\x2a\xc7\x1e\x69\xd8\xe5\x3f\xeb\x0a\xa0\xa5\xba\x8f\x4f\xb7\xed\x6f\x06\xf7\xbe\x3b\x1e\xa5\xd4\xec\x62\xe8\xdb\xf0\x37\xb2\x67\xd6\xec\x86\xff\x43\x0e\x75\x8e\xee\xa7\xd0\xd3\x49\xf4\xfd\x44\x8a\xf9\xd0\xa4\xc5\x63\x59\x8d\x8f\x27\xc0\x30\x2b\x34\xd0\xf4\xd7\x34\x5c\xfd\xdb\x46\xeb\x7c\x75\x40\xd6\x81\x04\x87\xe0\x1f\x95\x36\xce\xa6\xa6\x56\x30\xa5\x59\xd2\x7e\xe9\x7e\x70\x83\x69\xef\xd7\xb7\xc1\x7f\x02\x00\x00\xff\xff\xcc\x41\xdc\x32\xaf\x14\x00\x00")

func jsRunJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRunJs,
		"js/run.js",
	)
}

func jsRunJs() (*asset, error) {
	bytes, err := jsRunJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/run.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRun_mockJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xcd\x4a\x03\x51\x0c\x85\xf7\xf7\x29\x0e\x5d\x75\x68\x4d\x5d\x0a\xd2\x9d\x08\xba\xf6\x05\xee\x5c\xe2\xf4\xb6\x33\xc9\x90\xe4\x52\x41\x7c\x77\x19\xb4\x3f\xdb\x93\x2f\xe7\x7c\x69\xb7\xc3\xc7\xa1\x3a\xe6\x5c\x4e\x79\x60\x4c\x5a\x4e\x8e\x38\x30\xac\x09\x1d\x1d\x9f\x4d\x4a\x54\x15\x27\xbc\x05\x4a\x16\x18\xcf\x63\x2e\x37\x40\x0d\xc1\x1e\x55\x06\x4a\xe9\x82\xe3\x9c\x6b\xbc\xaa\xbd\x6b\xbf\x3e\x6a\xdf\xe1\x3b\x01\x45\xc5\x75\x64\x1a\x75\x58\x42\x92\x3c\x71\x97\x12\x60\x1c\xcd\x04\x61\x8d\xd3\xcf\x5d\x89\x35\x59\xc0\x2d\xf8\xaf\xe0\x9f\xbb\xfc\x62\x83\xd5\xc3\x0a\x1b\xbc\xe4\x60\x12\x3d\xaf\xbb\x6b\xc4\x54\x74\x9a\x6a\x90\xb7\xde\xc3\xaa\x0c\xeb\xc7\x2d\x9e\xba\xe7\x65\x80\xbf\x66\xb5\x70\xba\x49\x62\x7f\x67\x7c\xbd\x5b\x13\xec\x17\x8b\xf4\x1b\x00\x00\xff\xff\xf2\x0f\x52\xe5\x29\x01\x00\x00")

func jsRun_mockJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRun_mockJs,
		"js/run_mock.js",
	)
}

func jsRun_mockJs() (*asset, error) {
	bytes, err := jsRun_mockJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/run_mock.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRunnerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x8a\x14\x31\x10\x86\xef\x79\x8a\xdf\x3e\x75\xe3\xd2\xfb\x04\x73\x10\x5d\x51\xf0\xe8\x4d\x44\x42\x52\xdd\x1d\x08\x55\x6d\x55\x7a\x7b\x17\xd9\x77\x97\x74\xdc\x71\x66\xf0\xe0\x31\xf9\xc9\x97\xbf\x3e\xea\xfe\x1e\x5f\x97\x64\x48\x86\xb2\x10\x74\x63\x26\xc5\xae\x7e\x5d\x13\xcf\xb0\xa0\x69\x2d\xa3\x0b\xc2\x26\x99\xc6\x2c\x73\xdf\x7d\x11\x1f\x6b\xf8\xee\xfd\xe7\x0f\x08\xa2\xd4\x0d\xce\x4d\x1b\x87\x92\x84\x31\x25\xa5\x87\x47\xe2\xd2\xd3\x80\x5f\x0e\x78\xf4\x0a\xaa\x17\x9f\x3c\xc7\x4c\x8a\x13\x98\x76\x3c\x5c\x5c\xf5\x83\x73\xc0\x6b\x97\x59\xc8\x90\xf8\x28\x64\x41\x56\xaa\x87\x7d\x49\x61\xc1\xd4\x0f\xb5\x2a\x3d\x51\xd8\x0a\x45\x98\xa0\x2c\xbe\x60\x27\x04\xcf\xf0\x21\x90\x59\x43\xd5\xd7\xa2\x69\x4e\xec\x73\xfb\x1f\x7b\x2a\x8b\x6c\x05\x4a\x3f\xb7\xa4\x75\x84\x5e\xf4\x0e\x4a\x3e\xe7\xe7\x3b\xf8\x9c\x65\x4f\x3c\x0f\xc7\xdb\xcd\x48\x51\xa4\xc1\x56\x6f\x4d\x50\x03\x79\x95\x8d\xe3\xe8\x00\x7a\x5a\x45\x8b\x8d\x3f\x5a\x70\x02\xd5\x49\xd2\x84\xfe\x8d\xd2\x9c\xac\x90\x1e\x83\x5a\x73\x01\x5c\x99\x64\xf9\x03\x5c\x9a\x07\x43\xa4\x29\x31\xc5\xee\x10\x02\x28\x95\x4d\xd9\x01\x2f\xf5\x7c\x4d\xec\x2f\xa5\x0e\xee\x06\x7d\x84\x86\x2c\x3e\x52\x1c\xf1\xb1\xcd\xdb\xe1\x2d\x68\x2c\xcf\x2b\x0d\xe7\x9e\x97\x98\x6f\x2d\xfc\xfe\x5f\x6d\xcf\x7d\x28\x62\x12\xbd\x85\xdf\xb6\xaf\x7b\x30\x55\x43\xff\xf8\xaf\xe6\x53\x4f\x83\x7b\x71\xee\x55\xe9\x79\x91\x70\xfa\xbb\x54\xee\x77\x00\x00\x00\xff\xff\x0b\x8b\xea\xff\xb1\x02\x00\x00")

func jsRunnerJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRunnerJs,
		"js/runner.js",
	)
}

func jsRunnerJs() (*asset, error) {
	bytes, err := jsRunnerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/runner.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsWaitgroupJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xb1\x8e\x9c\x30\x10\x86\x7b\x3f\xc5\x5f\x82\x14\x99\x27\xd8\x32\x8a\x92\x22\x4d\x22\xa5\x88\x52\x18\x18\x62\x73\xc6\x83\xec\xe1\xd8\xd5\x89\x77\x3f\xd9\xbb\x2c\x48\x57\x5e\x61\xad\xd6\x78\xbe\xff\x9b\xb1\x9b\x06\x7f\x8c\x93\x6f\x91\x97\x19\xab\x71\x92\x30\x70\xc4\xb4\x78\x71\xb3\x27\x8c\xdc\x26\x08\x63\x70\xc1\x25\xab\xf1\x5d\xb0\x3a\xef\x21\x36\xf2\x0a\x13\x40\x31\x72\x54\x4d\x03\x93\x90\x98\x43\xfe\x35\xb9\x0c\x91\x66\x8e\x92\x9e\x87\xb4\x1a\x96\xd0\x89\xe3\x70\x44\x56\x35\xde\x14\x20\xd6\x25\x5d\xa2\x2e\xf8\xfb\x4f\x29\x20\x03\xfb\x3e\xaf\x8c\x0b\xb4\x16\xa4\x30\xc4\x52\xf1\xfc\x9f\xcb\xf7\xd2\x7c\xf4\x82\x1d\x5f\x8d\xdc\xde\xb9\x27\xb2\x9e\x97\x64\xcb\x17\x05\x6c\x8f\x88\xb8\x84\xbc\x12\xe8\x95\xe2\xad\x44\xb8\x50\x22\x0a\xfe\x0b\x4c\xe8\xf3\xdf\x70\x1a\x8d\x58\x9a\xb2\x48\xc7\xd3\xec\x49\x48\xef\x12\x19\x76\x92\xf8\x68\x30\x70\xfc\x6a\x3a\x5b\x3d\xc7\x50\x8d\xfb\x21\x60\xd4\xad\xe9\x5e\x72\x6c\xe8\xab\xba\x6c\x6e\xf5\x01\xc8\xf9\xd5\x59\x3d\x6f\x3c\xac\x96\x20\xce\xdf\x6f\xca\x44\x3a\xc4\xf0\x93\x85\x20\xd6\x48\x61\xa0\x67\x4a\x08\x2c\xa5\xed\xdc\x64\xb1\xc2\x6f\x4b\xb7\x3b\x73\x5a\x92\xa0\x25\x24\x31\x51\xa8\x07\x5d\x85\x62\x30\xde\xdf\x34\xaa\x5f\x44\xc7\xbd\x95\x66\x39\xe2\x07\xb7\x27\xef\x5a\x9d\x6c\x3f\x35\x8b\x67\xbb\x8f\x29\x6c\x6a\x53\x8a\xae\xe5\x41\xe9\xe3\xc1\x5e\x0e\x23\xf5\x1e\x00\x00\xff\xff\x60\x6f\x0e\xe7\xca\x02\x00\x00")

func jsWaitgroupJsBytes() ([]byte, error) {
	return bindataRead(
		_jsWaitgroupJs,
		"js/waitgroup.js",
	)
}

func jsWaitgroupJs() (*asset, error) {
	bytes, err := jsWaitgroupJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/waitgroup.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"js/event.js": jsEventJs,
	"js/job.js": jsJobJs,
	"js/mock8s.js": jsMock8sJs,
	"js/run.js": jsRunJs,
	"js/run_mock.js": jsRun_mockJs,
	"js/runner.js": jsRunnerJs,
	"js/waitgroup.js": jsWaitgroupJs,
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
	"js": &bintree{nil, map[string]*bintree{
		"event.js": &bintree{jsEventJs, map[string]*bintree{}},
		"job.js": &bintree{jsJobJs, map[string]*bintree{}},
		"mock8s.js": &bintree{jsMock8sJs, map[string]*bintree{}},
		"run.js": &bintree{jsRunJs, map[string]*bintree{}},
		"run_mock.js": &bintree{jsRun_mockJs, map[string]*bintree{}},
		"runner.js": &bintree{jsRunnerJs, map[string]*bintree{}},
		"waitgroup.js": &bintree{jsWaitgroupJs, map[string]*bintree{}},
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

