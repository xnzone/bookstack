---
title: "回溯算法"
date: 2022-07-15T15:43:42+08:00
tags: ["leetcode", "backtrack", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 809
---


## Letter Combinations of a Phone Number
[LeetCode](https://leetcode.com/problems/letter-combinations-of-a-phone-number)/[力扣](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number)

代码实现

```cpp
vector<string> letterCombinations(string digits) {
    if(digits.length() == 0) return {};
    map<char,string> num_to_char{
        {'2',"abc"},
        {'3',"def"},
        {'4',"ghi"},
        {'5',"jkl"},
        {'6',"mno"},
        {'7',"pqrs"},
        {'8',"tuv"},
        {'9',"wxyz"}
    };
    vector<string> res;
    string combination(digits.length(),' ');
    letterCombinations(digits,combination,0,num_to_char,res);
    return res;
}
void letterCombinations(string digits, string& combination, int idx, map<char,string>& num_to_char, vector<string>& res){
    if(idx == digits.length()){
        res.push_back(combination);
        return;
    }
    
    for(char c:num_to_char[digits[idx]]){
        combination[idx] = c;
        letterCombinations(digits,combination,idx+1,num_to_char,res);
    }
}
```

## Generate Parentheses
[LeetCode](https://leetcode.com/problems/generate-parentheses)/[力扣](https://leetcode-cn.com/problems/generate-parentheses)

代码实现

```cpp
vector<string> generateParenthesis(int n) {
    vector<string> result;
    if (n <= 0) return result;
    string tmpString = "(";
    addItem(result, tmpString, n-1, 1);
    return result;
}
    
void addItem(vector<string> &result,string &tmpString, int pre, int end) {
    if (pre == 0 && end == 0) {
        result.push_back(tmpString);
        return;
    }
    if (pre > 0) {
        tmpString.push_back('(');
        addItem(result, tmpString, pre-1, end+1);
        tmpString.pop_back();
    }
    if (end > 0) {
        tmpString.push_back(')');
        addItem(result, tmpString, pre, end-1);
        tmpString.pop_back();
    }
   return;
}
```

## Sudoku Solver
[LeetCode](https://leetcode.com/problems/sudoku-solver)/[力扣](https://leetcode-cn.com/problems/sudoku-solver)

代码实现

```cpp
void solveSudoku(vector<vector<char>>& board) {
    helper(board,0,0);
}
bool helper(vector<vector<char>>& board, int i, int j ) {
    if(i == 9) return true;
    if(j == 9) return helper(board, i + 1,0);
    if(board[i][j] != '.') return helper(board, i, j + 1);
    for(int k = 1; k <= 9; k++) {
        if(safe(board, i, j, k )){
            board[i][j] = k + '0';
            bool check = helper(board, i, j + 1);
            if(check) return true;
        } 
    }
    board[i][j] = '.';
    return false;
}
bool safe(vector<vector<char>>& board, int i,int j, int target) {
    for(int k = 0; k < 9; k++) {
        if(board[i][k] == target + '0' || board[k][j] == target + '0') return false;
    }
    int m = i / 3, n = j / 3;
    for(int k1 = m * 3; k1 < 3 * m + 3; k1++) {
        for(int k2 = n * 3; k2 < 3 * n + 3; k2++) {
            if(board[k1][k2] == target + '0') return false;
        }
    }
    return true;
}
```

## Combination Sum
[LeetCode](https://leetcode.com/problems/combination-sum)/[力扣](https://leetcode-cn.com/problems/combination-sum)

代码实现

```cpp
vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
    vector<int> temp;
    vector<vector<int>> res;
    helper(candidates, target,0, 0, temp, res);
    return res;
}

void helper(vector<int>& can, int target,int index, int sum, vector<int>& temp, vector<vector<int>>& res) {
    if(sum == target){
        res.push_back(temp);
        return;
    }
    if(sum > target) return;
    
    for(int i = index; i < can.size(); i++) {
        temp.push_back(can[i]);
        helper(can, target,i, sum + can[i], temp, res);
        temp.pop_back();
    }
}
```

## Combination Sum II
[LeetCode](https://leetcode.com/problems/combination-sum-ii)/[力扣](https://leetcode-cn.com/problems/combination-sum-ii)

代码实现

```cpp
vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
    vector<int> temp;
    vector<vector<int>> res;
    sort(candidates.begin(),candidates.end());
    helper(candidates, target, temp, res, 0, 0);
    return res;
}
void helper(vector<int>& can, int target, vector<int>& temp, vector<vector<int>>& res,int sum, int index) {
    if(sum == target){
        res.push_back(temp);
        return;
    }
    if(sum > target) return;
    
    for(int i = index; i<can.size();i++) {
        if(i > index && can[i] == can[i-1]) continue;
        temp.push_back(can[i]);
        helper(can,target, temp, res, sum + can[i], i + 1);
        temp.pop_back();
    }
}
```

## Permutations
[LeetCode](https://leetcode.com/problems/permutations)/[力扣](https://leetcode-cn.com/problems/permutations)

代码实现

```cpp
vector<vector<int>> permute(vector<int>& nums) {
    vector<int> temp;
    vector<vector<int>> res;
    vector<bool> visited(nums.size(), false);
    helper(nums, nums.size(), temp, res, visited);
    return res;
}
void helper(vector<int>& nums, int len, vector<int>& temp, vector<vector<int>>& res,vector<bool> visited) {
    if(len <= 0) {
        res.push_back(temp);
        return;
    }
    for(int i = 0; i < nums.size(); i++) {
        if(!visited[i]) {
            temp.push_back(nums[i]);
            visited[i] = true;
            helper(nums,len - 1, temp, res, visited);
            visited[i] = false;
            temp.pop_back();
        }
    }
}
```

## Permutations II
[LeetCode](https://leetcode.com/problems/permutations-ii)/[力扣](https://leetcode-cn.com/problems/permutations-ii)

代码实现

```cpp
vector<vector<int>> permuteUnique(vector<int>& nums) {
    vector<int> temp;
    unordered_map<int, int> m;
    for(int i = 0; i < nums.size(); i++) {
        m[nums[i]]++;
    }
    vector<vector<int>> res;
    helper(nums, nums.size(), temp, res, m);
    return res;
}
void helper(vector<int>& nums, int len, vector<int>& temp, vector<vector<int>>& res, unordered_map<int,int>& m) {
    if(len <= 0) {
        res.push_back(temp);
        return;
    }
    
    for(auto& it : m) {
        if(it.second <= 0) continue;
        it.second--;
        temp.push_back(it.first);
        helper(nums, len - 1, temp, res, m);
        it.second++;
        temp.pop_back();
    }
}
```

## N-Queens
[LeetCode](https://leetcode.com/problems/n-queens)/[力扣](https://leetcode-cn.com/problems/n-queens)

代码实现

```cpp
vector<vector<int>> placements;
vector<int> placement, br, bd1, bd2;
void dfs(int col, int n) {
    if (col == n) {
        placements.push_back(placement);
        return;
    }
    for (int row = 0; row < n; ++row) {
        if (br[row] || bd1[row + col]) continue;
        int id = row - col < 0 ? 2 * n + row - col : row - col;
        if (bd2[id]) continue;
        placement.push_back(row);
        br[row] = bd1[row + col] = bd2[id] = 1;
        dfs(col + 1, n);
        br[row] = bd1[row + col] = bd2[id] = 0;
        placement.pop_back();
    }
}
vector<vector<string>> solveNQueens(int n) {
    br.resize(n);
    bd1.resize(2 * n + 1);
    bd2.resize(2 * n + 1);
    dfs(0, n);
    vector<vector<string>> boards(placements.size(), vector<string>(n, string(n, '.')));
    for (int i = 0; i < placements.size(); ++i) {
        for (int j = 0; j < n; ++j) {
            boards[i][placements[i][j]][j] = 'Q';
        }
    }
    return boards;
}
```

## Combinations
[LeetCode](https://leetcode.com/problems/combinations)/[力扣](https://leetcode-cn.com/problems/combinations)

代码实现

```cpp
vector<vector<int>> combine(int n, int k) {
    vector<int> t;
    vector<vector<int>> r;
    vector<int> v(n+1,0);
    helper(t,r,1,n,k,v);
    return r;
}

void helper(vector<int>& t, vector<vector<int>>& r,int idx, int n, int k,vector<int>& v) {
    if(t.size() == k) {
        r.push_back(t);
        return;
    }
    for(int i = idx; i <= n; i++) {
        if(v[i] == 1) continue;
        t.push_back(i);
        v[i] = 1;
        helper(t, r,i, n, k,v);
        t.pop_back();
        v[i] = 0;
    }
}
```

## Subsets
[LeetCode](https://leetcode.com/problems/subsets)/[力扣](https://leetcode-cn.com/problems/subsets)

代码实现

```cpp
vector<vector<int>> subsets(vector<int>& nums) {
    vector<int> t;
    vector<vector<int>> r;
    helper(t,r,nums,0);
    return r;
}

void helper(vector<int>& t, vector<vector<int>>& r, vector<int>& nums, int idx) {
    if(t.size() <= nums.size()) {
        r.push_back(t);
    }
    if(t.size() > nums.size() || idx > nums.size()) {
        return;
    }
    
    for(int i = idx; i < nums.size(); i++) {
        t.push_back(nums[i]);
        helper(t,r,nums,i + 1);
        t.pop_back();
    }
}
```

## Subsets II
[LeetCode](https://leetcode.com/problems/subsets-ii)/[力扣](https://leetcode-cn.com/problems/subsets-ii)

代码实现

```cpp
vector<vector<int>> subsetsWithDup(vector<int>& nums) {
    sort(nums.begin(), nums.end());
    vector<int> t;
    vector<vector<int>> r;
    helper(t,r,nums,0);
    return r;
}

void helper(vector<int>& t, vector<vector<int>>& r, vector<int>& nums, int idx) {
    if(t.size() <= nums.size()) {
        r.push_back(t);
    }
    if(t.size() > nums.size() || idx > nums.size()) return;
    
    for(int i = idx; i < nums.size(); i++) {
        if(i > idx && nums[i] == nums[i - 1]) continue;
        t.push_back(nums[i]);
        helper(t,r,nums,i+1);
        t.pop_back();
    }
}
```

## Palindrome Partitioning
[LeetCode](https://leetcode.com/problems/palindrome-partitioning)/[力扣](https://leetcode-cn.com/problems/palindrome-partitioning)

代码实现

```cpp
void _solve(int ind, string &s, vector<vector<bool>> &dp, vector<string> &ans, vector<vector<string>> &res){
    if(ind == s.length()){
        res.push_back(ans);
        return;
    }
    for(int i = ind; i < s.length(); i++)
        if(dp[ind][i]){
            ans.push_back(s.substr(ind, i - ind + 1));
            _solve(i + 1, s, dp, ans, res);
            ans.pop_back();
        }
}
    
vector<vector<string>> partition(string s) {
    vector<vector<string>> res;
    vector<string> ans;
    int n = s.length();
    vector<vector<bool>> dp(n, vector<bool>(n));
    
    for(int i = 0; i < n; i++)
        dp[i][i] = true;
    for(int l = 2; l <= n; l++){
        for(int i = 0; i < n - l + 1; i++){
            int j = i + l - 1;
            if(l == 2)
                dp[i][j] = s[i] == s[j] ? true : false;
            else
                dp[i][j] = s[i] == s[j] && dp[i + 1][j - 1] ? true : false;
         }
    }
    _solve(0, s, dp, ans, res);
    return res;
}
```

## Combination Sum III
[LeetCode](https://leetcode.com/problems/combination-sum-iii)/[力扣](https://leetcode-cn.com/problems/combination-sum-iii)

代码实现

```cpp
vector<vector<int>> combinationSum3(int k, int n) {
    vector<int> t;
    vector<vector<int>> r;
    vector<int> visited(9,0);
    helper(k,n, t, r, 0,1, 0,visited);
    return r;
}
void helper(int k, int n, vector<int>& t, vector<vector<int>>& r, int sum,int index, int idx,vector<int>& visited) {
    if(idx == k && sum == n) {
        r.push_back(t);
        return;
    }
    if(sum > n) return;
    
    for(int i = index; i <= 9; i++) {
        if(visited[i-1] == 1) continue;
        t.push_back(i);
        visited[i-1] = 1;
        helper(k, n, t, r, sum + i,i + 1, idx + 1,visited);
        t.pop_back();
        visited[i-1] = 0;
    }
}
```

## Target Sum
[LeetCode](https://leetcode.com/problems/target-sum)/[力扣](https://leetcode-cn.com/problems/target-sum)

代码实现

```cpp
int findTargetSumWays(vector<int>& nums, int target) {
    long sum=0,n=nums.size();
    for(auto x:nums) sum+=x;
    if((sum+target)&1 or target>sum) return 0;
    long newsum=(sum+target)/2;
    vector<vector<int>>dp(n+1,vector<int>(newsum+1,0));
    dp[0][0]=1;
    for(int i=1 ; i<=n ; i++)
        for(int j=0 ; j<=newsum ; j++)
        {
            if(nums[i-1]>j) dp[i][j]=dp[i-1][j];
            else dp[i][j]=dp[i-1][j]+dp[i-1][j-nums[i-1]];
        }
    return dp[n][newsum];
}
```

## Letter Case Permutation
[LeetCode](https://leetcode.com/problems/letter-case-permutation)/[力扣](https://leetcode-cn.com/problems/letter-case-permutation)

代码实现

```cpp
vector<string> letterCasePermutation(string S) {
    vector<string> r;
    helper(r,S,0);
    return r;
}

void helper(vector<string>& r, string s, int idx) {
    if(idx == s.size()) {
        r.push_back(s);
        return;
    }
    
    helper(r, s, idx + 1);
    
    if(isalpha(s[idx])){
        s[idx]=isupper(s[idx]) ? tolower(s[idx]) : toupper(s[idx]) ;
        helper(r,s,idx+1);
    }
}
```