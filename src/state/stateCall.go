package state

import (
	"../chunk"
	"../luavalue"
)

// Load 构建出LuaClosure并且推入栈中
func (me *LuaState) Load(chunkBuffer []byte, chunkName, mode string) int {
	// 这里需要根据chunk生成函数原型
	var proto *chunk.Prototype = nil
	// 根据函数原型加工出Lua闭包
	closure := luavalue.NewLuaClosure(proto)
	// 闭包推入栈中
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
