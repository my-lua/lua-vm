package luavalue

// LuaClosure Lua闭包（函数）
type LuaClosure struct {
}

// Type s
func (me *LuaClosure) Type() ELuaType {
	return LuaTypeFunction
}

// Value s
func (me *LuaClosure) Value() interface{} {
	return *me
}

// ValueSrc s
func (me *LuaClosure) ValueSrc() LuaClosure {
	return me.Value().(LuaClosure)
}

// NewLuaClosure 构造函数
func NewLuaClosure() *LuaClosure {
	return &LuaClosure{}
}
