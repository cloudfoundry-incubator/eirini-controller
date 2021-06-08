// Code generated by counterfeiter. DO NOT EDIT.
package eventfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/eirini-controller/k8s/informers/event"
	"code.cloudfoundry.org/eirini-controller/k8s/reconciler"
	"code.cloudfoundry.org/lager"
	v1 "k8s.io/api/core/v1"
)

type FakeCrashEventGenerator struct {
	GenerateStub        func(context.Context, *v1.Pod, lager.Logger) (reconciler.CrashEvent, bool)
	generateMutex       sync.RWMutex
	generateArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.Pod
		arg3 lager.Logger
	}
	generateReturns struct {
		result1 reconciler.CrashEvent
		result2 bool
	}
	generateReturnsOnCall map[int]struct {
		result1 reconciler.CrashEvent
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCrashEventGenerator) Generate(arg1 context.Context, arg2 *v1.Pod, arg3 lager.Logger) (reconciler.CrashEvent, bool) {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.Pod
		arg3 lager.Logger
	}{arg1, arg2, arg3})
	stub := fake.GenerateStub
	fakeReturns := fake.generateReturns
	fake.recordInvocation("Generate", []interface{}{arg1, arg2, arg3})
	fake.generateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCrashEventGenerator) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *FakeCrashEventGenerator) GenerateCalls(stub func(context.Context, *v1.Pod, lager.Logger) (reconciler.CrashEvent, bool)) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = stub
}

func (fake *FakeCrashEventGenerator) GenerateArgsForCall(i int) (context.Context, *v1.Pod, lager.Logger) {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	argsForCall := fake.generateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCrashEventGenerator) GenerateReturns(result1 reconciler.CrashEvent, result2 bool) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 reconciler.CrashEvent
		result2 bool
	}{result1, result2}
}

func (fake *FakeCrashEventGenerator) GenerateReturnsOnCall(i int, result1 reconciler.CrashEvent, result2 bool) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 reconciler.CrashEvent
			result2 bool
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 reconciler.CrashEvent
		result2 bool
	}{result1, result2}
}

func (fake *FakeCrashEventGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCrashEventGenerator) recordInvocation(key string, args []interface{}) {
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

var _ event.CrashEventGenerator = new(FakeCrashEventGenerator)
