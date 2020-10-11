package api

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// Token used for any match reference
type Token struct {
	generateTime time.Time
	specificCode string
}

// SeedInit is random seed init
func SeedInit() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
}

func generateCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ret := make([]byte, 16)
	for i := range ret {
		ret[i] = letters[rand.Intn(len(letters))]
	}
	return string(ret)
}

func generateToken() Token {
	return Token{
		generateTime: time.Now(),
		specificCode: generateCode(),
	}
}

// IsExpire is Return True if it has expired. (30 minutes of validity.)
func (token *Token) IsExpire() bool {
	duration := time.Now().Sub(token.generateTime)
	times := duration.Minutes()
	if times >= 30 {
		return true
	}
	return false
}
