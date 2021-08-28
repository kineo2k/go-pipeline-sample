package task

import (
	"fmt"
	"go-pipeline-sample/service/spec"
	"image"
)

type Task struct {
	Spec     *spec.Spec
	Filename string
	Img      image.Image
	Ticket   chan<- *Result
}

func (t *Task) InputPath() string {
	return fmt.Sprintf("./in_%s.jpg", t.Filename)
}

func (t *Task) OutputPath() string {
	return fmt.Sprintf("./out_%s.jpg", t.Filename)
}
