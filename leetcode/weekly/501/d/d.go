package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
type edge struct{ to, wt int }

// 返回从起点 start 到每个点的最短路长度 dis
// 要求：没有负数边权
// 时间复杂度 O(n + mlogm)，注意堆中有 O(m) 个元素
func dijkstra(g [][]edge, start int, price int) []int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = math.MaxInt / 3 // 避免 p+dis1[j]+dis2[j] 加法溢出
	}
	dis[start] = 0
	h := hp{{0, start}}
	for len(h) > 0 {
		top := heap.Pop(&h).(pair)
		d, x := top.dis, top.x
		if d > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newD := d + e.wt
			if newD < price && newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}
	return dis
}

func minCost(n int, prices []int, roads [][]int) []int {
	g1 := make([][]edge, n)
	g2 := make([][]edge, n)
	for _, e := range roads {
		x, y, cost, tax := e[0], e[1], e[2], e[3]
		g1[x] = append(g1[x], edge{y, cost})
		g1[y] = append(g1[y], edge{x, cost})
		g2[x] = append(g2[x], edge{y, cost * tax})
		g2[y] = append(g2[y], edge{x, cost * tax})
	}

	ans := make([]int, n)
	for i, price := range prices {
		dis1 := dijkstra(g1, i, price)
		dis2 := dijkstra(g2, i, price)
		res := math.MaxInt
		for j, p := range prices {
			res = min(res, p+dis1[j]+dis2[j])
		}
		ans[i] = res
	}
	return ans
}

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
