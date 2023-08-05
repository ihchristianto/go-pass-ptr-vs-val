package main

import (
	"fmt"
	"os"
	"testing"
)

const COUNT = 10

func Benchmark(b *testing.B) {
	allResults := [][]*BenchResult{
		Bench(b, COUNT, "XSmallPtr", CreateXSmallPtr, PassXSmallPtr),
		Bench(b, COUNT, "XSmallVal", CreateXSmallVal, PassXSmallVal),
		Bench(b, COUNT, "SmallPtr", CreateSmallPtr, PassSmallPtr),
		Bench(b, COUNT, "SmallVal", CreateSmallVal, PassSmallVal),
		Bench(b, COUNT, "MediumPtr", CreateMediumPtr, PassMediumPtr),
		Bench(b, COUNT, "MediumVal", CreateMediumVal, PassMediumVal),
		Bench(b, COUNT, "LargePtr", CreateLargePtr, PassLargePtr),
		Bench(b, COUNT, "LargeVal", CreateLargeVal, PassLargeVal),
		Bench(b, COUNT, "XLargePtr", CreateXLargePtr, PassXLargePtr),
		Bench(b, COUNT, "XLargeVal", CreateXLargeVal, PassXLargeVal),
		Bench(b, COUNT, "XXLargePtr", CreateXXLargePtr, PassXXLargePtr),
		Bench(b, COUNT, "XXLargeVal", CreateXXLargeVal, PassXXLargeVal),
	}
	saveCsv(allResults)
}

// name, avg create, pass count, avg pass
const ROW_FMT = "%s,%d nsec,%d pass,%d nsec\n"
const FILENAME = "go_pass_ptr_vs_val_test.csv"

func saveCsv(allResults [][]*BenchResult) {
	csv := "Test Name,Avg Create,Pass Count,Avg Pass\n"
	for _, bySizeRes := range allResults {
		for _, res := range bySizeRes {
			csv += fmt.Sprintf(ROW_FMT, res.Name, res.Avg.Create, res.Avg.PassCnt, res.Avg.Pass)
		}
	}
	os.WriteFile(FILENAME, []byte(csv), 0644)
}
