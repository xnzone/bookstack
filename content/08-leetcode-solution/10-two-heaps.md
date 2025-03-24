---
title: "双堆问题"
date: 2022-07-20T15:43:42+08:00
tags: ["leetcode", "two heaps", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 810
---

## Find Median from Data Stream
[LeetCode](https://leetcode.com/problems/find-median-from-data-stream/)/[力扣](https://leetcode-cn.com/problems/find-median-from-data-stream/)

数组保存数据，add的时候直接在末尾插入，查找的时候，先排序，然后再算其中的结果

代码实现

```cpp
vector<int> data;
/** initialize your data structure here. */
MedianFinder() {
    
}

void addNum(int num) {
    data.push_back(num);
}

double findMedian() {
    sort(data.begin(), data.end());
    
    int n = data.size();
    if(n % 2 == 1) {
        return data[(n) / 2];
    }else {
        return (data[n/2] + data[(n/2) - 1]) * 0.5;
    }
}
```

上述方案应该过不了案例，插入的时候用二分查找，先找到要插入的位置，然后直接插入，find的时候直接计算就可以了

代码实现

```cpp
vector<int> data;

void addNum(int num){
    if(data.empty()){
        data.push_back(num);
    }else {
        data.insert(lower_bound(data.begin(),data.end(), num), num); // 二分查找到应该插入的位置
    }
}

double findMedian() {        
    int n = data.size();
    if(n % 2 == 1) {
        return data[(n) / 2];
    }else {
        return (data[n/2] + data[(n/2) - 1]) * 0.5;
    }
}
```

二叉堆，可以用两个stl的优先队列（内部是用二叉堆）来维护，一个队列是最大堆，一个队列是最小堆，最大堆保存的是小半部分的数字，最小堆保存的是大半部分的数字，这样队首的和就是中位数了

代码实现

```cpp
priority_queue<int, vector<int>, less<int>> max_heap;
priority_queue<int, vector<int>, greater<int>> min_heap;

void addNum(int num){
    max_heap.push(num);
    min_heap.push(max_heap.top());
    max_heap.pop();
    
    if(max_heap.size() < min_heap.size()){
        max_heap.push(min_heap.top());
        min_heap.pop();
    }
}

double findMedian() {
    return max_heap.size() > min_heap.size() ? max_heap.top() : (max_heap.top() + min_heap.top()) * 0.5;
}
```

## sliding window median
[LeetCode](https://leetcode.com/problems/sliding-window-median)/[力扣](https://leetcode-cn.com/problems/sliding-window-median)

将k中的数字用vector保存，分别定义add，del和find方法，分别表示添加和删除元素，以及查找中位数

代码实现

```cpp
vector<double> medianSlidingWindow(vector<int>& nums, int k) {
    vector<int> curs;
    vector<double> res;
    for(int i = 0; i < k; i++){
        add(curs, nums[i]);
    }
    res.push_back(find(curs));
    for(int i = k; i < nums.size(); i++){
        cout << "hi";
        del(curs, nums[i - k]);
        add(curs, nums[i]);
        res.push_back(find(curs));
    }
    return res;
}

double find(vector<int>& curs){
    int n = curs.size();
    if(n % 2 == 1){
        return curs[n / 2];
    }else{
        return curs[n / 2] /2.0 + curs[n / 2 - 1] / 2.0;
    }
}
void add(vector<int>& curs, int num){
    if(curs.empty()){
        curs.push_back(num);
    }else{
        curs.insert(lower_bound(curs.begin(), curs.end(), num), num);
    }
}
void del(vector<int>& curs, int num){
    curs.erase(lower_bound(curs.begin(), curs.end(), num));
}
```

用multiset来保存当前的数组，这是当前第一名的那个解法

代码实现

```cpp
vector<double> medianSlidingWindow(vector<int>& nums, int k) {
    vector<double> res;
    multiset<double> ms(nums.begin(), nums.begin() + k);
    auto mid = next(ms.begin(), k /  2);
    for (int i = k; ; ++i) {
        res.push_back((*mid + *prev(mid,  1 - k % 2)) / 2);        
        if (i == nums.size()) return res;
        ms.insert(nums[i]);
        if (nums[i] < *mid) --mid;
        if (nums[i - k] <= *mid) ++mid;
        ms.erase(ms.lower_bound(nums[i - k]));
    }
    
    return res;
}
```