package stages

import (
	"fmt"
	"github.com/disintegration/imaging"
	"go-pipeline-sample/service/pipeline/task"
)

type Resize struct {
}

func NewResize() *Resize {
	return &Resize{}
}

func (r *Resize) Process(in <-chan *task.Task) <-chan *task.Task {
	out := make(chan *task.Task, cap(in))

	go func() {
		defer close(out)
		for t := range in {
			fmt.Println("Resize Stage")

			r.resize(t)

			fmt.Println("  - Image Resize Success")

			out <- t
		}
	}()

	return out
}

func (r *Resize) resize(task *task.Task) {
	if task.Spec.Resize.KeepAspectRatio {
		if task.Spec.Resize.Width != 0 {
			task.Img = imaging.Resize(task.Img, task.Spec.Resize.Width, 0, imaging.Lanczos)
		} else {
			task.Img = imaging.Resize(task.Img, 0, task.Spec.Resize.Height, imaging.Lanczos)
		}
	} else {
		task.Img = imaging.Resize(task.Img, task.Spec.Resize.Width, task.Spec.Resize.Height, imaging.Lanczos)
	}
}
