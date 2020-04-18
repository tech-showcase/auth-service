package helper

import "testing"

func TestGenerate4CharsPassword(t *testing.T) {
	dummyInput := "hello world"
	expectedOutput := "cde9"

	output := Generate4CharsPassword(dummyInput)

	if output != expectedOutput{
		t.Fatal("unexpected output")
	}
}
