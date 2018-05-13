package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("give me something to hash, like 🦄\n")
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error scanning string to hash: %v\n", err)
		}
		// Removing the \n
		s = strings.Trim(s, " \n")

		fmt.Printf("the hash of `%s` is 0x%X\n\n", s, hash(s))
	}
}

func hash(txt string) []byte {
	msg := fmt.Sprintf(
		"\x19Ethereum Signed Message:\n%d%s",
		len(txt),
		txt,
	)
	return crypto.Keccak256([]byte(msg))
}
