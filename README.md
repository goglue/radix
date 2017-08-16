# radix
[![Build Status](https://travis-ci.org/goglue/radix.svg?branch=master)](https://travis-ci.org/goglue/radix)
[![GoDoc](https://godoc.org/github.com/goglue/radix?status.svg)](https://godoc.org/github.com/goglue/radix)
[![Coverage Status](https://coveralls.io/repos/github/goglue/radix/badge.svg?branch=master)](https://coveralls.io/github/goglue/radix?branch=master)

## ** Under development **

### Abstraction

This library implements Radix Tree in go, the usages for this library can vary as it accepts interfaces as node item

### Installation
```bash
$ go get github.com/goglue/radix
```

### How to use
```go

import "github.com/goglue/radix"

func main() {
    tree := radix.NewTree()
    tree.Add("someStringConsideredAsPath", 101)
    tree.Add("someStringConsideredAsPath1", 102)
    tree.Add("someStringConsideredAsPath2", 103)
    
    value, err := tree.Get("someStringConsideredAsPath2")
    if nil != err {
        // check error types at the end of the document
    }
    
    val := value.(int)
    println(val) // will output 102
}
```

### Error types

- _ErrNodeNotFound_: Returned when passing a path and it could not be found
- _ErrDuplicateNode_: Returned when trying to overwrite a path value
- _ErrNodeLabel_: Returned when a label is not passed for the node
- _ErrNodeValue_: Returned when trying to define a path with a nil value

### What is missing

- Delete
- Replace/Update
