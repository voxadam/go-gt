// Solve the Quadratic Assignment problem using simulated annealing

package main

import (
	"flag"
	"fmt"
	"os"
	.	"gt"
)

var verbose bool

func main() {
	k := flag.Int("k", 1, "Number of resolutions")
	m := flag.Int("m", 1000, "Number of iterations")
	ver := flag.Bool("v", false, "Verbose")
	flag.Parse()
	verbose = *ver
	in := os.Stdin
	if len(os.Args) > 1 {
		file := flag.Arg(0)
		var err os.Error
		in, err = os.Open(file)
		if in == nil {
			fmt.Printf("can't open file %s: %s\n", file, err.String())
			return
		}
	}
	n, a, b := Load(in)
	in.Close()
	v := make(Vector, n)
	Perm(v)
	for i := 0; i < *k; i++ {
		QAP_SolveSA(a, b, v, *m)
	}
	v.Print()
}