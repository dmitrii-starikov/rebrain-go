package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"module07/internal/generated/wrappers"
	"module07/internal/generator"
	"module07/internal/monitors"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"
)

func Task03() {
	yamlConfiTemplate, err := ioutil.ReadFile("./assets/template/config_template.yml")
	if err != nil {
		panic(err)
	}

	_ = generator.ConfigGenerate(string(yamlConfiTemplate), "./config.yml")
}

func Task04() {
	marshallerTemplate, err := ioutil.ReadFile("./assets/template/marshaller.gotmpl")
	if err != nil {
		panic(err)
	}

	_ = generator.MarshallerGenerator(
		string(marshallerTemplate),
		"Config",
		"./internal/config/config.go",
		"./internal/config/codegen_marshaller.go",
	)
}

func Task05() {
	fmt.Println("=== Log & Metrics Wrappers ===")

	simple := monitors.NewSimpleMonitor()
	fmt.Println("Base monitor (no wrappers):")
	fmt.Println("   Type:", simple.Type())
	fmt.Println("   Run() returns:", simple.Run())

	logged := wrappers.NewMonitorWithLog(simple, os.Stdout, os.Stderr)
	fmt.Println("\nLog wrapper only:")
	fmt.Println("   Type:", logged.Type())
	fmt.Println("   Run() returns:", logged.Run())

	metrics := wrappers.NewMonitorWithPrometheus(logged, "simple-monitor")
	fmt.Println("\nBoth wrappers: Log + Metrics")
	fmt.Println("   Type:", metrics.Type())
	fmt.Println("   Run() returns:", metrics.Run())

	fmt.Println("\nLog and Metrics wrappers work correctly!")
}

func Task05Star() {
	fmt.Println("\n=== Testing recovery wrapper with panic monitor ===")

	panicMon := monitors.NewPanicMonitor()
	simple := monitors.NewSimpleMonitor()
	recovered := wrappers.NewMonitorWithRecovery(panicMon)
	recoveredAndSimple := wrappers.NewMonitorWithRecovery(simple)

	fmt.Println("Testing monitor that should panic...")
	fmt.Println("Type:", recovered.Type())

	err := recovered.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Success! Panic recovered.")
	}

	fmt.Println()

	recoveredAndSimpleErr := recoveredAndSimple.Run()
	if recoveredAndSimpleErr != nil {
		fmt.Printf("Error: %v\n", recoveredAndSimpleErr)
	} else {
		fmt.Println("Success! No panic, all good!")
	}
}

func startMetricsServer(wg *sync.WaitGroup) {
	defer wg.Done()
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Metrics server started on :9091")
	http.ListenAndServe(":9091", nil)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go startMetricsServer(&wg)
	time.Sleep(100 * time.Millisecond)

	Task03()
	Task04()
	Task05()
	Task05Star()

	fmt.Println("Hello, from 07 module")
	wg.Wait()
}
