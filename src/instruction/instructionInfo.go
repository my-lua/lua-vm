package instruction

// LuaInstructionInfo 指令信息
type LuaInstructionInfo struct {
	testFlag byte
	setAFlag byte
	argBMode ELuaInstructionArgMode
	argCMode ELuaInstructionArgMode
	opMode   ELuaInstructionMode
	name     string
}

// TestFlag s
func (me *LuaInstructionInfo) TestFlag() byte {
	return me.testFlag
}

// SetAFlag s
func (me *LuaInstructionInfo) SetAFlag() byte {
	return me.setAFlag
}

// ArgBMode s
func (me *LuaInstructionInfo) ArgBMode() ELuaInstructionArgMode {
	return me.argBMode
}

// ArgCMode s
func (me *LuaInstructionInfo) ArgCMode() ELuaInstructionArgMode {
	return me.argCMode
}

// OpMode s
func (me *LuaInstructionInfo) OpMode() ELuaInstructionMode {
	return me.opMode
}

// Name s
func (me *LuaInstructionInfo) Name() string {
	return me.name
}

// LuaInstructionInfoTable 查询指令信息表
func LuaInstructionInfoTable(index int) LuaInstructionInfo {
	if index >= 0 && index < len(luaInstructionInfoTable) {
		return luaInstructionInfoTable[index]
	}
	panic("LuaInstructionInfo LuaInstructionInfoTable: 索引无效")
}

// luaInstructionInfoTable 指令信息表
var luaInstructionInfoTable = []LuaInstructionInfo{
	/*                 T  A  B                C                mode        name    */
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgN, IABC /* */, "MOVE    "}, // R(A) := R(B)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgN, IABx /* */, "LOADK   "}, // R(A) := Kst(Bx)
	LuaInstructionInfo{0, 1, InstructionArgN, InstructionArgN, IABx /* */, "LOADKX  "}, // R(A) := Kst(extra arg)
	LuaInstructionInfo{0, 1, InstructionArgU, InstructionArgU, IABC /* */, "LOADBOOL"}, // R(A) := (bool)B; if (C) pc++
	LuaInstructionInfo{0, 1, InstructionArgU, InstructionArgN, IABC /* */, "LOADNIL "}, // R(A), R(A+1), ..., R(A+B) := nil
	LuaInstructionInfo{0, 1, InstructionArgU, InstructionArgN, IABC /* */, "GETUPVAL"}, // R(A) := UpValue[B]
	LuaInstructionInfo{0, 1, InstructionArgU, InstructionArgK, IABC /* */, "GETTABUP"}, // R(A) := UpValue[B][RK(C)]
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgK, IABC /* */, "GETTABLE"}, // R(A) := R(B)[RK(C)]
	LuaInstructionInfo{0, 0, InstructionArgK, InstructionArgK, IABC /* */, "SETTABUP"}, // UpValue[A][RK(B)] := RK(C)
	LuaInstructionInfo{0, 0, InstructionArgU, InstructionArgN, IABC /* */, "SETUPVAL"}, // UpValue[B] := R(A)
	LuaInstructionInfo{0, 0, InstructionArgK, InstructionArgK, IABC /* */, "SETTABLE"}, // R(A)[RK(B)] := RK(C)
	LuaInstructionInfo{0, 1, InstructionArgU, InstructionArgU, IABC /* */, "NEWTABLE"}, // R(A) := {} (size = B,C)
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgK, IABC /* */, "SELF    "}, // R(A+1) := R(B); R(A) := R(B)[RK(C)]
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "ADD     "}, // R(A) := RK(B) + RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "SUB     "}, // R(A) := RK(B) - RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "MUL     "}, // R(A) := RK(B) * RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "MOD     "}, // R(A) := RK(B) % RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "POW     "}, // R(A) := RK(B) ^ RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "DIV     "}, // R(A) := RK(B) / RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "IDIV    "}, // R(A) := RK(B) // RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "BAND    "}, // R(A) := RK(B) & RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "BOR     "}, // R(A) := RK(B) | RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "BXOR    "}, // R(A) := RK(B) ~ RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "SHL     "}, // R(A) := RK(B) << RK(C)
	LuaInstructionInfo{0, 1, InstructionArgK, InstructionArgK, IABC /* */, "SHR     "}, // R(A) := RK(B) >> RK(C)
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgN, IABC /* */, "UNM     "}, // R(A) := -R(B)
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgN, IABC /* */, "BNOT    "}, // R(A) := ~R(B)
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgN, IABC /* */, "NOT     "}, // R(A) := not R(B)
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgN, IABC /* */, "LEN     "}, // R(A) := length of R(B)
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgR, IABC /* */, "CONCAT  "}, // R(A) := R(B).. ... ..R(C)
	LuaInstructionInfo{0, 0, InstructionArgR, InstructionArgN, IAsBx /**/, "JMP     "}, // pc+=sBx; if (A) close all upvalues >= R(A - 1)
	LuaInstructionInfo{1, 0, InstructionArgK, InstructionArgK, IABC /* */, "EQ      "}, // if ((RK(B) == RK(C)) ~= A) then pc++
	LuaInstructionInfo{1, 0, InstructionArgK, InstructionArgK, IABC /* */, "LT      "}, // if ((RK(B) <  RK(C)) ~= A) then pc++
	LuaInstructionInfo{1, 0, InstructionArgK, InstructionArgK, IABC /* */, "LE      "}, // if ((RK(B) <= RK(C)) ~= A) then pc++
	LuaInstructionInfo{1, 0, InstructionArgN, InstructionArgU, IABC /* */, "TEST    "}, // if not (R(A) <=> C) then pc++
	LuaInstructionInfo{1, 1, InstructionArgR, InstructionArgU, IABC /* */, "TESTSET "}, // if (R(B) <=> C) then R(A) := R(B) else pc++
	LuaInstructionInfo{0, 1, InstructionArgU, InstructionArgU, IABC /* */, "CALL    "}, // R(A), ... ,R(A+C-2) := R(A)(R(A+1), ... ,R(A+B-1))
	LuaInstructionInfo{0, 1, InstructionArgU, InstructionArgU, IABC /* */, "TAILCALL"}, // return R(A)(R(A+1), ... ,R(A+B-1))
	LuaInstructionInfo{0, 0, InstructionArgU, InstructionArgN, IABC /* */, "RETURN  "}, // return R(A), ... ,R(A+B-2)
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgN, IAsBx /**/, "FORLOOP "}, // R(A)+=R(A+2); if R(A) <?= R(A+1) then { pc+=sBx; R(A+3)=R(A) }
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgN, IAsBx /**/, "FORPREP "}, // R(A)-=R(A+2); pc+=sBx
	LuaInstructionInfo{0, 0, InstructionArgN, InstructionArgU, IABC /* */, "TFORCALL"}, // R(A+3), ... ,R(A+2+C) := R(A)(R(A+1), R(A+2));
	LuaInstructionInfo{0, 1, InstructionArgR, InstructionArgN, IAsBx /**/, "TFORLOOP"}, // if R(A+1) ~= nil then { R(A)=R(A+1); pc += sBx }
	LuaInstructionInfo{0, 0, InstructionArgU, InstructionArgU, IABC /* */, "SETLIST "}, // R(A)[(C-1)*FPF+i] := R(A+i), 1 <= i <= B
	LuaInstructionInfo{0, 1, InstructionArgU, InstructionArgN, IABx /* */, "CLOSURE "}, // R(A) := closure(KPROTO[Bx])
	LuaInstructionInfo{0, 1, InstructionArgU, InstructionArgN, IABC /* */, "VARARG  "}, // R(A), R(A+1), ..., R(A+B-2) = vararg
	LuaInstructionInfo{0, 0, InstructionArgU, InstructionArgU, IAx /*  */, "EXTRAARG"}, // extra (larger) argument for previous Code
}
