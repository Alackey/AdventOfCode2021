package day2

type line struct {
	direction string
	move      int
}

// Part1 returns the horizontal * depth position of the submarine
func Part1(lines []*line) (int, error) {
	var horizontal, depth int
	for _, line := range lines {
		switch line.direction {
		case "forward":
			horizontal += line.move
		case "down":
			depth += line.move
		case "up":
			depth -= line.move
		}
	}
	return horizontal * depth, nil
}

// Part1 returns the horizontal * depth position of the submarine using aim
func Part2(lines []*line) (int, error) {
	var horizontal, aim, depth int
	for _, line := range lines {
		switch line.direction {
		case "forward":
			horizontal += line.move
			depth += line.move * aim
		case "down":
			aim += line.move
		case "up":
			aim -= line.move
		}
	}
	return horizontal * depth, nil
}
