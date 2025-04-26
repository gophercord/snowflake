// Library for manipulating Discord snowflake IDs written in Go (Golang), used by Gophercord.
//
// See docs for [Snowflake] to see more information.
//
//   - [Snowflake reference] on Discord developer portal
//   - [Read more about snowflakes] on Wikipedia
//
// [Snowflake reference]: https://discord.com/developers/docs/reference#snowflakes
// [Read more about snowflakes]: https://en.wikipedia.org/wiki/Snowflake_ID
package snowflake

import (
	"bytes"
	"strconv"
	"time"
)

var (
	// Modifying these variables is NOT recommended.

	JSON_NULL        = []byte("null") // JSON null. Example: {"key": null}
	JSON_ZERO        = []byte("0")    // Zero without quotes (JSON integer). Example: {"key": 0}
	JSON_ZERO_QUOTED = []byte(`"0"`)  // Zero with quotes (JSON string). Example: {"key": "0"}
)

var (
	// If true, unquoted integers will be allowed in [ParseJSON] when parsing snowflake
	// from JSON.
	//
	//  snowflake.AllowUnquoted = true
	//  snowflake.ParseJSON("10")   // OK
	//  snowflake.ParseJSON(`"10"`) // OK
	//
	//  snowflake.AllowUnquoted = false
	//  snowflake.ParseJSON("10")   // ERROR because "10" is unquoted and unquoted
	//                              // integers are not allowed.
	//  snowflake.ParseJSON(`"10"`) // OK
	//
	// By default unquoted integers are allowed.
	AllowUnquoted = true

	// A Unix timestamp in milliseconds, represents the Discord epoch date and time.
	//
	//	// You can change epoch if needed
	//	snowflake.Epoch = 12345
	Epoch uint64 = 1420070400000
)

type (
	// Snowflake value. To get uint64 use:
	//
	//	v := uint64(snowflake)
	//	v := snowflake.Value() // not recommended, use first example
	//
	// If you are using Gophercord, use "ID" attribute instead of "SID".
	//
	// Snowflake bits are separated into groups:
	//
	//	 [000000100111000100000110010110101100000100][00001][00000][000010011001]
	//	64                                          22     17     12             0
	//
	// Where:
	//  1. Bits 0-12 is a sequence (incremented for every generated ID on process);
	//  2. Bits 12-17 is a internal process ID;
	//  3. Bits 17-22 is a internal worker ID;
	//  4. Bits 22-64 is a number of milliseconds since Discord epoch.
	Snowflake uint64
	Bit       bool    // One bit as a bool (where 1 is true and 0 is false).
	Bitmap    [64]Bit // List with length 64 of snowflake ID bits.
)

// # Method Unix() of Snowflake
//
// Returns snowflake creation date and time as milliseconds in unix format.
//
// Calculation formula:
//
//	(snowflake>>22) + Epoch
//
// # Return
//
//   - uint64: Unix timestamp in milliseconds.
//
// (No arguments, errors, and examples)
func (s Snowflake) UnixMilli() uint64 {
	return uint64(s>>22) + Epoch
}

// # Method Unix() of Snowflake
//
// Returns snowflake creation date and time as seconds in unix format.
//
// Calculation formula:
//
//	((snowflake>>22) + Epoch) / 1_000
//
// # Return
//
//   - uint64: Unix timestamp in seconds.
//
// (No arguments, errors, and examples)
func (s Snowflake) Unix() uint64 {
	return (uint64(s>>22) + Epoch) / 1_000
}

// # Method Time() of Snowflake
//
// Returns snowflake creation date and time.
//
// # Return
//
//   - [time.Time]: Snowflake creation date and time.
//
// # Examples
//
//	s := snowflake.Snowflake(1363292549053284505)
//	fmt.Println(s.Time().Year())       // 2025
//	fmt.Println(s.Time().Month())      // April
//	fmt.Println(int(s.Time().Month())) // 4
//
// (No arguments and errors)
func (s Snowflake) Time() time.Time {
	return time.UnixMilli(int64(s>>22) + int64(Epoch))
}

// # Method WorkerID() of Snowflake
//
// Returns snowflake internal worker ID.
//
// Calculation formula:
//
//	(snowflake & 0x3E0000) >> 17
//
// # Return
//
//   - uint8: Internal worker ID.
//
// # Examples
//
//	s := snowflake.Snowflake(1363292549053284505)
//	wid := s.ProcessID()
//	myWid := (s.Value() & 0x3E0000) >> 17
//	fmt.Println(wid == myWid) // true
//
// (No arguments and errors)
func (s Snowflake) WorkerID() uint8 {
	return uint8((s & 0x3E0000) >> 17)
}

// # Method ProcessID() of Snowflake
//
// Returns snowflake internal process ID.
//
// Calculation formula:
//
//	(snowflake & 0x1F000) >> 12
//
// # Return
//
//   - uint8: Internal process ID.
//
// # Examples
//
//	s := snowflake.Snowflake(1363292549053284505)
//	pid := s.ProcessID()
//	myPid := (s.Value() & 0x1F000) >> 12
//	fmt.Println(pid == myPid) // true
//
// (No arguments and errors)
func (s Snowflake) ProcessID() uint8 {
	return uint8((s & 0x1F000) >> 12)
}

// # Method Sequence() of Snowflake
//
// Returns snowflake sequence (incremented for every generated ID on process).
//
// Calculation formula:
//
//	snowflake & 0xFFF
//
// # Return
//
//   - uint16: Sequence number.
//
// # Examples
//
//	s := snowflake.Snowflake(1363292549053284505)
//	seq := s.Sequence()
//	mySeq := s.Value() & 0xFFF
//	fmt.Println(seq == mySeq) // true
//
// (No arguments and errors)
func (s Snowflake) Sequence() uint16 {
	return uint16(s & 0xFFF)
}

// # Method String() of Snowflake
//
// Returns snowflake ID converted to string.
//
// # Return
//
//   - string: Snowflake ID as string.
//
// # Examples
//
//	s := snowflake.Snowflake(1363292549053284505)
//	myString := s.String()
//	display := "Snowflake value: " + myString
//	fmt.Println(display)
//
// (No arguments and errors)
func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

// # Method Value() of Snowflake
//
// Returns snowflake ID converted to uint64.
//
// # Return
//
//   - uint64: Snowflake ID as uint64.
//
// (No errors, arguments, and examples)
func (s Snowflake) Value() uint64 {
	return uint64(s)
}

// # Method Bit(i) of Snowflake
//
// Returns a single bit from a snowflake, with the indexing starting from the right bit.
//
//	  64                                          22     17     12             0
//	   [000000100111000100000110010110101100000100][00001][00000][000010011001]
//	Last <----                                                           <--- First
//
// NOTE: The last bit index is 63 and not 64, as indexing starts from 0.
//
// # Arguments
//
//   - i uint8: Index of the bit. The index of the first bit is 0, and the last is 63.
//
// # Return
//
//   - Bit (bool): Bit value.
//
// # Examples
//
//	s := snowflake.Snowflake(0b0000001001/*and more*/0010011001)
//	fmt.Println(s.Bit(0))  // true
//	fmt.Println(s.Bit(1))  // false
//	fmt.Println(s.Bit(63)) // false
//	fmt.Println(s.Bit(64)) // ALWAYS false because index 64 is out of range
//
// (No errors)
func (s Snowflake) Bit(i uint8) Bit {
	return i < 64 && s&(1<<i) != 0
}

// # Method Bit(i) of Snowflake
//
// Returns a list of bits from snowflake ID.
//
// # Return
//
//   - Bitmap ([]Bit): List of bits.
//
// # Examples
//
//	// (in this example we use T as alias to true and F as alias to false)
//
//	s := snowflake.Snowflake(0b0000001001/*and more*/0010011001)
//	fmt.Println(s.Bitmap())
//	// {F, F, F, F, F, F, T, F, F, T, /*and more*/ F, F, T, F, F, T, T, F, F, T}
//
// (No arguments and errors)
func (s Snowflake) Bitmap() Bitmap {
	bits := Bitmap{}
	for i := range bits {
		bits[i] = s&(1<<i) != 0
	}
	return bits
}

// # Method MarshalJSON() of Snowflake
//
// Returns quoted snowflake ID value in JSON format.
//
// # Return
//
//   - []byte: Snowflake ID value in JSON-format encoded into bytes.
//   - error (always nil, but needed to implement interface. So, you can ignore the error value).
//
// # Examples
//
//	s := snowflake.Snowflake(1363292549053284505)
//	b, _ := s.MarshalJSON()
//
// (No arguments and errors)
func (s Snowflake) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(strconv.FormatUint(uint64(s), 10))), nil
}

// # Method UnmarshalJSON(b) of Snowflake
//
// Parses JSON with [ParseJSON] and changes the CURRENT snowflake ID value. Does NOT return a new
// snowflake ID.
//
// # Arguments
//
//   - v []byte: JSON-formatted string in bytes.
//
// # Errors
//
//   - [UnquotedIntegerError]: If the integer is not quoted and [AllowUnquoted] is false.
//   - [StringParseError]: If the string contains non-integer characters ([strconv.ParseUint]
//     returned an error when parsing the string).
//
// # Examples
//
//	s := snowflake.New()
//	s.UnmarshalJSON([]byte("1363292549053284505"))
//	fmt.Println(s) // 1363292549053284505
//
// (No return)
func (s *Snowflake) UnmarshalJSON(b []byte) error {
	snowflake, err := ParseJSON(b)
	if err != nil {
		return err
	}
	*s = snowflake
	return nil
}

// # Function ParseString(s)
//
// Parses a new snowflake from a string in integer format.
//
// # Arguments
//
//   - s string: The string contains only integer characters without sign, because snowflake
//     is uint64.
//
// # Return
//
//   - [Snowflake]: New snowflake parsed from argument "s".
//   - error
//
// # Errors
//
//   - [StringParseError]: If the string contains non-integer characters ([strconv.ParseUint]
//     returned an error when parsing the string).
//
// # Examples
//
//	s, _ := snowflake.ParseString("1363292549053284505")  // OK
//	s, _ := snowflake.ParseString("-1363292549053284505")
//	// ERROR: Snowflake can't be a negative integer.
//	s, _ := snowflake.ParseString("1234abcdef")
//	// ERROR: ParseString accepts only integers with base 10.
//	s, _ := snowflake.ParseString("not integer")
//	// ERROR: String contains non-integer characters.
func ParseString(s string) (Snowflake, error) {
	snowflake, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, &StringParseError{SnowflakeError: SnowflakeError{
			message: "unable to parse string as integer",
			err:     err,
		}}
	}
	return Snowflake(snowflake), nil
}

// # Wrapper for ParseString(s)
//
// Wrapper for [ParseString] function. Creates panic if [ParseString] returns an error.
func MustParseString(s string) Snowflake {
	snowflake, err := ParseString(s)
	if err != nil {
		panic(err)
	}
	return snowflake
}

// # Function ParseTime(t)
//
// Creates a new snowflake ID based on a [time.Time] with zero worker ID and process ID, and a
// sequence.
//
// # Arguments
//
//   - t [time.Time]: Time from which to parse a new snowflake ID
//
// # Return
//
//   - [Snowflake]: New snowflake parsed from argument "t".
//
// # Examples
//
//	s := snowflake.ParseTime(time.UnixMilli(0b1111))
//	fmt.Printf("%b", uint64(s))
//	//                   11110000000000000000000000
//	//                               |
//	//                               v
//	// [000000000000000000000000000000000000001111]0000000000000000000000
//	//             Snowflake time part                  Other parts
//
//	// NOTE: Other parts of the bits are always set to zero in Snowflake IDs created with
//	// ParseTime().
//
// (No errors)
func ParseTime(t time.Time) Snowflake {
	return Snowflake((t.UnixMilli() - int64(Epoch)) << 22)
}

// # Function ParseJSON(b)
//
// Parses a new snowflake from a JSON-formatted string (must be encoded as bytes). Can be an
// integer (if AllowUnquoted is true) or a quoted integer.
//
// # Arguments
//
//   - v []byte: JSON-formatted string in bytes.
//
// # Return
//
//   - [Snowflake]: New snowflake parsed from argument "v".
//   - error
//
// # Errors
//
//   - [UnquotedIntegerError]: If the integer is not quoted and [AllowUnquoted] is false.
//   - [StringParseError]: If the string contains non-integer characters ([strconv.ParseUint]
//     returned an error when parsing the string).
//
// # Examples
//
//	snowflake.AllowUnquoted = true
//	snowflake.ParseJSON("10")   // OK
//	snowflake.ParseJSON(`"10"`) // OK
//
//	snowflake.AllowUnquoted = false
//	snowflake.ParseJSON("10")   // ERROR because "10" is unquoted and unquoted
//	                            // integers are not allowed.
//	snowflake.ParseJSON(`"10"`) // OK
func ParseJSON(b []byte) (Snowflake, error) {
	if bytes.Equal(b, JSON_NULL) ||
		bytes.Equal(b, JSON_ZERO) ||
		bytes.Equal(b, JSON_ZERO_QUOTED) {
		return 0, nil
	}

	s := string(b)
	snowflake, err := strconv.Unquote(s)
	if err != nil {
		if !AllowUnquoted {
			return 0, &UnquotedIntegerError{SnowflakeError: SnowflakeError{
				message: "unquoted integer but unquoted integers are not allowed",
				err:     err,
			}}
		}
		snowflake = s
	}

	return ParseString(snowflake)
}

// # Wrapper for ParseJSON(b)
//
// Wrapper for [ParseJSON] function. Creates panic if [ParseJSON] returns an error.
func MustParseJSON(b []byte) Snowflake {
	snowflake, err := ParseJSON(b)
	if err != nil {
		panic(err)
	}
	return snowflake
}

// # Function Parse(v)
//
// Parses a new snowflake from string, uint64 or time.Time.
//
// Important: The use of this function is not recommended due to the use of type assertion.
// Instead, use [ParseString] and [ParseTime].
//
// # Arguments
//
//   - v string|uint64|time.Time: The value to convert into a snowflake ID.
//
// # Return
//
//   - [Snowflake]: New snowflake parsed from argument "v".
//   - error
//
// # Errors
//
//   - [StringParseError]: if the type of argument "v" is a string and the string contains
//     non-integer characters ([strconv.ParseUint] returned an error when parsing the string).
//
// # Examples:
//
//	myTime, myString, myUint := time.Now(), "1363292549053284505", 1363292549053284505
//
//	s, _ := snowflake.Parse(myTime)   // alternative to ParseTime(myTime)
//	s, _ := snowflake.Parse(myString) // alternative to ParseString(myString)
//	s, _ := snowflake.Parse(myUint)   // alternative to Snowflake(myUint)
func Parse[T string | uint64 | time.Time](v T) (Snowflake, error) {
	switch t := any(v).(type) {
	case string:
		return ParseString(t)
	case uint64:
		return Snowflake(t), nil
	case time.Time:
		return ParseTime(t), nil
	}

	return 0, nil
}

// # Wrapper for Parse(v)
//
// Wrapper for [Parse] function. Creates panic if [Parse] returns an error.
func MustParse[T string | uint64 | time.Time](v T) Snowflake {
	snowflake, err := Parse(v)
	if err != nil {
		panic(err)
	}
	return snowflake
}

// # Function New()
//
// Creates a new snowflake ID with all bits set to zero. You can also use Snowflake(0).
//
// # Return
//
//   - [Snowflake]: New snowflake with all bits set to zero.
//
// # Examples
//
//	s := snowflake.New()
//	// useful when using with UnmarshalJSON
//	s.UnmarshalJSON([]byte(`"10"`))
//
//	// New() returns empty snowflake
//	fmt.Println(s)             // 0
//	fmt.Println(s.UnixMilli()) // 0
//	fmt.Println(s.WorkerID())  // 0
//	fmt.Println(s.Sequence())  // 0
//
// (No arguments and errors)
func New() Snowflake {
	return Snowflake(0)
}
