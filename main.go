package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

const filepath string = "data.txt"

func main() {

	log.SetFlags(0)

	f, err := os.Open(filepath)
	check(err)

	printEvent("reading file...")
	chunks := readFileChunks(f)

	printEvent("generating merkle tree...")
	tree := NewMerkleTree(chunks)

	printEvent("Merkle Tree Root:", tree.RootHash())

	proofOfMembership(&tree)

	proofOfNonMembership(&tree)
}

func proofOfMembership(tree *MerkleTree) {
	log.Println()

	printEvent("verifying membership of", elems[len(elems)-1].hash)
	node := tree.ProveMembership(elems)
	bytesToPrint := 512
	if node != nil {
		color.Green("\nData exists in the merkle tree! Data Head (%d/%d bytes):", bytesToPrint, 8192000)
		log.Println(node.content[:bytesToPrint])
	} else {
		color.Red("\nNo such data.\n")
	}
}
func proofOfNonMembership(tree *MerkleTree) {
	log.Println()

	hash := "z62cacfb4c1d0f852f7925c0c96ae09b0b30e1a303a07b7b8d2d642f0ba91d7c"
	printEvent("verifying non-membership of", hash)
	exists := !tree.ProveNonMembership(hash)

	if exists {
		color.Red("\nData does exist in the merkle tree!")
	} else {
		color.Green("\nData does not exist!")

	}
}

func readFileChunks(f *os.File) []string {
	bytesToRead := 8192000

	var chunks []string
	b := make([]byte, bytesToRead)
	for {
		n, err := f.Read(b)
		if err != nil && err.Error() == "EOF" {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		chunks = append(chunks, string(b[:n]))
	}

	return chunks
}
