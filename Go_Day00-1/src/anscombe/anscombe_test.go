package anscombe

import "testing"

func TestCalcMean(t *testing.T) {
	numsArr := []int{1, 2, 3}
	mean := CalcMean(numsArr)
	if mean != 2 {
		t.Fatal()
	}
}

func TestCalcMedianOddNums(t *testing.T) {
	numsArr := []int{1, 2, 3}
	median := CalcMedian(numsArr)
	if median != 2 {
		t.Fatal()
	}
}

func TestCalcMedianEvenNums(t *testing.T) {
	numsArr := []int{1, 2, 3, 4}
	median := CalcMedian(numsArr)
	if median != 2.5 {
		t.Fatal()
	}
}

func TestCalcMode(t *testing.T) {
	numsArr := []int{1, 2, 3, 5, 5, 5, 6, 6, 6}
	mode := CalcMode(numsArr)
	if mode != 5 {
		t.Fatal()
	}
}

func TestCalcMode2(t *testing.T) {
	numsArr := []int{1}
	mode := CalcMode(numsArr)
	if mode != 1 {
		t.Fatal()
	}
}
