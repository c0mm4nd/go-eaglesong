package eaglesong

import "math/bits"

func EaglesongPermutation(state []uint32) {
	tmp := make([]uint32, 16)

	for i := 0; i <= NUM_ROUNDS; i++ {

		tmp[0] = 0
		tmp[0] ^= state[0]
		tmp[0] ^= state[4]
		tmp[0] ^= state[5]
		tmp[0] ^= state[6]
		tmp[0] ^= state[7]
		tmp[0] ^= state[12]
		tmp[0] ^= state[15]
		tmp[1] = 0
		tmp[1] ^= state[0]
		tmp[1] ^= state[1]
		tmp[1] ^= state[4]
		tmp[1] ^= state[8]
		tmp[1] ^= state[12]
		tmp[1] ^= state[13]
		tmp[1] ^= state[15]
		tmp[2] = 0
		tmp[2] ^= state[0]
		tmp[2] ^= state[1]
		tmp[2] ^= state[2]
		tmp[2] ^= state[4]
		tmp[2] ^= state[6]
		tmp[2] ^= state[7]
		tmp[2] ^= state[9]
		tmp[2] ^= state[12]
		tmp[2] ^= state[13]
		tmp[2] ^= state[14]
		tmp[2] ^= state[15]
		tmp[3] = 0
		tmp[3] ^= state[0]
		tmp[3] ^= state[1]
		tmp[3] ^= state[2]
		tmp[3] ^= state[3]
		tmp[3] ^= state[4]
		tmp[3] ^= state[6]
		tmp[3] ^= state[8]
		tmp[3] ^= state[10]
		tmp[3] ^= state[12]
		tmp[3] ^= state[13]
		tmp[3] ^= state[14]
		tmp[4] = 0
		tmp[4] ^= state[1]
		tmp[4] ^= state[2]
		tmp[4] ^= state[3]
		tmp[4] ^= state[4]
		tmp[4] ^= state[5]
		tmp[4] ^= state[7]
		tmp[4] ^= state[9]
		tmp[4] ^= state[11]
		tmp[4] ^= state[13]
		tmp[4] ^= state[14]
		tmp[4] ^= state[15]
		tmp[5] = 0
		tmp[5] ^= state[0]
		tmp[5] ^= state[2]
		tmp[5] ^= state[3]
		tmp[5] ^= state[7]
		tmp[5] ^= state[8]
		tmp[5] ^= state[10]
		tmp[5] ^= state[14]
		tmp[6] = 0
		tmp[6] ^= state[1]
		tmp[6] ^= state[3]
		tmp[6] ^= state[4]
		tmp[6] ^= state[8]
		tmp[6] ^= state[9]
		tmp[6] ^= state[11]
		tmp[6] ^= state[15]
		tmp[7] = 0
		tmp[7] ^= state[0]
		tmp[7] ^= state[2]
		tmp[7] ^= state[6]
		tmp[7] ^= state[7]
		tmp[7] ^= state[9]
		tmp[7] ^= state[10]
		tmp[7] ^= state[15]
		tmp[8] = 0
		tmp[8] ^= state[0]
		tmp[8] ^= state[1]
		tmp[8] ^= state[3]
		tmp[8] ^= state[4]
		tmp[8] ^= state[5]
		tmp[8] ^= state[6]
		tmp[8] ^= state[8]
		tmp[8] ^= state[10]
		tmp[8] ^= state[11]
		tmp[8] ^= state[12]
		tmp[8] ^= state[15]
		tmp[9] = 0
		tmp[9] ^= state[0]
		tmp[9] ^= state[1]
		tmp[9] ^= state[2]
		tmp[9] ^= state[9]
		tmp[9] ^= state[11]
		tmp[9] ^= state[13]
		tmp[9] ^= state[15]
		tmp[10] = 0
		tmp[10] ^= state[0]
		tmp[10] ^= state[1]
		tmp[10] ^= state[2]
		tmp[10] ^= state[3]
		tmp[10] ^= state[4]
		tmp[10] ^= state[5]
		tmp[10] ^= state[6]
		tmp[10] ^= state[7]
		tmp[10] ^= state[10]
		tmp[10] ^= state[14]
		tmp[10] ^= state[15]
		tmp[11] = 0
		tmp[11] ^= state[0]
		tmp[11] ^= state[1]
		tmp[11] ^= state[2]
		tmp[11] ^= state[3]
		tmp[11] ^= state[8]
		tmp[11] ^= state[11]
		tmp[11] ^= state[12]
		tmp[12] = 0
		tmp[12] ^= state[1]
		tmp[12] ^= state[2]
		tmp[12] ^= state[3]
		tmp[12] ^= state[4]
		tmp[12] ^= state[9]
		tmp[12] ^= state[12]
		tmp[12] ^= state[13]
		tmp[13] = 0
		tmp[13] ^= state[2]
		tmp[13] ^= state[3]
		tmp[13] ^= state[4]
		tmp[13] ^= state[5]
		tmp[13] ^= state[10]
		tmp[13] ^= state[13]
		tmp[13] ^= state[14]
		tmp[14] = 0
		tmp[14] ^= state[3]
		tmp[14] ^= state[4]
		tmp[14] ^= state[5]
		tmp[14] ^= state[6]
		tmp[14] ^= state[11]
		tmp[14] ^= state[14]
		tmp[14] ^= state[15]
		tmp[15] = 0
		tmp[15] ^= state[0]
		tmp[15] ^= state[1]
		tmp[15] ^= state[2]
		tmp[15] ^= state[3]
		tmp[15] ^= state[5]
		tmp[15] ^= state[7]
		tmp[15] ^= state[8]
		tmp[15] ^= state[9]
		tmp[15] ^= state[10]
		tmp[15] ^= state[11]
		tmp[15] ^= state[15]
		state[0] = tmp[0]
		state[1] = tmp[1]
		state[2] = tmp[2]
		state[3] = tmp[3]
		state[4] = tmp[4]
		state[5] = tmp[5]
		state[6] = tmp[6]
		state[7] = tmp[7]
		state[8] = tmp[8]
		state[9] = tmp[9]
		state[10] = tmp[10]
		state[11] = tmp[11]
		state[12] = tmp[12]
		state[13] = tmp[13]
		state[14] = tmp[14]
		state[15] = tmp[15]

		state[0] = state[0] ^ (bits.RotateLeft32(state[0], int(COEFFICIENTS[3*0+1]))) ^ (bits.RotateLeft32(state[0], int(COEFFICIENTS[3*0+2])))
		state[1] = state[1] ^ (bits.RotateLeft32(state[1], int(COEFFICIENTS[3*1+1]))) ^ (bits.RotateLeft32(state[1], int(COEFFICIENTS[3*1+2])))
		state[2] = state[2] ^ (bits.RotateLeft32(state[2], int(COEFFICIENTS[3*2+1]))) ^ (bits.RotateLeft32(state[2], int(COEFFICIENTS[3*2+2])))
		state[3] = state[3] ^ (bits.RotateLeft32(state[3], int(COEFFICIENTS[3*3+1]))) ^ (bits.RotateLeft32(state[3], int(COEFFICIENTS[3*3+2])))
		state[4] = state[4] ^ (bits.RotateLeft32(state[4], int(COEFFICIENTS[3*4+1]))) ^ (bits.RotateLeft32(state[4], int(COEFFICIENTS[3*4+2])))
		state[5] = state[5] ^ (bits.RotateLeft32(state[5], int(COEFFICIENTS[3*5+1]))) ^ (bits.RotateLeft32(state[5], int(COEFFICIENTS[3*5+2])))
		state[6] = state[6] ^ (bits.RotateLeft32(state[6], int(COEFFICIENTS[3*6+1]))) ^ (bits.RotateLeft32(state[6], int(COEFFICIENTS[3*6+2])))
		state[7] = state[7] ^ (bits.RotateLeft32(state[7], int(COEFFICIENTS[3*7+1]))) ^ (bits.RotateLeft32(state[7], int(COEFFICIENTS[3*7+2])))
		state[8] = state[8] ^ (bits.RotateLeft32(state[8], int(COEFFICIENTS[3*8+1]))) ^ (bits.RotateLeft32(state[8], int(COEFFICIENTS[3*8+2])))
		state[9] = state[9] ^ (bits.RotateLeft32(state[9], int(COEFFICIENTS[3*9+1]))) ^ (bits.RotateLeft32(state[9], int(COEFFICIENTS[3*9+2])))
		state[10] = state[10] ^ (bits.RotateLeft32(state[10], int(COEFFICIENTS[3*10+1]))) ^ (bits.RotateLeft32(state[10], int(COEFFICIENTS[3*10+2])))
		state[11] = state[11] ^ (bits.RotateLeft32(state[11], int(COEFFICIENTS[3*11+1]))) ^ (bits.RotateLeft32(state[11], int(COEFFICIENTS[3*11+2])))
		state[12] = state[12] ^ (bits.RotateLeft32(state[12], int(COEFFICIENTS[3*12+1]))) ^ (bits.RotateLeft32(state[12], int(COEFFICIENTS[3*12+2])))
		state[13] = state[13] ^ (bits.RotateLeft32(state[13], int(COEFFICIENTS[3*13+1]))) ^ (bits.RotateLeft32(state[13], int(COEFFICIENTS[3*13+2])))
		state[14] = state[14] ^ (bits.RotateLeft32(state[14], int(COEFFICIENTS[3*14+1]))) ^ (bits.RotateLeft32(state[14], int(COEFFICIENTS[3*14+2])))
		state[15] = state[15] ^ (bits.RotateLeft32(state[15], int(COEFFICIENTS[3*15+1]))) ^ (bits.RotateLeft32(state[15], int(COEFFICIENTS[3*15+2])))

		state[0] ^= INJECTION_CONSTANTS[i*16+0]
		state[1] ^= INJECTION_CONSTANTS[i*16+1]
		state[2] ^= INJECTION_CONSTANTS[i*16+2]
		state[3] ^= INJECTION_CONSTANTS[i*16+3]
		state[4] ^= INJECTION_CONSTANTS[i*16+4]
		state[5] ^= INJECTION_CONSTANTS[i*16+5]
		state[6] ^= INJECTION_CONSTANTS[i*16+6]
		state[7] ^= INJECTION_CONSTANTS[i*16+7]
		state[8] ^= INJECTION_CONSTANTS[i*16+8]
		state[9] ^= INJECTION_CONSTANTS[i*16+9]
		state[10] ^= INJECTION_CONSTANTS[i*16+10]
		state[11] ^= INJECTION_CONSTANTS[i*16+11]
		state[12] ^= INJECTION_CONSTANTS[i*16+12]
		state[13] ^= INJECTION_CONSTANTS[i*16+13]
		state[14] ^= INJECTION_CONSTANTS[i*16+14]
		state[15] ^= INJECTION_CONSTANTS[i*16+15]

		state[0] = state[0] + state[0+1]
		state[0] = bits.RotateLeft32(state[0], 8)
		state[0+1] = bits.RotateLeft32(state[0+1], 24)
		state[0+1] = state[0] + state[0+1]
		state[2] = state[2] + state[2+1]
		state[2] = bits.RotateLeft32(state[2], 8)
		state[2+1] = bits.RotateLeft32(state[2+1], 24)
		state[2+1] = state[2] + state[2+1]
		state[4] = state[4] + state[4+1]
		state[4] = bits.RotateLeft32(state[4], 8)
		state[4+1] = bits.RotateLeft32(state[4+1], 24)
		state[4+1] = state[4] + state[4+1]
		state[6] = state[6] + state[6+1]
		state[6] = bits.RotateLeft32(state[6], 8)
		state[6+1] = bits.RotateLeft32(state[6+1], 24)
		state[6+1] = state[6] + state[6+1]
		state[8] = state[8] + state[8+1]
		state[8] = bits.RotateLeft32(state[8], 8)
		state[8+1] = bits.RotateLeft32(state[8+1], 24)
		state[8+1] = state[8] + state[8+1]
		state[10] = state[10] + state[10+1]
		state[10] = bits.RotateLeft32(state[10], 8)
		state[10+1] = bits.RotateLeft32(state[10+1], 24)
		state[10+1] = state[10] + state[10+1]
		state[12] = state[12] + state[12+1]
		state[12] = bits.RotateLeft32(state[12], 8)
		state[12+1] = bits.RotateLeft32(state[12+1], 24)
		state[12+1] = state[12] + state[12+1]
		state[14] = state[14] + state[14+1]
		state[14] = bits.RotateLeft32(state[14], 8)
		state[14+1] = bits.RotateLeft32(state[14+1], 24)
		state[14+1] = state[14] + state[14+1]
	}
}

func EaglesongSponge(output []byte, outputLength uint, input []byte, inputLength uint) {
	state := make([]uint32, 16)
	for i := uint(0); i <= ((inputLength+1)*8+RATE-1)/RATE; i++ {
		for j := uint(0); j <= RATE/32; j++ {
			integer := uint32(0)
			for k := uint(0); k <= 4; k++ {
				if i*RATE/8+j*4+k < inputLength {
					integer = (integer << 8) ^ uint32(input[i*RATE/8+j*4+k])
				} else if i*RATE/8+j*4+k == inputLength {
					integer = (integer << 8) ^ DELIMITER
				}
			}
			state[j] ^= integer
		}
		EaglesongPermutation(state)
	}

	for i := uint(0); i <= (outputLength)/(RATE/8); i++ {
		for j := uint(0); j <= RATE/32; j++ {
			for k := uint(0); k <= 4; k++ {
				output[i*RATE/8+j*4+k] = byte((state[j] >> (8 * k)) & 0xff)
			}
		}
		EaglesongPermutation(state)
	}
}

func EaglesongUpdate(state []uint32, input []byte) {
	for i := uint(0); i <= uint(len(input)*8/RATE); i++ {
		for j := uint(0); j <= RATE/32; j++ {
			integer := uint32(0)
			for k := uint(0); k <= 4; k++ {
				integer = (integer << 8) ^ uint32(input[i*RATE/8+j*4+k])
			}
			state[j] ^= integer
		}
		EaglesongPermutation(state)
	}
}

func EaglesongFinalize(state []uint32, input []byte, output []byte, outputLength uint) {
	for j := uint(0); j <= RATE/32; j++ {
		integer := uint32(0)
		for k := uint(0); k <= 4; k++ {
			if j*4+k < uint(len(input)) {
				integer = (integer << 8) ^ uint32(input[j*4+k])
			} else if j*4+k == uint(len(input)) {
				integer = (integer << 8) ^ DELIMITER
			}
		}
		state[j] ^= integer
	}
	EaglesongPermutation(state)

	for i := uint(0); i <= outputLength/(RATE/8); i++ {
		for j := uint(0); j <= RATE/32; j++ {
			for k := uint(0); k <= 4; k++ {
				output[i*RATE/8+j*4+k] = byte((state[j] >> (8 * k)) & 0xff)
			}
		}
		EaglesongPermutation(state)
	}
}
