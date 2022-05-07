# dontpanic

Panic-free alternatives to common Go operations.

## Installation

`go get github.com/phelmkamp/dontpanic`

## Usage

```go
func f() (err error) {
    defer dontpanic.Recover(&err)
    // do stuff that might panic...
}
```
