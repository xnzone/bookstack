---
author: 黄健宏
title: RPUSHX
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 304
tags:
  - Redis
  - 列表
  - RPUSHX
---

# RPUSHX key value[

> 可用版本： >= 2.2.0
> 
> 时间复杂度： O(1)

将值 `value` 插入到列表 `key` 的表尾，当且仅当 `key` 存在并且是一个列表。

和 [RPUSH key value [value …]](../../02-redisdoc/03-list/04-rpushx/) 命令相反，当 `key` 不存在时， [RPUSHX](../../02-redisdoc/03-list/04-rpushx/) 命令什么也不做。

## 返回值

[RPUSHX](../../02-redisdoc/03-list/04-rpushx/) 命令执行之后，表的长度。

## 代码示例

```shell
# key不存在

redis> LLEN greet
(integer) 0

redis> RPUSHX greet "hello"     # 对不存在的 key 进行 RPUSHX，PUSH 失败。
(integer) 0

# key 存在且是一个非空列表

redis> RPUSH greet "hi"         # 先用 RPUSH 插入一个元素
(integer) 1

redis> RPUSHX greet "hello"     # greet 现在是一个列表类型，RPUSHX 操作成功。
(integer) 2

redis> LRANGE greet 0 -1
1) "hi"
2) "hello"
```