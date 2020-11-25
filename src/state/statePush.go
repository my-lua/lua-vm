package state

import "../luavalue"

// PushNil s
func (me *LuaState) PushNil() {
	me.Stack().Push(luavalue.NewLuaNil())
}

// PushBoolean s
func (me *LuaState) PushBoolean(value bool) {
	me.Stack().Push(luavalue.NewLuaBoolean(value))
}

// PushInteger s
func (me *LuaState) PushInteger(value int64) {
	me.Stack().Push(luavalue.NewLuaInteger(value))
}

// PushNumber s
func (me *LuaState) PushNumber(value float64) {
	me.Stack().Push(luavalue.NewLuaNumber(value))
}

// PushString s
func (me *LuaState) PushString(value string) {
	me.Stack().Push(luavalue.NewLuaString(value))
}
