package framework

import (
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"sync"
)

var (
	errExist        = errors.New("That worker already exists")
	errTypeNotFound = errors.New("Invalid worker type")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

type MakeWorker func(WorkerType) (Work, error)
type Manager struct {
	MakeWorker
	Workers
	sync.WaitGroup
}

func NewManager(fn MakeWorker) *Manager {
	return &Manager{
		MakeWorker: fn,
		Workers:    make(Workers),
	}
}

func (m *Manager) initWorker(work *Work) {

}

func (m *Manager) GetWorkerCount() uint16 {
	var cnt uint16 = 0
	for id, worker := range m.Workers {
		fmt.Println(id, worker)
		cnt++
	}
	return cnt
}

func (m *Manager) GetWorkerCountByType(workerType uint8) uint16 {
	var cnt uint16 = 0
	for id, worker := range m.Workers {
		fmt.Println(id, worker)
		cnt++
	}
	return cnt
}

func (m *Manager) PrintWorker() {
	for _, worker := range m.Workers {
		fmt.Println(worker)
	}
}

func (m *Manager) BeginWorker(id string, wt WorkerType) error {
	_, error := m.Workers.search(id)
	if error == nil {
		return errExist
	}

	work, err := m.MakeWorker(wt)
	if err != nil {
		return err
	}

	m.Workers[id] = work
	work.SetId(id)
	ctx := context.WithValue(context.Background(), EVENT_CHAN_KEY, make(chan Event, MAX_EVENTS))

	expire := work.GetExpire()
	if 0 < expire {
		var cancelFn context.CancelFunc
		ctx, cancelFn = context.WithTimeout(ctx, expire)
		work.SetContext(ctx)
		work.SetCancelFunc(cancelFn)
	} else {
		work.SetContext(ctx)
	}

	m.WaitGroup.Add(1)
	go func() {
		defer m.WaitGroup.Done()
		work.DoWork()
	}()

	return nil
}

func (m *Manager) EndWorker(id string) error {
	worker, err := m.Workers.search(id)
	if err != nil {
		return err
	}

	if worker.GetCancelFunc() != nil {
		worker.GetCancelFunc()()
	}

	return m.Workers.delete(id)
}

func (m *Manager) EndMapOfWorker() {
	for _, worker := range m.Workers {
		if worker.GetCancelFunc() != nil {
			worker.GetCancelFunc()()
		}
	}

	m.WaitGroup.Wait()
	m.Workers.clear()
}

func (m *Manager) CancelWorker(id string) error {
	worker, err := m.Workers.search(id)
	if err != nil {
		return err
	}

	cancel := worker.GetCancelFunc()
	cancel()
	return nil
}

func (m *Manager) CancelMapOfWorker() {
	for _, worker := range m.Workers {
		log.Println("Cancel")
		worker.GetContext().Err()
	}
}

func (m *Manager) Emit(id string, event Event) error {
	worker, err := m.Workers.search(id)
	if err != nil {
		return err
	}

	v := worker.GetContext().Value(EVENT_CHAN_KEY)
	if v == nil {
		return errChanNotFound
	}

	ch, ok := v.(chan Event)
	if !ok {
		return errDiffChanType
	}

	ch <- event
	return nil
}
