package task

import vec2d "github.com/flywave/go3d/float64/vec2"

type ProgressStore interface {
	Store(id string, progress interface{})
	Get(id string) interface{}
}

type ProgressLogger interface {
	LogMessage(msg string)
	LogStep(progress *TaskProgress)
	LogProgress(seed *TaskProgress, level int, bbox vec2d.Rect, tiles int)
	SetCurrentTaskId(id string)
	GetStore() ProgressStore
}
