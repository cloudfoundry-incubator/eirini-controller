// Code generated by counterfeiter. DO NOT EDIT.
package reconcilerfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/k8s/reconciler"
	v1 "k8s.io/api/core/v1"
	v1a "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FakeEventsClient struct {
	CreateStub        func(string, *v1.Event) (*v1.Event, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 *v1.Event
	}
	createReturns struct {
		result1 *v1.Event
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1.Event
		result2 error
	}
	ListStub        func(v1a.ListOptions) (*v1.EventList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 v1a.ListOptions
	}
	listReturns struct {
		result1 *v1.EventList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *v1.EventList
		result2 error
	}
	UpdateStub        func(string, *v1.Event) (*v1.Event, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 string
		arg2 *v1.Event
	}
	updateReturns struct {
		result1 *v1.Event
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *v1.Event
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventsClient) Create(arg1 string, arg2 *v1.Event) (*v1.Event, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 *v1.Event
	}{arg1, arg2})
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventsClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeEventsClient) CreateCalls(stub func(string, *v1.Event) (*v1.Event, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeEventsClient) CreateArgsForCall(i int) (string, *v1.Event) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventsClient) CreateReturns(result1 *v1.Event, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) CreateReturnsOnCall(i int, result1 *v1.Event, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1.Event
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) List(arg1 v1a.ListOptions) (*v1.EventList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 v1a.ListOptions
	}{arg1})
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventsClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeEventsClient) ListCalls(stub func(v1a.ListOptions) (*v1.EventList, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeEventsClient) ListArgsForCall(i int) v1a.ListOptions {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEventsClient) ListReturns(result1 *v1.EventList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *v1.EventList
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) ListReturnsOnCall(i int, result1 *v1.EventList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *v1.EventList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *v1.EventList
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) Update(arg1 string, arg2 *v1.Event) (*v1.Event, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 string
		arg2 *v1.Event
	}{arg1, arg2})
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventsClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeEventsClient) UpdateCalls(stub func(string, *v1.Event) (*v1.Event, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeEventsClient) UpdateArgsForCall(i int) (string, *v1.Event) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventsClient) UpdateReturns(result1 *v1.Event, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) UpdateReturnsOnCall(i int, result1 *v1.Event, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *v1.Event
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *v1.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventsClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventsClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ reconciler.EventsClient = new(FakeEventsClient)
