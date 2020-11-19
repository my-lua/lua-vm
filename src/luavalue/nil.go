package luavalue

// LuaNil Lua nil值
type LuaNil interface{}

// Type s
func (me LuaNil) Type() ELuaType {
	return LuaTypeBoolean
}

// Value s
func (me *LuaNil) Value() interface{} {
	return *me
}

// ValueSrc s
func (me *LuaNil) ValueSrc() interface{} {
	return nil
}
