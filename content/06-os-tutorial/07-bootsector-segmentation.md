---
author: cfenollosa  
title: Bootsector Segmentation
date: 2025-01-12 10:04:00
image: https://s2.loli.net/2025/09/24/mADxctyj3rV8LTq.jpg
cover: false
weight: 607
tags: ["os", "tutorial"]
---

*Concepts you may want to Google beforehand: segmentation*

**Goal: learn how to address memory with 16-bit real mode segmentation**

If you are comfortable with segmentation, skip this lesson.

We did segmentation
with `[org]` on lesson 3. Segmentation means that you can specify
an offset to all the data you refer to.

This is done by using special registers: `cs`, `ds`, `ss` and `es`, for
Code, Data, Stack and Extra (i.e. user-defined)

Beware: they are *implicitly* used by the CPU, so once you set some
value for, say, `ds`, then all your memory access will be offset by `ds`.
[Read more here](http://wiki.osdev.org/Segmentation)

Furthermore, to compute the real address we don't just join the two
addresses, but we *overlap* them: `segment << 4 + address`. For example,
if `ds` is `0x4d`, then `[0x20]` actually refers to `0x4d0 + 0x20 = 0x4f0`

Enough theory. Have a look at the code and play with it a bit.

Hint: We cannot `mov` literals to those registers, we have to
use a general purpose register before.