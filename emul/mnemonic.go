package emul

type Mnemonic uint16

const (
	Mov Mnemonic = iota
	Push
	Pop
	Inc
	Dec
	Neg
	Not
	Add
	Sub
	Mul
	Imul
	Lea
	Xor
	Or
	And
	Div
	Idiv
	Jmp
	Jz
	Jnz
	Js
	Jns
	Cmp
	Test
	Call
	Ret
)
