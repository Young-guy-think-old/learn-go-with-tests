package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var numbers = []int{1,2,3,4,5}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
        got := Sum(numbers)
        want := 15

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })
}


func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
	for i := 0; i < b.N; i++ {
		SumAll(numbers)
	}
}

func ExampleSum() {
	got := Sum(numbers)
	fmt.Println(strconv.Itoa(got))
	// Output: 15
}

func TestSumAll(t *testing.T) {

    got := SumAll([]int{1, 2}, []int{0, 9})
    want := []int{3, 9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func TestSumAllTails(t *testing.T) {
    checkSums := func(t testing.TB, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }

    t.Run("make the sums of tails of", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}
        checkSums(t, got, want)
    })

    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}
        checkSums(t, got, want)
    })
}