---
layout: post
title: 如何使用 javascript 发送浏览器通知🔔
subtitle: Window Notification API
tags: [javascript]
comments: true
---

**Notification API是存在window对象中，检查浏览器是否支持这个API**
```
    if (!window.Notification) {
        throw '浏览器不支持通知';
    }
    console.log('浏览器支持通知');
```

**使用requestPermission获取通知授权**
```
    //granted是允许，denied是拒绝
    Notification.requestPermission().then((permission) => {
    if ('denied' === permission) {
        throw '权限已拒绝!';
    }

    console.log('权限已获取 -> ' + permission)
})
```

**发送消息前可以先判断是否获得权限**
```
    if (window.Notification && Notification.permission === "granted") {
        console.log('已获得通知权限');
    }
```

**发送消息**
```
//消息就是Notification本身
const n = new Notification("hello world");

//5秒后关闭
setTimeout(n.close.bind(n), 5000);
```

**完整示例**

```
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Window Notification</title>
</head>
<body>
<div>
    <button id="get-notify-access-bth">Get Notionfication Access</button>
    <button id="send-notify-bth">Send Notionfication</button>
</div>
</body>
<script>
    const requestBth = document.getElementById('get-notify-access-bth');
    const sendBth = document.getElementById('send-notify-bth');

    requestBth.addEventListener('click', function (e) {
        e.preventDefault();
        if (!window.Notification) {
            throw '浏览器不支持通知';
        }
        Notification.requestPermission().then((permission) => {
            if ('denied' === permission) {
                throw '权限已拒绝!';
            }

            console.log('权限已获取 -> ' + permission)
        })
    });

    sendBth.addEventListener('click', function (e) {
        e.preventDefault();

        if (window.Notification && Notification.permission === "granted") {
            const n = new Notification("hello world");

            setTimeout(n.close.bind(n), 5000);
        }
    })

</script>
</html>
```

{: .box-note}
**谷歌浏览器无反应**

开启浏览器通知权限
![](/assets/img/posts/mac-enable-notify.png)