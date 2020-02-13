// Code generated by go-bindata. DO NOT EDIT.
// sources:
// config/public-chain-accounts.json (307B)
// config/status-chain-accounts.json (543B)
// config/test-data.json (84B)

package t

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

var _configPublicChainAccountsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xce\xbf\x0a\xc2\x30\x10\x80\xf1\x3d\x4f\x71\x64\xee\x90\xbb\xf4\xf2\xa7\x5b\x28\xed\x2b\x38\xc7\x5c\x8a\x43\x51\xb0\x15\x04\xe9\xbb\x4b\x71\x71\x51\xc4\xf1\x5b\x3e\x7e\x0f\x05\xa0\x53\x29\x97\xdb\x79\x45\xdd\xc1\xde\x00\xfa\x90\xe7\xb9\xae\x49\xe4\x5a\x97\x45\x77\xa0\xcd\x7d\xb4\x3c\x18\x4b\x2c\x59\x82\xaf\xe4\x1c\x96\xb6\x8e\x91\x51\x38\x78\xf2\xd5\x09\x07\x2b\x5c\x74\xf3\x5a\xf4\xa7\xfc\xe7\x40\x01\x6c\xcd\x9b\x8b\xbe\xba\x92\xc9\x18\x89\x90\x5c\x90\x68\x63\xb1\x39\x7a\x3a\x4e\xdc\x3a\x94\x1e\xcd\xe4\x63\x30\x43\xc0\xf6\xa3\xeb\xe7\xc1\xee\x52\x9b\x7a\x06\x00\x00\xff\xff\x70\xe6\xdd\xe7\x33\x01\x00\x00")

func configPublicChainAccountsJsonBytes() ([]byte, error) {
	return bindataRead(
		_configPublicChainAccountsJson,
		"config/public-chain-accounts.json",
	)
}

func configPublicChainAccountsJson() (*asset, error) {
	bytes, err := configPublicChainAccountsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/public-chain-accounts.json", size: 307, mode: os.FileMode(0644), modTime: time.Unix(1574771268, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x76, 0x5d, 0xc0, 0xfe, 0x57, 0x50, 0x18, 0xec, 0x2d, 0x61, 0x1b, 0xa9, 0x81, 0x11, 0x5f, 0x77, 0xf7, 0xb6, 0x67, 0x82, 0x1, 0x40, 0x68, 0x9d, 0xc5, 0x41, 0xaf, 0xce, 0x43, 0x81, 0x92, 0x96}}
	return a, nil
}

var _configStatusChainAccountsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xce\xbf\x4e\xc3\x40\x0c\xc7\xf1\x3d\x4f\x71\xba\xb9\x83\x63\xfb\x9c\xbb\x6e\x97\x90\xcc\x6c\xcc\xbe\x7f\x62\xa8\x28\x6a\x8a\x40\x42\x7d\x77\x54\x21\x24\x96\x54\x54\x6c\xf6\xf0\xfb\xea\xf3\xd9\x19\x63\x63\xce\xc7\xb7\x97\x73\x6f\xf7\xe6\xfa\x1b\x63\x9f\xf4\x70\xa8\xe7\x58\xca\xa9\xae\xab\xdd\x1b\x0b\x1f\x69\xe9\x85\xb3\x12\xf7\x84\xa2\x40\xc9\xf1\x90\xc1\x8d\xc4\x94\x70\xc6\xbe\x2e\xb1\x22\xa7\x60\x77\xdf\x89\xe9\x59\xff\x17\x78\xd4\x75\x7d\x3f\x9e\xca\x75\xfd\xfa\x73\x77\xc6\x5c\x76\xbf\xcc\x78\xd3\x8c\xe0\xd4\x15\x98\x07\x6c\x4d\x1c\x0c\x61\x99\x74\x1c\x39\x0a\xcd\x44\x05\xbd\x46\x79\xa8\x0d\xa7\x4d\xf3\xbd\x81\x3f\x99\xe9\xa6\x99\xb4\x10\x57\x09\xbe\xb0\x07\xd1\x56\xc0\x27\x72\x21\x05\x84\xe6\xb2\x24\xc1\x24\xbe\x56\xde\x34\xdf\x1b\xd8\x34\x77\x97\xee\x2b\x00\x00\xff\xff\x84\x46\x00\x06\x1f\x02\x00\x00")

func configStatusChainAccountsJsonBytes() ([]byte, error) {
	return bindataRead(
		_configStatusChainAccountsJson,
		"config/status-chain-accounts.json",
	)
}

func configStatusChainAccountsJson() (*asset, error) {
	bytes, err := configStatusChainAccountsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/status-chain-accounts.json", size: 543, mode: os.FileMode(0644), modTime: time.Unix(1574771268, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8e, 0xb3, 0x61, 0x51, 0x70, 0x3c, 0x12, 0x3e, 0xf1, 0x1c, 0x81, 0xfb, 0x9a, 0x7c, 0xe3, 0x63, 0xd0, 0x8f, 0x12, 0xc5, 0x2d, 0xf4, 0xea, 0x27, 0x33, 0xef, 0xca, 0xf9, 0x3f, 0x72, 0x44, 0xbf}}
	return a, nil
}

var _configTestDataJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\xf2\xcb\x4f\x49\x55\xb2\x52\x00\xb1\x15\x14\x94\x82\x2b\xf3\x92\x83\x53\x93\xf3\xf3\x52\x8a\x95\xac\x14\x8c\x0d\x74\x20\xc2\x1e\x21\x21\x01\x01\xf9\x45\x25\x4a\x56\x0a\x16\x66\x26\xa6\x50\xd1\xf0\x60\x84\x98\x19\x97\x82\x42\x2d\x57\x2d\x17\x20\x00\x00\xff\xff\x51\xca\x96\xb1\x54\x00\x00\x00")

func configTestDataJsonBytes() ([]byte, error) {
	return bindataRead(
		_configTestDataJson,
		"config/test-data.json",
	)
}

func configTestDataJson() (*asset, error) {
	bytes, err := configTestDataJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/test-data.json", size: 84, mode: os.FileMode(0644), modTime: time.Unix(1574771268, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xce, 0x9d, 0x80, 0xf5, 0x87, 0xfa, 0x57, 0x1d, 0xa1, 0xd5, 0x7a, 0x10, 0x3, 0xac, 0xd7, 0xf4, 0x64, 0x32, 0x96, 0x2b, 0xb7, 0x21, 0xb7, 0xa6, 0x80, 0x40, 0xe9, 0x65, 0xe3, 0xd6, 0xbd, 0x40}}
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
	"config/public-chain-accounts.json": configPublicChainAccountsJson,
	"config/status-chain-accounts.json": configStatusChainAccountsJson,
	"config/test-data.json":             configTestDataJson,
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
	"config": &bintree{nil, map[string]*bintree{
		"public-chain-accounts.json": &bintree{configPublicChainAccountsJson, map[string]*bintree{}},
		"status-chain-accounts.json": &bintree{configStatusChainAccountsJson, map[string]*bintree{}},
		"test-data.json":             &bintree{configTestDataJson, map[string]*bintree{}},
	}},
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
