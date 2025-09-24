---
author: 黄健宏
title: RPOP
date: 2024-03-07 15:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20306
tags:
  - Redis
  - 列表
  - RPOP
---

# RPOP key

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

移除并返回列表 `key` 的尾元素。

## 返回值

列表的尾元素。 当 `key` 不存在时，返回 `nil` 。

## 代码示例

```shell
redis> RPUSH mylist "one"
(integer) 1

redis> RPUSH mylist "two"
(integer) 2

redis> RPUSH mylist "three"
(integer) 3

redis> RPOP mylist           # 返回被弹出的元素
"three"

redis> LRANGE mylist 0 -1    # 列表剩下的元素
1) "one"
2) "two"
```