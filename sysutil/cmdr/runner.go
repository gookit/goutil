package cmdr

import (
	"fmt"
	"strings"

	"github.com/gookit/goutil/mathutil"
)

// Errs on run tasks. key is Task.ID
type Errs map[string]error

// Error string
func (e Errs) Error() string {
	var sb strings.Builder
	for name, err := range e {
		sb.WriteString(name)
		sb.WriteByte(':')
		sb.WriteString(err.Error())
		sb.WriteByte('\n')
	}
	return sb.String()
}

// IsEmpty error
func (e Errs) IsEmpty() bool {
	return len(e) == 0
}

// One error
func (e Errs) One() error {
	for _, err := range e {
		return err
	}
	return nil
}

// Task struct
type Task struct {
	err error

	// ID for task
	ID  string
	Cmd *Cmd

	// RunBefore hook
	RunBefore func() bool
	PrevCond  func(prev *Task) bool
}

// get task id by cmd.Name
func (t *Task) ensureID(idx int) {
	if t.ID != "" {
		return
	}

	id := t.Cmd.IDString()
	if t.Cmd.Name == "" {
		id += mathutil.String(idx)
	}
	t.ID = id
}

// Err get
func (t *Task) Err() error {
	return t.err
}

// Runner use for batch run multi task commands
type Runner struct {
	prev *Task
	// task name to index
	idMap map[string]int
	tasks []*Task
	// Errs on run tasks
	Errs Errs

	// IgnoreErr continue on error
	IgnoreErr  bool
	RunBefore  func(r *Runner) bool
	RunAfter   func(r *Runner)
	ListenPrev func(t *Task) bool
	EachBefore func(c *Cmd) bool
}

// NewRunner instance
func NewRunner() *Runner {
	return &Runner{
		idMap: make(map[string]int, 0),
		tasks: make([]*Task, 0),
		Errs:  make(Errs),
	}
}

// Add multitask at once
func (r *Runner) Add(tasks ...*Task) *Runner {
	for _, task := range tasks {
		r.AddTask(task)
	}
	return r
}

// AddTask add one task
func (r *Runner) AddTask(task *Task) *Runner {
	if task.Cmd == nil {
		panic("task command cannot be empty")
	}

	idx := len(r.tasks)
	task.ensureID(idx)

	// TODO check id repeat
	r.idMap[task.ID] = idx
	r.tasks = append(r.tasks, task)
	return r
}

// AddCmd commands
func (r *Runner) AddCmd(cmds ...*Cmd) *Runner {
	for _, cmd := range cmds {
		r.AddTask(&Task{Cmd: cmd})
	}
	return r
}

// Run all tasks
func (r *Runner) Run() error {
	if r.RunBefore != nil {
		if ok := r.RunBefore(r); !ok {
			return nil
		}
	}

	if r.RunAfter != nil {
		r.RunAfter(r)
	}

	// to run tasks
	r.runTasks()

	if r.Errs.IsEmpty() {
		return nil
	}
	return r.Errs
}

// Prev task instance after running
func (r *Runner) runTasks() {
	for _, task := range r.tasks {
		if task.RunBefore != nil && !task.RunBefore() {
			continue
		}

		if r.prev != nil {
			if r.ListenPrev != nil && !r.ListenPrev(r.prev) {
				continue
			}

			if task.PrevCond != nil && !task.PrevCond(r.prev) {
				continue
			}
		}

		cmd := task.Cmd
		if r.EachBefore != nil && r.EachBefore(cmd) {
			continue
		}

		// do running
		if err := cmd.Run(); err != nil {
			task.err = err
			r.Errs[task.ID] = err

			// not ignore error, stop.
			if !r.IgnoreErr {
				return
			}
		}

		if r.EachBefore != nil && !r.EachBefore(cmd) {
			continue
		}

		// store prev
		r.prev = task
	}
}

// TaskIDs get
func (r *Runner) TaskIDs() []string {
	ss := make([]string, 0, len(r.idMap))
	for id, _ := range r.idMap {
		ss = append(ss, id)
	}
	return ss
}

// Prev task instance after running
func (r *Runner) Prev() *Task {
	return r.prev
}

// Task get by id name
func (r *Runner) Task(id string) (*Task, error) {
	if idx, ok := r.idMap[id]; ok {
		return r.tasks[idx], nil
	}
	return nil, fmt.Errorf("task %q is not exists", id)
}
