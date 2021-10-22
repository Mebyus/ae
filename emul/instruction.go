package emul

type OpType uint8

const (
	ConstValue OpType = iota
	ConstAddress
	RegisterValue
	RegisterAddress
	IndexedAddress
)

type RegisterOperand uint8

const (
	Sp RegisterOperand = iota

	Ax
	Bx
	Cx
	Dx
)

type Scale uint8

const (
	Scale1 = 1
	Scale2 = 1 << 1
	Scale4 = 1 << 2
	Scale8 = 1 << 3
)

type IndexOperand struct {
	Offset uint16
	Base   RegisterOperand
	Index  RegisterOperand
	Scale  Scale
}

type LiteralOperand struct {
	Neg bool
	Abs uint16
}

type Operand struct {
	Type     OpType
	Literal  LiteralOperand
	Register RegisterOperand
	Index    IndexOperand
}

type Instruction struct {
	Mnemo Mnemonic
	SrcOp Operand
	DstOp Operand
}
