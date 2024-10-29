---
author: 黄健宏
title: MSETNX
date: 2024-03-10 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 117
tags: ["Redis", "字符串", "MSETNX"]
---

# MSETNX key value [key value …]

> 可用版本： >= 1.0.1
> 
> 时间复杂度： O(N)， 其中 N 为被设置的键数量。

当且仅当所有给定键都不存在时， 为所有给定键设置值。

即使只有一个给定键已经存在， `MSETNX` 命令也会拒绝执行对所有键的设置操作。

`MSETNX` 是一个原子性(atomic)操作， 所有给定键要么就全部都被设置， 要么就全部都不设置， 不可能出现第三种状态。

## 返回值

当所有给定键都设置成功时， 命令返回 `1` ； 如果因为某个给定键已经存在而导致设置未能成功执行， 那么命令返回 `0` 。

## 代码示例

对不存在的键执行 `MSETNX` 命令：

```shell
redis> MSETNX rmdbs "MySQL" nosql "MongoDB" key-value-store "redis"
(integer) 1

redis> MGET rmdbs nosql key-value-store
1) "MySQL"
2) "MongoDB"
3) "redis"
```

对某个已经存在的键进行设置：

```shell
redis> MSETNX rmdbs "Sqlite" language "python"  # rmdbs 键已经存在，操作失败
(integer) 0

redis> EXISTS language                          # 因为 MSETNX 命令没有成功执行
(integer) 0                                     # 所以 language 键没有被设置

redis> GET rmdbs                                # rmdbs 键也没有被修改
"MySQL"
```