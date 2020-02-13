// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1557732988_initialize_db.down.sql (72B)
// 1557732988_initialize_db.up.sql (234B)
// static.go (178B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __1557732988_initialize_dbDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xc8\x4c\x89\x4f\xca\xc9\xcf\xcf\x8d\xcf\x4c\xa9\xb0\xe6\x42\x95\x28\xc9\x2f\xc8\x4c\x46\x92\x08\x71\x74\xf2\x71\x55\x48\xcd\x2b\x4b\xcd\xc9\x2f\x48\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\x6b\x93\xaa\x08\x48\x00\x00\x00")

func _1557732988_initialize_dbDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1557732988_initialize_dbDownSql,
		"1557732988_initialize_db.down.sql",
	)
}

func _1557732988_initialize_dbDownSql() (*asset, error) {
	bytes, err := _1557732988_initialize_dbDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1557732988_initialize_db.down.sql", size: 72, mode: os.FileMode(0644), modTime: time.Unix(1574771268, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x77, 0x40, 0x78, 0xb7, 0x71, 0x3c, 0x20, 0x3b, 0xc9, 0xb, 0x2f, 0x49, 0xe4, 0xff, 0x1c, 0x84, 0x54, 0xa1, 0x30, 0xe3, 0x90, 0xf8, 0x73, 0xda, 0xb0, 0x2a, 0xea, 0x8e, 0xf1, 0x82, 0xe7, 0xd2}}
	return a, nil
}

var __1557732988_initialize_dbUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x48\xcd\x2b\x4b\xcd\xc9\x2f\x48\x2d\x56\xd0\xc8\x4c\x51\x70\x8a\x0c\x71\x75\x54\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x08\xf5\xf3\x0c\x0c\x75\xd5\x51\x48\x49\x2c\x49\x44\x93\xd3\x51\x28\xc9\x2f\xc8\x4c\xc6\x10\x4d\xca\xc9\xcf\xcf\x55\x70\xf2\x0c\xd1\x30\x35\x34\xd2\x84\x4b\x68\x5a\x73\x71\x41\xed\xf5\xf4\x73\x71\x8d\x50\xc8\x4c\x89\x07\x2b\x8d\xcf\x4c\xa9\x50\xf0\xf7\x43\x73\x87\x8b\x6b\xb0\x33\xd4\x2c\x4d\x6b\x0c\x8d\x60\x9b\xf1\x69\x04\x2b\xd0\xb4\xe6\x02\x04\x00\x00\xff\xff\xe2\x06\x6b\x16\xea\x00\x00\x00")

func _1557732988_initialize_dbUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1557732988_initialize_dbUpSql,
		"1557732988_initialize_db.up.sql",
	)
}

func _1557732988_initialize_dbUpSql() (*asset, error) {
	bytes, err := _1557732988_initialize_dbUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1557732988_initialize_db.up.sql", size: 234, mode: os.FileMode(0644), modTime: time.Unix(1574771268, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8f, 0xa, 0x31, 0xf, 0x94, 0xe, 0xd7, 0xd6, 0xaa, 0x22, 0xd6, 0x6c, 0x7a, 0xbc, 0xad, 0x6a, 0xed, 0x2e, 0x7a, 0xf0, 0x24, 0x81, 0x87, 0x14, 0xe, 0x1c, 0x8a, 0xf1, 0x45, 0xaf, 0x9e, 0x85}}
	return a, nil
}

var _staticGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8c\x41\x6a\xc3\x40\x0c\x45\xf7\x73\x8a\xbf\x6c\xa1\x1e\xed\x7b\x82\x52\x12\x08\x24\x17\x90\x6d\x21\x0b\xc7\x33\x46\x52\x72\xfe\x6c\x12\x42\x96\x8f\xc7\x7b\x44\x38\xf1\xb4\xb2\x0a\x22\x39\x6d\x82\x6c\xa3\xcc\xf1\xa2\xaf\xff\xf3\x0f\xfe\x2e\xc7\xc3\x37\x5c\xa2\xdf\x7c\x92\x80\x9b\x2e\x09\x6b\xd9\x91\x8b\x60\xb4\xc6\x6e\x12\x65\xff\x38\x95\x42\xa4\xfd\x57\xa5\x89\x73\x0a\xb4\x0f\xa3\xb5\x99\x93\x31\xec\xab\x62\x33\x75\x4e\xeb\x2d\x30\x74\xd4\x4a\xb5\xd2\xc6\x76\x0d\xf1\xbb\x38\xbd\x35\x3d\xb3\xaa\x1d\xb5\x3c\x02\x00\x00\xff\xff\xf4\xe4\x35\xe2\xb2\x00\x00\x00")

func staticGoBytes() ([]byte, error) {
	return bindataRead(
		_staticGo,
		"static.go",
	)
}

func staticGo() (*asset, error) {
	bytes, err := staticGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static.go", size: 178, mode: os.FileMode(0644), modTime: time.Unix(1574771268, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xab, 0x8a, 0xf4, 0x27, 0x24, 0x9d, 0x2a, 0x1, 0x7b, 0x54, 0xea, 0xae, 0x4a, 0x35, 0x40, 0x92, 0xb5, 0xf9, 0xb3, 0x54, 0x3e, 0x3a, 0x1a, 0x2b, 0xae, 0xfb, 0x9e, 0x82, 0xeb, 0x4c, 0xf, 0x6}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"1557732988_initialize_db.down.sql": _1557732988_initialize_dbDownSql,
	"1557732988_initialize_db.up.sql":   _1557732988_initialize_dbUpSql,
	"static.go":                         staticGo,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"1557732988_initialize_db.down.sql": &bintree{_1557732988_initialize_dbDownSql, map[string]*bintree{}},
	"1557732988_initialize_db.up.sql":   &bintree{_1557732988_initialize_dbUpSql, map[string]*bintree{}},
	"static.go":                         &bintree{staticGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
