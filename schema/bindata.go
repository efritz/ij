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

var _configYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\xc1\x8e\x9b\x30\x10\xbd\xf3\x15\x96\x5a\x29\x27\xb4\xea\x95\x73\x8f\x95\xd2\x43\xef\x91\x17\x4f\xc0\x5d\x3c\xa6\xe3\xc9\x6a\xf3\xf7\xc5\x8e\x20\x50\x1c\x20\x4e\xd3\xe6\x66\xe3\x99\x97\x79\xef\xcd\xd8\x79\x9e\x67\x99\x82\xa3\x46\xcd\xda\xa2\x2b\x32\x21\x00\xdf\x35\x59\x34\x80\xec\x97\x42\x28\x70\x25\xe9\xd6\x1f\x28\xc4\x8f\xfd\xd7\x7d\xd8\xe5\x73\x0b\x85\x90\x44\xf2\x1c\xd6\x9a\xc1\xb8\x4b\x40\xff\xd1\x31\x69\xac\xb2\x6e\x8f\xa5\x7b\x5b\x4f\x66\x5f\x7f\x42\xc9\x61\xa3\x25\xdb\x02\xb1\x86\x21\x25\x7c\x30\xa0\x1a\x96\xb7\x32\xcd\xd0\x2f\x5b\xda\xc8\x0a\x12\x63\x4b\x6b\x8c\x44\x95\x18\xed\x6a\x68\x9a\xd4\xd8\x70\x34\x31\xb8\xd3\x8f\xce\xad\xd5\x98\x9a\xa0\xb6\x8e\x51\x9a\x54\xda\x14\xb0\x2c\xeb\xed\xc1\xaf\xd6\x36\x20\xb1\x07\x07\xd9\x70\x5d\xd6\x50\xbe\x5d\x53\x7c\x26\x38\x16\x62\xf7\xe9\x65\x64\xd9\x97\xd1\xc9\xdd\x50\xfa\x1f\x16\xbe\x19\x3c\x3a\xd9\x07\x13\xfc\x3a\x69\x02\x75\x88\x66\x59\xab\xe2\xda\x10\xfe\x37\x69\x8a\x28\x4f\x52\xa9\xf0\x57\x64\xf3\xfd\x6a\x79\x71\x94\x8d\x03\xdf\x38\x33\x1a\x92\xfb\xe7\x31\x17\x77\x2e\x02\x7a\x97\xa9\x46\x26\xe8\x56\x70\x47\xf3\x7a\xbc\x0a\xa8\xef\x03\x96\xc4\x87\xae\x20\x6d\x53\x0b\x60\x6d\xc0\x9e\xd2\x5a\x61\x45\xa3\xb6\x91\xf8\xa0\x38\x5d\x81\xd5\x3d\xf4\xac\xb8\x2c\x6a\xf5\x80\xf1\x17\x3a\x64\x85\x8d\x00\xf3\x20\x1d\x0f\x4c\x1d\x7f\xd3\xfc\x0b\x22\x0f\x1e\x68\x37\xf4\x16\x96\x27\xa2\xbb\x06\xc5\x74\xdc\x3d\x49\x0f\xff\xbd\x9f\x67\x7d\xde\x3c\xd0\x3b\x48\x75\x58\xbb\x9b\x2d\xc2\xfe\x78\x0d\x8e\x70\x9e\xcf\x35\x8d\xeb\x3a\xd7\x76\x99\xa9\x1b\x1a\xdf\x24\x6c\x3b\x69\x5b\x88\x8b\x91\x37\x26\x70\x52\xf4\xb4\xd8\xc9\x63\x65\x56\x81\xd1\xf8\x0d\xb0\xe2\xba\x10\x5f\x62\x0f\xae\x2d\x25\x68\xd3\x5a\x5a\x7c\xa0\x6d\xd7\xed\xbe\x3b\x0b\x3e\xfe\x1f\x72\xd9\x9c\xd4\xe2\x70\x79\x16\xf4\x68\xac\x6c\x9c\x6a\x51\x77\x65\x0b\x1e\xed\x07\x8a\xbf\x4e\x9e\x0c\xe5\x21\x76\xd9\x92\xff\x7f\x07\x00\x00\xff\xff\x98\x98\xca\x81\x19\x0c\x00\x00")

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

	info := bindataFileInfo{name: "config.yaml", size: 3097, mode: os.FileMode(420), modTime: time.Unix(1534983731, 0)}
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

