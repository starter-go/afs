package support

import (
	"errors"
	"io"
	"io/fs"
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

func (inst *myCommonFileIO) openR(op *afs.Options) (*os.File, error) {
	op = inst.prepareOptionsForRead(op)
	path := inst.path.GetPath()
	return os.OpenFile(path, op.Flag, op.Permission)
}

func (inst *myCommonFileIO) openW(op *afs.Options) (*os.File, error) {
	op = inst.prepareOptionsForWrite(op)
	file := inst.path
	path := file.GetPath()
	if op.Mkdirs {
		dir := file.GetParent()
		if !dir.Exists() {
			dir.Mkdirs(op)
		}
	}
	if op.Create {
		op.Flag |= os.O_CREATE
	}
	return os.OpenFile(path, op.Flag, op.Permission)
}

func (inst *myCommonFileIO) OpenReader(op *afs.Options) (io.ReadCloser, error) {
	return inst.openR(op)
}

func (inst *myCommonFileIO) OpenWriter(op *afs.Options) (io.WriteCloser, error) {
	return inst.openW(op)
}

func (inst *myCommonFileIO) OpenSeekerR(op *afs.Options) (io.ReadSeekCloser, error) {
	return inst.openR(op)
}

func (inst *myCommonFileIO) OpenSeekerW(op *afs.Options) (afs.WriteSeekCloser, error) {
	return inst.openW(op)
}

func (inst *myCommonFileIO) OpenSeekerRW(op *afs.Options) (afs.ReadWriteSeekCloser, error) {
	return inst.openW(op)
}

func (inst *myCommonFileIO) ReadBinary(op *afs.Options) ([]byte, error) {

	op = inst.prepareOptionsForRead(op)

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

	op = inst.prepareOptionsForRead(op)

	data, err := inst.ReadBinary(op)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (inst *myCommonFileIO) WriteBinary(b []byte, op *afs.Options) error {

	op = inst.prepareOptionsForWrite(op)

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

	op = inst.prepareOptionsForWrite(op)

	data := []byte(text)
	return inst.WriteBinary(data, op)
}

func (inst *myCommonFileIO) logError(err error) {
	if err == nil {
		return
	}
	vlog.Warn(err)
}

func (inst *myCommonFileIO) prepareOptionsForWrite(ops *afs.Options) *afs.Options {
	if ops == nil {
		ops = &afs.Options{}
	}
	if ops.Permission == 0 {
		ops.Permission = fs.ModePerm
	}
	if ops.Flag == 0 {
		ops.Flag = os.O_TRUNC | os.O_WRONLY
	}
	return ops
}

func (inst *myCommonFileIO) prepareOptionsForRead(ops *afs.Options) *afs.Options {
	if ops == nil {
		ops = &afs.Options{}
	}
	if ops.Flag == 0 {
		ops.Flag = os.O_RDONLY
	}
	return ops
}
