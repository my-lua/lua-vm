package state

import (
	"fmt"

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

// callLuaClosure 调用闭包
func (me *LuaState) callLuaClosure(closure *luavalue.LuaClosure, nArgs, nResults int) {
	nRegs := int(closure.Prototype().MaxStackSize())
	nParams := int(closure.Prototype().NumParams())
	isVararg := closure.Prototype().IsVararg() == 1

	newStack := NewLuaStack(nRegs + 20)
	newStack.closure = closure

	funcAndArgs := me.stack.PopN(nArgs + 1)
	newStack.PushN(funcAndArgs[1:], nParams)
	newStack.top = nArgs
	if nArgs > nParams && isVararg {
		newStack.varArgs = funcAndArgs[nParams+1:]
	}

	me.PushStack(newStack)
	// 这里需要运行闭包
	me.PopStack()

	if nResults > 0 {

	}

	fmt.Println(nRegs, nParams, isVararg)
}
