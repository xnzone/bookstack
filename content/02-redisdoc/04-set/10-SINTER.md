---
author: 黄健宏
title: SINTER
date: 2024-03-07 15:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20410
tags:
  - Redis
  - 集合
  - SINTER
---

# SINTER key [key …]

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(N * M)， `N` 为给定集合当中基数最小的集合， `M` 为给定集合的个数。

返回一个集合的全部成员，该集合是所有给定集合的交集。

不存在的 `key` 被视为空集。

当给定集合当中有一个空集时，结果也为空集(根据集合运算定律)。

## 返回值

交集成员的列表。

## 代码示例

```shell
redis> SMEMBERS group_1
1) "LI LEI"
2) "TOM"
3) "JACK"

redis> SMEMBERS group_2
1) "HAN MEIMEI"
2) "JACK"

redis> SINTER group_1 group_2
1) "JACK"
```