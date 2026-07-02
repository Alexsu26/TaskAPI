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
}

// Graceful shutdown is deferred to T026; this worker currently exits with the process.
func (w *Worker) Start() {
	go func() {
		for event := range w.Events {
			w.Log.Info("task created event processed",
				"task_id", event.TaskID,
				"user_id", event.UserID,
				"title", event.Title)
		}
	}()
}

func (w *Worker) PublishTaskCreated(event TaskCreatedEvent) {
	select {
	case w.Events <- event:
	default:
		w.Log.Warn("task created event dropped")
	}
}
