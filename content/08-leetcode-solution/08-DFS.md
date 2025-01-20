---
title: "深度优先搜索"
date: 2022-06-30T15:43:42+08:00
tags: ["leetcode", "dfs", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 808
---


## Validate Binary Search Tree
[LeetCode](https://leetcode.com/problems/validate-binary-search-tree)/[力扣](https://leetcode-cn.com/problems/validate-binary-search-tree)

递归

```c++
bool isValidBST(TreeNode* root) {
    double lower = DBL_MIN;
    double upper = DBL_MAX;
    return helper(root, lower, upper);
}

bool helper(TreeNode* root, double lower, double upper) {
    if(root == nullptr) return true;
    
    int val = root->val;
    if(lower != DBL_MIN && val <= lower) return false;
    if(upper !=DBL_MAX && val >= upper) return false;
    
    if(!helper(root->left, lower, val)) return false;
    if(!helper(root->right, val, upper)) return false;
    return true;
}
```

## Same Tree
[LeetCode](https://leetcode.com/problems/same-tree)/[力扣](https://leetcode-cn.com/problems/same-tree)

递归

```c++
bool isSameTree(TreeNode* p, TreeNode* q) {
    if(p == nullptr && q == nullptr) return true;
    
    if(p == nullptr || q == nullptr) return false;
    
    if(p->val != q->val) return false;
    
    return isSameTree(p->left,q->left) && isSameTree(p->right,q->right);
}
```

## Maximum Depth of Binary Tree
[LeetCode](https://leetcode.com/problems/maximum-depth-of-binary-tree)/[力扣](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree)

```c++
int maxDepth(TreeNode* root) {
    if(root ==  nullptr)
        return 0;
    int nleft = 0;
    int nright = 0;
    TreeNode *pRoot = root;
    nleft += maxDepth(pRoot->left);
    nright += maxDepth(pRoot->right);
    return (nleft > nright ? nleft : nright) + 1 ;
}
```

## Construct Binary Tree from Preorder and Inorder Traversal
[LeetCode](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal)/[力扣](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal)

就是重建树。要了解到前序遍历（中左右）和中序遍历（左中右）的特点

```c++
TreeNode* buildTree(vector<int>& pre, vector<int>& in) {
    return helper(pre, 0, pre.size() - 1, in, 0, in.size() - 1);
}

TreeNode* helper(vector<int>& pre, int pleft, int pright, vector<int>& ino, int ileft, int iright) {
    if(pleft > pright || ileft > iright) return nullptr;
    TreeNode* p = new TreeNode(pre[pleft]);
    if(pleft == pright) return p;
    int index = ileft;
    for(; index <= iright; index++) {
        if(ino[index] == pre[pleft])break;
    }
    int n = index - ileft;
    p->left = helper(pre, pleft + 1, pleft + n, ino, ileft, ileft + n - 1);
    p->right = helper(pre, pleft + n + 1, pright, ino, ileft + n + 1, iright);

    return p;
}
```

## Path Sum
[LeetCode](https://leetcode.com/problems/path-sum)/[力扣](https://leetcode-cn.com/problems/path-sum)

```c++
bool hasPathSum(TreeNode* root, int sum) {
    if(root == nullptr) return false;
    return helper(root, sum, 0);
}

bool helper(TreeNode* root, int sum, int temp) {
    if(root == nullptr) return false;
    temp += root->val;
    if(temp == sum && !root->left && !root->right) return true;
    return helper(root->left, sum, temp) || helper(root->right, sum, temp);
}
```

## Path Sum II
[LeetCode](https://leetcode.com/problems/path-sum-ii)/[力扣](https://leetcode-cn.com/problems/path-sum-ii)

```c++
vector<vector<int>> pathSum(TreeNode* root, int sum) {
    vector<int> tv;
    vector<vector<int>> res;
    helper(root, sum, 0, 0, tv, res);
    return res;
}
void helper(TreeNode* root, int sum, int temp,int len, vector<int>& tv, vector<vector<int>>& res) {
    if(root == nullptr) return;
    temp += root->val;
    tv.push_back(root->val);
    if(temp == sum && !root->left && !root->right){
        res.push_back(tv);
    }
    helper(root->left, sum, temp, len + 1, tv, res);
    while(tv.size() > len + 1){
        tv.pop_back();
    }
    helper(root->right, sum, temp, len + 1, tv, res);
}
```

## Binary Tree Maximum Path Sum
[LeetCode](https://leetcode.com/problems/binary-tree-maximum-path-sum)/[力扣](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum)

```c++
int res = INT_MIN;
int maxPathSum(TreeNode* root) {
    helper(root);
    return res;
}

int helper(TreeNode* root) {
    if(root == nullptr) return 0;
    
    int left = helper(root->left);
    int right = helper(root->right);
    
    res = max(res, root->val);
    res = max(res, root->val + left);
    res = max(res, root->val + right);
    res = max(res, root->val + left + right);
    
    return max(root->val, max(root->val + left,root->val + right));
}
```

## Implement Trie (Prefix Tree)
[LeetCode](https://leetcode.com/problems/implement-trie-prefix-tree)/[力扣](https://leetcode-cn.com/problems/implement-trie-prefix-tree)

```c++
class Trie {
public:
    struct newTrie
    {
        char data;
        unordered_map<char,newTrie*> child;
        bool last;
        newTrie(char c)
        {
            data=c;
            last=false;
        }
    };
    newTrie*root;
    /** Initialize your data structure here. */
    Trie()
    {   
        root=new newTrie('\0');
    }
    
    /** Inserts a word into the trie. */
    void insert(string word) {
        
        newTrie *temp=root;
        for(int i=0;i<word.length();i++)
        {
            if(temp->child.find(word[i])!=temp->child.end())
            {
                temp=temp->child[word[i]];
                if(i==word.length()-1)
                    temp->last=true;
            }
            else
            {
                temp->child[word[i]]=new newTrie(word[i]);
                temp=temp->child[word[i]];
                if(i==word.length()-1)
                    temp->last=true;
            }
        }
    }
    
    //Returns if the word is in the trie. 
	//which is only possible when last character of word is in last node of one of branches of prefix tree
    bool search(string word) {
        newTrie *temp=root;
        for(int i=0;i<word.length();i++)
        {
            if(temp->child.find(word[i])!=temp->child.end())
            {
               if(i!=word.length()-1)    
                    temp=temp->child[word[i]];
               else if(temp->child[word[i]]->last==true)
                   return true;
            }
            else
                return false;
        }
        return false;
    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    bool startsWith(string prefix) {
        newTrie *temp=root;
        for(int i=0;i<prefix.length();i++)
        {
            if(temp->child.find(prefix[i])!=temp->child.end())
            {
                temp=temp->child[prefix[i]];
            }
            else
                return false;
        }
        return true;
    }
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */
```

## Word Search II
[LeetCode](https://leetcode.com/problems/word-search-ii)/[力扣](https://leetcode-cn.com/problems/word-search-ii)

```c++
vector<string> findWords(vector<vector<char>>& board, vector<string>& words) {
    vector<string> res;
    for(int i = 0; i < words.size();i++){
        if(helper(board,words[i]))res.push_back(words[i]);
    }
    return res;
}
bool helper(vector<vector<char>>& board, string word) {
    for(int i = 0; i < board.size(); i++) {
        for(int j = 0; j < board[0].size(); j++){
            if(existedHelper(i,j,board,word,0)) return true;
        }
    }
    return false;
    
}

bool existedHelper(int x, int y, vector<vector<char>>& board, string& word, int len) {
    if(x < 0 || y < 0 || x >= board.size() || y >= board[0].size() || board[x][y] != word[len]) return false;
    
    if(len == word.size() - 1) return true;
    
    char t = board[x][y];
    board[x][y] = '0';
    bool exist = existedHelper(x - 1, y, board, word, len + 1) ||
        existedHelper(x, y - 1, board, word, len + 1) ||
        existedHelper(x + 1, y, board, word, len + 1) ||
        existedHelper(x, y + 1, board, word, len + 1);
    
    board[x][y] = t;
    return exist;
}
```

## Invert Binary Tree
[LeetCode](https://leetcode.com/problems/invert-binary-tree)/[力扣](https://leetcode-cn.com/problems/invert-binary-tree)

```c++
TreeNode* invertTree(TreeNode* root) {
    if(root == nullptr) return root;
    swap(root->left, root->right);
    invertTree(root->left);
    invertTree(root->right);
    return root;
}
```

## Kth Smallest Element in a BST
[LeetCode](https://leetcode.com/problems/kth-smallest-element-in-a-bst)/[力扣](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst)

```c++
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
```

## Lowest Common Ancestor of a Binary Search Tree
[LeetCode](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree)/[力扣](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree)

```c++
TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
    if(root == nullptr) return root;
    int maxVal = max(p->val,q->val);
    int minVal = min(p->val,q->val);
    if(maxVal >= root->val && minVal <= root->val) return root;
    else if(root->val > maxVal) return lowestCommonAncestor(root->left,p,q);
    else return lowestCommonAncestor(root->right,p,q);
}
```

## Lowest Common Ancestor of a Binary Tree
[LeetCode](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree)/[力扣](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree)

```c++
TreeNode* res = nullptr;
TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
    helper(root, p, q);
    return res;
}

bool helper(TreeNode* root, TreeNode* p, TreeNode* q) {
    if(root == nullptr) return false;
    
    int left = helper(root->left,p,q) ? 1 : 0;
    int right = helper(root->right, p,q) ? 1 : 0;
    int mid = (root == p || root == q) ? 1 : 0;
    
    if(left + right  + mid >= 2){
        res = root;
    }
    
    return left + mid + right > 0 ;
}
```

## Serialize and Deserialize Binary Tree
[LeetCode](https://leetcode.com/problems/serialize-and-deserialize-binary-tree)/[力扣](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree)

```c++
// Encodes a tree to a single string.
string serialize(TreeNode* root) {
    stringstream ss;
    serializeHelper(ss, root);
    return ss.str();
}
void serializeHelper(stringstream& ss, TreeNode* cur) {
    if (!cur) {
        ss << "# ";
        return;
    }
    ss << cur->val << " ";
    serializeHelper(ss, cur->left);
    serializeHelper(ss, cur->right);
}

// Decodes your encoded data to tree.
TreeNode* deserialize(string data) {
    stringstream ss(data);
    TreeNode* root = nullptr;
    deserializeHelper(ss, root);
    return root;
}
void deserializeHelper(stringstream& ss, TreeNode* &cur) {
    string node;
    ss >> node;
    if (node == "" || node == "#") {
        cur = nullptr;
        return;
    }
    stringstream sss(node);
    int data;
    sss >> data;
    cur = new TreeNode();
    cur->val = data;
    cur->left = cur->right = nullptr;
    deserializeHelper(ss, cur->left);
    deserializeHelper(ss, cur->right);
}
```

## Path Sum III
[LeetCode](https://leetcode.com/problems/path-sum-iii)/[力扣](https://leetcode-cn.com/problems/path-sum-iii)

```c++
int pathSum(TreeNode* root, int sum) {
    if(root == nullptr) return 0;
    return helper(root, sum) + pathSum(root->left, sum) + pathSum(root->right, sum);
}

int helper(TreeNode* root, int sum) {
    if(root == nullptr) return 0;
    
    int left = helper(root->left, sum - root->val);
    int right = helper(root->right, sum - root->val);
    
    return sum == root->val ? left + right + 1 : left + right;
}
```

## Diameter of Binary Tree
[LeetCode](https://leetcode.com/problems/diameter-of-binary-tree)/[力扣](https://leetcode-cn.com/problems/diameter-of-binary-tree)

```c++
int diameterOfBinaryTree(TreeNode* root) {
    depth = 1;
    helper(root);
    return depth - 1;
}
int depth = 0;
int helper(TreeNode* root){
    if(root == nullptr) return 0;
    int left = helper(root->left);
    int right = helper(root->right);
    depth = max(depth, left + right + 1);
    return max(left, right) + 1;
}
```

## Subtree of Another Tree
[LeetCode](https://leetcode.com/problems/subtree-of-another-tree)/[力扣](https://leetcode-cn.com/problems/subtree-of-another-tree)

```c++
bool isSubtree(TreeNode* s, TreeNode* t) {
    if(s == nullptr && t == nullptr) return true;
    if(s == nullptr || t == nullptr) return false;
    return helper(s,t) || isSubtree(s->left, t) || isSubtree(s->right, t);
}

bool helper(TreeNode* root, TreeNode* t) {
    if(root == nullptr && t == nullptr) return true;
    
    if(root == nullptr || t == nullptr) return false;
    
    if(root->val != t->val) return false;
    
    return helper(root->left,t->left) && helper(root->right, t->right);
}
```

## Merge Two Binary Trees
[LeetCode](https://leetcode.com/problems/merge-two-binary-trees)/[力扣](https://leetcode-cn.com/problems/merge-two-binary-trees)

```c++
TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
    if(t1 == nullptr) return t2;
    if(t2 == nullptr) return t1;
    t1->val += t2->val;
    t1->left = mergeTrees(t1->left,t2->left);
    t1->right = mergeTrees(t1->right, t2->right);
    return t1;
}
```

## Maximum Binary Tree
[LeetCode](https://leetcode.com/problems/maximum-binary-tree)/[力扣](https://leetcode-cn.com/problems/maximum-binary-tree)

```c++
TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
    return helper(nums, 0, nums.size() - 1);
}
TreeNode* helper(vector<int>& nums, int left, int right) {
    if(left > right) return nullptr;
    int mindex = left;
    for(int i = left; i <= right; i++) {
        if(nums[i] > nums[mindex])mindex = i;
    }
    TreeNode* p = new TreeNode(nums[mindex]);
    p->left = helper(nums, left, mindex - 1);
    p->right = helper(nums, mindex + 1, right);
    return p;
}
```

## Maximum Width of Binary Tree
[LeetCode](https://leetcode.com/problems/maximum-width-of-binary-tree)/[力扣](https://leetcode-cn.com/problems/maximum-width-of-binary-tree)

```c++
int widthOfBinaryTree(TreeNode* root) {
    if(root == nullptr) return 0;
    queue<TreeNode*> q;
    q.push(root);
    int maxVal = 1;
    while(!q.empty()) {
        int n = q.size();
        maxVal = max(n, maxVal);
        for(int i = 0; i < n; i++) {
            TreeNode* t = q.front(); q.pop();
            if(t == nullptr) continue;
            q.push(t->left);
            q.push(t->right);
        }
    }
    return maxVal;
}
```