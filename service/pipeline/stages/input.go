package stages

import (
	"fmt"
	"github.com/disintegration/imaging"
	"go-pipeline-sample/service/pipeline/task"
	"image"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Process(in <-chan *task.Task) <-chan *task.Task {
	out := make(chan *task.Task, cap(in))

	go func() {
		defer close(out)
		for t := range in {
			fmt.Println("Input Stage")

			if i.download(t) {
				fmt.Println("  - Download Success")
				t.Img = i.open(t)
			} else {
				fmt.Println("  - Download Failure")
			}

			out <- t
		}
	}()

	return out
}

func (i *Input) download(task *task.Task) bool {
	url := task.Spec.Input.Url
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	dir, _ := filepath.Split(task.InputPath())
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return false
	}

	out, err := os.Create(task.InputPath())
	if err != nil {
		return false
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err == nil
}

func (i *Input) open(task *task.Task) image.Image {
	src, err := imaging.Open(task.InputPath())
	if err != nil {
		fmt.Println("  - Image Open Failure")
		return nil
	}

	fmt.Println("  - Image Open Success")

	return src
}
