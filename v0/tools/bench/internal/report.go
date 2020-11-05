package internal

import (
	"fmt"
	"io"
)

const (
	Line2  = "Cycles:    "
	Line3  = "Iterations:"
	Line4  = "Metrics:"
	Line5  = "  Function %s:"
	Line6  = "    Mean:           "
	Line7  = "    Std Deviation:  "
	Line8  = "    Margin of Error:"
	Line9  = "Output (TSV):"
	Line10 = "--------------------------------------------------------------------------------"
	Line11 = "\t%s\t%s"
	Line12 = "Cycle\tAvg NS/Call\tAvg NS/Call"
	Line13 = "%d\t%d\t%d"
)

type Report struct {
	Bytify     Stats
	StdLib     Stats
	Cycles     []Cycle
	Iterations uint64
}

type Stats struct {
	Name    string
	AvgTime float64
	StdDev  float64
	MoE     float64
}

type Cycle struct {
	Tool1 uint16
	Tool2 uint16
}

func PrintReport(rep *Report, out io.Writer) {
	mustPrint(out)
	mustPrint(out, Line2, len(rep.Cycles))
	mustPrint(out, Line3, rep.Iterations)
	mustPrint(out, Line4)

	mustPrint(out, fmt.Sprintf(Line5, rep.Bytify.Name))
	mustPrint(out, Line6, rep.Bytify.AvgTime)
	mustPrint(out, Line7, rep.Bytify.StdDev)
	mustPrint(out, Line8, rep.Bytify.MoE)

	mustPrint(out, fmt.Sprintf(Line5, rep.StdLib.Name))
	mustPrint(out, Line6, rep.StdLib.AvgTime)
	mustPrint(out, Line7, rep.StdLib.StdDev)
	mustPrint(out, Line8, rep.StdLib.MoE)

	mustPrint(out, Line9)
	mustPrint(out, Line10)
	mustPrint(out, fmt.Sprintf(Line11, rep.Bytify.Name, rep.StdLib.Name))
	mustPrint(out, Line12)

	for i := range rep.Cycles {
		tmp := &rep.Cycles[i]
		mustPrint(out, fmt.Sprintf(Line13, i, tmp.Tool1, tmp.Tool2))
	}
}

func mustPrint(out io.Writer, things... interface{}) {
	if _, err := fmt.Fprintln(out, things...); err != nil {
		panic(err)
	}
}