package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Tx struct {
	FromID string `json:"from"`
	ToID   string `json:"to"`
	Value  uint64 `json:"value"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {

	privateKey, err := crypto.LoadECDSA("zblock/accounts/kennedy.ecdsa")
	if err != nil {
		return fmt.Errorf("unable to laod private key for node: %w", err)
	}

	tx := Tx{
		FromID: "0xF01813E4B85e178A83e29B8E7bF26BD830a25f32",
		ToID:   "Aaron",
		Value:  1000,
	}

	// get a slice of byte easily of tx
	data, err := json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
	}

	// get a 32 bytes hash of the byte slice of data above
	v := crypto.Keccak256(data)

	sig, err := crypto.Sign(v, privateKey)
	if err != nil {
		return fmt.Errorf("unable to sign: %w", err)
	}

	// printout the signature
	fmt.Println("SIG:", hexutil.Encode(sig))
	// ==================================================================================
	// OVER THE WIRE

	// returns the public key from the signature - using the ECDSA
	publicKey, err := crypto.SigToPub(v, sig)
	if err != nil {
		return fmt.Errorf("unable to pub: %w", err)
	}

	// return the ethereum common address := maybe by returning the first twenty bytes
	fmt.Println("PUB", crypto.PubkeyToAddress(*publicKey).String())

	// ==================================================================================

	tx = Tx{
		FromID: "0xF01813E4B85e178A83e29B8E7bF26BD830a25f32",
		ToID:   "Frank",
		Value:  250,
	}

	// get a slice of byte easily of tx
	data, err = json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
	}

	// get a 32 bytes hash of the byte slice of data above
	v2 := crypto.Keccak256(data)

	sig2, err := crypto.Sign(v2, privateKey)
	if err != nil {
		return fmt.Errorf("unable to sign: %w", err)
	}

	// printout the signature
	fmt.Println("SIG:", hexutil.Encode(sig2))
	// ==================================================================================
	// OVER THE WIRE

	tx2 := Tx{
		FromID: "0xF01813E4B85e178A83e29B8E7bF26BD830a25f32",
		ToID:   "Frank",
		Value:  250,
	}

	// get a slice of byte easily of tx
	data, err = json.Marshal(tx2)
	if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
	}

	// get a 32 bytes hash of the data slice
	v2 = crypto.Keccak256(data)

	// returns the public key from the signature - using the ECDSA
	publicKey, err = crypto.SigToPub(v2, sig2)
	if err != nil {
		return fmt.Errorf("unable to pub: %w", err)
	}

	// return the ethereum common address := maybe by returning the first twenty bytes
	fmt.Println("PUB", crypto.PubkeyToAddress(*publicKey).String())

	return nil
}

// Get the data
// get the byte slice
// get the 32byte hash of the byte slice (using keccak or sha256)
// then sign the 32 byte hash with your private key
// then forward it to the node for processing
