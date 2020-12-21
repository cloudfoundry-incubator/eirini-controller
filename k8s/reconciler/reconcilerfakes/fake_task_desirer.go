// Code generated by counterfeiter. DO NOT EDIT.
package reconcilerfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/k8s/reconciler"
	"code.cloudfoundry.org/eirini/k8s/shared"
	"code.cloudfoundry.org/eirini/opi"
)

type FakeTaskDesirer struct {
	DesireStub        func(string, *opi.Task, ...shared.Option) error
	desireMutex       sync.RWMutex
	desireArgsForCall []struct {
		arg1 string
		arg2 *opi.Task
		arg3 []shared.Option
	}
	desireReturns struct {
		result1 error
	}
	desireReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskDesirer) Desire(arg1 string, arg2 *opi.Task, arg3 ...shared.Option) error {
	fake.desireMutex.Lock()
	ret, specificReturn := fake.desireReturnsOnCall[len(fake.desireArgsForCall)]
	fake.desireArgsForCall = append(fake.desireArgsForCall, struct {
		arg1 string
		arg2 *opi.Task
		arg3 []shared.Option
	}{arg1, arg2, arg3})
	stub := fake.DesireStub
	fakeReturns := fake.desireReturns
	fake.recordInvocation("Desire", []interface{}{arg1, arg2, arg3})
	fake.desireMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskDesirer) DesireCallCount() int {
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	return len(fake.desireArgsForCall)
}

func (fake *FakeTaskDesirer) DesireCalls(stub func(string, *opi.Task, ...shared.Option) error) {
	fake.desireMutex.Lock()
	defer fake.desireMutex.Unlock()
	fake.DesireStub = stub
}

func (fake *FakeTaskDesirer) DesireArgsForCall(i int) (string, *opi.Task, []shared.Option) {
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	argsForCall := fake.desireArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTaskDesirer) DesireReturns(result1 error) {
	fake.desireMutex.Lock()
	defer fake.desireMutex.Unlock()
	fake.DesireStub = nil
	fake.desireReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskDesirer) DesireReturnsOnCall(i int, result1 error) {
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

func (fake *FakeTaskDesirer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskDesirer) recordInvocation(key string, args []interface{}) {
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

var _ reconciler.TaskDesirer = new(FakeTaskDesirer)
