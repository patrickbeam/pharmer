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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x5f\x4f\xdb\x3c\x14\xc6\xef\xfb\x29\x8e\x72\x9d\xb7\x2f\x49\x4b\x55\xe5\x2e\x84\x6c\x2a\x0c\x16\xd1\x32\x54\x4d\x28\x72\x93\x03\x33\x4d\x6c\xcf\x76\x0b\x1d\xea\x77\x9f\x9c\xd0\xa4\xa4\xa1\x14\xb4\x9b\xfe\x39\xc7\x7e\x9e\xdf\x89\x73\x7c\x9e\x3b\x00\x16\x23\x39\x5a\x1e\x58\x82\x24\x73\xd4\x96\x6d\x62\xc8\x96\x96\x07\x3f\x3b\x00\x00\x56\x8a\xcb\x22\x0a\x60\xfd\x26\x9b\x5f\x42\xf2\xd4\xea\x00\xdc\x16\xeb\x25\xde\x53\xce\x54\xb5\xe7\xb9\xf8\x04\xb0\x32\x9e\x10\x4d\x39\x33\x06\x97\xd3\x00\x2e\x50\x4b\x6e\xc3\xe5\xd4\x86\xeb\xb1\xff\x22\x56\x09\x98\x55\xd7\x63\x08\x89\xd2\x75\xea\x0f\x67\x58\x2b\x17\x21\x7c\x94\x8e\xf5\xf2\xf7\xb6\xf8\x5e\xdb\x6f\xfb\x8e\x09\x83\x33\xae\xd0\x86\xc0\xdf\x67\x7b\x83\xef\xd8\xaa\x87\xe4\x03\xb6\x7e\xae\x34\xca\x94\xe4\x36\x5c\x7e\x6b\xb3\x0c\xaf\xe1\x26\x1c\x4f\xf6\x5a\x92\x5c\x7d\xc0\x72\xc2\xe7\x2b\x6e\xc3\x59\xd4\x66\xe7\x8f\x47\x3e\x84\xfe\x3b\x86\x4c\xea\xa6\x61\x75\xca\x94\x29\x4d\x58\x82\x93\x95\xc0\x96\xb3\x56\xf3\x85\xf1\x99\x11\x89\x39\x6a\x92\xc5\x47\xb5\x53\x8a\x2a\x91\x54\x54\xa0\xd3\x28\x84\x23\x0f\xfa\xf0\x34\x1c\xc0\xa0\x3f\xa3\x1a\x02\x2e\x51\xd9\x30\xfc\x7a\x02\xa7\xa7\x57\x3d\xb8\xf2\x2f\xea\xfd\x09\xd1\x78\xcf\xe5\xca\x6c\x3e\x21\x12\xcd\x8b\x44\xb2\xad\xbc\x30\xde\xfd\xba\x6c\x92\x5b\x1e\x0c\x6b\x7f\xaa\xe6\x26\x70\xd4\xfa\x0c\x77\xc8\x9d\xfd\xe4\x4e\x2b\x79\xcf\xfd\x87\xe8\x3d\xb7\xc9\xee\xb8\x07\xc2\xbb\xfb\xe1\x5d\x0f\xdc\x16\x7a\xf7\x78\x50\xe2\xf7\x21\x0c\x82\xcf\x94\xe0\x36\x6b\x70\x8f\x07\xcd\x22\xdc\xe1\xa1\x47\xd0\xdb\x5f\x45\xcf\x03\x67\xb0\x5b\x85\xe3\x0e\x37\x55\x7c\xa2\x02\x67\xd0\xa8\xc0\x71\x77\x5e\xa1\xea\x18\xaa\xb6\x48\x24\xa6\xc8\x34\x25\x59\x4b\x53\x08\xc9\x97\x34\x45\x69\x8c\xa3\xfa\x86\xdd\x28\x8a\x8c\xac\xbe\x70\x99\x13\x6d\x16\xdc\x51\xcc\xd2\x3a\x4f\x18\xe3\xba\x68\x6e\x23\xfc\x5c\x37\xa9\xf8\x45\x64\x8e\xb2\x4b\x84\x50\x09\x4f\xb1\x9b\xf0\xfc\xff\x24\x5b\x98\x0b\xe7\xbf\x1a\xc7\x48\x6e\x7a\x79\x5d\xa9\x16\x26\xaf\xbb\xbe\x96\x2e\xef\xfe\x84\xb3\x3b\x7a\x5f\x20\xfb\xc1\x79\x38\x89\xa3\xab\xef\x67\x61\x30\x89\x47\xa7\x15\x5d\xa9\xc5\x65\x5e\xcf\x8e\x58\x48\xfe\x80\x89\x8e\x69\xfa\x7a\xd9\x83\x2a\x4f\xee\x25\xdf\x54\xc9\xc8\x0c\x0b\xd8\xa8\xcc\xc3\xa8\xb1\x9f\x32\xb1\x28\x1e\x90\xc6\x27\x6d\x55\x99\xb5\x7d\x78\x05\x7e\x34\x8a\xcf\xc3\xe9\x5e\x7c\x22\x68\x3c\xc7\x55\x3b\x3b\x11\xf4\xbc\x99\xab\xc0\xfd\x68\x04\x3b\xd9\x8a\x5a\x10\xa5\x1e\xb9\x4c\xb7\xc8\xdf\xb8\x61\xe7\x8b\x19\x4a\x86\x1a\xd5\x0f\x94\xaa\x7d\xa4\x2e\xcb\x8c\x11\x76\xba\xc3\xee\xdb\x97\x6c\x23\x5b\xce\xf4\xad\xb7\xc8\xcc\x75\x0f\xb4\x5c\x60\x8d\x6d\x26\xfc\x4e\xac\x98\xf5\x65\xb4\xb3\xcd\x5f\x70\x77\xd6\x9d\xbf\x01\x00\x00\xff\xff\xcc\x1a\xeb\x5b\x47\x08\x00\x00")

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

	info := bindataFileInfo{name: "cloud.json", size: 2119, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
