---
author: 黄健宏
title: HSET
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 41
tags: ["Redis", "哈希表", "HSET"]
---
# HSET hash field value

> 可用版本： >= 2.0.0
> 
> 时间复杂度： O(1)

将哈希表 `hash` 中域 `field` 的值设置为 `value` 。

如果给定的哈希表并不存在， 那么一个新的哈希表将被创建并执行 `HSET` 操作。

如果域 `field` 已经存在于哈希表中， 那么它的旧值将被新值 `value` 覆盖。

## 返回值

当 `HSET` 命令在哈希表中新创建 `field` 域并成功为它设置值时， 命令返回 `1` ； 如果域 `field` 已经存在于哈希表， 并且 `HSET` 命令成功使用新值覆盖了它的旧值， 那么命令返回 `0` 。

## 代码示例
设置一个新域：

{{< highlight shell >}}
redis> HSET website google "www.g.cn"
(integer) 1

redis> HGET website google
"www.g.cn"
{{< /highlight >}}

对一个已存在的域进行更新：

{{< highlight shell >}}
redis> HSET website google "www.google.com"
(integer) 0

redis> HGET website google
"www.google.com"
{{< /highlight >}}