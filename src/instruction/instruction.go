package instruction

import "fmt"

// LuaInstruction Lua指令（为一个32位无符号数）
type LuaInstruction uint32

// ABC 获取ABC模式下的参数（6，8，9，9）
func (me LuaInstruction) ABC() (a, b, c int) {
	a = int(uint32(me) >> 6 & 0xff)
	c = int(uint32(me) >> 14 & 0x1ff)
	b = int(uint32(me) >> 23)
	return a, b, c
}

// ABx 获取ABx模式下的参数（6，8，18）
func (me LuaInstruction) ABx() (a, bx int) {
	a = int(uint32(me) >> 6 & 0xff)
	bx = int(uint32(me) >> 14)
	return a, bx
}

// MaxArgBx （比如说255）
const MaxArgBx = 1<<18 - 1

// MaxArgSBx （比如说127）
const MaxArgSBx = MaxArgBx >> 1

// AsBx 获取AsBx模式下的参数（6，8，18）
func (me LuaInstruction) AsBx() (a, sbx int) {
	a, bx := me.ABx()
	// 比如[0 ~ 126][127][128 ~ 255]，可以表示127个负数和128个正数，一个0
	return a, bx - MaxArgSBx
}

// Ax 获取Ax模式下的参数（6，26）
func (me LuaInstruction) Ax() int {
	return int(uint32(me) >> 6)
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

// Operands 获取指令操作数信息
func (me LuaInstruction) Operands() string {
	switch me.OpMode() {
	case IABC:
		a, b, c := me.ABC()
		var outText = fmt.Sprintf("%d", a)
		if me.BMode() != InstructionArgN {
			// 9位数的最高位为1的话，视为常量表索引，低8位按负数输出
			if b > 0xff {
				outText += fmt.Sprintf(" %d", -1-(b&0xff))
			} else {
				outText += fmt.Sprintf(" %d", b)
			}
		}
		if me.CMode() != InstructionArgN {
			if c > 0xff {
				outText += fmt.Sprintf(" %d", -1-(c&0xff))
			} else {
				outText += fmt.Sprintf(" %d", c)
			}
		}
		return fmt.Sprintf("%s [ABC]", outText)
	case IABx:
		a, bx := me.ABx()
		var outText = fmt.Sprintf("%d", a)
		if me.BMode() == InstructionArgK {
			outText += fmt.Sprintf(" %d", -1-bx)
		} else if me.BMode() == InstructionArgU {
			outText += fmt.Sprintf(" %d", bx)
		}
		return fmt.Sprintf("%s [ABx]", outText)
	case IAsBx:
		a, sbx := me.AsBx()
		return fmt.Sprintf("%d %d [AsBx]", a, sbx)
	case IAx:
		ax := me.Ax()
		return fmt.Sprintf("%d [Ax]", -1-ax)
	default:
		return ""
	}
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
