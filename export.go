package eaglesong

func Eaglesong(input []byte) (output []byte) {
	EaglesongSponge(output, uint(len(output)), input, uint(len(input)))
}
