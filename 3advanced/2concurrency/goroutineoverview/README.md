# Goroutine

## What is Goroutine?
- Function of instructions executed independently (sounds similar to thread?)
- Uses `go` keyword followed by a function call `f()`

## Goroutine vs Thread
- Architecture
  ```
  Thread                   vs                  Goroutine
  Scheudled by kernal (OS)                     Scheduled by Go runtime, talking to kernal
  Hardware dependent                           Hardware independent, abstracted by Go runtime
  ```
- Goroutine Performance
  - Fast startup time
  - Fast context switch
  - Less memory consumption

