package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T)  {
	repeated := Repeat("b", 6)
	expected := "bbbbbb"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 6)
    }
}

func ExampleRepeat() {
    repeat := Repeat("a", 6)
    fmt.Println(repeat)
    // Output: aaaaaa
}


