---
author: 黄健宏
title: LPUSHX
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 2302
tags:
  - Redis
  - 列表
  - LPUSHX
---

# LPUSHX key value

> 可用版本： >= 2.2.0
> 
> 时间复杂度： O(1)

将值 `value` 插入到列表 `key` 的表头，当且仅当 `key` 存在并且是一个列表。

和 [LPUSH key value [value …]](../../02-redisdoc/03-list/02-lpushx/) 命令相反，当 `key` 不存在时， [LPUSHX](../../02-redisdoc/03-list/02-lpushx/) 命令什么也不做。

## 返回值

[LPUSHX](../../02-redisdoc/03-list/02-lpushx/) 命令执行之后，表的长度。

## 代码示例

```shell
# 对空列表执行 LPUSHX

redis> LLEN greet                       # greet 是一个空列表
(integer) 0

redis> LPUSHX greet "hello"             # 尝试 LPUSHX，失败，因为列表为空
(integer) 0

# 对非空列表执行 LPUSHX

redis> LPUSH greet "hello"              # 先用 LPUSH 创建一个有一个元素的列表
(integer) 1

redis> LPUSHX greet "good morning"      # 这次 LPUSHX 执行成功
(integer) 2

redis> LRANGE greet 0 -1
1) "good morning"
2) "hello"
```