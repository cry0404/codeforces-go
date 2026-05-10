$s$ 不含子序列 $\texttt{110}$ 或 $\texttt{011}$ ，意味着对于 $s$ 中的任意一个 $\texttt{0}$，其左侧不能有超过一个 $\texttt{1}$，右侧也不能有超过一个 $\texttt{1}$。

为了满足题目要求，要么让 $s$ 没有 $\texttt{0}$，要么让 $s$ 至多有一个 $\texttt{1}$（除了一种特殊情况）。

设 $n$ 是 $s$ 的长度，设 $c_0$ 是 $s$ 中的 $\texttt{0}$ 的个数，那么 $n-c_0$ 是 $s$ 中的 $\texttt{1}$ 的个数。

分类讨论：

- 把 $s$ 中的 $\texttt{0}$ 全变成 $\texttt{1}$，可以满足要求。操作 $c_0$ 次。
- 把 $s$ 中的 $\texttt{1}$ 变成 $\texttt{0}$，至多留一个 $\texttt{1}$ 不变，可以满足要求。操作 $\max(n-c_0-1, 0)$ 次。和 $0$ 取最大值是为了兼容 $s$ 没有 $\texttt{1}$ 的情况。特别地，如果 $s[0] = s[n-1] = \texttt{1}$，那么可以留下 $s[0]$ 和 $s[n-1]$ 不变，其余全变成 $\texttt{0}$，操作 $\max(n-c_0-2,0)$ 次。

两种情况取最小值，即为答案。

[本题视频讲解](https://www.bilibili.com/video/BV1uB5L6YEeX/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minFlips(self, s: str) -> int:
        c0 = s.count('0')
        c1 = len(s) - c0 - 1
        if s[0] == '1' and s[-1] == '1':
            c1 -= 1
        return min(c0, max(c1, 0))
```

```java [sol-Java]
class Solution {
    public int minFlips(String S) {
        char[] s = S.toCharArray();
        int n = s.length;

        int c0 = 0;
        for (char ch : s) {
            c0 += '1' - ch;
        }

        int c1 = n - c0 - 1;
        if (s[0] == '1' && s[n - 1] == '1') {
            c1--;
        }

        return Math.min(c0, Math.max(c1, 0));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minFlips(string s) {
        int n = s.size();
        int c0 = ranges::count(s, '0');
        int c1 = n - c0 - 1;
        if (s[0] == '1' && s[n - 1] == '1') {
            c1--;
        }
        return min(c0, max(c1, 0));
    }
};
```

```go [sol-Go]
func minFlips(s string) int {
	n := len(s)
	c0 := strings.Count(s, "0")
	c1 := n - c0 - 1
	if s[0] == '1' && s[n-1] == '1' {
		c1--
	}
	return min(c0, max(c1, 0))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面贪心与思维题单的「**§5.7 分类讨论**」。

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
