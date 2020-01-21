package menupad_test

import (
	"testing"
	"github.com/misostack/menupad"
	"fmt"
)

type TimeSlotTest struct {
	hours uint8
	minutes uint8
	expected uint8
}

func TestCalculate(t *testing.T) {
	var x, y int
	// static case
	x = 5
	y = 10

	got := menupad.Calculate(uint(x), uint(y))
	expected := uint(x + y)
	if got != expected {
		t.Fatalf("Expected %v, got %v", got, expected)
	}
}

func TestSubtract(t *testing.T) {
	x,y := 5, 10

	got := menupad.Subtract(x, y)
	expected := x - y
	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestTimeSlotID(t *testing.T){
	var testcases = []TimeSlotTest{
		{0, 10, 1},
		{1, 10, 7},
		{2, 40, 16},
	}
	for _, tc := range testcases {
		got := menupad.TimeSlotID(tc.hours, tc.minutes)
		fmt.Printf("TimeSlotID(%v, %v) = %v\n", tc.hours, tc.minutes, got)
		if got != tc.expected {
			t.Errorf("TimeSlotID(%v, %v) = %v; expected : %v", tc.hours, tc.minutes, got, tc.expected)
		}
	}
}