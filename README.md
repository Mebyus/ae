# Assembler emulator

## Syntax
```
// comment
label:
$imm -> imm
%reg -> R[reg]
imm -> M[imm]
(%reg) -> M[R[reg]]
imm(%reg1, %reg2, s) -> M[imm+R[reg1]+R[reg2]*s]
.section // data, text
[label:] mnemonic [operands] [// comment]
```

## Registers
```
AX
BX
CX
DX

SP
PC
```

## Flags
```
CF
ZF
SF
OF
```

## Instructions
```
mov S, D
push S
pop D

inc D
dec D
neg D
not D

add S, D
sub S, D
imul S, D
lea S, D
xor S, D
or S, D
and S, D

mul S, D
idiv S, D
div S, D

jmp L
jz L
jnz L
js L
jns L

cmp A, B
test A, B

call L
ret
```