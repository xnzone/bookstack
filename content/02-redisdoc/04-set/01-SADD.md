---
author: 黄健宏
title: SADD
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 401
tags:
  - Redis
  - 集合
  - SADD
---

# SADD key member [member …]

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(N)， `N` 是被添加的元素的数量。

将一个或多个 `member` 元素加入到集合 `key` 当中，已经存在于集合的 `member` 元素将被忽略。

假如 `key` 不存在，则创建一个只包含 `member` 元素作成员的集合。

当 `key` 不是集合类型时，返回一个错误。

Note

在Redis2.4版本以前， [SADD](https://bookstack.xnzone.eu.org/02-redisdoc/04-set/01-sadd/) 只接受单个 `member` 值。

## 返回值

被添加到集合中的新元素的数量，不包括被忽略的元素。

## 代码示例

```shell
# 添加单个元素

redis> SADD bbs "discuz.net"
(integer) 1

# 添加重复元素

redis> SADD bbs "discuz.net"
(integer) 0

# 添加多个元素

redis> SADD bbs "tianya.cn" "groups.google.com"
(integer) 2

redis> SMEMBERS bbs
1) "discuz.net"
2) "groups.google.com"
3) "tianya.cn"
```