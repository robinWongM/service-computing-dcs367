package sort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	sorted := Sort([]int{5, 4, 3, 2, 1})
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("expected %v but got %v", expected, sorted)
	}
}

func generateRandomArray(arrayLen int) []int {
	var a []int
	for i := 0; i < arrayLen; i++ {
		a = append(a, rand.Intn(1000000))
	}
	return a
}

func BenchmarkSort(b *testing.B) {
	numbers := generateRandomArray(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(numbers)
	}
}
