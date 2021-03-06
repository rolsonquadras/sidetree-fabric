// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"
)

type RestConfig struct {
	SidetreeListenURLStub        func() (string, error)
	sidetreeListenURLMutex       sync.RWMutex
	sidetreeListenURLArgsForCall []struct{}
	sidetreeListenURLReturns     struct {
		result1 string
		result2 error
	}
	sidetreeListenURLReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RestConfig) SidetreeListenURL() (string, error) {
	fake.sidetreeListenURLMutex.Lock()
	ret, specificReturn := fake.sidetreeListenURLReturnsOnCall[len(fake.sidetreeListenURLArgsForCall)]
	fake.sidetreeListenURLArgsForCall = append(fake.sidetreeListenURLArgsForCall, struct{}{})
	fake.recordInvocation("SidetreeListenURL", []interface{}{})
	fake.sidetreeListenURLMutex.Unlock()
	if fake.SidetreeListenURLStub != nil {
		return fake.SidetreeListenURLStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.sidetreeListenURLReturns.result1, fake.sidetreeListenURLReturns.result2
}

func (fake *RestConfig) SidetreeListenURLCallCount() int {
	fake.sidetreeListenURLMutex.RLock()
	defer fake.sidetreeListenURLMutex.RUnlock()
	return len(fake.sidetreeListenURLArgsForCall)
}

func (fake *RestConfig) SidetreeListenURLReturns(result1 string, result2 error) {
	fake.SidetreeListenURLStub = nil
	fake.sidetreeListenURLReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *RestConfig) SidetreeListenURLReturnsOnCall(i int, result1 string, result2 error) {
	fake.SidetreeListenURLStub = nil
	if fake.sidetreeListenURLReturnsOnCall == nil {
		fake.sidetreeListenURLReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.sidetreeListenURLReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *RestConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sidetreeListenURLMutex.RLock()
	defer fake.sidetreeListenURLMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RestConfig) recordInvocation(key string, args []interface{}) {
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
