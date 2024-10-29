---
author: 黄健宏
title: GETSET
date: 2024-03-10 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 107 
tags: ["Redis", "字符串", "STRLEN"]
---

# STRLEN key

> 可用版本： >= 2.2.0
> 
> 复杂度： O(1)

返回键 `key` 储存的字符串值的长度。

## 返回值

`STRLEN` 命令返回字符串值的长度。

当键 `key` 不存在时， 命令返回 `0` 。

当 `key` 储存的不是字符串值时， 返回一个错误。

## 代码示例

获取字符串值的长度：

```shell
redis> SET mykey "Hello world"
OK

redis> STRLEN mykey
(integer) 11
```

不存在的键的长度为 `0` ：

```shell
redis> STRLEN nonexisting
(integer) 0
```