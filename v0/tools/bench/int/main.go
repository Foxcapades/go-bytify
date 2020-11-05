package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/foxcapades/go-bytify/v0/bytify"
)

const (
	cycles     = 100
	iterations = 10_000_000

	bufferSize = 6

	confidence float64 = 1.96
)

var inputValue int8 = -1

func main() {
	funcs := [...]string{
		"bytify.Int8ToBytes",
		"strconv.AppendInt",
	}

	results := [2][cycles]time.Duration{}

	for cycle := 0; cycle < cycles; cycle++ {
		results[0][cycle] = Test1()
		fmt.Print(".")
		results[1][cycle] = Test2()
		fmt.Print(".")
	}

	fmt.Println()
	fmt.Println("Report:")
	fmt.Println("  Cycles:    ", cycles)
	fmt.Println("  Iterations:", iterations)
	fmt.Println("  Metrics:")

	raw := [2][cycles]float64{}
	vals := [2]values{}
	for i := 0; i < 2; i++ {
		fmt.Println("   ", funcs[i], ":")

		for j := range results[i] {
			raw[i][j] = float64(results[i][j].Nanoseconds() / iterations)
		}

		vals[i] = stdDev(raw[i])

		fmt.Println("      Mean:           ", vals[i].mean)
		fmt.Println("      Std Deviation:  ", vals[i].stdDev)
		fmt.Println("      Margin of Error:", marginOfError(&vals[i]))
	}

	fmt.Println("  Cycles (TSV):")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("\t%s\t%s", funcs[0], funcs[1]))
	fmt.Println("Cycle\tAvg NS Per Call\tAvg NS Per Call")

	for i := 0; i < cycles; i++ {
		fmt.Println(fmt.Sprintf("%d\t%.0f\t%.0f", i,
			raw[0][i],
			raw[1][i],
		))
	}
}

func Test1() time.Duration {
	buffer := make([]byte, bufferSize)

	start := time.Now()
	for i := 0; i < iterations; i++ {
		bytify.Int8ToBytes(inputValue, buffer)
	}
	return time.Now().Sub(start)
}

func Test2() time.Duration {
	buffer := make([]byte, bufferSize)

	start := time.Now()
	for i := 0; i < iterations; i++ {
		strconv.AppendInt(buffer, int64(inputValue), 10)
	}
	return time.Now().Sub(start)
}

////////////////////////////////////////////////////////////////////////////////

type values struct {
	mean   float64
	stdDev float64
}

func stdDev(vals [cycles]float64) (out values) {
	var n, sum, ssq float64

	n = float64(cycles)

	for i := range vals {
		sum += vals[i]
		ssq += vals[i] * vals[i]
	}

	out.mean = sum / n
	out.stdDev = math.Sqrt(ssq/(n-1) - out.mean*out.mean)

	return
}

func marginOfError(v *values) float64 {
	n := float64(cycles)

	return confidence * (v.stdDev / math.Sqrt(n))
}
