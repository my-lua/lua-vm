package state

import "../luavalue"

func (me *LuaState) Type(index int) luavalue.ELuaType {
	return luavalue.LuaTypeNil
}

func (me *LuaState) IsNone(index int) bool {
	return false
}

func (me *LuaState) IsNil(index int) bool {
	return false
}

func (me *LuaState) IsNoneOrNil(index int) bool {
	return false
}

func (me *LuaState) IsBoolean(index int) bool {
	return false
}

func (me *LuaState) IsInteger(index int) bool {
	return false
}

func (me *LuaState) IsNumber(index int) bool {
	return false
}

func (me *LuaState) IsString(index int) bool {
	return false
}

func (me *LuaState) IsTable(index int) bool {
	return false
}

func (me *LuaState) IsFunction(index int) bool {
	return false
}

func (me *LuaState) IsThread(index int) bool {
	return false
}
