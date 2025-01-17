---
author: cfenollosa  
title: Environment
date: 2025-01-12 10:04:00
image: /covers/os-tutorial.jpg
cover: true
weight: 601
tags: ["os", "tutorial"]
---

*Concepts you may want to Google beforehand: linux, mac, terminal, compiler, emulator, nasm, qemu*

**Goal: Install the software required to run this tutorial**

I'm working on a Mac, though Linux is better because it will have all the standard tools already
available for you.

On a mac, [install Homebrew](http://brew.sh) and then `brew install qemu nasm`

Don't use the Xcode developer tools `nasm` if you have them installed, they won't work for the most cases. Always use `/usr/local/bin/nasm`

On some systems qemu is split into multiple binaries. You may want
to call `qemu-system-x86_64 binfile`