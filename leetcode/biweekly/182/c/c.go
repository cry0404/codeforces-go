package main

import "maps"

// https://space.bilibili.com/206214
func minGenerations(points [][]int, target []int) int {
	type point struct{ x, y, z int }
	tar := point{target[0], target[1], target[2]}

	cur := make(map[point]struct{}, len(points))
	for _, p := range points {
		cur[point{p[0], p[1], p[2]}] = struct{}{}
	}

	for ans := 0; ; ans++ {
		if _, ok := cur[tar]; ok {
			return ans
		}

		nxt := maps.Clone(cur)
		for p := range cur {
			for q := range cur {
				nxt[point{(p.x + q.x) / 2, (p.y + q.y) / 2, (p.z + q.z) / 2}] = struct{}{}
			}
		}

		if len(nxt) == len(cur) {
			return -1
		}

		cur = nxt
	}
}
