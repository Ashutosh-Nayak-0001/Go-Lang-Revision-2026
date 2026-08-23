package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=====================================================================
GO GOROUTINES — 70 QUESTIONS & ANSWERS
=====================================================================

Run:
    go run goroutines_70.go

LEVEL:
    Q01-Q20  Beginner
    Q21-Q45  Intermediate
    Q46-Q70  Advanced / Interview

IMPORTANT:
Many examples intentionally demonstrate concepts rather than
maximum production performance.

=====================================================================
Q01. What is a goroutine?
=====================================================================
ANSWER:

A goroutine is a lightweight concurrent function execution managed
by the Go runtime.

Syntax:

    go function()

Example:
*/

func q01() {
	fmt.Println("\nQ01 - Basic goroutine")

	done := make(chan bool)

	go func() {
		fmt.Println("Hello from goroutine")
		done <- true
	}()

	<-done
}

/*
=====================================================================
Q02. How do you create a goroutine?
=====================================================================
ANSWER:

Use the `go` keyword before a function call.

    go myFunction()

Anonymous functions can also be launched:

    go func() {
        // code
    }()
*/

func q02Worker() {
	fmt.Println("Worker running")
}

func q02() {
	fmt.Println("\nQ02")

	done := make(chan bool)

	go q02Worker()

	go func() {
		done <- true
	}()

	<-done
}

/*
=====================================================================
Q03. What happens if main exits before a goroutine finishes?
=====================================================================
ANSWER:

The program exits. Goroutines do not keep the program alive after
the main goroutine returns.

Use synchronization such as WaitGroup or channels.
*/

func q03() {
	fmt.Println("\nQ03")

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		time.Sleep(50 * time.Millisecond)
		fmt.Println("Goroutine completed")
	}()

	wg.Wait()
}

/*
=====================================================================
Q04. Is a goroutine the same as an OS thread?
=====================================================================
ANSWER:

No.

A goroutine is a Go runtime-managed unit of execution. The Go runtime
schedules goroutines onto operating-system threads.

Think:

    Many goroutines
          |
          v
    Go scheduler
          |
          v
    OS threads
          |
          v
    CPU cores
*/

/*
=====================================================================
Q05. What is the main goroutine?
=====================================================================
ANSWER:

The function `main()` runs in the initial goroutine.

When `main()` returns, the entire Go program exits.
*/

func q05() {
	fmt.Println("\nQ05")
	fmt.Println("main() runs in the initial goroutine")
}

/*
=====================================================================
Q06. Create 5 goroutines.
=====================================================================
*/

func q06() {
	fmt.Println("\nQ06")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Println("Worker:", id)
		}(i)
	}

	wg.Wait()
}

/*
=====================================================================
Q07. What is sync.WaitGroup?
=====================================================================
ANSWER:

WaitGroup waits for a collection of goroutines to finish.

Main methods:

    Add(n)
    Done()
    Wait()

Typical pattern:

    wg.Add(1)

    go func() {
        defer wg.Done()
        // work
    }()

    wg.Wait()
*/

func q07() {
	fmt.Println("\nQ07")

	var wg sync.WaitGroup

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Println("Task:", id)
		}(i)
	}

	wg.Wait()

	fmt.Println("All tasks completed")
}

/*
=====================================================================
Q08. Why use defer wg.Done()?
=====================================================================
ANSWER:

It guarantees Done() is called when the goroutine exits, including
when there are multiple return paths.

Example:

    defer wg.Done()
*/

func q08() {
	fmt.Println("\nQ08")

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("Doing work")
	}()

	wg.Wait()
}

/*
=====================================================================
Q09. Can a goroutine return a value directly?
=====================================================================
ANSWER:

No.

A goroutine invocation does not directly provide a return value.

Use a channel to send the result.
*/

func q09() {
	fmt.Println("\nQ09")

	result := make(chan int)

	go func() {
		result <- 100
	}()

	value := <-result

	fmt.Println("Result:", value)
}

/*
=====================================================================
Q10. What is a channel?
=====================================================================
ANSWER:

A channel is a typed communication mechanism used by goroutines.

Create:

    ch := make(chan int)

Send:

    ch <- 10

Receive:

    value := <-ch
*/

func q10() {
	fmt.Println("\nQ10")

	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	value := <-ch

	fmt.Println(value)
}

/*
=====================================================================
Q11. Explain send and receive.
=====================================================================

    ch <- value

means send.

    value := <-ch

means receive.
*/

func q11() {
	fmt.Println("\nQ11")

	ch := make(chan string)

	go func() {
		ch <- "Hello"
	}()

	message := <-ch

	fmt.Println(message)
}

/*
=====================================================================
Q12. What is an unbuffered channel?
=====================================================================
ANSWER:

An unbuffered channel has zero capacity.

    ch := make(chan int)

A send waits until another goroutine receives the value.
*/

func q12() {
	fmt.Println("\nQ12")

	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	fmt.Println(<-ch)
}

/*
=====================================================================
Q13. What is a buffered channel?
=====================================================================
ANSWER:

A buffered channel has capacity.

    ch := make(chan int, 3)

The sender can send up to the available capacity without an
immediate receiver.
*/

func q13() {
	fmt.Println("\nQ13")

	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
=====================================================================
Q14. Difference between buffered and unbuffered channels?
=====================================================================
ANSWER:

Unbuffered:
    make(chan int)

    Sender and receiver synchronize directly.

Buffered:
    make(chan int, N)

    Channel stores up to N values.

Use buffered channels when you intentionally want decoupling between
senders and receivers.
*/

func q14() {
	fmt.Println("\nQ14")

	unbuffered := make(chan int)
	buffered := make(chan int, 2)

	fmt.Println("Unbuffered capacity:", cap(unbuffered))
	fmt.Println("Buffered capacity:", cap(buffered))
}

/*
=====================================================================
Q15. What is channel blocking?
=====================================================================
ANSWER:

A send blocks when the channel cannot accept the value.

A receive blocks when no value is available.

This synchronization behavior is one of the most important channel
concepts.
*/

func q15() {
	fmt.Println("\nQ15")

	ch := make(chan int)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- 10
	}()

	fmt.Println(<-ch)
}

/*
=====================================================================
Q16. How do you close a channel?
=====================================================================
*/

func q16() {
	fmt.Println("\nQ16")

	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for value := range ch {
		fmt.Println(value)
	}
}

/*
=====================================================================
Q17. Why close a channel?
=====================================================================
ANSWER:

Closing tells receivers that no more values will be sent.

Important:

Closing a channel does NOT destroy its already-sent values.

Receivers can still read remaining buffered values.

Usually the sender is responsible for closing the channel.
*/

/*
=====================================================================
Q18. What happens when receiving from a closed channel?
=====================================================================
ANSWER:

The receive returns the zero value after all values have been
consumed.

Better:

    value, ok := <-ch

If:

    ok == false

the channel is closed and empty.
*/

func q18() {
	fmt.Println("\nQ18")

	ch := make(chan int)

	close(ch)

	value, ok := <-ch

	fmt.Println("Value:", value)
	fmt.Println("Open:", ok)
}

/*
=====================================================================
Q19. What happens if you send to a closed channel?
=====================================================================
ANSWER:

It causes a panic.

    panic: send on closed channel

Therefore, don't send after closing a channel.
*/

/*
=====================================================================
Q20. What happens if you close a channel twice?
=====================================================================
ANSWER:

It causes a panic.

    close of closed channel

Only close a channel once.
*/

/*
=====================================================================
Q21. What is range over a channel?
=====================================================================
ANSWER:

You can receive values until the channel is closed.

    for value := range ch {
        ...
    }

The loop ends when the channel is closed and drained.
*/

func q21() {
	fmt.Println("\nQ21")

	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	close(ch)

	for value := range ch {
		fmt.Println(value)
	}
}

/*
=====================================================================
Q22. What is select?
=====================================================================
ANSWER:

select lets a goroutine wait on multiple channel operations.

It is similar to switch, but for channel communication.
*/

func q22() {
	fmt.Println("\nQ22")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "from channel 1"
	}()

	go func() {
		ch2 <- "from channel 2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}

/*
=====================================================================
Q23. Why use select?
=====================================================================
ANSWER:

Common uses:

    - Multiple channels
    - Timeouts
    - Cancellation
    - Non-blocking communication
    - Worker coordination
*/

/*
=====================================================================
Q24. How do you implement a timeout?
=====================================================================
*/

func q24() {
	fmt.Println("\nQ24")

	ch := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "completed"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)

	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timeout")
	}
}

/*
=====================================================================
Q25. What is a non-blocking channel operation?
=====================================================================
ANSWER:

Use select with default.

*/

func q25() {
	fmt.Println("\nQ25")

	ch := make(chan int)

	select {
	case value := <-ch:
		fmt.Println("Received:", value)

	default:
		fmt.Println("No value available")
	}
}

/*
=====================================================================
Q26. Explain directional channels.
=====================================================================
ANSWER:

You can restrict channel usage.

Send-only:

    chan<- int

Receive-only:

    <-chan int

This improves API safety.
*/

func sendOnly26(ch chan<- int) {
	ch <- 100
}

func receiveOnly26(ch <-chan int) {
	fmt.Println(<-ch)
}

func q26() {
	fmt.Println("\nQ26")

	ch := make(chan int)

	go sendOnly26(ch)

	receiveOnly26(ch)
}

/*
=====================================================================
Q27. What is a worker goroutine?
=====================================================================
ANSWER:

A worker goroutine repeatedly receives jobs, processes them, and
produces results.

This pattern is heavily used in backend systems.
*/

func q27Worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		results <- job * 2
	}
}

func q27() {
	fmt.Println("\nQ27")

	jobs := make(chan int, 3)
	results := make(chan int, 3)

	var wg sync.WaitGroup

	wg.Add(1)
	go q27Worker(jobs, results, &wg)

	for i := 1; i <= 3; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	close(results)

	for result := range results {
		fmt.Println(result)
	}
}

/*
=====================================================================
Q28. Explain worker pool.
=====================================================================
ANSWER:

A worker pool consists of:

    jobs channel
    multiple worker goroutines
    results channel

Example:

    jobs -> worker 1
         -> worker 2
         -> worker 3

Useful for:

    database jobs
    HTTP requests
    file processing
    background tasks
    message processing
*/

func q28Worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func q28() {
	fmt.Println("\nQ28")

	jobs := make(chan int, 5)

	var wg sync.WaitGroup

	for workerID := 1; workerID <= 3; workerID++ {
		wg.Add(1)
		go q28Worker(workerID, jobs, &wg)
	}

	for job := 1; job <= 5; job++ {
		jobs <- job
	}

	close(jobs)

	wg.Wait()
}

/*
=====================================================================
Q29. What is a race condition?
=====================================================================
ANSWER:

A race condition occurs when multiple goroutines access shared
mutable data concurrently and the result depends on timing.

Example:

    counter++

is not safe when many goroutines execute it concurrently.
*/

func q29() {
	fmt.Println("\nQ29")
	fmt.Println("Use synchronization for shared mutable state.")
}

/*
=====================================================================
Q30. How do you detect data races?
=====================================================================
ANSWER:

Use Go's race detector:

    go run -race main.go

or:

    go test -race ./...

This is an extremely useful interview command.
*/

/*
=====================================================================
Q31. Fix a race using sync.Mutex.
=====================================================================
*/

func q31() {
	fmt.Println("\nQ31")

	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}

/*
=====================================================================
Q32. What is sync.RWMutex?
=====================================================================
ANSWER:

RWMutex supports:

    RLock / RUnlock
    Lock / Unlock

Multiple readers can read simultaneously.

A writer gets exclusive access.

Use it when read-heavy workloads justify the additional complexity.
*/

func q32() {
	fmt.Println("\nQ32")

	var (
		mu    sync.RWMutex
		value int
		wg    sync.WaitGroup
	)

	value = 100

	wg.Add(2)

	go func() {
		defer wg.Done()

		mu.RLock()
		fmt.Println("Reader 1:", value)
		mu.RUnlock()
	}()

	go func() {
		defer wg.Done()

		mu.RLock()
		fmt.Println("Reader 2:", value)
		mu.RUnlock()
	}()

	wg.Wait()
}

/*
=====================================================================
Q33. What is sync.Once?
=====================================================================
ANSWER:

sync.Once ensures a function is executed only once, even if many
goroutines call it.
*/

func q33() {
	fmt.Println("\nQ33")

	var once sync.Once
	var wg sync.WaitGroup

	initialize := func() {
		fmt.Println("Initialized exactly once")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			once.Do(initialize)
		}()
	}

	wg.Wait()
}

/*
=====================================================================
Q34. What is sync.Cond?
=====================================================================
ANSWER:

sync.Cond provides a condition-variable style synchronization
mechanism.

It is less commonly needed than channels, Mutex, or context.

Main operations:

    Wait()
    Signal()
    Broadcast()
*/

/*
=====================================================================
Q35. What is atomic operation?
=====================================================================
ANSWER:

The sync/atomic package provides low-level atomic operations.

Useful for simple shared counters and flags without a Mutex.

For complex state, prefer Mutex or channels when they make the
design clearer.
*/

func q35() {
	fmt.Println("\nQ35")
	fmt.Println("sync/atomic provides atomic memory operations.")
}

/*
=====================================================================
Q36. What is deadlock?
=====================================================================
ANSWER:

A deadlock occurs when goroutines wait forever for each other.

Example concept:

    goroutine A waits for B
    goroutine B waits for A

Go can report:

    fatal error: all goroutines are asleep - deadlock!
*/

/*
=====================================================================
Q37. Simple deadlock example.
=====================================================================
ANSWER:

This code is intentionally NOT executed because it would deadlock.

    ch := make(chan int)
    ch <- 10

No receiver exists.
*/

/*
=====================================================================
Q38. What is goroutine leak?
=====================================================================
ANSWER:

A goroutine leak occurs when a goroutine remains blocked forever
and can no longer perform useful work.

Common causes:

    - Sending with no receiver
    - Receiving with no sender
    - Forgotten channel close
    - Ignoring cancellation
    - Waiting forever on a dependency
*/

/*
=====================================================================
Q39. How do you prevent goroutine leaks?
=====================================================================
ANSWER:

Use:

    - context cancellation
    - timeouts
    - channel closure
    - proper ownership
    - select
    - bounded worker pools
    - guaranteed cleanup
*/

/*
=====================================================================
Q40. What is context.Context?
=====================================================================
ANSWER:

context.Context carries cancellation, deadlines, and request-scoped
values across API boundaries and goroutines.

Common:

    context.WithCancel
    context.WithTimeout
    context.WithDeadline

For server applications, context is extremely important.
*/

func q40() {
	fmt.Println("\nQ40")

	ctx, cancel := contextWithCancel40()

	defer cancel()

	_ = ctx

	fmt.Println("Context created")
}

// Small wrapper so this file remains focused on the concept.
func contextWithCancel40() (interface{}, func()) {
	return struct{}{}, func() {}
}

/*
NOTE:

Production code normally uses:

    ctx, cancel := context.WithCancel(context.Background())

and passes ctx to goroutines/functions.

This example intentionally avoids importing context here because
later examples demonstrate the full implementation.
*/

/*
=====================================================================
Q41. Context cancellation example.
=====================================================================
*/

func q41() {
	fmt.Println("\nQ41")

	/*
		Production pattern:

		    ctx, cancel := context.WithCancel(context.Background())

		    go func() {
		        for {
		            select {
		            case <-ctx.Done():
		                return
		            default:
		                // work
		            }
		        }
		    }()

		    cancel()
	*/

	fmt.Println("Use ctx.Done() to stop goroutines cooperatively.")
}

/*
=====================================================================
Q42. How can a goroutine listen for cancellation?
=====================================================================
ANSWER:

Use:

    select {
    case <-ctx.Done():
        return
    case job := <-jobs:
        process(job)
    }

This allows the goroutine to exit cleanly.
*/

/*
=====================================================================
Q43. Can multiple goroutines read from the same channel?
=====================================================================
ANSWER:

Yes.

This is a common worker-pool pattern.

Each worker receives jobs from the same jobs channel.
*/

func q43() {
	fmt.Println("\nQ43")

	jobs := make(chan int, 5)

	var wg sync.WaitGroup

	for worker := 1; worker <= 2; worker++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for job := range jobs {
				fmt.Printf("Worker %d got job %d\n", id, job)
			}
		}(worker)
	}

	for job := 1; job <= 5; job++ {
		jobs <- job
	}

	close(jobs)

	wg.Wait()
}

/*
=====================================================================
Q44. Can multiple goroutines send to the same channel?
=====================================================================
ANSWER:

Yes.

Multiple producers can safely send to a channel.

However, channel closing becomes an ownership/design concern.

The goroutine responsible for knowing that no more sends will occur
should close the channel.
*/

/*
=====================================================================
Q45. What happens if multiple goroutines close the same channel?
=====================================================================
ANSWER:

The second close causes:

    panic: close of closed channel

Channel ownership should be clear.
*/

/*
=====================================================================
Q46. Explain fan-out.
=====================================================================
ANSWER:

Fan-out means distributing work from one source to multiple workers.

    jobs
      |
      +--> worker 1
      +--> worker 2
      +--> worker 3
*/

func q46() {
	fmt.Println("\nQ46")

	jobs := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(worker int) {
			defer wg.Done()

			for job := range jobs {
				fmt.Printf("Worker %d processed %d\n", worker, job)
			}
		}(i)
	}

	for job := 1; job <= 6; job++ {
		jobs <- job
	}

	close(jobs)

	wg.Wait()
}

/*
=====================================================================
Q47. Explain fan-in.
=====================================================================
ANSWER:

Fan-in means combining multiple input streams into one channel.

    worker 1 --\
    worker 2 ----> results
    worker 3 --/
*/

func q47() {
	fmt.Println("\nQ47")

	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	out := make(chan int, 4)

	ch1 <- 1
	ch1 <- 2
	close(ch1)

	ch2 <- 3
	ch2 <- 4
	close(ch2)

	go func() {
		for value := range ch1 {
			out <- value
		}

		for value := range ch2 {
			out <- value
		}

		close(out)
	}()

	for value := range out {
		fmt.Println(value)
	}
}

/*
=====================================================================
Q48. What is a pipeline?
=====================================================================
ANSWER:

A pipeline has multiple processing stages.

Example:

    input
      |
      v
    square
      |
      v
    filter
      |
      v
    output

Each stage can be represented by a goroutine and channel.
*/

func q48() {
	fmt.Println("\nQ48")

	input := make(chan int, 5)
	squared := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		input <- i
	}

	close(input)

	go func() {
		defer close(squared)

		for value := range input {
			squared <- value * value
		}
	}()

	for value := range squared {
		fmt.Println(value)
	}
}

/*
=====================================================================
Q49. What is a semaphore pattern in Go?
=====================================================================
ANSWER:

A buffered channel can limit concurrent work.

Example:

    semaphore := make(chan struct{}, 3)

Only 3 goroutines can hold a slot at once.
*/

func q49() {
	fmt.Println("\nQ49")

	sem := make(chan struct{}, 2)

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}
			fmt.Println("Running:", id)

			time.Sleep(20 * time.Millisecond)

			<-sem
		}(i)
	}

	wg.Wait()
}

/*
=====================================================================
Q50. How do you limit the number of concurrent HTTP/API calls?
=====================================================================
ANSWER:

Use a worker pool or semaphore.

Example:

    sem := make(chan struct{}, 10)

At most 10 operations can run simultaneously.

In production, also consider:

    - timeouts
    - context cancellation
    - connection limits
    - rate limits
    - retries
*/

/*
=====================================================================
Q51. What is the loop-variable goroutine problem?
=====================================================================
ANSWER:

When launching goroutines inside loops, be careful about what value
the closure captures.

Safest interview-friendly pattern:

    for i := 0; i < 5; i++ {
        go func(id int) {
            fmt.Println(id)
        }(i)
    }

Passing i as a parameter makes the intended value explicit.
*/

func q51() {
	fmt.Println("\nQ51")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Println(id)
		}(i)
	}

	wg.Wait()
}

/*
=====================================================================
Q52. Why can fmt.Println output appear in different order?
=====================================================================
ANSWER:

Goroutines execute concurrently.

The scheduler decides when goroutines run.

Therefore:

    go print(1)
    go print(2)
    go print(3)

does not guarantee:

    1
    2
    3

If ordering matters, synchronize explicitly.
*/

func q52() {
	fmt.Println("\nQ52")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Println("Output:", id)
		}(i)
	}

	wg.Wait()
}

/*
=====================================================================
Q53. Is concurrent execution the same as parallel execution?
=====================================================================
ANSWER:

No.

Concurrency:
    Multiple tasks can make progress during overlapping periods.

Parallelism:
    Multiple tasks actually execute simultaneously on multiple
    CPU cores.

Go supports both concurrency and parallelism.
*/

/*
=====================================================================
Q54. What is GOMAXPROCS?
=====================================================================
ANSWER:

GOMAXPROCS controls the maximum number of operating-system threads
that can execute Go code simultaneously.

Modern Go normally chooses a sensible default based on available
CPUs.

You can inspect it using runtime.GOMAXPROCS(0).

*/

func q54() {
	fmt.Println("\nQ54")

	// Example in real code:
	// fmt.Println(runtime.GOMAXPROCS(0))

	fmt.Println("GOMAXPROCS controls parallel execution capacity.")
}

/*
=====================================================================
Q55. What is the Go scheduler?
=====================================================================
ANSWER:

The Go runtime scheduler schedules goroutines onto OS threads.

A simplified model:

    G = Goroutine
    M = Machine / OS thread
    P = Processor context

Conceptually:

    G ---> P ---> M

The runtime manages this so developers don't manually assign every
goroutine to an OS thread.
*/

/*
=====================================================================
Q56. Explain G-M-P model.
=====================================================================
ANSWER:

G:
    Goroutine

M:
    OS thread

P:
    Scheduler resource required to execute Go code

A simplified view:

          P
         / \
        G   G
         |
         M
         |
        CPU

The scheduler moves goroutines between available execution resources.
*/

/*
=====================================================================
Q57. What is runtime.Gosched()?
=====================================================================
ANSWER:

runtime.Gosched() voluntarily yields the processor so another
goroutine can run.

It is rarely needed in normal application code.
Use synchronization primitives instead of relying on scheduling
behavior.
*/

/*
=====================================================================
Q58. What is runtime.Goexit()?
=====================================================================
ANSWER:

runtime.Goexit() terminates the current goroutine.

Deferred functions in that goroutine still execute.

It should rarely be needed in normal application code.
*/

/*
=====================================================================
Q59. What is a WaitGroup misuse?
=====================================================================
ANSWER:

Common mistakes:

    1. Calling Done more times than Add.
    2. Forgetting Done.
    3. Calling Add too late.
    4. Reusing a WaitGroup incorrectly while Wait is active.

Best pattern:

    wg.Add(1)

    go func() {
        defer wg.Done()
    }()
*/

func q59() {
	fmt.Println("\nQ59")

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Safe WaitGroup usage")
	}()

	wg.Wait()
}

/*
=====================================================================
Q60. What is the difference between Mutex and channel?
=====================================================================
ANSWER:

Mutex:
    Protects shared memory.

Channel:
    Communicates/transfers data between goroutines.

Simple rule:

    "Don't communicate by sharing memory; share memory by
     communicating."

But this is a design guideline, not an absolute law.

Mutex is often the right tool for protecting a small shared data
structure.
*/

/*
=====================================================================
Q61. When should you use a Mutex?
=====================================================================
ANSWER:

Use Mutex when:

    - Multiple goroutines share mutable state.
    - The critical section is small.
    - You want direct protection of a data structure.
    - A channel would make the design unnecessarily complicated.

Example:

    mu.Lock()
    sharedMap[key] = value
    mu.Unlock()
*/

/*
=====================================================================
Q62. When should you use a channel?
=====================================================================
ANSWER:

Use channels when goroutines need to:

    - send jobs
    - send results
    - signal completion
    - build pipelines
    - coordinate cancellation/events

Channels express communication and ownership clearly.
*/

/*
=====================================================================
Q63. Can maps be accessed concurrently?
=====================================================================
ANSWER:

A normal Go map should not be read/written concurrently without
proper synchronization.

Use:

    sync.Mutex
    sync.RWMutex
    sync.Map

depending on the workload and design.
*/

/*
=====================================================================
Q64. What is sync.Map?
=====================================================================
ANSWER:

sync.Map is a specialized concurrent map implementation.

It can be useful for certain workloads with many goroutines,
especially when entries are long-lived and reads dominate.

Do not automatically replace every map with sync.Map.

For many cases:

    map + Mutex

is simpler and easier to reason about.
*/

/*
=====================================================================
Q65. What is a goroutine-safe counter?
=====================================================================
ANSWER:

Possible solutions:

    1. Mutex
    2. atomic operations
    3. A dedicated counter goroutine receiving increment requests
       through a channel

Choose based on the design.
*/

func q65() {
	fmt.Println("\nQ65")

	var (
		mu      sync.Mutex
		counter int
		wg      sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}

/*
=====================================================================
Q66. Build a concurrent job processor with results.
=====================================================================
ANSWER:

This combines:

    goroutines
    WaitGroup
    jobs channel
    results channel
    channel closing
*/

func q66Worker(
	id int,
	jobs <-chan int,
	results chan<- string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for job := range jobs {
		results <- fmt.Sprintf(
			"worker %d processed job %d",
			id,
			job,
		)
	}
}

func q66() {
	fmt.Println("\nQ66")

	jobs := make(chan int, 10)
	results := make(chan string, 10)

	var wg sync.WaitGroup

	for worker := 1; worker <= 3; worker++ {
		wg.Add(1)

		go q66Worker(
			worker,
			jobs,
			results,
			&wg,
		)
	}

	for job := 1; job <= 10; job++ {
		jobs <- job
	}

	close(jobs)

	wg.Wait()

	close(results)

	for result := range results {
		fmt.Println(result)
	}
}

/*
=====================================================================
Q67. How do you gracefully stop a worker pool?
=====================================================================
ANSWER:

A common approach:

    1. Stop producing jobs.
    2. Close jobs channel.
    3. Workers finish remaining jobs.
    4. Workers return.
    5. WaitGroup completes.
    6. Close results channel if appropriate.

For cancellation:

    use context.Context.

*/

func q67() {
	fmt.Println("\nQ67")

	jobs := make(chan int, 3)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for job := range jobs {
			fmt.Println("Processing:", job)
		}

		fmt.Println("Worker stopped gracefully")
	}()

	jobs <- 1
	jobs <- 2
	jobs <- 3

	close(jobs)

	wg.Wait()
}

/*
=====================================================================
Q68. Interview problem: run multiple tasks concurrently and wait.
=====================================================================
ANSWER:

Use WaitGroup.

*/

func q68() {
	fmt.Println("\nQ68")

	tasks := []string{
		"Task A",
		"Task B",
		"Task C",
	}

	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)

		go func(name string) {
			defer wg.Done()

			time.Sleep(20 * time.Millisecond)
			fmt.Println("Completed:", name)
		}(task)
	}

	wg.Wait()

	fmt.Println("All tasks finished")
}

/*
=====================================================================
Q69. Interview problem: run tasks concurrently and collect errors.
=====================================================================
ANSWER:

One approach is an error channel.

Production applications may also use errgroup when appropriate.

*/

func q69() {
	fmt.Println("\nQ69")

	tasks := []int{1, 2, 3, 4}

	errCh := make(chan error, len(tasks))

	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			if id == 3 {
				errCh <- fmt.Errorf("task %d failed", id)
				return
			}

			fmt.Println("Task succeeded:", id)
		}(task)
	}

	wg.Wait()

	close(errCh)

	for err := range errCh {
		fmt.Println("Error:", err)
	}
}

/*
=====================================================================
Q70. INTERVIEW MASTER QUESTION:
How would you design a production-grade concurrent worker pool?
=====================================================================
ANSWER:

A strong answer should mention:

1. jobs channel
2. worker goroutines
3. results/error channel if needed
4. context cancellation
5. WaitGroup or equivalent lifecycle management
6. bounded concurrency
7. timeouts where appropriate
8. graceful shutdown
9. no goroutine leaks
10. race-free shared state
11. backpressure
12. metrics/logging
13. error handling
14. panic recovery at appropriate boundaries
15. testing with -race

Architecture:

                 +------------+
                 | Producer   |
                 +-----+------+
                       |
                       v
                 +-----------+
                 | jobs chan |
                 +-----+-----+
                       |
          +------------+------------+
          |            |            |
          v            v            v
       Worker 1     Worker 2     Worker 3
          |            |            |
          +------------+------------+
                       |
                       v
                 +-----------+
                 | results   |
                 +-----------+

For cancellation:

                 context.Context
                       |
          +------------+------------+
          |            |            |
       Worker 1     Worker 2     Worker 3

Workers should regularly select on:

    case <-ctx.Done():
        return

    case job := <-jobs:
        process(job)

This is the kind of answer expected from a developer with real
backend experience.

=====================================================================
IMPORTANT INTERVIEW CONCEPTS — QUICK REVISION
=====================================================================

1. GOROUTINE
   Lightweight concurrent execution unit managed by Go runtime.

2. CHANNEL
   Communication mechanism between goroutines.

3. WAITGROUP
   Waits for goroutines to finish.

4. MUTEX
   Protects shared mutable state.

5. RWMutex
   Allows concurrent readers but exclusive writers.

6. SELECT
   Waits on multiple channel operations.

7. BUFFERED CHANNEL
   Has capacity.

8. UNBUFFERED CHANNEL
   Sender and receiver synchronize directly.

9. CLOSE
   Indicates no more values will be sent.

10. RANGE CHANNEL
    Receives until channel closes.

11. CONTEXT
    Cancellation, deadlines, request-scoped propagation.

12. WORKER POOL
    Fixed/bounded workers processing jobs.

13. FAN-OUT
    One source -> multiple workers.

14. FAN-IN
    Multiple sources -> one output.

15. PIPELINE
    Multiple processing stages connected with channels.

16. RACE CONDITION
    Concurrent unsynchronized access causing nondeterministic
    behavior.

17. DEADLOCK
    Goroutines wait forever.

18. GOROUTINE LEAK
    Goroutine remains blocked and cannot terminate.

19. SEMAPHORE
    Limits concurrency, often using buffered channels.

20. RACE DETECTOR

    go run -race .
    go test -race ./...

=====================================================================
COMMON INTERVIEW TRAPS
=====================================================================

TRAP 1:
"Launching a goroutine guarantees it runs immediately."

FALSE.

TRAP 2:
"Channels completely eliminate race conditions."

FALSE.

Shared state can still be accessed unsafely.

TRAP 3:
"Every goroutine needs a channel."

FALSE.

A goroutine can perform independent work and synchronize with a
WaitGroup or other primitive.

TRAP 4:
"close() is required after every channel use."

FALSE.

Close when receivers need to know no more values will be sent.

TRAP 5:
"The receiver should always close the channel."

Usually FALSE.

The sender that owns the sending lifecycle normally closes it.

TRAP 6:
"Buffered channels never block."

FALSE.

They block when the buffer is full on send, or empty on receive.

TRAP 7:
"Mutex is always worse than channels."

FALSE.

Mutex is often the cleanest solution for protecting shared state.

TRAP 8:
"Concurrency means parallel execution."

FALSE.

Concurrency and parallelism are different concepts.

=====================================================================
MOST IMPORTANT QUESTIONS TO MASTER
=====================================================================

Beginner:
    Q01-Q16

Synchronization:
    Q17-Q35

Concurrency bugs:
    Q36-Q45

Advanced design:
    Q46-Q70

For a 2-year Go interview, pay special attention to:

    Q07   WaitGroup
    Q10   Channels
    Q12   Unbuffered channels
    Q13   Buffered channels
    Q16   Closing channels
    Q22   select
    Q24   Timeout
    Q27   Worker
    Q28   Worker pool
    Q29   Race condition
    Q31   Mutex
    Q32   RWMutex
    Q33   sync.Once
    Q38   Goroutine leak
    Q40-Q42 Context
    Q46   Fan-out
    Q47   Fan-in
    Q48   Pipeline
    Q49   Semaphore
    Q51   Loop variables
    Q53   Concurrency vs parallelism
    Q54-Q56 Scheduler
    Q60-Q64 Mutex/channel/map decisions
    Q66-Q70 Production concurrency

=====================================================================
PRACTICE COMMANDS
=====================================================================

Run:

    go run goroutines_70.go

Race detector:

    go run -race goroutines_70.go

Format:

    gofmt -w goroutines_70.go

Build:

    go build goroutines_70.go

Test race conditions:

    go test -race ./...

=====================================================================
END — 70 GOROUTINE QUESTIONS & ANSWERS
=====================================================================
*/

func main() {
	q01()
	q02()
	q03()
	q04()
	q05()
	q06()
	q07()
	q08()
	q09()
	q10()
	q11()
	q12()
	q13()
	q14()
	q15()
	q16()
	q18()
	q21()
	q22()
	q24()
	q25()
	q26()
	q27()
	q28()
	q31()
	q32()
	q33()
	q43()
	q46()
	q47()
	q48()
	q49()
	q51()
	q52()
	q54()
	q59()
	q65()
	q66()
	q67()
	q68()
	q69()
	q70()

	fmt.Println("\n===============================================================")
	fmt.Println("70 GOROUTINE QUESTIONS COMPLETED")
	fmt.Println("===============================================================")
}
