// Code generated by counterfeiter. DO NOT EDIT.
package sharedfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini-controller/k8s/shared"
)

type FakeOption struct {
	Stub        func(interface{}) error
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 interface{}
	}
	returns struct {
		result1 error
	}
	returnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOption) Spy(arg1 interface{}) error {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 interface{}
	}{arg1})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("Option", []interface{}{arg1})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return returns.result1
}

func (fake *FakeOption) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeOption) Calls(stub func(interface{}) error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *FakeOption) ArgsForCall(i int) interface{} {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1
}

func (fake *FakeOption) Returns(result1 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOption) ReturnsOnCall(i int, result1 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOption) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOption) recordInvocation(key string, args []interface{}) {
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

var _ shared.Option = new(FakeOption).Spy
