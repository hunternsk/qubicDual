package farm

import "sync"

type WorkerType struct {
	Name string
	FsId int
}

type Workers struct {
	mx sync.RWMutex
	m  map[int]WorkerType
}

func NewWorkers() *Workers {
	return &Workers{
		m: make(map[int]WorkerType),
	}
}

func (w *Workers) Load(id int) (WorkerType, bool) {
	w.mx.RLock()
	defer w.mx.RUnlock()
	val, ok := w.m[id]
	return val, ok
}

func (w *Workers) Store(id int, val WorkerType) {
	w.mx.Lock()
	defer w.mx.Unlock()
	w.m[id] = val
}

func (w *Workers) SetFs(id, fsId int) {
	w.mx.Lock()
	defer w.mx.Unlock()
	_tmp := w.m[id]
	_tmp.FsId = fsId
	w.m[id] = _tmp
}

func (w *Workers) Len() int {
	w.mx.Lock()
	defer w.mx.Unlock()
	return len(w.m)
}

func (w *Workers) GetAll() map[int]WorkerType {
	w.mx.Lock()
	defer w.mx.Unlock()
	return w.m
}
