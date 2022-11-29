package main

import (
	"bytes"
	"fmt"
)

func main() {
	word := []byte("Lorem Ipsu 也称乱数假文或者哑元文本 印刷及排版领域所常用的虚拟文字。")

	fmt.Printf("%s =  %[1]x\n", word)
	var size int

	for i := range word {
		if i > 0 {
			size = 1
			break
		}
	}
	copy(word[:size], bytes.ToUpper(word[:size]))
	fmt.Printf("%s = % [1]x\n", word)
}
