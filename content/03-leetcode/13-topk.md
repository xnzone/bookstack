---
title: "第K最大值"
date: 2022-08-18T00:43:42+08:00
tags: ["leetcode", "top k", "c++"]
image: /covers/leetcode.png
cover: false 
weight: 15 
---

## Kth Kth Largest Element in an Array
[LeetCode](https://leetcode.com/problems/kth-largest-element-in-an-array)/[力扣](https://leetcode-cn.com/problems/kth-largest-element-in-an-array)

用优先队列来存k个元素，然后遍历数组，如果比堆顶元素大，则将堆顶元素弹出，并将元素放入

{{< highlight cpp >}}
int findKthLargest(vector<int>& nums, int k) {
    priority_queue<int, vector<int>, greater<int>> min_heap;
    for(int i = 0; i < k; ++i) {
        min_heap.push(nums[i]);
    }
    
    for(int i = k; i < nums.size(); ++i) {
        int min = min_heap.top();
        if (nums[i] > min) {
            min_heap.pop();
            min_heap.push(nums[i]);
        }
    }
    return min_heap.top();
}
{{< /highlight  >}}

## Kth Smallest Element in a BST
[LeetCode](https://leetcode.com/problems/kth-smallest-element-in-a-bst)/[力扣](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst)

{{< highlight cpp >}}
int index = -1;
int res;
int kthSmallest(TreeNode* root, int k) {
    helper(root,k);
    
    return res;
    
}

void helper(TreeNode* root,int& k) {
    if(root == nullptr) return;
    helper(root->left,k);
    index++;
    if (index == k - 1) res = root->val;
    else helper(root->right,k);
}
{{< /highlight  >}}

## Top K Frequent Elements
[LeetCode](https://leetcode.com/problems/top-k-frequent-elements)/[力扣](https://leetcode-cn.com/problems/top-k-frequent-elements)

{{< highlight cpp >}}
struct compare{
    bool operator()(pair<int, int>& p1, pair<int,int>& p2) {
        return p1.second > p2.second;
    }
};
vector<int> topKFrequent(vector<int>& nums, int k) {
    map<int, int> map;
    for(int num : nums) {
        map[num]++;
    }
    
    priority_queue<pair<int,int>,vector<pair<int,int>>, greater<pair<int,int>>> q;
    auto it = map.begin();
    for(int i = 0; i < k; ++i) {
        q.push({it->second, it->first});
        it++;
    }
    while(it != map.end()){
        q.push({it->second, it->first});
        q.pop();
        it++;
    }
    
    vector<int> res;
    while(!q.empty()) {
        pair<int,int> p = q.top(); q.pop();
        res.push_back(p.second);
    }
    return res;
}
{{< /highlight  >}}

## Sort Characters By Frequency
[LeetCode](https://leetcode.com/problems/sort-characters-by-frequency)/[力扣](https://leetcode-cn.com/problems/sort-characters-by-frequency)

{{< highlight cpp >}}
struct compare {
    bool operator()(const pair<int,int> p1, const pair<int,int> p2) {
        return p1.first < p2.first;
    }  
};
string frequencySort(string s) {
    map<char,int> map;
    for(char c : s) map[c]++;
    
    priority_queue<pair<int,int>, vector<pair<int,int>>, compare> p;
    
    for(auto it = map.begin(); it != map.end(); ++it){
        p.push({it->second, it->first});
    }
    
    int index = 0;
    while(!p.empty()) {
        auto q = p.top(); p.pop();
        for(int i = 0; i < q.first; ++i) {
            s[index++] = q.second;
        }
    }
    
    return s;
}
{{< /highlight  >}}

## Course Schedule III
[LeetCode](https://leetcode.com/problems/course-schedule-iii)/[力扣](https://leetcode.com/problems/course-schedule-iii)

{{< highlight cpp >}}
static bool cmp(vector<int> & a,vector<int> & b){
    if(a[1] == b[1]){
        return a[0] < b[0];
    }
    return a[1] < b[1];
}

int scheduleCourse(vector<vector<int>>& courses) {
    int curr = 0;
    
    priority_queue<int,vector<int>,less<int>> pq;
    sort(courses.begin(),courses.end(),cmp);
    
    for(auto c : courses){
        if(c[0] + curr <= c[1]){
            curr += c[0];
            pq.push(c[0]);
        }else{
            if(!pq.empty() && pq.top() > c[0]){
                curr += c[0] - pq.top();
                pq.pop();
                pq.push(c[0]);
            }
        }
    }
    
    return pq.size();
}
{{< /highlight  >}}

## Find K Closest Elements
[LeetCode](https://leetcode.com/problems/find-k-closest-elements)/[力扣](https://leetcode-cn.com/problems/find-k-closest-elements)

{{< highlight cpp >}}
vector<int> findClosestElements(vector<int>& arr, int k, int x) {
  int n=arr.size();
  vector<int>res;
  //i have to find the k closeset elements 
  if(x <= arr[0]){
      //first k elements is the answer
    return vector<int>(arr.begin(),arr.begin()+k);
  }

  if(x>=arr[n-1]){
    return vector<int>(arr.begin()+n-k,arr.end());
  }

  int index=lower_bound(arr.begin(),arr.end(),x)-arr.begin();
  int low=max(0,index-k);
  int high=min(n-1,index+k-1);
  while(high-low+1>k){   
    if(x-arr[low] >   arr[high]-x)
      low++;
    else
      high--;
  }

  return vector<int>(arr.begin()+low,arr.begin()+high+1);
}
{{< /highlight  >}}

## Reorganize String
[LeetCode](https://leetcode.com/problems/reorganize-string)/[力扣](https://leetcode-cn.com/problems/reorganize-string)

{{< highlight cpp >}}
string reorganizeString(string S) {
    int n = S.size();
    map<char, int> map;
    char maxC;
    int maxN = 0;
    
    for(char c:S) {
        map[c]++;
        if(map[c] > maxN) {
            maxC = c;
            maxN = map[c];
        }
    }
    
    if ((n+1)/ maxN < 2) return "";
    
    int gap = (n + 1) / maxN;
    for(int i = 0; i < n; i = i + gap) {
        S[i] = maxC;
        map[maxC]--;
    }
    for(int i = 0; i < gap; i++) {
        int j = i + 1;
        for(auto it = map.begin(); it != map.end(); ++it) {
            if(it->first == maxC){
                continue;
            }
            int count = it->second;
            while(count > 0 && j < n) {
                S[j] = it->first;
                map[it->first]--;
                j += gap;
                count--;
            }
            if (j >= n) break;
        }
    }
    
    return S;
}
{{< /highlight  >}}

## Maximum Frequency Stack
[LeetCode](https://leetcode.com/problems/maximum-frequency-stack)/[力扣](https://leetcode-cn.com/problems/maximum-frequency-stack)

{{< highlight cpp >}}
unordered_map<int,int> freq;
unordered_map<int,stack<int>> freq_map;
int mfreq=0;
FreqStack() {
    
}

void push(int x) {
    mfreq=max(mfreq,++freq[x]);
    freq_map[freq[x]].push(x);
}

int pop() {
    auto temp=freq_map[mfreq].top();
    freq_map[mfreq].pop();
    if(freq_map[freq[temp]--].empty())
        --mfreq;
    return temp;
}
{{< /highlight  >}}

## K Closest Points to Origin
[LeetCode](https://leetcode.com/problems/k-closest-points-to-origin)/[力扣](https://leetcode-cn.com/problems/k-closest-points-to-origin)

{{< highlight cpp >}}
struct compare {
    bool operator()(vector<int> v1, vector<int> v2) {
        return v1[0] * v1[0] + v1[1] * v1[1] < v2[0] * v2[0] + v2[1] * v2[1];
    }  
};
vector<vector<int>> kClosest(vector<vector<int>>& points, int K) {
    vector<vector<int>> res;
    priority_queue<vector<int>, vector<vector<int>>, compare> pq;
    for(int i = 0; i < K; ++i) {
        pq.push(points[i]);
    }
    
    for(int i = K; i < points.size(); i++) {
        pq.push(points[i]);
        pq.pop();
    }
    
    while(!pq.empty()) {
        res.push_back(pq.top());
        pq.pop();
    }
    return res;
}
{{< /highlight  >}}

