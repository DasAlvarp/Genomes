package main

import(
	"fmt"
	"math/rand"
)


func main() {

}

//returns one string
func generateSequence(length int)(string){
	st := ""
	for x:= 0; x < length; x++{
		st += getNucleotide()
	}
	return st
}

//gets one character for a nucleotide
func getNucleotide() (string) {
	switch rand.Intn(4) {
	case 0:
		return "a"
	case 1:
		return "b"
	case 2:
		return "c"
	default:
		return "d"
	}
	
}