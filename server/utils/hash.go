package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashUrl(url string) string {		// hash the url using sha256 algorithm
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hashsum := hasher.Sum(nil)
	hashhex := hex.EncodeToString(hashsum)
	return hashhex[:16]
	// fmt.Println("hash: ", hashhex)
	// return "hi"
}