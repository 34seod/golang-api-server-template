package testutils

import "math/rand"

func GenerateRandomString(n int) string {
	lettersAndNumbers := []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersAndNumbers[rand.Intn(len(lettersAndNumbers))]
	}
	return string(b)
}

func GenerateRandomStringPtr(n int) *string {
	lettersAndNumbers := []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersAndNumbers[rand.Intn(len(lettersAndNumbers))]
	}
	str := string(b)
	return &str
}
