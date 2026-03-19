# Web IDL for Go

## Installation

```sh
go get go.jcbhmr.com/webidl
```

## Usage

```go
func myWebAlgorithm(cause webidl.DOMException) webidl.DOMException {
    // The my web algorithm steps given DOMException cause are:

    // 1. Store cause's message as message
    message := cause.Message()

    // 2. Create a new DOMException exception from the concatenation of "cause: " and message.
    exception := webidl.NewDOMException("cause: " + message)

    // 3. Return exception
    return exception
}
```
