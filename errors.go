package snowflake

import "fmt"

// Base error struct for all other snowflake errors. Implements error interface.
type SnowflakeError struct {
	message string
	err     error
}

// Required to implement error interface. Returns formatted error.
func (s *SnowflakeError) Error() string {
	return fmt.Sprintf("%s (original error: %s)", s.message, s.err.Error())
}

// This function returns original error.
func (s *SnowflakeError) OriginalError() error {
	return s.err
}

// Used in:
//
//	snowflake.ParseString() // when strconv.ParseUint unable to parse string as uint64.
type StringParseError struct{ SnowflakeError }

// Used in:
//
//	snowflake.ParseJSON() // when JSON is a unquoted integer and unquoted integers are
//	                      // not allowed.
type UnquotedIntegerError struct{ SnowflakeError }
