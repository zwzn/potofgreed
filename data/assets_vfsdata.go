// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 6, 30, 15, 58, 19, 597999232, time.UTC),
		},
		"/common.go": &vfsgen۰CompressedFileInfo{
			name:             "common.go",
			modTime:          time.Date(2019, 7, 6, 14, 2, 19, 582513031, time.UTC),
			uncompressedSize: 1079,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xd1\x6e\x9b\x30\x14\x7d\xf6\xfd\x8a\x3b\xa4\x55\x20\xb5\xa0\x6d\x6f\x95\xfa\xd4\x64\x2c\x52\x36\xb5\x4a\xaa\xbd\x4c\x9a\x3c\x30\x60\x0d\x6c\x6a\x5f\xa6\x45\x11\xff\x3e\x19\x13\x12\x46\xb5\x8e\xa7\xe3\xe3\x73\x7c\xcf\xbd\xd8\x2d\xcf\x7e\xf2\x52\xe0\xf7\x11\x00\xc8\xa6\xd5\x86\x30\x04\x16\x14\x0d\x05\xc0\x02\x25\x28\xa9\x88\x5a\x87\xb5\x0d\x00\x58\x50\x4a\xaa\xba\x1f\x71\xa6\x9b\xa4\xd4\xba\xac\x45\xd2\x75\x32\x0f\xfe\xda\x31\xbc\xad\x6e\x4a\xdd\x56\xc2\x58\xbf\x7a\xae\x6f\x4a\xfd\x9f\xb2\xc4\x88\x9a\x1f\x02\x88\x00\xe8\xd0\x0a\xfc\xac\x73\x51\xa3\x25\xd3\x65\x84\x47\x60\x9b\x15\xba\xa2\xf1\xd3\xd3\x66\x05\xfd\x28\x7a\xee\x84\x39\x8c\xa2\x63\x0f\x90\x24\xb8\x23\x23\x55\x79\xaf\x9b\x96\x4b\x83\xd2\x62\x67\x45\x8e\xa4\xd1\x0a\x6e\xb2\x0a\x0b\x6d\x9c\x41\xaa\x12\xa5\xc2\xd4\xd5\x7f\xdc\xfa\xd3\xe6\xd6\x73\xe9\xf5\x23\x9e\x3e\xef\x04\xf6\x65\xbd\xa0\xb6\xfb\x25\xb5\x54\xa5\x4b\x55\xba\x54\xdd\x6b\x45\x5c\x2a\x7b\x41\xed\x88\x1b\xb2\x5f\x25\x55\x13\xb5\x56\xb9\x27\x26\x95\x1f\xc1\x46\xd1\xbf\xfb\x97\x8a\xec\xa2\xfb\x0b\xd7\xac\x75\xa9\xe8\xc3\xfb\xa1\xe3\x11\x6d\xf7\x13\x9a\xb8\x74\xe2\xd2\x13\xe7\xb3\x7c\xac\x35\x7f\x25\x4d\xe1\x24\xcb\x3c\x33\xe7\x2c\xd1\x60\x18\x33\x4d\x78\xbb\xbf\xc0\x17\x7c\x7a\xc1\xa7\x67\x7e\x9c\xd4\xea\x95\x41\xe5\x2f\xcc\x69\xf5\x62\xa8\xf3\xe5\x74\xb1\x66\x57\xb5\xe8\x54\x86\x9f\xb8\xca\x6b\x61\xc2\x08\xdd\xe3\x8a\xc7\xa5\x33\xdb\xac\x12\x0d\xbf\x46\x61\x0c\xde\xde\xe1\xf8\x24\xe2\x07\x6e\xac\xd8\x0d\x7b\xa1\xff\xbd\xa1\x57\x46\xd7\x78\x35\x5c\xfc\x63\x1f\x01\x93\xc5\x60\x7c\x73\x87\x4a\xd6\xee\x38\x56\x34\x14\x3f\x18\xa9\xa8\x08\x83\xb7\xbf\xbe\xa9\x60\x38\x3a\x02\xc6\xb4\x8d\xd7\xbf\x25\x85\xef\x22\x60\x3d\x00\x33\x82\x3a\xa3\xf0\x6a\x78\x7a\xa7\x48\x47\x5f\xf3\x16\x7d\xb5\x1e\x7a\xf8\x13\x00\x00\xff\xff\x2e\x32\x22\x54\x37\x04\x00\x00"),
		},
		"/model.go": &vfsgen۰CompressedFileInfo{
			name:             "model.go",
			modTime:          time.Date(2019, 7, 6, 13, 43, 7, 305720343, time.UTC),
			uncompressedSize: 779,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x51\xcd\x8a\xdb\x30\x10\x3e\x47\x4f\xf1\x91\x43\x49\x42\x49\xa0\x7d\x82\xb2\xa1\xb0\xd0\x5e\x5a\xef\xa9\x14\x31\x2b\x8f\x1d\x51\x5b\x4a\xa5\x11\xec\x12\xf6\xdd\x8b\x2c\x39\x6e\x53\xda\x3d\x79\x34\xf8\xfb\x9d\x33\x99\x1f\xd4\x33\x74\x1d\x94\xb2\xe3\xd9\x07\xc1\x46\xad\xd6\xc6\x3b\xe1\x27\x59\x2b\xb5\x5a\xf7\x56\x4e\xe9\x71\x6f\xfc\x78\xe8\xbd\xef\x07\x3e\xa4\x64\xdb\xb5\xda\x2a\x25\xcf\x67\x86\x6e\xf0\xd9\xb7\x3c\x28\x75\x38\x40\x77\xc9\x19\xb1\xde\xe9\x28\x14\x44\x97\x65\xf3\x21\xf4\x11\x14\x18\x72\x62\x50\xe8\xd3\xc8\x4e\x22\x3a\x1f\xa6\x8d\x6e\x60\xfc\x38\x92\x6b\x67\xca\x09\x10\x25\x24\x23\xb8\xa8\xd5\xfd\x11\x59\x74\xff\xf0\x70\x7f\x54\x2f\x95\x14\x9d\x75\x6d\x04\xe5\xd1\xba\x89\xa8\x25\xa1\x47\x8a\xac\xb2\x0d\x6c\x76\x3f\x13\x87\xe7\x2d\x74\xb3\x31\xf2\x84\x9a\x6a\x7f\x57\xbe\x6f\xb3\x93\x58\xc5\xb6\xd8\xe9\x26\x4b\x05\x96\x14\x1c\xde\xe8\xe6\xf2\x72\x95\xfa\x68\x07\xe1\x00\x9b\xd5\xba\x32\x8b\x47\x64\x0a\xe6\x34\x27\xb8\x6a\xd7\x08\x15\x53\x42\x5c\x66\xa6\xaf\x13\xe6\xd5\x3e\xca\x6f\xb7\xad\xfc\x06\x5e\xba\xf9\x64\x47\x2b\x80\x75\xf2\xfe\x9d\x5a\xdd\xf9\xe4\x04\xd8\xd5\x67\x35\xb1\x9b\xed\x5c\x13\x7d\xe1\x98\x06\xc9\x89\xb2\x60\x28\x2f\xdf\x81\x30\x75\x96\xc7\x7f\xe4\xaa\xc8\xc5\x40\x91\xac\x82\x47\x12\x02\xbe\x7d\xd7\x8d\xfa\x33\x72\x6d\x8b\xe3\x2d\x2f\xc8\xb5\x98\x5a\x1f\x22\x68\x18\x30\xb2\x98\x93\x75\x3d\x02\x1b\xe3\x43\x1b\xff\xba\x66\x61\xfc\xef\x4d\x97\xaa\xb6\x8b\xe7\xe5\xbc\xf3\x6a\xba\xf1\xaf\x00\x00\x00\xff\xff\x6c\x37\xa2\xe7\x0b\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/common.go"].(os.FileInfo),
		fs["/model.go"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
