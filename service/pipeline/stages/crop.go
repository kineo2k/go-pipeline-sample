package stages

import (
	"fmt"
	"github.com/disintegration/imaging"
	"go-pipeline-sample/service/pipeline/task"
)

type Crop struct {
}

func NewCrop() *Crop {
	return &Crop{}
}

func (c *Crop) Process(in <-chan *task.Task) <-chan *task.Task {
	out := make(chan *task.Task, cap(in))

	go func() {
		defer close(out)
		for t := range in {
			fmt.Println("Crop Stage")

			if c.crop(t) {
				fmt.Println("  - Image Crop Success")
			} else {
				fmt.Println("  - Image Crop Passed")
			}

			out <- t
		}
	}()

	return out
}

func (c *Crop) crop(task *task.Task) bool {
	if task.Spec.Crop.Anchor == "none" {
		return false
	}

	cropSize := 0
	if task.Img.Bounds().Size().X < task.Img.Bounds().Size().Y {
		cropSize = task.Img.Bounds().Size().X
	} else {
		cropSize = task.Img.Bounds().Size().Y
	}

	switch task.Spec.Crop.Anchor {
	case "top":
		task.Img = imaging.CropAnchor(task.Img, cropSize, cropSize, imaging.Top)
	case "bottom":
		task.Img = imaging.CropAnchor(task.Img, cropSize, cropSize, imaging.Bottom)
	case "center":
		task.Img = imaging.CropCenter(task.Img, cropSize, cropSize)
	}

	return true
}
