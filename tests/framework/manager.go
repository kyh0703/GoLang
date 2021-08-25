package framework

import "fmt"

// type TransFn func(job *Job, event Event)

// type JobType struct {
// 	fn TransFn
// }

// func NewManager() *Manager {
// 	w := &Manager{
// 		transObjs: make(map[string]*JobType),
// 		transJobs: make(map[string]*Job),
// 	}

// 	return w
// }

// func (m *Manager) AddJobType(tName string, fn TransFn) *JobType {
// 	trans := &JobType{fn: fn}
// 	m.transObjs[tName] = trans
// 	return trans
// }

// /**
//  * 신규 Job 생성
//  */
// func (m *Manager) RunJob(tName string, invokeId string, dur time.Duration, event Event) bool {
// 	trans, ok := m.transObjs[tName]
// 	if !ok {
// 		return false
// 	}
// 	_, ok = m.transJobs[invokeId]
// 	if ok {
// 		return false
// 	}

// 	job := &Job{
// 		eventCh:   make(chan Event, 100),
// 		keepEvent: make([]Event, 0),
// 	}

// 	m.transJobs[invokeId] = job

// 	job.ctx = context.Background()
// 	if dur > 0 {
// 		job.ctx, job.cancelFn = context.WithTimeout(job.ctx, dur)
// 	}

// 	m.wg.Add(1)

// 	go func() {
// 		defer m.wg.Done()
// 		if job, ok := m.transJobs[invokeId]; ok {
// 			trans.fn(job, event)
// 			fmt.Printf("Job Done\n")
// 		}
// 	}()

// 	return true
// }

// /**
//  * Job에 Event 등록
//  */
// func (m *Manager) Emmit(invokeId string, event Event) bool {
// 	job, ok := m.transJobs[invokeId]
// 	if !ok {
// 		return false
// 	}

// 	job.Emmit(event)
// 	return true
// }

// /**
//  * Job 취소
//  * @param  {[type]} m *Manager) EndJob(invokeId string [description]
//  * @return {[type]}   [description]
//  */
// func (m *Manager) EndJob(invokeId string) {
// 	job, ok := m.transJobs[invokeId]
// 	if !ok {
// 		return
// 	}

// 	if job.cancelFn != nil {
// 		job.cancelFn()
// 	}

// 	delete(m.transJobs, invokeId)
// }

////////////////////////////////////////
type MakeWorker func(WorkerType) *Worker
type Manager struct {
	MakeWorker
	Workers
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

func (m *Manager) BeginWorker(id string, workerType uint8) error {
	return nil
}

func (m *Manager) EndWorker() {
}
