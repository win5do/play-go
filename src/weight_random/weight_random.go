package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Weight struct {
	value  int
	weight int
}

func WeightRandom(w []Weight) int {
	sum := 0

	for _, i := range w {
		sum += i.weight
	}

	seed := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(seed)
	r := rd.Float64()
	r *= float64(sum)
	fmt.Println("random:", r)

	scale := 0
	for _, i := range w {
		scale += i.weight
		if r < float64(scale) {
			return i.value
		}
	}
	return w[len(w)-1].value
}

func main() {
	w := []Weight{
		{1, 2},
		{2, 4},
		{3, 6},
	}
	r := WeightRandom(w)
	fmt.Println("result:", r)
}
