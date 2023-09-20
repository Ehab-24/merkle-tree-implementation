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
	{hash: "349ddc594c2f21c67faf35a8f4275719449b2314356842d60a4be5e3e299da41", direction: Left},
	{hash: "a5925e7b615e4c904f8ae036e88ad26e8f4eb358f4bac944675ce134bad02ce9", direction: Left},
	{hash: "4d608c1fd840aa9dd005403b71c295b13481d754d22c76f056c7bfb917f9cbcd", direction: Left},
	{hash: "376797797561b820ade42d3273d1c6811d4351b71e76f893333121d6c7c405e8", direction: Left},
	{hash: "e76d2dbcd31d40d4ee87b850d1a806b3ebbec7f3fabeecdc2d2cb0f33846d827", direction: Left},
	{hash: "5efa87d2a8ad7be128600d56fade5bd8afb61d555369a232e14fe43e30f17667", direction: Left},
	{hash: "fdf954e5ce8cb5bf32ca292f9de64e80a7cdecf904de4fe50991222dd75202a6", direction: Right},
	{hash: "d135075ee5b8e87729e96821e9b798ee69d3ad0b8c68cc70cd462522b6ae7f6d", direction: None},
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

		fmt.Printf("\r%d/1024 MB", (chunk+1)*8)
	}
	return nil
}
