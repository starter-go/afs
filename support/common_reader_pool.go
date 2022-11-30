package support

import (
	"fmt"
	"io"
	"sync"

	"bitwormhole.com/starter/afs"
)

// todo ... 还没有完全实现！

type myCommonReaderPool struct {
	size   int
	table  map[string]*myCommonReaderPoolFile
	closed bool
	mu     sync.Mutex

	countOpen             int
	countClean            int
	countEntityWithNil    int
	countEntityWithoutNil int
}

func (inst *myCommonReaderPool) _Impl() afs.ReaderPool {
	return inst
}

func (inst *myCommonReaderPool) open(size int) {
	inst.table = make(map[string]*myCommonReaderPoolFile)
	inst.size = size
}

func (inst *myCommonReaderPool) getFile(file afs.Path, create bool) (*myCommonReaderPoolFile, error) {
	tab := inst.table
	if tab == nil {
		return nil, fmt.Errorf("this pool is closed")
	}
	key := file.GetPath()
	item := tab[key]
	if item == nil {
		if !create {
			return nil, fmt.Errorf("no cahced file: %v", key)
		}
		item = &myCommonReaderPoolFile{pool: inst}
		item.init(file)
		tab[key] = item
	}
	return item, nil
}

func (inst *myCommonReaderPool) Clean() {

	inst.mu.Lock()
	defer func() {
		inst.mu.Unlock()
	}()

	// scan all
	keepKeys := []string{}
	removeKeys := []string{}
	tab1 := inst.table
	for key, item := range tab1 {
		if item == nil {
			removeKeys = append(removeKeys, key)
		} else {
			item.clean()
			if item.isKeepAlive() {
				keepKeys = append(keepKeys, key)
			} else {
				removeKeys = append(removeKeys, key)
			}
		}
	}

	// make new table
	if len(removeKeys) > 0 {
		tab2 := make(map[string]*myCommonReaderPoolFile, 0)
		for _, key := range keepKeys {
			tab2[key] = tab1[key]
		}
		inst.table = tab2
	}
}

func (inst *myCommonReaderPool) Close() error {

	inst.mu.Lock()
	defer func() {
		inst.mu.Unlock()
	}()

	tab := inst.table
	inst.closed = true
	inst.table = nil
	if tab == nil {
		return nil
	}
	errlist := make([]error, 0)
	for _, f := range tab {
		if f == nil {
			continue
		}
		err := f.Close()
		if err != nil {
			errlist = append(errlist, err)
		}
	}
	if len(errlist) > 0 {
		return errlist[0]
	}
	return nil
}

func (inst *myCommonReaderPool) OpenReader(file afs.Path, op *afs.Options) (io.ReadSeekCloser, error) {
	inst.mu.Lock()
	defer func() {
		inst.mu.Unlock()
	}()
	if inst.closed {
		return nil, fmt.Errorf("this pool is closed")
	}
	f, err := inst.getFile(file, true)
	if err != nil {
		return nil, err
	}
	return f.open(op)
}

////////////////////////////////////////////////////////////////////////////////

type myCommonReaderPoolFile struct {
	pool       *myCommonReaderPool
	file       afs.Path
	items      []*myCommonReaderPoolEntity
	countOpen  int
	countClose int
	closed     bool
}

func (inst *myCommonReaderPoolFile) init(file afs.Path) {
	inst.file = file
}

func (inst *myCommonReaderPoolFile) clean() {
	// todo ...
	return
}

func (inst *myCommonReaderPoolFile) isKeepAlive() bool {
	// todo ...
	return false
}

func (inst *myCommonReaderPoolFile) countEntity() (withNil, withoutNil int) {

	withNil = 0
	withoutNil = 0

	list := inst.items
	if list == nil {
		return 0, 0
	}

	return withNil, withoutNil
}

func (inst *myCommonReaderPoolFile) Close() error {
	inst.closed = true
	inst.items = nil
	return nil
}

func (inst *myCommonReaderPoolFile) findReadyEntity() *myCommonReaderPoolEntity {
	list := inst.items
	for _, item := range list {
		if !item.busy {
			return item
		}
	}
	return nil
}

func (inst *myCommonReaderPoolFile) open(op *afs.Options) (io.ReadSeekCloser, error) {
	ent := inst.findReadyEntity()
	if ent == nil {
		ent = &myCommonReaderPoolEntity{}
		ent.init(inst)
		inst.items = append(inst.items, ent)
	}
	return ent.open(op)
}

////////////////////////////////////////////////////////////////////////////////

type myCommonReaderPoolEntity struct {
	file       *myCommonReaderPoolFile
	source     io.ReadSeekCloser
	busy       bool
	countOpen  int
	countClose int
	options    *afs.Options
}

func (inst *myCommonReaderPoolEntity) init(file *myCommonReaderPoolFile) error {
	inst.file = file
	return nil
}

func (inst *myCommonReaderPoolEntity) destroy() error {
	src := inst.source
	inst.source = nil
	if src == nil {
		return nil
	}
	return src.Close()
}

func (inst *myCommonReaderPoolEntity) open(op *afs.Options) (io.ReadSeekCloser, error) {
	if op != nil {
		inst.options = op
	}
	if inst.busy {
		return nil, fmt.Errorf("the file entity is busy")
	}
	se := &myCommonReaderPoolSession{}
	se.init(inst)
	err := se.open()
	if err != nil {
		return nil, err
	}
	return se.facade, nil
}

func (inst *myCommonReaderPoolEntity) onSessionOpen(session *myCommonReaderPoolSession) error {

	src := inst.source
	if src == nil {
		// open a new reader
		file := inst.file.file
		s2, err := file.GetIO().OpenSeekerR(inst.options)
		if err != nil {
			return err
		}
		src = s2
		inst.source = src
	}

	session.source = src

	inst.countOpen++
	inst.busy = true
	return nil
}

func (inst *myCommonReaderPoolEntity) onSessionClose(session *myCommonReaderPoolSession) error {
	session.source = nil
	inst.countClose++
	inst.busy = false
	if inst.isKeepAlive() {
		return nil
	}
	return inst.destroy()
}

func (inst *myCommonReaderPoolEntity) isKeepAlive() bool {
	c1 := inst.file.closed
	c2 := inst.file.pool.closed
	if c1 {
		return false
	}
	if c2 {
		return false
	}
	return true
}

////////////////////////////////////////////////////////////////////////////////

type myCommonReaderPoolSession struct {
	entity *myCommonReaderPoolEntity
	source io.ReadSeekCloser
	facade io.ReadSeekCloser
	opened bool
	closed bool
}

func (inst *myCommonReaderPoolSession) init(entity *myCommonReaderPoolEntity) {
	inst.entity = entity
	inst.facade = inst
}

func (inst *myCommonReaderPoolSession) open() error {
	if inst.opened {
		return nil
	}
	inst.opened = true
	return inst.entity.onSessionOpen(inst)
}

func (inst *myCommonReaderPoolSession) Read(b []byte) (int, error) {
	src := inst.source
	if src == nil {
		return 0, fmt.Errorf("this stream is closed")
	}
	return src.Read(b)
}

func (inst *myCommonReaderPoolSession) Seek(offset int64, ref int) (int64, error) {
	src := inst.source
	if src == nil {
		return 0, fmt.Errorf("this stream is closed")
	}
	return src.Seek(offset, ref)
}

func (inst *myCommonReaderPoolSession) Close() error {
	if inst.closed {
		return nil
	}
	inst.closed = true
	inst.source = nil
	return inst.entity.onSessionClose(inst)
}

////////////////////////////////////////////////////////////////////////////////
