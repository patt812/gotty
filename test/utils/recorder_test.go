package utils_test

import (
	"fmt"
	"gotty/test/utils"
	"testing"
)

func TestToString(t *testing.T) {
	rec := utils.NewLogRecorder()

	fmt.Println("Hello, World!")

	got := rec.ToString()
	if got != "Hello, World!\n" {
		t.Errorf("Expected 'Hello, World!' in the output, but got: %s", got)
	}
}

func TestToString_MultipleLines(t *testing.T) {
	rec := utils.NewLogRecorder()

	fmt.Println("Hello, World!")
	fmt.Println("Goodbye, World!")

	got := rec.ToString()
	if got != "Hello, World!\nGoodbye, World!\n" {
		t.Errorf("Expected 'Hello, World!\nGoodbye, World!' in the output, but got: %s", got)
	}
}

func TestToArray_MultipleLines(t *testing.T) {
	rec := utils.NewLogRecorder()

	fmt.Println("Hello, World!")
	fmt.Println("Goodbye, World!")

	got := rec.ToArray()
	want := []string{"Hello, World!", "Goodbye, World!"}

	for i, line := range want {
		if line != got[i] {
			t.Errorf("Expected '%s' in line %d, but got '%s'", line, i+1, got[i])
		}
	}
}
