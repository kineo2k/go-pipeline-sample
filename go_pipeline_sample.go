package main

import (
	"embed"
	"go-pipeline-sample/server"
	"go-pipeline-sample/tm"
)

//go:embed statics
var statics embed.FS

func main() {
	tm.Init(statics)

	server.EchoStart(7063)
}
