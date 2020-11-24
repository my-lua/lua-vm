package state

import (
	"fmt"

	"../luavalue"
)

// LuaStack Lua栈
type LuaStack struct {
	top     int
	slots   []luavalue.ILuaValue
	prev    *LuaStack
	closure *luavalue.LuaClosure
	varArgs []luavalue.ILuaValue
	pc      int
}

// Top 获取栈顶位置索引（内部索引）
func (me *LuaStack) Top() int {
	return me.top
}

// Prev 获取上一个函数调用栈
func (me *LuaStack) Prev() *LuaStack {
	return me.prev
}

// Closure 栈所属于的闭包
func (me *LuaStack) Closure() *luavalue.LuaClosure {
	return me.closure
}

// VarArgs 可变参数列表
func (me *LuaStack) VarArgs() []luavalue.ILuaValue {
	return me.varArgs
}

// PC 程序计数器
func (me *LuaStack) PC() int {
	return me.pc
}

// Size 获取栈容量
func (me *LuaStack) Size() int {
	return len(me.slots)
}

// Check 检查栈中是否有n个或以上空闲空间，如果没有则扩容
func (me *LuaStack) Check(n int) {
	var diff = n - (me.Size() - me.top)
	for i := 0; i < diff; i++ {
		me.slots = append(me.slots, luavalue.NewLuaNil())
	}
}

// Push 向栈顶压入一个Lua值
func (me *LuaStack) Push(value luavalue.ILuaValue) {
	if me.top >= me.Size() {
		panic("LuaStack Push: 栈空间不足")
	}
	me.slots[me.top] = value
	me.top++
}

// Pop 从栈顶弹出一个Lua值
func (me *LuaStack) Pop() luavalue.ILuaValue {
	if me.top < 1 {
		panic("LuaStack Pop: 栈已为空")
	}
	me.top--
	result := me.slots[me.top]
	me.slots[me.top] = luavalue.NewLuaNil()
	return result
}

// PushN 连续推入n个Lua值
func (me *LuaStack) PushN(values []luavalue.ILuaValue, n int) {
	nValues := len(values)
	if n < 0 {
		n = nValues
	}
	for i := 0; i < n; i++ {
		if i < nValues {
			me.Push(values[i])
		} else {
			me.Push(luavalue.NewLuaNil())
		}
	}
}

// PopN 连续弹出n个Lua值
func (me *LuaStack) PopN(n int) []luavalue.ILuaValue {
	result := make([]luavalue.ILuaValue, n)
	for i := n - 1; i >= 0; i-- {
		result[i] = me.Pop()
	}
	return result
}

// AbsIndex 相对索引转换成为绝对索引（内部索引）
func (me *LuaStack) AbsIndex(index int) int {
	if index >= 0 {
		return index
	}
	return me.top + index + 1
}

// IndexIsValid 判断索引是否在有效范围之内（外部索引）
func (me *LuaStack) IndexIsValid(index int) bool {
	absIndex := me.AbsIndex(index)
	return absIndex > 0 && absIndex <= me.top
}

// Get 根据索引从栈内获取Lua值（外部索引）
func (me *LuaStack) Get(index int) luavalue.ILuaValue {
	if me.IndexIsValid(index) {
		return me.slots[me.AbsIndex(index)-1]
	}
	return luavalue.NewLuaNil()
}

// Set 设置Lua值到栈中指定索引（外部索引）
func (me *LuaStack) Set(index int, value luavalue.ILuaValue) {
	if me.IndexIsValid(index) {
		me.slots[me.AbsIndex(index)-1] = value
		return
	}
	panic("LuaStack Set: 错误的索引")
}

// Reverse 特定栈区间倒序
func (me *LuaStack) Reverse(fromIndex, toIndex int) {
	slots := me.slots
	for fromIndex < toIndex {
		slots[fromIndex], slots[toIndex] = slots[toIndex], slots[fromIndex]
		fromIndex++
		toIndex--
	}
}

// NewLuaStack 构造函数
func NewLuaStack(size int) *LuaStack {
	return &LuaStack{
		top:     0,
		slots:   make([]luavalue.ILuaValue, size),
		prev:    nil,
		closure: nil,
		varArgs: []luavalue.ILuaValue{},
		pc:      0,
	}
}

// Print 打印栈内值
func (me *LuaStack) Print() {
	for _, value := range me.slots[:me.top] {
		fmt.Printf("[%s]", value.GetString())
	}
	fmt.Println("")
}
