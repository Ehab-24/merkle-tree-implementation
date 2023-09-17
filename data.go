package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/fatih/color"
)

type Direction string

const (
	Left  Direction = "left"
	Right Direction = "right"
	None  Direction = ""
)

type ProofElement struct {
	hash      string
	direction Direction
}

var elems = []ProofElement{
	{hash: "772babda48037d65975d13a09ceba104176ef929f187f9a617536adccfc734cc", direction: Left},
	{hash: "715b3fa6ac146b7cf8aff2df6e3150b04469804c914d1b7607ddccb58b8c8fd4", direction: Left},
	{hash: "e202c950c7fc96edbae85bbb57350bc995a614cd4fa6e169cf9e59a84842ff07", direction: Left},
	{hash: "7810fc89c2ad4f6d05e4a99fd9cf4633e836f7ef368b1aaa9930b10fe95cf97f", direction: Left},
	{hash: "a34f85d63faad5772633e6cf4afbdf923ace04ac30b7ecaa50daf93805e75e3b", direction: Left},
	{hash: "b14a669ad6b7dd7aaeab427eb0ee01ae0a238007f911d067e180b89c58ae42f2", direction: Left},
	{hash: "bc5878e05f7d6578a395e9f5691c0cdc9a82b34b11eb991c7c582b11cb6885fa", direction: Left},
	{hash: "e62cacfb4c1d0f852f7925c0c96ae09b0b30e1a303a07b7b8d2d642f0ba91d7c", direction: None},
}

func writeTestFile(file *os.File) error {
	literals := []rune{
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',
		' ',
		'.',
		'"',
		'\'',
		':',
		'?',
		'!',
		'(',
		')',
		'[',
		']',
	}

	// 128 chunks (8192000 bytes each) makes 1GB
	data := make([]byte, 8192000)

	color.Magenta("Writing Test File")

	for chunk := 0; chunk < 128; chunk++ {
		for i := 0; i < len(data); i++ {
			literal := literals[rand.Intn(len(literals))]
			data[i] = byte(literal)
		}
		_, err := file.WriteString(string(data))
		if err != nil {
			return err
		}

		fmt.Printf("\rchunks: %d/128", chunk+1)
	}
	return nil
}
