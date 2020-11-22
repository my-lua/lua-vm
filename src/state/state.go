package state

// LuaState Lua状态实例
type LuaState struct {
	stack *LuaStack
}

// PushValue s
func (me *LuaState) PushValue(index int) {

}

// ReplaceValue s
func (me *LuaState) ReplaceValue(index int) {

}

// Insert s
func (me *LuaState) Insert(index int) {

}

// Remove s
func (me *LuaState) Remove(index int) {

}

// NewLuaState 构造函数创建一个LuaState
func NewLuaState() *LuaState {
	return &LuaState{
		stack: NewLuaStack(20),
	}
}
