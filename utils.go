package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

func hash256(val string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(val)))
}

func ensureEven(hashes *[]string) {
	if len(*hashes)%2 == 1 {
		*hashes = append(*hashes, (*hashes)[len(*hashes)-1])
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printEvent(args ...string) {
	magenta := color.New(color.FgMagenta).SprintFunc()
	log.Println(magenta("event"), strings.Join(args, " "))
}

func isLeaf(n MerkleNode) bool {
	return n.left == nil && n.right == nil
}

func printHelpManual() {
	log.Println("Available commands:\n\t1. make create-file - create a random test file, size: 1GB\n\t2. make run -\n\t\ta) create a merkletree using 'data.txt' generated in (1)\n\t\tb) print the merkle tree (hashes only)\n\t\tc) run membership and non-membership proofs using the generated merkle tree and hardcoded hashes")
}
