package binchunk

const (
	LUA_SIGNATURE = "\x1bLua"
	LUAC_VERSION  = 0x53
	LUAC_FORMAT   = 0
	LUAC_DATA     = "\x19\x93\r\n\x1a\n"

	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

const (
	TAG_NIL       = 0x00 // nil
	TAG_BOOLEAN   = 0x01 // boolean ，存储 字节 0 or 1
	TAG_NUMBER    = 0x03 // number , 存储 浮点数
	TAG_INTEGER   = 0x13 // integer , 存储 整数
	TAG_SHORT_STR = 0x04 // string , 存储 短字符串
	TAG_LONG_STR  = 0x14 // string , 存储 长字符串
)

type binaryChunk struct {
	header
	sizeUpvalues byte
	mainFunc     *Prototype
}

type header struct {
	// LUA_SIGNATURE chunk 的魔数， 四字节，是 ESC、L、u、a 的ASCII 码， 0x1B4C7561  \x1bLua
	signature [4]byte

	// LUAC_VERSION 版本号，大版本号 乘以 16 加 小版本号 ，
	// 		   e.g. lua 5.4.3 => 5*16+4 = 84 ，转16进制 => 0x54
	version byte

	// LUAC_FORMAT 格式号， 官方实现中是 0x00
	format byte

	// LUAC_DATA  6个字节，依次是
	//   - lua 1.0 发布的年份（0x19 0x93）
	//   - 回车符 （0x0D）
	//   - 换行符 （0x0A）
	//   - 替换符 （0x1A）
	//   - 换行符 （0x0A）
	luacData [6]byte

	// 下边五个字节分别记录 cint、size_t、lua虚拟机指令、Lua整数、Lua浮点数占用的字节数
	cintSize        byte // CINT_SIZE
	sizetSize       byte // CSIZET_SIZE
	instructionSize byte // INSTRUCTION_SIZE
	luaIntegerSize  byte // LUA_INTEGER_SIZE
	luaNumberSize   byte // LUA_NUMBER_SIZE

	// LUAC_INT 存放 Lua 整数值 0x5678，为了检测二进制 chunk 的大小端方式
	luaInt int64

	// LUAC_NUM 存放 Lua 浮点数值 370.5，用于检测二进制 chunk 的浮点数格式
	//        目前主流平台一般采用 IEEE 754 浮点数格式
	luaNum float64
}

// Prototype function prototype
// 标为 debug 的字段，属于调试信息，对于程序的执行是不必要的，如果使用 -s 选项编译，则会被去掉
type Prototype struct {
	// 源文件名，主函数的 Prototype 中才有值
	Source string // debug

	//  LineDefined 起行号， LastLineDefined 止行号
	LineDefined     uint32
	LastLineDefined uint32

	// 固定参数个数
	NumParams byte

	// 是否有变长参数
	IsVararg byte

	// 寄存器数量
	MaxStackSize byte

	// Lua 虚拟机 指令表
	Code []uint32

	// 常量表，用于存放字面量，不同常量有不同的 tag 标记
	// TAG_NIL, TAG_BOOLEAN, TAG_NUMBER, TAG_INTEGER, TAG_SHORT_STR, TAG_LONG_STR
	Constants []interface{}

	Upvalues []Upvalue

	// 子函数原型表
	Protos []*Prototype

	// 行号表，行号和 指令表 Code 中的指令一一对应，按 cint 类型存储
	LineInfo []uint32 // debug

	// 局部变量表
	LocVars []LocVar // debug

	// Upvalue 名列表，分别记录每个 Upvalue 在源码中的名字，与 Upvalues 中的 Upvalue 一一对应
	UpvalueNames []string // debug
}

type Upvalue struct {
	Instack byte
	Idx     byte
}

// LocVar 局部变量信息
type LocVar struct {
	VarName string // 局部变量名
	StartPC uint32 // 起索引
	EndPC   uint32 // 止索引
}

func Dump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte() // 跳过 Upvalue 数量
	return reader.readProto("")
}
