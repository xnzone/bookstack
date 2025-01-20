---
title: "链表翻转"
date: 2022-06-28T15:43:42+08:00
tags: ["leetcode", "reverse list", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 806
---


## Swap Nodes in Pairs
[LeetCode](https://leetcode.com/problems/swap-nodes-in-pairs)/[力扣](https://leetcode-cn.com/problems/swap-nodes-in-pairs)

- 递归
- 交换之后，直接交换下一个节点

```c++
ListNode* swapPairs(ListNode* head) {
    if(head && head->next) {
        swap(head->val, head->next->val);
        swapPair(head->next->next);
    }
    return head;
}
```

- 非递归，两两交换
- 看作先删除，然后再插入

```c++
ListNode* swapPairs(ListNode* head) {
    if(head == nullptr) return head;
    ListNode p(0);
    ListNode *node = &p;
    node->next = head;
    ListNode *cur = head, *pre = node;
    while(cur != nullptr && cur->next != nullptr) {
        ListNode *t = cur->next;
        cur->next = t->next;
        pre->next = t;
        t->next = cur;
        pre = cur;
        cur = cur->next;
    }
    return node->next;
}
```

## Reverse Nodes in k-Group
[LeetCode](https://leetcode.com/problems/reverse-nodes-in-k-group)/[力扣](https://leetcode-cn.com/problems/reverse-nodes-in-k-group)

- 中间链表反转
- 记住前后面指针，然后拼接三个链表

```c++
ListNode* reverseKGroup(ListNode* head, int k) {
    ListNode po(0);
    ListNode* p = &po;
    ListNode* prev = p;
    ListNode* cur = head;
    ListNode* cend = nextNode(head, k);
    ListNode* next = cend->next;
    
    while(cur != nullptr) {
        cend->next = nullptr;
        ListNode* rcur = reverse(cur);
        prev->next = rcur;
        prev = cur;
        cur = next;
        cend = nextNode(next, k);
        if(cend == nullptr) {
            prev->next = next;
            next = nullptr;
            break;
        }else{
            next = cend->next;
        }
    }
    
    return p->next;;
}

ListNode* nextNode(ListNode* node, int k) {
    for(int i = 0; i < k - 1; i++){
        if(node == nullptr) return node;
        node = node->next;
    }
    return node;
}

ListNode* reverse(ListNode* head){
    if(head == nullptr) return head;
    ListNode* prev = nullptr;
    ListNode* cur = head;
    ListNode* next = head->next;
    
    while(cur != nullptr) {
        cur->next = prev;
        prev = cur;
        cur = next;
        if(next != nullptr)next = next->next;
    }
    return prev;
}
```

## Rotate List
[LeetCode](https://leetcode.com/problems/rotate-list)/[力扣](https://leetcode-cn.com/problems/rotate-list)

记录链表总长度，根据总长度与K的余数，再看从哪里开始断开放到队列首部

```c++
ListNode* rotateRight(ListNode* head, int k) {
    if(head == nullptr || k == 0) return head;
    int n = 0;
    ListNode* phead = head;
    ListNode* pend = head;
    while(pend->next != nullptr){
        pend = pend->next;
        n++;
    }
    k = (n + 1) -  k % (n + 1);
    for(int i = 0; i < k - 1; i++){
        phead = phead->next;
    }
    pend->next = head;
    ListNode* res = phead->next;
    phead->next = nullptr;
    return res;
}
```

## Reverse Linked List II
[LeetCode](https://leetcode.com/problems/reverse-linked-list-ii)/[力扣](https://leetcode-cn.com/problems/reverse-linked-list-ii)

参考链表反转，将prev替换成nullptr就可以了

```c++
 ListNode* reverseBetween(ListNode* head, int m, int n) {
    if(head == nullptr) return head;
    ListNode p(0);
    ListNode* node = &p;
    node->next = head;
    ListNode* begin = node;
    ListNode* end = node;
    
    for(int i = 0; i < n; i++) {
        if(i == m - 1) {
            begin = end;
        }
        end = end->next;
    }
    begin->next = reverse(begin->next, end->next);
    return node->next;
    
}

ListNode* reverse(ListNode* begin, ListNode* end){
    if(begin == end) return begin;
    ListNode* prev = end;
    ListNode* cur = begin;
    while(cur != end){
        ListNode* next = cur->next;
        cur->next = prev;
        prev = cur;
        cur = next;
    }
    return prev;
}
```

## Reverse Linked List
[LeetCode](https://leetcode.com/problems/reverse-linked-list)/[力扣](https://leetcode-cn.com/problems/reverse-linked-list)

非递归，三个指针，每次都是后面指向前面一个

```c++
ListNode* reverseList(ListNode* head) {
    ListNode* prev = nullptr;
    ListNode* cur = head;
    while(cur != nullptr) {
        ListNode* next = cur->next;
        cur->next = prev;
        prev = cur;
        cur = next;
    }
    return prev;
}
```

## Odd Even Linked List
[LeetCode](https://leetcode.com/problems/odd-even-linked-list)/[力扣](https://leetcode-cn.com/problems/odd-even-linked-list)

用两个指针保存奇偶的链表，把偶指针加到奇指针末尾

```c++
ListNode* oddEvenList(ListNode* head) {
    if(head == nullptr || head->next == nullptr) return head;
    ListNode* odd = head;
    ListNode* even = head->next;
    ListNode* eh = head->next;
    while(even != nullptr && even->next != nullptr) {
        odd->next = even->next;
        odd = odd->next;
        even->next = odd->next;
        even = even->next;
    }
    odd->next = eh;
    return head;
}
```

分别保存几个指针，就来回倒腾就可以了

```c++
ListNode* oddEvenList(ListNode* head) {
    if(head == nullptr || head->next == nullptr) return head;
    ListNode* op = head;
    ListNode* cur = head->next->next;
    ListNode* eh = head->next;
    ListNode* ep = head->next;
    
    while(cur != nullptr) {
        ListNode* t = cur->next;
        op->next = cur;
        ep->next = t;
        cur->next = eh;
        op = cur;
        ep = t;
        cur = t == nullptr ? nullptr : t->next;
    }
    return head;
}
```