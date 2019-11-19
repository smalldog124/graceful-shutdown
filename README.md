# GRACEFULL-SHUTDOWN

เป็น Code ตัวอย่างสำหรับการทำ service gracefull shutdow โดย service จะมี timeout 5 วินาที สำหรับทำ prosecess ที่ค้างอยู่ให้เสร็จ code อ้างอิงจาก [gin-gonic/examples](https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go)

## สิ่งที่ต้องมีเพื่อ run service

```text
1. golang version 1.10+
2. docker
```

## วิธีการใช้งาน(เฉพาะ linux os)

1. build docker images
```text
make build
```

2. run docker (no port :3000)
```text
make run
```

3. ดู log service
```text
make logservice
```

4. load request เข้าไปยัง service
```text
make load-request
```

5. stop service ที่ run อยู่
```text
make stop
```

## ผลที่ได้

log service
```log
2019/11/19 21:03:48 save to database
[GIN] 2019/11/19 - 21:03:48 | 200 |    2.0101757s |      172.17.0.1 | GET      /ping
2019/11/19 21:03:52 save to database
[GIN] 2019/11/19 - 21:03:52 | 200 |    2.0078826s |      172.17.0.1 | GET      /ping
2019/11/19 21:03:56 save to database
[GIN] 2019/11/19 - 21:03:56 | 200 |    2.0090411s |      172.17.0.1 | GET      /ping
2019/11/19 21:04:00 save to database
[GIN] 2019/11/19 - 21:04:00 | 200 |    2.0040282s |      172.17.0.1 | GET      /ping
2019/11/19 21:04:03 Shutdown Server ...
2019/11/19 21:04:04 save to database
[GIN] 2019/11/19 - 21:04:04 | 200 |    2.0063833s |      172.17.0.1 | GET      /ping
2019/11/19 21:04:04 Program stopped successful
```

log load request
```log
2019/11/19 21:03:46 start request
2019/11/19 21:03:48 resuest succesfull
2019/11/19 21:03:50 start request
2019/11/19 21:03:52 resuest succesfull
2019/11/19 21:03:54 start request
2019/11/19 21:03:56 resuest succesfull
2019/11/19 21:03:58 start request
2019/11/19 21:04:00 resuest succesfull
2019/11/19 21:04:02 start request
2019/11/19 21:04:04 resuest succesfull
2019/11/19 21:04:06 start request
2019/11/19 21:04:06 request error:  Get http://localhost:3000/ping: dial tcp [::1]:3000: connect: connection refused
```