## Description

[rasberry-dashboard](https://github.com/yangqi93/raspberry-dashboard) 参考了 [rpi-dashboard](https://github.com/nxez/pi-dashboard) 的设计和 WebUi，使用了 Go 语言的方式实现了相同的功能。

Pi Dashboard (Pi 仪表盘) 是一个开源的 IoT 设备监控工具，目前主要针对树莓派平台，也尽可能兼容其他类树莓派硬件产品。你只需要在树莓派上安装好 Docker 环境，即可方便的部署一个 Pi 仪表盘，通过炫酷的 WebUI 来监控树莓派的状态！

目前已加入的监测项目有：

- CPU 基本信息、状态和使用率等实时数据

- 内存、缓存、SWAP分区使用的实时数据

- SD卡（磁盘）的占用情况

- 实时负载数据

- 实施进程数据

- 网络接口的实时数据

- 树莓派IP、运行时间、操作系统、HOST 等基础信息

## Rreview

![image](https://user-images.githubusercontent.com/14936391/236113344-cfcd72ab-9c54-40fa-84ec-6e1e62d91491.png)

## Install
1. touch conf/config.yaml
2. vim config.yaml
```yaml
     hostName: xxx.xxx.com
```
3. vim docker-compose.yaml
```yaml
version: "3.7"
services:
  app:
    image: yangqigo/go-pi-dashboard
    restart: always
    volumes:
      - ./conf:/conf
      - ./logs:/logs
    ports:
      - "1024:9001" 
```
4. docker-compose up -d
5. open http://yourip:1024
