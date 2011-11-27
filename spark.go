// Sparklines for your command line
// Inspired by https://github.com/holman/spark
// Copyright (c) 2011 Damian Gryski <damian@gryski.com>
// Licensed under the GPLv3 or, at your option, any later version

package main

import (
	"bufio"
	"flag"
	"math"
	"os"
	"strconv"
	"strings"
)

func graph(vals []float32) string {

	max := float32(0.0)
	min := float32(math.MaxFloat32)

	for _, v := range vals {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}

	graph := make([]string, 0, len(vals))
	// the last box doesn't look great on my mac :(
	bars := []string{"__", "▁ ", "▂ ", "▃ ", "▄ ", "▅ ", "▆ ", "▇ ", "█ "}
	scale := float32(len(bars)-1) / (max - min)
	for _, v := range vals {
		h := int(scale * float32(v-min))
		graph = append(graph, bars[h])
	}

	return strings.Join(graph, "")
}

func main() {

	flag.Parse()

	var vals []float32

	if flag.NArg() != 0 {
		// scan from parameters
		for _, s := range flag.Args() {
			fl, _ := strconv.Atof32(s)
			vals = append(vals, fl)
		}
	} else {
		// scan from stdin
		br := bufio.NewReader(os.Stdin)
		for {
			s, err := br.ReadString('\n')
			if s == "" && err != nil {
				break
			}
			fields := strings.Fields(s)

			for _, f := range fields {
				fl, _ := strconv.Atof32(f)
				vals = append(vals, fl)
			}
		}
	}

	os.Stdout.Write([]byte("|" + graph(vals) + "|\n"))
}
