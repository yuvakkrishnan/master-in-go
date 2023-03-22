package main

<<<<<<< HEAD
import (
	"os"
)

func main() {
	files := os.OpenFile("text.txt")
=======
import "fmt"

func reverse(s string) string {
	letter := []rune(s)
	for i, j := 0, len(letter)-1; i < j; i, j = i+1, j-1 {
		letter[i], letter[j] = letter[j], letter[i]
	}
	return string(letter)
}

func main() {
	d := "Hello"
	e := reverse(d)
	fmt.Println(e)
>>>>>>> 825db90e4d42816866d191d319ada5f3665140f7
}
