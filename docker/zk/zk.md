# Zookpeer

## 环境准备

- 执行make brun
- 启动客户端

```bash
docker run -it --rm \
    --link zk1.local:zk1 \
    --link zk2.local:zk2 \
    --link zk3.local:zk3 \
    --net zk_default \
    zookeeper:3.6 zkCli.sh -server zk1:2181,zk2:2181,zk3:2181
```

## 命令

- create
