package luavalue

// ILuaValue Lua值约束接口
type ILuaValue interface {
	Type() ELuaType
	Value() interface{}

	ToLuaBoolean() *LuaBoolean
	ToLuaInteger() *LuaInteger
	ToLuaNumber() *LuaNumber
	ToLuaString() *LuaString

	GetBoolean() bool
	GetInteger() int64
	GetNumber() float64
	GetString() string
}
