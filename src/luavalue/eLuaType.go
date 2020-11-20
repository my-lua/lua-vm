package luavalue

// ELuaType Lua数据类型枚举
type ELuaType int

const (
	// LuaTypeNone -1
	LuaTypeNone ELuaType = iota - 1
	LuaTypeNil
	LuaTypeBoolean
	LuaTypeInteger
	LuaTypeNumber
	LuaTypeString
	LuaTypeTable
	LuaTypeFunction
	LuaTypeUserData
	LuaTypeLightUserData
	LuaTypeThread
)
