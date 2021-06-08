// Code generated by counterfeiter. DO NOT EDIT.
package reconcilerfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/eirini-controller/api"
	"code.cloudfoundry.org/eirini-controller/k8s/reconciler"
	"code.cloudfoundry.org/eirini-controller/k8s/shared"
	v1 "code.cloudfoundry.org/eirini-controller/pkg/apis/eirini/v1"
)

type FakeTaskWorkloadClient struct {
	DeleteStub        func(context.Context, string) (string, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	deleteReturns struct {
		result1 string
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DesireStub        func(context.Context, string, *api.Task, ...shared.Option) error
	desireMutex       sync.RWMutex
	desireArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 *api.Task
		arg4 []shared.Option
	}
	desireReturns struct {
		result1 error
	}
	desireReturnsOnCall map[int]struct {
		result1 error
	}
	GetStatusStub        func(context.Context, string) (v1.TaskStatus, error)
	getStatusMutex       sync.RWMutex
	getStatusArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getStatusReturns struct {
		result1 v1.TaskStatus
		result2 error
	}
	getStatusReturnsOnCall map[int]struct {
		result1 v1.TaskStatus
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskWorkloadClient) Delete(arg1 context.Context, arg2 string) (string, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskWorkloadClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeTaskWorkloadClient) DeleteCalls(stub func(context.Context, string) (string, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeTaskWorkloadClient) DeleteArgsForCall(i int) (context.Context, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTaskWorkloadClient) DeleteReturns(result1 string, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskWorkloadClient) DeleteReturnsOnCall(i int, result1 string, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskWorkloadClient) Desire(arg1 context.Context, arg2 string, arg3 *api.Task, arg4 ...shared.Option) error {
	fake.desireMutex.Lock()
	ret, specificReturn := fake.desireReturnsOnCall[len(fake.desireArgsForCall)]
	fake.desireArgsForCall = append(fake.desireArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 *api.Task
		arg4 []shared.Option
	}{arg1, arg2, arg3, arg4})
	stub := fake.DesireStub
	fakeReturns := fake.desireReturns
	fake.recordInvocation("Desire", []interface{}{arg1, arg2, arg3, arg4})
	fake.desireMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskWorkloadClient) DesireCallCount() int {
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	return len(fake.desireArgsForCall)
}

func (fake *FakeTaskWorkloadClient) DesireCalls(stub func(context.Context, string, *api.Task, ...shared.Option) error) {
	fake.desireMutex.Lock()
	defer fake.desireMutex.Unlock()
	fake.DesireStub = stub
}

func (fake *FakeTaskWorkloadClient) DesireArgsForCall(i int) (context.Context, string, *api.Task, []shared.Option) {
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	argsForCall := fake.desireArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTaskWorkloadClient) DesireReturns(result1 error) {
	fake.desireMutex.Lock()
	defer fake.desireMutex.Unlock()
	fake.DesireStub = nil
	fake.desireReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskWorkloadClient) DesireReturnsOnCall(i int, result1 error) {
	fake.desireMutex.Lock()
	defer fake.desireMutex.Unlock()
	fake.DesireStub = nil
	if fake.desireReturnsOnCall == nil {
		fake.desireReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.desireReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskWorkloadClient) GetStatus(arg1 context.Context, arg2 string) (v1.TaskStatus, error) {
	fake.getStatusMutex.Lock()
	ret, specificReturn := fake.getStatusReturnsOnCall[len(fake.getStatusArgsForCall)]
	fake.getStatusArgsForCall = append(fake.getStatusArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStatusStub
	fakeReturns := fake.getStatusReturns
	fake.recordInvocation("GetStatus", []interface{}{arg1, arg2})
	fake.getStatusMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskWorkloadClient) GetStatusCallCount() int {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	return len(fake.getStatusArgsForCall)
}

func (fake *FakeTaskWorkloadClient) GetStatusCalls(stub func(context.Context, string) (v1.TaskStatus, error)) {
	fake.getStatusMutex.Lock()
	defer fake.getStatusMutex.Unlock()
	fake.GetStatusStub = stub
}

func (fake *FakeTaskWorkloadClient) GetStatusArgsForCall(i int) (context.Context, string) {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	argsForCall := fake.getStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTaskWorkloadClient) GetStatusReturns(result1 v1.TaskStatus, result2 error) {
	fake.getStatusMutex.Lock()
	defer fake.getStatusMutex.Unlock()
	fake.GetStatusStub = nil
	fake.getStatusReturns = struct {
		result1 v1.TaskStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskWorkloadClient) GetStatusReturnsOnCall(i int, result1 v1.TaskStatus, result2 error) {
	fake.getStatusMutex.Lock()
	defer fake.getStatusMutex.Unlock()
	fake.GetStatusStub = nil
	if fake.getStatusReturnsOnCall == nil {
		fake.getStatusReturnsOnCall = make(map[int]struct {
			result1 v1.TaskStatus
			result2 error
		})
	}
	fake.getStatusReturnsOnCall[i] = struct {
		result1 v1.TaskStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskWorkloadClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskWorkloadClient) recordInvocation(key string, args []interface{}) {
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

var _ reconciler.TaskWorkloadClient = new(FakeTaskWorkloadClient)
