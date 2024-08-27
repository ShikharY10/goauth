package utils

import (
	"crypto/sha256"
	"errors"
)

// High level hash function. It takes paylaod and return SHA256 hash of it.
// iteration specifies how many the paylaod will go the hash function.
func HashWithSHA256(payload []byte, iteration int) ([]byte, error) {

	if iteration == 0 {
		return nil, errors.New("iteration should greater or equal to 1")
	}

	hashAlgo := sha256.New()
	var hashDigest []byte = payload

	for i := 0; i < iteration; i++ {
		hashAlgo.Write(hashDigest)
		hashDigest = hashAlgo.Sum(nil)
	}
	return hashDigest, nil
}
