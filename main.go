package main

/*
*alvarpq 12/4/2016
 */

import(
	"fmt"
	"strconv"
	"math/rand"
	"os"
)

//input params:1 max number of sequences,2 length of sequence, 3number of motifs in, 4max length of motif,5number of possible changes to a motif,6activeSubRegion
func main() {
	fmt.Print("oWo\n")


	max,_:=strconv.Atoi(os.Args[4])
	length,_:=strconv.Atoi(os.Args[3])
	var motifs []string
	motifs = make([]string, length)
	changes,_:=strconv.Atoi(os.Args[5])
	for x:=0;x< length;x++{
		motifs[x] = generateRandom(rand.Intn(max - 10) + 10)
	}
	max,_ = strconv.Atoi(os.Args[1])

	for x:=0;x<max;x++{
		seq,_ :=strconv.Atoi(os.Args[6])
		sequenceGenome(changes, seq, motifs)
	}
}

func sequenceGenome(changeNum int, maxLength int, motifs []string){
	max,_:=strconv.Atoi(os.Args[4])
	machC,_:=strconv.Atoi(os.Args[2])
	shortestMotif := max
	for x:=0;x<len(motifs);x++{
		if len(motifs[x]) < shortestMotif{
			shortestMotif = len(motifs[x])
		}
	}
	genome := generateRandom(rand.Intn(machC/2))
	gStart := len(genome)
	motC := make(chan string)
	randC := make(chan string)

	motifNums := make(chan int)
	go getMotifs(motifs,changeNum,motC, motifNums)

	go getRandomStrings(maxLength,randC)

	for ;len(genome)<maxLength + gStart;{
		select {
		case sequence:= <-motC:
			cardinal:= <- motifNums
			printer :="> m"  +strconv.Itoa(cardinal) +  " "+ strconv.Itoa(len(genome)) + ", " + strconv.Itoa(len(genome) + len(sequence) - 1)+ ": " + sequence[2:]
			println(printer)
			genome += sequence[2:]
		case sequence:= <-randC:
			genome += sequence

		}
	}

	genome += generateRandom(machC - len(genome))
	fmt.Println(genome + ":(")

}

//improvises on a motif
func improvise(changes int, motif string)(string){
	improvised := motif//variable to jam on

	for x:=0;x<changes;x++{
		improvised = replace(motif, rand.Intn(len(motif)))
	}

	if improvised == motif{
		return improvise(changes, motif)
	}else {
		return improvised
	}
}

func replace(motif string, replaceInd int)(string){
	ret :="";
	for x:=0;x<replaceInd;x++{
		ret+=string(motif[x])
	}
	ret+=getNucleotide()
	for x:=replaceInd + 1;x< len(motif);x++{
		ret+=string(motif[x])
	}
	return ret

}

//returns one string
func generateRandom(length int)(string){
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

func getMotifs(motifs []string,changes int, motifChannel chan <- string, motifNums chan <- int){
	for{
		motif := rand.Intn(len(motifs))
		if(motif < 10){
			motifChannel <- "0" + strconv.Itoa(motif) + improvise(changes, motifs[motif])

		}else {
			motifChannel <- strconv.Itoa(motif) + improvise(changes, motifs[motif])
		}
		motifNums <- motif
	}
}

func getRandomStrings(maxLength int, randomChannel chan <- string){
	for{
		randomChannel <-generateRandom(rand.Intn(maxLength))
	}
}