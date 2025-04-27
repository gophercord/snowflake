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
