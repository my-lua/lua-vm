package state

// PushStack 函数调用栈压栈
func (me *LuaState) PushStack(newStack *LuaStack) {
	newStack.prev = me.stack
	me.stack = newStack
}

// PopStack 函数调用栈弹栈
func (me *LuaState) PopStack() {
	oldStack := me.stack
	if me.stack != nil {
		me.stack = me.stack.prev
		oldStack.prev = nil
	}
	panic("LuaState PopStack: 栈已为空")
}
