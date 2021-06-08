// Code generated by counterfeiter. DO NOT EDIT.
package stsetfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/eirini-controller/k8s/stset"
	v1 "k8s.io/api/core/v1"
	v1a "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FakeSecretsClient struct {
	CreateStub        func(context.Context, string, *v1.Secret) (*v1.Secret, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 *v1.Secret
	}
	createReturns struct {
		result1 *v1.Secret
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1.Secret
		result2 error
	}
	DeleteStub        func(context.Context, string, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	SetOwnerStub        func(context.Context, *v1.Secret, v1a.Object) (*v1.Secret, error)
	setOwnerMutex       sync.RWMutex
	setOwnerArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.Secret
		arg3 v1a.Object
	}
	setOwnerReturns struct {
		result1 *v1.Secret
		result2 error
	}
	setOwnerReturnsOnCall map[int]struct {
		result1 *v1.Secret
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecretsClient) Create(arg1 context.Context, arg2 string, arg3 *v1.Secret) (*v1.Secret, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 *v1.Secret
	}{arg1, arg2, arg3})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretsClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSecretsClient) CreateCalls(stub func(context.Context, string, *v1.Secret) (*v1.Secret, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeSecretsClient) CreateArgsForCall(i int) (context.Context, string, *v1.Secret) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecretsClient) CreateReturns(result1 *v1.Secret, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretsClient) CreateReturnsOnCall(i int, result1 *v1.Secret, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1.Secret
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretsClient) Delete(arg1 context.Context, arg2 string, arg3 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSecretsClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSecretsClient) DeleteCalls(stub func(context.Context, string, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeSecretsClient) DeleteArgsForCall(i int) (context.Context, string, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecretsClient) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecretsClient) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecretsClient) SetOwner(arg1 context.Context, arg2 *v1.Secret, arg3 v1a.Object) (*v1.Secret, error) {
	fake.setOwnerMutex.Lock()
	ret, specificReturn := fake.setOwnerReturnsOnCall[len(fake.setOwnerArgsForCall)]
	fake.setOwnerArgsForCall = append(fake.setOwnerArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.Secret
		arg3 v1a.Object
	}{arg1, arg2, arg3})
	stub := fake.SetOwnerStub
	fakeReturns := fake.setOwnerReturns
	fake.recordInvocation("SetOwner", []interface{}{arg1, arg2, arg3})
	fake.setOwnerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretsClient) SetOwnerCallCount() int {
	fake.setOwnerMutex.RLock()
	defer fake.setOwnerMutex.RUnlock()
	return len(fake.setOwnerArgsForCall)
}

func (fake *FakeSecretsClient) SetOwnerCalls(stub func(context.Context, *v1.Secret, v1a.Object) (*v1.Secret, error)) {
	fake.setOwnerMutex.Lock()
	defer fake.setOwnerMutex.Unlock()
	fake.SetOwnerStub = stub
}

func (fake *FakeSecretsClient) SetOwnerArgsForCall(i int) (context.Context, *v1.Secret, v1a.Object) {
	fake.setOwnerMutex.RLock()
	defer fake.setOwnerMutex.RUnlock()
	argsForCall := fake.setOwnerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecretsClient) SetOwnerReturns(result1 *v1.Secret, result2 error) {
	fake.setOwnerMutex.Lock()
	defer fake.setOwnerMutex.Unlock()
	fake.SetOwnerStub = nil
	fake.setOwnerReturns = struct {
		result1 *v1.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretsClient) SetOwnerReturnsOnCall(i int, result1 *v1.Secret, result2 error) {
	fake.setOwnerMutex.Lock()
	defer fake.setOwnerMutex.Unlock()
	fake.SetOwnerStub = nil
	if fake.setOwnerReturnsOnCall == nil {
		fake.setOwnerReturnsOnCall = make(map[int]struct {
			result1 *v1.Secret
			result2 error
		})
	}
	fake.setOwnerReturnsOnCall[i] = struct {
		result1 *v1.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretsClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.setOwnerMutex.RLock()
	defer fake.setOwnerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecretsClient) recordInvocation(key string, args []interface{}) {
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

var _ stset.SecretsClient = new(FakeSecretsClient)
