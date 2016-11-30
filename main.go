package main

import(
	"fmt"
	"strconv"
	"math/rand"
	"os"
	"time"
)

//input params:1 max number of sequences,2 length of sequence, 3number of motifs in, 4max length of motif,5number of possible changes to a motif
func main() {
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
		seq,_ :=strconv.Atoi(os.Args[2])
		go sequenceGenome(changes, seq, motifs)
	}
	time.Sleep(99999999)

}

func sequenceGenome(changeNum int, maxLength int, motifs []string){
	fmt.Print("oWo\n")
	max,_:=strconv.Atoi(os.Args[4])
	sequence:=make(chan string)

	shortestMotif := max
	for x:=0;x<len(motifs);x++{
		if len(motifs[x]) < shortestMotif{
			shortestMotif = len(motifs[x])
		}
	}
	genome :=""
	for x:=0;x<100;x++{
		jam := improvise(changeNum, motifs[rand.Intn(len(motifs))])
		random := generateRandom(rand.Intn(max))
		select {
		case sequence <-jam:
			fmt.Println(jam)
		case sequence <-random:
			fmt.Println(random)

		default:
			fmt.Println(jam + "D:")


		}
		fmt.Println(x)
		latestStr:=<-sequence
		close(sequence)

		genome += latestStr
	}
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