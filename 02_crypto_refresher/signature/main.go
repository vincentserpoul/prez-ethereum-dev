package main

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generate a private key - ethereum style
	sk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("error generating the private key: %v", err)
	}

	// Get the address
	addr := crypto.PubkeyToAddress(sk.PublicKey)

	fmt.Printf("your public address is %s\n\n", addr.String())

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("give me something to sign\n")
		s, errRS := reader.ReadString('\n')
		if errRS != nil {
			log.Fatalf("error scanning string to sign: %v\n", errRS)
		}
		// Removing the \n
		s = strings.Trim(s, " \n")

		sig, errSig := sign(s, sk)
		if errSig != nil {
			log.Fatalf("%v\n", errSig)
		}

		fmt.Printf("the signature of `%s` is a 65-byte [R || S || V] format where V is 0 or 1.\n0x%X\n\n", s, sig)
	}
}

func sign(txt string, sk *ecdsa.PrivateKey) ([]byte, error) {
	// First hash the message
	h := hash(txt)

	// Then sign it
	sig, errSig := crypto.Sign(h, sk)
	if errSig != nil {
		return nil, fmt.Errorf("sign(%s): %v", txt, errSig)
	}

	return sig, nil
}

func hash(txt string) []byte {
	msg := fmt.Sprintf(
		"\x19Ethereum Signed Message:\n%d%s",
		len(txt),
		txt,
	)
	return crypto.Keccak256([]byte(msg))
}
