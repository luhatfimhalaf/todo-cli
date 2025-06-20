package tasks

import (
	"testing"
)

func TestAddTask(t *testing.T) {
    tl := &TaskList{}
    err := tl.AddTask("Test Task", 1)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if len(tl.Tasks) != 1 {
        t.Errorf("Expected 1 task, got %d", len(tl.Tasks))
    }
    if tl.Tasks[0].Description != "Test Task" || tl.Tasks[0].Priority != 1 {
        t.Errorf("Task details incorrect: got %v", tl.Tasks[0])
    }
}