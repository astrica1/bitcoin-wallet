package seed

import (
	"log"

	"github.com/tyler-smith/go-bip39"
)

type EntropySize int

const (
	EntropySize128 EntropySize = 128 // 12 words
	EntropySize160 EntropySize = 160 // 15 words
	EntropySize192 EntropySize = 192 // 18 words
	EntropySize224 EntropySize = 224 // 21 words
	EntropySize256 EntropySize = 256 // 24 words
)

func generateMnemonic(entropySize EntropySize) string {
	entropy, err := bip39.NewEntropy(int(entropySize))
	if err != nil {
		log.Fatal(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatal(err)
	}

	return mnemonic
}

func GenerateSeed(wordCount int) []byte {
	var entropySize EntropySize
	switch wordCount {
	case 12:
		entropySize = EntropySize128
	case 15:
		entropySize = EntropySize160
	case 18:
		entropySize = EntropySize192
	case 21:
		entropySize = EntropySize224
	case 24:
		entropySize = EntropySize256
	default:
		log.Fatal("Invalid word count")
		entropySize = EntropySize256
		return nil
	}

	mnemonic := generateMnemonic(entropySize)

	seed := bip39.NewSeed(mnemonic, "")
	return seed
}
