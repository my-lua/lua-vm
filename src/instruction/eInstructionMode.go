package instruction

// ELuaInstructionMode 指令编码模式枚举
type ELuaInstructionMode int

const (
	// IABC 6 8 9 9
	IABC ELuaInstructionMode = iota
	// IABx 6 8 18
	IABx
	// IAsBx 6 8 18
	IAsBx
	// IAx 6 26
	IAx
)
