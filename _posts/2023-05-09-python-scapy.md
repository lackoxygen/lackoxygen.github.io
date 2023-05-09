---
layout: post
title: python scpay
subtitle: scpay抓包跟解析包内容
tags: [ python ]
comments: true
---

### 使用`scpay`捕获、分析、操作网络数据包

**安装**

```
pip3 install scapy
```

**抓包**

- 必须使用管理员启动，否则会导致权限不足
- `iface` 网卡名称
- `count` 抓包次数，0不限制次数
- `filter` 过滤数据包
- `prn` 每个数据包的回调函数
- `store` 捕获的数据放入内存中
- `stop_filter` 返回True 停止抓包

```
from scapy.all import *

packets = sniff(
    iface=iface,
    count=0,
    filter='tcp',
    prn=packet_callback,
    store=True,
    stop_filter=stop_callback
)

def packet_callback(packet):
    print(packet.summary())

def stop_callback(packet):
    return packet.haslayer(TCP)

```

**发送网络数据包**
```
from scapy.all import *

# 构造数据包
packet = IP(dst="www.google.com")/ICMP()

# 发送数据包
send(packet)

```

**构造网络数据包**
```
from scapy.all import *

# 构造数据包
packet = IP(src="192.168.1.1", dst="192.168.1.2")/TCP(sport=1234, dport=80)/"GET / HTTP/1.1\r\nHost: www.google.com\r\n\r\n"

# 发送数据包
send(packet)
```

**解析数据包**

```
from scapy.all import *
from scapy.layers.l2 import Ether

from scapy.layers.inet import IP
from scapy.layers.inet import TCP
from scapy.layers.inet import UDP
from scapy.layers.inet import ICMP
from scapy.layers.inet6 import IPv6

sniff(prn=parse_packet, iface='en0', count=1, filter="tcp port 80")

def parse_packet(pack: Ether):
    if pack.haslayer(IP):  # ip
        print("ip")
        print(pack.getlayer(IP))
    if pack.haslayer(TCP):  # tcp
        print("tcp")
    if pack.haslayer(UDP):  # udp
        print("udp")
    if pack.haslayer(ICMP):  # icmp
        print("icmp")
    if pack.haslayer(IPv6):  # ipv6
        print('ipv6')

    if pack.haslayer(Raw):  # 获取数据包
        print(pack.getlayer(Raw))

    if pack.haslayer(Padding):  # 数据对齐
        print(pack.getlayer(Padding))

```

**raw数据包不包括以下内容**

- Frame
- Ethernet
- Internet Protocol
- Transmission control Protocol