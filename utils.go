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
