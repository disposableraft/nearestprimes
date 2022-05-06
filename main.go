package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	limit, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	p := getNearestPrimes(x, limit)
	fmt.Fprintf(os.Stdout, "%v\n", p)
}

func getNearestPrimes(x int, limit int) []int {
	if x == 0 {
		return []int{}
	}
	// a set-like data structure; using struct{} saves memory
	primes := make(map[int]struct{})

	if isPrime(x) {
		primes[x] = struct{}{}
	}

	gp := findGreaterPrimes(x+1, limit)
	lp := findLesserPrimes(x-1, limit)

	for i := range gp {
		primes[i] = struct{}{}
	}

	for i := range lp {
		primes[i] = struct{}{}
	}

	return toSortedSlice(primes)

}

func toSortedSlice(x map[int]struct{}) []int {
	var p []int
	for i := range x {
		p = append(p, i)
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i] < p[j]
	})
	return p
}

func findGreaterPrimes(x int, limit int) map[int]struct{} {
	primes := make(map[int]struct{})

	for i := x; ; i++ {
		if len(primes) == limit {
			return primes
		}
		if isPrime(i) {
			primes[i] = struct{}{}
		}
	}
}

func findLesserPrimes(x int, limit int) map[int]struct{} {
	primes := make(map[int]struct{})

	for i := x; ; i-- {
		if len(primes) == limit {
			return primes
		}
		if isPrime(i) {
			primes[i] = struct{}{}
		}
	}
}

func isPrime(x int) bool {
	limit := math.Round(math.Sqrt(float64(x)))
	// bug here where 4 was prime because < 4
	for i := 2; i <= int(limit); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
