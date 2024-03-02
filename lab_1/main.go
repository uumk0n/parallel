package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const (
	n = 6000
	m = 1000
)

func firstMethod(n int) (float64, string) {
	minElapsed := time.Hour
	maxElapsed := time.Nanosecond
	var elapsedSum time.Duration

	arr := make([][]float64, n)

	diffs := make([]time.Duration, 1000)

	var variance float64

	for k := 0; k < m; k++ {
		start := time.Now()

		for j := 0; j < n; j++ {
			arr[j] = make([]float64, n)
			for i := 0; i < n; i++ {
				arr[j][i] = float64(i) / float64(j+1)
			}
		}
		elapsed := time.Since(start)
		diffs[k] = elapsed
		elapsedSum += elapsed
		if elapsed < minElapsed {
			minElapsed = elapsed
		}
		if elapsed > maxElapsed {
			maxElapsed = elapsed
		}
	}

	for _, diff := range diffs {
		variance += math.Abs((elapsedSum.Seconds() / 1000) - diff.Seconds())
	}
	variance /= 1000

	return elapsedSum.Seconds() / 10, "1 метод"
}

func secondMethod(n int) (float64, string) {
	minElapsed := time.Hour
	maxElapsed := time.Nanosecond
	var elapsedSum time.Duration

	arr := make([][]float64, n)

	diffs := make([]time.Duration, 1000)

	var variance float64

	for k := 0; k < m; k++ {
		start := time.Now()
		for i := 0; i < n; i++ {
			arr[i] = make([]float64, n)
			for j := 0; j < n; j++ {
				arr[i][j] = float64(i) / float64(j+1)
			}
		}
		elapsed := time.Since(start)
		diffs[k] = elapsed
		elapsedSum += elapsed
		if elapsed < minElapsed {
			minElapsed = elapsed
		}
		if elapsed > maxElapsed {
			maxElapsed = elapsed
		}
	}

	for _, diff := range diffs {
		variance += math.Abs((elapsedSum.Seconds() / 1000) - diff.Seconds())
	}
	variance /= 1000

	return elapsedSum.Seconds() / 10, "2 метод"
}

func thirdMethod(n int) (float64, string) {
	minElapsed := time.Hour
	maxElapsed := time.Nanosecond
	var elapsedSum time.Duration

	arr := make([][]float64, n)

	diffs := make([]time.Duration, 1000)

	var variance float64

	for k := 0; k < m; k++ {
		start := time.Now()
		for i := n - 1; i > 0; i-- {
			arr[i] = make([]float64, n)
			for j := n - 1; j > 0; j-- {
				arr[i][j] = float64(i) / float64(j+1)
			}
		}
		elapsed := time.Since(start)
		diffs[k] = elapsed
		elapsedSum += elapsed
		if elapsed < minElapsed {
			minElapsed = elapsed
		}
		if elapsed > maxElapsed {
			maxElapsed = elapsed
		}
	}

	for _, diff := range diffs {
		variance += math.Abs((elapsedSum.Seconds() / 1000) - diff.Seconds())
	}
	variance /= 1000

	return elapsedSum.Seconds() / 10, "3 метод"
}

func fourthMethod(n int) (float64, string) {
	minElapsed := time.Hour
	maxElapsed := time.Nanosecond
	var elapsedSum time.Duration

	arr := make([][]float64, n)

	diffs := make([]time.Duration, 1000)

	var variance float64

	for k := 0; k < m; k++ {
		start := time.Now()

		for j := n - 1; j > 0; j-- {
			arr[j] = make([]float64, n)
			for i := n - 1; i > 0; i-- {
				arr[j][i] = float64(i) / float64(j+1)
			}
		}
		elapsed := time.Since(start)
		diffs[k] = elapsed
		elapsedSum += elapsed
		if elapsed < minElapsed {
			minElapsed = elapsed
		}
		if elapsed > maxElapsed {
			maxElapsed = elapsed
		}
	}

	for _, diff := range diffs {
		variance += math.Abs((elapsedSum.Seconds() / 1000) - diff.Seconds())
	}
	variance /= 1000

	return elapsedSum.Seconds() / 10, "4 метод"
}

func main() {
	for n := 1000; n <= 3000; n += 1000 {
		fmt.Println(n)
		n := n
		wg := sync.WaitGroup{}
		wg.Add(4)
		go func() {
			fmt.Println(firstMethod(n))
			wg.Done()
		}()
		go func() {
			fmt.Println(secondMethod(n))
			wg.Done()
		}()
		go func() {
			fmt.Println(thirdMethod(n))
			wg.Done()
		}()
		go func() {
			fmt.Println(fourthMethod(n))
			wg.Done()
		}()
		wg.Wait()
	}

}
