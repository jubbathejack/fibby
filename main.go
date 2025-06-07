package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func fibSeq(inNum int) *big.Int {

	numX := big.NewInt(0)
	numY := big.NewInt(1)

	for i := 2; i <= inNum; i++ {
		numX, numY = numY, new(big.Int).Add(numX, numY)
	}

	return numY
}

func main() {

	help := flag.Bool("h", false, "Show this help message")
	helpLong := flag.Bool("help", false, "Show this help message")
	flag.Parse()

	if *help || *helpLong || len(os.Args) == 1 {
		fmt.Println("Usage: fibby [options] arg")
		fmt.Println("Options:")
		fmt.Println("  -h, --help   Show this help message")
		fmt.Println("Args:")
		fmt.Println("  An valid int32 number denoting the sequence number you would like to return.")
		return
	}

	seqTopInt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fibSeq(seqTopInt))

}
