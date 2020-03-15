# segment

[![Build Status](https://travis-ci.org/opentogo/segment.svg?branch=master)](https://travis-ci.org/opentogo/segment)
[![GoDoc](https://godoc.org/github.com/opentogo/segment?status.png)](https://godoc.org/github.com/opentogo/segment)
[![codecov](https://codecov.io/gh/opentogo/segment/branch/master/graph/badge.svg)](https://codecov.io/gh/opentogo/segment)
[![Go Report Card](https://goreportcard.com/badge/github.com/opentogo/segment)](https://goreportcard.com/report/github.com/opentogo/segment)
[![Open Source Helpers](https://www.codetriage.com/opentogo/segment/badges/users.svg)](https://www.codetriage.com/opentogo/segment)

This package was inspired in Michel Martens' [seg](https://github.com/soveran/seg) Crystal library and it offer a segment matcher for paths in Go.

## Installation

```bash
go get github.com/opentogo/segment
```

## Usage

Consider this interactive session:

```go
package main

import (
	"fmt"

	"github.com/opentogo/segment"
)

func main() {
	seg := segment.NewSegment("/users/42", "")

	fmt.Printf("%#v\n", seg)
	// => segment.Segment{path:"/users/42", size:9, pos:0}

	fmt.Println(seg.Previous())
	// =>

	fmt.Println(seg.Current())
	// => /users/42

	fmt.Println(seg.Consume("users"))
	// => true

	fmt.Println(seg.Previous())
	// => /users

	fmt.Println(seg.Current())
	// => /42

	fmt.Println(seg.Consume("42"))
	// => true

	fmt.Println(seg.Previous())
	// => /users/42

	fmt.Println(seg.Current())
	// =>
}
```

As you can see, the command fails and the `prev` and `curr` strings are not altered. Now we'll see how to capture segment values:

```go
package main

import (
	"fmt"

	"github.com/opentogo/segment"
)

func main() {
	var (
		captures = map[string]string{}
		seg      = segment.NewSegment("/users/42", "")
	)

	fmt.Printf("%#v\n", seg)
	// => segment.Segment{path:"/users/42", size:9, pos:0}

	fmt.Println(seg.Previous())
	// =>

	fmt.Println(seg.Current())
	// => /users/42

	seg.Capture("foo", captures)

	fmt.Println(seg.Previous())
	// => /users

	fmt.Println(seg.Current())
	// => /42

	seg.Capture("bar", captures)

	fmt.Println(seg.Previous())
	// => /users/42

	fmt.Println(seg.Current())
	// =>

	fmt.Println(captures)
	// => map[bar:42 foo:users]
}
```

It is also possible to `extract` the next segment from the path. The method `extract` returns the next segment, if available, or nil otherwise:

```go
package main

import (
	"fmt"

	"github.com/opentogo/segment"
)

func main() {
	seg := segment.NewSegment("/users/42", "")

	fmt.Printf("%#v\n", seg)
	// => segment.Segment{path:"/users/42", size:9, pos:0}

	fmt.Println(seg.Previous())
	// =>

	fmt.Println(seg.Current())
	// => /users/42

	fmt.Println(seg.Extract())
	// => users

	fmt.Println(seg.Previous())
	// => /users

	fmt.Println(seg.Current())
	// => /42

	fmt.Println(seg.Extract())
	// => 42

	fmt.Println(seg.Previous())
	// => /users/42

	fmt.Println(seg.Current())
	// =>
}
```

You can also go back by using the methods `retract` and `restore`, which are the antidote to `extract` and `consume` respectively.

Let's see how `retract` works:

```go
package main

import (
	"fmt"

	"github.com/opentogo/segment"
)

func main() {
	seg := segment.NewSegment("/users/42", "")

	fmt.Printf("%#v\n", seg)
	// => segment.Segment{path:"/users/42", size:9, pos:0}

	fmt.Println(seg.Previous())
	// =>

	fmt.Println(seg.Current())
	// => /users/42

	fmt.Println(seg.Extract())
	// => users

	fmt.Println(seg.Previous())
	// => /users

	fmt.Println(seg.Current())
	// => /42

	fmt.Println(seg.Retract())
	// => users

	fmt.Println(seg.Previous())
	// =>

	fmt.Println(seg.Current())
	// => /users/42
}
```

And now `restore`:

```go
package main

import (
	"fmt"

	"github.com/opentogo/segment"
)

func main() {
	seg := segment.NewSegment("/users/42", "")

	fmt.Printf("%#v\n", seg)
	// => segment.Segment{path:"/users/42", size:9, pos:0}

	fmt.Println(seg.Previous())
	// =>

	fmt.Println(seg.Current())
	// => /users/42

	fmt.Println(seg.Extract())
	// => users

	fmt.Println(seg.Previous())
	// => /users

	fmt.Println(seg.Current())
	// => /42

	fmt.Println(seg.Restore("foo"))
	// => false

	fmt.Println(seg.Restore("users"))
	// => true

	fmt.Println(seg.Previous())
	// =>

	fmt.Println(seg.Current())
	// => /users/42
}
```

## Contributors

- [rogeriozambon](https://github.com/rogeriozambon) Rogério Zambon - creator, maintainer
