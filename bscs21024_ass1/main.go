package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

const filepath string = "data.txt"

func main() {
	log.SetFlags(0)

	if len(os.Args) == 2 {
		command := os.Args[1]
		printEvent("writing into data.txt...")
		execCommand(command)
		color.Green("\nFile written successfully!")
		log.Println("\nHint: run 'make run' to run the program")
		return
	}

	f, err := os.Open(filepath)
	check(err)

	printEvent("reading file...")
	chunks := readFileChunks(f)

	printEvent("generating merkle tree...")
	tree := NewMerkleTree(chunks)

	tree.Print()

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
		color.Green("Data exists in the merkle tree! Data Head (%d/%d bytes):", bytesToPrint, 8192000)
		log.Println(node.content[:bytesToPrint])
	} else {
		color.Red("No such data.\n")
	}
}

func proofOfNonMembership(tree *MerkleTree) {
	log.Println()

	hash := "d72518be626086284a0003d7365aa035e76eb1cab7646c7a506a4143af2fe5fd"
	printEvent("verifying non-membership of", hash)
	exists := !tree.ProveNonMembership(hash)

	if exists {
		color.Red("Data exists in the merkle tree!\n")
	} else {
		color.Green("Data does not exist!\n")
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

func execCommand(command string) {
	if command == "create-file" {
		file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		check(err)
		writeTestFile(file)
	}
}
