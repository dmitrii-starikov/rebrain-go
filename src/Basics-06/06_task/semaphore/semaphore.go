package semaphore

import "sync/atomic"

type Semaphore struct {
	// chan struct{} better than int in src/Basics-06/05_task/semaphore/semaphore.go
	slots chan struct{}
	tasks int32
}

func NewSemaphore(slots int) *Semaphore {
	c := make(chan struct{}, slots)
	return &Semaphore{
		slots: c,
		tasks: 0,
	}
}

func (m *Semaphore) P() {
	m.slots <- struct{}{}
}

func (m *Semaphore) V() {
	<-m.slots
	atomic.AddInt32(&m.tasks, -1)
}

func (m *Semaphore) NewTask() {
	atomic.AddInt32(&m.tasks, 1)
}

func (m *Semaphore) IsWorking() bool {
	return atomic.LoadInt32(&m.tasks) > 0 || len(m.slots) > 0
}
