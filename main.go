package main

import (
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

	seqTopInt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fibSeq(seqTopInt))

}
