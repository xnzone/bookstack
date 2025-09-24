---
title: "二分搜索"
date: 2022-08-04T00:43:42+08:00
tags: ["leetcode", "binary search", "c++"]
image: https://s2.loli.net/2025/09/24/aNzcSETVp5LlgHd.jpg
cover: false
weight: 811
---

## Search in Rotated Sorted Array
[LeetCode](https://leetcode.com/problems/search-in-rotated-sorted-array)/[力扣](https://leetcode-cn.com/problems/search-in-rotated-sorted-array)

先判断中间的和尾部的数字大小，再判断target和首尾中三个数字大小关系，如此便能进行二分搜索

代码实现

```cpp
int search(vector<int>& nums, int target) {
    if (nums.empty()) return -1;
    int left = 0, right = nums.size() - 1;
    while(left < right) {
        if(nums[mid] > nums[right]) {
            if(target > nums[mid] || target < nums[left]){
                left = mid + 1;
            }else {
                right = mid;
            }
        } else {
            if (target > nums[mid] && target <= nums[right]){
                left = mid + 1;
            } else {
                right = mid;
            }
        }
    }
    return nums[left] == target ? left : -1;
}
```

## Search a 2D Matrix
[LeetCode](https://leetcode.com/problems/search-a-2d-matrix)/[力扣](https://leetcode.com/problems/search-a-2d-matrix)

先找到所在的行，再找所在的列

代码实现

```cpp
bool searchMatrix(vector<vector<int>>& matrix, int target) {
    int m = matrix.size();
    if (m == 0) return false;
    int n = matrix[0].size();
    if (n == 0) return false;
    
    int row = 0;
    int left = 0, right = m;
    
    while(left < right) {
        int mid = left + (right - left) / 2;
        
        if (matrix[mid][0] <= target && matrix[mid][n - 1] >= target) {
            row = mid;
            break;
        } else if (matrix[mid][0] > target) {
            right = mid;
        } else {
            left = mid + 1;
        }
    }
    
    left = 0; right = n;
    while(left < right) {
        int mid = left + (right - left) / 2;
        
        if(matrix[row][mid] == target){
            return true;
        }else if(matrix[row][mid] < target) {
            left = mid + 1;
        }else {
            right = mid;
        }
    }
    return false;
}
```

## Search in Rotated Sorted Array II
[LeetCode](https://leetcode.com/problems/search-in-rotated-sorted-array-ii)/[力扣](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii)

代码实现

```cpp
bool search(vector<int>& nums, int target) {
    int n = nums.size();
    int l = 0;
    int r = n;
    int m;
    while(l!=r){
        m = l+(r-l)/2;
        if(nums[m] == target)
            return true;
        if(nums[l] == nums[m]){
            l++;continue;
        }
        if(nums[l]<nums[m]){
            if(nums[l]<=target && target<nums[m])
                r = m;
            else
                l = m+1;
        }
        else{
            if(nums[m]<target && target<=nums[r-1])
                l = m+1;
            else
                r = m;
        }
    }
    return false;
}
```

## Find Minimum in Rotated Sorted Array
[LeetCode](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array)/[力扣](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array)

代码实现

```cpp
int findMin(vector<int>& nums) {
    int left = 0;
    int right = nums.size() - 1;
    if(nums[left] <= nums[right]) return nums[0];
    
    while(left < right) {
        int mid = left + (right - left) / 2;
        if(nums[mid] > nums[left]){
            left = mid;
        } else {
            right = mid;
        }
    }
    
    return nums[left + 1];
    
}
```

## Find Peak Element
[LeetCode](https://leetcode.com/problems/find-peak-element)/[力扣](https://leetcode-cn.com/problems/find-peak-element)

代码实现

```cpp
int findPeakElement(vector<int>& nums) {
    return search(nums, 0, nums.size() -1);
}
int search(vector<int>& nums, int l, int r) {
    if (l == r) return l;
    int mid = l + (r - l) / 2;
    if(nums[mid] > nums[mid + 1])
        return search(nums, l, mid);
    
    return search(nums, mid+1, r);
}
```

## Count of Range Sum
[LeetCode](https://leetcode.com/problems/count-of-range-sum)/[力扣](https://leetcode-cn.com/problems/count-of-range-sum)

代码实现

```cpp
int countRangeSum(vector<int>& nums, int lower, int upper) {
    int res = 0;
    long long sum = 0;
    multiset<long long> sums;
    sums.insert(0);
    for (int i = 0; i < nums.size(); ++i) {
        sum += nums[i];
        res += distance(sums.lower_bound(sum - upper), sums.upper_bound(sum - lower));
        sums.insert(sum);
    }
    return res;
}
```

## Find Smallest Letter Greater Than Target
[LeetCode](https://leetcode.com/problems/find-smallest-letter-greater-than-target)/[力扣](https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target)

代码实现

```cpp
char nextGreatestLetter(vector<char>& letters, char target) {
    int left = 0;
    int right = letters.size() - 1;
    while(left <= right) {
        int mid = left + (right - left) / 2;
        if(letters[mid] == target){
            left = mid + 1;
            // break;
        }else if (letters[mid] < target) {
            left = mid + 1;
        }else {
            right = mid - 1;
        }
    }
    
    return left >= letters.size() ? letters[0] : letters[left];
}
```

## Binary Search
[LeetCode](https://leetcode.com/problems/binary-search)/[力扣](https://leetcode-cn.com/problems/binary-search)


代码实现

```cpp
int search(vector<int>& nums, int target) {
    int left = 0;
    int right = nums.size() - 1;
    
    while(left <= right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] == target) return mid;
        if (nums[mid] < target) left = mid + 1;
        else right = mid - 1;
    }
    
    return -1;
}
```


## Peak Index in a Mountain Array
[LeetCode](https://leetcode.com/problems/peak-index-in-a-mountain-array)/[力扣](https://leetcode-cn.com/problems/peak-index-in-a-mountain-array)

代码实现

```cpp
int peakIndexInMountainArray(vector<int>& A) {
    int left = 0;
    int right = A.size() - 1;
    while(left <= right) {
        int mid = left + (right - left) / 2;
        if (A[mid] > A[mid + 1]) {
            right = mid - 1;
        }else {
            left = mid + 1;
        }
    }
    return left;
}
```