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

// OpCode 获取指令操作码（左边6个字节）
func (me LuaInstruction) OpCode() ELuaInstruction {
	return ELuaInstruction(uint32(me) & 0x3f)
}

// OpName 获取指令名称
func (me LuaInstruction) OpName() string {
	info := LuaInstructionInfoTable(int(me.OpCode()))
	return info.Name()
}

// OpMode 获取指令模式
func (me LuaInstruction) OpMode() ELuaInstructionMode {
	info := LuaInstructionInfoTable(int(me.OpCode()))
	return info.OpMode()
}

// BMode 获取操作数B的使用模式
func (me LuaInstruction) BMode() ELuaInstructionArgMode {
	info := LuaInstructionInfoTable(int(me.OpCode()))
	return info.ArgBMode()
}

// CMode 获取操作数C的使用模式
func (me LuaInstruction) CMode() ELuaInstructionArgMode {
	info := LuaInstructionInfoTable(int(me.OpCode()))
	return info.ArgCMode()
}
