---
author: 黄健宏
title: GETSET
date: 2024-03-10 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 106 
tags: ["Redis", "字符串", "GETSET"]
---

# GETSET key value[

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

将键 `key` 的值设为 `value` ， 并返回键 `key` 在被设置之前的旧值。

## 返回值

返回给定键 `key` 的旧值。

如果键 `key` 没有旧值， 也即是说， 键 `key` 在被设置之前并不存在， 那么命令返回 `nil` 。

当键 `key` 存在但不是字符串类型时， 命令返回一个错误。

## 代码示例

{{< highlight shell >}}
redis> GETSET db mongodb    # 没有旧值，返回 nil
(nil)

redis> GET db
"mongodb"

redis> GETSET db redis      # 返回旧值 mongodb
"mongodb"

redis> GET db
"redis"
{{< /highlight >}}