package state

import (
	"../instruction"
)

// PC 当前程序计数器
func (me *LuaState) PC() int {
	return me.Stack().PC()
}

// AddPC 增加程序计数器
func (me *LuaState) AddPC(n int) {
	me.stack.pc += n
}

// Fetch 获取当前指令并后移程序计数器
func (me *LuaState) Fetch() instruction.LuaInstruction {
	inst := me.Stack().Closure().Prototype().Instructions()[me.PC()]
	me.stack.pc++
	return inst
}

// GetConst 获取常量
func (me *LuaState) GetConst(index int) {
	value := me.Stack().Closure().Prototype().Constants()[index]
	me.stack.Push(value)
}

// GetRK 获取寄存器或者常量
func (me *LuaState) GetRK(index int) {
	if index < 0xff {
		me.GetConst(index & 0xff)
	} else {
		me.PushValue(index)
	}
}

// RegisterCount 当前栈帧的最大寄存器数量
func (me *LuaState) RegisterCount() int {
	return int(me.Stack().Closure().Prototype().MaxStackSize())
}
