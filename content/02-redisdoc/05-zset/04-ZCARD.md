---
author: 黄健宏
title: ZCARD 
date: 2024-10-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 504
tags:
  - Redis
  - 集合
  - ZCARD  
---

# ZCARD key

> 可用版本： >= 1.2.0
> 
> 时间复杂度: O(1)

返回有序集 `key` 的基数。

## 返回值

当 `key` 存在且是有序集类型时，返回有序集的基数。 当 `key` 不存在时，返回 `0` 。

## 代码示例

```shell
redis > ZADD salary 2000 tom    # 添加一个成员
(integer) 1

redis > ZCARD salary
(integer) 1

redis > ZADD salary 5000 jack   # 再添加一个成员
(integer) 1

redis > ZCARD salary
(integer) 2

redis > EXISTS non_exists_key   # 对不存在的 key 进行 ZCARD 操作
(integer) 0

redis > ZCARD non_exists_key
(integer) 0
```