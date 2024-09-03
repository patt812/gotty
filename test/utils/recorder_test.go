package utils_test

import (
	"fmt"
	"gotty/test/utils"
	"testing"
)

func TestToAnsiString(t *testing.T) {
	rec := utils.NewLogRecorder()

	fmt.Println("Hello, World!")

	got := rec.ToAnsiString()
	if got != "Hello, World!\n" {
		t.Errorf("Expected 'Hello, World!' in the output, but got: %s", got)
	}
}

func TestToAnsiString_MultipleLines(t *testing.T) {
	rec := utils.NewLogRecorder()

	fmt.Println("Hello, World!")
	fmt.Println("Goodbye, World!")

	got := rec.ToAnsiString()
	if got != "Hello, World!\nGoodbye, World!\n" {
		t.Errorf("Expected 'Hello, World!\nGoodbye, World!' in the output, but got: %s", got)
	}
}

func TestToAnsiArray_MultipleLines(t *testing.T) {
	rec := utils.NewLogRecorder()

	fmt.Println("Hello, World!")
	fmt.Println("Goodbye, World!")

	got := rec.ToAnsiArray()
	want := []string{"Hello, World!", "Goodbye, World!"}

	for i, line := range want {
		if line != got[i] {
			t.Errorf("Expected '%s' in line %d, but got '%s'", line, i+1, got[i])
		}
	}
}

func TestToString(t *testing.T) {
	rec := utils.NewLogRecorder()

	fmt.Println("\x1b[31mHello, World!\x1b[0m")

	got := rec.ToString()
	if got != "Hello, World!\n" {
		t.Errorf("Expected 'Hello, World!' in the output, but got: %s", got)
	}
}

func TestToArray(t *testing.T) {
	rec := utils.NewLogRecorder()

	fmt.Println("\x1b[31mHello, World!\x1b[0m")
	fmt.Println("\x1b[32mGoodbye, World!\x1b[0m")

	got := rec.ToArray()
	want := []string{"Hello, World!", "Goodbye, World!"}

	for i, line := range want {
		if line != got[i] {
			t.Errorf("Expected '%s' in line %d, but got '%s'", line, i+1, got[i])
		}
	}
}
