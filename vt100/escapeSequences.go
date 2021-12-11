package vt100

import (
	"fmt"
	"strings"
)

const Yellow = `33`
const Cyan = `36`

func SaveCursor() {
	fmt.Print("\x1b7")
}

func RestoreCursor() {
	fmt.Print("\x1b8")
}

func EraseCurrentLine() {
	fmt.Print("\x1b[2K")
}

func Print(text string, attributes ...string) {
	attributeString := ""

	for _, attribute := range attributes {
		attributeString += attribute + ";"
	}
	attributeString = strings.Trim(attributeString, ";")

	fmt.Printf("\x1b[%vm%v\x1b[0m", attributeString, text)
}
