//go:build ycharts

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/wenooij/nuggit"
	"github.com/wenooij/nuggit-spec/charts/ycharts"
	"github.com/wenooij/nuggit/runtime"
	"github.com/wenooij/nuggit/v1alpha"
)

func main() {
	flag.Parse()

	g := ycharts.IndicatorChromedp()

	r := &runtime.StageRunner{
		Graphs:  []*nuggit.Graph{g},
		Factory: &v1alpha.Factory{},
	}
	if err := r.Run(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println(r.Results[0]["out"].(string))
}
