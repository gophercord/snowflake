package snowflake_test

import (
	"testing"

	"github.com/gophercord/snowflake"
)

const (
	T       = true
	F       = false
	example = snowflake.Snowflake(175928847299117209)
)

func TestBit(t *testing.T) {
	tests := []struct {
		Index    uint8
		WantsBit snowflake.Bit
	}{
		// normal
		{0, T}, {1, F}, {2, F}, {3, T}, {4, T}, {5, F}, {6, F}, {63, F}, {62, F}, {61, F}, {57, T},
		// out of range
		{70, F}, {65, F}, {100, F}, {200, F}, {255, F},
	}

	for i, test := range tests {
		result := example.Bit(test.Index)

		if result != test.WantsBit {
			t.Errorf("FAIL TestBit[%d]: snowflake.Snowflake<%b>.Bit(%d) wanted %v, got %v",
				i, example, test.Index, test.WantsBit, result)
		}
	}
}

func TestBitmap(t *testing.T) {
	type want struct {
		Index uint8
		Bit   snowflake.Bit
	}

	tests := []struct {
		Snowflake snowflake.Snowflake
		Wants     []want
	}{
		{
			example,
			[]want{
				{0, T}, {1, F}, {2, F}, {3, T},
				{63, F}, {62, F}, {61, F}, {60, F},
				{57, T}, {56, F}, {28, F}, {27, F},
				{24, T},
			},
		},
		{
			1363292549053284505,
			[]want{
				{0, T}, {1, F}, {2, F}, {3, T},
				{63, F}, {62, F}, {61, F}, {60, T},
				{57, T}, {56, F}, {28, F}, {27, T},
				{24, T},
			},
		},
	}

	for i, test := range tests {
		bitmap := test.Snowflake.Bitmap()

		for _, want := range test.Wants {
			bit := bitmap[want.Index]

			if bit != want.Bit {
				t.Errorf("FAIL TestBitmap[%d]: snowflake.Snowflake<%b>.Bitmap()[%d] wanted %v, "+
					"got %v (bitmap=%v)",
					i, test.Snowflake, want.Index, want.Bit, bit, bitmap)
			}
		}
	}
}

func TestParseString(t *testing.T) {
	tests := []struct {
		Input    string
		WantsErr bool
	}{
		{"1000", false},
		{"175928847299117209", false},
		{"-1", true},
		{"1000000000000000000000000", true},
		{"abcdef", true},
	}

	for i, test := range tests {
		_, err := snowflake.ParseString(test.Input)

		if err == nil && test.WantsErr {
			t.Errorf("FAIL TestParseJSON[%d]: string<%v> wanted error!=nil but "+
				"error IS nil",
				i, test.Input)
		} else if err != nil && !test.WantsErr {
			t.Errorf("FAIL TestParseJSON[%d]: string<%v> wanted error=nil but "+
				"error is NOT nil",
				i, test.Input)
		}
	}
}

func TestParseJSON(t *testing.T) {
	tests := []struct {
		Input         []byte
		AllowUnquoted bool
		WantsErr      bool
	}{
		{[]byte("10"), false, true},
		{[]byte(`"10"`), false, false},
		{[]byte("175928847299117209"), false, true},
		{[]byte("175928847299117209"), true, false},
		{[]byte(`"175928847299117209"`), false, false},
		{[]byte("1000000000000000000000000"), false, true},
		{[]byte("1000000000000000000000000"), true, true},
		{[]byte(`"1000000000000000000000000"`), false, true},
		{[]byte("-1"), false, true},
		{[]byte("-1"), true, true},
		{[]byte(`"-1"`), false, true},
		{[]byte(`"0"`), false, false},
		{[]byte("0"), false, false},
		{[]byte("null"), true, false},
		{[]byte("null"), false, false},
		{[]byte("integer"), true, true},
		{[]byte("."), true, true},
		{[]byte("3.14"), true, true},
	}

	for i, test := range tests {
		snowflake.AllowUnquoted = test.AllowUnquoted
		_, err := snowflake.ParseJSON(test.Input)

		if err == nil && test.WantsErr {
			t.Errorf("FAIL TestParseJSON[%d AllowUnquoted=%v]: []byte<%v> wanted error!=nil but "+
				"error IS nil",
				i, test.AllowUnquoted, test.Input)
		} else if err != nil && !test.WantsErr {
			t.Errorf("FAIL TestParseJSON[%d AllowUnquoted=%v]: []byte<%v> wanted error=nil but "+
				"error is NOT nil",
				i, test.AllowUnquoted, test.Input)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	s := snowflake.New()
	s.UnmarshalJSON([]byte("175928847299117209"))

	if s != 175928847299117209 {
		t.Errorf("FAIL TestUnmarshalJSON: invalid snowflake value after JSON unmarshal")
	}

	err := s.UnmarshalJSON([]byte("1000000000000000000000000"))

	if err == nil {
		t.Errorf("FAIL TestUnmarshalJSON: snowflake second JSON unmarshal must return error")
	}
	if s != 175928847299117209 {
		t.Errorf("FAIL TestUnmarshalJSON: snowflake value must remain unchanged after the " +
			"second JSON unmarshal")
	}

	// No more tests needed for UnmarshalJSON, because UnmarshalJSON is based on ParseJSON
}
