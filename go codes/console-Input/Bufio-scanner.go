package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scanner() []string {
	scanner := bufio.NewScanner(os.Stdin)
	storeValue := make([]string, 0, 0)
	for scanner.Scan() {
		storeValue = append(storeValue, scanner.Text())
		break
	}
	return storeValue
}

func subStrCon(a, b string) {
if strings.Compare(a, b)


}
func main() {
	res := scanner()
	for v := range res {

	}
}
