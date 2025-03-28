---
author: 黄健宏
title: SREM
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20405
tags:
  - Redis
  - 集合
  - SREM
---

# SREM key member [member …]

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(N)， `N` 为给定 `member` 元素的数量。

移除集合 `key` 中的一个或多个 `member` 元素，不存在的 `member` 元素会被忽略。

当 `key` 不是集合类型，返回一个错误。

Note

在 Redis 2.4 版本以前， [SREM](../../02-redisdoc/04-set/05-srem/) 只接受单个 `member` 值。

## 返回值

被成功移除的元素的数量，不包括被忽略的元素。

## 代码示例

```shell
# 测试数据

redis> SMEMBERS languages
1) "c"
2) "lisp"
3) "python"
4) "ruby"

# 移除单个元素

redis> SREM languages ruby
(integer) 1

# 移除不存在元素

redis> SREM languages non-exists-language
(integer) 0

# 移除多个元素

redis> SREM languages lisp python c
(integer) 3

redis> SMEMBERS languages
(empty list or set)
```