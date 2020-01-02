// Code generated by counterfeiter. DO NOT EDIT.
package k8sfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/k8s"
	"code.cloudfoundry.org/eirini/metrics"
)

type FakeEmitter struct {
	EmitStub        func(metrics.Message)
	emitMutex       sync.RWMutex
	emitArgsForCall []struct {
		arg1 metrics.Message
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmitter) Emit(arg1 metrics.Message) {
	fake.emitMutex.Lock()
	fake.emitArgsForCall = append(fake.emitArgsForCall, struct {
		arg1 metrics.Message
	}{arg1})
	fake.recordInvocation("Emit", []interface{}{arg1})
	fake.emitMutex.Unlock()
	if fake.EmitStub != nil {
		fake.EmitStub(arg1)
	}
}

func (fake *FakeEmitter) EmitCallCount() int {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	return len(fake.emitArgsForCall)
}

func (fake *FakeEmitter) EmitCalls(stub func(metrics.Message)) {
	fake.emitMutex.Lock()
	defer fake.emitMutex.Unlock()
	fake.EmitStub = stub
}

func (fake *FakeEmitter) EmitArgsForCall(i int) metrics.Message {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	argsForCall := fake.emitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEmitter) recordInvocation(key string, args []interface{}) {
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

var _ k8s.Emitter = new(FakeEmitter)