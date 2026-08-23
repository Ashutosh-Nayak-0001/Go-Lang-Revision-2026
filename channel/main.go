package main

import (
	"fmt"
	"sync"
	"time"
)

/*
======================================================================
GO CHANNELS — 60 QUESTIONS & ANSWERS
======================================================================

This file is designed for Go coding-interview practice.

LEVEL:
    Q01-Q20  Beginner
    Q21-Q40  Intermediate
    Q41-Q60  Advanced / Interview

Run:
    go run channels_60.go

Race detector:
    go run -race channels_60.go

======================================================================
Q01. What is a channel in Go?
======================================================================

ANSWER:
A channel is a typed communication mechanism that allows goroutines
to safely exchange values.

Create:
    ch := make(chan int)

Send:
    ch <- 10

Receive:
    value := <-ch

Channels are mainly used for communication and synchronization
between goroutines.
*/

func q01() {
	fmt.Println("\nQ01 - Basic channel")

	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	value := <-ch
	fmt.Println("Received:", value)
}

/*
======================================================================
Q02. How do you create a channel?
======================================================================

ANSWER:

    ch := make(chan int)

The channel type is chan int.

A channel can carry only values of its declared type.
*/

func q02() {
	fmt.Println("\nQ02")

	ch := make(chan string)

	go func() {
		ch <- "Hello Go"
	}()

	fmt.Println(<-ch)
}

/*
======================================================================
Q03. How do you send a value to a channel?
======================================================================

ANSWER:

    ch <- value
*/

func q03() {
	fmt.Println("\nQ03")

	ch := make(chan int)

	go func() {
		ch <- 50
	}()

	fmt.Println(<-ch)
}

/*
======================================================================
Q04. How do you receive a value from a channel?
======================================================================

ANSWER:

    value := <-ch
*/

func q04() {
	fmt.Println("\nQ04")

	ch := make(chan int)

	go func() {
		ch <- 75
	}()

	value := <-ch

	fmt.Println(value)
}

/*
======================================================================
Q05. What is an unbuffered channel?
======================================================================

ANSWER:

An unbuffered channel has zero capacity.

    ch := make(chan int)

A send blocks until another goroutine receives the value.

An unbuffered channel provides direct synchronization between sender
and receiver.
*/

func q05() {
	fmt.Println("\nQ05")

	ch := make(chan int)

	go func() {
		fmt.Println("Sending...")
		ch <- 10
		fmt.Println("Send completed")
	}()

	value := <-ch
	fmt.Println("Received:", value)
}

/*
======================================================================
Q06. What is a buffered channel?
======================================================================

ANSWER:

A buffered channel has a fixed capacity.

    ch := make(chan int, 3)

The sender can send values while capacity is available.
*/

func q06() {
	fmt.Println("\nQ06")

	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
======================================================================
Q07. Difference between buffered and unbuffered channel?
======================================================================

ANSWER:

Unbuffered:
    make(chan int)

    Send waits for receiver.

Buffered:
    make(chan int, 3)

    Send can continue until buffer becomes full.

Interview line:

"An unbuffered channel synchronizes sender and receiver directly,
while a buffered channel provides temporary storage up to its
capacity."
*/

func q07() {
	fmt.Println("\nQ07")

	unbuffered := make(chan int)
	buffered := make(chan int, 3)

	fmt.Println("Unbuffered capacity:", cap(unbuffered))
	fmt.Println("Buffered capacity:", cap(buffered))
}

/*
======================================================================
Q08. What does cap(channel) return?
======================================================================

ANSWER:

cap(ch) returns channel capacity.

For:

    ch := make(chan int, 5)

cap(ch) == 5

For an unbuffered channel:

    cap(ch) == 0
*/

func q08() {
	fmt.Println("\nQ08")

	ch := make(chan int, 5)

	fmt.Println("Capacity:", cap(ch))
}

/*
======================================================================
Q09. What does len(channel) return?
======================================================================

ANSWER:

For a buffered channel, len(ch) returns the number of elements
currently waiting in the channel buffer.

It does NOT return the capacity.

Example:

    ch := make(chan int, 5)
    ch <- 10
    ch <- 20

len(ch) == 2
cap(ch) == 5
*/

func q09() {
	fmt.Println("\nQ09")

	ch := make(chan int, 5)

	ch <- 10
	ch <- 20

	fmt.Println("Length:", len(ch))
	fmt.Println("Capacity:", cap(ch))
}

/*
======================================================================
Q10. What happens when an unbuffered send has no receiver?
======================================================================

ANSWER:

The sender blocks.

If no goroutine can ever receive, the program can deadlock.
*/

func q10() {
	fmt.Println("\nQ10")

	ch := make(chan int)

	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(<-ch)
	}()

	ch <- 100

	fmt.Println("Send completed")
}

/*
======================================================================
Q11. What happens when receiving from an empty channel?
======================================================================

ANSWER:

The receiver blocks until a value becomes available or the channel
is closed.
*/

func q11() {
	fmt.Println("\nQ11")

	ch := make(chan int)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- 200
	}()

	fmt.Println(<-ch)
}

/*
======================================================================
Q12. What happens when a buffered channel becomes full?
======================================================================

ANSWER:

A new send blocks until another goroutine receives a value.
*/

func q12() {
	fmt.Println("\nQ12")

	ch := make(chan int, 2)

	ch <- 10
	ch <- 20

	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Received:", <-ch)
	}()

	ch <- 30

	fmt.Println("Third send completed")
}

/*
======================================================================
Q13. How do you close a channel?
======================================================================

ANSWER:

    close(ch)

Closing means:

"No more values will be sent."

It does not mean that existing values disappear.
*/

func q13() {
	fmt.Println("\nQ13")

	ch := make(chan int, 2)

	ch <- 10
	ch <- 20

	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
======================================================================
Q14. What happens when you receive from a closed channel?
======================================================================

ANSWER:

If the channel has remaining values, those values are returned first.

After all values are consumed, receiving returns the zero value.

Better pattern:

    value, ok := <-ch

When ok == false, the channel is closed and empty.
*/

func q14() {
	fmt.Println("\nQ14")

	ch := make(chan int, 1)

	ch <- 10
	close(ch)

	value, ok := <-ch
	fmt.Println(value, ok)

	value, ok = <-ch
	fmt.Println(value, ok)
}

/*
======================================================================
Q15. What happens when you send to a closed channel?
======================================================================

ANSWER:

It causes a panic:

    panic: send on closed channel

Never send after the channel has been closed.
*/

/*
======================================================================
Q16. What happens if you close a channel twice?
======================================================================

ANSWER:

It causes a panic:

    panic: close of closed channel

A channel should be closed only once.
*/

/*
======================================================================
Q17. Who should close a channel?
======================================================================

ANSWER:

Usually the sender, or the goroutine that owns the sending lifecycle,
should close the channel.

The receiver normally should not close a channel because the receiver
usually does not know whether other senders still exist.
*/

/*
======================================================================
Q18. What is the comma-ok syntax for channels?
======================================================================

ANSWER:

    value, ok := <-ch

If:

    ok == true

a value was received.

If:

    ok == false

the channel is closed and empty.
*/

func q18() {
	fmt.Println("\nQ18")

	ch := make(chan string)

	close(ch)

	value, ok := <-ch

	fmt.Println("Value:", value)
	fmt.Println("OK:", ok)
}

/*
======================================================================
Q19. How do you range over a channel?
======================================================================

ANSWER:

    for value := range ch {
        ...
    }

The loop ends when the channel is closed and all buffered values
have been consumed.
*/

func q19() {
	fmt.Println("\nQ19")

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
======================================================================
Q20. What happens if range is used on a channel that is never closed?
======================================================================

ANSWER:

The range loop waits forever after all available values are consumed.

This can cause a goroutine leak or deadlock depending on the design.

Always make channel ownership and closing behavior clear.
*/

/*
======================================================================
Q21. What is a send-only channel?
======================================================================

ANSWER:

A send-only channel is:

    chan<- int

A function receiving this type can send values but cannot receive.
*/

func q21Send(ch chan<- int) {
	ch <- 100
}

func q21() {
	fmt.Println("\nQ21")

	ch := make(chan int)

	go q21Send(ch)

	fmt.Println(<-ch)
}

/*
======================================================================
Q22. What is a receive-only channel?
======================================================================

ANSWER:

A receive-only channel is:

    <-chan int

The function can receive but cannot send.
*/

func q22Receive(ch <-chan int) {
	fmt.Println(<-ch)
}

func q22() {
	fmt.Println("\nQ22")

	ch := make(chan int)

	go func() {
		ch <- 200
	}()

	q22Receive(ch)
}

/*
======================================================================
Q23. Why use directional channels?
======================================================================

ANSWER:

Directional channels improve API safety.

Example:

    func producer() <-chan int

The caller can only receive from the returned channel.

This prevents accidental sends or closes from the wrong side.
*/

/*
======================================================================
Q24. Can a bidirectional channel be passed to a send-only parameter?
======================================================================

ANSWER:

Yes.

A normal channel:

    chan int

can be used where:

    chan<- int

is expected.

Similarly, it can be passed to:

    <-chan int
*/

func q24() {
	fmt.Println("\nQ24")

	ch := make(chan int)

	go func() {
		ch <- 500
	}()

	fmt.Println(<-ch)
}

/*
======================================================================
Q25. What is select with channels?
======================================================================

ANSWER:

select waits for multiple channel operations.

Example:

    select {
    case v := <-ch1:
        ...
    case v := <-ch2:
        ...
    }

If multiple cases are ready, select chooses one pseudo-randomly.
*/

func q25() {
	fmt.Println("\nQ25")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Channel 1"
	}()

	go func() {
		ch2 <- "Channel 2"
	}()

	select {
	case value := <-ch1:
		fmt.Println(value)
	case value := <-ch2:
		fmt.Println(value)
	}
}

/*
======================================================================
Q26. How do you implement a channel timeout?
======================================================================

ANSWER:

Use select with time.After or a context deadline.
*/

func q26() {
	fmt.Println("\nQ26")

	ch := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "Done"
	}()

	select {
	case value := <-ch:
		fmt.Println(value)

	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timeout")
	}
}

/*
======================================================================
Q27. What is a non-blocking receive?
======================================================================

ANSWER:

Use select with default.
*/

func q27() {
	fmt.Println("\nQ27")

	ch := make(chan int)

	select {
	case value := <-ch:
		fmt.Println("Received:", value)

	default:
		fmt.Println("No value available")
	}
}

/*
======================================================================
Q28. What is a non-blocking send?
======================================================================

ANSWER:

Use select with default.
*/

func q28() {
	fmt.Println("\nQ28")

	ch := make(chan int, 1)

	select {
	case ch <- 100:
		fmt.Println("Sent successfully")

	default:
		fmt.Println("Channel is full")
	}
}

/*
======================================================================
Q29. How do you check whether a channel is closed?
======================================================================

ANSWER:

Receive using comma-ok:

    value, ok := <-ch

There is no general "isClosed(ch)" operation that is safe for
concurrent use.

The proper design is usually to receive from the channel or use
ownership rules.
*/

/*
======================================================================
Q30. Can a closed channel be selected?
======================================================================

ANSWER:

Yes.

A receive from a closed channel is immediately ready.

Therefore, if you repeatedly select a closed channel without removing
that case, it can continuously win.

Typical solution:
set the channel variable to nil after detecting closure.
*/

/*
======================================================================
Q31. What is a nil channel?
======================================================================

ANSWER:

A nil channel is:

    var ch chan int

Its value is nil.

Sending to or receiving from a nil channel blocks forever.

Closing a nil channel causes a panic.
*/

func q31() {
	fmt.Println("\nQ31")

	var ch chan int

	fmt.Println("Channel is nil:", ch == nil)
}

/*
======================================================================
Q32. Why is nil channel useful with select?
======================================================================

ANSWER:

A nil channel disables its select case.

Example:

    if ch == nil {
        // corresponding select case effectively disabled
    }

This is useful for dynamically enabling/disabling channel cases.
*/

/*
======================================================================
Q33. What happens with close(nil channel)?
======================================================================

ANSWER:

It panics.

    close(ch)

when ch == nil causes:

    panic: close of nil channel
*/

/*
======================================================================
Q34. What is the zero value of a channel?
======================================================================

ANSWER:

The zero value of a channel is nil.

    var ch chan int

You cannot use it for normal communication until initialized:

    ch = make(chan int)
*/

/*
======================================================================
Q35. Can you compare channels?
======================================================================

ANSWER:

Yes.

Channels can be compared with == and !=.

A channel can be compared with nil.

Two channel values are equal when they refer to the same channel.
*/

func q35() {
	fmt.Println("\nQ35")

	ch1 := make(chan int)
	ch2 := ch1

	fmt.Println(ch1 == ch2)
}

/*
======================================================================
Q36. What is a channel deadlock?
======================================================================

ANSWER:

A deadlock occurs when goroutines wait forever for channel operations
that can never happen.

Classic example:

    ch := make(chan int)
    ch <- 10

No receiver exists, so the send blocks forever.
*/

/*
======================================================================
Q37. Explain this deadlock:

    ch := make(chan int)
    fmt.Println(<-ch)

ANSWER:

The receive blocks forever because:

    - channel is unbuffered
    - no goroutine sends a value
    - main goroutine waits forever

The runtime can report:

    fatal error: all goroutines are asleep - deadlock!
*/

/*
======================================================================
Q38. What is a goroutine leak caused by a channel?
======================================================================

ANSWER:

Example:

    ch := make(chan int)

    go func() {
        ch <- 10
    }()

If nobody ever receives from ch, the goroutine remains blocked.

That goroutine can stay alive unnecessarily.

This is a goroutine leak.
*/

/*
======================================================================
Q39. How do you avoid channel leaks?
======================================================================

ANSWER:

Use:

    - context cancellation
    - timeout
    - proper receiver lifecycle
    - channel closure
    - select
    - bounded queues
    - guaranteed cleanup

Every goroutine should have a clear termination path.
*/

/*
======================================================================
Q40. What is a worker pool using channels?
======================================================================

ANSWER:

A worker pool usually has:

    jobs channel
    multiple worker goroutines
    optional results channel

Architecture:

    jobs
      |
      +--> Worker 1
      +--> Worker 2
      +--> Worker 3
      |
    results
*/

func q40Worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processed job %d\n", id, job)
	}
}

func q40() {
	fmt.Println("\nQ40")

	jobs := make(chan int, 5)

	var wg sync.WaitGroup

	for worker := 1; worker <= 3; worker++ {
		wg.Add(1)
		go q40Worker(worker, jobs, &wg)
	}

	for job := 1; job <= 5; job++ {
		jobs <- job
	}

	close(jobs)

	wg.Wait()
}

/*
======================================================================
Q41. What is fan-out?
======================================================================

ANSWER:

Fan-out means distributing work from one channel to multiple workers.

    jobs
      |
      +--> worker 1
      +--> worker 2
      +--> worker 3
*/

func q41() {
	fmt.Println("\nQ41")

	jobs := make(chan int)

	var wg sync.WaitGroup

	for worker := 1; worker <= 3; worker++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for job := range jobs {
				fmt.Printf("Worker %d got %d\n", id, job)
			}
		}(worker)
	}

	for job := 1; job <= 6; job++ {
		jobs <- job
	}

	close(jobs)

	wg.Wait()
}

/*
======================================================================
Q42. What is fan-in?
======================================================================

ANSWER:

Fan-in combines multiple streams into one output channel.

    ch1 ----\
    ch2 -----> out
    ch3 ----/
*/

func q42() {
	fmt.Println("\nQ42")

	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	out := make(chan int, 4)

	ch1 <- 1
	ch1 <- 2
	close(ch1)

	ch2 <- 3
	ch2 <- 4
	close(ch2)

	var wg sync.WaitGroup

	wg.Add(2)

	forward := func(input <-chan int) {
		defer wg.Done()

		for value := range input {
			out <- value
		}
	}

	go forward(ch1)
	go forward(ch2)

	go func() {
		wg.Wait()
		close(out)
	}()

	for value := range out {
		fmt.Println(value)
	}
}

/*
======================================================================
Q43. What is a channel pipeline?
======================================================================

ANSWER:

A pipeline has multiple stages connected by channels.

Example:

    input -> square -> output

Each stage can run as a goroutine.
*/

func q43() {
	fmt.Println("\nQ43")

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
======================================================================
Q44. How do you use a channel as a semaphore?
======================================================================

ANSWER:

A buffered channel can limit the number of concurrent operations.

    sem := make(chan struct{}, 2)

At most 2 goroutines can acquire a slot.
*/

func q44() {
	fmt.Println("\nQ44")

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
======================================================================
Q45. How do you send results through a channel?
======================================================================

ANSWER:

Create a result channel and have goroutines send their results.

The receiver reads from the channel.

This avoids direct return values from goroutines.
*/

func q45() {
	fmt.Println("\nQ45")

	results := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		go func(n int) {
			results <- n * n
		}(i)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Result:", <-results)
	}
}

/*
======================================================================
Q46. How do you collect errors through a channel?
======================================================================

ANSWER:

Create an error channel.

Important:
Buffering the error channel can prevent a worker from becoming stuck
while the main goroutine waits for all workers.
*/

func q46() {
	fmt.Println("\nQ46")

	errCh := make(chan error, 3)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			if id == 2 {
				errCh <- fmt.Errorf("worker %d failed", id)
			}
		}(i)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		fmt.Println("Error:", err)
	}
}

/*
======================================================================
Q47. What is backpressure in channels?
======================================================================

ANSWER:

Backpressure happens when producers generate data faster than
consumers can process it.

A bounded buffered channel naturally applies backpressure:

    producer -> bounded channel -> consumer

When the buffer fills, the producer blocks.

This can prevent unlimited memory growth.
*/

/*
======================================================================
Q48. Why can an unbounded queue be dangerous?
======================================================================

ANSWER:

If producers keep generating work faster than consumers process it,
memory usage can grow continuously.

A bounded channel provides a natural limit.
*/

/*
======================================================================
Q49. How do you implement a timeout for multiple channel operations?
======================================================================

ANSWER:

Use select.

    select {
    case result := <-resultCh:
        ...
    case <-time.After(timeout):
        ...
    }

In production request flows, context deadlines are often preferable.
*/

func q49() {
	fmt.Println("\nQ49")

	resultCh := make(chan string)

	go func() {
		time.Sleep(30 * time.Millisecond)
		resultCh <- "success"
	}()

	select {
	case result := <-resultCh:
		fmt.Println(result)

	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	}
}

/*
======================================================================
Q50. How do you disable a select case?
======================================================================

ANSWER:

Assign nil to the channel.

Example:

    ch = nil

A send or receive on a nil channel blocks forever, so its select
case is effectively disabled.

This is useful for dynamically controlling select loops.
*/

/*
======================================================================
Q51. What is the common select loop pattern?
======================================================================

ANSWER:

A common pattern is:

    for {
        select {
        case value := <-ch:
            ...
        case <-done:
            return
        }
    }

This lets a goroutine process messages while also listening for
shutdown.
*/

/*
======================================================================
Q52. How can a done channel signal cancellation?
======================================================================

ANSWER:

A closed channel can broadcast a signal to all receivers.

Example:

    done := make(chan struct{})

    close(done)

Any goroutine doing:

    <-done

will immediately unblock.

This is a classic cancellation/signaling pattern.
*/

func q52() {
	fmt.Println("\nQ52")

	done := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		select {
		case <-done:
			fmt.Println("Worker stopped")
		}
	}()

	close(done)

	wg.Wait()
}

/*
======================================================================
Q53. Can a channel be used only for signaling?
======================================================================

ANSWER:

Yes.

A channel can carry an actual value:

    ch <- data

or simply signal an event:

    close(done)

For signaling, chan struct{} is common because struct{} carries no
meaningful data and uses zero-sized values.
*/

/*
======================================================================
Q54. Why is chan struct{} commonly used?
======================================================================

ANSWER:

Because struct{} has zero size and communicates an event without
allocating meaningful payload data.

Common:

    done := make(chan struct{})

Then:

    close(done)

to broadcast completion/cancellation.
*/

/*
======================================================================
Q55. What happens if select has multiple ready cases?
======================================================================

ANSWER:

If multiple cases are ready, Go chooses one pseudo-randomly.

Do not depend on a particular case always winning.
*/

func q55() {
	fmt.Println("\nQ55")

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	select {
	case value := <-ch1:
		fmt.Println("ch1:", value)
	case value := <-ch2:
		fmt.Println("ch2:", value)
	}
}

/*
======================================================================
Q56. What happens if select has only default?
======================================================================

ANSWER:

The default executes immediately.

This is useful for non-blocking channel operations.
*/

func q56() {
	fmt.Println("\nQ56")

	ch := make(chan int)

	select {
	case value := <-ch:
		fmt.Println(value)

	default:
		fmt.Println("No value; continue immediately")
	}
}

/*
======================================================================
Q57. What is a common mistake with closing channels in worker pools?
======================================================================

ANSWER:

Closing the jobs channel while producers are still sending causes:

    panic: send on closed channel

The owner of the sending lifecycle should close the channel only
after all sends are finished.

With multiple producers, use coordination such as a WaitGroup to
know when all producers have stopped before closing the channel.
*/

/*
======================================================================
Q58. How do you build multiple producers safely?
======================================================================

ANSWER:

Use a producer WaitGroup.

Pattern:

    producerWG.Add(numberOfProducers)

    each producer:
        defer producerWG.Done()
        send jobs

    separate goroutine:
        producerWG.Wait()
        close(jobs)

This ensures the channel is not closed while producers are sending.
*/

func q58() {
	fmt.Println("\nQ58")

	jobs := make(chan int, 10)

	var producerWG sync.WaitGroup

	producer := func(start int) {
		defer producerWG.Done()

		for i := start; i < start+3; i++ {
			jobs <- i
		}
	}

	producerWG.Add(2)

	go producer(1)
	go producer(10)

	go func() {
		producerWG.Wait()
		close(jobs)
	}()

	for job := range jobs {
		fmt.Println("Job:", job)
	}
}

/*
======================================================================
Q59. Interview coding problem:
Create a channel that produces numbers 1 to 10.
======================================================================

ANSWER:
Return a receive-only channel.

This demonstrates:

    goroutine
    channel
    channel closure
    range
    directional channel
*/

func q59Generate() <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := 1; i <= 10; i++ {
			out <- i
		}
	}()

	return out
}

func q59() {
	fmt.Println("\nQ59")

	for value := range q59Generate() {
		fmt.Println(value)
	}
}

/*
======================================================================
Q60. INTERVIEW MASTER:
Design a production-style channel worker system.
======================================================================

ANSWER:

A strong interview answer should cover:

1. Jobs channel
2. Bounded capacity
3. Fixed number of workers
4. Results/error handling
5. Context cancellation
6. Timeout/deadline
7. Clear channel ownership
8. Proper channel closing
9. WaitGroup for lifecycle
10. No goroutine leaks
11. Backpressure
12. Graceful shutdown
13. Race-free shared state
14. Monitoring/logging
15. Testing with race detector

Typical architecture:

                    +-------------+
                    | Producers   |
                    +------+------+
                           |
                           v
                    +-------------+
                    | jobs channel|
                    +------+------+
                           |
              +------------+------------+
              |            |            |
              v            v            v
           Worker 1     Worker 2     Worker 3
              |            |            |
              +------------+------------+
                           |
                           v
                    +-------------+
                    | results     |
                    +-------------+

Cancellation:

    select {
    case <-ctx.Done():
        return

    case job := <-jobs:
        process(job)
    }

Important interview statement:

"Channels should be used to define clear ownership and communication
between goroutines. Shared mutable state should still be protected
with Mutex or atomic operations where necessary."

======================================================================
CHANNEL INTERVIEW QUICK REVISION
======================================================================

1. Channel:
   Communication between goroutines.

2. Unbuffered:
   Capacity 0; sender and receiver synchronize.

3. Buffered:
   Has capacity; sender can proceed while buffer has space.

4. Send:
   ch <- value

5. Receive:
   value := <-ch

6. Close:
   close(ch)

7. Closed receive:
   value, ok := <-ch

8. Range:
   for value := range ch

9. Send-only:
   chan<- T

10. Receive-only:
    <-chan T

11. Capacity:
    cap(ch)

12. Buffered length:
    len(ch)

13. select:
    Multiple channel operations.

14. nil channel:
    Send/receive blocks forever.

15. Closed channel:
    Receive is immediately ready.

16. Send to closed:
    Panic.

17. Close twice:
    Panic.

18. Fan-out:
    One input -> many workers.

19. Fan-in:
    Many inputs -> one output.

20. Pipeline:
    Multiple stages connected by channels.

21. Backpressure:
    Producer slows when bounded buffer is full.

22. Semaphore:
    Channel limits concurrent operations.

23. done channel:
    Signaling/cancellation.

24. chan struct{}:
    Common zero-payload signal.

25. Worker pool:
    Fixed/bounded workers consuming jobs.

======================================================================
MOST IMPORTANT QUESTIONS FOR YOUR 2-YEAR GO INTERVIEW
======================================================================

MASTER THESE FIRST:

Q05  Unbuffered channel
Q06  Buffered channel
Q07  Buffered vs unbuffered
Q09  len vs cap
Q10  Blocking send
Q11  Blocking receive
Q13  Closing
Q14  Comma-ok
Q17  Who closes channel?
Q19  Range over channel
Q21  Send-only
Q22  Receive-only
Q25  Select
Q26  Timeout
Q27  Non-blocking receive
Q28  Non-blocking send
Q31  Nil channel
Q36  Deadlock
Q38  Goroutine leak
Q40  Worker pool
Q41  Fan-out
Q42  Fan-in
Q43  Pipeline
Q44  Semaphore
Q47  Backpressure
Q52  Done channel
Q55  Multiple ready select cases
Q57  Channel ownership
Q58  Multiple producers
Q59  Channel generator
Q60  Production worker system

======================================================================
COMMON INTERVIEW TRAPS
======================================================================

TRAP 1:
"Closing a channel deletes all its values."

FALSE.

Buffered values can still be received after close.

TRAP 2:
"Receiver should always close the channel."

FALSE.

The sender/owner normally closes it.

TRAP 3:
"Buffered channels never block."

FALSE.

They block when full on send and empty on receive.

TRAP 4:
"nil channel is the same as closed channel."

FALSE.

Nil:
    send/receive blocks forever.

Closed:
    receive succeeds immediately with zero value after draining.

TRAP 5:
"select guarantees the first case."

FALSE.

If multiple cases are ready, one is selected pseudo-randomly.

TRAP 6:
"Channels eliminate all race conditions."

FALSE.

Shared memory can still have races.

TRAP 7:
"Every channel should be closed."

FALSE.

Close when the receiver needs a signal that no more values will
arrive.

TRAP 8:
"close() is a synchronization replacement for everything."

FALSE.

Use the appropriate primitive: channels, Mutex, atomic, WaitGroup,
Context, etc.

======================================================================
PRACTICE COMMANDS
======================================================================

Run:

    go run channels_60.go

Format:

    gofmt -w channels_60.go

Race detector:

    go run -race channels_60.go

Build:

    go build channels_60.go

======================================================================
END — 60 GO CHANNEL QUESTIONS & ANSWERS
======================================================================
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
	q18()
	q19()
	q21()
	q22()
	q24()
	q25()
	q26()
	q27()
	q28()
	q31()
	q35()
	q40()
	q41()
	q42()
	q43()
	q44()
	q45()
	q46()
	q49()
	q52()
	q55()
	q56()
	q58()
	q59()

	fmt.Println("\n==============================================================")
	fmt.Println("60 GO CHANNEL QUESTIONS COMPLETED")
	fmt.Println("==============================================================")
}
