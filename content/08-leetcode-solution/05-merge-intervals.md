---
title: "区间合并"
date: 2022-06-27T15:43:42+08:00
tags: ["leetcode", "merge intervals", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 805
---

## Merge Intervals
[LeetCode](https://leetcode.com/problems/merge-intervals)/[力扣](https://leetcode-cn.com/problems/merge-intervals)

- 先排序
- 然后遍历组合
- 如果满足合并条件， 当前区间的末尾(两个取大的)

代码实现

```cpp
vector<vector<int>> merge(vector<vector<int>>& intervals) {
    sort(intervals.begin(), intervals.end());
    vector<vector<int>> res;
    int n = intervals.size();
    if(n <= 1) return intervals;
    vector<int> cur = intervals[0];
    for(int i = 1; i < n; i++) {
        if(cur[1] < intervals[i][0]) {
            res.push_back(cur);
            cur = intervals[i];
        } else {
            cur[1] = max(cur[1], intervals[i][1]);
        }
    }
    res.push_back(cur);
    return res;
}
```

## Insert Interval
[LeetCode](https://leetcode.com/problems/insert-interval)/[力扣](https://leetcode-cn.com/problems/insert-interval)

- 需要注意给定的组合都是非重叠的
- 前一个最后一个数与下一个第一个数相等时，不能合并

代码实现

```cpp
vector<vector<int>> insert(vector<vector<int>>& intervals, vector<int>& newInterval) {
    vector<vector<int>> res;
    for(int i = 0; i < intervals.size(); i++) {
        if(intervals[i][1] < newInterval[0]) {
            res.push_back(intervals[i]);
        } else if (intervals[i][0] > newInterval[1]) {
            res.push_back(newInterval);
            newInterval = intervals[i];
        } else {
            newInterval = {min(intervals[i][0], newInterval[0]), max(intervals[i][1], newInterval[1])};
        }
    }
    res.push_back(newInterval);
    return res;
}
```

## Non-overlapping Intervals
[LeetCode](https://leetcode.com/problems/non-overlapping-intervals)/[力扣](https://leetcode-cn.com/problems/non-overlapping-intervals)

代码实现

```cpp
int eraseOverlapIntervals(vector<vector<int>>& intervals) {
    int cnt = 0;
    sort(intervals.begin(), intervals.end());
    for(int i = 1; i < intervals.size(); i++) {
        if(intervals[i - 1][1] > intervals[i][0]) {
            intervals[i][1] = min(intervals[i - 1][1], intervals[i][1]);
            cnt++;
        }
    }
    return cnt;
}
```

## Minimum Number of Arrows to Burst Balloons
[LeetCode](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons)/[力扣](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons)

- 用一个数组保存当前范围
- 如果下一个在这个范围内，就更新这个范围大小
- 如果不在这个范围内，就结果加一
- 最后把范围更新为新的数组组合就行

代码实现

```cpp
int findMinArrowShots(vector<vector<int>>& points) {
    sort(points.begin(), points.end());
    int n = points.size();
    if(n <= 1) return n;
    vector<int> cur = points[0];
    int cnt = 1;
    for(int i = 1; i < n; i++) {
        if(points[i][0] > cur[1]) {
            cur = points[i];
            cnt++;
        } else {
            cur[0] = max(points[i][0], cur[0]);
            cur[1] = min(points[i][1], cur[1]);
        }
    }
    return cnt;
}
```

## Task Scheduler
[LeetCode](https://leetcode.com/problems/task-scheduler)/[力扣](https://leetcode-cn.com/problems/task-scheduler)

代码实现

```cpp
int leastInterval(vector<char>& tasks, int n) {
    vector<int> map(26, 0);
    for(char c : tasks) {
        map[c - 'A']++;
    }
    sort(map.begin(), map.end());
    int max_val = map[25] - 1, idle_slots = max_val * n;
    for(int i = 24; i >= 0 && map[i] > 0; i--) {
        idle_slots -= min(map[i], max_val);
    }
    return idle_slots > 0 ? idle_slots + tasks.size() : tasks.size();
}
```

## Interval List Intersections
[LeetCode](https://leetcode.com/problems/interval-list-intersections)/[力扣](https://leetcode-cn.com/problems/interval-list-intersections)

代码实现

```cpp
vector<vector<int>> intervalIntersection(vector<vector<int>>& A, vector<vector<int>>& B) {
    vector<vector<int>> res;
    int i = 0, j = 0;
    while(i < A.size() && j < B.size()) {
        int lo = max(A[i][0], B[j][0]);
        int hi = min(A[i][1], B[j][1]);
        if(lo <= hi) {
            res.push_back(lo, hi);
        }
        if(A[i][1] < B[j][1]) {
            i++;
        } else {
            j++;
        }
    }
    return res;
}
```