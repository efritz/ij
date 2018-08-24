// Code generated by go-bindata.
// sources:
// asset.go
// config.yaml
// DO NOT EDIT!

package schema

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

var _assetGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\xce\x48\xcd\x4d\xe4\xe2\xd2\xd7\x4f\xcf\xb7\x4a\x4f\xcd\x4b\x2d\x4a\x2c\x49\x55\x48\xcf\xd7\x4d\xca\xcc\x4b\x49\x2c\x49\x54\xd0\x2d\xc8\x4e\x87\xaa\x52\xd0\xe3\x02\x04\x00\x00\xff\xff\x26\x1f\xf5\x4c\x37\x00\x00\x00")

func assetGoBytes() ([]byte, error) {
	return bindataRead(
		_assetGo,
		"asset.go",
	)
}

func assetGo() (*asset, error) {
	bytes, err := assetGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "asset.go", size: 55, mode: os.FileMode(420), modTime: time.Unix(1534806267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\xcb\xae\x9b\x30\x10\xdd\xf3\x15\x48\xad\x94\x15\xba\xea\x96\x75\x97\x95\xd2\x45\xf7\xc8\xb1\x07\x70\xe3\x07\x1d\x9b\xdb\x9b\xbf\x2f\x26\x82\x98\x62\x02\x38\x49\xeb\x9d\x1f\x33\xc7\x73\x66\xe6\xd8\x59\x96\x25\x09\x83\x92\x2b\x6e\xb9\x56\x26\x4f\xd2\x14\xd4\x3b\x47\xad\x24\x28\xeb\xa6\x69\xca\xc0\x50\xe4\x8d\x3b\x90\xa7\x3f\x8e\x5f\x8f\xfd\xaa\xbd\x34\x90\xa7\x04\x91\x5c\xfa\x39\xb7\x20\xcd\xd5\x60\xd8\x34\x16\xb9\xaa\x92\x6e\xcd\x12\x73\x5e\x77\xa6\x4f\x3f\x81\xda\x7e\xa1\x41\xdd\x00\x5a\x0e\xa3\x4b\xf8\xb0\xa0\xd8\x38\x5d\xf2\x34\x43\xbf\x2e\x71\x49\x2a\x88\xb4\xa5\x5a\x4a\xa2\x58\xa4\xb5\xa9\x41\x88\x58\xdb\xfe\x68\xa4\x71\x97\x3f\xbc\x34\x9a\xab\x58\x07\xad\x01\x8c\x34\xfd\xad\xf1\x6c\x1a\x42\x63\x29\xaf\xb5\xb1\x8a\xc8\x58\x73\x06\x96\xd0\x7a\xbb\xf1\x49\x6b\x01\x44\x0d\xe0\x40\x84\xad\x69\x0d\xf4\x7c\x73\xf1\x19\xa1\xcc\xd3\xc3\xa7\x37\xaf\x5b\xde\xbc\x93\x87\x91\xf5\xbf\xba\x67\xd1\xd8\x3b\x39\x18\x23\xfc\x6a\x39\x02\x2b\x82\x5e\xd6\xa2\xb8\xf5\xa2\x1b\x93\x7e\x5c\xe0\x09\x3e\x1a\x8d\xd6\x47\x2b\x4a\x2e\x36\xb1\xae\x15\x1c\x4b\xdf\x7d\x16\x02\xf0\x37\xa6\xd7\x0b\x5e\x31\x78\x4d\xc2\x58\xcf\x18\x11\xdf\x6f\xa2\x90\x96\x44\x18\x70\xd2\x32\xcb\x56\xb4\xc2\x3c\xd6\xe7\x5d\x9f\x01\xbe\x93\xd8\x56\x47\xe8\x66\xb0\x43\xde\x1c\x5e\x05\x38\x28\x85\x25\x5d\x22\xbb\x80\xb8\x8e\x0d\xc0\x72\x09\xba\x8d\x13\x8b\x95\x1c\x35\x82\xa8\xa7\xc8\x7f\x6c\x4b\x77\xf4\x54\x7b\xc8\x5d\x69\xa5\x60\x3f\xf7\x18\x4f\x90\x81\x15\x2e\x7b\x98\x07\xc9\x7c\x40\x5a\x4f\x50\x6a\x84\xc2\xbb\xc5\x6e\x17\xa4\xec\x3a\xe5\x21\x0f\xee\x3b\xf1\x2f\xb2\x59\x38\xa0\xc3\x28\x0f\x8a\xb6\x88\xbb\x24\x79\x5a\x85\x2f\x2a\x0a\xb7\x3f\xbc\x1c\x83\xdf\xac\xcf\xf1\x58\x2f\xc5\xda\x07\x6c\x22\xe7\x41\x29\xcf\xe6\x85\xe5\xc6\xbc\xb8\xdc\x98\x16\xd8\x7d\xa6\x16\x72\xbc\x48\xd8\x76\xd2\xb6\x10\x17\x22\xcf\x27\x70\x12\xf4\x34\xd8\xc9\x8f\x74\x16\x81\xe4\xea\x1b\xa8\xca\xd6\x79\xfa\x25\x09\x84\xb2\x25\x04\x2e\xdd\xfb\xfc\x9c\xbc\xed\xfb\x1d\x5c\x7f\x06\xff\x07\x99\x8a\x96\xdd\x55\xb8\x57\x41\x7b\xb2\xb2\x51\x5a\x83\xd5\x95\xdc\xa9\xd1\x41\x50\xdc\x8b\xf8\x62\x28\x07\x71\x48\xee\xd5\xff\x9f\x00\x00\x00\xff\xff\xf4\xcb\x29\xdb\xfe\x0d\x00\x00")

func configYamlBytes() ([]byte, error) {
	return bindataRead(
		_configYaml,
		"config.yaml",
	)
}

func configYaml() (*asset, error) {
	bytes, err := configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.yaml", size: 3582, mode: os.FileMode(420), modTime: time.Unix(1535079220, 0)}
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
	"asset.go": assetGo,
	"config.yaml": configYaml,
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
	"asset.go": &bintree{assetGo, map[string]*bintree{}},
	"config.yaml": &bintree{configYaml, map[string]*bintree{}},
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

