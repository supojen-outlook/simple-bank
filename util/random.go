package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// summary: generate random integer
func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// summary: generate random name
func RandomName() string {
	names := []string{
		"David", "Brian", "Jacky", "Eric", "Jackson", "Johnson", "Miranda", "Bryan",
	}
	n := len(names)
	index := rand.Intn(n)
	return names[index]
}

// summary: generate random amount of oney
func RandomMoney() int64 {
	return RandInt(100, 1000)
}

// summary: generate random currency
func RandomCurrency() string {
	currencies := []string{
		"USD", "EUR", "CAD",
	}
	n := len(currencies)
	index := rand.Intn(n)
	return currencies[index]
}
