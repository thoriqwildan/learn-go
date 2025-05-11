package questions

func AddBonus(list []int, add *int) {
	for i, _ := range list {
		list[i] += *add
	}
}