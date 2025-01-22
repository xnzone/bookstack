---
title: "双指针"
date: 2022-06-25T15:43:42+08:00
tags: ["leetcode", "two points", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 803
---

## Two Sum
[LeetCode](https://leetcode.com/problems/two-sum)/[力扣](https://leetcode-cn.com/problems/two-sum)

- 先排序
- 排序完，用两个指针分别指向第一个和最后一个
- 根据和，分别移动前后指针

上述方法需要额外保存排序前的位置，所以可解，但不推荐，现给出排序好的解法(不是当前题目解法)

```cpp
vector<int> twoSum(vector<int>& nums, int target) {
    int n = nums.size();
    vectort<int> res;
    for(int i = 0, j = n - 1; i < j;) {
        int t = nums[i] + nums[j];
        if(t > target) {
            j--;
        } else if(t < target) {
            i++;
        } else {
            res.push_back(i);
            res.push_back(j);
            break;
        }
    }
    return res;
}
```

- 用map保存已经计算过数值位置
- 如果找到target - 当前数在map中，直接输出结果
- 如果不存在，则把结果放入map中

```cpp
vector<int> twoSum(vector<int>& nums, int target) {
    int n = nums.size();
    vector<int> res;
    map<int, int> nmap;

    for(int i = 0; i < n; i++) {
        int t = target - nums[i];
        if(nmap.find(t) != nmap.end()) {
            res.push_back(nmap[t]);
            res.push_back(i);
            break;
        }
        nmap[nums[i]] = i;
    }
    return res;
}
```

## Container With Most Water
[LeetCode](https://leetcode.com/problems/container-with-most-water)/[力扣](https://leetcode-cn.com/problems/container-with-most-water)

- 前后指针遍历
- 哪边数字小，移动哪边

```cpp
int maxArea(vector<int>& height) {
    int res = -1, i = 0, j = height.size() - 1;
    while(i <= j) {
        int area = (j - i) * min(height[i], height[j]);
        res = max(res, area);
        if(height[i] < height[j]) {
            i++;
        } else {
            j--;
        }
    }
    return res;
}
```

## 3Sum
[LeetCode](https://leetcode.com/problems/3sum)/[力扣](https://leetcode-cn.com/problems/3sum)

- 分别用两个数组保存出现的次数和去重后排序的数组
- 然后计算最终是否满足条件
- 也可以用map来保存上述两个值，过程是一样的

```cpp
vector<vector<int>> threeSum(vector<int>& nums) {
  vector<vector<int>> res;
  if(nums.size()==0) return res;
  int max=0, min=0;
  // 这里是计算数组的最大值和最小值
  for(int d:nums){
      max=max<d?d:max;
      min=min>d?d:min;
  }
  // 下面这些是计算最大最小值里面的这个数字出现了多少次，比如
  // nums = [-1,0,1,2,-1,-4], tb=[1,0,0,2,1,1,1]
  int tb_size=max-min+1, m=-min;
  vector<char> tb(tb_size, 0);
  for(int d:nums){
      if(++tb[d+m]>3)tb[d+m]=3;
  }
  if(tb[m]==2) tb[m]=1;
  // 下面这些是计算不重复的数组并排序过的，比如
  // nums = [-1,0,1,2,-1,-4], v= [-4,-1,0,1,2]
  vector<int> v;
  for(int i=0;i<tb_size;++i)
      if(tb[i]>0) v.push_back(i-m);
  int vsz=size(v);
  // 下面这些我没看懂
  for(int i=0;i<vsz;++i){
      int vi=v[i];
      if(vi>0) break;
      for(int j=i;j<vsz;++j){
          int vj=v[j];
          int vk=-vi-vj;  // 求相反数
          if(vk<vj) break;   //说明第三个数已经到第二个数后面去了，所以直接break
          if(vk>max||tb[vk+m]==0) continue;  //说明相反数在里面并不存在，直接进行下一次循环
          if(tb[vj+m]>1||vi<vj && vj<vk) res.push_back({vi,vj,vk}); // 求出的结果是否合法
      }
  }
  return res;
}
```

## 3Sum Closest
[LeetCode](https://leetcode.com/problems/3sum-closest)/[力扣](https://leetcode-cn.com/problems/3sum-closest)

- 第一个从`0~n`遍历，第二个从`i~n`遍历
- 如果sum大于target， 则右指针移动
- 如果sum小于target，则左指针移动

```cpp
int threeSumClosest(vector<int>& nums, int target) {
    std::sort(nums.begin(), nums.end());
    int n = nums.size();
    if(n < 3) return 0;
    int res = nums[0] + nums[1] + nums[2];
    for(int i = 0; i < n; i++) {
        int j = i + 1, k = n - 1;
        while(j < k) {
            int sum = nums[i] + nums[j] + nums[k];
            res = abs(target - sum) < abs(res - target) ? sum : res;
            if(sum < target) {
                j++;
            } else if(sum > target) {
                k--;
            } else {
                res = target;
                return res;
            }
        }
    }
    return res;
}
```

## Trapping Rain Water
[LeetCode](https://leetcode.com/problems/trapping-rain-water)/[力扣](https://leetcode-cn.com/problems/trapping-rain-water)

- 只要左右都有黑色方块，那么积水一直会增加
- 所以可以用双向指针解决问题

```cpp
int trap(vector<int>& height) {
    int res = 0, i = 0, j = height.size() - 1;
    int lmax = 0, rmax = 0;
    while(i < j) {
        if(height[i] < height[j]) {
            height[i] >= lmax ? (lmax = height[i]) : res += (lmax - height[i]);
            i++; 
        } else {
            height[j] >= rmax ? (rmax = height[j]) : res += (rmax - height[j]);
            j--;
        }
    }
    return res;
}
```

- 动态规划
- 从左到右，找到当前最大值`left_max`
- 从右到左，找到当前最大值`right_max`
- 遍历，找到`left_max`和`right_max`最小值，减去`height`就是积水

```cpp
int trap(vector<int>& height) {
    if(height.size() == 0) return 0;
    int res = 0, n = height.size();
    vector<int> left_max(n, 0), right_max(n, 0);
    left_max[0] = height[0];
    for(int i = 1; i < n; i++) {
        left_max[i] = max(height[i], left_max[i - 1]);
    }
    right_max[n - 1] = height[n - 1];
    for(int i = n - 2; i >= 0; i++) {
        right_max[i] = max(right_max[i + 1], height[i]);
    }

    for(int i = 1; i < n - 1; i++) {
        res += min(left_max[i], right[i]) - height[i];
    }
    return res;
}
```

## Sort Colors
[LeetCode](https://leetcode.com/problems/sort-colors)/[力扣](https://leetcode-cn.com/problems/sort-colors)

- 分别用前后指针代表0和2所在的位置，然后遍历中间的
- 如果是0， 跟左指针换
- 如果是2， 跟右指针换

```cpp
void sortColors(vector<int>& nums) {
    int n = nums.size(), n0 = 0, n2 = n - 1;
    int i = 0;
    while( i <= n2) {
        if(nums[i] == 0 && i != 0) {
            swap(nums[i], nums[n0++]);
        } else if(nums[i] == 2 && i != n2) {
            swap(nums[i], nums[n2--]);
        } else {
            i++;
        }
    }
}
```

## Minimum Window Substring
[LeetCode](https://leetcode.com/problems/minimum-window-substring)/[力扣](https://leetcode-cn.com/problems/minimum-window-substring)

- 双指针，先后两个指针
- 先移动右指针，用一个vector或map保存字符的个数
- 当所有字符全部存在时，移动左指针
- 直到不满足条件时，跳出循环

```cpp
string minWindow(string s, string t) {
    vector<int> count(128, 0);
    for(char c : t) count[c]++;

    int remain = t.size(), left = 0, right = 0, minStart = 0, minLen = INT_MAX;

    while(right < s.size()) {
        count[s[right]]--;
        if(count[s[right]] >= 0)remain--;
        right++;
        while(remain == 0) {
            if(right - left < minLen) {
                minStart = left;
                minLen = right - left;
            }
            count[s[left]]++;
            if(count[s[left]] > 0) remain++;
            left++;
        }
    }
    return minLen < INT_MAX ? s.substr(minStart, minLen) : "";
}
```

## Remove Duplicates from Sorted List
[LeetCode](https://leetcode.com/problems/remove-duplicates-from-sorted-list)/[力扣](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list)

- 一个指向当前的不重复尾部
- 一个指向下一个遍历
- 如果相等， 则 `left -> next = right -> next`
- 否则 `left = right`

```cpp
ListNode* deleteDuplicates(ListNode* head) {
    if(head == nullptr) return head;
    ListNode* left = head;
    ListNode* right = head->next;

    while(right != nullptr) {
        if(left -> val == right -> val) {
            left->next = right->next;
        } else {
            left = right;
        }
        right = right->next;
    }
    return head;
}
```

## Subarray Product Less Than K
[LeetCode](https://leetcode.com/problems/subarray-product-less-than-k)/[力扣](https://leetcode-cn.com/problems/subarray-product-less-than-k)

- 用一个数字保存乘积，先移动右指针
- 当乘积大于K时，一直除左边的
- 直到程序小于K或`left = right`时
- 每次结果是`left - right + 1`

```cpp
int numSubarrayProductLessThanK(vector<int>& nums, int k) {
    if(k == 0) return 0;
    int size = nums.size(), mul = 0, left = 0, right = 0, res = 0;

    while(right < size) {
        mul *= nums[right];
        while(mul >= k && left <= right) {
            mul / = nums[left++];
        }
        res += right - left + 1;
        ++right;
    }
    return res;
}
```


## Backspace String Compare
[LeetCode](https://leetcode.com/problems/backspace-string-compare)/[力扣](https://leetcode-cn.com/problems/backspace-string-compare)

- 从后面开始遍历，用一个数字记录当前space的个数
- 然后保存新的字符串即可

```cpp
string backspace(string str) {
    string res = "";
    int count = 0;

    for(int i = str.size() - 1; i >= 0; i--) {
        if(str[i] == '#') {
            count++;
            continue;
        } 
        if(count > 0 ) {
            count--;
        } else {
            res = str[i] + res;
        }
    }
    return res;
}

bool backspaceCompare(string S, string T) {
    return backspace(S) == backspace(T);
}
```

## Squares of a Sorted Array
[LeetCode](https://leetcode.com/problems/squares-of-a-sorted-array)/[力扣](https://leetcode-cn.com/problems/squares-of-a-sorted-array)

- 先找到大于等于0的分界点，然后计算平方值
- 如果左边大，则压入右边，右边++
- 如果右边大，则压入左边，左边--
- 最后把剩下的全都计算平方加进去

```cpp
vector<int> sortedSquares(vector<int>& A) {
    vector<int> res;
    int more = 0, less = 0;
    for(int i = 0; i < A.size(); i++) {
        if(A[i] >= 0 ) {
            more = i;
            less = i - 1;
            break;
        }
    }
    while(less >= 0 && more < A.size()) {
        int m = A[less] * A[less];
        int n = A[more] * A[more];
        if(m > n) {
            res.push_back(n);
            more++;
        }
        if(m < n) {
            res.push_back(m);
            less--;
        }
        if(m == n) {
            res.push_back(m);
            if(more != less) res.push_back(n);
            more++;
            less--;
        }
    }
    while(less >= 0) {
        res.push_back(A[less] * A[less]);
        less--;
    }
    while(more < A.size()) {
        res.push_back(A[more] * A[more]);
        more++;
    }
    return res;
}
```