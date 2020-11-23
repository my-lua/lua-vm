package state

import (
	"../luavalue"
)

// Load s
func (me *LuaState) Load(chunk []byte, chunkName, mode string) int {
	closure := luavalue.NewLuaClosure()
	me.stack.Push(closure)
	return 0
}

// Call s
func (me *LuaState) Call(nArgs, nResults int) {
	closure := me.stack.Get(-nArgs - 1)
	if closure.Type() == luavalue.LuaTypeFunction {
		me.callLuaClosure(closure.(*luavalue.LuaClosure), nArgs, nResults)
	}
	panic("LuaState Call: 目标索引处的值并非函数")
}

// callLuaClosure s
func (me *LuaState) callLuaClosure(closure *luavalue.LuaClosure, nArgs, nResults int) {

}
