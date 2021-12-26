package state

import (
	"fmt"
	. "luagogo/ch04/src/luago/api"
	"testing"
)

func TestNew(t *testing.T) {
	//  [true]
	// 	[true][10]
	// 	[true][10][nil]
	// 	[true][10][nil]["hello"]
	// 	[true][10][nil]["hello"][true]
	// 	[true][10][true]["hello"]
	// 	[true][10][true]["hello"][nil][nil]
	// 	[true][10][true][nil][nil]
	// 	[true]
	ls := New()

	ls.PushBoolean(true)
	printStack(ls) // [true]
	ls.PushInteger(10)
	printStack(ls) // [true][10]
	ls.PushNil()
	printStack(ls) // [true][10][nil]
	ls.PushString("hello")
	printStack(ls) // [true][10][nil]["hello"]
	ls.PushValue(-4)
	printStack(ls) // [true][10][nil]["hello"][true]
	ls.Replace(3)
	printStack(ls) // [true][10][true]["hello"]
	ls.SetTop(6)
	printStack(ls) // [true][10][true]["hello"][nil][nil]
	ls.Remove(-3)
	printStack(ls) // [true][10][true][nil][nil]
	ls.SetTop(-5)
	printStack(ls) // [true]
}

func printStack(ls LuaState) {
	top := ls.GetTop()
	for i := 1; i <= top; i++ {
		t := ls.Type(i)
		switch t {
		case LUA_TBOOLEAN:
			fmt.Printf("[%t]", ls.ToBoolean(i))
		case LUA_TNUMBER:
			fmt.Printf("[%g]", ls.ToNumber(i))
		case LUA_TSTRING:
			fmt.Printf("[%q]", ls.ToString(i))
		default: // other values
			fmt.Printf("[%s]", ls.TypeName(t))
		}
	}
	fmt.Println()
}
