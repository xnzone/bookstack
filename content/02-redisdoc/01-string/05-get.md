---
author: 黄健宏
title: GET
date: 2024-03-07 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 105 
tags: ["Redis", "字符串", "GET"]
---

# GET key

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

返回与键 `key` 相关联的字符串值。

## 返回值

如果键 `key` 不存在， 那么返回特殊值 `nil` ； 否则， 返回键 `key` 的值。

如果键 `key` 的值并非字符串类型， 那么返回一个错误， 因为 `GET` 命令只能用于字符串值。

## 代码示例

对不存在的键 `key` 或是字符串类型的键 `key` 执行 `GET` 命令：

{{< highlight shell >}}
redis> GET db
(nil)

redis> SET db redis
OK

redis> GET db
"redis"
{{< /highlight >}}

对不是字符串类型的键 `key` 执行 `GET` 命令：

{{< highlight shell >}}
redis> DEL db
(integer) 1

redis> LPUSH db redis mongodb mysql
(integer) 3

redis> GET db
(error) ERR Operation against a key holding the wrong kind of value
{{< /highlight >}}