---
author: xnzone 
title: Docker
date: 1906-01-01
image: https://s2.loli.net/2025/09/24/B9kM2IQLyPN4DRw.jpg
cover: false 
weight: 1202
tags: ["xv6", "os"]
---

## Dockerfile

{{< highlight Dockerfile >}}
FROM ubuntu:16.04

RUN apt-get -qq update

RUN apt-get install -y git build-essential gdb gcc-multilib tmux

RUN git clone http://web.mit.edu/ccutler/www/qemu.git -b 6.828-2.3.0

RUN apt-get install -y libsdl1.2-dev libtool-bin libglib2.0-dev libz-dev libpixman-1-dev

RUN cd qemu && ./configure --disable-kvm --target-list="i386-softmmu x86_64-softmmu" && make && make install && cd ..

ADD ./jos jos

WORKDIR jos

CMD ["/bin/bash"]
{{< /highlight >}}

## startup

{{< highlight bash >}}
joswd=$(pwd)
echo ${joswd}
docker run --rm -it -v ${joswd}/jos:/jos xv6
{{< /highlight >}}