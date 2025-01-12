## Exercise 1 - Theory questions

### Concepts

What is the difference between _concurrency_ and _parallelism_?

> \*In parallelism several tasks run simultaneously, typically on multiple cores or processors
>
> Concurrency is the ability of a system to handle multiple task at the same time, but they might not be executed simultaneously. They might be interleaved

What is the difference between a _race condition_ and a _data race_?

> _Data race is when multiple threads access the same shared memory location, and there is no synchronization to control acess to the shared memory_
>
> _A race condition is when the result depends on ordering in time_

_Very_ roughly - what does a _scheduler_ do, and how does it do it?

> _¨The scheduler decides what thread to run next_

### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?

> It is faster, less queue task time

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?

> _Fibers are used to divide the thread into several tasks, so that independent tasks can do their own work, wiothout using functions. It also is ismilar to blocking (which usually requires the use of more threads), and you do not need to use casllbacks or async/await (why?)_

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?

> _Both. It might make a program faster, but requires careful consideration conserning synchronization._

What do you think is best - _shared variables_ or _message passing_?

> Depends on the task. Shared variables is more efficient for tasks with lots of shared data, but isd more complex and mighjt be hard to scale.
>
> Message passing is less efficient, but is less complex and scales well.
