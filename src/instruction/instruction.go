package instruction

// LuaInstruction Lua指令
type LuaInstruction uint32

// ABC 获取ABC模式下的参数（6，8，9，9）
func (me LuaInstruction) ABC() (a, b, c int) {
	a = int(me >> 6 & 0xff)
	c = int(me >> 14 & 0x1ff)
	b = int(me >> 23)
	return a, b, c
}

// ABx 获取ABx模式下的参数（6，8，18）
func (me LuaInstruction) ABx() (a, bx int) {
	a = int(me >> 6 & 0xff)
	bx = int(me >> 14)
	return a, bx
}

// Ax 获取Ax模式下的参数（6，26）
func (me LuaInstruction) Ax() int {
	return int(me >> 6)
}
