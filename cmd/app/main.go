package main

import (
	"fmt"
	"log"

	"github.com/astrica1/bitcoin-wallet/internal/app/seed"
	"github.com/astrica1/bitcoin-wallet/internal/app/wallet"
)

func main() {
	seed, mnemonic := seed.GenerateSeed(24)

	masterKey, err := wallet.ExtractKey(seed)
	if err != nil {
		log.Fatal(err)
	}

	legacyAddress, err := wallet.GetLegacyAddress(masterKey)
	if err != nil {
		log.Fatal(err)
	}

	segwitAddress, err := wallet.GetSegwitAddress(masterKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seed Words:", mnemonic)
	fmt.Println("Legacy Address:", legacyAddress)
	fmt.Println("Segwit Address:", segwitAddress)
}
