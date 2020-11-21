package luavalue

// LuaTable s
type LuaTable struct {
	tArray []ILuaValue
	tMap   map[ILuaValue]ILuaValue
}

// Len s
func (me *LuaTable) Len() int64 {
	return int64(len(me.tArray))
}

// AbsIndex s
func (me *LuaTable) AbsIndex(index int64) int64 {
	return index - 1
}

// IndexIsValid s
func (me *LuaTable) IndexIsValid(index int64) bool {
	absIndex := me.AbsIndex(index)
	if absIndex >= 0 && absIndex < me.Len() {
		return true
	}
	return false
}

// Get 根据Key从表内取出某个Lua值
func (me *LuaTable) Get(key ILuaValue) ILuaValue {
	if (key.Type() == LuaTypeInteger ||
		key.Type() == LuaTypeNumber) &&
		me.IndexIsValid(key.GetInteger()) {
		return me.tArray[me.AbsIndex(key.GetInteger())]
	}
	if result, found := me.tMap[key]; found {
		return result
	}
	return NewLuaNil()
}

func (me *LuaTable) shrinkArray() {
	for i := len(me.tArray) - 1; i >= 0; i-- {
		curItem := me.tArray[i]
		if curItem.Type() == LuaTypeNil {
			me.tArray = me.tArray[:i]
		} else {
			break
		}
	}
}

func (me *LuaTable) expandArray() {
	for i := len(me.tArray) + 1; true; i++ {
		key := NewLuaInteger(int64(i))
		if value, found := me.tMap[key]; found {
			delete(me.tMap, key)
			me.tArray = append(me.tArray, value)
		} else {
			break
		}
	}
}

// Put 根据Key和Value设置某个Lua值到表内
func (me *LuaTable) Put(key, value ILuaValue) {
	if key.Type() == LuaTypeInteger || key.Type() == LuaTypeNumber {
		index := key.GetInteger()
		absIndex := me.AbsIndex(index)
		// 如果索引在数组范围内
		if me.IndexIsValid(index) {
			me.tArray[absIndex] = value
			if absIndex == me.Len()-1 && value.Type() == LuaTypeNil {
				me.shrinkArray()
			}
			return
		} else if absIndex == me.Len() {
			delete(me.tMap, key)
			if value.Type() != LuaTypeNil {
				me.tArray = append(me.tArray, value)
				me.expandArray()
			}
			return
		}
	}

	if value.Type() != LuaTypeNil {
		if me.tMap == nil {
			me.tMap = make(map[ILuaValue]ILuaValue, 8)
		}
		me.tMap[key] = value
	} else {
		delete(me.tMap, key)
	}
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

// NewLuaTable 构造函数
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
