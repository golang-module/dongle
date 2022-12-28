package base45

import (
	"testing"
)

type testPair struct {
	decoded string
	encoded string
}

type malformedTestPair struct {
	encoded string
	error   error
}

var pairs = []testPair{
	{
		decoded: "AB",
		encoded: "BB8",
	},
	{
		decoded: "Hello!!",
		encoded: "%69 VD92EX0",
	},
	{
		decoded: "base-45",
		encoded: "UJCLQE7W581",
	},
	{
		decoded: "ietf!",
		encoded: "QED8WEX0",
	},
	{
		decoded: "Lorem ipsum dolor sit amet",
		encoded: "$T9ZKE ZD$ED$QE ZDGVC*VDBJEPQESUEBEC7$C",
	},
}

var malformedPairs = []malformedTestPair{
	{
		encoded: "A",
		error:   InvalidLengthError{length: 1, mod: 1},
	},
	{
		encoded: "FAILEd",
		error:   InvalidCharacterError{char: 'd', position: 5},
	},
	{
		encoded: "GGW",
		error:   IllegalBase45ByteError{position: 0},
	},
	{
		encoded: ":::",
		error:   IllegalBase45ByteError{position: 0},
	},
	{
		encoded: "ABC:::",
		error:   IllegalBase45ByteError{position: 1},
	},
}

func testEqual(t *testing.T, msg string, args ...interface{}) bool {
	t.Helper()
	if args[len(args)-2] != args[len(args)-1] {
		t.Errorf(msg, args...)
		return false
	}
	return true
}

func TestEncode(t *testing.T) {
	for _, p := range pairs {
		got := Encode(p.decoded)
		testEqual(t, "Encode(%q) = %q, want %q", p.decoded, got, p.encoded)
	}
}

func TestDecode(t *testing.T) {
	for _, p := range pairs {
		got, err := Decode(p.encoded)
		if err != nil {
			t.Failed()
		}
		testEqual(t, "Decode(%q) = %q, want %q", p.encoded, got, p.decoded)
	}
}

func TestDecodeMalformed(t *testing.T) {
	for _, p := range malformedPairs {
		_, err := Decode(p.encoded)
		if err == nil {
			t.Failed()
		}
		switch err := err.(type) {
		case InvalidLengthError:
			testEqual(t, "Error(%s) = %s, want %s", p.error.Error(), err.Error(), p.error.Error())
		case InvalidCharacterError:
			testEqual(t, "Error(%s) = %s, want %s", p.error.Error(), err.Error(), p.error.Error())
		case IllegalBase45ByteError:
			testEqual(t, "Error(%s) = %s, want %s", p.error.Error(), err.Error(), p.error.Error())
		}
	}
}
