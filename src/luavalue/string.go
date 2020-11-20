package luavalue

// LuaString s
type LuaString string

// Type s
func (me *LuaString) Type() ELuaType {
	return LuaTypeString
}

// Value s
func (me *LuaString) Value() interface{} {
	return *me
}

// ValueSrc s
func (me *LuaString) ValueSrc() string {
	return string(*me)
}

// GetBoolean s
func (me *LuaString) GetBoolean() bool {
	if len(me.ValueSrc()) > 0 {
		return true
	}
	return false
}

// GetInteger s
func (me *LuaString) GetInteger() int64 {
	return 0
}

// GetNumber s
func (me *LuaString) GetNumber() float64 {
	return 0.0
}

// GetString s
func (me *LuaString) GetString() string {
	return me.ValueSrc()
}

// ToLuaBoolean s
func (me *LuaString) ToLuaBoolean() LuaBoolean {
	return NewLuaBoolean(me.GetBoolean())
}

// ToLuaInteger s
func (me *LuaString) ToLuaInteger() LuaInteger {
	return NewLuaInteger(me.GetInteger())
}

// ToLuaNumber s
func (me *LuaString) ToLuaNumber() LuaNumber {
	return NewLuaNumber(me.GetNumber())
}

// ToLuaString s
func (me *LuaString) ToLuaString() LuaString {
	return NewLuaString(me.GetString())
}

// NewLuaString s
func NewLuaString(value string) LuaString {
	return LuaString(value)
}
