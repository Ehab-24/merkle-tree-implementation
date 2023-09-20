package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

type MerkleNode struct {
	hash    string
	content string
	left    *MerkleNode
	right   *MerkleNode
}

type MerkleTree struct {
	root MerkleNode
}

func createMerkleRoot(nodes []MerkleNode) MerkleNode {
	n := len(nodes)
	if n == 2 {
		return NewNode(hash256(nodes[0].hash+nodes[1].hash), &nodes[0], &nodes[1], "")
	}

	half := n / 2
	left := createMerkleRoot(nodes[:half])
	right := createMerkleRoot(nodes[half:])

	// sort nodes in alphabetical order
	if left.hash > right.hash {
		temp := left
		left = right
		right = temp
	}

	hash := hash256(left.hash + right.hash)

	return NewNode(hash, &left, &right, "")
}

func NewNode(hash string, left *MerkleNode, right *MerkleNode, content string) MerkleNode {
	return MerkleNode{
		hash:    hash,
		left:    left,
		right:   right,
		content: content,
	}
}

func NewMerkleTree(chunks []string) MerkleTree {
	hashes := make([]string, len(chunks))
	for i, chunk := range chunks {
		hashes[i] = hash256(chunk)
	}

	leaves := make([]MerkleNode, len(hashes))
	for i, hash := range hashes {
		leaves[i] = NewNode(hash, nil, nil, chunks[i])
	}

	return MerkleTree{
		root: createMerkleRoot(leaves),
	}
}

func printTreeRec(n *MerkleNode, level int) {
	if n == nil {
		return
	}

	bullet := fmt.Sprintf("%d ", level)
	if level == 0 {
		bullet = ""
	}

	str := n.hash
	if n.content != "" {
		str = n.hash + fmt.Sprintf(" (Leaf)")
	}

	cyan := color.New(color.FgCyan).SprintFunc()
	log.Printf("%*s%s", level*2+9, cyan(bullet), str)

	printTreeRec(n.left, level+1)
	printTreeRec(n.right, level+1)
}

func (t *MerkleTree) Print() {
	printEvent("\nMerkle Tree")
	printTreeRec(&t.root, 0)
}

func (t *MerkleTree) RootHash() string {
	return t.root.hash
}

/************************************************
 * Membership proofs
************************************************/

func (t *MerkleTree) ProveMembership(elems []ProofElement) *MerkleNode {

	currentNode := t.root
	for i := range elems {
		if elems[i].hash != currentNode.hash {
			log.Println("found")
			return nil
		}

		if i < len(elems)-1 && elems[i].direction == Left {
			currentNode = *currentNode.left
		} else if i < len(elems)-1 {
			currentNode = *currentNode.right
		}
	}

	return &currentNode
}

/*
 * All nodes in the tree are sorted in alphabetical order meaning, for any inner node N, N.left.hash < N.right.hash
 */
func findNode(n *MerkleNode, targetHash string) *MerkleNode {
	if n == nil {
		return nil
	}
	if n.hash == targetHash {
		return n
	}
	left := findNode(n.left, targetHash)
	right := findNode(n.right, targetHash)

	if left != nil {
		return left
	}
	return right
}

func (t *MerkleTree) ProveNonMembership(hash string) bool {
	n := findNode(&t.root, hash)
	return n == nil
}
