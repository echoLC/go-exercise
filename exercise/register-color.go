package exercise

func RegisterColor(colors []string) int {
	colorMap := map[string]int{"black": 0, "brown": 1, "red": 2, "orange": 3, "yellow": 4, "green": 5, "blue": 6, "violet": 7, "grey": 8, "white": 9}

	res := 0

	for i := 1; i >= 0; i-- {
		value := colors[i]
		colorValue, existed := colorMap[value]

		if (existed) {
			if i == 1 {
				res += colorValue
			} else {
				res += 10 * colorValue
			}
		} 
	}

	return res 
}