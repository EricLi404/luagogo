package state

type luaStack struct {
	slots []luaValue // 值
	top   int        // 栈顶索引
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

// 检查 栈中是否还能再推入 n 个值，如果不能，则为栈扩容
func (s *luaStack) check(n int) {
	free := len(s.slots) - s.top
	for i := free; i < n; i++ {
		s.slots = append(s.slots, nil)
	}
}

func (s *luaStack) push(val luaValue) {
	if s.top == len(s.slots) {
		panic("stack overflow!")
	}
	s.slots[s.top] = val
	s.top++
}

func (s *luaStack) pop() luaValue {
	if s.top < 1 {
		panic("stack underflow!")
	}
	s.top--
	val := s.slots[s.top]
	s.slots[s.top] = nil
	return val
}

func (s *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return idx + s.top + 1
}

func (s *luaStack) isValid(idx int) bool {
	absIdx := s.absIndex(idx)
	return absIdx > 0 && absIdx <= s.top
}

func (s *luaStack) get(idx int) luaValue {
	absIdx := s.absIndex(idx)
	if absIdx > 0 && absIdx <= s.top {
		return s.slots[absIdx-1]
	}
	return nil
}

func (s *luaStack) set(idx int, val luaValue) {
	absIdx := s.absIndex(idx)
	if absIdx > 0 && absIdx <= s.top {
		s.slots[absIdx-1] = val
		return
	}
	panic("invalid index!")
}

func (s *luaStack) reverse(from, to int) {
	slots := s.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
