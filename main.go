package main

import(
	"fmt"
	"strconv"
	"math/rand"
	"os"
	"time"
)


func main() {

	max,_ := strconv.Atoi(os.Args[1])
	for x:=0;x<max;x++{
		seq,_ :=strconv.Atoi(os.Args[2])
		go sequenceGenome(generateSequence(seq))
	}
	time.Sleep(10000)
}

func sequenceGenome(sequence string){
	fmt.Print(sequence + "oWo\n")
}

//returns one string
func generateSequence(length int)(string){
	st := ""
	for x:= 0; x < length; x++{
		st += getNucleotide()
	}
//	fmt.Print(st + "*___*")
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