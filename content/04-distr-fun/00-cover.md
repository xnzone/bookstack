---
author: xnzone 
title: 趣谈分布式系统 
date: 2022-09-10 10:23:32
image: /covers/distr-fun.png
cover: true
weight: 4
tags: ["distributed system", "CAP", "分布式系统"]
---

## 介绍

>翻译自[Distributed System For Fun & Profit](http://book.mixu.net/distsys/)

我希望有一个文章能够尽可能多涵盖最近分布式系统的内容。例如Amazon's Dynamo，Google's BitTable 和 MapReduce, Apache's Hadoop等

这个文章，我已经尝试提供一个更容易理解的分布式系统。对我而言，这意味着两件事：介绍一些阅读更多高级文章时需要的关键概念，提供一个涵盖足够细节的叙述而不会陷入更深的细节中。现在是2013年了，你已经可以联网了，你可以选择性阅读更多你感兴趣的话题

在我看来，分布式程序主要处理两个分布式的结果
- 信息以光速传递
- 独立的业务独立失败

换句话说，分布式程序核心是处理距离和多件事。这些约束限制了系统设计空间，我希望在阅读完这本书，你可以对距离，时间和一致性模型有更好的理解

这个文章专注于分布式程序和系统的概念，这些概念，你需要在商业数据中心系统中理解。涵盖所有的事情可能会比较疯狂。你将会学习一些关键的协议和算法(例如涵盖该学科中被很多论文引用的)，包括一些最终一致性的解决方案，这也被写进了一些大学课程的教材中，比如CRDT和CALM理论

我希望你能够喜欢它，如果你想说谢谢，关注我的Github或者Twitter。如果你发现了错误，请在Github上提交一个pull request

-------------------------------------------------------------------------------------
### 1. 概述
第一章包括高层次的分布式系统重要术语和概念。它包含一个高级目标，比如伸缩性，可用性，性能和容灾；这些是如何难以实现，以及抽象、模型、分区和复制是如何发挥作用的

### 2. 上下层抽象
第二章更深入介绍抽象和不可能的结果。从Nietzsche引论开始，然后介绍系统模型和经典系统模型中的大量抽象。然后讨论CAP理论和概括FLP不可能结果。最后探讨CAP理论的实现，这是浏览其他一致性模型时应该了解的。还有一些一致性模型也会被讨论

### 3. 时序
理解分布式系统一部分就是关于理解时序。如果我们无法理解和模拟时间，我们的系统将是失败的。第三章讨论时间，顺序和时钟，同时也包括他们的使用(例如向量时钟和失败探测)

### 4. 复制：防止分歧
第四章介绍复制问题，以及可以执行的两种方法。已经证实大多数相关特征都可以用这个简单表征来讨论。然后，讨论维护副本一致性的方法，从最小容错(2PC)到Paxos

### 5. 复制：接受分歧
第五章讨论副本的弱一致性保证。它介绍了一个基本的妥协场景，分区副本可以达到最终一致性。然后用Amazon's Dynamo作为例子，讨论弱一致性的实现。最后，讨论无序编程的两种观点:CRDTs和CALM理论

### 附录
附录包括进一步阅读建议

