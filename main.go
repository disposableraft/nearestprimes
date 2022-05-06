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

	np := findNeighborPrimes(x, limit)

	for i := range np {
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

func findNeighborPrimes(x int, limit int) map[int]struct{} {
	primes := make(map[int]struct{})

	for i := x + 1; len(primes) < limit; i++ {
		if isPrime(i) {
			primes[i] = struct{}{}
		}
	}

	for i := x - 1; len(primes) < limit*2; i-- {
		if isPrime(i) {
			primes[i] = struct{}{}
		}
	}

	return primes
}

func isPrime(x int) bool {
	if x == 2 {
		return true
	}

	if x%2 == 0 {
		return false
	}

	limit := math.Round(math.Sqrt(float64(x)))

	for i := 3; i <= int(limit); i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}
