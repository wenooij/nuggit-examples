//go:build ycharts

package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/wenooij/nuggit"
	"github.com/wenooij/nuggit-spec/charts/ycharts"
	"github.com/wenooij/nuggit/runtime"
	"github.com/wenooij/nuggit/v1alpha"
	"golang.org/x/net/html"
)

func main() {
	flag.Parse()

	g := ycharts.IndicatorChromedp()

	r := &runtime.StageRunner{
		Graphs:  []*nuggit.Graph{g},
		Factory: v1alpha.Factory{},
	}
	if err := r.Run(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println(strings.TrimSpace(r.Results[0]["out"].(*html.Node).FirstChild.Data))
}
