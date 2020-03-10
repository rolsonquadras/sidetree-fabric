// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/document"
)

type DocumentResolver struct {
	ResolveDocumentStub        func(idOrDocument string) (document.Document, error)
	resolveDocumentMutex       sync.RWMutex
	resolveDocumentArgsForCall []struct {
		idOrDocument string
	}
	resolveDocumentReturns struct {
		result1 document.Document
		result2 error
	}
	resolveDocumentReturnsOnCall map[int]struct {
		result1 document.Document
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DocumentResolver) ResolveDocument(idOrDocument string) (document.Document, error) {
	fake.resolveDocumentMutex.Lock()
	ret, specificReturn := fake.resolveDocumentReturnsOnCall[len(fake.resolveDocumentArgsForCall)]
	fake.resolveDocumentArgsForCall = append(fake.resolveDocumentArgsForCall, struct {
		idOrDocument string
	}{idOrDocument})
	fake.recordInvocation("ResolveDocument", []interface{}{idOrDocument})
	fake.resolveDocumentMutex.Unlock()
	if fake.ResolveDocumentStub != nil {
		return fake.ResolveDocumentStub(idOrDocument)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.resolveDocumentReturns.result1, fake.resolveDocumentReturns.result2
}

func (fake *DocumentResolver) ResolveDocumentCallCount() int {
	fake.resolveDocumentMutex.RLock()
	defer fake.resolveDocumentMutex.RUnlock()
	return len(fake.resolveDocumentArgsForCall)
}

func (fake *DocumentResolver) ResolveDocumentArgsForCall(i int) string {
	fake.resolveDocumentMutex.RLock()
	defer fake.resolveDocumentMutex.RUnlock()
	return fake.resolveDocumentArgsForCall[i].idOrDocument
}

func (fake *DocumentResolver) ResolveDocumentReturns(result1 document.Document, result2 error) {
	fake.ResolveDocumentStub = nil
	fake.resolveDocumentReturns = struct {
		result1 document.Document
		result2 error
	}{result1, result2}
}

func (fake *DocumentResolver) ResolveDocumentReturnsOnCall(i int, result1 document.Document, result2 error) {
	fake.ResolveDocumentStub = nil
	if fake.resolveDocumentReturnsOnCall == nil {
		fake.resolveDocumentReturnsOnCall = make(map[int]struct {
			result1 document.Document
			result2 error
		})
	}
	fake.resolveDocumentReturnsOnCall[i] = struct {
		result1 document.Document
		result2 error
	}{result1, result2}
}

func (fake *DocumentResolver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.resolveDocumentMutex.RLock()
	defer fake.resolveDocumentMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DocumentResolver) recordInvocation(key string, args []interface{}) {
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
