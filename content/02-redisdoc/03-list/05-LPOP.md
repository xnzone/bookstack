---
author: 黄健宏
title: LPOP
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20305
tags:
  - Redis
  - 列表
  - LPOP
---

# LPOP key

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

移除并返回列表 `key` 的头元素。

## 返回值

列表的头元素。 当 `key` 不存在时，返回 `nil` 。

## 代码示例

```shell
redis> LLEN course
(integer) 0

redis> RPUSH course algorithm001
(integer) 1

redis> RPUSH course c++101
(integer) 2

redis> LPOP course  # 移除头元素
"algorithm001"
```