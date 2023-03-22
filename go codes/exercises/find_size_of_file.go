/* A Go program that prints the size of a file is on the fridge. It calls the
os.Stat function, which returns an os.FileInfo value, and possibly an
error value. Then it calls the Size method on the FileInfo value to get the
file size.
But the original program uses the _ blank identifier to ignore the error value
from os.Stat. If an error occurs (which could happen if the file doesn’t
exist), this will cause the program to fail.
Reconstruct the extra code snippets to make a program that works just like
the original one, but also checks for an error from os.Stat. If the error
from os.Stat is not nil, the error should be reported, and the program
should exit. Discard the magnet with the _ blank identifier; it won’t be used
in the finished program. */

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {

// 	fileInfo, err := os.Stat("flie.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("File Found !!!", fileInfo)
// 	fmt.Println(fileInfo.Size())
// }
