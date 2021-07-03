package main

import (
	"math"
	"math/rand"
	"time"
)

func RoundPlus(n float64, decimals int) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func GenerateRandomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	random := min + rand.Float64()*(max-min)

	return random
}

func GenerateRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(max-min+1) + min

	return random
}
