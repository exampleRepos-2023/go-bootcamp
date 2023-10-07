# Testing

## What is testing?
Testing is the process of evaluating and verifying the program works as expected.

## Why testing?
- Find bugs you may not know
- Prevent bugs to be introduced

[//]: # (## Types of testing)

[//]: # (- Unit testing)

[//]: # (- Integration testing)

[//]: # (- Functional testing)

[//]: # (- Performance testing)

[//]: # (- ...)

[//]: # ()
[//]: # (## Test coverage)

[//]: # (You are expected to write tests for exported components,)

[//]: # (and additional test cases for additional conditions)

[//]: # ()
[//]: # (`$ go test` - run tests)

[//]: # ()
[//]: # (`$ go test -cover` - run tests + show test coverage percentage)

[//]: # (```)

[//]: # (func NeedBeCovered&#40;v bool&#41; string {)

[//]: # (    if v {)

[//]: # (        return "true")

[//]: # (    })

[//]: # (    return "false")

[//]: # (})

[//]: # ()
[//]: # (func needNotBeCovered&#40;&#41; {)

[//]: # (    // ...)

[//]: # (})

[//]: # (```)
