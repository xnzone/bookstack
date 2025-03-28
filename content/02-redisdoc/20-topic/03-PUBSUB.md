---
author: 黄健宏
title: 发布与订阅
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 22003
tags:
  - Redis
  - 功能文档
  - 发布与订阅
---

# 发布与订阅（pub/sub）

> **Note:**  
> 本文档翻译自： [http://redis.io/topics/pubsub](http://redis.io/topics/pubsub) 。

[SUBSCRIBE channel [channel …]](../pubsub/subscribe.html#subscribe) 、 [UNSUBSCRIBE [channel [channel …]]](../pubsub/unsubscribe.html#unsubscribe) 和 [PUBLISH channel message](../pubsub/publish.html#publish) 三个命令实现了[发布与订阅信息泛型](http://en.wikipedia.org/wiki/Publish/subscribe)（Publish/Subscribe messaging paradigm）， 在这个实现中， 发送者（发送信息的客户端）不是将信息直接发送给特定的接收者（接收信息的客户端）， 而是将信息发送给频道（channel）， 然后由频道将信息转发给所有对这个频道感兴趣的订阅者。

发送者无须知道任何关于订阅者的信息， 而订阅者也无须知道是那个客户端给它发送信息， 它只要关注自己感兴趣的频道即可。

对发布者和订阅者进行解构（decoupling）， 可以极大地提高系统的扩展性（scalability）， 并得到一个更动态的网络拓扑（network topology）。

比如说， 要订阅频道 `foo` 和 `bar` ， 客户端可以使用频道名字作为参数来调用 [SUBSCRIBE channel [channel …]](../pubsub/subscribe.html#subscribe) 命令：

```bash
redis> SUBSCRIBE foo bar
```

当有客户端发送信息到这些频道时， Redis 会将传入的信息推送到所有订阅这些频道的客户端里面。

正在订阅频道的客户端不应该发送除 [SUBSCRIBE channel [channel …]](../pubsub/subscribe.html#subscribe) 和 [UNSUBSCRIBE [channel [channel …]]](../pubsub/unsubscribe.html#unsubscribe) 之外的其他命令。 其中， [SUBSCRIBE channel [channel …]](../pubsub/subscribe.html#subscribe) 可以用于订阅更多频道， 而 [UNSUBSCRIBE [channel [channel …]]](../pubsub/unsubscribe.html#unsubscribe) 则可以用于退订已订阅的一个或多个频道。

[SUBSCRIBE channel [channel …]](../pubsub/subscribe.html#subscribe) 和 [UNSUBSCRIBE [channel [channel …]]](../pubsub/unsubscribe.html#unsubscribe) 的执行结果会以信息的形式返回， 客户端可以通过分析所接收信息的第一个元素， 从而判断所收到的内容是一条真正的信息， 还是 [SUBSCRIBE channel [channel …]](../pubsub/subscribe.html#subscribe) 或 [UNSUBSCRIBE [channel [channel …]]](../pubsub/unsubscribe.html#unsubscribe) 命令的操作结果。

## 信息的格式

频道转发的每条信息都是一条带有三个元素的多条批量回复（multi-bulk reply）。

信息的第一个元素标识了信息的类型：

- `subscribe` ： 表示当前客户端成功地订阅了信息第二个元素所指示的频道。 而信息的第三个元素则记录了目前客户端已订阅频道的总数。
    
- `unsubscribe` ： 表示当前客户端成功地退订了信息第二个元素所指示的频道。 信息的第三个元素记录了客户端目前仍在订阅的频道数量。 当客户端订阅的频道数量降为 `0` 时， 客户端不再订阅任何频道， 它可以像往常一样， 执行任何 Redis 命令。
    
- `message` ： 表示这条信息是由某个客户端执行 [PUBLISH channel message](../pubsub/publish.html#publish) 命令所发送的， 真正的信息。 信息的第二个元素是信息来源的频道， 而第三个元素则是信息的内容。
    

举个例子， 如果客户端执行以下命令：

```bash
redis> SUBSCRIBE first second
```

那么它将收到以下回复：

```bash
1) "subscribe"
2) "first"
3) (integer) 1

1) "subscribe"
2) "second"
3) (integer) 2
```

如果在这时， 另一个客户端执行以下 [PUBLISH channel message](../pubsub/publish.html#publish) 命令：

```bash
redis> PUBLISH second Hello
```

那么之前订阅了 `second` 频道的客户端将收到以下信息：

```bash
1) "message"
2) "second"
3) "hello"
```

当订阅者决定退订所有频道时， 它可以执行一个无参数的 [UNSUBSCRIBE [channel [channel …]]](../pubsub/unsubscribe.html#unsubscribe) 命令：

```bash
redis> UNSUBSCRIBE
```

这个命令将接到以下回复：

```bash
1) "unsubscribe"
2) "second"
3) (integer) 1

1) "unsubscribe"
2) "first"
3) (integer) 0
```

## 订阅模式

Redis 的发布与订阅实现支持模式匹配（pattern matching）： 客户端可以订阅一个带 `*` 号的模式， 如果某个/某些频道的名字和这个模式匹配， 那么当有信息发送给这个/这些频道的时候， 客户端也会收到这个/这些频道的信息。

比如说，执行命令

```bash
redis> PSUBSCRIBE news.*
```

的客户端将收到来自 `news.art.figurative` 、 `news.music.jazz` 等频道的信息。

客户端订阅的模式里面可以包含多个 glob 风格的通配符， 比如 `*` 、 `?` 和 `[...]` ， 等等。

执行命令

```bash
redis> PUNSUBSCRIBE news.*
```

将退订 `news.*` 模式， 其他已订阅的模式不会被影响。

通过订阅模式接收到的信息， 和通过订阅频道接收到的信息， 这两者的格式不太一样：

- 通过订阅模式而接收到的信息的类型为 `pmessage` ： 这代表有某个客户端通过 [PUBLISH channel message](../pubsub/publish.html#publish) 向某个频道发送了信息， 而这个频道刚好匹配了当前客户端所订阅的某个模式。 信息的第二个元素记录了被匹配的模式， 第三个元素记录了被匹配的频道的名字， 最后一个元素则记录了信息的实际内容。
    

客户端处理 [PSUBSCRIBE pattern [pattern …]](../pubsub/psubscribe.html#psubscribe) 和 [PUNSUBSCRIBE [pattern [pattern …]]](../pubsub/punsubscribe.html#punsubscribe) 返回值的方式， 和客户端处理 [SUBSCRIBE channel [channel …]](../pubsub/subscribe.html#subscribe) 和 [UNSUBSCRIBE [channel [channel …]]](../pubsub/unsubscribe.html#unsubscribe) 的方式类似： 通过对信息的第一个元素进行分析， 客户端可以判断接收到的信息是一个真正的信息， 还是 [PSUBSCRIBE pattern [pattern …]](../pubsub/psubscribe.html#psubscribe) 或 [PUNSUBSCRIBE [pattern [pattern …]]](../pubsub/punsubscribe.html#punsubscribe) 命令的返回值。

## 通过频道和模式接收同一条信息

如果客户端订阅的多个模式匹配了同一个频道， 或者客户端同时订阅了某个频道、以及匹配这个频道的某个模式， 那么它可能会多次接收到同一条信息。

举个例子， 如果客户端执行了以下命令：

```bash
SUBSCRIBE foo
PSUBSCRIBE f*
```

那么当有信息发送到频道 `foo` 时， 客户端将收到两条信息： 一条来自频道 `foo` ，信息类型为 `message` ； 另一条来自模式 `f*` ，信息类型为 `pmessage` 。

## 订阅总数

在执行 [SUBSCRIBE channel [channel …]](../pubsub/subscribe.html#subscribe) 、 [UNSUBSCRIBE [channel [channel …]]](../pubsub/unsubscribe.html#unsubscribe) 、 [PSUBSCRIBE pattern [pattern …]](../pubsub/psubscribe.html#psubscribe) 和 [PUNSUBSCRIBE [pattern [pattern …]]](../pubsub/punsubscribe.html#punsubscribe) 命令时， 返回结果的最后一个元素是客户端目前仍在订阅的频道和模式总数。

当客户端退订所有频道和模式， 也即是这个总数值下降为 `0` 的时候， 客户端将退出订阅与发布状态。

## 编程示例

Pieter Noordhuis 提供了一个使用 EventMachine 和 Redis 编写的 [高性能多用户网页聊天软件](https://gist.github.com/348262) ， 这个软件很好地展示了发布与订阅功能的用法。

## 客户端库实现提示

因为所有接收到的信息都会包含一个信息来源：

- 当信息来自频道时，来源是某个频道；
    
- 当信息来自模式时，来源是某个模式。
    

因此， 客户端可以用一个哈希表， 将特定来源和处理该来源的回调函数关联起来。 当有新信息到达时， 程序就可以根据信息的来源， 在 O(1) 复杂度内， 将信息交给正确的回调函数来处理。