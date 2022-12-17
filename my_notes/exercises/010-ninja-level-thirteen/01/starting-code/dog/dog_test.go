package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	e := Years(10)
	if e != 70 {
		// "Expected" "Got" idiom
		t.Error("Expected", 70, "Got", e)
	}
}

func TestYearsTwo(t *testing.T) {
	e := Years(10)
	if e != 70 {
		t.Error("Expected", 70, "got", e)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}
