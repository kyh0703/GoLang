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
	receive := make(<-chan Event, 10)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "recv_chan", receive)
	ctx = context.WithValue(ctx, "test", 1)

	expire := work.GetExpire()
	if 0 < expire {
		var cancelFn context.CancelFunc
		ctx, cancelFn = context.WithTimeout(ctx, expire)
		work.SetContext(ctx)
		work.SetCancelFunc(cancelFn)
	} else {
		work.SetContext(ctx)
	}

	fmt.Println(work.GetContext().Value("test"))

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

func (m *Manager) EndWorkers() {
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

func (m *Manager) CancelWorkers() {
	for _, worker := range m.Workers {
		log.Println("Cancel")
		worker.GetContext().Err()
	}
}

// func (m *Manager) Emit(id string, event struct{}) error {
// 	worker, err := m.Workers.search(id)
// 	if err != nil {
// 		return err
// 	}
// }
