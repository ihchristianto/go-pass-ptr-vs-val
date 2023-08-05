# go-pass-ptr-vs-val
Benchmarking Golang Pass-By-Pointer vs Pass-By-Value

Requires min Go 1.20 for `benchmark.Elapsed`
Simply run `test.sh`

## Summary thus far

## Background

There are so many discussion about whether to use struct pointer or value in Golang. Many of the discussion focused on how struct pointers is "bad" in terms of performance due to initialization in heap and garbage collecting.

But I have a feeling that pointer is "better" in performance when passed around between functions as the only thing that needs to be copied is one 64-bit integer (or whatever integer length your CPU architecture dictates) that is the pointer itself, as opposed to copying the entire struct.

## The Benchmark

To test out the hypothesis, the benchmark is done by simply passing and returning the same struct over and over again in 2 forms: pointer and value.

We will try different sizes of struct: 2, 4, 8, 16, 24, and 32 fields
Also how many times it is passed & returned around: 2, 4, 8, 16, 24 or 32 times

We will measure the time it takes for each pointer and value to be created
And we will also measure the time it takes for each pointer and value to be passed around

## Results
