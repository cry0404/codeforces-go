package main

// https://space.bilibili.com/206214
func scoreValidator(events []string) []int {
	score, counter := 0, 0

	for _, s := range events {
		if s == "W" {
			counter++
			if counter == 10 {
				break
			}
		} else if len(s) > 1 { // "WD" "NB"
			score++
		} else { // 数字
			score += int(s[0] - '0')
		}
	}

	return []int{score, counter}
}
