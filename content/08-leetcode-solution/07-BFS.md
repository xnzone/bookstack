---
title: "广度优先搜索"
date: 2022-06-29T15:43:42+08:00
tags: ["leetcode", "bfs", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 807
---


## Binary Tree Level Order Traversal
[LeetCode](https://leetcode.com/problems/binary-tree-level-order-traversal)/[力扣](https://leetcode-cn.com/problems/binary-tree-level-order-traversal)

递归

```c++
void levelOrderHelper(vector<vector<int>> &res, int level, TreeNode *root) {
    if (root == NULL) return;
    if (res.size() == level) res.push_back(vector<int>());
    res[level].push_back(root->val);
    levelOrderHelper(res, level + 1, root->left);
    levelOrderHelper(res, level + 1, root->right);
}

vector<vector<int>> levelOrder(TreeNode* root) {
    vector<vector<int>> res;
    levelOrderHelper(res, 0, root);
    return res;
}
```

非递归，层序遍历，用queue保存当前层

```c++
vector<vector<int>> levelOrder(TreeNode* root) {
    vector<vector<int>> res;
    if(root == nullptr) return res;
    
    queue<TreeNode*> q;
    q.push(root);
    int count = 1; // 当前层的个数
    vector<int> cur; // 当前层的vector
    int next = 0; // 下一层的个数
    while(!q.empty()){
        TreeNode* temp = q.front();
        cur.push_back(temp->val);
        if(temp->left != nullptr){
            next++;
            q.push(temp->left);
        }
        if(temp->right != nullptr){
            next++;
            q.push(temp->right);
        }
        q.pop();
        count--;
        if(count == 0){
            count = next;
            next = 0;
            res.push_back(cur);
            cur.clear();// 清空vector
        }
    }
    return res;
}
```

## Binary Tree Zigzag Level Order Traversal
[LeetCode](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal)/[力扣](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal)


- 层序遍历，用一个标识表示是否从左开始
- 如果是从左开始，直接push_back
- 如果是从右开始，直接insert到第一个

```c++
 vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
    vector<vector<int>> res;
    if (root == nullptr) return res;
    
    queue<TreeNode*> q;
    q.push(root);
    int count = 1;
    vector<int> cur;
    int next = 0;
    bool left = true;
    while(!q.empty()){
        TreeNode* temp = q.front();
        if (left) 
            cur.push_back(temp->val);
        else
            cur.insert(cur.begin(),temp->val);
        if(temp->left != nullptr){
            next++;
            q.push(temp->left);
        }
        if(temp->right != nullptr){
            next++;
            q.push(temp->right);
        }
        q.pop();
        count--;
        if(count == 0){
            // if(!left) reverse(cur.begin(),cur.end());
            left = !left;
            count = next;
            next = 0;
            res.push_back(cur);
            cur.clear();
        }
    }
    return res;
}
```

## Binary Tree Level Order Traversal II
[LeetCode](https://leetcode.com/problems/binary-tree-level-order-traversal-ii)/[力扣](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii)

层序遍历，最后直接翻转vector就可以了

```c++
vector<vector<int>> levelOrderBottom(TreeNode* root) {
    vector<vector<int>> res;
    if(root == nullptr) return res;
    
    queue<TreeNode*> q;
    vector<int> cur;
    q.push(root);
    int count = 1;
    int next = 0;
    while(!q.empty()) {
        TreeNode* temp = q.front();q.pop();
        cur.push_back(temp->val);
        if(temp->left){next++;q.push(temp->left);}
        if(temp->right){next++;q.push(temp->right);}
        count--;
        if(count == 0){
            count = next;
            next = 0;
            res.push_back(cur);
            cur.clear();
        }
    }
    reverse(res.begin(), res.end());
    return res;
}
```

## Minimum Depth of Binary Tree
[LeetCode](https://leetcode.com/problems/minimum-depth-of-binary-tree)/[力扣](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree)

递归求解，直接求左子树的高度，然后再求右子树的高度，两个取小的

```c++
int minDepth(TreeNode* root) {
    if (root == nullptr) return 0;
    if (root->left == nullptr && root->right == nullptr) return 1;
    if (!root->left) return minDepth(root->right) + 1;
    if (!root->right) return minDepth(root->left) + 1;
    
    return min(minDepth(root->left),minDepth(root->right)) + 1;
}
```

## Populating Next Right Pointers in Each Node
[LeetCode](https://leetcode.com/problems/populating-next-right-pointers-in-each-node)/[力扣](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node)


层序遍历

递归

```c++
Node* connect(Node* root) {
    if(root == nullptr) return root;
    root->next = nullptr;
    connectHelper(root);
    return root;
}

void connectHelper(Node* root) {
    if(root == nullptr) return;
    
    if (root->left && root->right) {
        root->left->next = root->right;
    }
    
    if(root->next && root->right){
        root->right->next = root->next->left;
    }
    connectHelper(root->left);
    connectHelper(root->right);
}
```

非递归

```c++
Node* connect(Node* root) {
    if (!root) return root;
    Node* level = root;
    
    while (level) {
        Node* list = level;
        while (list) {
            if (list->left) {
                list->left->next = list->right;
            }
            if (list->next && list->right) {
                list->right->next = list->next->left;
            }
            list = list->next;
        }
        level = level->left;
    }
    
    return root;
}
```

## Populating Next Right Pointers in Each Node II
[LeetCode](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii)/[力扣](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii)


递归 层序遍历做成递归形式

```c++
void solve(Node* node, int lvl, vector<Node*> &track){
    if(node==NULL){
        return;
    }
    
    solve(node->right,lvl+1,track);
    solve(node->left,lvl+1,track);
    
    if(track[lvl]==NULL){
        track[lvl]=node;
    }else{
        node->next=track[lvl];
        track[lvl]=node;
    }
    
}

Node* connect(Node* root) {
    vector<Node*> track(100000,NULL);
    
    solve(root,0,track);
    return root;
}
```

层序遍历，记录当前层的node，然后组成一个链表

```c++
Node* connect(Node* root) {
    if(root == nullptr) return root;
    queue<Node*> q;
    q.push(root);
    while(!q.empty()) {
        int count = q.size();
        Node* cur = q.front();q.pop();
        if(cur->left != nullptr)q.push(cur->left);
        if(cur->right != nullptr) q.push(cur->right);
        while(count > 0){
            if(count == 1) break;
            Node* tmp = q.front();q.pop();
            cur->next = tmp;
            cur = tmp;
            if(tmp->left != nullptr)q.push(tmp->left);
            if(tmp->right != nullptr) q.push(tmp->right);
            count--;
        }
        cur->next = nullptr;
    }
    return root;
}
```


## Binary Tree Right Side View
[LeetCode](https://leetcode.com/problems/binary-tree-right-side-view)/[力扣](https://leetcode-cn.com/problems/binary-tree-right-side-view)

层序遍历，压入最右边那个

```c++
vector<int> rightSideView(TreeNode* root) {
    vector<int> ans;
    if(!root) return ans;
    
    queue<TreeNode*> q;
    
    q.push(root);
    
    while(!q.empty()) {
        int size = q.size();
        int value;
        while(size--) {
            TreeNode* node = q.front();
            value = node->val;
            q.pop();
            if(node->left) q.push(node->left);
            if(node->right) q.push(node->right);
        }
        ans.push_back(value);
    }
    
    return ans;
}
```

## Number of Islands
[LeetCode](https://leetcode.com/problems/number-of-islands)/[力扣](https://leetcode-cn.com/problems/number-of-islands)

用一个标记表示有没有访问过

```c++
int numIslands(vector<vector<char>>& grid) {
    int res = 0;
    int m = grid.size();
    int n = m == 0 ? 0 : grid[0].size();
    if (m == 0 || n == 0) {
        return res;
    }
    
    bool visited[m][n];
    for(int i = 0; i < m; ++i){
        for(int j = 0; j < n; ++j){
            visited[i][j] = false;
        }
    }
    
    queue<int> q;
    int dx[] = {-1, 0, 1, 0}, dy[] = {0, 1, 0, -1};
    int x = 0, y = 0, xx = 0, yy = 0;
    
    for(int i = 0; i < m; ++i) {
        for(int j = 0; j < n; ++j) {
            if(grid[i][j] == '1' && !visited[i][j]){
                q.push(i);
                q.push(j);
                visited[i][j] = true;
                res++;
                while(!q.empty()) {
                    x = q.front(); q.pop();
                    y = q.front(); q.pop();
                    for(int k = 0; k < 4; k++){
                        xx = x + dx[k];
                        yy = y + dy[k];
                        if(xx < 0 || xx >= m || yy < 0 || yy >= n) continue;
                        if(grid[xx][yy] == '1' && !visited[xx][yy]){
                            q.push(xx);
                            q.push(yy);
                            visited[xx][yy] = true;
                        }
                    }
                }
            }
        }
    }
    return res;
}
```

## Average of Levels in Binary Tree
[LeetCode](https://leetcode.com/problems/average-of-levels-in-binary-tree)/[力扣](https://leetcode-cn.com/problems/average-of-levels-in-binary-tree)

记录每一层的和，求平均值

```c++
vector<double> averageOfLevels(TreeNode* root) {
    vector<double> res;
    if(root == nullptr) return res;
    queue<TreeNode*> q;
    q.push(root);
    
    while(!q.empty()) {
        int n = q.size();
        double sum = 0;
        for(int i = 0; i < n; ++i) {
            TreeNode* temp = q.front(); q.pop();
            sum += temp->val;
            if(temp->left != nullptr) q.push(temp->left);
            if(temp->right != nullptr) q.push(temp->right);
        }
        res.push_back(sum / n);
    }
    
    return res;
    
}
```

## All Nodes Distance K in Binary Tree
[LeetCode](https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree)/[力扣](https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree)

- 找到目标节点相对于当前节点的深度，然后分四种情况判断
- 如果当前节点就是目标节点，就再往下寻找k层即可
- 如果目标节点在当前节点的左分支，且深度为L，则只需要在当前节点的右边节点去寻找`K - L -1`深度即可
- 如果目标节点在当前节点的右分支，处理情况类似于左分支
- 如果当目标节点不在当前节点子分支下面，直接返回

```c++
vector<int> res;
TreeNode* target;
int K;
vector<int> distanceK(TreeNode* root, TreeNode* target, int K) {
    this->target = target;
    this->K = K;
    dfs(root);
    return this->res;
}
// dfs 
int dfs(TreeNode* root) {
    if(root == nullptr) return -1;
    if(root == this->target){
        subtree_helper(root, 0);
        return 1;
    } 
    
    int l = dfs(root->left), r = dfs(root->right);
    if (l != -1) {
        if(l == K) this->res.push_back(root->val);
        subtree_helper(root->right, l + 1);
        return l + 1;
    } else if (r != -1) {
        if(r == K) this->res.push_back(root->val);
        subtree_helper(root->left, r + 1);
        return r + 1;
    } else {
        return -1;
    }
}
// 从节点出发，查找深度为 l - k - 1 的节点
void subtree_helper(TreeNode* root, int dist) {
    if(root == nullptr) return;
    if(dist == K){
        this->res.push_back(root->val);
    } else {
        subtree_helper(root->left, dist + 1);
        subtree_helper(root->right, dist + 1);
    }
}
```