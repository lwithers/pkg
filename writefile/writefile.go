/*
Package writefile provides simple support routines for writing data to a temporary
file before renaming it into place. This avoids pitfalls such as writing partial
content to a file and then being interruted, or trying to write to a program that
is currently being executed, etc.

This package will correctly dereference symlinks and in the case of overwriting
will retain permissions from the original, underlying file.
*/
package writefile

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

const (
	// MaxSymlinkDeref is the maximum number of symlinks that we will
	// dereference before giving up.
	MaxSymlinkDeref = 16
)

// New opens a file for writing. It returns the final filename which should be
// passed to Commit; this may differ from the passed target filename if the
// target is actually a symlink.
//
// Any existing file will not be altered in any way until Commit() is called.
// Data will be written to a temporary file in the same directory as the
// target.
func New(targetFname string) (finalFname string, f *os.File, err error) {
	var (
		info os.FileInfo
		tgt  string
	)
	finalFname = targetFname
	for i := 0; i < MaxSymlinkDeref; i++ {
		info, err = os.Lstat(finalFname)
		switch {
		case err != nil && !os.IsNotExist(err):
			return

		case os.IsNotExist(err), info.Mode().IsRegular():
			f, err = NewNoDeref(finalFname)
			return

		case info.Mode()&os.ModeType == os.ModeSymlink:
			tgt, err = os.Readlink(finalFname)
			if err != nil {
				return
			}
			if filepath.IsAbs(tgt) {
				finalFname = tgt
			} else {
				finalFname = filepath.Clean(filepath.Join(filepath.Dir(finalFname), tgt))
			}

		default:
			err = &os.PathError{
				Op:   "open",
				Path: targetFname,
				Err:  errors.New("not a regular file"),
			}
			return
		}
	}
	err = &os.PathError{
		Op:   "open",
		Path: targetFname,
		Err:  unix.ELOOP,
	}
	return
}

// Abort writing to a file. Performs a Close() and unlinks the temporary file.
func Abort(f *os.File) {
	_ = f.Close()
	_ = os.Remove(f.Name())
}

// NewNoDeref is similar to New, but will not dereference symlinks and will
// allow them to be overwritten.
func NewNoDeref(finalFname string) (*os.File, error) {
	return ioutil.TempFile(filepath.Dir(finalFname), ".new.")
}

// Commit the temporary file. This will close the file f and then rename it
// into place once it has been ensured the data is on disk. It will retain
// permissions from the original file if present.
func Commit(finalFname string, f *os.File) error {
	if err := f.Sync(); err != nil {
		os.Remove(f.Name())
		return err
	}

	// if the final destination file already exists, try to inherit its
	// permissions, but don't return an error if we fail
	if st, err := os.Stat(finalFname); err == nil { // NB: inverted
		f.Chmod(st.Mode() & os.ModePerm)
	}

	if err := f.Close(); err != nil {
		os.Remove(f.Name())
		return err
	}

	if err := os.Rename(f.Name(), finalFname); err != nil {
		os.Remove(f.Name())
		return err
	}

	return nil
}
