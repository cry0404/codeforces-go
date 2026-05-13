把所有点保存到一个哈希集合中，每次循环，枚举集合中的所有点对，生成新的点。

如果没有生成新的点，那么继续循环也不会生成新的点，返回 $-1$。

如果 $\textit{target}$ 在集合中，返回循环次数。

[本题视频讲解](https://www.bilibili.com/video/BV1uB5L6YEeX/?t=4m54s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minGenerations(self, points: List[List[int]], target: List[int]) -> int:
        tar = tuple(target)
        cur = set(map(tuple, points))
        for ans in count(0):
            if tar in cur:
                return ans
            nxt = cur.copy()
            for (x, y, z), (a, b, c) in combinations(cur, 2):  # 枚举 cur 中的所有点对
                nxt.add(((x + a) // 2, (y + b) // 2, (z + c) // 2))
            if len(nxt) == len(cur):  # 没有产生新的点
                return -1
            cur = nxt
```

```java [sol-Java]
class Solution {
    public int minGenerations(int[][] points, int[] target) {
        int tar = hash(target[0], target[1], target[2]);

        Set<Integer> cur = new HashSet<>();
        for (int[] p : points) {
            cur.add(hash(p[0], p[1], p[2]));
        }

        for (int ans = 0; ; ans++) {
            if (cur.contains(tar)) {
                return ans;
            }

            Set<Integer> nxt = new HashSet<>(cur);
            for (int p : cur) {
                int px = p >> 6, py = p >> 3 & 7, pz = p & 7;
                for (int q : cur) { // 枚举 cur 中的所有点对 (p, q)
                    int qx = q >> 6, qy = q >> 3 & 7, qz = q & 7;
                    nxt.add(hash((px + qx) / 2, (py + qy) / 2, (pz + qz) / 2));
                }
            }

            if (nxt.size() == cur.size()) { // 没有产生新的点
                return -1;
            }

            cur = nxt;
        }
    }

    private int hash(int x, int y, int z) {
        return x << 6 | y << 3 | z; // 每个数占用 3 个比特位
    }
}
```

```cpp [sol-C++]
class Solution {
    int hash(int x, int y, int z) {
        return x << 6 | y << 3 | z; // 每个数占用 3 个比特位
    }

public:
    int minGenerations(vector<vector<int>>& points, vector<int>& target) {
        int tar = hash(target[0], target[1], target[2]);

        unordered_set<int> cur;
        for (auto& p : points) {
            cur.insert(hash(p[0], p[1], p[2]));
        }

        for (int ans = 0; ; ans++) {
            if (cur.contains(tar)) {
                return ans;
            }

            auto nxt = cur;
            for (int p : cur) {
                int px = p >> 6, py = p >> 3 & 7, pz = p & 7;
                for (int q : cur) { // 枚举 cur 中的所有点对 (p, q)
                    int qx = q >> 6, qy = q >> 3 & 7, qz = q & 7;
                    nxt.insert(hash((px + qx) / 2, (py + qy) / 2, (pz + qz) / 2));
                }
            }

            if (nxt.size() == cur.size()) { // 没有产生新的点
                return -1;
            }

            cur = move(nxt);
        }
    }
};
```

```go [sol-Go]
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
			for q := range cur { // 枚举 cur 中的所有点对 (p, q)
				nxt[point{(p.x + q.x) / 2, (p.y + q.y) / 2, (p.z + q.z) / 2}] = struct{}{}
			}
		}

		if len(nxt) == len(cur) { // 没有产生新的点
			return -1
		}

		cur = nxt
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(U^9)$，其中 $U\le 7$。至多有 $U^3\le 343$ 个点。在结束循环前，每次循环至少产生一个新的点，所以至多循环 $\mathcal{O}(U^3)$ 次，每次枚举点对花费 $\mathcal{O}((U^3)^2) = \mathcal{O}(U^6)$ 时间，所以总的时间复杂度为 $\mathcal{O}(U^9)$。
- 空间复杂度：$\mathcal{O}(U^3)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
