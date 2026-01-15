package syncs

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// ErrGroup is a collection of goroutines working on subtasks that
// are part of the same overall task.
//
// Refers:
//
//	https://github.com/neilotoole/errgroup
//	https://github.com/fatih/semgroup
type ErrGroup struct {
	*errgroup.Group
}

// NewCtxErrGroup instance
func NewCtxErrGroup(ctx context.Context, limit ...int) (*ErrGroup, context.Context) {
	egg, ctx1 := errgroup.WithContext(ctx)
	if len(limit) > 0 && limit[0] > 0 {
		egg.SetLimit(limit[0])
	}

	eg := &ErrGroup{Group: egg}
	return eg, ctx1
}

// NewErrGroup instance
func NewErrGroup(limit ...int) *ErrGroup {
	eg := &ErrGroup{Group: new(errgroup.Group)}

	if len(limit) > 0 && limit[0] > 0 {
		eg.SetLimit(limit[0])
	}
	return eg
}

// Add one or more handler at once
func (g *ErrGroup) Add(handlers ...func() error) {
	for _, handler := range handlers {
		g.Go(handler)
	}
}

// Go add a handler function. if panic occurs, will catch it and return as error
func (g *ErrGroup) Go(fn func() error) {
	// Group.Go: 调用给定函数的新流程。
	// 第一次“离开”的调用必须在等待之前发生。它会一直阻塞，直到新 goroutine 能够被添加，而该组中活跃 goroutine 数量不超过配置限制。
	// 第一次返回非零错误的调用会取消该组的上下文，如果该组是由调用 WithContext 创建的。错误将由等待返回。
	g.Group.Go(func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic recover: %v", r)
			}
		}()
		err = fn()
		return
	})
}

// TryGo calls the given function in a new goroutine only if the number of
// active goroutines in the group is currently below the configured limit.
//
// The return value reports whether the goroutine was started.
//
// If panic occurs, will catch it and return as error
func (g *ErrGroup) TryGo(fn func() error) bool {
	// 只有当组中当前活跃的 goroutine 数量低于配置限制时，才会在新的 Goroutine 中调用该函数。
	// 返回值报告是否启动了goroutine。
	return g.Group.TryGo(func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic recover: %v", r)
			}
		}()
		err = fn()
		return
	})
}
