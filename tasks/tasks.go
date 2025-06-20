package tasks

import (
	"errors"
	"time"
)

type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    Priority    int       `json:"priority"`
    CreatedAt   time.Time `json:"created_at"`
}

type TaskList struct {
    Tasks []Task `json:"tasks"`
}

// AddTask menambahkan tugas baru ke daftar
func (tl *TaskList) AddTask(description string, priority int) error {
    if description == "" {
        return errors.New("description cannot be empty")
    }
    if priority < 1 || priority > 3 {
        return errors.New("priority must be 1 (high), 2 (medium), or 3 (low)")
    }
    newID := 1
    if len(tl.Tasks) > 0 {
        newID = tl.Tasks[len(tl.Tasks)-1].ID + 1
    }
    task := Task{
        ID:          newID,
        Description: description,
        Completed:   false,
        Priority:    priority,
        CreatedAt:   time.Now(),
    }
    tl.Tasks = append(tl.Tasks, task)
    return nil
}

// CompleteTask menandai tugas sebagai selesai berdasarkan ID
func (tl *TaskList) CompleteTask(id int) bool {
    for i, task := range tl.Tasks {
        if task.ID == id {
            tl.Tasks[i].Completed = true
            return true
        }
    }
    return false
}

// DeleteTask menghapus tugas berdasarkan ID
func (tl *TaskList) DeleteTask(id int) bool {
    for i, task := range tl.Tasks {
        if task.ID == id {
            tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
            return true
        }
    }
    return false
}

// ListTasks mengembalikan tugas berdasarkan status
func (tl *TaskList) ListTasks(showCompleted bool, showPending bool) []Task {
    var result []Task
    for _, task := range tl.Tasks {
        if showCompleted && task.Completed {
            result = append(result, task)
        }
        if showPending && !task.Completed {
            result = append(result, task)
        }
    }
    return result
}

// EditTask mengubah deskripsi tugas berdasarkan ID
func (tl *TaskList) EditTask(id int, description string) error {
    if description == "" {
        return errors.New("description cannot be empty")
    }
    for i, task := range tl.Tasks {
        if task.ID == id {
            tl.Tasks[i].Description = description
            return nil
        }
    }
    return errors.New("task not found")
}

// SetPriority mengatur prioritas tugas berdasarkan ID
func (tl *TaskList) SetPriority(id int, priority int) error {
    if priority < 1 || priority > 3 {
        return errors.New("priority must be 1 (high), 2 (medium), or 3 (low)")
    }
    for i, task := range tl.Tasks {
        if task.ID == id {
            tl.Tasks[i].Priority = priority
            return nil
        }
    }
    return errors.New("task not found")
}