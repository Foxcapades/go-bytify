package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type data struct {
	function function
	input    input
	count    uint
	sum      float64
}

type function = string
type input = string

type dataMap map[input]map[function]*data

var patt = regexp.MustCompile("[^\\d]")

func main() {
	files, err := filepath.Glob("docs/benchmarks/*.tsv")
	if err != nil {
		panic(err)
	}

	data := make(dataMap, 10)

	for _, file := range files {
		processFile(file, data)
	}

	inputs := make([]string, 0, len(data))
	for k := range data {
		inputs = append(inputs, k)
	}
	sort.Slice(inputs, func(i, j int) bool {
		if inputs[i][0] == 'Z' {
			return true
		}
		if inputs[j][0] == 'Z' {
			return false
		}

		if inputs[i][3] == 'U' && inputs[j][3] == 'I' {
			return false
		}
		if inputs[i][3] == 'I' && inputs[j][3] == 'U' {
			return true
		}

		a, err := strconv.Atoi(patt.ReplaceAllString(inputs[i], ""))
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(patt.ReplaceAllString(inputs[j], ""))
		if err != nil {
			panic(err)
		}

		return a < b
	})

	fmt.Println("Function\tInput\tAvg Ns per Op\tAvg Ops per Sec")
	for _, input := range inputs {
		funcs := make([]string, 0, len(data[input]))
		for fn := range data[input] {
			funcs = append(funcs, fn)
		}
		sort.Slice(funcs, func(i, j int) bool {
			if funcs[i][0] == 'A' && funcs[j][0] == 'A' {
				if funcs[i][6] == 'U' && funcs[j][6] == 'I' {
					return false
				}
				if funcs[i][6] == 'I' && funcs[j][6] == 'U' {
					return true
				}
			}
			if funcs[i][0] == 'A' {
				return true
			}
			if funcs[j][0] == 'A' {
				return false
			}

			a, err := strconv.Atoi(patt.ReplaceAllString(funcs[i], ""))
			if err != nil {
				panic(err)
			}

			b, err := strconv.Atoi(patt.ReplaceAllString(funcs[j], ""))
			if err != nil {
				panic(err)
			}

			return a < b
		})

		for _, fn := range funcs {
			dat := data[input][fn]
			ns := dat.sum / float64(dat.count)
			fmt.Printf("%s\t%s\t%.2f\t%.2f\r\n", fn, input, ns, float64(time.Second)/ns)
		}
	}
}

const (
	funcPos  = 0
	inputPos = 1
	timePos  = 3
)

func processFile(path string, dataMap dataMap) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		txt := scan.Text()
		if len(txt) == 0 {
			continue
		}

		split := strings.Split(txt, "\t")

		time, err := strconv.ParseFloat(split[timePos], 64)
		if err != nil {
			panic(err)
		}

		if _, ok := dataMap[split[inputPos]]; !ok {
			dataMap[split[inputPos]] = make(map[function]*data, 10)
		}

		fn := split[funcPos][9:]

		if dat, ok := dataMap[split[inputPos]][fn]; ok {
			dat.count++
			dat.sum += time
		} else {
			dataMap[split[inputPos]][fn] = &data{
				function: fn,
				input:    split[inputPos],
				count:    1,
				sum:      time,
			}
		}
	}

	if scan.Err() != nil {
		panic(err)
	}
}
