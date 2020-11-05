package main

import (
	"math"
	"os"
	"time"

	"github.com/gosuri/uiprogress"

	"github.com/foxcapades/go-bytify/v0/tools/bench/internal"
)

const (
	cycles             = 100
	iterations         = 10_000_000
	confidence float64 = 1.96
)

var (
	i8Vals  = []int64{-128, 0, 127}
	i16Vals = append(i8Vals, -32768, 32767)
	i32Vals = append(i16Vals, -2147483648, 2147483647)
	i64Vals = append(i32Vals, -9223372036854775808, 9223372036854775807)

	bfFuncs = []*internal.ITestConfig{
		{"Int8ToBytes", internal.Int8ToBytes, i8Vals, internal.StdLibIntToBytes},
		{"Int16ToBytes", internal.Int16ToBytes, i16Vals, internal.StdLibIntToBytes},
		{"Int32ToBytes", internal.Int32ToBytes, i32Vals, internal.StdLibIntToBytes},
		{"Int64ToBytes", internal.Int64ToBytes, i64Vals, internal.StdLibIntToBytes},
		{"Int8ToByteSlice", internal.Int8ToByteSlice, i8Vals, internal.StdLibIntToByteSlice},
		{"Int16ToByteSlice", internal.Int16ToByteSlice, i16Vals, internal.StdLibIntToByteSlice},
		{"Int32ToByteSlice", internal.Int32ToByteSlice, i32Vals, internal.StdLibIntToByteSlice},
		{"Int64ToByteSlice", internal.Int64ToByteSlice, i64Vals, internal.StdLibIntToByteSlice},
	}
)

func main() {
	uiprogress.Start()

	for _, cn := range bfFuncs {
		name := cn.Name
		out, err := os.Create("docs/benchmarks/" + name + ".txt")
		if err != nil {
			panic(err)
		}

		bar := uiprogress.AddBar(cycles)
		bar.AppendCompleted()
		bar.PrependElapsed()
		bar.Width = 100
		bar.PrependFunc(func(*uiprogress.Bar) string { return name })

		rep := new(internal.Report)
		rep.Iterations = iterations
		rep.Bytify.Name = "bytify." + name
		rep.StdLib.Name = "strconv.AppendInt"

		results := [2][cycles]time.Duration{}
		inputs := cn.Inputs.([]int64)
		ln := len(inputs)

		for cycle := 0; cycle < cycles; cycle++ {
			v := inputs[cycle%ln]
			results[0][cycle] = cn.Func(v, iterations)
			results[1][cycle] = cn.Competitor(v, iterations)
			bar.Incr()
		}

		raw := [...][]float64{
			make([]float64, cycles),
			make([]float64, cycles),
		}
		vals := [2]*internal.Stats{&rep.Bytify, &rep.StdLib}

		for i := 0; i < 2; i++ {
			for j := range results[i] {
				raw[i][j] = float64(results[i][j].Nanoseconds() / iterations)
			}

			internal.StdDev(raw[i], vals[i])
			vals[i].MoE = confidence * (vals[i].StdDev / math.Sqrt(float64(cycles)))
		}

		rep.Cycles = make([]internal.Cycle, cycles)
		for i := 0; i < cycles; i++ {
			rep.Cycles[i].Tool1 = uint16(raw[0][i])
			rep.Cycles[i].Tool2 = uint16(raw[1][i])
		}

		internal.PrintReport(rep, out)

		_ = out.Close()
	}
}
