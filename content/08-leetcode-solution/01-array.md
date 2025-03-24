---
title: "数组"
date: 2022-06-23T15:43:42+08:00
tags: ["leetcode", "array", "c++"]
author: xnzone
image: /covers/leetcode-solution.jpg
cover: false
weight: 801
---

## First Missing Positive
[LeetCode](https://leetcode.com/problems/first-missing-positive)/[力扣](https://leetcode-cn.com/problems/first-missing-positive)

- 用一个数组保存`0~n`是否有数，假设都为1
- 遍历数组，如果出现了，设为0，大于n或者小于0，忽略
- 最后遍历一遍就是结果，如果跑到最后，直接输出`n+1`

代码实现

```cpp
int firstMissingPositive(vector<int> &nums) {
    int n = nums.size();
    vector<int> tmp(n + 1, 1);
    for (int i = 0; i < n; i++) {
        if (nums[i] < 0 || nums[i] > n) continue;
        tmp[i] = 0;
    }
    for (int i = 1; i <= n; i++) {
        if (tmp[i] == 1) return i;
    }
    return n + 1;
}
```


## Rotate Image
[LeetCode](https://leetcode.com/problems/rotate-image)/[力扣](https://leetcode-cn.com/problems/rotate-image)

- 分层旋转，每次旋转交换四个数字。用一个中间变量保存其中一个数
- 一层旋转结束了，就往里旋转，直到层数大于`n / 2`
- 注意判断特殊情况

```cpp
void rotate(vector<vector<int>> &matrix) {
    int n = matrix.size();
    int layer = 0;
    while(layer < n / 2) {
        for (int i = 0; i < n - 2 * layer - 1; i++) {
            int tmp = matrix[layer][layer + i];
            matrix[layer][i + layer] = matrix[n - layer -i - 1][layer];
            matrix[n - layer - i - 1][layer] = matrix[n - layer - 1][n - layer -i - 1];
            matrix[n - layer - 1][n - layer - i - 1] = matrix[layer + i][n - layer - 1];
            matrix[layer + i][n - layer - 1] = tmp; 
        }
        layer++;
    }
}
```

- 直接新建一个数组保存

```cpp
void rotate(vector<vector<int>> &matrix) {
    vector<vector<int>> tmp(matrix);
    for (int i = 0; i < matrix.size(); i++) {
        for (int j = 0; j < matrix[i].size(); j++) {
            matrix[i][j] = temp[matrix.size() - j - 1][i];
        }
    }
}
```

- 先reserve数组
- 再对角线交换`i`, `j`

```cpp
void rotate(vector<vector<int>> &matrix) {
    reserve(matrix.begin(), matrix.end());
    for (int i = 0; i < matrix.size(); i++) {
        for (int j = 0; j < matrix[i].size(); j++) {
            swap(matrix[i][j], matrix[j][i]);
        }
    }
}
```

## Spiral Matrix
[LeetCode](https://leetcode.com/problems/spiral-matrix)/[力扣](https://leetcode-cn.com/problems/spiral-matrix)

- 分层输出，每一层输出上、右、下、左四个方向
- 输出完之后，上、左加一，右、下减一
- 如此循环，直到上大于下或左大于右

![回形输出](https://leetcode.com/problems/spiral-matrix/Figures/54_spiralmatrix.png)

```cpp
vector<int> spiralOrder(vector<vector<int>> &matrix) {
    vector<int> res;
    if (matrix.size() <= 0) return res;
    int top = 0, bot = matrix.size() - 1;
    int left = 0, righ = matrix[0].size() - 1;
    while(top <= bot && left <= righ) {
        for(int i = 0; i < left; i <= righ; i++) res.push_back(matrix[top][i]);
        for(int i = top + 1; i <= bot; i++) res.push_back(matrix[i][righ]);
        if (top < bot; && left < righ) {
            for (int i = righ - 1; i > left; i--) res.push_back(matrix[bot][i]);
            for (int i = bot; i > top; i--) res.push_back(matrix[i][left]);
        }
        top++; bot--; left++;righ--;
    }
    return res;
}
```

## Set Matrix Zeroes
[LeetCode](https://leetcode.com/problems/set-matrix-zeroes)/[力扣](https://leetcode-cn.com/problems/set-matrix-zeroes)

- 遍历矩阵，遇到0，把行和列都变成0
- 用map保存需要变换的行和列

```cpp
void setZeroes(vector<vector<int>> &matrix) {
    map<int, int> mi;
    map<int, int> mj;
    int m = matrix.size(), n = matrix[0].szie();
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            if (matrix[i][j] == 0) {
                mi[i] = 1; mj[j] = 1;
            }
        }
    }
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            if (mi[i] != 0 || mj[j] != 0) {
                matrix[i][j] = 0;
            }
        }
    }
}
```

## Word Search
[LeetCode](https://leetcode.com/problems/word-search)/[力扣](https://leetcode-cn.com/problems/word-search)

- 递归实现。如果前一个节点满足条件，则把其他周围四个节点分别进行递归
- 递归之前需要用一个标志位表示该节点已经被访问过了。递归之后，将该标志位归零

```cpp
bool exist(vector<vector<char>> &board, string word) {
    for(int i = 0; i < board.size(); i++) {
        for(int j = 0; j < board[0].size(); j++) {
            if (helper(i, j, board, word, 0)) return true;
        }
    }
    return false;
}
bool helper(int x, int y, vector<vector<char>> &board, string &word, int idx) {
    if (x < 0 || y < 0 || x >= board.size() || y >= board[0].size() || board[x][y] != word[idx]) return false; 
    if (idx == word.size() - 1) return true;

    char t = board[x][y];
    board[x][y] = '0';
    bool check = helper(x - 1, y, board, word, idx + 1) ||
    helper(x, y - 1, board, word, idx + 1) ||
    helper(x + 1, y, board, word, idx + 1) ||
    helper(x, y + 1, board, word, idx + 1);
    board[x][y] = t;
    return check
}
```

## Longest Consecutive Sequence
[LeetCode](https://leetcode.com/problems/longest-consecutive-sequence)/[力扣](https://leetcode-cn.com/problems/longest-consecutive-sequence)

- 用set保存当前的数
- 遍历set， 遍历到某个数时，查找这个数减1是否存在set中
- 如果存在，进入下一个循环
- 如果不存在，说明比当前这个set小的连续序列没有，就递增当前这个数
- 每次递增时，用一个变量保存当前递增的个数，直到递增的数字不在set中出现为止
- 最后取递增个数和当前结果的最大值

```cpp
int longestConsecutive(vector<int>& nums) {
    unordered_set<int> s;
    for (int i = 0; i < nums.size(); i++) {
        s.insert(nums[i]);
    }
    int res = 0;
    for (int i = 0; i < nums.size(); i++) {
        int streak = 0;
        if(s.find(nums[i] - 1) != s.end()) continue;
        streak = 1;
        int temp = nums[i]
        while(s.find(++a) != s.end()) {
            streak++;
        }
        res = max(res, streak);
    }
    return res;
}
```

## Single Number
[LeetCode](https://leetcode.com/problems/single-number)/[力扣](https://leetcode-cn.com/problems/single-number)

这题本身很简单，但是有个巧妙的数学解法

- A异或A = 0
- A异或0 = A
- A异或B异或A = A异或A异或B = B

```cpp
int singleNumber(vector<int>& nums) {
    int a = 0;
    for(int i = 0; i < nums.size(); i++) {
        a ^= nums[i];
    }
    return a;
}
```

## Contains Duplicate
[LeetCode](https://leetcode.com/problems/contains-duplicate)/[力扣](https://leetcode-cn.com/problems/contains-duplicate)

- 排序，排序完遍历，看前后两个数是否相等

- 另一种解法，用map或set保存遍历过的，如果已经存在map中，直接返回true(以空间换时间)

```cpp
bool containsDunplicate(vector<int>& nums) {
    if(nums.size() <= 0) return false;
    sort(nums.begin(), nums.end());
    for(int i = 0; i < nums.size() - 1; i++) {
        if (nums[i] == nums[i+ 1]) return true;
    }
    return false; 
}
```

## Product of Array Except Self
[LeetCode](https://leetcode.com/problems/product-of-array-except-self)/[力扣](https://leetcode-cn.com/problems/product-of-array-except-self)

- 先从左向右遍历，但只保存`nums[i - 1]` 和 `prod[i - 1]`的乘积
- 用`prod`从右往左与`nums`相乘
- 最后结果就是除去了当前这个数的乘积

```cpp
vector<int> productExeceptSelf(vector<int>& nums) {
    int n = nums.size();
    vector<int> res(nums);
    res[0] = 1;
    for(int i = 1; i < n; i++) {
        res[i] = nums[i - 1] * res[i - 1];
    }
    int R = 1;
    for(i = n - 1; i >= 0; i--) {
        res[i] = res[i] * R;
        R *= nums[i];
    }
    return res;
}
```

## Missing Number
[LeetCode](https://leetcode.com/problems/missing-number)/[力扣](https://leetcode-cn.com/problems/missing-number)

- 直接求和
- 把理论值和实际值相减

```cpp
int missingNumber(vector<int>& nums) {
    int sum = 0;
    int n = nums.size();
    for(int i = 0; i < nums.size(); i++) {
        sum += nums[i];
    }
    return n * (n - 1) / 2 - sum;
}
```

## Find the Duplicate Number
[LeetCode](https://leetcode.com/problems/find-the-duplicate-number)/[力扣](https://leetcode-cn.com/problems/find-the-duplicate-number)

- 快慢指针，可以看作一个环形的链表
- 一个快指针，总能遇到慢指针
- 最后慢指针从头开始，快慢指针每次运动一步
- 两个相等的时候，快指针就是结果

```cpp
int findDuplicate(vector<int>& nums) {
    int tortoise = nums[0], hare = nums[0];
    do {
        tortoise = nums[tortoise];
        hare = nums[nums[hare]];
    } while(tortoise != hare);

    tortoise = nums[0];
    while(tortoise != hare) {
        tortoise = nums[tortoise];
        hare = nums[hare];
    }
    return hare;
}
```

## Find All Duplicates in an Array
[LeetCode](https://leetcode.com/problems/find-all-duplicates-in-an-array)/[力扣](https://leetcode-cn.com/problems/find-all-duplicates-in-an-array)

- 修改数组，将出现的那个地方改成负数
- 后续遍历到，如果为负数，就是重复了

```cpp
vector<int> findDuplicates(vector<int>& nums) {
    vector<int> v;
    for(int i = 0; i < nums.size(); i++){
        int t = abs(nums[i]) - 1;
        if(nums[i] < 0) {
            v.push_back(t + 1);
        }
        nums[t] = -nums[t];
    }
    return v;
}
```

## Find All Numbers Disappeared in an Array
[LeetCode](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array)/[力扣](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array)

- 先把所有数字改成负数
- 再遍历，如果位置为正数，则说明数字没有出现

```cpp
vector<int> findDisappearedNumbers(vector<int>& nums) {
    vector<int> res;
    for(int i = 0; i < nums.size(); i++) {
        nums[abs(nums[i]) - 1] = -abs(nums[abs(nums[i]) - 1]);
    }
    for(int i = 0; i < nums.size(); i++) {
        if(nums[i] > 0 ) {
            res.push_back(i + 1);
        }
    }
    return res;'
}
```

## Circular Array Loop
[LeetCode](https://leetcode.com/problems/circular-array-loop)/[力扣](https://leetcode-cn.com/problems/circular-array-loop)

```cpp
int next(vector<int>& nums, int idx) {
    int res = (idx + nums[idx]);
    while(res < 0) res += nums.size();
    while(res >= nums.size()) res -= nums.size();
    return res;
}

bool circularArrayLoop(vector<int>& nums) {
    for(int i = 0; i < nums.size(); i++) {
        if (nums[i] == 0) continue;
        int j = i, k = next(nums, i);
        while(nums[k] * nums[i] > 0 && nums[next(nums, k)] * nums[i] > 0) {
            if(j == k) {
                if(j != next(nums, j)) return true;
                else break;
            }
            j = next(nums, j);
            k = next(nums, next(nums, k));
        }
        j = i;
        while(nums[j] * nums[i] > 0 ) {
            nums[j] = 0;
            j = next(nums, j);
        }
    }
    return false;
}
```

## Shortest Unsorted Continuous Subarray
[LeetCode](https://leetcode.com/problems/shortest-unsorted-continuous-subarray)/[力扣](https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray)

- 先排序
- 把辅助的数组与原数组比较

```cpp
int findUnsortedSubarray(vector<int>& nums) {
    vector<int> helper = nums;
    sort(helper.begin(), helper.end());
    int left = nums.size(), right = 0;
    for(int i = 0; i < nums.size(); i++) {
        if(helper[i] != nums[i]) {
            left = min(left, i);
            right = max(right, i);
        }
    }
    return right - left >= 0 ? right - left + 1 : 0;
}
```


## Number of Matching Subsequences
[LeetCode](https://leetcode.com/problems/number-of-matching-subsequences)/[力扣](https://leetcode-cn.com/problems/number-of-matching-subsequences)

- 用map或vector保存字典
- 然后用二分法查找

```cpp
int numMatchingSubseq(string S, vector<string>& words) {
    int res = 0;
    vector<vector<int>> vc(128, vector<int>(0, 0));
    for(int i = 0; i < S.size(); i++) {
        vc[S[i]].push_back(i);
    }
    for(string w: words) {
        int pos = -1;
        bool sub = true;
        for(char c: w) {
            auto it = upper_bound(vc[c].begin(), vc[c].end(), pos);
            if (it == vc[c].end()) {
                sub = false;
            } else {
                pos = *it;
            }
        }
        if (sub) res++;
    }
    return res;
}
```