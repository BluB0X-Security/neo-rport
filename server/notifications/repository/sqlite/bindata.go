// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 001_init.down.sql (29B)
// 001_init.up.sql (1.302kB)

package sqlite

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\x8e\xcf\xc9\x4f\xb7\x06\x04\x00\x00\xff\xff\x5d\x68\xf4\x86\x1d\x00\x00\x00")

func _001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initDownSql,
		"001_init.down.sql",
	)
}

func _001_initDownSql() (*asset, error) {
	bytes, err := _001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.down.sql", size: 29, mode: os.FileMode(0644), modTime: time.Unix(1688094241, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1e, 0xdc, 0x32, 0xb, 0xf2, 0x33, 0xb6, 0x50, 0x9e, 0x36, 0x9a, 0x12, 0x8c, 0xea, 0x4e, 0x29, 0x51, 0xab, 0x6d, 0x90, 0x1f, 0x8d, 0x6c, 0x7b, 0x4e, 0x7e, 0xf, 0x26, 0x3d, 0x48, 0x1a, 0x96}}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4f\x4f\xdb\x4e\x10\xbd\xfb\x53\xbc\x5f\x2e\x10\xc9\xf9\x09\x21\xb5\x17\xd4\x83\x1b\x5c\x81\x1a\x02\x0a\xa6\xe2\x66\x6d\xbc\x93\x64\x8b\xbd\xeb\xce\x8e\x81\xf4\xd3\x57\xbb\xf9\x43\x13\x02\xed\xa1\x3e\xce\xce\xbc\x37\xf3\x66\x9e\x87\x93\x3c\x2b\x72\x14\xd9\xe7\x51\x0e\xeb\xc4\xcc\x4c\xa5\xc4\x38\xeb\xcb\xda\xcd\x71\x9c\x00\xd8\x89\x97\x46\x63\x78\x91\x4d\x8e\x4f\x3f\xf6\x31\xbe\x2e\x30\xbe\x1b\x8d\x30\xbc\xc8\x87\x5f\x71\xbc\x9f\xf8\xdf\x27\x1c\x1d\xf5\x53\x0c\x06\xfb\x18\x29\x6c\xd7\x4c\x89\xe1\x66\x90\x05\x21\x90\x91\x15\x5e\xa6\xe8\x6a\xa3\x23\xad\x98\x86\xbc\xa8\xa6\xc5\x79\x56\xe4\xc5\xe5\x55\xfe\x42\x78\x9e\x7f\xc9\xee\x46\x05\x86\x77\x93\x49\x3e\x2e\xca\xf0\x7a\x5b\x64\x57\x37\x69\xa8\x1c\x0c\x5e\x8a\x53\x68\x25\x04\x65\x75\x8c\x05\xc2\x86\xbc\x57\x73\x82\x36\xbe\x55\x52\x2d\x8c\x9d\x47\xc2\xca\x59\x21\x2b\xc5\xb2\x25\x7c\xcb\x26\x71\xcc\x0f\x27\x7d\xbc\xa6\xed\xf5\xd2\x58\xc1\x34\x23\x26\x5b\x51\x98\x76\x53\x12\x94\x39\x58\xb2\xea\x8c\x69\xb6\x29\x49\xd1\xb2\x9b\xd6\xd4\xc0\x68\x2c\x94\xd5\xa4\xe1\x1e\x89\x31\x5d\x46\x55\x54\x2d\xc4\xc6\xce\xe1\x89\x1f\x4d\x45\xff\xaf\x74\x61\x65\x7d\xeb\x58\x50\xe4\xf7\xc5\x5b\x4c\x7f\xfb\x05\xad\x36\x80\x29\xc8\xc8\x82\x18\x3d\xdf\x48\xdb\x83\x63\x58\xb5\xd2\x2c\xb4\xe3\x2b\x36\xad\xac\x07\xaf\x4c\x6b\xc8\x8a\xff\x07\x4d\x44\x51\x36\x78\x69\x18\xd6\xa8\xda\xfc\x24\x0d\xc5\xac\x96\x49\x64\xf4\x12\xd6\xb8\xd5\xf8\xe4\xf5\xf5\xad\x32\x7e\xbb\x39\xff\xa3\x36\x42\xd0\x8e\xbc\x3d\x12\xd4\xe6\x81\x40\xb6\x6b\xfc\x1a\xb1\x9b\x7e\xa7\xea\x1d\x15\x63\xd6\xd4\xe9\xe5\x3b\x29\x31\xc7\x75\x6f\xa3\x84\x46\x9e\x16\x4a\x28\xec\x95\x98\x1d\xc3\x78\x30\x49\xc7\x36\xac\xdb\xd6\x4b\x3c\x2d\xc8\xbe\x7e\x4b\xfa\x67\xc9\x60\x90\x24\x6b\x8f\x5e\x8e\xcf\xf3\x7b\x18\xfd\x5c\xee\xfa\xd4\xe8\xe4\x7a\x7c\xc8\xbb\x7b\x96\xeb\x9f\xfd\x11\x6a\x6b\x9a\x38\xd5\x61\xd4\x6d\x4e\xc0\x0b\x22\x8b\x92\xce\x97\x95\xd3\x94\x22\x9c\xcd\x3a\x82\x10\x79\x99\x73\x7d\xd1\xb7\x57\xc5\x4d\xbc\xe6\x60\x7d\x8e\x21\x7a\x36\xb2\x4a\xde\x3d\xb3\x80\xdd\x48\x5b\x46\x5d\x7c\xba\xd6\x67\xe3\xdd\x2d\xf0\x8c\x5d\xb3\x0f\x1d\x6b\x23\x4a\xc9\xe4\x5b\x67\x3d\xa5\x98\x19\xf6\x82\xd3\x87\x40\xe3\x45\x87\x9d\x85\x9f\x82\x17\x4d\xcc\xbb\xdc\xbf\x02\x00\x00\xff\xff\x7b\x63\x51\x66\x16\x05\x00\x00")

func _001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initUpSql,
		"001_init.up.sql",
	)
}

func _001_initUpSql() (*asset, error) {
	bytes, err := _001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.up.sql", size: 1302, mode: os.FileMode(0644), modTime: time.Unix(1688526544, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2b, 0xe3, 0x53, 0x78, 0x94, 0x5e, 0x40, 0x31, 0xab, 0xe7, 0x91, 0x66, 0x0, 0x1c, 0xb, 0x6b, 0x93, 0x4c, 0xee, 0x3f, 0xc0, 0x31, 0xf9, 0xf2, 0x14, 0xea, 0x49, 0x29, 0x43, 0xb2, 0x14, 0x2c}}
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
	"001_init.down.sql": _001_initDownSql,
	"001_init.up.sql":   _001_initUpSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"001_init.down.sql": {_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql":   {_001_initUpSql, map[string]*bintree{}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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
