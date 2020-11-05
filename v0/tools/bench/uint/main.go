package main

import (
	"math"
	"os"
	"time"

	"github.com/gosuri/uiprogress"

	"github.com/foxcapades/go-bytify/v0/tools/bench/internal"
)

const (
	cycles     = 100
	iterations = 10_000_000
	confidence float64 = 1.96
)

var (
	u8Vals = []uint64{
		0,
		0x7F,
		0xFF,
	}

	u16Vals = append(u8Vals,
		0x7FFF,
		0xFFFF,
	)

	u32Vals = append(u16Vals,
		0x7FFF_FFFF,
		0xFFFF_FFFF,
	)

	u64Vals = append(u32Vals,
		0x7FFF_FFFF_FFFF_FFFF,
		0xFFFF_FFFF_FFFF_FFFF,
	)

	bfFuncs = []*internal.TestConfig{
		{"Uint8ToBytes",  internal.Uint8ToBytes, u8Vals, internal.StdLibToBytes },
		{"Uint16ToBytes", internal.Uint16ToBytes, u16Vals, internal.StdLibToBytes },
		{"Uint32ToBytes", internal.Uint32ToBytes, u32Vals, internal.StdLibToBytes },
		{"Uint64ToBytes", internal.Uint64ToBytes, u64Vals, internal.StdLibToBytes },
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
		rep.StdLib.Name = "strconv.AppendUint"

		results := [2][cycles]time.Duration{}
		inputs := cn.Inputs.([]uint64)
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
