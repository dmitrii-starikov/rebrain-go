package monitors

import "fmt"

type panicMonitor struct {
	shouldPanic bool
}

func NewPanicMonitor() Monitor {
	return &panicMonitor{shouldPanic: true}
}

func (m *panicMonitor) Type() string {
	return "panic_monitor"
}

func (m *panicMonitor) Run() error {
	if m.shouldPanic {
		panic("something went wrong!")
	}
	fmt.Println("Running without panic")
	return nil
}
