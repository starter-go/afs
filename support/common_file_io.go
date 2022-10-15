package support

import (
	"errors"
	"io"
	"os"
	"strconv"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/vlog"
)

type myCommonFileIO struct {
	path afs.Path
}

func (inst *myCommonFileIO) _Impl() afs.FileIO {
	return inst
}

func (inst *myCommonFileIO) Path() afs.Path {
	return inst.path
}

func (inst *myCommonFileIO) OpenReader(op *afs.Options) (io.ReadCloser, error) {

	if op == nil {
		op = &afs.Options{}
	}

	path := inst.path.GetPath()
	flag := os.O_RDONLY
	return os.OpenFile(path, flag, op.Permission)
}

func (inst *myCommonFileIO) OpenWriter(op *afs.Options) (io.WriteCloser, error) {
	if op == nil {
		op = &afs.Options{}
	}
	file := inst.path
	path := file.GetPath()
	if op.Mkdirs {
		dir := file.GetParent()
		if !dir.Exists() {
			dir.Mkdirs(op)
		}
	}
	flag := os.O_CREATE | os.O_WRONLY
	return os.OpenFile(path, flag, op.Permission)
}

func (inst *myCommonFileIO) ReadBinary(op *afs.Options) ([]byte, error) {
	r, err := inst.OpenReader(op)
	if err != nil {
		return nil, err
	}
	defer func() {
		err2 := r.Close()
		inst.logError(err2)
	}()
	return io.ReadAll(r)
}

func (inst *myCommonFileIO) ReadText(op *afs.Options) (string, error) {
	data, err := inst.ReadBinary(op)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (inst *myCommonFileIO) WriteBinary(b []byte, op *afs.Options) error {

	if b == nil {
		return errors.New("data buffer is nil")
	}

	w, err := inst.OpenWriter(op)
	if err != nil {
		return err
	}
	defer func() {
		err2 := w.Close()
		inst.logError(err2)
	}()

	cnt, err := w.Write(b)
	if err != nil {
		return err
	}

	size := len(b)
	if cnt != size {
		want := strconv.Itoa(size)
		have := strconv.Itoa(cnt)
		return errors.New("bad io size, want:" + want + " have:" + have)
	}

	return nil
}

func (inst *myCommonFileIO) WriteText(text string, op *afs.Options) error {
	data := []byte(text)
	return inst.WriteBinary(data, op)
}

func (inst *myCommonFileIO) logError(err error) {
	if err == nil {
		return
	}
	vlog.Warn(err)
}
