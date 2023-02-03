package recisub

import (
	"sync"

	"github.com/vela-ssoc/manager/model"
	"github.com/yuin/gopher-lua"
)

type subUser struct {
	rec         *model.Recipient
	evtMutex    sync.Mutex
	evtState    *lua.LState
	evtComplied bool
	evtError    error
	rskMutex    sync.Mutex
	rskState    *lua.LState
	rskComplied bool
	rskError    error
}

func (su *subUser) EvalEvent(evt *model.Event) bool {
	return true
}

func (su *subUser) EvalRisk(rsk *model.Risk) bool {
	return true
}

//
//func (su *subUser) compileEvent() (*lua.LState, error) {
//	su.evtMutex.Lock()
//	defer su.evtMutex.Unlock()
//
//	if su.evtComplied {
//		return su.evtState, su.evtError
//	}
//
//	code := su.rec.EventCode
//	if len(code) == 0 {
//		su.evtComplied = true
//		return nil, nil
//	}
//
//	state := su.newState()
//	err := state.DoString(string(code))
//
//	return nil, nil
//}
//
//func (su *subUser) compile(code []byte) (*lua.LState, error) {
//	state := su.newState()
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer func() {
//		cancel()
//		state.RemoveContext()
//	}()
//
//	state.SetContext(ctx)
//	if err := state.DoString(string(code)); err != nil {
//		return nil, err
//	}
//
//	fn, ok := state.Get(-1).(*lua.LFunction)
//	if !ok {
//		state.NewThread()
//		return nil, nil
//	}
//
//
//	return nil, nil
//}
//
//func (*subUser) newState() *lua.LState {
//	opt := lua.Options{
//		CallStackSize:       32,
//		SkipOpenLibs:        true,
//		IncludeGoStackTrace: true,
//	}
//	return lua.NewState(opt)
//}
