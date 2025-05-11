package questions

func Swap(a, b *int) {
	*a, *b = *b, *a
}