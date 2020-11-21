package luavalue

// LuaBoolean Lua布尔值
type LuaBoolean bool

// Type s
func (me LuaBoolean) Type() ELuaType {
	return LuaTypeBoolean
}

// Value s
func (me LuaBoolean) Value() interface{} {
	return me
}

// ValueSrc s
func (me LuaBoolean) ValueSrc() bool {
	return bool(me)
}

// GetBoolean s
func (me LuaBoolean) GetBoolean() bool {
	return me.ValueSrc()
}

// GetInteger s
func (me LuaBoolean) GetInteger() int64 {
	if me.ValueSrc() {
		return 1
	}
	return 0
}

// GetNumber s
func (me LuaBoolean) GetNumber() float64 {
	if me.ValueSrc() {
		return 1.0
	}
	return 0.0
}

// GetString s
func (me LuaBoolean) GetString() string {
	if me.ValueSrc() {
		return "true"
	}
	return "false"
}

// ToLuaBoolean s
func (me LuaBoolean) ToLuaBoolean() LuaBoolean {
	return NewLuaBoolean(me.GetBoolean())
}

// ToLuaInteger s
func (me LuaBoolean) ToLuaInteger() LuaInteger {
	return NewLuaInteger(me.GetInteger())
}

// ToLuaNumber s
func (me LuaBoolean) ToLuaNumber() LuaNumber {
	return NewLuaNumber(me.GetNumber())
}

// ToLuaString s
func (me LuaBoolean) ToLuaString() LuaString {
	return NewLuaString(me.GetString())
}

// NewLuaBoolean 构造函数
func NewLuaBoolean(value bool) LuaBoolean {
	return LuaBoolean(value)
}
