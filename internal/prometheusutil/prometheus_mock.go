// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.

package prometheusutil

import (
	"context"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	model "github.com/prometheus/common/model"
	"sync"
)

// MockPrometheusQuerier is a mock impelementation of the PrometheusQuerier
// interface (from the package
// github.com/sourcegraph/sourcegraph/internal/prometheusutil) used for unit
// testing.
type MockPrometheusQuerier struct {
	// QueryRangeFunc is an instance of a mock function object controlling
	// the behavior of the method QueryRange.
	QueryRangeFunc *PrometheusQuerierQueryRangeFunc
}

// NewMockPrometheusQuerier creates a new mock of the PrometheusQuerier
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockPrometheusQuerier() *MockPrometheusQuerier {
	return &MockPrometheusQuerier{
		QueryRangeFunc: &PrometheusQuerierQueryRangeFunc{
			defaultHook: func(context.Context, string, v1.Range) (model.Value, v1.Warnings, error) {
				return nil, nil, nil
			},
		},
	}
}

// NewMockPrometheusQuerierFrom creates a new mock of the
// MockPrometheusQuerier interface. All methods delegate to the given
// implementation, unless overwritten.
func NewMockPrometheusQuerierFrom(i PrometheusQuerier) *MockPrometheusQuerier {
	return &MockPrometheusQuerier{
		QueryRangeFunc: &PrometheusQuerierQueryRangeFunc{
			defaultHook: i.QueryRange,
		},
	}
}

// PrometheusQuerierQueryRangeFunc describes the behavior when the
// QueryRange method of the parent MockPrometheusQuerier instance is
// invoked.
type PrometheusQuerierQueryRangeFunc struct {
	defaultHook func(context.Context, string, v1.Range) (model.Value, v1.Warnings, error)
	hooks       []func(context.Context, string, v1.Range) (model.Value, v1.Warnings, error)
	history     []PrometheusQuerierQueryRangeFuncCall
	mutex       sync.Mutex
}

// QueryRange delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockPrometheusQuerier) QueryRange(v0 context.Context, v1 string, v2 v1.Range) (model.Value, v1.Warnings, error) {
	r0, r1, r2 := m.QueryRangeFunc.nextHook()(v0, v1, v2)
	m.QueryRangeFunc.appendCall(PrometheusQuerierQueryRangeFuncCall{v0, v1, v2, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the QueryRange method of
// the parent MockPrometheusQuerier instance is invoked and the hook queue
// is empty.
func (f *PrometheusQuerierQueryRangeFunc) SetDefaultHook(hook func(context.Context, string, v1.Range) (model.Value, v1.Warnings, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// QueryRange method of the parent MockPrometheusQuerier instance inovkes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *PrometheusQuerierQueryRangeFunc) PushHook(hook func(context.Context, string, v1.Range) (model.Value, v1.Warnings, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *PrometheusQuerierQueryRangeFunc) SetDefaultReturn(r0 model.Value, r1 v1.Warnings, r2 error) {
	f.SetDefaultHook(func(context.Context, string, v1.Range) (model.Value, v1.Warnings, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *PrometheusQuerierQueryRangeFunc) PushReturn(r0 model.Value, r1 v1.Warnings, r2 error) {
	f.PushHook(func(context.Context, string, v1.Range) (model.Value, v1.Warnings, error) {
		return r0, r1, r2
	})
}

func (f *PrometheusQuerierQueryRangeFunc) nextHook() func(context.Context, string, v1.Range) (model.Value, v1.Warnings, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *PrometheusQuerierQueryRangeFunc) appendCall(r0 PrometheusQuerierQueryRangeFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of PrometheusQuerierQueryRangeFuncCall objects
// describing the invocations of this function.
func (f *PrometheusQuerierQueryRangeFunc) History() []PrometheusQuerierQueryRangeFuncCall {
	f.mutex.Lock()
	history := make([]PrometheusQuerierQueryRangeFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// PrometheusQuerierQueryRangeFuncCall is an object that describes an
// invocation of method QueryRange on an instance of MockPrometheusQuerier.
type PrometheusQuerierQueryRangeFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 v1.Range
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 model.Value
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 v1.Warnings
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c PrometheusQuerierQueryRangeFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c PrometheusQuerierQueryRangeFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}
