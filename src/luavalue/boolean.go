package luavalue

// LuaBoolean Lua布尔值
type LuaBoolean bool

// Type s
func (me *LuaBoolean) Type() ELuaType {
	return LuaTypeBoolean
}

// Value s
func (me *LuaBoolean) Value() interface{} {
	return *me
}

// ValueSrc s
func (me *LuaBoolean) ValueSrc() bool {
	return bool(*me)
}
