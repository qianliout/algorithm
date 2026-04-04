toolName: view_files
            
status: success
          
            
filePath: /Users/liuqianli/work/golang/src/algorithm/202601/40combinationSum2/combinationSum2.go
          

        
            
toolName: Skill
            
status: success
          
          
**核心思路（这段剪枝在做什么）**
- 背景：输入已排序，DFS 采用“对于下标 i：先不选 i，再考虑选 i”的二叉展开；used 记录当前路径里“相邻相同值的上一个元素是否被用过”。
- 条件含义：[combinationSum2.go:L33-L35](file:///Users/liuqianli/work/golang/src/algorithm/202601/40combinationSum2/combinationSum2.go#L33-L35) 的判断
  - i > 0：当前有前驱元素。
  - candidates[i] == candidates[i-1]：当前值和前一个值相同。
  - !used[i-1]：前一个相同值没有被放入当前路径（说明我们在“同一层级”里又遇到了一份相同的数）。
- 为什么要“return”：在这套“先不选 i 再尝试选 i”的二叉搜索里，这个 return 等价于“跳过 i 的‘选’分支”。因为“i-1 没被选时，再选 i”会生成与“选 i-1，跳过 i”完全相同的组合，从而重复。直接返回可避免进入这个重复的“选 i”分支。

简短例子
- candidates 排序后为 [1,1,2]，target=2。
- 当来到第二个 1（i=1）时，如果上一个 1 未在路径中（!used[0]），那么“选择 i=1 这个 1”会和“选择 i=0 的 1”形成相同组合的另一条路径，所以剪掉这条“选 i=1”的分支。

**与常见写法的差异**
- 本实现是“二叉 include/exclude 形”的 DFS，故用 used[i-1] 判断同层去重，命中时用 return 结束当前节点的“选分支”。
- 更常见的组合型写法是“for 枚举起点 start…n-1”，在循环里用 continue 去重与 break 剪枝（见下节）。两种思路等价，但去重判断位置与语义略不同：
  - 二叉展开：用 “相同值且前一个没选” 判重 → return 掉“选分支”。
  - for 枚举：用 “i > start 且相同值” 判重 → continue 跳过本次 i。

**还可以怎么剪枝/去重（更主流且简洁的方式）**
- 循环枚举 + 两条经典剪枝：
  - 去重：在循环层判重，只保留某层的第一个相同值
    - if i > start && candidates[i] == candidates[i-1] { continue }
  - 目标越界提前停止（排序前提下成立）
    - if candidates[i] > remain { break }
- 参考实现（思路示意，等价于你现在的功能）：

```go
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    res := [][]int{}
    path := []int{}
    var dfs func(start, remain int)
    dfs = func(start, remain int) {
        if remain == 0 {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        for i := start; i < len(candidates); i++ {
            // 层内去重：同一层只用第一个相同数字
            if i > start && candidates[i] == candidates[i-1] {
                continue
            }
            // 单调剪枝：后面只会更大，直接停止这一层
            if candidates[i] > remain {
                break
            }
            path = append(path, candidates[i])
            dfs(i+1, remain-candidates[i]) // 每个数只能用一次，所以从 i+1 开始
            path = path[:len(path)-1]
        }
    }
    dfs(0, target)
    return res
}
```

**在你当前写法中还能加的有效剪枝**
- 单调剪枝（与上面等价）：在准备“选 i”前，如果 s + candidates[i] > target，可以直接 return，因为数组已排序，选更靠后的只会更大：
  - 放在 L36 之前加：如果 s + candidates[i] > target { return }
- 小提示：你文件里 L29 的 if used[i] { return } 在这套二叉结构中基本不会为真（进入 dfs(i, …) 时 used[i] 总是被复位为 false），可以去掉以简化逻辑。

**另一种思路：按值分组计数去重（更强的去重与剪枝）**
- 先把 candidates 压缩为唯一值数组 vals 及其计数 cnts。对每个值 v，枚举选 0..k 次（k ≤ cnt 且 v*k ≤ remain），递归到下一值。
- 优点：天然无重复组合，且一次性跳过整段相同值，搜索树更浅；缺点：实现稍多几行，适合大批重复数的输入。

**小结**
- 你这段判断的本质是“同层去重”：前一个相同值没被选时，跳过当前相同值的“选分支”，避免重复组合。
- 更主流且直观的做法是在 for 枚举层使用“i > start && a[i] == a[i-1] → continue”，并结合“a[i] > remain → break”的单调剪枝。
- 如输入中重复较多，可考虑“按值计数”的分组枚举方式，进一步减少搜索分支。
