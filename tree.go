package main

import f "fmt"
import "flag"
import "strconv"
import "os"

func main() {
	flag.Parse()
	s := flag.Arg(0)
	i, err := strconv.Atoi(s)
	if err != nil {

		f.Println(err)
		os.Exit(2)
	}

	buildTree(i)

}

func buildTree(size int) {

	for i := size; i > 0; i-- {

		for z := i; z > 0; z-- {
			f.Print(" ")
		}
		f.Print("*")
		for j := 0; j < (size-i)*2; j++ {
			f.Print("*")
		}
		f.Println("")
	}
	for i := 0; i < size; i++ {
		f.Print(" ")
	}
	f.Print("*")
}
