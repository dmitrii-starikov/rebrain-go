Reading and writing to a channel are blocking. This means the following:

* When a goroutine is about to write to or read from a channel, it blocks and checks 
if the action is possible.
* **(Read, when other writes)** If the channel is **unbuffered**, writing to it is only possible if another goroutine is reading from it (and if it hasn't already been closed).
* **(Write, when other reads)** If the channel is **unbuffered**, reading from it is only possible if another goroutine is writing to it (or if it has already been closed).
* Therefore, it's incorrect to think of a channel as a "message pipe," although this 
analogy is often used to explain it.
* If the channel is **buffered**, writing to it is only possible if `len(ch) < cap(ch)` (the number of elements in it is less than the maximum, the channel is not full).
* If the channel is **buffered**, reading from it is only possible if it contains **at least one element**.

---

The source file `go/src/runtime/chan.go` describes the structure of a channel. A channel consists of:
* a counter of current elements;
* a value for the channel's capacity;
* A pointer to an array of elements;
* The element's size;
* The "channel is closed" value;
* A pointer to the element's type;
* The send index;
* The receive index;
* A list of goroutines reading the channel;
* A list of goroutines writing to the channel;
* A mutex.

```go
type hchan struct {
      	qcount   uint           // total data in the queue	
      	dataqsiz uint           // size of the circular queue	
      	buf      unsafe.Pointer // points to an array of dataqsiz elements	
      	elemsize uint16	
      	closed   uint32	
      	elemtype *_type         // element type	
      	sendx    uint           // send index	
      	recvx    uint           // receive index	
      	recvq    waitq          // list of recv waiters	
      	sendq    waitq          // list of send waiters	
      	lock     mutex
}
```
## Dining philosophers problem

A classic example of a deadlock is the "Dining philosophers problem."

It's stated like this: 
_five philosophers are seated at a round table, each with a plate of spaghetti in front of 
them. Each has a fork to the left of their plate. Thus, there are five plates and five 
forks on the table. Each philosopher can either eat or think. To eat, two forks are 
needed. Taking away forks is prohibited._

The goal is to ensure that each philosopher can both think and eat (that is, control of 
resources must be properly distributed). In the Dining philosophers problem, a deadlock 
occurs when each philosopher takes one fork (say, the left one) and waits for the second 
(the right one). However, since everyone has already grabbed one, no one can get a second 
one—and everyone is forever stuck in a waiting state. Formally, four conditions must be 
met for a deadlock to occur:

* **Mutual exclusion:** a fork is an exclusive resource; only one philosopher can hold it at a time.
* **Holding and waiting:** a philosopher who takes the left fork holds it and requests the right one without 
releasing the first one.
* **No forced preemption:** a philosopher cannot be "reflexively" held by the system—until they release the fork themselves, it will not be taken from them.
* **Circular waiting:** a ring is formed: "Philosopher 1 is waiting for fork y2, 2—y3, …, 5—y1."

### Ways to avoid deadlock
* Limit the number of concurrent executions. Introduce an additional semaphore 
`table := newSemaphore(4)`. Before taking forks, the philosopher attempts `table.Acquire()`, 
and after finishing eating, he calls `table.Release()`. Since there are 5 philosophers and 
4 allowed "at the table," at least one will always be able to take both forks and eat, then 
release resources, breaking the waiting loop.
* Alternate the order of acquisition. For example, even philosophers take the left fork 
first, then the right, while odd philosophers take the opposite order. This breaks the 
waiting loop.
* Host. A table "host" is introduced, issuing eating permissions, guaranteeing that no 
more than two philosophers are at the table at a time (or some other safe limit).
* Failure to hold. If a philosopher fails to take a second fork during the waiting 
period, they replace the first one and try again after a random interval. 
This reduces the risk of deadlock, replacing deadlock with potential starvation, 
which is easier to monitor and prevent.

One possible solution to the "Dining philosophers problem" is semaphores.
In general, **a mutex**, behaviorally, could be considered **a special case of a semaphore**:

* A mutex guarantees access to only one thread/goroutine. It is initialized as "unlocked,"
and any attempt to lock it when it is already locked will result in a wait.
* A semaphore can allow N simultaneous acquisitions. It is initialized to N (the capacity), 
and each acquire decrements the available counter by 1, while at 0, subsequent acquisitions 
are blocked.

## Links
- [Edsger W. Dijkstra](https://ru.wikipedia.org/wiki/%D0%94%D0%B5%D0%B9%D0%BA%D1%81%D1%82%D1%80%D0%B0,_%D0%AD%D0%B4%D1%81%D0%B3%D0%B5%D1%80_%D0%92%D0%B8%D0%B1%D0%B5)
- [Dining philosophers problem](https://ru.wikipedia.org/wiki/%D0%97%D0%B0%D0%B4%D0%B0%D1%87%D0%B0_%D0%BE%D0%B1_%D0%BE%D0%B1%D0%B5%D0%B4%D0%B0%D1%8E%D1%89%D0%B8%D1%85_%D1%84%D0%B8%D0%BB%D0%BE%D1%81%D0%BE%D1%84%D0%B0%D1%85)
- [Semaphore](https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D0%BC%D0%B0%D1%84%D0%BE%D1%80_(%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%BC%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5))

# Buffered vs Unbuffered Channels
| Aspect	| Unbuffered Channel                                                               |	Buffered Channel |
|---|----------------------------------------------------------------------------------|---|
| Declaration  | make(chan T) or make(chan T, 0)                                                  | make(chan T, N) where N > 0 |
|Capacity| 0 (no storage)                                                                   |N elements can be stored|	
|Send behavior| Blocks until receiver is ready                                                   |Blocks only when buffer is full|	
|Receive behavior| Blocks until sender is ready                                                     |Blocks only when buffer is empty|
|Synchronization| Provides guaranteed synchronization ("handshake")                                |No synchronization guarantee|
|Use cases| • Signaling between goroutines • Guaranteed handoff • Synchronization primitives |• Task queues • Work pools • Decoupling sender/receiver • Rate limiting|
|Analogy| Hand-to-hand delivery (both parties must be present)                       |Mailbox with capacity (sender drops, receiver picks up later)|

## Key Takeaways
- **Unbuffered channels** ensure that send and receive happen simultaneously - perfect for synchronization when you need guarantees.
- **Buffered channels** decouple the sender and receiver, allowing them to operate at different speeds - ideal for building pipelines and queues.

Choose based on your needs: use unbuffered for coordination, buffered for decoupling!

An **unbuffered channel** is like a walkie-talkie:
* You can only speak (and want to be listened) when you press a button and someone is listening.
* Synchronization is needed: "You to me, I to you."
* Can be reused.

A **buffered channel** is like email:
* Send and forget.
* The recipient will read it when it's convenient.
* There's a buffer of a certain size.