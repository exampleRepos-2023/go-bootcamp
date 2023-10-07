# Concurrency

## CPU, Core, Thread, Scheduler, Process
- Core is a processing unit in a CPU that execute the instructions. One CPU can have multiple CPU cores.
- Thread is a sequence of instructions that can be executed independently.
- Each CPU core can execute one thread at a time, but multiple threads can be assigned to a CPU core.
- Scheduler decides which thread gets executed and when. Even with one CPU core, the scheduler still allows for context switching, which has an appearance of multiple threads running at the same time.
- Process is a running program (often used to describe an independently running task as well).

A process is a running program whose instructions can be grouped into multiple threads; those threads can be scheduled to run independently on multiple CPU cores.

## Concurrency vs Parallelism
- Concurrency is DEALING with many things at once; parallelism is DOING many things at once.
- Concurrent programs don't have to be parallel; parallel programs likely have a concurrency model.

