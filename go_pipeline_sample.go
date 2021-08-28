package main

import (
	"embed"
	"go-pipeline-sample/server"
	"go-pipeline-sample/service/pipeline"
	"go-pipeline-sample/tm"
)

//go:embed statics
var statics embed.FS

func main() {
	tm.Init(statics)

	pipeline.GetInstance().Start()

	server.EchoStart(7063)
}
