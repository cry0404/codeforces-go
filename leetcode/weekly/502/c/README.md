[二维 ST 表原理讲解](https://blog.nowcoder.net/n/3eccd1386a8846398d3bee62b485309b)

预处理**二维 ST 表**，即可 $\mathcal{O}(1)$ 算出任意子矩阵的最大值。

需要注意的是，本题不能计入子矩阵的四个角，怎么办？

![w502c.png](https://pic.leetcode.cn/1778988029-AamMpU-w502c.png){:width=300px}

我们可以把蓝色区域视作**两个子矩阵的并集**。以上图（示例 1）为例说明：

- 计算以 $2$ 为中心的 $3$ 行 $5$ 列的子矩阵的最大值。
- 计算以 $2$ 为中心的 $5$ 行 $3$ 列的子矩阵的最大值。
- 这两个最大值的最大值，即为蓝色区域的最大值。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def countLocalMaximums(self, matrix: list[list[int]]) -> int:
        n, m = len(matrix), len(matrix[0])
        wn, wm = n.bit_length(), m.bit_length()

        # st[k1][k2][i][j] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
        st = [[[[0] * m for _ in range(n)] for _ in range(wm)] for _ in range(wn)]

        # 初始值
        st[0][0] = matrix

        # 单独计算 k1 = 0
        for k2 in range(1, wm):
            half = 1 << (k2 - 1)
            for i in range(n):
                for j in range(m - (1 << k2) + 1):
                    st[0][k2][i][j] = max(st[0][k2 - 1][i][j], st[0][k2 - 1][i][j + half])

        for k1 in range(1, wn):
            half = 1 << (k1 - 1)
            for k2 in range(wm):
                for i in range(n - (1 << k1) + 1):
                    for j in range(m - (1 << k2) + 1):
                        st[k1][k2][i][j] = max(st[k1 - 1][k2][i][j], st[k1 - 1][k2][i + half][j])

        # 返回子矩阵最大值
        # 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
        def query(r1: int, c1: int, r2: int, c2: int) -> int:
            r1 = max(r1, 0)
            c1 = max(c1, 0)
            r2 = min(r2, n)
            c2 = min(c2, m)
            k1 = (r2 - r1).bit_length() - 1
            k2 = (c2 - c1).bit_length() - 1
            # 视作四个子矩阵的并集
            return max(st[k1][k2][r1][c1],
                       st[k1][k2][r2 - (1 << k1)][c1],
                       st[k1][k2][r1][c2 - (1 << k2)],
                       st[k1][k2][r2 - (1 << k1)][c2 - (1 << k2)])

        ans = 0
        for i, row in enumerate(matrix):
            for j, x in enumerate(row):
                if x > 0 and max(query(i - x, j - x + 1, i + x + 1, j + x), query(i - x + 1, j - x, i + x, j + x + 1)) <= x:
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countLocalMaximums(int[][] matrix) {
        int n = matrix.length;
        int m = matrix[0].length;
        int wn = 32 - Integer.numberOfLeadingZeros(n);
        int wm = 32 - Integer.numberOfLeadingZeros(m);

        // st[k1][k2][i][j] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
        int[][][][] st = new int[wn][wm][n][m];

        // 初始值
        st[0][0] = matrix;

        // 单独计算 k1 = 0
        for (int k2 = 1; k2 < wm; k2++) {
            int half = 1 << (k2 - 1);
            for (int i = 0; i < n; i++) {
                for (int j = 0; j <= m - (1 << k2); j++) {
                    st[0][k2][i][j] = Math.max(st[0][k2 - 1][i][j], st[0][k2 - 1][i][j + half]);
                }
            }
        }

        for (int k1 = 1; k1 < wn; k1++) {
            int half = 1 << (k1 - 1);
            for (int k2 = 0; k2 < wm; k2++) {
                for (int i = 0; i <= n - (1 << k1); i++) {
                    for (int j = 0; j <= m - (1 << k2); j++) {
                        st[k1][k2][i][j] = Math.max(st[k1 - 1][k2][i][j], st[k1 - 1][k2][i + half][j]);
                    }
                }
            }
        }

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                int x = matrix[i][j];
                if (x == 0) {
                    continue;
                }
                int max1 = query(st, i - x, j - x + 1, i + x + 1, j + x, n, m);
                int max2 = query(st, i - x + 1, j - x, i + x, j + x + 1, n, m);
                if (Math.max(max1, max2) <= x) {
                    ans++;
                }
            }
        }
        return ans;
    }

    // 返回子矩阵最大值
    // 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
    private int query(int[][][][] st, int r1, int c1, int r2, int c2, int n, int m) {
        r1 = Math.max(r1, 0);
        c1 = Math.max(c1, 0);
        r2 = Math.min(r2, n);
        c2 = Math.min(c2, m);
        int k1 = 31 - Integer.numberOfLeadingZeros(r2 - r1);
        int k2 = 31 - Integer.numberOfLeadingZeros(c2 - c1);
        // 视作四个子矩阵的并集
        return Math.max(
                Math.max(st[k1][k2][r1][c1], st[k1][k2][r2 - (1 << k1)][c1]),
                Math.max(st[k1][k2][r1][c2 - (1 << k2)], st[k1][k2][r2 - (1 << k1)][c2 - (1 << k2)])
        );
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countLocalMaximums(vector<vector<int>>& matrix) {
        int n = matrix.size(), m = matrix[0].size();
        int wn = bit_width(1u * n), wm = bit_width(1u * m);

        // st[i][j][k1][k2] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
        vector st(n, vector<array<array<int, 8>, 8>>(m));

        // 初始值
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                st[i][j][0][0] = matrix[i][j];
            }
        }

        // 单独计算 k1 = 0
        for (int k2 = 1; k2 < wm; k2++) {
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < m - (1 << k2) + 1; j++) {
                    st[i][j][0][k2] = max(st[i][j][0][k2 - 1], st[i][j + (1 << (k2 - 1))][0][k2 - 1]);
                }
            }
        }

        for (int k1 = 1; k1 < wn; k1++) {
            for (int k2 = 0; k2 < wm; k2++) {
                for (int i = 0; i < n - (1 << k1) + 1; i++) {
                    for (int j = 0; j < m - (1 << k2) + 1; j++) {
                        st[i][j][k1][k2] = max(st[i][j][k1 - 1][k2], st[i + (1 << (k1 - 1))][j][k1 - 1][k2]);
                    }
                }
            }
        }

        // 返回子矩阵最大值
        // 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
        auto query = [&](int r1, int c1, int r2, int c2) -> int {
            r1 = max(r1, 0);
            c1 = max(c1, 0);
            r2 = min(r2, n);
            c2 = min(c2, m);
            int k1 = bit_width(1u * (r2 - r1)) - 1;
            int k2 = bit_width(1u * (c2 - c1)) - 1;
            // 视作四个子矩阵的并集
            return max({st[r1][c1][k1][k2], st[r2 - (1 << k1)][c1][k1][k2], st[r1][c2 - (1 << k2)][k1][k2], st[r2 - (1 << k1)][c2 - (1 << k2)][k1][k2]});
        };

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                int x = matrix[i][j];
                if (x > 0 && max(query(i - x, j - x + 1, i + x + 1, j + x), query(i - x + 1, j - x, i + x, j + x + 1)) <= x) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countLocalMaximums(matrix [][]int) (ans int) {
	n, m := len(matrix), len(matrix[0])
	wn, wm := bits.Len(uint(n)), bits.Len(uint(m))
	const MX = 8
	// st[i][j][k1][k2] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
	st := make([][][MX][MX]int, n)
	for i := range st {
		st[i] = make([][MX][MX]int, m)
	}
	for i, row := range matrix {
		for j, x := range row {
			st[i][j][0][0] = x // 初始值
		}
	}
	// 单独计算 k1 = 0
	for k2 := 1; k2 < wm; k2++ {
		for i := range n {
			for j := range m - 1<<k2 + 1 {
				st[i][j][0][k2] = max(st[i][j][0][k2-1], st[i][j+1<<(k2-1)][0][k2-1])
			}
		}
	}
	for k1 := 1; k1 < wn; k1++ {
		for k2 := range wm {
			for i := range n - 1<<k1 + 1 {
				for j := range m - 1<<k2 + 1 {
					st[i][j][k1][k2] = max(st[i][j][k1-1][k2], st[i+1<<(k1-1)][j][k1-1][k2])
				}
			}
		}
	}

	// 返回子矩阵最大值
	// 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
	query := func(r1, c1, r2, c2 int) int {
		r1 = max(r1, 0)
		c1 = max(c1, 0)
		r2 = min(r2, n)
		c2 = min(c2, m)
		k1 := bits.Len8(uint8(r2-r1)) - 1
		k2 := bits.Len8(uint8(c2-c1)) - 1
		// 视作四个子矩阵的并集
		return max(st[r1][c1][k1][k2], st[r2-1<<k1][c1][k1][k2], st[r1][c2-1<<k2][k1][k2], st[r2-1<<k1][c2-1<<k2][k1][k2])
	}

	for i, row := range matrix {
		for j, x := range row {
			if x > 0 && max(query(i-x, j-x+1, i+x+1, j+x), query(i-x+1, j-x, i+x, j+x+1)) <= x {
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm\log n\log m)$，其中 $n$ 和 $m$ 分别是 $\textit{matrix}$ 的行数和列数。瓶颈在预处理 ST 表上。
- 空间复杂度：$\mathcal{O}(nm\log n\log m)$。

## 专题训练

见下面数据结构题单的「**§8.7 ST 表**」。

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
