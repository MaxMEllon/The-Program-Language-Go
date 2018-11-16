package word

import (
	// "fmt"
	"math/rand"
	"testing"
	"time"
	// "unicode"
)

const MAX_UTF_BYTE_LENGTH = 6

func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(50) + MAX_UTF_BYTE_LENGTH*2
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	i := 0
	for {
		if i >= len(runes)/2 {
			break
		}
		r := rune(rng.Intn(0x1000))
		if r == runes[i] {
			continue
		}
		runes[i] = r
		i++
	}
	return string(runes)
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(50)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		// fmt.Printf("%s\n", p)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome: %q", p)
		}
	}
}

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsNotPalindrome: %q", p)
		}
	}
}
