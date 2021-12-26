package state

// GetTop [-0, +0, –] 返回栈顶索引
// http://www.lua.org/manual/5.3/manual.html#lua_gettop
func (s *luaState) GetTop() int {
	return s.stack.top
}

// AbsIndex [-0, +0, –] 将索引转换为绝对索引
// http://www.lua.org/manual/5.3/manual.html#lua_absindex
func (s *luaState) AbsIndex(idx int) int {
	return s.stack.absIndex(idx)
}

// CheckStack [-0, +0, –] 检查容量，发现容量不够时自动扩容
// http://www.lua.org/manual/5.3/manual.html#lua_checkstack
func (s *luaState) CheckStack(n int) bool {
	s.stack.check(n)
	return true // never fails
}

// Pop [-n, +0, –] 弹出 n 个值
// http://www.lua.org/manual/5.3/manual.html#lua_pop
func (s *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		s.stack.pop()
	}
}

// Copy [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_copy
func (s *luaState) Copy(fromIdx, toIdx int) {
	val := s.stack.get(fromIdx)
	s.stack.set(toIdx, val)
}

// PushValue [-0, +1, –] 把指定索引处的值 复制一份推到 栈顶
// http://www.lua.org/manual/5.3/manual.html#lua_pushvalue
func (s *luaState) PushValue(idx int) {
	val := s.stack.get(idx)
	s.stack.push(val)
}

// Replace [-1, +0, –] 将栈顶值弹出，覆盖指定索引处的值
// http://www.lua.org/manual/5.3/manual.html#lua_replace
func (s *luaState) Replace(idx int) {
	val := s.stack.pop()
	s.stack.set(idx, val)
}

// Insert [-1, +1, –] 将栈顶值弹出，插入指定索引处
// http://www.lua.org/manual/5.3/manual.html#lua_insert
func (s *luaState) Insert(idx int) {
	s.Rotate(idx, 1)
}

// Remove [-1, +0, –] 删除指定索引处的值
// http://www.lua.org/manual/5.3/manual.html#lua_remove
func (s *luaState) Remove(idx int) {
	s.Rotate(idx, -1)
	s.Pop(1)
}

// Rotate [-0, +0, –] 将[idx,top]区间内的元素朝栈顶方向旋转 n 个位置
// e.g. [a,b,c,d,e] -->(2,1)--> [a,e,b,c,d]
// http://www.lua.org/manual/5.3/manual.html#lua_rotate
func (s *luaState) Rotate(idx, n int) {
	t := s.stack.top - 1           /* end of stack segment being rotated */
	p := s.stack.absIndex(idx) - 1 /* start of segment */
	var m int                      /* end of prefix */
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	s.stack.reverse(p, m)   /* reverse the prefix with length 'n' */
	s.stack.reverse(m+1, t) /* reverse the suffix */
	s.stack.reverse(p, t)   /* reverse the entire segment */
}

// SetTop [-?, +?, –] 将栈顶索引设置为指定值
// http://www.lua.org/manual/5.3/manual.html#lua_settop
func (s *luaState) SetTop(idx int) {
	newTop := s.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}

	n := s.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			s.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i > n; i-- {
			s.stack.push(nil)
		}
	}
}
