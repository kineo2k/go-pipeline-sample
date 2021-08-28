package pipeline

import (
	"fmt"
	"go-pipeline-sample/service/pipeline/stages"
	"go-pipeline-sample/service/pipeline/task"
	"go-pipeline-sample/service/spec"
	"math"
	"math/rand"
	"time"
)

const maxQueueTask = 4

var uniqueInstance *executor

type executor struct {
	Queue    chan *task.Task
	Pipeline []Stage
}

func GetInstance() *executor {
	if uniqueInstance == nil {
		uniqueInstance = new(executor)
		uniqueInstance.Queue = make(chan *task.Task, maxQueueTask)
		uniqueInstance.Pipeline = []Stage{
			stages.NewInput(),
			stages.NewResize(),
			stages.NewCrop(),
			stages.NewEffect(),
			stages.NewOutput(),
		}
	}

	return uniqueInstance
}

func (e *executor) Start() {
	go func() {
		for {
			select {
			case t := <-e.Queue:
				in := make(chan *task.Task, 1)
				in <- t
				close(in)

				e.processPipeline(in)
			}
		}
	}()
}

func (e *executor) Enqueue(spec *spec.Spec) <-chan *task.Result {
	ticket := make(chan *task.Result)
	e.Queue <- &task.Task{
		Spec:     spec,
		Filename: e.randInt(),
		Img:      nil,
		Ticket:   ticket,
	}

	return ticket
}

func (e *executor) processPipeline(in <-chan *task.Task) {
	var nextChannel <-chan *task.Task
	for _, pipe := range e.Pipeline {
		if nextChannel == nil {
			nextChannel = pipe.Process(in)
		} else {
			nextChannel = pipe.Process(nextChannel)
		}
	}
}

func (e *executor) randInt() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return fmt.Sprintf("%d", r.Intn(math.MaxInt))
}
