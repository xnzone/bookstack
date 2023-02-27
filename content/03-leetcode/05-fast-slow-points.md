---
title: "快慢指针"
date: 2022-06-26T15:43:42+08:00
tags: ["leetcode", "fast and slow points", "c++"]
image: /covers/leetcode.png
cover: false 
weight: 7 
---

## Add Two Numbers
[LeetCode](https://leetcode.com/problems/add-two-numbers)/[力扣](https://leetcode-cn.com/problems/add-two-numbers)

- 模拟两个数相加
- 用一个数表示进位

{{< highlight cpp >}}
ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    ListNode *t1 = l1, *t2 = l2;
    int n1 = 0, n2 = 0, c = 0;
    ListNode *prev = nullptr;
    while(t1 && t2) {
        int sum = t1->val + t2 -> val + c;
        t1->val = t2 -> val = sum % 10;
        c = sum / 10;
        n1++; n2++;
        prev = t1;
        t1 = t1 -> next; t2 = t2 -> next;
    }
    while(t1) {
        int sum = t1 -> val + c;
        t1 -> val = sum % 10;
        c = sum / 10;
        n1++;
        prev = t1;
        t1 = t1 -> next;
    }
    while(t2) {
        int sum = t2 -> val + c;
        t2 -> val = sum % 10;
        c = sum / 10;
        n2++;
        prev = t2;
        t2 = t2 -> next;
    }
    if (c > 0) {
        ListNode *cry = new ListNode(c);
        prev->next = cry;
    }
    return n1 >= n2 ? l1 : l2;
}
{{< /highlight  >}}

## Remove Nth Node From End of List
[LeetCode](https://leetcode.com/problems/remove-nth-node-from-end-of-list)/[力扣](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list)

- 一个指针先向前移动`n - 1`长度
- 然后两个指针一起移动
- 直到第一个指针移动到了末尾
- 慢指针后面那个就是要被移除的，执行删除即可

{{< highlight cpp >}}
ListNode* removeNthFromEnd(ListNode* head, int n) {
    ListNode* p = new ListNode(0);
    p->next = head;
    if(n <= 0) return head;
    ListNode *fast = p, *slow = p;
    for(int i = 0; i < n && fast != nullptr; i++) {
        fast = fast -> next;
    }
    while(fast != nullptr && fast -> next != nullptr) {
        fast = fast -> next;
        slow = slow -> next;
    }
    slow -> next = slow -> next == nullptr ? nullptr : slow -> next -> next;
    return p -> next;
}
{{< /highlight  >}}

## Remove Duplicates from Sorted List
[LeetCode](https://leetcode.com/problems/remove-duplicates-from-sorted-list)/[力扣](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list)

- 两个指针移动
- 如果两个指针值相同， 直接把slow指针的next指向fast的next
- 否则 slow = fast，fast = fast -> next

{{< highlight cpp >}}
ListNode* deleteDuplicates(ListNode* head) {
    if(head == nullptr) return head;
    ListNode *left = head, *right = head -> next;

    while(right != nullptr) {
        if(left -> val == right -> val) {
            left - > next = right -> next;
        } else {
            left = right;
        }
        right = right -> next;
    }
    return head;
}
{{< /highlight  >}}

## Linked List Cycle
[LeetCode](https://leetcode.com/problems/linked-list-cycle)/[力扣](https://leetcode-cn.com/problems/linked-list-cycle)

- 快指针走两步， 慢指针走一步
- 如果慢指针追上了快指针，则有环
- 否则，无环

{{< highlight cpp >}}
bool hasCycle(ListNode* head) {
    if(head == nullptr || head -> next == nullptr) return false;
    ListNode *fast = head -> next, *slow = head;

    while(fast != nullptr && fast -> next != nullptr && slow != nullptr) {
        if (fast == slow) return true;
        fast = fast -> next -> next;
        slow = slow -> next;
    }
    return false;
}
{{< /highlight  >}}

## Linked List Cycle II
[LeetCode](https://leetcode.com/problems/linked-list-cycle-ii)/[力扣](https://leetcode-cn.com/problems/linked-list-cycle-ii)

- 先用快慢指针找到两个相等的位置
- 然后慢指针从头开始走， 快指针继续走
- 两者再次相遇的地方就是环开始的地方
- 否则就没有环

{{< highlight cpp >}}
ListNode* detectCycle(ListNode* head) {
    if(head == nullptr || head -> next == nullptr) return nullptr;
    ListNode *slow = head, *fast = head;

    do {
        if(fast == nullptr || fast -> next == nullptr) return nullptr;
        slow = slow -> next;
        fast = fast -> next -> next;
    } while(fast != slow)
    slow = head;
    while(fast != slow && fast != nullptr && slow != nullptr) {
        slow = slow -> next;
        fast = fast -> next;
    }
    return slow;
}
{{< /highlight  >}}

## Reorder List
[LeetCode](https://leetcode.com/problems/reorder-list)/[力扣](https://leetcode-cn.com/problems/reorder-list)

- 先用快慢指针找到中点位置
- 然后后半部分链表翻转
- 最后两个链表合并

{{< highlight cpp >}}
ListNode* reverse(ListNode* head) {
    ListNode *prev = nullptr, *cur = head;
    while(cur != nullptr) {
        ListNode* next = cur -> next;
        cur -> next = prev;
        prev = cur;
        cur = next;
    }
    return prev;
}
void recorderList(ListNode* head) {
    if(head == nullptr || head -> next != nullptr) return;
    ListNode *slow = head, *fast = head -> next;
    while(fast != nullptr && fast -> next != nullptr) {
        slow = slow -> next;
        fast = fast -> next;
    }
    fast = reverse(slow);
    slow = head;
    while(fast != nullptr && slow != nullptr) {
        ListNode* st = slow -> next, *ft = fast -> next;
        slow -> next = fast;
        fast -> next = st;
        slow = st;
        fast =ft;
    }
}
{{< /highlight  >}}

## Sort List
[LeetCode](https://leetcode.com/problems/sort-list)/[力扣](https://leetcode-cn.com/problems/sort-list)

- 归并排序

{{< highlight cpp >}}
ListNode* merge(ListNode* p1, ListNode* p2) {
    ListNode* t, *head = p1;
    while(p2 && p1) {
        if(p1 -> val > p2 -> val) {
            int temp = p1 -> val;
            p1 -> val = p2 -> val;
            p2 -> val = temp;

            t = p2 -> next;
            p2 -> next = p1 -> next;
            p1 -> next = p2;
            p2 = t;
            p1 = p1 -> next;
        } else {
            p1 = p1 -> next;
        }
    }
    if(!p1 && p2) {
        t = head;
        while(t ->next) t = t -> next;
        t -> next =p2;
    }
    return head;
}

ListNode* sort(ListNode* head) {
    if(head -> next == nullptr) return head;
    ListNode *slow = head, *fast = head;
    while(fast -> next != nullptr && fast -> next -> next != nullptr) {
        slow = slow -> next;
        fast = fast -> next -> next;
    }
    ListNode* head2 = slow -> next;
    slow -> next = nullptr;
    sort(head);
    sort(head2);
    return merge(head, head2);
}
ListNode* sortList(ListNode* head) {
    if(head == nullptr) return head;
    return sort(head);
}
{{< /highlight  >}}

## Remove Linked List Elements
[LeetCode](https://leetcode.com/problems/remove-linked-list-elements)/[力扣](https://leetcode-cn.com/problems/remove-linked-list-elements)

- 删除前后，链表不要断

{{< highlight cpp >}}
ListNode* removeElements(ListNode* head, int val) {
    ListNode *prev = head;
    while(prev != nullptr && prev -> val == val) {
        prev = prev -> next;
    }
    head = prev;
    if(head == nullptr) return head;
    ListNode* next = prev -> next;
    while(next != nullptr) {
        if(next -> val == val) {
            prev -> next = next -> next;
            next -> next -> next;
            continue;
        }
        prev = prev -> next;
        next = next -> next;
    }
    return head;
}
{{< /highlight  >}}

## Palindrome Linked List
[LeetCode](https://leetcode.com/problems/palindrome-linked-list)/[力扣](https://leetcode-cn.com/problems/palindrome-linked-list)

- 递归， 递归到最后，相当于首尾逐个比较

{{< highlight cpp >}}
int ans = 1;
ListNode *root;
void helper(ListNode *head) {
    if(head == nullptr) return;
    helper(head->next);
    if(head -> val  != root -> val) ans = 0;
    root = root -> next;
}

bool isPalindrome(ListNode* head) {
    if(head == nullptr || head -> next == nullptr) return true;
    root =head;
    helper(head);
    return ans == 1;
}
{{< /highlight  >}}

- 快慢指针找到后面的链表
- 然后翻转后半部分
- 再逐一比较

{{< highlight cpp >}}
ListNode *reverse(ListNode *head) {
    ListNode prev = nullptr;
    while(head) {
        ListNode* next = head-> next;
        head -> next = prev;
        prev = head;
        head = next;
    }
    return prev;
}

bool isPalindrome(ListNode *head) {
    if(!head) return true;

    ListNode *slow = head, *fast = head;

    while(fast && fast -> next && fast -> next -> next) {
        fast = fast -> next -> next;
        slow = slow -> next;
    }
    fast = slow -> next;
    slow -> next  = nullptr;
    for(ListNode *left = head, *right = reverse(fast); left && right; left = left -> next, right = right -> next) {
        if(left -> val != right -> val) return false;
    }
    slow -> next = fast;
    return true;
}

{{< /highlight  >}}

## Middle of the Linked List
[LeetCode](https://leetcode.com/problems/middle-of-the-linked-list)/[力扣](https://leetcode-cn.com/problems/middle-of-the-linked-list)

- 快慢指针直接解决

{{< highlight cpp >}}
ListNode *middleNode(ListNode *head) {
    if(head == nullptr) return head;
    ListNode *prev = head, *next = head;
    while(next != nullptr && next -> next != nullptr) {
        prev = prev -> next;
        next = next -> next -> next;
    }
    return prev;
}
{{< /highlight  >}}