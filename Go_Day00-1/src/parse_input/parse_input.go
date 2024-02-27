package parse_input

import (
	"bufio"
	"flag"
	"os"
	"strconv"
)

func ParsFlags() (mean, median, mode, ds bool) {
	flag.BoolVar(&mean, "mean", false, "")
	flag.BoolVar(&median, "median", false, "")
	flag.BoolVar(&mode, "mode", false, "")
	flag.BoolVar(&ds, "ds", false, "")

	flag.Parse()

	if mean == false && median == false && mode == false && ds == false {
		mean, median, mode, ds = true, true, true, true
	}

	return
}

func ParsTerm() ([]int, error) {
	var nums []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}
