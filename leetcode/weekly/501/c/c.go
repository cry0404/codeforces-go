package main

// https://space.bilibili.com/206214
const mx = 100_001
var divisors [mx][]int

func init() {
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}
}

func minArraySum(nums []int) (ans int64) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	for x, c := range cnt { // 遍历 cnt 而不是 nums，这样重复元素只会计算一次
		for _, d := range divisors[x] { // 从小到大枚举 x 的因子 d
			if cnt[d] > 0 {
				ans += int64(d) * int64(c) // 把 x 变成 d 是最优的
				break
			}
		}
	}
	return
}
