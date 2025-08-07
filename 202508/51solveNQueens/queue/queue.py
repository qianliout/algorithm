


class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        ans=[]
        col = [0]*n # 
        def valid(r,c):
            for R in range(r):
                C = col[R]
                if r+c ==R+C or r-c ==R-C:
                    return False
            return True

        # r 表示当前要枚举的行号，s 表示还可以使用的列号
        def dfs(r,s):
            if r==n:
                ans.append(['.'*c+'Q'+'.'*(n-c-1) for c in col])
                return
            for c in s:
                if valid(r,c):
                    col[r] = c
                    dfs(r+1,s-{c})
        dfs(0,set(range(n)))            

        return ans   
