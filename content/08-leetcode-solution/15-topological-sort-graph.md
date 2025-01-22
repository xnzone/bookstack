---
title: "拓扑排序和图论"
date: 2022-09-10T00:43:42+08:00
tags: ["leetcode", "topological sort and graph", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 815
---

## Course Schedule
[LeetCode](https://leetcode.com/problems/course-schedule)/[力扣](https://leetcode-cn.com/problems/course-schedule)

```cpp
bool res = true;
bool canFinish(int n, vector<vector<int>>& preps) {
    vector<vector<int>> graph(n, vector<int>());
    for(int i = 0; i < preps.size(); i++) {
        graph[preps[i][1]].push_back(preps[i][0]);
    }
    vector<int> visited(n,0);
    
    for(int i = 0; i < n; i++) {
        if(visited[i] == 0) 
            helper(graph,i,visited);
    }
    return res;
}

void helper(vector<vector<int> >&graph, int idx, vector<int>&visited) {
    if(!res) return;
    visited[idx] = 2;
    for(int i = 0; i < graph[idx].size(); i++) {
        if(visited[graph[idx][i]] == 0) {
            helper(graph,graph[idx][i],visited);
        }else if(visited[graph[idx][i]] == 2) {
            res = false;
            return;
        }
    }
    visited[idx] = 1;
}
```

## Course Schedule II
[LeetCode](https://leetcode.com/problems/course-schedule-ii)/[力扣](https://leetcode-cn.com/problems/course-schedule-ii)

```cpp
bool res = true;
bool canFinish(int n, vector<vector<int>>& preps) {
    vector<vector<int>> graph(n, vector<int>());
    for(int i = 0; i < preps.size(); i++) {
        graph[preps[i][1]].push_back(preps[i][0]);
    }
    vector<int> visited(n,0);
    
    for(int i = 0; i < n; i++) {
        if(visited[i] == 0) 
            helper(graph,i,visited);
    }
    return res;
}

void helper(vector<vector<int> >&graph, int idx, vector<int>&visited) {
    if(!res) return;
    visited[idx] = 2;
    for(int i = 0; i < graph[idx].size(); i++) {
        if(visited[graph[idx][i]] == 0) {
            helper(graph,graph[idx][i],visited);
        }else if(visited[graph[idx][i]] == 2) {
            res = false;
            return;
        }
    }
    visited[idx] = 1;
}
```

## Minimum Height Trees
[LeetCode](https://leetcode.com/problems/minimum-height-trees)/[力扣](https://leetcode-cn.com/problems/minimum-height-trees)

```cpp
vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
            vector<int> leaves;
    vector<vector<int>> graph(n,vector<int>());
    
    for(auto ed : edges)
    {
        graph[ed[0]].push_back(ed[1]);
        graph[ed[1]].push_back(ed[0]);
    }
    //Prepare the list of leaves to remove
    for(int i=0;i<n;i++)
    {
        if(graph[i].size()==1)
            leaves.push_back(i);
    }
    while(n>2)
    {
        n-=leaves.size();
        vector<int> new_leaves;
        for(auto leaf : leaves)
        {
            int adj = graph[leaf][0];
            vector<int>::iterator itr = remove(graph[adj].begin(),graph[adj].end(),leaf);
            graph[adj].erase(itr,graph[adj].end());
            if(graph[adj].size()==1)
                new_leaves.push_back(adj);
        }
        leaves = new_leaves;
    }
    if(leaves.size()==0)
        return {0};
    return leaves;
}
```


## Clone Graph
[LeetCode](https://leetcode.com/problems/clone-graph)/[力扣](https://leetcode-cn.com/problems/clone-graph)

```cpp
unordered_map<Node*, Node*> m;
Node* cloneGraph(Node* node) {
    if(node == nullptr) return node;
    Node* root;
    if(m.find(node) != m.end()) {
        root = m[node];
        return root;
    }
    root = new Node(node->val);
    m[node] = root;
    for(auto x:node->neighbors) {
        Node* r = cloneGraph(x);
        if(r) root->neighbors.push_back(r);
    }
    return root;
}
```

## Pacific Atlantic Water Flow
[LeetCode](https://leetcode.com/problems/pacific-atlantic-water-flow)/[力扣](https://leetcode-cn.com/problems/pacific-atlantic-water-flow)

```cpp
bool res = true;
bool canFinish(int n, vector<vector<int>>& preps) {
    vector<vector<int>> graph(n, vector<int>());
    for(int i = 0; i < preps.size(); i++) {
        graph[preps[i][1]].push_back(preps[i][0]);
    }
    vector<int> visited(n,0);
    
    for(int i = 0; i < n; i++) {
        if(visited[i] == 0) 
            helper(graph,i,visited);
    }
    return res;
}

void helper(vector<vector<int> >&graph, int idx, vector<int>&visited) {
    if(!res) return;
    visited[idx] = 2;
    for(int i = 0; i < graph[idx].size(); i++) {
        if(visited[graph[idx][i]] == 0) {
            helper(graph,graph[idx][i],visited);
        }else if(visited[graph[idx][i]] == 2) {
            res = false;
            return;
        }
    }
    visited[idx] = 1;
}
```
