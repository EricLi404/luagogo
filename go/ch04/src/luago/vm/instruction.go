package vm

const MAXARG_Bx = 1<<18 - 1       // 262143
const MAXARG_sBx = MAXARG_Bx >> 1 // 131071

/*
 31       22       13       5    0
  +-------+^------+-^-----+-^-----
  |b=9bits |c=9bits |a=8bits|op=6|  iABC
  +-------+^------+-^-----+-^-----
  |    bx=18bits    |a=8bits|op=6|  iABx
  +-------+^------+-^-----+-^-----
  |   sbx=18bits    |a=8bits|op=6|  iAsBx
  +-------+^------+-^-----+-^-----
  |    ax=26bits            |op=6|  iAx
  +-------+^------+-^-----+-^-----
 31      23      15       7      0
*/

// Instruction 一条虚拟机指令，对应 binchunk.Prototype.Code
//             在 此模块中 解析时，为了便于操作，转换为了 OpCode 来操作
type Instruction uint32

// OpCodesIdx 获取指令操作码， 也即 OpCodes 数组的下标
func (i Instruction) OpCodesIdx() int {
	return int(i & 0x3F)
}

// ABC ，IABC 编码模式 提取参数
func (i Instruction) ABC() (a, b, c int) {
	a = int(i >> 6 & 0xFF)
	c = int(i >> 14 & 0x1FF)
	b = int(i >> 23 & 0x1FF)
	return
}

// ABx ，IABx 编码模式 提取参数
func (i Instruction) ABx() (a, bx int) {
	a = int(i >> 6 & 0xFF)
	bx = int(i >> 14)
	return
}

// AsBx ， IAsBx 编码模式 提取参数
func (i Instruction) AsBx() (a, sbx int) {
	a, bx := i.ABx()
	// 偏移二进制码， 也叫 Excess-K ，将有符号整数编码为比特序列
	return a, bx - MAXARG_sBx
}

// Ax , IAx 编码模式 提取参数
func (i Instruction) Ax() int {
	return int(i >> 6)
}

// OpName 获取 指令的 名称,
func (i Instruction) OpName() string {
	return OpCodes[i.OpCodesIdx()].name
}

// OpMode 获取指令 编码模式
func (i Instruction) OpMode() byte {
	return OpCodes[i.OpCodesIdx()].opMode
}

// BMode 获取 参数B 的操作模式
func (i Instruction) BMode() byte {
	return OpCodes[i.OpCodesIdx()].argBMode
}

// CMode 获取 参数C 的操作模式
func (i Instruction) CMode() byte {
	return OpCodes[i.OpCodesIdx()].argCMode
}
