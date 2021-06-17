// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/jesseduffield/lazygit/pkg/commands"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/commands/types"
)

type FakeIDiffMgr struct {
	DiffEndArgsStub        func(string, string, bool, string) string
	diffEndArgsMutex       sync.RWMutex
	diffEndArgsArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
	}
	diffEndArgsReturns struct {
		result1 string
	}
	diffEndArgsReturnsOnCall map[int]struct {
		result1 string
	}
	LoadDiffFilesStub        func(string, string, bool) ([]*models.CommitFile, error)
	loadDiffFilesMutex       sync.RWMutex
	loadDiffFilesArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	loadDiffFilesReturns struct {
		result1 []*models.CommitFile
		result2 error
	}
	loadDiffFilesReturnsOnCall map[int]struct {
		result1 []*models.CommitFile
		result2 error
	}
	ShowFileDiffStub        func(string, string, bool, string, bool) (string, error)
	showFileDiffMutex       sync.RWMutex
	showFileDiffArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
		arg5 bool
	}
	showFileDiffReturns struct {
		result1 string
		result2 error
	}
	showFileDiffReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ShowFileDiffCmdObjStub        func(string, string, bool, string, bool, bool) types.ICmdObj
	showFileDiffCmdObjMutex       sync.RWMutex
	showFileDiffCmdObjArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
		arg5 bool
		arg6 bool
	}
	showFileDiffCmdObjReturns struct {
		result1 types.ICmdObj
	}
	showFileDiffCmdObjReturnsOnCall map[int]struct {
		result1 types.ICmdObj
	}
	WorktreeFileDiffStub        func(*models.File, bool, bool) string
	worktreeFileDiffMutex       sync.RWMutex
	worktreeFileDiffArgsForCall []struct {
		arg1 *models.File
		arg2 bool
		arg3 bool
	}
	worktreeFileDiffReturns struct {
		result1 string
	}
	worktreeFileDiffReturnsOnCall map[int]struct {
		result1 string
	}
	WorktreeFileDiffCmdObjStub        func(models.IFile, bool, bool) types.ICmdObj
	worktreeFileDiffCmdObjMutex       sync.RWMutex
	worktreeFileDiffCmdObjArgsForCall []struct {
		arg1 models.IFile
		arg2 bool
		arg3 bool
	}
	worktreeFileDiffCmdObjReturns struct {
		result1 types.ICmdObj
	}
	worktreeFileDiffCmdObjReturnsOnCall map[int]struct {
		result1 types.ICmdObj
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIDiffMgr) DiffEndArgs(arg1 string, arg2 string, arg3 bool, arg4 string) string {
	fake.diffEndArgsMutex.Lock()
	ret, specificReturn := fake.diffEndArgsReturnsOnCall[len(fake.diffEndArgsArgsForCall)]
	fake.diffEndArgsArgsForCall = append(fake.diffEndArgsArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.DiffEndArgsStub
	fakeReturns := fake.diffEndArgsReturns
	fake.recordInvocation("DiffEndArgs", []interface{}{arg1, arg2, arg3, arg4})
	fake.diffEndArgsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIDiffMgr) DiffEndArgsCallCount() int {
	fake.diffEndArgsMutex.RLock()
	defer fake.diffEndArgsMutex.RUnlock()
	return len(fake.diffEndArgsArgsForCall)
}

func (fake *FakeIDiffMgr) DiffEndArgsCalls(stub func(string, string, bool, string) string) {
	fake.diffEndArgsMutex.Lock()
	defer fake.diffEndArgsMutex.Unlock()
	fake.DiffEndArgsStub = stub
}

func (fake *FakeIDiffMgr) DiffEndArgsArgsForCall(i int) (string, string, bool, string) {
	fake.diffEndArgsMutex.RLock()
	defer fake.diffEndArgsMutex.RUnlock()
	argsForCall := fake.diffEndArgsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeIDiffMgr) DiffEndArgsReturns(result1 string) {
	fake.diffEndArgsMutex.Lock()
	defer fake.diffEndArgsMutex.Unlock()
	fake.DiffEndArgsStub = nil
	fake.diffEndArgsReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIDiffMgr) DiffEndArgsReturnsOnCall(i int, result1 string) {
	fake.diffEndArgsMutex.Lock()
	defer fake.diffEndArgsMutex.Unlock()
	fake.DiffEndArgsStub = nil
	if fake.diffEndArgsReturnsOnCall == nil {
		fake.diffEndArgsReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.diffEndArgsReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeIDiffMgr) LoadDiffFiles(arg1 string, arg2 string, arg3 bool) ([]*models.CommitFile, error) {
	fake.loadDiffFilesMutex.Lock()
	ret, specificReturn := fake.loadDiffFilesReturnsOnCall[len(fake.loadDiffFilesArgsForCall)]
	fake.loadDiffFilesArgsForCall = append(fake.loadDiffFilesArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.LoadDiffFilesStub
	fakeReturns := fake.loadDiffFilesReturns
	fake.recordInvocation("LoadDiffFiles", []interface{}{arg1, arg2, arg3})
	fake.loadDiffFilesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIDiffMgr) LoadDiffFilesCallCount() int {
	fake.loadDiffFilesMutex.RLock()
	defer fake.loadDiffFilesMutex.RUnlock()
	return len(fake.loadDiffFilesArgsForCall)
}

func (fake *FakeIDiffMgr) LoadDiffFilesCalls(stub func(string, string, bool) ([]*models.CommitFile, error)) {
	fake.loadDiffFilesMutex.Lock()
	defer fake.loadDiffFilesMutex.Unlock()
	fake.LoadDiffFilesStub = stub
}

func (fake *FakeIDiffMgr) LoadDiffFilesArgsForCall(i int) (string, string, bool) {
	fake.loadDiffFilesMutex.RLock()
	defer fake.loadDiffFilesMutex.RUnlock()
	argsForCall := fake.loadDiffFilesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIDiffMgr) LoadDiffFilesReturns(result1 []*models.CommitFile, result2 error) {
	fake.loadDiffFilesMutex.Lock()
	defer fake.loadDiffFilesMutex.Unlock()
	fake.LoadDiffFilesStub = nil
	fake.loadDiffFilesReturns = struct {
		result1 []*models.CommitFile
		result2 error
	}{result1, result2}
}

func (fake *FakeIDiffMgr) LoadDiffFilesReturnsOnCall(i int, result1 []*models.CommitFile, result2 error) {
	fake.loadDiffFilesMutex.Lock()
	defer fake.loadDiffFilesMutex.Unlock()
	fake.LoadDiffFilesStub = nil
	if fake.loadDiffFilesReturnsOnCall == nil {
		fake.loadDiffFilesReturnsOnCall = make(map[int]struct {
			result1 []*models.CommitFile
			result2 error
		})
	}
	fake.loadDiffFilesReturnsOnCall[i] = struct {
		result1 []*models.CommitFile
		result2 error
	}{result1, result2}
}

func (fake *FakeIDiffMgr) ShowFileDiff(arg1 string, arg2 string, arg3 bool, arg4 string, arg5 bool) (string, error) {
	fake.showFileDiffMutex.Lock()
	ret, specificReturn := fake.showFileDiffReturnsOnCall[len(fake.showFileDiffArgsForCall)]
	fake.showFileDiffArgsForCall = append(fake.showFileDiffArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.ShowFileDiffStub
	fakeReturns := fake.showFileDiffReturns
	fake.recordInvocation("ShowFileDiff", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.showFileDiffMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIDiffMgr) ShowFileDiffCallCount() int {
	fake.showFileDiffMutex.RLock()
	defer fake.showFileDiffMutex.RUnlock()
	return len(fake.showFileDiffArgsForCall)
}

func (fake *FakeIDiffMgr) ShowFileDiffCalls(stub func(string, string, bool, string, bool) (string, error)) {
	fake.showFileDiffMutex.Lock()
	defer fake.showFileDiffMutex.Unlock()
	fake.ShowFileDiffStub = stub
}

func (fake *FakeIDiffMgr) ShowFileDiffArgsForCall(i int) (string, string, bool, string, bool) {
	fake.showFileDiffMutex.RLock()
	defer fake.showFileDiffMutex.RUnlock()
	argsForCall := fake.showFileDiffArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeIDiffMgr) ShowFileDiffReturns(result1 string, result2 error) {
	fake.showFileDiffMutex.Lock()
	defer fake.showFileDiffMutex.Unlock()
	fake.ShowFileDiffStub = nil
	fake.showFileDiffReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeIDiffMgr) ShowFileDiffReturnsOnCall(i int, result1 string, result2 error) {
	fake.showFileDiffMutex.Lock()
	defer fake.showFileDiffMutex.Unlock()
	fake.ShowFileDiffStub = nil
	if fake.showFileDiffReturnsOnCall == nil {
		fake.showFileDiffReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.showFileDiffReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeIDiffMgr) ShowFileDiffCmdObj(arg1 string, arg2 string, arg3 bool, arg4 string, arg5 bool, arg6 bool) types.ICmdObj {
	fake.showFileDiffCmdObjMutex.Lock()
	ret, specificReturn := fake.showFileDiffCmdObjReturnsOnCall[len(fake.showFileDiffCmdObjArgsForCall)]
	fake.showFileDiffCmdObjArgsForCall = append(fake.showFileDiffCmdObjArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
		arg5 bool
		arg6 bool
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	stub := fake.ShowFileDiffCmdObjStub
	fakeReturns := fake.showFileDiffCmdObjReturns
	fake.recordInvocation("ShowFileDiffCmdObj", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.showFileDiffCmdObjMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIDiffMgr) ShowFileDiffCmdObjCallCount() int {
	fake.showFileDiffCmdObjMutex.RLock()
	defer fake.showFileDiffCmdObjMutex.RUnlock()
	return len(fake.showFileDiffCmdObjArgsForCall)
}

func (fake *FakeIDiffMgr) ShowFileDiffCmdObjCalls(stub func(string, string, bool, string, bool, bool) types.ICmdObj) {
	fake.showFileDiffCmdObjMutex.Lock()
	defer fake.showFileDiffCmdObjMutex.Unlock()
	fake.ShowFileDiffCmdObjStub = stub
}

func (fake *FakeIDiffMgr) ShowFileDiffCmdObjArgsForCall(i int) (string, string, bool, string, bool, bool) {
	fake.showFileDiffCmdObjMutex.RLock()
	defer fake.showFileDiffCmdObjMutex.RUnlock()
	argsForCall := fake.showFileDiffCmdObjArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeIDiffMgr) ShowFileDiffCmdObjReturns(result1 types.ICmdObj) {
	fake.showFileDiffCmdObjMutex.Lock()
	defer fake.showFileDiffCmdObjMutex.Unlock()
	fake.ShowFileDiffCmdObjStub = nil
	fake.showFileDiffCmdObjReturns = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeIDiffMgr) ShowFileDiffCmdObjReturnsOnCall(i int, result1 types.ICmdObj) {
	fake.showFileDiffCmdObjMutex.Lock()
	defer fake.showFileDiffCmdObjMutex.Unlock()
	fake.ShowFileDiffCmdObjStub = nil
	if fake.showFileDiffCmdObjReturnsOnCall == nil {
		fake.showFileDiffCmdObjReturnsOnCall = make(map[int]struct {
			result1 types.ICmdObj
		})
	}
	fake.showFileDiffCmdObjReturnsOnCall[i] = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeIDiffMgr) WorktreeFileDiff(arg1 *models.File, arg2 bool, arg3 bool) string {
	fake.worktreeFileDiffMutex.Lock()
	ret, specificReturn := fake.worktreeFileDiffReturnsOnCall[len(fake.worktreeFileDiffArgsForCall)]
	fake.worktreeFileDiffArgsForCall = append(fake.worktreeFileDiffArgsForCall, struct {
		arg1 *models.File
		arg2 bool
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.WorktreeFileDiffStub
	fakeReturns := fake.worktreeFileDiffReturns
	fake.recordInvocation("WorktreeFileDiff", []interface{}{arg1, arg2, arg3})
	fake.worktreeFileDiffMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIDiffMgr) WorktreeFileDiffCallCount() int {
	fake.worktreeFileDiffMutex.RLock()
	defer fake.worktreeFileDiffMutex.RUnlock()
	return len(fake.worktreeFileDiffArgsForCall)
}

func (fake *FakeIDiffMgr) WorktreeFileDiffCalls(stub func(*models.File, bool, bool) string) {
	fake.worktreeFileDiffMutex.Lock()
	defer fake.worktreeFileDiffMutex.Unlock()
	fake.WorktreeFileDiffStub = stub
}

func (fake *FakeIDiffMgr) WorktreeFileDiffArgsForCall(i int) (*models.File, bool, bool) {
	fake.worktreeFileDiffMutex.RLock()
	defer fake.worktreeFileDiffMutex.RUnlock()
	argsForCall := fake.worktreeFileDiffArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIDiffMgr) WorktreeFileDiffReturns(result1 string) {
	fake.worktreeFileDiffMutex.Lock()
	defer fake.worktreeFileDiffMutex.Unlock()
	fake.WorktreeFileDiffStub = nil
	fake.worktreeFileDiffReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIDiffMgr) WorktreeFileDiffReturnsOnCall(i int, result1 string) {
	fake.worktreeFileDiffMutex.Lock()
	defer fake.worktreeFileDiffMutex.Unlock()
	fake.WorktreeFileDiffStub = nil
	if fake.worktreeFileDiffReturnsOnCall == nil {
		fake.worktreeFileDiffReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.worktreeFileDiffReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeIDiffMgr) WorktreeFileDiffCmdObj(arg1 models.IFile, arg2 bool, arg3 bool) types.ICmdObj {
	fake.worktreeFileDiffCmdObjMutex.Lock()
	ret, specificReturn := fake.worktreeFileDiffCmdObjReturnsOnCall[len(fake.worktreeFileDiffCmdObjArgsForCall)]
	fake.worktreeFileDiffCmdObjArgsForCall = append(fake.worktreeFileDiffCmdObjArgsForCall, struct {
		arg1 models.IFile
		arg2 bool
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.WorktreeFileDiffCmdObjStub
	fakeReturns := fake.worktreeFileDiffCmdObjReturns
	fake.recordInvocation("WorktreeFileDiffCmdObj", []interface{}{arg1, arg2, arg3})
	fake.worktreeFileDiffCmdObjMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIDiffMgr) WorktreeFileDiffCmdObjCallCount() int {
	fake.worktreeFileDiffCmdObjMutex.RLock()
	defer fake.worktreeFileDiffCmdObjMutex.RUnlock()
	return len(fake.worktreeFileDiffCmdObjArgsForCall)
}

func (fake *FakeIDiffMgr) WorktreeFileDiffCmdObjCalls(stub func(models.IFile, bool, bool) types.ICmdObj) {
	fake.worktreeFileDiffCmdObjMutex.Lock()
	defer fake.worktreeFileDiffCmdObjMutex.Unlock()
	fake.WorktreeFileDiffCmdObjStub = stub
}

func (fake *FakeIDiffMgr) WorktreeFileDiffCmdObjArgsForCall(i int) (models.IFile, bool, bool) {
	fake.worktreeFileDiffCmdObjMutex.RLock()
	defer fake.worktreeFileDiffCmdObjMutex.RUnlock()
	argsForCall := fake.worktreeFileDiffCmdObjArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIDiffMgr) WorktreeFileDiffCmdObjReturns(result1 types.ICmdObj) {
	fake.worktreeFileDiffCmdObjMutex.Lock()
	defer fake.worktreeFileDiffCmdObjMutex.Unlock()
	fake.WorktreeFileDiffCmdObjStub = nil
	fake.worktreeFileDiffCmdObjReturns = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeIDiffMgr) WorktreeFileDiffCmdObjReturnsOnCall(i int, result1 types.ICmdObj) {
	fake.worktreeFileDiffCmdObjMutex.Lock()
	defer fake.worktreeFileDiffCmdObjMutex.Unlock()
	fake.WorktreeFileDiffCmdObjStub = nil
	if fake.worktreeFileDiffCmdObjReturnsOnCall == nil {
		fake.worktreeFileDiffCmdObjReturnsOnCall = make(map[int]struct {
			result1 types.ICmdObj
		})
	}
	fake.worktreeFileDiffCmdObjReturnsOnCall[i] = struct {
		result1 types.ICmdObj
	}{result1}
}

func (fake *FakeIDiffMgr) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.diffEndArgsMutex.RLock()
	defer fake.diffEndArgsMutex.RUnlock()
	fake.loadDiffFilesMutex.RLock()
	defer fake.loadDiffFilesMutex.RUnlock()
	fake.showFileDiffMutex.RLock()
	defer fake.showFileDiffMutex.RUnlock()
	fake.showFileDiffCmdObjMutex.RLock()
	defer fake.showFileDiffCmdObjMutex.RUnlock()
	fake.worktreeFileDiffMutex.RLock()
	defer fake.worktreeFileDiffMutex.RUnlock()
	fake.worktreeFileDiffCmdObjMutex.RLock()
	defer fake.worktreeFileDiffCmdObjMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIDiffMgr) recordInvocation(key string, args []interface{}) {
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

var _ commands.IDiffMgr = new(FakeIDiffMgr)
