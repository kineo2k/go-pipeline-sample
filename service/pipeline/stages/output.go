package stages

import (
	"fmt"
	"github.com/disintegration/imaging"
	"go-pipeline-sample/service/pipeline/task"
)

type Output struct {
}

func NewOutput() *Output {
	return &Output{}
}

func (o *Output) Process(in <-chan *task.Task) <-chan *task.Task {
	out := make(chan *task.Task, cap(in))

	go func() {
		defer close(out)
		for t := range in {
			fmt.Println("Output Stage")

			err := imaging.Save(t.Img, t.OutputPath(), imaging.JPEGQuality(90))
			if err != nil {
				fmt.Println(err)
				fmt.Println("  - Image Save Failure")
			} else {
				fmt.Println("  - Image Save Success")
			}

			t.Ticket <- &task.Result{
				OutputPath: t.OutputPath(),
			}
			close(t.Ticket)

			out <- t
		}
	}()

	return out
}
