---
layout: post
title: vmware-workstation挂载目录
subtitle: vmware-workstation挂载目录
tags: [vmware]
comments: true
---

**执行挂载命令**

{: .box-warning}
**挂载的文件夹必须是空目录.**
```
sudo vmhgfs-fuse .host:/ /mnt/hgfs
```

**设置开机自动挂载**

```
# 编辑fatab
sudo vim /etc/fstab

# 增加加一行，把主机的web目录挂载到虚拟机的web目录 
host:/web   /web   fuse.vmhgfs-fuse   allow_other,uid=1000,gid=1000,umask=022   0   0
```

