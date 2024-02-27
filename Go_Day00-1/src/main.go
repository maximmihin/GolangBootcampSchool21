package main

import (
	"apg1_00/anscombe"
	"apg1_00/parse_input"
	"fmt"
	"log"
)

func main() {

	mean, median, mode, ds := parse_input.ParsFlags()

	numsArr, err := parse_input.ParsTerm()
	if err != nil {
		log.Fatal(err)
		return
	}

	if mean {
		fmt.Println("Mean:", anscombe.CalcMean(numsArr))
	}

	if median {
		fmt.Println("Median:", anscombe.CalcMedian(numsArr))
	}

	if mode {
		fmt.Println("Mode:", anscombe.CalcMode(numsArr))
	}

	if ds {
		fmt.Println("DS:", anscombe.CalcDS(numsArr))
	}
}
