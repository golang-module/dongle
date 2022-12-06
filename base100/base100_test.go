package base100

import "testing"

func TestInvalidInput(t *testing.T) {
	if _, err := Decode([]byte("aaaa")); err != invalidDataError {
		t.Errorf("Expected ErrInvalidData but got %v", err)
	}

	if _, err := Decode([]byte("aaa")); err != invalidLengthError {
		t.Errorf("Expected ErrInvalidLength but got %v", err)
	}
}

func TestCoverError(t *testing.T) {
	const message = "are you happy now, code coverage?"
	err := invalidInputError{message}
	if err.Error() != message {
		t.Errorf("(InvalidInputError).Error(): Expected %v, got %v", message, err.Error())
	}
}

var tests = []struct {
	name  string
	text  string
	emoji string
}{
	{
		"ASCII",
		"hello",
		"👟👜👣👣👦",
	},
	{
		"Cyrillic",
		"РАШ Бэ",
		"📇💗📇💇📇💟🐗📇💈📈💄",
	},
	{
		"HelloUnicode",
		"Hello, 世界",
		"🐿👜👣👣👦🐣🐗📛💯💍📞💌💃",
	},
}

func TestDecode(t *testing.T) {
	for _, test := range tests {
		res, err := Decode([]byte(test.emoji))
		if err != nil {
			t.Errorf("%v: Unexpected error: %v", test.name, err)
		}

		if string(res) != test.text {
			t.Errorf("%v: Expected to get '%v', got '%v'", test.name, test.text, string(res))
		}
	}
}

func TestEncode(t *testing.T) {
	for _, test := range tests {
		res := Encode([]byte(test.text))

		if string(res) != test.emoji {
			t.Errorf("%v: Expected to get '%v', got '%v'", test.name, test.emoji, res)
		}
	}
}

func TestFlow(t *testing.T) {
	text := []byte("the quick brown fox 😂😂👌👌👌 over the lazy dog привет")

	res, err := Decode(Encode(text))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(res) != string(text) {
		t.Errorf("Expected to get '%v', got '%v'", string(text), string(res))
	}
}
