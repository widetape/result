# Result

Result is a Go library that only provides a simple `Result` object, that
is intended to replace conventional value-error pair returning.

## Motivation

It is too easy to bypass error checking, which can lead to a nil reference panic
or even worse to an undefined behaviour when the result of a function
is returned by value, rather than a reference.

## Usage

Conceptually, there are two conditions of a result: fake and real.

A fake result contains no value, and has an error. If your try to
unwrap a fake result via the `Value` method, you will get a runtime
panic. To check for an error you should first check the result via
`Error` method, and see if it returns an error or a `nil`.

On the other hand a real result does contain a value, and has no error
and can be safely unwrapped via the `Value` method.

> [!WARNING]
> Making a result directly via struct construction (i.e `result.Result[any]{}`)
> is unsupported and leads to undefined behaviour.

```go
package main

import (
    "math/rand"

    "github.com/widetape/result"
)

func realResult() result.Result[int] {
    return result.Real(rand.Int())
}

func fakeResult() result.Result[int] {
    return result.Fake(errors.New("no int this time"))
}

func main() {
    real := realResult()
    if err := real.Error(); err != nil {
        panic(err) // <- this won't happen
    }
    println(real.Value()) // this will output a random integer.

    fake := fakeResult()
    println(fake.Value()) // panic! one cannot simply use a fake result's non-existent value.
    if err := fake.Error(); err != nil {
        println(err.Error()) // "no int this time"
    }
}
```
