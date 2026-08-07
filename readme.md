# Golang Basics - Learning Progress

Repository for learning Go through the **Golang Basics** course

## Project Structure

```
rebrain-go/
└── src/
    ├── Basics-02/           # Introduction
    │   ├── main.go
    │   └── hello.go
    ├── Basics-03/           # Language Fundamentals
    │   ├── 01_vars_n_types/
    │   ├── 02_block_pointers/
    │   ├── 03_slices/
    │   ├── etc.
    └── etc.
```

---

## Module Progress

### Module 02: Introduction. Environment Setup
- [x] Environment setup
- [x] Hello World

### Module 03: Language Fundamentals
- [x] **Block 01:** Variables. Data Types
  - `01_vars_n_types/` - [defineVars.go](src/Basics-03/01_vars_n_types/defineVars.go)
  - Tasks: [01_task](src/Basics-03/01_vars_n_types/01_task/main.go), [02_task](src/Basics-03/01_vars_n_types/02_task/main.go)
  - [Useful links](src/Basics-03/01_vars_n_types/links.md)

- [x] **Block 02:** Pointers in Go
  - `02_block_pointers/` - [main.go](src/Basics-03/02_block_pointers/main.go)
  - Tasks: [01_task](src/Basics-03/02_block_pointers/01_task.go), [02_task](src/Basics-03/02_block_pointers/02_task.go)

- [x] **Block 03:** Slices
  - `03_slices/` - [main.go](src/Basics-03/03_slices/main.go)
  - Task: [03_task](src/Basics-03/03_slices/03_task.go)

- [x] **Block 04:** Maps in Go
  - `04_maps/` - [main.go](src/Basics-03/04_maps/main.go)
  - Task: [04_task](src/Basics-03/04_maps/04_task.go)
  - [Useful links](src/Basics-03/04_maps/links.md)

- [x] **Block 05:** Language Constructs and Functions
  - `05_language_structures_and_functions/` - [main.go](src/Basics-03/05_language_structures_and_functions/main.go)
  - Task: [05_task.go](src/Basics-03/05_language_structures_and_functions/05_task.go)
  - [Useful links](src/Basics-03/05_language_structures_and_functions/links.md)

- [x] **Block 06:** defer - Function Exit Handling
  - `06_defer/` - [main.go](src/Basics-03/06_defer/main.go)
  - Task: [06_task.go](src/Basics-03/06_defer/06_task.go)
  - [Useful links](src/Basics-03/06_defer/links.md)

- [x] **Block 07:** Panic and Recovery
  - `07_panic/` - [main.go](src/Basics-03/07_panic/main.go)
  - Task: [07_task.go](src/Basics-03/07_panic/07_task.go)

- [x] **Block 08:** Error Handling
  - `08_errors/` - [main.go](src/Basics-03/08_errors/main.go)
  - Tasks: [08_task.go](src/Basics-03/08_errors/08_task.go), [homework.go](src/Basics-03/08_errors/homework.go)
  - [Useful links](src/Basics-03/08_errors/links.md)

### Module 04: Modules and Packages
- [x] **Block 01:** Scope, Initialization with init()
  - `01_scopes_and_initialization/` - [main.go](src/Basics-04/01_scopes_and_initialization/main.go)
  - Packages: `color/`, `color1/`, `color2/`, `colorInit/`, `wordz/`
  - Homework: [09_homework/](src/Basics-04/01_scopes_and_initialization/09_homework/)
  - [Useful links](src/Basics-04/01_scopes_and_initialization/links.md)

- [x] **Block 02:** Working with Dependencies, go mod
  - `02_go_mod/` - includes vendoring examples
  - Homework: [02_homework/](src/Basics-04/02_go_mod/02_homework/) with `homework_package/`

- [x] **Block 03:** Creating Modules and Versioning
  - `03_creating_modules_and_versioning/` - [task.go](src/Basics-04/03_creating_modules_and_versioning/task.go)
  - [Useful links](src/Basics-04/03_creating_modules_and_versioning/links.md)

- [x] **Block 04:** Project Layout (Structure)
  - `04_project/` - Standard Go project structure

### Module 05: Structs and Interfaces
- [x] **Block 01:** Structs in Go
  - `01_structs/` - [main.go](src/Basics-05/01_structs/cmd/myapp/main.go)
  - Internal: [config.go](src/Basics-05/01_structs/internal/config.go), [customer.go](src/Basics-05/01_structs/internal/customer.go)
  - Task: [01_task/](src/Basics-05/01_task/)
  - [Useful links](src/Basics-05/01_structs/links.md)

- [x] **Block 02:** Struct Methods
  - `02_methods/` - [main.go](src/Basics-05/02_methods/cmd/myapp/main.go)
  - Internal: [customer.go](src/Basics-05/02_methods/internal/customer.go)
  - Task: [02_task/](src/Basics-05/02_task/)
  - [Useful links](src/Basics-05/02_methods/links.md)

- [x] **Block 03:** Interfaces and Duck Typing
  - `03_interfaces_and_duck_typing/` - [main.go](src/Basics-05/03_interfaces_and_duck_typing/cmd/myapp/main.go)
  - Internal: [customer.go](src/Basics-05/03_interfaces_and_duck_typing/internal/customer.go), [partner.go](src/Basics-05/03_interfaces_and_duck_typing/internal/partner.go), [debtor.go](src/Basics-05/03_interfaces_and_duck_typing/internal/debtor.go)
  - Task: [03_task/](src/Basics-05/03_task/) with `discounter` interface
  - [Useful links](src/Basics-05/03_interfaces_and_duck_typing/links.md)

- [x] **Block 04:** Empty Interface
  - `04_empty_interface/` - [main.go](src/Basics-05/04_empty_interface/cmd/myapp/main.go)
  - Internal: [customer.go](src/Basics-05/04_empty_interface/internal/customer.go), [partner.go](src/Basics-05/04_empty_interface/internal/partner.go), [debtor.go](src/Basics-05/04_empty_interface/internal/debtor.go)
  - Task: [04_task/](src/Basics-05/04_task/)
  - [Useful links](src/Basics-05/04_empty_interface/links.md)

- [x] **Block 05:** Composite Inheritance
  - `05_inheritance/` - [main.go](src/Basics-05/05_inheritance/main.go)
  - Task: [05_task/](src/Basics-05/05_task/) with `overdue` type

- [x] **Block 06:** Advanced Error Handling
  - `06_advanced_error_handling/` - [main.go](src/Basics-05/06_advanced_error_handling/main.go)
  - Task: [06_task/homework.go](src/Basics-05/06_task/homework.go)
  - [Useful links](src/Basics-05/06_advanced_error_handling/links.md)

### Module 06: Concurrency
- [x] **Block 01:** Goroutines
  - `01_goroutines/` - [goroutines.go](src/Basics-06/01_goroutines/goroutines.go), [channel.go](src/Basics-06/01_goroutines/channel.go), [select.go](src/Basics-06/01_goroutines/select.go)
  - Examples: [sum.go](src/Basics-06/01_goroutines/sum.go), [fibonacci.go](src/Basics-06/01_goroutines/fibonacci.go), [webcrawler.go](src/Basics-06/01_goroutines/webcrawler.go), [mutex.go](src/Basics-06/01_goroutines/mutex.go), [numbers.go](src/Basics-06/01_goroutines/numbers.go)
  - Task: [01_task/goroutines.go](src/Basics-06/01_task/goroutines.go)
  - [Useful links](src/Basics-06/01_goroutines/links.md)

- [x] **Block 02:** Go Scheduler
  - `02_scheduler/` - [procs.go](src/Basics-06/02_scheduler/procs.go)
  - Task: [02_task/](src/Basics-06/02_task/) with `gosched/` and `procs/` examples, screenshots
  - [task.md](src/Basics-06/02_task/task.md)
  - [Useful links](src/Basics-06/02_scheduler/links.md)

- [x] **Block 03:** Race Conditions
  - `03_race_condition/` - [7dollars.go](src/Basics-06/03_race_condition/7dollars.go), [mutex7dollars.go](src/Basics-06/03_race_condition/mutex7dollars.go)
  - Task: [03_task/](src/Basics-06/03_task/) with `cache/` (unsafe and thread-safe implementations)
  - [README.md](src/Basics-06/03_race_condition/README.md)

- [x] **Block 04:** sync and atomic Packages
  - `04_sync_n_atomic/` - [rwmutex.go](src/Basics-06/04_sync_n_atomic/rwmutex.go)
  - Task: [04_task/](src/Basics-06/04_task/) with `cache/rw-safe.go`
  - [Useful links](src/Basics-06/04_sync_n_atomic/links.md)

- [x] **Block 05:** Channels pt.1. Deadlocks
  - `05_deadlocks/` - [channel.go](src/Basics-06/05_deadlocks/channel.go), [unbuf_pingpong.go](src/Basics-06/05_deadlocks/unbuf_pingpong.go), [semaphore.go](src/Basics-06/05_deadlocks/semaphore.go)
  - Task: [05_task/](src/Basics-06/05_task/) with `cache/rw-safe.go` and `semaphore/semaphore.go`
  - [Useful links](src/Basics-06/05_deadlocks/links.md)

- [x] **Block 06:** Channels pt.2. Context
  - `06_context/` - [context.go](src/Basics-06/06_context/context.go), [timeout.go](src/Basics-06/06_context/timeout.go), [goroutines.go](src/Basics-06/06_context/goroutines.go), [select.go](src/Basics-06/06_context/select.go)
  - Task: [06_task/](src/Basics-06/06_task/) with `cache/rw-safe.go`, `semaphore/semaphore.go`, screenshots and [result.md](src/Basics-06/06_task/result.md)
  - [Useful links](src/Basics-06/06_context/links.md)

- [x] **Block 07:** sync.Pool
  - `07_pool/` - [json.go](src/Basics-06/07_pool/json.go)
  - Task: [07_task/](src/Basics-06/07_task/) - [increment.go](src/Basics-06/07_task/increment.go), [increment_test.go](src/Basics-06/07_task/increment_test.go)
  - Screenshots: [without.png](src/Basics-06/07_task/without.png), [tests.png](src/Basics-06/07_task/tests.png)
  - [commands.md](src/Basics-06/07_task/commands.md), [Useful links](src/Basics-06/07_task/links.md) and [Useful links](src/Basics-06/07_pool/links.md)

- [x] **Block 08:** errgroup Package
  - `08_errgroup/` - [errgroup.go](src/Basics-06/08_errgroup/errgroup.go)
  - Task: [08_task/](src/Basics-06/08_task/) - [links.go](src/Basics-06/08_task/links.go), [readme.md](src/Basics-06/08_task/readme.md)

### ⏳ Module 07: Testing, Benchmarks and Profiling
- [x] **Block 01:** Unit Testing in Go
- [x] **Block 02:** Mocks, Stubs and Generation with GoMock
- [x] **Block 03:** Table Driven Tests vs Closure Driven Tests*
- [x] **Block 04:** Test Coverage
- [x] **Block 05:** Benchmarks
- [x] **Block 06:** Profiling with pprof

### ⏳ Module 08: Code Generation
- [x] **Block 01:** Reflection
- [x] **Block 02:** AST (Abstract Syntax Tree)
- [x] **Block 03:** Templates
- [ ] **Block 04:** Solving Reflection Issues
- [ ] **Block 05:** Wrapping

### ⏳ Module 09: Final Project
- [ ] **Block 01:** Final Project
- [ ] **Block 02:** Course Summary

