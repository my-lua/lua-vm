package luavalue

// LuaTable s
type LuaTable struct {
}

// Type s
func (*LuaTable) Type() ELuaType {
	return LuaTypeTable
}

// Value s
func (*LuaTable) Value() interface{} {
	return "[table]"
}

// ValueSrc s
func (me *LuaTable) ValueSrc() string {
	return me.Value().(string)
}

// GetBoolean s
func (me *LuaTable) GetBoolean() bool {
	return true
}

// GetInteger s
func (me *LuaTable) GetInteger() int64 {
	return 1
}

// GetNumber s
func (me *LuaTable) GetNumber() float64 {
	return 1.0
}

// GetString s
func (me *LuaTable) GetString() string {
	return me.ValueSrc()
}

// ToLuaBoolean s
func (me *LuaTable) ToLuaBoolean() LuaBoolean {
	return NewLuaBoolean(me.GetBoolean())
}

// ToLuaInteger s
func (me *LuaTable) ToLuaInteger() LuaInteger {
	return NewLuaInteger(me.GetInteger())
}

// ToLuaNumber s
func (me *LuaTable) ToLuaNumber() LuaNumber {
	return NewLuaNumber(me.GetNumber())
}

// ToLuaString s
func (me *LuaTable) ToLuaString() LuaString {
	return NewLuaString(me.GetString())
}

// NewLuaTable s
func NewLuaTable(nArr, nRec int) {

}
