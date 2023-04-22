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
	path    afs.Path
	context *myFSContext
}

func (inst *myCommonFileIO) _Impl() afs.FileIO {
	return inst
}

func (inst *myCommonFileIO) Path() afs.Path {
	return inst.path
}

func (inst *myCommonFileIO) openR(op *afs.Options) (*os.File, error) {

	op = inst.prepareOptions(op, &afs.Options{
		Create:    false,
		Directory: false,
		File:      true,
		Read:      true,
		Write:     false,
	})

	path := inst.path.GetPath()
	return os.OpenFile(path, op.Flag, op.Permission)
}

func (inst *myCommonFileIO) openW(op *afs.Options) (*os.File, error) {

	op = inst.prepareOptions(op, &afs.Options{
		Create:    true,
		Directory: false,
		File:      true,
		Read:      false,
		Write:     true,
	})

	file := inst.path
	if op.Mkdirs {
		dir := file.GetParent().GetPath()
		os.MkdirAll(dir, op.Permission)
	}
	path := file.GetPath()
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

	op = inst.prepareOptions(op, &afs.Options{
		Create:    false,
		Directory: false,
		File:      true,
		Read:      true,
		Write:     false,
	})

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

	op = inst.prepareOptions(op, &afs.Options{
		Create:    false,
		Directory: false,
		File:      true,
		Read:      true,
		Write:     false,
	})

	data, err := inst.ReadBinary(op)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (inst *myCommonFileIO) WriteBinary(b []byte, op *afs.Options) error {

	op = inst.prepareOptions(op, &afs.Options{
		Create:    true,
		Directory: false,
		File:      true,
		Read:      false,
		Write:     true,
	})

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

	op = inst.prepareOptions(op, &afs.Options{
		Create:    true,
		Directory: false,
		File:      true,
		Read:      false,
		Write:     true,
	})

	data := []byte(text)
	return inst.WriteBinary(data, op)
}

func (inst *myCommonFileIO) logError(err error) {
	if err == nil {
		return
	}
	vlog.Warn(err)
}

func (inst *myCommonFileIO) prepareOptions(have, want *afs.Options) *afs.Options {
	return inst.context.common.PrepareOptions(inst.path, have, want)
}
