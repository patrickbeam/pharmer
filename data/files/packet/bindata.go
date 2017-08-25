// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package packet

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x41\x6f\xe2\x3c\x10\xbd\xf3\x2b\x46\x3e\x47\x55\x63\x68\x3e\x94\x5b\x4a\xf3\xad\xda\xdd\x56\xa8\xb4\xaa\xd0\xaa\x42\x26\x99\xad\x42\x12\x3b\xb2\x0d\xdb\x2c\xe2\xbf\xaf\x9c\x40\x02\x21\xa4\xb4\xda\x0b\xc1\x7e\x9e\x79\xef\xd9\x9e\xf1\xba\x07\x40\x38\x4b\x91\xb8\x40\x32\x16\xc4\xa8\x89\x65\xe6\x90\xaf\x88\x0b\x3f\x7b\x00\x00\x24\xc4\x15\xe9\x01\xbc\x16\x88\xc4\xb7\x48\x70\x55\xa1\xeb\xe2\x17\x80\x24\x22\x60\x3a\x12\xdc\xa4\x7a\x98\x8e\xe0\x1e\xb5\x14\x16\x3c\x4c\x2d\x78\x9e\x78\x45\xda\x62\x5d\x99\xc0\xac\x7a\x9e\x80\xcf\x94\xae\xa1\x3f\x82\x63\x9d\xb9\x98\xc2\xdf\xd2\x26\xdb\xe1\x6b\xf1\xdd\x58\xa7\x79\x27\x8c\xc3\x9d\x50\x68\xc1\xc8\xeb\xa2\x7d\xc1\x0f\x68\xd5\x22\xf8\x04\xad\x97\x2a\x8d\x32\x64\xa9\x05\x0f\x3f\xda\x28\xfd\x67\x78\xf1\x27\x4f\x9d\x94\x2c\x55\x9f\xa0\x7c\x12\x71\x2e\x2c\xb8\x1b\xb7\xd1\x79\x93\x5b\x0f\x7c\xef\x03\x42\x2e\x75\x93\xb0\x3a\xe5\x88\x2b\xcd\x78\x80\x4f\x79\x86\x2d\x67\xad\xe2\xa5\xe1\x99\x33\x89\x29\x6a\x96\xcc\x2e\x6b\xa6\x10\x55\x20\xa3\xac\x12\x3a\x1d\xfb\x70\xe9\xc2\x00\xde\x87\x0e\x38\x83\x79\xa4\x61\x24\x24\x2a\x0b\x86\xdf\xae\xe1\xe6\xe6\xb1\x0f\x8f\xde\x7d\x1d\x1f\x30\x8d\x6f\x42\xe6\x26\xf8\x9a\x49\x34\x17\x89\x25\x7b\x78\x66\xb8\x07\xb5\x6d\x96\x12\x17\x86\x35\x7f\xa4\x62\x33\x71\xd9\xba\x87\x47\xca\xed\x6e\xe5\x76\xab\xf2\x3e\xfd\x87\xd2\xfb\xb4\xa9\xdd\xa6\x67\x8a\xa7\xdd\xe2\xa9\x0b\xb4\x45\x3d\xbd\x72\x4a\xf9\x03\xf0\x47\xa3\xaf\x58\xa0\x4d\x0f\xf4\xca\x69\x9a\xa0\xc3\x73\x8f\xa0\xdf\xed\xa2\xef\x82\xed\x1c\xbb\xb0\xe9\x70\xe7\xe2\x0b\x0e\x6c\xa7\xe1\xc0\xa6\x47\x57\xa8\x3a\x86\xaa\x2c\x02\x89\x21\x72\x1d\xb1\xa4\xa5\x28\x74\x9e\x15\x7d\x74\x5c\xf7\xd1\x5d\xb6\x2c\x61\xf9\xff\x42\xa6\x4c\x9b\x05\xbf\x22\x4c\xc2\x1a\x2f\x86\x87\xb5\xb9\xae\xfe\x01\x90\x85\x2a\x37\x83\x65\xd1\x2c\xc6\xbc\x0a\x2c\x8b\x78\xdb\xbc\x5b\xc1\x84\xcd\x31\x29\x1a\xc2\xf8\x16\xbe\x37\xd1\x88\x67\xcb\x42\x8f\xc6\x77\x4d\x2a\x64\x63\x75\xcb\xc8\xa4\x58\x60\xa0\x67\x51\xd8\xae\xe4\x14\x5e\x89\x19\x97\x0b\xe0\x36\x3c\x4f\xcf\x89\x16\x15\x2f\xe7\x28\x39\xea\xa2\x3f\x95\x4a\xc9\x0a\xa5\x3a\x78\x9d\xf6\x3d\xec\x50\x43\x61\x5f\xfc\x77\x31\xd8\xa3\x6f\x5e\xbc\x23\x5c\x0b\x91\xd4\x44\xdb\x49\x23\x81\x85\x69\x1d\xd0\xb6\x89\xdb\xb7\xf4\x20\xd0\xbc\xa8\x2e\x68\xb9\xc4\x23\x9b\xe5\xd7\x98\xdd\xf4\x36\x7f\x03\x00\x00\xff\xff\x3f\xda\x11\x35\xa0\x07\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 1952, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": {cloudJson, map[string]*bintree{}},
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