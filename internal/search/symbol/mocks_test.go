// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package symbol

import (
	"context"
	"sync"

	zoekt "github.com/sourcegraph/zoekt"
	query "github.com/sourcegraph/zoekt/query"
)

// MockStreamer is a mock implementation of the Streamer interface (from the
// package github.com/sourcegraph/zoekt) used for unit testing.
type MockStreamer struct {
	// CloseFunc is an instance of a mock function object controlling the
	// behavior of the method Close.
	CloseFunc *StreamerCloseFunc
	// ListFunc is an instance of a mock function object controlling the
	// behavior of the method List.
	ListFunc *StreamerListFunc
	// SearchFunc is an instance of a mock function object controlling the
	// behavior of the method Search.
	SearchFunc *StreamerSearchFunc
	// StreamSearchFunc is an instance of a mock function object controlling
	// the behavior of the method StreamSearch.
	StreamSearchFunc *StreamerStreamSearchFunc
	// StringFunc is an instance of a mock function object controlling the
	// behavior of the method String.
	StringFunc *StreamerStringFunc
}

// NewMockStreamer creates a new mock of the Streamer interface. All methods
// return zero values for all results, unless overwritten.
func NewMockStreamer() *MockStreamer {
	return &MockStreamer{
		CloseFunc: &StreamerCloseFunc{
			defaultHook: func() {
				return
			},
		},
		ListFunc: &StreamerListFunc{
			defaultHook: func(context.Context, query.Q, *zoekt.ListOptions) (r0 *zoekt.RepoList, r1 error) {
				return
			},
		},
		SearchFunc: &StreamerSearchFunc{
			defaultHook: func(context.Context, query.Q, *zoekt.SearchOptions) (r0 *zoekt.SearchResult, r1 error) {
				return
			},
		},
		StreamSearchFunc: &StreamerStreamSearchFunc{
			defaultHook: func(context.Context, query.Q, *zoekt.SearchOptions, zoekt.Sender) (r0 error) {
				return
			},
		},
		StringFunc: &StreamerStringFunc{
			defaultHook: func() (r0 string) {
				return
			},
		},
	}
}

// NewStrictMockStreamer creates a new mock of the Streamer interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockStreamer() *MockStreamer {
	return &MockStreamer{
		CloseFunc: &StreamerCloseFunc{
			defaultHook: func() {
				panic("unexpected invocation of MockStreamer.Close")
			},
		},
		ListFunc: &StreamerListFunc{
			defaultHook: func(context.Context, query.Q, *zoekt.ListOptions) (*zoekt.RepoList, error) {
				panic("unexpected invocation of MockStreamer.List")
			},
		},
		SearchFunc: &StreamerSearchFunc{
			defaultHook: func(context.Context, query.Q, *zoekt.SearchOptions) (*zoekt.SearchResult, error) {
				panic("unexpected invocation of MockStreamer.Search")
			},
		},
		StreamSearchFunc: &StreamerStreamSearchFunc{
			defaultHook: func(context.Context, query.Q, *zoekt.SearchOptions, zoekt.Sender) error {
				panic("unexpected invocation of MockStreamer.StreamSearch")
			},
		},
		StringFunc: &StreamerStringFunc{
			defaultHook: func() string {
				panic("unexpected invocation of MockStreamer.String")
			},
		},
	}
}

// NewMockStreamerFrom creates a new mock of the MockStreamer interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockStreamerFrom(i zoekt.Streamer) *MockStreamer {
	return &MockStreamer{
		CloseFunc: &StreamerCloseFunc{
			defaultHook: i.Close,
		},
		ListFunc: &StreamerListFunc{
			defaultHook: i.List,
		},
		SearchFunc: &StreamerSearchFunc{
			defaultHook: i.Search,
		},
		StreamSearchFunc: &StreamerStreamSearchFunc{
			defaultHook: i.StreamSearch,
		},
		StringFunc: &StreamerStringFunc{
			defaultHook: i.String,
		},
	}
}

// StreamerCloseFunc describes the behavior when the Close method of the
// parent MockStreamer instance is invoked.
type StreamerCloseFunc struct {
	defaultHook func()
	hooks       []func()
	history     []StreamerCloseFuncCall
	mutex       sync.Mutex
}

// Close delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStreamer) Close() {
	m.CloseFunc.nextHook()()
	m.CloseFunc.appendCall(StreamerCloseFuncCall{})
	return
}

// SetDefaultHook sets function that is called when the Close method of the
// parent MockStreamer instance is invoked and the hook queue is empty.
func (f *StreamerCloseFunc) SetDefaultHook(hook func()) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Close method of the parent MockStreamer instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StreamerCloseFunc) PushHook(hook func()) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StreamerCloseFunc) SetDefaultReturn() {
	f.SetDefaultHook(func() {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StreamerCloseFunc) PushReturn() {
	f.PushHook(func() {
		return
	})
}

func (f *StreamerCloseFunc) nextHook() func() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StreamerCloseFunc) appendCall(r0 StreamerCloseFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StreamerCloseFuncCall objects describing
// the invocations of this function.
func (f *StreamerCloseFunc) History() []StreamerCloseFuncCall {
	f.mutex.Lock()
	history := make([]StreamerCloseFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StreamerCloseFuncCall is an object that describes an invocation of method
// Close on an instance of MockStreamer.
type StreamerCloseFuncCall struct{}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StreamerCloseFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StreamerCloseFuncCall) Results() []interface{} {
	return []interface{}{}
}

// StreamerListFunc describes the behavior when the List method of the
// parent MockStreamer instance is invoked.
type StreamerListFunc struct {
	defaultHook func(context.Context, query.Q, *zoekt.ListOptions) (*zoekt.RepoList, error)
	hooks       []func(context.Context, query.Q, *zoekt.ListOptions) (*zoekt.RepoList, error)
	history     []StreamerListFuncCall
	mutex       sync.Mutex
}

// List delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStreamer) List(v0 context.Context, v1 query.Q, v2 *zoekt.ListOptions) (*zoekt.RepoList, error) {
	r0, r1 := m.ListFunc.nextHook()(v0, v1, v2)
	m.ListFunc.appendCall(StreamerListFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the List method of the
// parent MockStreamer instance is invoked and the hook queue is empty.
func (f *StreamerListFunc) SetDefaultHook(hook func(context.Context, query.Q, *zoekt.ListOptions) (*zoekt.RepoList, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// List method of the parent MockStreamer instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StreamerListFunc) PushHook(hook func(context.Context, query.Q, *zoekt.ListOptions) (*zoekt.RepoList, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StreamerListFunc) SetDefaultReturn(r0 *zoekt.RepoList, r1 error) {
	f.SetDefaultHook(func(context.Context, query.Q, *zoekt.ListOptions) (*zoekt.RepoList, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StreamerListFunc) PushReturn(r0 *zoekt.RepoList, r1 error) {
	f.PushHook(func(context.Context, query.Q, *zoekt.ListOptions) (*zoekt.RepoList, error) {
		return r0, r1
	})
}

func (f *StreamerListFunc) nextHook() func(context.Context, query.Q, *zoekt.ListOptions) (*zoekt.RepoList, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StreamerListFunc) appendCall(r0 StreamerListFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StreamerListFuncCall objects describing the
// invocations of this function.
func (f *StreamerListFunc) History() []StreamerListFuncCall {
	f.mutex.Lock()
	history := make([]StreamerListFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StreamerListFuncCall is an object that describes an invocation of method
// List on an instance of MockStreamer.
type StreamerListFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 query.Q
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 *zoekt.ListOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *zoekt.RepoList
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StreamerListFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StreamerListFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StreamerSearchFunc describes the behavior when the Search method of the
// parent MockStreamer instance is invoked.
type StreamerSearchFunc struct {
	defaultHook func(context.Context, query.Q, *zoekt.SearchOptions) (*zoekt.SearchResult, error)
	hooks       []func(context.Context, query.Q, *zoekt.SearchOptions) (*zoekt.SearchResult, error)
	history     []StreamerSearchFuncCall
	mutex       sync.Mutex
}

// Search delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStreamer) Search(v0 context.Context, v1 query.Q, v2 *zoekt.SearchOptions) (*zoekt.SearchResult, error) {
	r0, r1 := m.SearchFunc.nextHook()(v0, v1, v2)
	m.SearchFunc.appendCall(StreamerSearchFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Search method of the
// parent MockStreamer instance is invoked and the hook queue is empty.
func (f *StreamerSearchFunc) SetDefaultHook(hook func(context.Context, query.Q, *zoekt.SearchOptions) (*zoekt.SearchResult, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Search method of the parent MockStreamer instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StreamerSearchFunc) PushHook(hook func(context.Context, query.Q, *zoekt.SearchOptions) (*zoekt.SearchResult, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StreamerSearchFunc) SetDefaultReturn(r0 *zoekt.SearchResult, r1 error) {
	f.SetDefaultHook(func(context.Context, query.Q, *zoekt.SearchOptions) (*zoekt.SearchResult, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StreamerSearchFunc) PushReturn(r0 *zoekt.SearchResult, r1 error) {
	f.PushHook(func(context.Context, query.Q, *zoekt.SearchOptions) (*zoekt.SearchResult, error) {
		return r0, r1
	})
}

func (f *StreamerSearchFunc) nextHook() func(context.Context, query.Q, *zoekt.SearchOptions) (*zoekt.SearchResult, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StreamerSearchFunc) appendCall(r0 StreamerSearchFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StreamerSearchFuncCall objects describing
// the invocations of this function.
func (f *StreamerSearchFunc) History() []StreamerSearchFuncCall {
	f.mutex.Lock()
	history := make([]StreamerSearchFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StreamerSearchFuncCall is an object that describes an invocation of
// method Search on an instance of MockStreamer.
type StreamerSearchFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 query.Q
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 *zoekt.SearchOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *zoekt.SearchResult
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StreamerSearchFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StreamerSearchFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StreamerStreamSearchFunc describes the behavior when the StreamSearch
// method of the parent MockStreamer instance is invoked.
type StreamerStreamSearchFunc struct {
	defaultHook func(context.Context, query.Q, *zoekt.SearchOptions, zoekt.Sender) error
	hooks       []func(context.Context, query.Q, *zoekt.SearchOptions, zoekt.Sender) error
	history     []StreamerStreamSearchFuncCall
	mutex       sync.Mutex
}

// StreamSearch delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockStreamer) StreamSearch(v0 context.Context, v1 query.Q, v2 *zoekt.SearchOptions, v3 zoekt.Sender) error {
	r0 := m.StreamSearchFunc.nextHook()(v0, v1, v2, v3)
	m.StreamSearchFunc.appendCall(StreamerStreamSearchFuncCall{v0, v1, v2, v3, r0})
	return r0
}

// SetDefaultHook sets function that is called when the StreamSearch method
// of the parent MockStreamer instance is invoked and the hook queue is
// empty.
func (f *StreamerStreamSearchFunc) SetDefaultHook(hook func(context.Context, query.Q, *zoekt.SearchOptions, zoekt.Sender) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// StreamSearch method of the parent MockStreamer instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *StreamerStreamSearchFunc) PushHook(hook func(context.Context, query.Q, *zoekt.SearchOptions, zoekt.Sender) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StreamerStreamSearchFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, query.Q, *zoekt.SearchOptions, zoekt.Sender) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StreamerStreamSearchFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, query.Q, *zoekt.SearchOptions, zoekt.Sender) error {
		return r0
	})
}

func (f *StreamerStreamSearchFunc) nextHook() func(context.Context, query.Q, *zoekt.SearchOptions, zoekt.Sender) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StreamerStreamSearchFunc) appendCall(r0 StreamerStreamSearchFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StreamerStreamSearchFuncCall objects
// describing the invocations of this function.
func (f *StreamerStreamSearchFunc) History() []StreamerStreamSearchFuncCall {
	f.mutex.Lock()
	history := make([]StreamerStreamSearchFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StreamerStreamSearchFuncCall is an object that describes an invocation of
// method StreamSearch on an instance of MockStreamer.
type StreamerStreamSearchFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 query.Q
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 *zoekt.SearchOptions
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 zoekt.Sender
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StreamerStreamSearchFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StreamerStreamSearchFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// StreamerStringFunc describes the behavior when the String method of the
// parent MockStreamer instance is invoked.
type StreamerStringFunc struct {
	defaultHook func() string
	hooks       []func() string
	history     []StreamerStringFuncCall
	mutex       sync.Mutex
}

// String delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStreamer) String() string {
	r0 := m.StringFunc.nextHook()()
	m.StringFunc.appendCall(StreamerStringFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the String method of the
// parent MockStreamer instance is invoked and the hook queue is empty.
func (f *StreamerStringFunc) SetDefaultHook(hook func() string) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// String method of the parent MockStreamer instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StreamerStringFunc) PushHook(hook func() string) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StreamerStringFunc) SetDefaultReturn(r0 string) {
	f.SetDefaultHook(func() string {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StreamerStringFunc) PushReturn(r0 string) {
	f.PushHook(func() string {
		return r0
	})
}

func (f *StreamerStringFunc) nextHook() func() string {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StreamerStringFunc) appendCall(r0 StreamerStringFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StreamerStringFuncCall objects describing
// the invocations of this function.
func (f *StreamerStringFunc) History() []StreamerStringFuncCall {
	f.mutex.Lock()
	history := make([]StreamerStringFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StreamerStringFuncCall is an object that describes an invocation of
// method String on an instance of MockStreamer.
type StreamerStringFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 string
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StreamerStringFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StreamerStringFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
