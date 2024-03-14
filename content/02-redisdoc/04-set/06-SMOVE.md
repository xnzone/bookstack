---
author: 黄健宏
title: SMOVE
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 406
tags:
  - Redis
  - 集合
  - SMOVE
---

# SMOVE source destination member

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(1)

将 `member` 元素从 `source` 集合移动到 `destination` 集合。

[SMOVE](https://bookstack.xnzone.eu.org/02-redisdoc/04-set/06-smove) 是原子性操作。

如果 `source` 集合不存在或不包含指定的 `member` 元素，则 [SMOVE](https://bookstack.xnzone.eu.org/02-redisdoc/04-set/06-smove) 命令不执行任何操作，仅返回 `0` 。否则， `member` 元素从 `source` 集合中被移除，并添加到 `destination` 集合中去。

当 `destination` 集合已经包含 `member` 元素时， [SMOVE](https://bookstack.xnzone.eu.org/02-redisdoc/04-set/06-smove) 命令只是简单地将 `source` 集合中的 `member` 元素删除。

当 `source` 或 `destination` 不是集合类型时，返回一个错误。

## 返回值

如果 `member` 元素被成功移除，返回 `1` 。 如果 `member` 元素不是 `source` 集合的成员，并且没有任何操作对 `destination` 集合执行，那么返回 `0` 。

## 代码示例

{{< highlight shell >}}
redis> SMEMBERS songs
1) "Billie Jean"
2) "Believe Me"

redis> SMEMBERS my_songs
(empty list or set)

redis> SMOVE songs my_songs "Believe Me"
(integer) 1

redis> SMEMBERS songs
1) "Billie Jean"

redis> SMEMBERS my_songs
1) "Believe Me"
{{< /highlight >}}