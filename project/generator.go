package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		f := min + rand.Float64()*(max-min)
		sf := fmt.Sprintf("%.2f", f)
		if s, err := strconv.ParseFloat(sf, 64); err == nil {
			res[i] = s
		}
	}
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randFloats(1.10, 500.00, 500))
}
