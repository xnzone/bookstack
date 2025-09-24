---
author: 黄健宏
title: HDEL
date: 2024-03-07 15:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20205
tags:
  - Redis
  - 哈希表
  - HDEL
---

# HDEL

**HDEL key field [field …]**

删除哈希表 `key` 中的一个或多个指定域，不存在的域将被忽略。

Note

在Redis2.4以下的版本里， `HDEL` 每次只能删除单个域，如果你需要在一个原子时间内删除多个域，请将命令包含在 [MULTI](../../02-redisdoc/11-transaction/01-multi) / [EXEC](../../02-redisdoc/11-transaction/02-exec) 块内。

**可用版本：**

>= 2.0.0

**时间复杂度:**

O(N)， `N` 为要删除的域的数量。

**返回值:**

被成功移除的域的数量，不包括被忽略的域。


## 测试数据

```shell
redis> HGETALL abbr
1) "a"
2) "apple"
3) "b"
4) "banana"
5) "c"
6) "cat"
7) "d"
8) "dog"
```

删除单个域

```shell
redis> HDEL abbr a
(integer) 1
```

删除不存在的域

```shell
redis> HDEL abbr not-exists-field
(integer) 0
```

删除多个域

```shell
redis> HDEL abbr b c
(integer) 2

redis> HGETALL abbr
1) "d"
2) "dog"
```