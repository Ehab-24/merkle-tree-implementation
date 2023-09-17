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

	printEvent("verifying membership of", elems[len(elems)-1].hash)
	node := tree.ProveMembership(elems)

	bytesToPrint := 512
	if node != nil {
		color.Green("\nData exists in the merkle tree! Head (%d/%d bytes):", bytesToPrint, 8192000)
		log.Println(node.content[:bytesToPrint])
	} else {
		color.Red("\nNo such data.")
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
