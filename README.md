# kafka demo

## 探针下载

到洞态网站 Add Agent 页面下载以下两个 agent 放到根目录下：

1. 选择 go, 下载 `dongtai-go-agent-config.yaml`
2. 选择 java, 下载 `dongtai-agent.jar`


## 启动

```
docker-compose up -d
```

## 漏洞触发

命令执行

Java 生产者 -> Java 消费者

* `Runtime.exec()`: http://localhost:8810/kafka/publish?message=whoami
* `ProcessBuilder.start()`: http://localhost:8810/kafka/publish?message=whoami&topic=addUserV3

Go 生产者 -> Java 消费者

* `Runtime.exec()`: http://localhost:8811/kafka/publish?message=whoami
* `ProcessBuilder.start()`: http://localhost:8811/kafka/publish?message=whoami&topic=addUserV3

Web 页面

http://127.0.0.1:8811/kafka

> 如果非本机访问靶场： 项目根目录执行 `sed -i "s#127.0.0.1#服务实际地址#g" golang/index.html`
