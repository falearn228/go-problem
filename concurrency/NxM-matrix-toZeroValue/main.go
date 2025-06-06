package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

const N = 5
const M = 6

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	startch := make(chan struct{})

	a := make([][]int, N)
	for i := range a {
		a[i] = make([]int, M)
		for j := range a[i] {
			a[i][j] = rand.Intn(10)
		}
	}

	printMatrix(a)
	expected := 0
	copya := make([]int, M)
	for i := range a {
		copy(copya, a[i])
		sort.Ints(copya)
		left, right := 0, M-1

		for left <= right {
			middle := left + (right-left)/2
			if copya[middle] == expected {
				col := simpleSearch(a[i])
				wg.Add(1)
				go zeroRowCol(a, i, col, mu, wg, startch)
			}
			if copya[middle] > expected {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	close(startch)
	wg.Wait()
	printMatrix(a)
}

func simpleSearch(a []int) int {
	result := 0
	for i := range a {
		if a[i] == 0 {
			result = i
			break
		}
	}
	return result
}

func zeroRowCol(a [][]int, row, col int, mu *sync.Mutex, wg *sync.WaitGroup, startch <-chan struct{}) {
	defer wg.Done()
	<-startch

	mu.Lock()
	// Зануляем строку
	for j := range a[row] {
		a[row][j] = 0
	}

	// Зануляем столбец
	for i := range a {
		a[i][col] = 0
	}

	mu.Unlock()
}

func printMatrix(a [][]int) {
	for _, row := range a {
		fmt.Println(row)
	}
	fmt.Println("-----------")
}
