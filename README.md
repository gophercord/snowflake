<h1>Moved to <a href="https://github.com/gorocord/snowflake">gorocord/snowflake</a></h1>

<div align="center">
<h1>Gophercord/Snowflake</h1>
<p>Library for manipulating Discord snowflake IDs written in Go (Golang). Used by Gophercord.</p>
<img width="169.7" height="150" style="padding-bottom: 20px;" src=".etc/pictures/gopher/gopher-with-discord-logo.png">

[![Go Reference](https://pkg.go.dev/badge/github.com/gophercord/snowflake.svg)](https://pkg.go.dev/github.com/gophercord/snowflake)
[![Go Report](https://goreportcard.com/badge/github.com/gophercord/snowflake)](https://goreportcard.com/report/github.com/gophercord/snowflake)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/gophercord/snowflake/blob/master/LICENSE)
[![Code examples](https://img.shields.io/badge/Code%20examples-grey.svg)](.examples)

</div>

---

1. [About](#about)
    1. [What is snowflake](#what-is-snowflake)
    2. [Snowflake structure](#snowflake-structure)
2. [Getting started](#getting-started)
    1. [Installing snowflake](#installing-snowflake)
    2. [Code example](#code-example)
3. [License](#license)

---

## About
[![Reference](https://img.shields.io/badge/Discord%20Developers-Reference-blue.svg?logo=discord)](https://discord.com/developers/docs/reference#snowflakes)
[![Wikipedia](https://img.shields.io/badge/Wikipedia-Snowflake%20ID-blue.svg?logo=wikipedia)](https://en.wikipedia.org/wiki/Snowflake_ID)

### What is snowflake
Snowflake is a unique identifier format used by Discord, Twitter (now X) and other platforms. This library provides tools for parsing Discord snowflake IDs.

### Snowflake structure
Snowflake is a 64-bit integer without sign (in Go, this is a uint64 type). Snowflake bits are separated into groups:
```
 [000000100111000100000110010110101100000100][00001][00000][000010011001]
64                                          22     17     12             0
```
Where:
1. Bits 0-12 is a sequence (incremented for every generated ID on process);
2. Bits 12-17 is a internal process ID;
3. Bits 17-22 is a internal worker ID;
4. Bits 22-64 is a number of milliseconds since Discord epoch.

## Getting started
### Installing snowflake
Type this command in your terminal to install:
```bash
$ go get github.com/gophercord/snowflake
```

### Code example
[![Playground](https://img.shields.io/badge/Try%20this%20code%20on-Go%20Playground-blue.svg)](https://go.dev/play/p/HsJ461Pf7Pn)

Simple gophercord/snowflake usage example. You may like to see the [.examples](.examples) directory for more examples.

```golang
package main

import (
	"fmt"

	"github.com/gophercord/snowflake"
)

func main() {
	var err error

	// Creating new snowflake from uint64
	s := snowflake.Snowflake(1363292549053284505)

	// Accessing snowflake attributes
	fmt.Println(
		"Created at:", s.Time(),
		"\n  Seconds:", s.Unix(),
		"\n  Milliseconds:", s.UnixMilli(),
		"\nWorker ID:", s.WorkerID(),
		"\nProcess ID:", s.ProcessID(),
		"\nSequence:", s.Sequence(),
		"\n==============================",
	)

	// You can parse a snowflake from a string, JSON, or a [time.Time]

	// Parsing new snowflake ID from a string
	//
	// NOTE: The string must contain only digits without any signs (because Snowflake is a
	// uint64 type)
	s2, _ := snowflake.ParseString("10")
	fmt.Println("parsed from string:", s2)

	// Parsing new snowflake ID from JSON
	s3, _ := snowflake.ParseJSON([]byte(`"134"`))
	s4, _ := snowflake.ParseJSON([]byte("134")) // unquoted integer

	fmt.Println("parsed from JSON:", s3, s4)

	// You can deny unquoted integers in JSON (by default, unquoted integers are allowed)
	snowflake.AllowUnquoted = false

	_, err = snowflake.ParseJSON([]byte("42"))
	fmt.Println("JSON parse error:", err)
	// The error is not nil, because unquoted integers are not allowed

	// Allow unquoted integers
	snowflake.AllowUnquoted = true

	_, err = snowflake.ParseJSON([]byte("42"))
	fmt.Println("JSON parse error:", err)
	// The error is nil because unquoted integers are now allowed
}
```

## License
This software is licensed under the MIT License. For more information, see [LICENSE](./LICENSE.md).
