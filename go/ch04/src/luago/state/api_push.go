package state

// PushNil [-0, +1, –]
// http://www.lua.org/manual/5.3/manual.html#lua_pushnil
func (s *luaState) PushNil() {
	s.stack.push(nil)
}

// PushBoolean [-0, +1, –]
// http://www.lua.org/manual/5.3/manual.html#lua_pushboolean
func (s *luaState) PushBoolean(b bool) {
	s.stack.push(b)
}

// PushInteger [-0, +1, –]
// http://www.lua.org/manual/5.3/manual.html#lua_pushinteger
func (s *luaState) PushInteger(n int64) {
	s.stack.push(n)
}

// PushNumber [-0, +1, –]
// http://www.lua.org/manual/5.3/manual.html#lua_pushnumber
func (s *luaState) PushNumber(n float64) {
	s.stack.push(n)
}

// PushString [-0, +1, m]
// http://www.lua.org/manual/5.3/manual.html#lua_pushstring
func (s *luaState) PushString(str string) {
	s.stack.push(str)
}
