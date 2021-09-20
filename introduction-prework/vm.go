package vm

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
//
func compute(memory []byte) {

	registers := [3]byte{8, 0, 0} // PC, R1 and R2

	var pcJump byte = 3
	// Keep looping, like a physical computer's clock
	for {
		op := memory[registers[0]] // fetch the opcode
		arg1 := memory[registers[0]+1]
		arg2 := memory[registers[0]+2]
		pcJump = 3
		switch op {
		case Load:
			registers[arg1] = memory[arg2]
		case Store:
			memory[arg2] = registers[arg1]
		case Add:
			result := registers[arg1] + registers[arg2]
			if result > 255 {
				result = result % 255
			}
			registers[arg1] = result
		case Sub:
			registers[arg1] = registers[arg1] - registers[arg2]
		case Jump:
			registers[0] = arg1
			pcJump = 0
		case Beqz:
			if registers[arg1] == 0 {
				pcJump += arg2
			}
		case Addi:
			result := registers[arg1] + arg2
			if result > 255 {
				result = result % 255
			}
			registers[memory[registers[0]+1]] = result
		case Subi:
			result := registers[memory[registers[0]+1]] - memory[registers[0]+2]
			registers[memory[registers[0]+1]] = result
		case Halt:
			return
		}
		registers[0] += pcJump
	}
}
