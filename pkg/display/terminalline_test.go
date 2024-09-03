package display_test

import (
	"fmt"
	"gotty/pkg/display"
	"gotty/test/utils"
	"reflect"
	"testing"

	"github.com/fatih/color"
)

func TestSetText(t *testing.T) {
	rec := utils.NewLogRecorder()

	sut := display.NewTerminalLine(1)
	sut.SetText("FooBar")

	got := rec.ToAnsiString()
	want := "\x1b[1;0H\r\x1b[K" + "FooBar"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected display manager to be %+v, but got %+v", want, got)
	}
}

func TestClear(t *testing.T) {
	rec := utils.NewLogRecorder()

	sut := display.NewTerminalLine(1)
	sut.Clear()

	got := rec.ToAnsiString()
	want := "\x1b[1;0H\r\x1b[K"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected display manager to be %+v, but got %+v", want, got)
	}
}

func TestPaintText(t *testing.T) {
	got := display.PaintText(color.FgRed, "Foo")
	want := "Foo"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected display manager to be %+v, but got %+v", want, got)
	}
}

// func TestShowMissMessage(t *testing.T) {
// 	rec := utils.NewLogRecorder()

// 	sut := display.NewTerminalLine(1)
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()
// 		sut.ShowMissMessage()
// 	}()

// 	wg.Wait()

// 	got := rec.ToArray()
// 	want := []string{"MISS!"}

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("Expected display manager to be %+v, but got %+v", want, got)
// 	}
// }

func TestShowProgressBar(t *testing.T) {
	type testCase struct {
		current int
		total   int
		want    []string
	}

	testCases := []testCase{
		{
			current: 1,
			total:   1,
			want:    []string{"1 / 1 [====================>]"},
		},
		{
			current: 1,
			total:   2,
			want:    []string{"1 / 2 [==========>----------]"},
		},
		{
			current: 1,
			total:   10,
			want:    []string{"1 / 10 [==>------------------]"},
		},
		{
			current: 2,
			total:   10,
			want:    []string{"2 / 10 [====>----------------]"},
		},
		{
			current: 9,
			total:   10,
			want:    []string{"9 / 10 [==================>--]"},
		},
		{
			current: 10,
			total:   10,
			want:    []string{"10 / 10 [====================>]"},
		},
		{
			current: 1,
			total:   1,
			want:    []string{"1 / 1 [====================>]"},
		},
		{
			current: 1,
			total:   20,
			want:    []string{"1 / 20 [=>-------------------]"},
		},
		{
			current: 1,
			total:   100,
			want:    []string{"1 / 100 [>--------------------]"},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d_%d", tc.current, tc.total), func(t *testing.T) {
			rec := utils.NewLogRecorder()

			display.ShowProgressBar(tc.current, tc.total, display.NewTerminalLine(1))

			got := rec.ToArray()

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected display manager to be %+v, but got %+v", tc.want, got)
			}
		})
	}
}
