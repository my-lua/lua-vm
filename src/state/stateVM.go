package state

// PC 当前程序计数器
func (me *LuaState) PC() int {
	return me.stack.pc
}

// AddPC 增加程序计数器
func (me *LuaState) AddPC(n int) {
	me.stack.pc += n
}

// Fetch 获取当前指令并后移程序计数器
func (me *LuaState) Fetch() uint32 {
	p := me.Stack().Closure().Prototype()
	me.stack.pc++
	return 0
}
