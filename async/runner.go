/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package async

import (
	"sync"
)

// Runner is an abstraction to make it easy to start and stop groups of things that can be
// described by a single function which waits on a channel close to exit.
type Runner struct {
	lock      sync.Mutex
	loopFuncs []func(stop <-chan struct{})
	stop      <-chan struct{}
	innerStop chan struct{}
}

// NewRunner makes a runner for the given function(s). The function(s) should loop until
// the channel is closed.
func NewRunner(f ...func(stop <-chan struct{})) *Runner {
	return &Runner{loopFuncs: f}
}

// 注意 使用 StopChan 的时候, 此时 Stop 由外部控制, 不能调用 Stop 函数
func (r *Runner) WithStopChan(stop <-chan struct{}) *Runner {
	r.stop = stop
	return r
}

// Add function
func (r *Runner) Add(f ...func(stop <-chan struct{})) *Runner {
	r.loopFuncs = append(r.loopFuncs, f...)
	return r
}

// Start begins running.
func (r *Runner) Start() {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.innerStop == nil {
		r.innerStop = make(chan struct{})
		stop := r.stop
		if stop == nil {
			stop = r.innerStop
		}
		for i := range r.loopFuncs {
			go r.loopFuncs[i](stop)
		}
	}
}

// Stop stops running.
func (r *Runner) Stop() {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.innerStop != nil {
		close(r.innerStop)
		r.innerStop = nil
	}
	if r.stop != nil {
		panic("cannot Call Stop, stop is controlled by outside stop chan")
	}
}
