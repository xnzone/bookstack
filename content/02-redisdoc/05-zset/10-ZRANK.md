---
author: 黄健宏
title: ZRANK
date: 2024-12-16 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 510
tags:
  - Redis
  - 集合
  - ZRANK
---


# ZRANK key member

> 可用版本： >= 2.0.0
> 
> 时间复杂度: O(log(N))

返回有序集 `key` 中成员 `member` 的排名。其中有序集成员按 `score` 值递增(从小到大)顺序排列。

排名以 `0` 为底，也就是说， `score` 值最小的成员排名为 `0` 。

使用 [ZREVRANK key member](../../05-zset/11-ZREVZRANK) 命令可以获得成员按 `score` 值递减(从大到小)排列的排名。

## 返回值

如果 `member` 是有序集 `key` 的成员，返回 `member` 的排名。 如果 `member` 不是有序集 `key` 的成员，返回 `nil` 。

## 代码示例

```bash
redis> ZRANGE salary 0 -1 WITHSCORES        # 显示所有成员及其 score 值
1) "peter"
2) "3500"
3) "tom"
4) "4000"
5) "jack"
6) "5000"

redis> ZRANK salary tom                     # 显示 tom 的薪水排名，第二
(integer) 1
```
