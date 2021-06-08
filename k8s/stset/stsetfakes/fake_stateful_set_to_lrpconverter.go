// Code generated by counterfeiter. DO NOT EDIT.
package stsetfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini-controller/api"
	"code.cloudfoundry.org/eirini-controller/k8s/stset"
	v1 "k8s.io/api/apps/v1"
)

type FakeStatefulSetToLRPConverter struct {
	ConvertStub        func(v1.StatefulSet) (*api.LRP, error)
	convertMutex       sync.RWMutex
	convertArgsForCall []struct {
		arg1 v1.StatefulSet
	}
	convertReturns struct {
		result1 *api.LRP
		result2 error
	}
	convertReturnsOnCall map[int]struct {
		result1 *api.LRP
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatefulSetToLRPConverter) Convert(arg1 v1.StatefulSet) (*api.LRP, error) {
	fake.convertMutex.Lock()
	ret, specificReturn := fake.convertReturnsOnCall[len(fake.convertArgsForCall)]
	fake.convertArgsForCall = append(fake.convertArgsForCall, struct {
		arg1 v1.StatefulSet
	}{arg1})
	stub := fake.ConvertStub
	fakeReturns := fake.convertReturns
	fake.recordInvocation("Convert", []interface{}{arg1})
	fake.convertMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStatefulSetToLRPConverter) ConvertCallCount() int {
	fake.convertMutex.RLock()
	defer fake.convertMutex.RUnlock()
	return len(fake.convertArgsForCall)
}

func (fake *FakeStatefulSetToLRPConverter) ConvertCalls(stub func(v1.StatefulSet) (*api.LRP, error)) {
	fake.convertMutex.Lock()
	defer fake.convertMutex.Unlock()
	fake.ConvertStub = stub
}

func (fake *FakeStatefulSetToLRPConverter) ConvertArgsForCall(i int) v1.StatefulSet {
	fake.convertMutex.RLock()
	defer fake.convertMutex.RUnlock()
	argsForCall := fake.convertArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStatefulSetToLRPConverter) ConvertReturns(result1 *api.LRP, result2 error) {
	fake.convertMutex.Lock()
	defer fake.convertMutex.Unlock()
	fake.ConvertStub = nil
	fake.convertReturns = struct {
		result1 *api.LRP
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetToLRPConverter) ConvertReturnsOnCall(i int, result1 *api.LRP, result2 error) {
	fake.convertMutex.Lock()
	defer fake.convertMutex.Unlock()
	fake.ConvertStub = nil
	if fake.convertReturnsOnCall == nil {
		fake.convertReturnsOnCall = make(map[int]struct {
			result1 *api.LRP
			result2 error
		})
	}
	fake.convertReturnsOnCall[i] = struct {
		result1 *api.LRP
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetToLRPConverter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convertMutex.RLock()
	defer fake.convertMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatefulSetToLRPConverter) recordInvocation(key string, args []interface{}) {
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

var _ stset.StatefulSetToLRPConverter = new(FakeStatefulSetToLRPConverter)
