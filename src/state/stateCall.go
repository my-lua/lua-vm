package state

import (
	"../luavalue"
)

// Load s
func (me *LuaState) Load(chunk []byte, chunkName, mode string) int {
	closure := luavalue.NewLuaClosure()
	me.stack.Push(closure)
	return 0
}

// Call s
func (me *LuaState) Call(nArgs, nResults int) {

}
