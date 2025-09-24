---
author: xnzone 
title: 排序 
date: 2023-02-22 10:04:00
image: https://s2.loli.net/2025/09/24/QCDLM3PdlaS4jgv.png
cover: false
math: true
weight: 908
tags: ["数据结构", "算法", "排序"]
---

## 插入排序

### 算法思想

每次将一个待排序的记录按其关键字大小插入到前面已排好序的子序列中，直到全部记录插入完成

### 代码实现

- 最好时间复杂度：O(n)
- 最坏时间复杂度：O(n^2)
- 算法稳定性：稳定

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-insert.jpg)

**优化**
- 先用折半思想找到插入位置，再移动元素

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-insert-better.jpg)

## 希尔排序

### 算法思想

- 先追求表中元素部分有序，再逼近全局有序

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-shell-base.jpg)

### 代码实现

- 最坏时间复杂度：O(n^2),当n在某个范围时，可达O(n^1.3)
- 算法稳定性：不稳定

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-shell.jpg)

## 冒泡排序

### 算法思想

- 从后往前(或从前往后)两两比较相邻元素的值，若为逆序，则交换他们，直到序列比较完

### 代码实现

- 最好时间复杂度：O(n) 有序的情况
- 最坏时间复杂度：O(n^2) 逆序的情况
- 算法稳定性：稳定

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-bubble-swap.jpg)
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-bubble.jpg)

## 快速排序

### 算法思想

- 在待排序列表中，任意选取一个元素pivot作为枢轴
- 通过一趟排序将待排序表划分为两部分
- 一部分全小于pivot，一部分全大于pivot
- 此时pivot放在了正确排序的位置，整个过程为一次划分
- 分别递归对两个子表进行重复上述过程，直到每部分内只有一个元素或者空为止

### 代码实现

- 最好时间复杂度：O($n\log_2n$)
- 最坏时间复杂度：O(n^2) 原本有序或逆序
- 最好空间复杂度：O($\log_2n$)
- 最坏空间复杂度：O(n)
- 算法稳定性：不稳定
- 所有排序算法中平均性能最优的


![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-quick-part.jpg)
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-quick.jpg)

## 选择排序

### 算法思想

每一趟再待排序元素中选取关键字最小(或最大)的元素加入有序子序列

### 代码实现

- 时间复杂度：O(n^2)
- 算法稳定性：不稳定

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-select.jpg)

## 堆排序

- 最大堆：完全二叉树中，任意根 >= 左、右
- 最小堆：完全二叉树中，任意根 <= 左、右

### 建立最大堆

- 所有非终端节点都检查一遍，是否满足最大堆的要求，不满足则调整
- 在顺序存储的完全二叉树中，非终端节点编号 i <= [n / 2]
- 检查当前节点是否满足，不满足，将当前节点与更大的一个孩子互换
- 若元素互换破坏了下一级的堆，则采用相同的方法继续调整

### 最大堆代码

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-max-heap-build.jpg)
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-max-heap-head.jpg)

### 最大堆排序

- 每一趟在待排序元素中，选取关键字最大的元素加入有序子序列
- 将堆顶元素加入有序子序列(与待排序序列中的最后一个元素交换)
- 并将待排序元素序列再次调整为最大堆

- 建堆时间复杂度：O(n)
- 排序时间复杂度：O($n\log_2n$)
- 算法稳定性：不稳定

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-max-heap.jpg)

### 插入和删除

**插入**

- 对于最小堆，新元素放到表尾，与父节点比较，若新元素比父节点小，则将二者互换

**删除**

- 被删除的元素用堆底替代
- 让该元素不断下坠，直到无法下坠为止

## 归并排序

### 算法思想

- 把两个或多个已经有序的序列合并成一个
- 对于两个有序序列，将i、j指针指向序列的表头，选择更小的一个放入k所指的位置
- k++，i/j指向更小元素的指针++
- 只剩一个子表未合并时，可以将该表的剩余元素全部加到总表
- m路归并：每选出一个小的元素，需要对比关键字m-1次
- 核心操作：把数组内的两个有序序列归并为一个

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-merge-base.jpg)

### 代码实现

- 时间复杂度：O($n\log_2n$)
- 空间复杂度：O(n)
- 算法稳定性：稳定

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-merge-sort.jpg)
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-merge.jpg)

## 基数排序

### 算法思想

- 初始化： 设置r个空队列
- 按照各个关键字位权重递增的次序，对d个关键字分别左分配和收集
- 分配：顺序扫描各个元素，若当前处理的关键字位=x，则将元素插入Qx队尾
- 收集：把Qr-1,Qr-2,...,Q0各个队列中的节点一次出队并连接

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-basic-base-0.jpg)
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-basic-base-1.jpg)

### 算法分析和应用

- 空间复杂度：O($r$)
- 时间复杂度：O(d(n+r))
- 算法稳定性：稳定
- 学校有10000名学生，将学生信息按照年龄递减排序
- 给十亿的身份证号排序
- 数据元素的关键字可以方便拆分为d组，且d较小
- 每组关键字的取值范围不大，即r较小
- 数据元素个数n很大

## 外部排序

### 原理

- 数据元素太多，无法一次全部读入内存进行排序
- 使用归并排序的方法，最少只需在内存或只能分配3块大小的缓冲区即可对任意一个大文件进行排序
- 归并排序要求各个子序列有序，每次读入两个块的内容，进行内部排序后写回磁盘

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/merge-out.jpg)

### 多路归并

- 采用多路归并可以减少归并趟数，从而减少磁盘I/O
- 对r个初始归并段，做k路归并，则归并树可用k叉树表示，若树高为h，则归并趟数=h-l=[$\log_kr$]
- k越大，r越小，归并数越少
- 负面影响：内存开销增加，内部归并时间增加(可使用败者树优化)


## 败者树

### 定义

- 可视为完全二叉树(多了一个头)
- k个叶节点分别是当前参加比较的元素
- 非叶子节点用来记忆左右子树中的失败者
- 胜者往上继续比较，一直到根节点
- 即失败者留在这一回合，胜利者进入下一回合比拼

### 败者树在多路平衡归并中的应用

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-fail-tree-0.jpg)
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-fail-tree-1.jpg)

### 败者树存储结构

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/sort-fail-tree-struct.jpg)

## 置换-选择排序

- 可以让每个初始归并段的长度超过内存工作区大小限制

### 算法步骤

文件为FI，初始归并短输出文件为FO，内存工作区为WA

- 从FI输入w个记录到工作区WA
- 从WA中选出其中关键字最小的记录，记为MINIMAX记录
- 将MINIMAX记录输出到FO中去
- 若FI不空，则从FI输入下一个记录到WA中
- 从WA中所有关键字比MINIMAX记录的关键字大的记录中选出最小关键字记录，作为新的MINMAX记录
- 重复(3~5)，直到WA中选不出新的MINIMAX为止，由此得到一个初始归并段，输出一个归并段的结束标志到FO中去
- 重复(2~6),直到WA为空，得到全部初始归并段

