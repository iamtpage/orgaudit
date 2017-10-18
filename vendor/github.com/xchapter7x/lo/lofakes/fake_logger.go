// This file was generated by counterfeiter
package lofakes

import (
	"sync"

	"github.com/xchapter7x/lo"
)

type FakeLogger struct {
	CriticalStub        func(args ...interface{})
	criticalMutex       sync.RWMutex
	criticalArgsForCall []struct {
		args []interface{}
	}
	CriticalfStub        func(format string, args ...interface{})
	criticalfMutex       sync.RWMutex
	criticalfArgsForCall []struct {
		format string
		args   []interface{}
	}
	DebugStub        func(args ...interface{})
	debugMutex       sync.RWMutex
	debugArgsForCall []struct {
		args []interface{}
	}
	DebugfStub        func(format string, args ...interface{})
	debugfMutex       sync.RWMutex
	debugfArgsForCall []struct {
		format string
		args   []interface{}
	}
	ErrorStub        func(args ...interface{})
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
		args []interface{}
	}
	ErrorfStub        func(format string, args ...interface{})
	errorfMutex       sync.RWMutex
	errorfArgsForCall []struct {
		format string
		args   []interface{}
	}
	FatalStub        func(args ...interface{})
	fatalMutex       sync.RWMutex
	fatalArgsForCall []struct {
		args []interface{}
	}
	FatalfStub        func(format string, args ...interface{})
	fatalfMutex       sync.RWMutex
	fatalfArgsForCall []struct {
		format string
		args   []interface{}
	}
	InfoStub        func(args ...interface{})
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		args []interface{}
	}
	InfofStub        func(format string, args ...interface{})
	infofMutex       sync.RWMutex
	infofArgsForCall []struct {
		format string
		args   []interface{}
	}
	NoticeStub        func(args ...interface{})
	noticeMutex       sync.RWMutex
	noticeArgsForCall []struct {
		args []interface{}
	}
	NoticefStub        func(format string, args ...interface{})
	noticefMutex       sync.RWMutex
	noticefArgsForCall []struct {
		format string
		args   []interface{}
	}
	PanicStub        func(args ...interface{})
	panicMutex       sync.RWMutex
	panicArgsForCall []struct {
		args []interface{}
	}
	PanicfStub        func(format string, args ...interface{})
	panicfMutex       sync.RWMutex
	panicfArgsForCall []struct {
		format string
		args   []interface{}
	}
	WarningStub        func(args ...interface{})
	warningMutex       sync.RWMutex
	warningArgsForCall []struct {
		args []interface{}
	}
	WarningfStub        func(format string, args ...interface{})
	warningfMutex       sync.RWMutex
	warningfArgsForCall []struct {
		format string
		args   []interface{}
	}
}

func (fake *FakeLogger) Critical(args ...interface{}) {
	fake.criticalMutex.Lock()
	fake.criticalArgsForCall = append(fake.criticalArgsForCall, struct {
		args []interface{}
	}{args})
	fake.criticalMutex.Unlock()
	if fake.CriticalStub != nil {
		fake.CriticalStub(args...)
	}
}

func (fake *FakeLogger) CriticalCallCount() int {
	fake.criticalMutex.RLock()
	defer fake.criticalMutex.RUnlock()
	return len(fake.criticalArgsForCall)
}

func (fake *FakeLogger) CriticalArgsForCall(i int) []interface{} {
	fake.criticalMutex.RLock()
	defer fake.criticalMutex.RUnlock()
	return fake.criticalArgsForCall[i].args
}

func (fake *FakeLogger) Criticalf(format string, args ...interface{}) {
	fake.criticalfMutex.Lock()
	fake.criticalfArgsForCall = append(fake.criticalfArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.criticalfMutex.Unlock()
	if fake.CriticalfStub != nil {
		fake.CriticalfStub(format, args...)
	}
}

func (fake *FakeLogger) CriticalfCallCount() int {
	fake.criticalfMutex.RLock()
	defer fake.criticalfMutex.RUnlock()
	return len(fake.criticalfArgsForCall)
}

func (fake *FakeLogger) CriticalfArgsForCall(i int) (string, []interface{}) {
	fake.criticalfMutex.RLock()
	defer fake.criticalfMutex.RUnlock()
	return fake.criticalfArgsForCall[i].format, fake.criticalfArgsForCall[i].args
}

func (fake *FakeLogger) Debug(args ...interface{}) {
	fake.debugMutex.Lock()
	fake.debugArgsForCall = append(fake.debugArgsForCall, struct {
		args []interface{}
	}{args})
	fake.debugMutex.Unlock()
	if fake.DebugStub != nil {
		fake.DebugStub(args...)
	}
}

func (fake *FakeLogger) DebugCallCount() int {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return len(fake.debugArgsForCall)
}

func (fake *FakeLogger) DebugArgsForCall(i int) []interface{} {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return fake.debugArgsForCall[i].args
}

func (fake *FakeLogger) Debugf(format string, args ...interface{}) {
	fake.debugfMutex.Lock()
	fake.debugfArgsForCall = append(fake.debugfArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.debugfMutex.Unlock()
	if fake.DebugfStub != nil {
		fake.DebugfStub(format, args...)
	}
}

func (fake *FakeLogger) DebugfCallCount() int {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return len(fake.debugfArgsForCall)
}

func (fake *FakeLogger) DebugfArgsForCall(i int) (string, []interface{}) {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return fake.debugfArgsForCall[i].format, fake.debugfArgsForCall[i].args
}

func (fake *FakeLogger) Error(args ...interface{}) {
	fake.errorMutex.Lock()
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
		args []interface{}
	}{args})
	fake.errorMutex.Unlock()
	if fake.ErrorStub != nil {
		fake.ErrorStub(args...)
	}
}

func (fake *FakeLogger) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeLogger) ErrorArgsForCall(i int) []interface{} {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return fake.errorArgsForCall[i].args
}

func (fake *FakeLogger) Errorf(format string, args ...interface{}) {
	fake.errorfMutex.Lock()
	fake.errorfArgsForCall = append(fake.errorfArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.errorfMutex.Unlock()
	if fake.ErrorfStub != nil {
		fake.ErrorfStub(format, args...)
	}
}

func (fake *FakeLogger) ErrorfCallCount() int {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	return len(fake.errorfArgsForCall)
}

func (fake *FakeLogger) ErrorfArgsForCall(i int) (string, []interface{}) {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	return fake.errorfArgsForCall[i].format, fake.errorfArgsForCall[i].args
}

func (fake *FakeLogger) Fatal(args ...interface{}) {
	fake.fatalMutex.Lock()
	fake.fatalArgsForCall = append(fake.fatalArgsForCall, struct {
		args []interface{}
	}{args})
	fake.fatalMutex.Unlock()
	if fake.FatalStub != nil {
		fake.FatalStub(args...)
	}
}

func (fake *FakeLogger) FatalCallCount() int {
	fake.fatalMutex.RLock()
	defer fake.fatalMutex.RUnlock()
	return len(fake.fatalArgsForCall)
}

func (fake *FakeLogger) FatalArgsForCall(i int) []interface{} {
	fake.fatalMutex.RLock()
	defer fake.fatalMutex.RUnlock()
	return fake.fatalArgsForCall[i].args
}

func (fake *FakeLogger) Fatalf(format string, args ...interface{}) {
	fake.fatalfMutex.Lock()
	fake.fatalfArgsForCall = append(fake.fatalfArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.fatalfMutex.Unlock()
	if fake.FatalfStub != nil {
		fake.FatalfStub(format, args...)
	}
}

func (fake *FakeLogger) FatalfCallCount() int {
	fake.fatalfMutex.RLock()
	defer fake.fatalfMutex.RUnlock()
	return len(fake.fatalfArgsForCall)
}

func (fake *FakeLogger) FatalfArgsForCall(i int) (string, []interface{}) {
	fake.fatalfMutex.RLock()
	defer fake.fatalfMutex.RUnlock()
	return fake.fatalfArgsForCall[i].format, fake.fatalfArgsForCall[i].args
}

func (fake *FakeLogger) Info(args ...interface{}) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		args []interface{}
	}{args})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		fake.InfoStub(args...)
	}
}

func (fake *FakeLogger) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeLogger) InfoArgsForCall(i int) []interface{} {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return fake.infoArgsForCall[i].args
}

func (fake *FakeLogger) Infof(format string, args ...interface{}) {
	fake.infofMutex.Lock()
	fake.infofArgsForCall = append(fake.infofArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.infofMutex.Unlock()
	if fake.InfofStub != nil {
		fake.InfofStub(format, args...)
	}
}

func (fake *FakeLogger) InfofCallCount() int {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	return len(fake.infofArgsForCall)
}

func (fake *FakeLogger) InfofArgsForCall(i int) (string, []interface{}) {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	return fake.infofArgsForCall[i].format, fake.infofArgsForCall[i].args
}

func (fake *FakeLogger) Notice(args ...interface{}) {
	fake.noticeMutex.Lock()
	fake.noticeArgsForCall = append(fake.noticeArgsForCall, struct {
		args []interface{}
	}{args})
	fake.noticeMutex.Unlock()
	if fake.NoticeStub != nil {
		fake.NoticeStub(args...)
	}
}

func (fake *FakeLogger) NoticeCallCount() int {
	fake.noticeMutex.RLock()
	defer fake.noticeMutex.RUnlock()
	return len(fake.noticeArgsForCall)
}

func (fake *FakeLogger) NoticeArgsForCall(i int) []interface{} {
	fake.noticeMutex.RLock()
	defer fake.noticeMutex.RUnlock()
	return fake.noticeArgsForCall[i].args
}

func (fake *FakeLogger) Noticef(format string, args ...interface{}) {
	fake.noticefMutex.Lock()
	fake.noticefArgsForCall = append(fake.noticefArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.noticefMutex.Unlock()
	if fake.NoticefStub != nil {
		fake.NoticefStub(format, args...)
	}
}

func (fake *FakeLogger) NoticefCallCount() int {
	fake.noticefMutex.RLock()
	defer fake.noticefMutex.RUnlock()
	return len(fake.noticefArgsForCall)
}

func (fake *FakeLogger) NoticefArgsForCall(i int) (string, []interface{}) {
	fake.noticefMutex.RLock()
	defer fake.noticefMutex.RUnlock()
	return fake.noticefArgsForCall[i].format, fake.noticefArgsForCall[i].args
}

func (fake *FakeLogger) Panic(args ...interface{}) {
	fake.panicMutex.Lock()
	fake.panicArgsForCall = append(fake.panicArgsForCall, struct {
		args []interface{}
	}{args})
	fake.panicMutex.Unlock()
	if fake.PanicStub != nil {
		fake.PanicStub(args...)
	}
}

func (fake *FakeLogger) PanicCallCount() int {
	fake.panicMutex.RLock()
	defer fake.panicMutex.RUnlock()
	return len(fake.panicArgsForCall)
}

func (fake *FakeLogger) PanicArgsForCall(i int) []interface{} {
	fake.panicMutex.RLock()
	defer fake.panicMutex.RUnlock()
	return fake.panicArgsForCall[i].args
}

func (fake *FakeLogger) Panicf(format string, args ...interface{}) {
	fake.panicfMutex.Lock()
	fake.panicfArgsForCall = append(fake.panicfArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.panicfMutex.Unlock()
	if fake.PanicfStub != nil {
		fake.PanicfStub(format, args...)
	}
}

func (fake *FakeLogger) PanicfCallCount() int {
	fake.panicfMutex.RLock()
	defer fake.panicfMutex.RUnlock()
	return len(fake.panicfArgsForCall)
}

func (fake *FakeLogger) PanicfArgsForCall(i int) (string, []interface{}) {
	fake.panicfMutex.RLock()
	defer fake.panicfMutex.RUnlock()
	return fake.panicfArgsForCall[i].format, fake.panicfArgsForCall[i].args
}

func (fake *FakeLogger) Warning(args ...interface{}) {
	fake.warningMutex.Lock()
	fake.warningArgsForCall = append(fake.warningArgsForCall, struct {
		args []interface{}
	}{args})
	fake.warningMutex.Unlock()
	if fake.WarningStub != nil {
		fake.WarningStub(args...)
	}
}

func (fake *FakeLogger) WarningCallCount() int {
	fake.warningMutex.RLock()
	defer fake.warningMutex.RUnlock()
	return len(fake.warningArgsForCall)
}

func (fake *FakeLogger) WarningArgsForCall(i int) []interface{} {
	fake.warningMutex.RLock()
	defer fake.warningMutex.RUnlock()
	return fake.warningArgsForCall[i].args
}

func (fake *FakeLogger) Warningf(format string, args ...interface{}) {
	fake.warningfMutex.Lock()
	fake.warningfArgsForCall = append(fake.warningfArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.warningfMutex.Unlock()
	if fake.WarningfStub != nil {
		fake.WarningfStub(format, args...)
	}
}

func (fake *FakeLogger) WarningfCallCount() int {
	fake.warningfMutex.RLock()
	defer fake.warningfMutex.RUnlock()
	return len(fake.warningfArgsForCall)
}

func (fake *FakeLogger) WarningfArgsForCall(i int) (string, []interface{}) {
	fake.warningfMutex.RLock()
	defer fake.warningfMutex.RUnlock()
	return fake.warningfArgsForCall[i].format, fake.warningfArgsForCall[i].args
}

var _ lo.Logger = new(FakeLogger)