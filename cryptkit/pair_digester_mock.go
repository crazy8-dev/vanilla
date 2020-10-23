package cryptkit

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/vanilla/longbits"
)

// PairDigesterMock implements PairDigester
type PairDigesterMock struct {
	t minimock.Tester

	funcDigestPair          func(digest0 longbits.FoldableReader, digest1 longbits.FoldableReader) (d1 Digest)
	inspectFuncDigestPair   func(digest0 longbits.FoldableReader, digest1 longbits.FoldableReader)
	afterDigestPairCounter  uint64
	beforeDigestPairCounter uint64
	DigestPairMock          mPairDigesterMockDigestPair

	funcGetDigestMethod          func() (d1 DigestMethod)
	inspectFuncGetDigestMethod   func()
	afterGetDigestMethodCounter  uint64
	beforeGetDigestMethodCounter uint64
	GetDigestMethodMock          mPairDigesterMockGetDigestMethod

	funcGetDigestSize          func() (i1 int)
	inspectFuncGetDigestSize   func()
	afterGetDigestSizeCounter  uint64
	beforeGetDigestSizeCounter uint64
	GetDigestSizeMock          mPairDigesterMockGetDigestSize
}

// NewPairDigesterMock returns a mock for PairDigester
func NewPairDigesterMock(t minimock.Tester) *PairDigesterMock {
	m := &PairDigesterMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DigestPairMock = mPairDigesterMockDigestPair{mock: m}
	m.DigestPairMock.callArgs = []*PairDigesterMockDigestPairParams{}

	m.GetDigestMethodMock = mPairDigesterMockGetDigestMethod{mock: m}

	m.GetDigestSizeMock = mPairDigesterMockGetDigestSize{mock: m}

	return m
}

type mPairDigesterMockDigestPair struct {
	mock               *PairDigesterMock
	defaultExpectation *PairDigesterMockDigestPairExpectation
	expectations       []*PairDigesterMockDigestPairExpectation

	callArgs []*PairDigesterMockDigestPairParams
	mutex    sync.RWMutex
}

// PairDigesterMockDigestPairExpectation specifies expectation struct of the PairDigester.DigestPair
type PairDigesterMockDigestPairExpectation struct {
	mock    *PairDigesterMock
	params  *PairDigesterMockDigestPairParams
	results *PairDigesterMockDigestPairResults
	Counter uint64
}

// PairDigesterMockDigestPairParams contains parameters of the PairDigester.DigestPair
type PairDigesterMockDigestPairParams struct {
	digest0 longbits.FoldableReader
	digest1 longbits.FoldableReader
}

// PairDigesterMockDigestPairResults contains results of the PairDigester.DigestPair
type PairDigesterMockDigestPairResults struct {
	d1 Digest
}

// Expect sets up expected params for PairDigester.DigestPair
func (mmDigestPair *mPairDigesterMockDigestPair) Expect(digest0 longbits.FoldableReader, digest1 longbits.FoldableReader) *mPairDigesterMockDigestPair {
	if mmDigestPair.mock.funcDigestPair != nil {
		mmDigestPair.mock.t.Fatalf("PairDigesterMock.DigestPair mock is already set by Set")
	}

	if mmDigestPair.defaultExpectation == nil {
		mmDigestPair.defaultExpectation = &PairDigesterMockDigestPairExpectation{}
	}

	mmDigestPair.defaultExpectation.params = &PairDigesterMockDigestPairParams{digest0, digest1}
	for _, e := range mmDigestPair.expectations {
		if minimock.Equal(e.params, mmDigestPair.defaultExpectation.params) {
			mmDigestPair.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDigestPair.defaultExpectation.params)
		}
	}

	return mmDigestPair
}

// Inspect accepts an inspector function that has same arguments as the PairDigester.DigestPair
func (mmDigestPair *mPairDigesterMockDigestPair) Inspect(f func(digest0 longbits.FoldableReader, digest1 longbits.FoldableReader)) *mPairDigesterMockDigestPair {
	if mmDigestPair.mock.inspectFuncDigestPair != nil {
		mmDigestPair.mock.t.Fatalf("Inspect function is already set for PairDigesterMock.DigestPair")
	}

	mmDigestPair.mock.inspectFuncDigestPair = f

	return mmDigestPair
}

// Return sets up results that will be returned by PairDigester.DigestPair
func (mmDigestPair *mPairDigesterMockDigestPair) Return(d1 Digest) *PairDigesterMock {
	if mmDigestPair.mock.funcDigestPair != nil {
		mmDigestPair.mock.t.Fatalf("PairDigesterMock.DigestPair mock is already set by Set")
	}

	if mmDigestPair.defaultExpectation == nil {
		mmDigestPair.defaultExpectation = &PairDigesterMockDigestPairExpectation{mock: mmDigestPair.mock}
	}
	mmDigestPair.defaultExpectation.results = &PairDigesterMockDigestPairResults{d1}
	return mmDigestPair.mock
}

//Set uses given function f to mock the PairDigester.DigestPair method
func (mmDigestPair *mPairDigesterMockDigestPair) Set(f func(digest0 longbits.FoldableReader, digest1 longbits.FoldableReader) (d1 Digest)) *PairDigesterMock {
	if mmDigestPair.defaultExpectation != nil {
		mmDigestPair.mock.t.Fatalf("Default expectation is already set for the PairDigester.DigestPair method")
	}

	if len(mmDigestPair.expectations) > 0 {
		mmDigestPair.mock.t.Fatalf("Some expectations are already set for the PairDigester.DigestPair method")
	}

	mmDigestPair.mock.funcDigestPair = f
	return mmDigestPair.mock
}

// When sets expectation for the PairDigester.DigestPair which will trigger the result defined by the following
// Then helper
func (mmDigestPair *mPairDigesterMockDigestPair) When(digest0 longbits.FoldableReader, digest1 longbits.FoldableReader) *PairDigesterMockDigestPairExpectation {
	if mmDigestPair.mock.funcDigestPair != nil {
		mmDigestPair.mock.t.Fatalf("PairDigesterMock.DigestPair mock is already set by Set")
	}

	expectation := &PairDigesterMockDigestPairExpectation{
		mock:   mmDigestPair.mock,
		params: &PairDigesterMockDigestPairParams{digest0, digest1},
	}
	mmDigestPair.expectations = append(mmDigestPair.expectations, expectation)
	return expectation
}

// Then sets up PairDigester.DigestPair return parameters for the expectation previously defined by the When method
func (e *PairDigesterMockDigestPairExpectation) Then(d1 Digest) *PairDigesterMock {
	e.results = &PairDigesterMockDigestPairResults{d1}
	return e.mock
}

// DigestPair implements PairDigester
func (mmDigestPair *PairDigesterMock) DigestPair(digest0 longbits.FoldableReader, digest1 longbits.FoldableReader) (d1 Digest) {
	mm_atomic.AddUint64(&mmDigestPair.beforeDigestPairCounter, 1)
	defer mm_atomic.AddUint64(&mmDigestPair.afterDigestPairCounter, 1)

	if mmDigestPair.inspectFuncDigestPair != nil {
		mmDigestPair.inspectFuncDigestPair(digest0, digest1)
	}

	mm_params := &PairDigesterMockDigestPairParams{digest0, digest1}

	// Record call args
	mmDigestPair.DigestPairMock.mutex.Lock()
	mmDigestPair.DigestPairMock.callArgs = append(mmDigestPair.DigestPairMock.callArgs, mm_params)
	mmDigestPair.DigestPairMock.mutex.Unlock()

	for _, e := range mmDigestPair.DigestPairMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.d1
		}
	}

	if mmDigestPair.DigestPairMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDigestPair.DigestPairMock.defaultExpectation.Counter, 1)
		mm_want := mmDigestPair.DigestPairMock.defaultExpectation.params
		mm_got := PairDigesterMockDigestPairParams{digest0, digest1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDigestPair.t.Errorf("PairDigesterMock.DigestPair got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDigestPair.DigestPairMock.defaultExpectation.results
		if mm_results == nil {
			mmDigestPair.t.Fatal("No results are set for the PairDigesterMock.DigestPair")
		}
		return (*mm_results).d1
	}
	if mmDigestPair.funcDigestPair != nil {
		return mmDigestPair.funcDigestPair(digest0, digest1)
	}
	mmDigestPair.t.Fatalf("Unexpected call to PairDigesterMock.DigestPair. %v %v", digest0, digest1)
	return
}

// DigestPairAfterCounter returns a count of finished PairDigesterMock.DigestPair invocations
func (mmDigestPair *PairDigesterMock) DigestPairAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDigestPair.afterDigestPairCounter)
}

// DigestPairBeforeCounter returns a count of PairDigesterMock.DigestPair invocations
func (mmDigestPair *PairDigesterMock) DigestPairBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDigestPair.beforeDigestPairCounter)
}

// Calls returns a list of arguments used in each call to PairDigesterMock.DigestPair.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDigestPair *mPairDigesterMockDigestPair) Calls() []*PairDigesterMockDigestPairParams {
	mmDigestPair.mutex.RLock()

	argCopy := make([]*PairDigesterMockDigestPairParams, len(mmDigestPair.callArgs))
	copy(argCopy, mmDigestPair.callArgs)

	mmDigestPair.mutex.RUnlock()

	return argCopy
}

// MinimockDigestPairDone returns true if the count of the DigestPair invocations corresponds
// the number of defined expectations
func (m *PairDigesterMock) MinimockDigestPairDone() bool {
	for _, e := range m.DigestPairMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DigestPairMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDigestPairCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDigestPair != nil && mm_atomic.LoadUint64(&m.afterDigestPairCounter) < 1 {
		return false
	}
	return true
}

// MinimockDigestPairInspect logs each unmet expectation
func (m *PairDigesterMock) MinimockDigestPairInspect() {
	for _, e := range m.DigestPairMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PairDigesterMock.DigestPair with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DigestPairMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDigestPairCounter) < 1 {
		if m.DigestPairMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PairDigesterMock.DigestPair")
		} else {
			m.t.Errorf("Expected call to PairDigesterMock.DigestPair with params: %#v", *m.DigestPairMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDigestPair != nil && mm_atomic.LoadUint64(&m.afterDigestPairCounter) < 1 {
		m.t.Error("Expected call to PairDigesterMock.DigestPair")
	}
}

type mPairDigesterMockGetDigestMethod struct {
	mock               *PairDigesterMock
	defaultExpectation *PairDigesterMockGetDigestMethodExpectation
	expectations       []*PairDigesterMockGetDigestMethodExpectation
}

// PairDigesterMockGetDigestMethodExpectation specifies expectation struct of the PairDigester.GetDigestMethod
type PairDigesterMockGetDigestMethodExpectation struct {
	mock *PairDigesterMock

	results *PairDigesterMockGetDigestMethodResults
	Counter uint64
}

// PairDigesterMockGetDigestMethodResults contains results of the PairDigester.GetDigestMethod
type PairDigesterMockGetDigestMethodResults struct {
	d1 DigestMethod
}

// Expect sets up expected params for PairDigester.GetDigestMethod
func (mmGetDigestMethod *mPairDigesterMockGetDigestMethod) Expect() *mPairDigesterMockGetDigestMethod {
	if mmGetDigestMethod.mock.funcGetDigestMethod != nil {
		mmGetDigestMethod.mock.t.Fatalf("PairDigesterMock.GetDigestMethod mock is already set by Set")
	}

	if mmGetDigestMethod.defaultExpectation == nil {
		mmGetDigestMethod.defaultExpectation = &PairDigesterMockGetDigestMethodExpectation{}
	}

	return mmGetDigestMethod
}

// Inspect accepts an inspector function that has same arguments as the PairDigester.GetDigestMethod
func (mmGetDigestMethod *mPairDigesterMockGetDigestMethod) Inspect(f func()) *mPairDigesterMockGetDigestMethod {
	if mmGetDigestMethod.mock.inspectFuncGetDigestMethod != nil {
		mmGetDigestMethod.mock.t.Fatalf("Inspect function is already set for PairDigesterMock.GetDigestMethod")
	}

	mmGetDigestMethod.mock.inspectFuncGetDigestMethod = f

	return mmGetDigestMethod
}

// Return sets up results that will be returned by PairDigester.GetDigestMethod
func (mmGetDigestMethod *mPairDigesterMockGetDigestMethod) Return(d1 DigestMethod) *PairDigesterMock {
	if mmGetDigestMethod.mock.funcGetDigestMethod != nil {
		mmGetDigestMethod.mock.t.Fatalf("PairDigesterMock.GetDigestMethod mock is already set by Set")
	}

	if mmGetDigestMethod.defaultExpectation == nil {
		mmGetDigestMethod.defaultExpectation = &PairDigesterMockGetDigestMethodExpectation{mock: mmGetDigestMethod.mock}
	}
	mmGetDigestMethod.defaultExpectation.results = &PairDigesterMockGetDigestMethodResults{d1}
	return mmGetDigestMethod.mock
}

//Set uses given function f to mock the PairDigester.GetDigestMethod method
func (mmGetDigestMethod *mPairDigesterMockGetDigestMethod) Set(f func() (d1 DigestMethod)) *PairDigesterMock {
	if mmGetDigestMethod.defaultExpectation != nil {
		mmGetDigestMethod.mock.t.Fatalf("Default expectation is already set for the PairDigester.GetDigestMethod method")
	}

	if len(mmGetDigestMethod.expectations) > 0 {
		mmGetDigestMethod.mock.t.Fatalf("Some expectations are already set for the PairDigester.GetDigestMethod method")
	}

	mmGetDigestMethod.mock.funcGetDigestMethod = f
	return mmGetDigestMethod.mock
}

// GetDigestMethod implements PairDigester
func (mmGetDigestMethod *PairDigesterMock) GetDigestMethod() (d1 DigestMethod) {
	mm_atomic.AddUint64(&mmGetDigestMethod.beforeGetDigestMethodCounter, 1)
	defer mm_atomic.AddUint64(&mmGetDigestMethod.afterGetDigestMethodCounter, 1)

	if mmGetDigestMethod.inspectFuncGetDigestMethod != nil {
		mmGetDigestMethod.inspectFuncGetDigestMethod()
	}

	if mmGetDigestMethod.GetDigestMethodMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetDigestMethod.GetDigestMethodMock.defaultExpectation.Counter, 1)

		mm_results := mmGetDigestMethod.GetDigestMethodMock.defaultExpectation.results
		if mm_results == nil {
			mmGetDigestMethod.t.Fatal("No results are set for the PairDigesterMock.GetDigestMethod")
		}
		return (*mm_results).d1
	}
	if mmGetDigestMethod.funcGetDigestMethod != nil {
		return mmGetDigestMethod.funcGetDigestMethod()
	}
	mmGetDigestMethod.t.Fatalf("Unexpected call to PairDigesterMock.GetDigestMethod.")
	return
}

// GetDigestMethodAfterCounter returns a count of finished PairDigesterMock.GetDigestMethod invocations
func (mmGetDigestMethod *PairDigesterMock) GetDigestMethodAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetDigestMethod.afterGetDigestMethodCounter)
}

// GetDigestMethodBeforeCounter returns a count of PairDigesterMock.GetDigestMethod invocations
func (mmGetDigestMethod *PairDigesterMock) GetDigestMethodBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetDigestMethod.beforeGetDigestMethodCounter)
}

// MinimockGetDigestMethodDone returns true if the count of the GetDigestMethod invocations corresponds
// the number of defined expectations
func (m *PairDigesterMock) MinimockGetDigestMethodDone() bool {
	for _, e := range m.GetDigestMethodMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetDigestMethodMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetDigestMethodCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetDigestMethod != nil && mm_atomic.LoadUint64(&m.afterGetDigestMethodCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetDigestMethodInspect logs each unmet expectation
func (m *PairDigesterMock) MinimockGetDigestMethodInspect() {
	for _, e := range m.GetDigestMethodMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to PairDigesterMock.GetDigestMethod")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetDigestMethodMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetDigestMethodCounter) < 1 {
		m.t.Error("Expected call to PairDigesterMock.GetDigestMethod")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetDigestMethod != nil && mm_atomic.LoadUint64(&m.afterGetDigestMethodCounter) < 1 {
		m.t.Error("Expected call to PairDigesterMock.GetDigestMethod")
	}
}

type mPairDigesterMockGetDigestSize struct {
	mock               *PairDigesterMock
	defaultExpectation *PairDigesterMockGetDigestSizeExpectation
	expectations       []*PairDigesterMockGetDigestSizeExpectation
}

// PairDigesterMockGetDigestSizeExpectation specifies expectation struct of the PairDigester.GetDigestSize
type PairDigesterMockGetDigestSizeExpectation struct {
	mock *PairDigesterMock

	results *PairDigesterMockGetDigestSizeResults
	Counter uint64
}

// PairDigesterMockGetDigestSizeResults contains results of the PairDigester.GetDigestSize
type PairDigesterMockGetDigestSizeResults struct {
	i1 int
}

// Expect sets up expected params for PairDigester.GetDigestSize
func (mmGetDigestSize *mPairDigesterMockGetDigestSize) Expect() *mPairDigesterMockGetDigestSize {
	if mmGetDigestSize.mock.funcGetDigestSize != nil {
		mmGetDigestSize.mock.t.Fatalf("PairDigesterMock.GetDigestSize mock is already set by Set")
	}

	if mmGetDigestSize.defaultExpectation == nil {
		mmGetDigestSize.defaultExpectation = &PairDigesterMockGetDigestSizeExpectation{}
	}

	return mmGetDigestSize
}

// Inspect accepts an inspector function that has same arguments as the PairDigester.GetDigestSize
func (mmGetDigestSize *mPairDigesterMockGetDigestSize) Inspect(f func()) *mPairDigesterMockGetDigestSize {
	if mmGetDigestSize.mock.inspectFuncGetDigestSize != nil {
		mmGetDigestSize.mock.t.Fatalf("Inspect function is already set for PairDigesterMock.GetDigestSize")
	}

	mmGetDigestSize.mock.inspectFuncGetDigestSize = f

	return mmGetDigestSize
}

// Return sets up results that will be returned by PairDigester.GetDigestSize
func (mmGetDigestSize *mPairDigesterMockGetDigestSize) Return(i1 int) *PairDigesterMock {
	if mmGetDigestSize.mock.funcGetDigestSize != nil {
		mmGetDigestSize.mock.t.Fatalf("PairDigesterMock.GetDigestSize mock is already set by Set")
	}

	if mmGetDigestSize.defaultExpectation == nil {
		mmGetDigestSize.defaultExpectation = &PairDigesterMockGetDigestSizeExpectation{mock: mmGetDigestSize.mock}
	}
	mmGetDigestSize.defaultExpectation.results = &PairDigesterMockGetDigestSizeResults{i1}
	return mmGetDigestSize.mock
}

//Set uses given function f to mock the PairDigester.GetDigestSize method
func (mmGetDigestSize *mPairDigesterMockGetDigestSize) Set(f func() (i1 int)) *PairDigesterMock {
	if mmGetDigestSize.defaultExpectation != nil {
		mmGetDigestSize.mock.t.Fatalf("Default expectation is already set for the PairDigester.GetDigestSize method")
	}

	if len(mmGetDigestSize.expectations) > 0 {
		mmGetDigestSize.mock.t.Fatalf("Some expectations are already set for the PairDigester.GetDigestSize method")
	}

	mmGetDigestSize.mock.funcGetDigestSize = f
	return mmGetDigestSize.mock
}

// GetDigestSize implements PairDigester
func (mmGetDigestSize *PairDigesterMock) GetDigestSize() (i1 int) {
	mm_atomic.AddUint64(&mmGetDigestSize.beforeGetDigestSizeCounter, 1)
	defer mm_atomic.AddUint64(&mmGetDigestSize.afterGetDigestSizeCounter, 1)

	if mmGetDigestSize.inspectFuncGetDigestSize != nil {
		mmGetDigestSize.inspectFuncGetDigestSize()
	}

	if mmGetDigestSize.GetDigestSizeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetDigestSize.GetDigestSizeMock.defaultExpectation.Counter, 1)

		mm_results := mmGetDigestSize.GetDigestSizeMock.defaultExpectation.results
		if mm_results == nil {
			mmGetDigestSize.t.Fatal("No results are set for the PairDigesterMock.GetDigestSize")
		}
		return (*mm_results).i1
	}
	if mmGetDigestSize.funcGetDigestSize != nil {
		return mmGetDigestSize.funcGetDigestSize()
	}
	mmGetDigestSize.t.Fatalf("Unexpected call to PairDigesterMock.GetDigestSize.")
	return
}

// GetDigestSizeAfterCounter returns a count of finished PairDigesterMock.GetDigestSize invocations
func (mmGetDigestSize *PairDigesterMock) GetDigestSizeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetDigestSize.afterGetDigestSizeCounter)
}

// GetDigestSizeBeforeCounter returns a count of PairDigesterMock.GetDigestSize invocations
func (mmGetDigestSize *PairDigesterMock) GetDigestSizeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetDigestSize.beforeGetDigestSizeCounter)
}

// MinimockGetDigestSizeDone returns true if the count of the GetDigestSize invocations corresponds
// the number of defined expectations
func (m *PairDigesterMock) MinimockGetDigestSizeDone() bool {
	for _, e := range m.GetDigestSizeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetDigestSizeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetDigestSizeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetDigestSize != nil && mm_atomic.LoadUint64(&m.afterGetDigestSizeCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetDigestSizeInspect logs each unmet expectation
func (m *PairDigesterMock) MinimockGetDigestSizeInspect() {
	for _, e := range m.GetDigestSizeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to PairDigesterMock.GetDigestSize")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetDigestSizeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetDigestSizeCounter) < 1 {
		m.t.Error("Expected call to PairDigesterMock.GetDigestSize")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetDigestSize != nil && mm_atomic.LoadUint64(&m.afterGetDigestSizeCounter) < 1 {
		m.t.Error("Expected call to PairDigesterMock.GetDigestSize")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PairDigesterMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDigestPairInspect()

		m.MinimockGetDigestMethodInspect()

		m.MinimockGetDigestSizeInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PairDigesterMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *PairDigesterMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDigestPairDone() &&
		m.MinimockGetDigestMethodDone() &&
		m.MinimockGetDigestSizeDone()
}
