package main

import (
	"testing"
	"time"
)

type BenchResult struct {
	Name  string
	Loops []BenchEntry
	Avg   BenchEntry
}

func (br *BenchResult) CountAvg() *BenchResult {
	for i := range br.Loops {
		br.Avg.Create += br.Loops[i].Create
		br.Avg.Pass += br.Loops[i].Pass
		br.Avg.PassCnt = br.Loops[i].PassCnt
	}
	n := time.Duration(len(br.Loops))
	br.Avg.Create /= n
	br.Avg.Pass /= n
	return br
}

type BenchEntry struct {
	Create  time.Duration
	PassCnt int
	Pass    time.Duration
}

var PASS_DEPTH = []int{2, 4, 6, 8, 16, 24, 32}

func Bench[T any](b *testing.B, cnt int, name string, create func() T, pass func(T) T) []*BenchResult {
	b.StopTimer()
	b.ResetTimer()

	final_res := make([]*BenchResult, len(PASS_DEPTH))
	for i, depth := range PASS_DEPTH {
		res := &BenchResult{
			Name:  name,
			Loops: make([]BenchEntry, cnt),
		}

		for n := 0; n < cnt; n++ {
			e := BenchEntry{}

			b.StartTimer()
			v := create()
			b.StopTimer()
			e.Create = b.Elapsed()

			b.ResetTimer()

			b.StartTimer()
			for i := 0; i < depth; i++ {
				v = pass(v)
			}
			b.StopTimer()
			e.Pass = b.Elapsed()
			e.PassCnt = depth
		}

		res = res.CountAvg()

		final_res[i] = res
	}
	return final_res

}
