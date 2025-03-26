package main

import (
	"errors"
	"io"
	"io/fs"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func FileExists(file_name string) (fs.FileInfo, bool, error) {
	info, err := os.Stat(file_name)
	if err == nil {
		return info, true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return info, false, nil
	}
	return info, false, err
}

func CopyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}

	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
