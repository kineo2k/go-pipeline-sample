package stages

import (
	"fmt"
	"github.com/disintegration/imaging"
	"go-pipeline-sample/service/pipeline/task"
)

type Effect struct {
}

func NewEffect() *Effect {
	return &Effect{}
}

func (e *Effect) Process(in <-chan *task.Task) <-chan *task.Task {
	out := make(chan *task.Task, cap(in))

	go func() {
		defer close(out)
		for t := range in {
			fmt.Println("Effect Stage")

			if e.effect(t) {
				fmt.Println("  - Image Effect Success")
			} else {
				fmt.Println("  - Image Effect Passed")
			}

			out <- t
		}
	}()

	return out
}

func (e *Effect) effect(task *task.Task) bool {
	if task.Spec.Effect.Type == "none" {
		return false
	}

	switch task.Spec.Effect.Type {
	case "blur":
		task.Img = imaging.Blur(task.Img, 4)
	case "sharpening":
		task.Img = imaging.Sharpen(task.Img, 4)
	case "brightness":
		task.Img = imaging.AdjustBrightness(task.Img, 30)
	}

	return true
}
