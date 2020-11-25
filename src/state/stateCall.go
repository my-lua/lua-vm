package state

import (
	"fmt"

	"../luavalue"
)

// Load 构建出LuaClosure并且推入栈中
func (me *LuaState) Load(chunkBuffer []byte, chunkName, mode string) int {
	// 这里需要根据chunk生成函数原型
	var proto *luavalue.Prototype = nil
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

func (me *LuaState) runLuaClosure() {
	for {
		break
	}
}

// callLuaClosure 调用闭包
func (me *LuaState) callLuaClosure(closure *luavalue.LuaClosure, nArgs, nResults int) {
	nRegs := int(closure.Prototype().MaxStackSize())
	nParams := int(closure.Prototype().NumParams())
	isVararg := closure.Prototype().IsVararg() == 1

	newStack := NewLuaStack(nRegs + 20)
	newStack.closure = closure

	funcAndArgs := me.stack.PopN(nArgs + 1)
	// 函数参数入栈
	newStack.PushN(funcAndArgs[1:], nParams)
	// 根据函数原型内的MaxStackSize设置栈顶
	newStack.top = nRegs
	// 如果实际传入参数数量大于函数原型内定义的参数数量，且函数原型内isVararg为真
	// 则收集varArgs参数到新的栈
	if nArgs > nParams && isVararg {
		newStack.varArgs = funcAndArgs[nParams+1:]
	}

	me.PushStack(newStack)
	// 这里需要运行闭包
	me.runLuaClosure()
	me.PopStack()

	if nResults > 0 {
		// 获取函数运行后留在栈顶的参数
		results := newStack.PopN(newStack.Top() - nRegs)
		// 检查或尝试扩容主调栈
		me.stack.Check(len(results))
		// 推入主调栈，多退少补
		me.stack.PushN(results, nResults)
	}

	fmt.Println(nRegs, nParams, isVararg)
}
