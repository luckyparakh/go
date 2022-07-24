package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	// init
	ele := []int{9, 8, 7, 6}
	// execution
	BubbleSort(ele)
	//validation
	assert.NotNil(t, ele)
	assert.Equal(t, len(ele), 4)
	assert.Equal(t, 6, ele[0])
	assert.Equal(t, 7, ele[1])
	assert.Equal(t, 8, ele[2])
	assert.Equal(t, 9, ele[3])
}

func getElement(n int) []int {
	ele := make([]int, n)
	i := 0
	for j := n - 1; j == 0; j-- {
		ele[i] = j
		i++
	}
	return ele
}

func BenchmarkBubbleSort10(b *testing.B) {
	ele := getElement(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(ele)
	}
}

func BenchmarkSort10(b *testing.B) {
	ele := getElement(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(ele)
	}
}

func BenchmarkBubbleSort50(b *testing.B) {
	ele := getElement(50)
	for i := 0; i < b.N; i++ {
		BubbleSort(ele)
	}
}

func BenchmarkSort50(b *testing.B) {
	ele := getElement(50)
	for i := 0; i < b.N; i++ {
		sort.Ints(ele)
	}
}


func BenchmarkBubbleSort1000(b *testing.B) {
	ele := getElement(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(ele)
	}
}

func BenchmarkSort1000(b *testing.B) {
	ele := getElement(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(ele)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	ele := getElement(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(ele)
	}
}
