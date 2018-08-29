// Code generated by go-bindata.
// sources:
// schema/build.yaml
// schema/config.yaml
// schema/ecr-registry.yaml
// schema/gcr-registry.yaml
// schema/logout.yaml
// schema/metaplan.yaml
// schema/override.yaml
// schema/plan.yaml
// schema/push.yaml
// schema/run.yaml
// schema/server-registry.yaml
// schema/task.yaml
// DO NOT EDIT!

package asset

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
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

var _schemaBuildYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  type:
    description: TODO
    type: string
    enum:
      - build
  extends:
    description: TODO
    type: string
  dockerfile:
    description: TODO
    type: string
  tags:
    description: TODO
    oneOf:
      - type: string
      - type: array
        items:
          type: string
  labels:
    description: TODO
    oneOf:
      - type: string
      - type: array
        items:
          type: string
  arguments:
    description: TODO
    oneOf:
      - type: string
      - type: array
        items:
          type: string
additionalProperties: false
`)

func schemaBuildYamlBytes() ([]byte, error) {
	return _schemaBuildYaml, nil
}

func schemaBuildYaml() (*asset, error) {
	bytes, err := schemaBuildYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/build.yaml", size: 629, mode: os.FileMode(420), modTime: time.Unix(1535420764, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaConfigYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  extends:
    description: TODO
    type: string
  ssh-identities:
    description: TODO
    oneOf:
      - type: string
      - type: array
        items:
          type: string
  environment:
    description: TODO
    type: array
    items:
      type: string
  registries:
    description: TODO
    type: array
    items:
      description: TODO
      type: object
  import:
    description: TODO
    type: object
    properties:
      files:
        description: TODO
        oneOf:
          - type: string
          - type: array
            items:
              type: string
      excludes:
        description: TODO
        oneOf:
          - type: string
          - type: array
            items:
              type: string
    additionalProperties: false
  export:
    description: TODO
    type: object
    properties:
      files:
        description: TODO
        oneOf:
          - type: string
          - type: array
            items:
              type: string
      excludes:
        description: TODO
        oneOf:
          - type: string
          - type: array
            items:
              type: string
    additionalProperties: false
  workspace:
    description: TODO
    type: string
  tasks:
    description: TODO
    type: object
  plans:
    description: TODO
    type: object
  metaplans:
    description: TODO
    type: object
additionalProperties: false
`)

func schemaConfigYamlBytes() ([]byte, error) {
	return _schemaConfigYaml, nil
}

func schemaConfigYaml() (*asset, error) {
	bytes, err := schemaConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/config.yaml", size: 1453, mode: os.FileMode(420), modTime: time.Unix(1535577570, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaEcrRegistryYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  type:
    description: TODO
    type: string
    enum:
      - ecr
  access_key_id:
    description: TODO
    type: string
  secret_access_key:
    description: TODO
    type: string
  account_id:
    description: TODO
    type: string
  region:
    description: TODO
    type: string
  role:
    description: TODO
    type: string
additionalProperties: false
`)

func schemaEcrRegistryYamlBytes() ([]byte, error) {
	return _schemaEcrRegistryYaml, nil
}

func schemaEcrRegistryYaml() (*asset, error) {
	bytes, err := schemaEcrRegistryYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/ecr-registry.yaml", size: 422, mode: os.FileMode(420), modTime: time.Unix(1535494834, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGcrRegistryYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  type:
    description: TODO
    type: string
    enum:
      - gcr # TODO - make region configurable
  key_file:
    description: TODO
    type: string
additionalProperties: false
`)

func schemaGcrRegistryYamlBytes() ([]byte, error) {
	return _schemaGcrRegistryYaml, nil
}

func schemaGcrRegistryYaml() (*asset, error) {
	bytes, err := schemaGcrRegistryYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/gcr-registry.yaml", size: 242, mode: os.FileMode(420), modTime: time.Unix(1535494795, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaLogoutYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  type:
    description: TODO
    type: string
    enum:
      - logout
  extends:
    description: TODO
    type: string
  servers:
    description: TODO
    oneOf:
      - type: string
      - type: array
        items:
          type: string
additionalProperties: false
`)

func schemaLogoutYamlBytes() ([]byte, error) {
	return _schemaLogoutYaml, nil
}

func schemaLogoutYaml() (*asset, error) {
	bytes, err := schemaLogoutYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/logout.yaml", size: 333, mode: os.FileMode(420), modTime: time.Unix(1535481404, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaMetaplanYaml = []byte(`---

title: TODO
description: TODO
type: array
items:
  description: TODO
  type: string
`)

func schemaMetaplanYamlBytes() ([]byte, error) {
	return _schemaMetaplanYaml, nil
}

func schemaMetaplanYaml() (*asset, error) {
	bytes, err := schemaMetaplanYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/metaplan.yaml", size: 89, mode: os.FileMode(420), modTime: time.Unix(1535420628, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaOverrideYaml = []byte(`---

# TODO - other options such as workspace configs, force sequential, etc

title: TODO
description: TODO
type: object
properties:
  registries:
    description: TODO
    type: array
    items:
      description: TODO
      type: object
  ssh-identities:
    description: TODO
    oneOf:
      - type: string
      - type: array
        items:
          type: string
  environment:
    description: TODO
    type: array
    items:
      type: string
  import:
    description: TODO
    type: object
    properties:
      excludes:
        description: TODO
        oneOf:
          - type: string
          - type: array
            items:
              type: string
    additionalProperties: false
  export:
    description: TODO
    type: object
    properties:
      excludes:
        description: TODO
        oneOf:
          - type: string
          - type: array
            items:
              type: string
    additionalProperties: false
additionalProperties: false`)

func schemaOverrideYamlBytes() ([]byte, error) {
	return _schemaOverrideYaml, nil
}

func schemaOverrideYaml() (*asset, error) {
	bytes, err := schemaOverrideYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/override.yaml", size: 977, mode: os.FileMode(420), modTime: time.Unix(1535577563, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaPlanYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  extend:
    description: TODO
    type: boolean
  stages:
    description: TODO
    type: array
    items:
      description: TODO
      type: object
      properties:
        name:
          description: TODO
          type: string
        before_stage:
          description: TODO
          type: string
        after_stage:
          description: TODO
          type: string
        tasks:
          description: TODO
          type: array
          items:
            stage_task:
              description: TODO
              oneOf:
                - type: string
                - type: object
                  properties:
                    name:
                      description: TODO
                      type: string
                    environment:
                      description: TODO
                      type: array
                      items:
                        type: string
                  additionalProperties: false
                  required:
                    - name
        run-mode:
          description: TODO
          type: string
          enum:
            - always
            - on-success
            - on-failure
        concurrent:
          description: TODO
          type: boolean
        environment:
          description: TODO
          type: array
          items:
            type: string
      additionalProperties: false
      required:
        - name
  environment:
    description: TODO
    type: array
    items:
      type: string
additionalProperties: false
`)

func schemaPlanYamlBytes() ([]byte, error) {
	return _schemaPlanYaml, nil
}

func schemaPlanYaml() (*asset, error) {
	bytes, err := schemaPlanYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/plan.yaml", size: 1583, mode: os.FileMode(420), modTime: time.Unix(1535464102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaPushYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  type:
    description: TODO
    type: string
    enum:
      - push
  extends:
    description: TODO
    type: string
  images:
    description: TODO
    oneOf:
      - type: string
      - type: array
        items:
          type: string
additionalProperties: false
`)

func schemaPushYamlBytes() ([]byte, error) {
	return _schemaPushYaml, nil
}

func schemaPushYaml() (*asset, error) {
	bytes, err := schemaPushYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/push.yaml", size: 330, mode: os.FileMode(420), modTime: time.Unix(1535420756, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaRunYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  type:
    description: TODO
    type: string
    enum:
      - run
  extends:
    description: TODO
    type: string
  image:
    description: TODO
    type: string
  command:
    description: TODO
    type: string
  shell:
    description: TODO
    type: string
  script:
    description: TODO
    type: string
  entrypoint:
    description: TODO
    type: string
  user:
    description: TODO
    type: string
  workspace:
    description: TODO
    type: string
  hostname:
    description: TODO
    type: string
  detach:
    description: TODO
    type: boolean
  healthcheck:
    description: TODO
    type: object
    properties:
      command:
        description: TODO
        type: string
      interval:
        description: TODO
        type: string
      retries:
        description: TODO
        type: integer
      start_period:
        description: TODO
        type: string
      timeout:
        description: TODO
        type: string
    additionalProperties: false
  environment:
    description: TODO
    type: array
    items:
      type: string
  required_environment:
    description: TODO
    type: array
    items:
      type: string
  export_environment_file:
    description: TODO
    oneOf:
      - type: string
      - type: array
        items:
          type: string
additionalProperties: false
`)

func schemaRunYamlBytes() ([]byte, error) {
	return _schemaRunYaml, nil
}

func schemaRunYaml() (*asset, error) {
	bytes, err := schemaRunYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/run.yaml", size: 1388, mode: os.FileMode(420), modTime: time.Unix(1535420851, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaServerRegistryYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  server:
    description: TODO
    type: string
  username:
    description: TODO
    type: string
  password:
    description: TODO
    type: string
  password_file:
    description: TODO
    type: string
additionalProperties: false
`)

func schemaServerRegistryYamlBytes() ([]byte, error) {
	return _schemaServerRegistryYaml, nil
}

func schemaServerRegistryYaml() (*asset, error) {
	bytes, err := schemaServerRegistryYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/server-registry.yaml", size: 295, mode: os.FileMode(420), modTime: time.Unix(1535495163, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTaskYaml = []byte(`---

title: TODO
description: TODO
type: object
properties:
  type:
    description: TODO
    type: string
    enum:
      - remove
  extends:
    description: TODO
    type: string
  images:
    description: TODO
    oneOf:
      - type: string
      - type: array
        items:
          type: string
additionalProperties: false`)

func schemaTaskYamlBytes() ([]byte, error) {
	return _schemaTaskYaml, nil
}

func schemaTaskYaml() (*asset, error) {
	bytes, err := schemaTaskYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/task.yaml", size: 331, mode: os.FileMode(420), modTime: time.Unix(1535420757, 0)}
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
	"schema/build.yaml": schemaBuildYaml,
	"schema/config.yaml": schemaConfigYaml,
	"schema/ecr-registry.yaml": schemaEcrRegistryYaml,
	"schema/gcr-registry.yaml": schemaGcrRegistryYaml,
	"schema/logout.yaml": schemaLogoutYaml,
	"schema/metaplan.yaml": schemaMetaplanYaml,
	"schema/override.yaml": schemaOverrideYaml,
	"schema/plan.yaml": schemaPlanYaml,
	"schema/push.yaml": schemaPushYaml,
	"schema/run.yaml": schemaRunYaml,
	"schema/server-registry.yaml": schemaServerRegistryYaml,
	"schema/task.yaml": schemaTaskYaml,
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
	"schema": &bintree{nil, map[string]*bintree{
		"build.yaml": &bintree{schemaBuildYaml, map[string]*bintree{}},
		"config.yaml": &bintree{schemaConfigYaml, map[string]*bintree{}},
		"ecr-registry.yaml": &bintree{schemaEcrRegistryYaml, map[string]*bintree{}},
		"gcr-registry.yaml": &bintree{schemaGcrRegistryYaml, map[string]*bintree{}},
		"logout.yaml": &bintree{schemaLogoutYaml, map[string]*bintree{}},
		"metaplan.yaml": &bintree{schemaMetaplanYaml, map[string]*bintree{}},
		"override.yaml": &bintree{schemaOverrideYaml, map[string]*bintree{}},
		"plan.yaml": &bintree{schemaPlanYaml, map[string]*bintree{}},
		"push.yaml": &bintree{schemaPushYaml, map[string]*bintree{}},
		"run.yaml": &bintree{schemaRunYaml, map[string]*bintree{}},
		"server-registry.yaml": &bintree{schemaServerRegistryYaml, map[string]*bintree{}},
		"task.yaml": &bintree{schemaTaskYaml, map[string]*bintree{}},
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

