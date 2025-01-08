package wallet

import (
	"fmt"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

const (
	LEGACY_PATH = "m/44'/0'/0'/0/0"
	SEGWIT_PATH = "m/84'/0'/0'/0/0"
)

func ExtractKey(seed []byte) (masterKey *hdkeychain.ExtendedKey, err error) {
	masterKey, err = hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return masterKey, nil
}

func GetLegacyAddress(masterKey *hdkeychain.ExtendedKey) (string, error) {
	legacyKey, err := deriveKey(masterKey, LEGACY_PATH)
	if err != nil {
		return "", err
	}
	legacyPubKey, _ := legacyKey.ECPubKey()
	legacyAddress, err := btcutil.NewAddressPubKey(legacyPubKey.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	return legacyAddress.EncodeAddress(), nil
}

func GetSegwitAddress(masterKey *hdkeychain.ExtendedKey) (string, error) {
	segwitKey, err := deriveKey(masterKey, SEGWIT_PATH)
	if err != nil {
		return "", err
	}
	segwitPubKey, _ := segwitKey.ECPubKey()
	witnessProg := btcutil.Hash160(segwitPubKey.SerializeCompressed())
	segwitAddress, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	return segwitAddress.EncodeAddress(), nil
}

func deriveKey(masterKey *hdkeychain.ExtendedKey, path string) (*hdkeychain.ExtendedKey, error) {
	parts := strings.Split(path, "/")
	key := masterKey
	for _, part := range parts[1:] {
		index := uint32(0)
		hardened := false
		if strings.HasSuffix(part, "'") {
			hardened = true
			part = strings.TrimSuffix(part, "'")
		}
		fmt.Sscanf(part, "%d", &index)
		if hardened {
			index += hdkeychain.HardenedKeyStart
		}
		var err error
		key, err = key.Derive(index)
		if err != nil {
			return nil, err
		}
	}

	return key, nil
}
