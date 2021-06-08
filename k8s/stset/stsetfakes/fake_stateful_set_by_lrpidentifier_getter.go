// Code generated by counterfeiter. DO NOT EDIT.
package stsetfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/eirini-controller/api"
	"code.cloudfoundry.org/eirini-controller/k8s/stset"
	v1 "k8s.io/api/apps/v1"
)

type FakeStatefulSetByLRPIdentifierGetter struct {
	GetByLRPIdentifierStub        func(context.Context, api.LRPIdentifier) ([]v1.StatefulSet, error)
	getByLRPIdentifierMutex       sync.RWMutex
	getByLRPIdentifierArgsForCall []struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
	}
	getByLRPIdentifierReturns struct {
		result1 []v1.StatefulSet
		result2 error
	}
	getByLRPIdentifierReturnsOnCall map[int]struct {
		result1 []v1.StatefulSet
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatefulSetByLRPIdentifierGetter) GetByLRPIdentifier(arg1 context.Context, arg2 api.LRPIdentifier) ([]v1.StatefulSet, error) {
	fake.getByLRPIdentifierMutex.Lock()
	ret, specificReturn := fake.getByLRPIdentifierReturnsOnCall[len(fake.getByLRPIdentifierArgsForCall)]
	fake.getByLRPIdentifierArgsForCall = append(fake.getByLRPIdentifierArgsForCall, struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
	}{arg1, arg2})
	stub := fake.GetByLRPIdentifierStub
	fakeReturns := fake.getByLRPIdentifierReturns
	fake.recordInvocation("GetByLRPIdentifier", []interface{}{arg1, arg2})
	fake.getByLRPIdentifierMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStatefulSetByLRPIdentifierGetter) GetByLRPIdentifierCallCount() int {
	fake.getByLRPIdentifierMutex.RLock()
	defer fake.getByLRPIdentifierMutex.RUnlock()
	return len(fake.getByLRPIdentifierArgsForCall)
}

func (fake *FakeStatefulSetByLRPIdentifierGetter) GetByLRPIdentifierCalls(stub func(context.Context, api.LRPIdentifier) ([]v1.StatefulSet, error)) {
	fake.getByLRPIdentifierMutex.Lock()
	defer fake.getByLRPIdentifierMutex.Unlock()
	fake.GetByLRPIdentifierStub = stub
}

func (fake *FakeStatefulSetByLRPIdentifierGetter) GetByLRPIdentifierArgsForCall(i int) (context.Context, api.LRPIdentifier) {
	fake.getByLRPIdentifierMutex.RLock()
	defer fake.getByLRPIdentifierMutex.RUnlock()
	argsForCall := fake.getByLRPIdentifierArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStatefulSetByLRPIdentifierGetter) GetByLRPIdentifierReturns(result1 []v1.StatefulSet, result2 error) {
	fake.getByLRPIdentifierMutex.Lock()
	defer fake.getByLRPIdentifierMutex.Unlock()
	fake.GetByLRPIdentifierStub = nil
	fake.getByLRPIdentifierReturns = struct {
		result1 []v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetByLRPIdentifierGetter) GetByLRPIdentifierReturnsOnCall(i int, result1 []v1.StatefulSet, result2 error) {
	fake.getByLRPIdentifierMutex.Lock()
	defer fake.getByLRPIdentifierMutex.Unlock()
	fake.GetByLRPIdentifierStub = nil
	if fake.getByLRPIdentifierReturnsOnCall == nil {
		fake.getByLRPIdentifierReturnsOnCall = make(map[int]struct {
			result1 []v1.StatefulSet
			result2 error
		})
	}
	fake.getByLRPIdentifierReturnsOnCall[i] = struct {
		result1 []v1.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetByLRPIdentifierGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getByLRPIdentifierMutex.RLock()
	defer fake.getByLRPIdentifierMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatefulSetByLRPIdentifierGetter) recordInvocation(key string, args []interface{}) {
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

var _ stset.StatefulSetByLRPIdentifierGetter = new(FakeStatefulSetByLRPIdentifierGetter)
