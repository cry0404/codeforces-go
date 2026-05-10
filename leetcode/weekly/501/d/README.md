本题 $n\le 1000$，可以暴力一点。

对于每个 $i$，枚举在商店 $j=0,1,2,\ldots,n-1$ 购买苹果。我们需要知道：

- 从 $i$ 到 $j$ 的边权为 $\textit{cost}$ 的最短路长度。
- 从 $j$ 到 $i$ 的边权为 $\textit{cost}\cdot \textit{tax}$ 的最短路长度。由于图是无向图，从 $j$ 到 $i$ 的最短路长度等于从 $i$ 到 $j$ 的最短路长度。

用 Dijkstra 算法求最短路长度。

⚠**剪枝**：由于答案至多为 $\textit{prices}[i]$，所以若算出的最短路长度 $\ge \textit{prices}[i]$，则不入堆。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 返回从起点 start 到每个点的最短路长度 dis，如果节点 x 不可达，则 dis[x] = inf
# 要求：没有负数边权
# 时间复杂度 O(n + mlogm)，其中 n 是 g 的长度，m 是 edges 的长度。注意堆中有 O(m) 个元素
def shortestPathDijkstra(g: list[list[tuple[int, int]]], start: int, price: int) -> list[int]:
    dis = [inf] * len(g)
    dis[start] = 0  # 起点到自己的距离是 0
    h = [(0, start)]  # 堆中保存 (起点到节点 x 的最短路长度，节点 x)

    while h:
        dis_x, x = heappop(h)
        if dis_x > dis[x]:  # x 之前出堆过
            continue
        for y, wt in g[x]:
            new_dis_y = dis_x + wt
            if new_dis_y < price and new_dis_y < dis[y]:  # 剪枝
                dis[y] = new_dis_y  # 更新 x 的邻居的最短路
                # 懒更新堆：只插入数据，不更新堆中数据
                # 相同节点可能有多个不同的 new_dis_y，除了最小的 new_dis_y，其余值都会触发上面的 continue
                heappush(h, (new_dis_y, y))

    return dis

class Solution:
    def minCost(self, n: int, prices: list[int], roads: list[list[int]]) -> list[int]:
        g1 = [[] for _ in range(n)]
        g2 = [[] for _ in range(n)]
        for x, y, cost, tax in roads:
            g1[x].append((y, cost))
            g1[y].append((x, cost))
            g2[x].append((y, cost * tax))
            g2[y].append((x, cost * tax))

        ans = [0] * n
        for i, price in enumerate(prices):
            dis1 = shortestPathDijkstra(g1, i, price)
            dis2 = shortestPathDijkstra(g2, i, price)
            ans[i] = min(p + d1 + d2 for p, d1, d2 in zip(prices, dis1, dis2))
        return ans
```

```java [sol-Java]
class Solution {
    private record Pair(int x, long y) {
    }

    public int[] minCost(int n, int[] prices, int[][] roads) {
        List<Pair>[] g1 = new ArrayList[n];
        List<Pair>[] g2 = new ArrayList[n];
        Arrays.setAll(g1, _ -> new ArrayList<>());
        Arrays.setAll(g2, _ -> new ArrayList<>());
        for (int[] e : roads) {
            int x = e[0], y = e[1], cost = e[2], tax = e[3];
            g1[x].add(new Pair(y, cost));
            g1[y].add(new Pair(x, cost));
            g2[x].add(new Pair(y, (long) cost * tax));
            g2[y].add(new Pair(x, (long) cost * tax));
        }

        int[] ans = new int[n];
        for (int i = 0; i < n; i++) {
            int price = prices[i];
            long[] dis1 = shortestPathDijkstra(g1, i, price);
            long[] dis2 = shortestPathDijkstra(g2, i, price);
            long res = Long.MAX_VALUE;
            for (int j = 0; j < n; j++) {
                res = Math.min(res, prices[j] + dis1[j] + dis2[j]);
            }
            ans[i] = (int) res;
        }

        return ans;
    }

    // 返回从起点 start 到每个点的最短路长度 dis
    // 要求：没有负数边权
    // 时间复杂度 O(n + mlogm)，注意堆中有 O(m) 个元素
    private long[] shortestPathDijkstra(List<Pair>[] g, int start, int price) {
        long[] dis = new long[g.length];
        Arrays.fill(dis, Long.MAX_VALUE / 3); // 避免 prices[j] + dis1[j] + dis2[j] 溢出
        // 堆中保存 (节点 x, 起点到节点 x 的最短路长度)
        PriorityQueue<Pair> pq = new PriorityQueue<>(Comparator.comparingLong(a -> a.y));
        dis[start] = 0; // 起点到自己的距离是 0
        pq.offer(new Pair(start, 0));

        while (!pq.isEmpty()) {
            Pair p = pq.poll();
            int x = p.x;
            long disX = p.y;
            if (disX > dis[x]) { // x 之前出堆过
                continue;
            }
            for (Pair e : g[x]) {
                int y = e.x;
                long newDisY = disX + e.y;
                if (newDisY < price && newDisY < dis[y]) {
                    dis[y] = newDisY; // 更新 x 的邻居的最短路
                    // 懒更新堆：只插入数据，不更新堆中数据
                    // 相同节点可能有多个不同的 newDisY，除了最小的 newDisY，其余值都会触发上面的 continue
                    pq.offer(new Pair(y, newDisY));
                }
            }
        }

        return dis;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 返回从起点 start 到每个点的最短路长度 dis
    // 要求：没有负数边权
    // 时间复杂度 O(n + mlogm)，注意堆中有 O(m) 个元素
    vector<long long> shortestPathDijkstra(vector<vector<pair<int, long long>>>& g, int start, int price) {
        vector<long long> dis(g.size(), LLONG_MAX / 3); // 避免 prices[j] + dis1[j] + dis2[j] 溢出
        // 堆中保存 (起点到节点 x 的最短路长度，节点 x)
        priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq;
        dis[start] = 0; // 起点到自己的距离是 0
        pq.emplace(0, start);

        while (!pq.empty()) {
            auto [dis_x, x] = pq.top();
            pq.pop();
            if (dis_x > dis[x]) { // x 之前出堆过
                continue;
            }
            for (auto& [y, wt] : g[x]) {
                auto new_dis_y = dis_x + wt;
                if (new_dis_y < price && new_dis_y < dis[y]) {
                    dis[y] = new_dis_y; // 更新 x 的邻居的最短路
                    // 懒更新堆：只插入数据，不更新堆中数据
                    // 相同节点可能有多个不同的 new_dis_y，除了最小的 new_dis_y，其余值都会触发上面的 continue
                    pq.emplace(new_dis_y, y);
                }
            }
        }

        return dis;
    }

public:
    vector<int> minCost(int n, vector<int>& prices, vector<vector<int>>& roads) {
        vector<vector<pair<int, long long>>> g1(n);
        vector<vector<pair<int, long long>>> g2(n);
        for (auto& e : roads) {
            int x = e[0], y = e[1], cost = e[2], tax = e[3];
            g1[x].emplace_back(y, cost);
            g1[y].emplace_back(x, cost);
            g2[x].emplace_back(y, 1LL * cost * tax);
            g2[y].emplace_back(x, 1LL * cost * tax);
        }

        vector<int> ans(n);
        for (int i = 0; i < n; i++) {
            int price = prices[i];
            vector<long long> dis1 = shortestPathDijkstra(g1, i, price);
            vector<long long> dis2 = shortestPathDijkstra(g2, i, price);
            long long res = LLONG_MAX;
            for (int j = 0; j < n; j++) {
                res = min(res, prices[j] + dis1[j] + dis2[j]);
            }
            ans[i] = res;
        }
        return ans;
    }
};
```

```go [sol-Go]
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
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+m\log m))$，其中 $m$ 是 $\textit{roads}$ 的长度。在使用**懒更新堆**的情况下，跑一次 Dijkstra 算法的时间复杂度为 $\mathcal{O}(n + m\log m)$。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 专题训练

见下面图论题单的「**§3.1 单源最短路：Dijkstra 算法**」。

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
