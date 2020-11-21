package luavalue

// LuaTable s
type LuaTable struct {
	tArray []ILuaValue
	tMap   map[ILuaValue]ILuaValue
}

// Get 根据Key从表内取出某个Lua值
func (me *LuaTable) Get(key ILuaValue) ILuaValue {
	var index int64 = -1
	switch key.Type() {
	case LuaTypeInteger:
		index = key.(*LuaInteger).GetInteger()
		break
	case LuaTypeNumber:
		index = int64(key.(*LuaNumber).GetNumber())
		break
	}
	if index >= 1 && index <= int64(len(me.tArray)) {
		return me.tArray[index-1]
	}
	result := me.tMap[key]
	if result == nil {
		return NewLuaNil()
	}
	return result
}

// Put 根据Key和Value设置某个Lua值到表内
func (me *LuaTable) Put(key, value ILuaValue) {

}

// Type s
func (me *LuaTable) Type() ELuaType {
	return LuaTypeTable
}

// Value s
func (me *LuaTable) Value() interface{} {
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
func (me *LuaTable) ToLuaBoolean() *LuaBoolean {
	return NewLuaBoolean(me.GetBoolean())
}

// ToLuaInteger s
func (me *LuaTable) ToLuaInteger() *LuaInteger {
	return NewLuaInteger(me.GetInteger())
}

// ToLuaNumber s
func (me *LuaTable) ToLuaNumber() *LuaNumber {
	return NewLuaNumber(me.GetNumber())
}

// ToLuaString s
func (me *LuaTable) ToLuaString() *LuaString {
	return NewLuaString(me.GetString())
}

// NewLuaTable s
func NewLuaTable(nArr, nRec int) *LuaTable {
	result := &LuaTable{}
	if nArr > 0 {
		result.tArray = make([]ILuaValue, 0, nArr)
	}
	if nRec > 0 {
		result.tMap = make(map[ILuaValue]ILuaValue, nRec)
	}
	return result
}
