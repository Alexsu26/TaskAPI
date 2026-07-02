package worker

import "log/slog"

type TaskCreatedEvent struct {
	TaskID int64
	UserID int64
	Title  string
}

type Worker struct {
	Events chan TaskCreatedEvent
	Log    *slog.Logger
	done   chan struct{}
}

func New(log *slog.Logger, bufferSize int) *Worker {
	return &Worker{
		Events: make(chan TaskCreatedEvent, bufferSize),
		Log:    log,
		done:   make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go func() {
		defer close(w.done)
		for event := range w.Events {
			w.Log.Info("task created event processed",
				"task_id", event.TaskID,
				"user_id", event.UserID,
				"title", event.Title)
		}
	}()
}

func (w *Worker) Stop() {
	close(w.Events)
	<-w.done
}

func (w *Worker) PublishTaskCreated(event TaskCreatedEvent) {
	select {
	case w.Events <- event:
	default:
		w.Log.Warn("task created event dropped")
	}
}
