---
author: mixu
title: "趣谈分布式系统"
date: 2024-09-10 10:23:32
image: /covers/distr-fun.jpg
cover: true
weight: 3
tags:
  - distributed
  - system
  - CAP
  - 分布式系统
---

本文翻译自[Distributed systems: for fun and profit](http://book.mixu.net/distsys/)，仅限学习交流，如有侵权，联系本人删除

## 引言

我希望能写一篇文章，把许多现代分布式系统背后的思想汇集起来，比如 Amazon 的 Dynamo、Google 的 BigTable 和 MapReduce，以及 Apache 的 Hadoop 等。

在这篇文章中，我试图以更容易理解的方式介绍分布式系统。对我来说，这意味着两件事：一是介绍你在阅读更专业的文献时所需要的关键概念；二是提供一个贯穿全篇的叙述，覆盖足够的细节，让你理解系统的运作，而不至于在细节中迷失。现在是2013年，你可以利用互联网选择性深入阅读你感兴趣的主题。

在我看来，大部分分布式编程都与分布式带来的两个结果有关：
- 信息传播受限于光速
- 独立事件独立失败

换句话说，分布式编程的核心在于处理距离和多个独立组件。这些限制决定了系统设计的可能性空间。我希望在阅读本篇后，你能更好地理解距离、时间和一致性模型如何相互作用的。

这篇文章主要关注分布式编程和数据中心商业系统中需要理解的概念。试图覆盖一切内容是不太可行的。你将学习许多关键协议和算法（包括分布式系统领域引用最多的论文），还将了解一些尚未被大学教科书广泛收录的新概念，比如 CRDT 和 CALM 定理。

希望你喜欢这篇文章！如果你想表达感谢，可以在 [Github](https://github.com/mixu/) 或 [Twitter](https://twitter.com/mikitotakada) 上关注我。如果发现错误，[请在 Github 提交 pull request](https://github.com/mixu/distsysbook/issues)。


### 1. 基础知识

[第一章](https://book.mixu.net/distsys/single-page.html#intro) 概括性地介绍了分布式系统，涵盖了一些重要术语和概念，包括扩展性、可用性、性能、延迟和容错性等目标，以及这些目标为何难以实现，还探讨了抽象和模型、分区和复制的作用。

---

### 2. 抽象的层次

[第二章](https://book.mixu.net/distsys/single-page.html#abstractions) 深入探讨了抽象和不可能性结论。它以尼采的引言开篇，介绍了系统模型及典型系统模型中的假设。随后讨论了 CAP 定理并总结了 FLP 不可能性结论，然后转向 CAP 定理的影响，其中之一是探索其他一致性模型的必要性，文中讨论了一些一致性模型。

---

### 3. 时间与顺序

理解分布式系统的一大关键是理解时间与顺序。若我们未能正确理解和建模时间，系统就会失败。[第三章](https://book.mixu.net/distsys/single-page.html#time) 探讨了时间与顺序以及时钟的相关内容，以及时间、顺序和时钟的各种用途（如向量时钟和故障检测器）。

---

### 4. 复制：防止分歧

[第四章](https://book.mixu.net/distsys/single-page.html#replication) 介绍了复制问题，以及两种基本的实现方式。事实上，大多数相关特性可以通过这些简单的方式进行讨论。随后，文中探讨了从最低容错性（2PC）到 Paxos 的单副本一致性复制方法。

---

### 5. 复制：接受分歧

[第五章](https://book.mixu.net/distsys/single-page.html#eventual) 探讨了弱一致性保证的复制问题。它引入了一个基本的对账场景，即分区的副本尝试达成一致。然后以 Amazon 的 Dynamo 为例，讨论了弱一致性保证的系统设计。最后，讨论了无序编程的两种视角：CRDT 和 CALM 定理。

---

### 附录

[附录](https://book.mixu.net/distsys/single-page.html#appendix) 提供了进一步阅读的推荐资料。

---

*: This is a [lie](https://en.wikipedia.org/wiki/Statistical_independence). [This post by Jay Kreps elaborates](http://blog.empathybox.com/post/19574936361/getting-real-about-distributed-system-reliability).