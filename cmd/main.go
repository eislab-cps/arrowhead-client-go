package main

import (
	"github.com/eislab-cps/arrowhead-client-go/internal/cli"
	"github.com/eislab-cps/arrowhead-client-go/pkg/build"
)

var (
	BuildVersion string = ""
	BuildTime    string = ""
)

func main() {
	build.BuildVersion = BuildVersion
	build.BuildTime = BuildTime
	cli.Execute()
}
