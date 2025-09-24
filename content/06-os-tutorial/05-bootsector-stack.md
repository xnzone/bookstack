---
author: cfenollosa  
title: Bootsector Stack
date: 2025-01-12 10:04:00
image: https://s2.loli.net/2025/09/24/mADxctyj3rV8LTq.jpg
cover: false
weight: 605
tags: ["os", "tutorial"]
---

*Concepts you may want to Google beforehand: stack*

**Goal: Learn how to use the stack**

The usage of the stack is important, so we'll write yet another boot sector
with an example.

Remember that the `bp` register stores the base address (i.e. bottom) of the stack,
and `sp` stores the top, and that the stack grows downwards from `bp` (i.e. `sp` gets
decremented)

This lesson is quite straightforward, so jump ahead to the code.

I suggest that you try accessing in-stack memory addresses by yourself, 
at different points in the code, and see what happens.