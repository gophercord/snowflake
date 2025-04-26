package snowflake_test

import (
	"snowflake"
	"testing"
)

const T, F = true, false

var exampleSnowflake = snowflake.Snowflake(175928847299117209)

type testBit struct {
	Index    uint8
	WantsBit snowflake.Bit
}

var testsBit = []testBit{
	// normal
	{0, T}, {1, F}, {2, F}, {3, T}, {4, T}, {5, F}, {6, F}, {63, F}, {62, F}, {61, F}, {57, T},
	// out of range
	{70, F}, {65, F}, {100, F}, {200, F}, {255, F},
}

func TestBit(t *testing.T) {
	for i, test := range testsBit {
		result := exampleSnowflake.Bit(test.Index)

		if result != test.WantsBit {
			t.Errorf("TestBit[%d] snowflake.Snowflake<%b>.Bit(%d) wanted %v, got %v",
				i, exampleSnowflake, test.Index, test.WantsBit, result)
		}
	}
}
