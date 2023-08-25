//go:build ychartsdot
//go:build !windows

package main

import (
	"flag"
	"fmt"

	"github.com/wenooij/nuggit-spec/charts/ycharts"
	"github.com/wenooij/nuggit/graphs"
	"github.com/wenooij/nuggit/graphviz"
)

func main() {
	flag.Parse()

	g := ycharts.IndicatorChromedp()

	gz := graphviz.Grapher{Graph: graphs.FromGraph(g)}
	data, err := gz.DOT()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
