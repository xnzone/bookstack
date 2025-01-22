---
title: "多路合并"
date: 2022-08-31T00:43:42+08:00
tags: ["leetcode", "top k", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 813
---

## Merge Two Sorted Lists
[LeetCode](https://leetcode.com/problems/merge-two-sorted-lists)/[力扣](https://leetcode-cn.com/problems/merge-two-sorted-lists)

```cpp
ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
    if(l2 == nullptr) return l1;
    if(l1 == nullptr) return l2;
    
    if(l1->val > l2->val) {
        l2->next = mergeTwoLists(l1,l2->next);
        return l2;
    } else {
        l1->next = mergeTwoLists(l1->next,l2);
        return l1;
    }
}
```

## Merge k Sorted Lists
[LeetCode](https://leetcode.com/problems/merge-k-sorted-lists)/[力扣](https://leetcode-cn.com/problems/merge-k-sorted-lists)


```cpp
ListNode* mergeKLists(vector<ListNode*>& lists) {
  priority_queue<ListNode*, vector<ListNode*>,compare> q;
  for(auto l : lists){
      if(l) q.push(l);
  }
  
  ListNode pre(0);
  ListNode* node = &pre;
  while(!q.empty()) {
      ListNode* top = q.top(); q.pop();
      node->next = top;
      node = node->next;
      if(top->next)q.push(top->next);
  }
  return pre.next;
}
// 最小堆
struct compare {
    bool operator()(const ListNode* l1, const ListNode* l2) {
        return l1->val > l2->val;
    }
};
```

```cpp
ListNode* mergeKLists(vector<ListNode*>& lists) {
    ListNode* list = nullptr;
    for(int i = 0; i < lists.size(); i++){
        list =  mergeTwoList(list, lists[i]);
    }
    return list;
}
ListNode* mergeTwoList(ListNode* l1, ListNode* l2) {
    if(l1 == nullptr) return l2;
    if(l2 == nullptr) return l1;
    
    if(l1->val < l2->val){
        l1->next = mergeTwoList(l1->next,l2);
        return l1;
    }else {
        l2->next = mergeTwoList(l1, l2->next);
        return l2;
    }
}
```

## Find K Pairs with Smallest Sums
[LeetCode](https://leetcode.com/problems/find-k-pairs-with-smallest-sums)/[力扣](https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums)

```cpp
struct compare {
    bool operator()(vector<int>& v1, vector<int>& v2) {
        return v1[0] + v1[1] < v2[0] + v2[1];
    }  
};
vector<vector<int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {
    priority_queue<vector<int>,vector<vector<int>>,compare> pq;
    
    for(int i = 0; i < nums1.size() && i < k; i++) {
        for(int j = 0; j < nums2.size() && j < k; j++) {
            pq.push({nums1[i],nums2[j]});
            if(pq.size() > k) pq.pop();
        }
    }
    vector<vector<int>> res;
    while(!pq.empty()) {
        res.push_back(pq.top());pq.pop();
    }
    return res;
}
```

## Kth Smallest Element in a Sorted Matrix
[LeetCode](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix)/[力扣](https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix)

```cpp
struct Node {
    int row;
    int col;
    int val;
    Node(int r, int c, int v) {
        row = r;col = c; val = v;
    }
};
struct compare {
    bool operator()(Node* n1, Node* n2) {
        return n1->val > n2->val;
    }  
};
int kthSmallest(vector<vector<int>>& matrix, int k) {
    int m = matrix.size(), n = matrix[0].size();
    int lo = matrix[0][0], hi = matrix.back().back();
    while (lo < hi) {
        int cnt = 0, j = n-1;
        int mid = lo + (hi-lo) / 2;
        for (int i = 0; i < matrix.size(); ++i) {
            while(j >= 0 && mid < matrix[i][j]) --j;
            cnt += j+1;
        }
        if (cnt < k) lo = mid + 1;
        else         hi = mid;
    }
    return lo;
}  
```

## Smallest Range Covering Elements from K Lists
[LeetCode](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists)/[力扣](https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists)


```cpp
vector<int> smallestRange(vector<vector<int>>& nums) {
  int curMin = INT_MAX, curMax = INT_MIN;
  
  priority_queue<VI, vector<VI>, Comp> pq; // min head
  for (auto& arr : nums) {
    pq.push({arr.begin(), arr.end()});
    curMin = min(curMin, arr[0]);
    curMax = max(curMax, arr[0]);
  }
  
  vector<int> range = {curMin, curMax};
  while (true) {
    auto p = pq.top(); pq.pop(); // top holds smallest value in the k-value collection, and try to replace it with next larger one
    if (++p[0] == p[1]) break; // an array is running out
    
    pq.push({p[0], p[1]});
    curMin = *pq.top()[0];
    curMax = max(curMax, *p[0]);
    if (curMax - curMin < range[1] - range[0]) range = {curMin, curMax};
  }
  
  return range;
}
  
typedef vector<vector<int>::iterator> VI; // (curItr, endItr)

struct Comp { 
  bool operator()(const VI& a, const VI& b) { return *a[0] > *b[0]; }
};
```