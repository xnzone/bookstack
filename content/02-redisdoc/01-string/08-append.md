---
author: 黄健宏
title: APPEND
date: 2024-03-10 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 308 
tags: ["Redis", "字符串", "APPEND"]
---

# APPEND key value

> 可用版本： >= 2.0.0
> 
> 时间复杂度： 平摊O(1)

如果键 `key` 已经存在并且它的值是一个字符串， `APPEND` 命令将把 `value` 追加到键 `key` 现有值的末尾。

如果 `key` 不存在， `APPEND` 就简单地将键 `key` 的值设为 `value` ， 就像执行 `SET key value` 一样。

## 返回值

追加 `value` 之后， 键 `key` 的值的长度。

## 示例代码

对不存在的 `key` 执行 `APPEND` ：

{{< highlight shell >}}
redis> EXISTS myphone               # 确保 myphone 不存在
(integer) 0

redis> APPEND myphone "nokia"       # 对不存在的 key 进行 APPEND ，等同于 SET myphone "nokia"
(integer) 5                         # 字符长度
{{< /highlight >}}

对已存在的字符串进行 `APPEND` ：

{{< highlight shell >}}
redis> APPEND myphone " - 1110"     # 长度从 5 个字符增加到 12 个字符
(integer) 12

redis> GET myphone
"nokia - 1110"
{{< /highlight >}}