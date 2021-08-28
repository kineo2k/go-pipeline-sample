package pipeline

import "go-pipeline-sample/service/pipeline/task"

type Stage interface {
	Process(in <-chan *task.Task) <-chan *task.Task
}
