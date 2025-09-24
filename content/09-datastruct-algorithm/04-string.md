---
author: xnzone 
title: 字符串 
date: 2023-02-22 10:04:00
image: https://s2.loli.net/2025/09/24/QCDLM3PdlaS4jgv.png
cover: false
weight: 904 
tags: ["数据结构", "算法", "字符串"]
---

## 定义和基本操作

### 定义

- 由零个或多个字符组成的优先序列

### 基本操作

- `StrAssign(&T, chars)` 赋值。把字符串T赋值为chars
- `StrCopy(&T, S)` 复制。由S复制到T
- `StrEmpty(S)` 判空
- `StrLength(S)` 求长度
- `ClearString(&S)` 清空
- `DestroyString(&S)` 销毁。回收存储空间
- `Contact(&T,&T1,&T2)` 联接，返回S1和S2联接的新串
- `SubString(&sub, S, pos, len)` 求子串，第pos字符开始长度为len的子串
- `Index(S,T)` 定位操作，返回S中存在T相同的子串的位置，否则返回-1
- `StrCompare(S,T)` 比较操作

### 字符集编码

- 任何数据在计算机中一定是二进制
- 字符集：英文字符，ASCII字符，中英文，Unicode字符集
- 基于同一个字符集可以有不同编码方式：UTF-8，GBK
- 不同的编码方式，字符占用的空间不同

## 存储结构

### 顺序存储

- 方案1: 一个数组存储字符，一个int变量存储实际长度
- 方案2: 数组的ch[0]充当长度，优点：字符的位序和数组下标相同
- 方案3: 没有长度变量，已字符'\0'表示结尾，缺点：获取长度需要遍历
- 方案4: ch[0]废弃不用，外加一个长度变量

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/str-array.jpg)

### 链式存储

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/str-list.jpg)

### 基本操作实现-第四种方案

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/str-substr.jpg)

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/str-compare.jpg)

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/str-index.jpg)

## 匹配问题-朴素算法

### 方法1

- 将主串中所有长度为m的子串依次与模式串对比，直到找到一个完全匹配的子串，或所有子串都不匹配为止
- 主串长度为n，模式串长度为m， 最多对比 n - m + 1个子串

### 方法2

- 使用两个指针i和j进行匹配，如果当前子串匹配失败，则主串指针i指向下一个子串的第一个位置，模式串指针j回到模式串的第一个位置，即 i = i - j + 2; j = 1;
- 如果 j > T.length, 则当前串匹配成功，返回第一个字符的位置即 i - T.length


```c++
int Index(SString S, SString T) {
    int i = 1, j = 1;
    while(i <= S.length && j <= T.length) {
        if (S.ch[i] == T.ch[j]) {
            i++;
            j++;
        } else {
            i = i - j + 2;
            j = 1;
        }
    }
    if (j > T.length) {
        return i - T.length;
    } else {
        return -1;
    }
}
```

## 匹配问题-KMP算法

- 原理：减少i指针的回溯，通过已经计算好的next指针，提高算法效率
- 精髓：利用好已经匹配过的模式串信息
- next数组记录了当第几个元素匹配失败的时候，j的取值

步骤

- 根据模式串T，求出next数组
- 利用next数组进行匹配（主串指针不回溯）

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/str-kmp.jpg)

### 求next数组

- 当模式串的第j个字符匹配失败时，从模式串的第next[j]的继续往后匹配

步骤

- 第一个字符不匹配时，只能匹配下一个子串，因此next[1]都无脑写0`if(j==0){i++;j++}`
- 第二个字符不匹配时，应尝试匹配模式串的第一个字符，因此next[2]都无脑写1
- 第三个字符及之后的，在不匹配的位置前边，划一根美丽的分界线模式串一步步往后退，直到分界线之前能对上，此时j指向哪儿，next数组值就是多少


```c++
void get_next(SString T, int next[]) {
    int i = 1, j = 0;
    next[1] = 0;
    while(i < T.length) {
        if (j == 0 || T.ch[i] == T.ch[j]) {
            ++i;
            ++j;
            next[i] = j;
        } else {
            j = next[j];
        }
    }
}
```