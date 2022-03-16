1. 获取etcd-manager的代码
git clone https://github.com/shiguanghuxian/etcd-manage

2. 安装docker和 docker-compose

curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

> docker-compose 1.27 报错 segment错误，换低版本即可
curl -L "https://github.com/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose


3. 修改docker-compose.yml 配置

因为我们只需要管理后台，不需要etcd的安装，需要把安装etcd部分注释掉

### docker-compose.yml
```yaml
version: '3'

services:
  etcd:
   image: "quay.io/coreos/etcd:v3.3"
   container_name: "etcdv3.3"
   environment:
     ETCD_ADVERTISE_CLIENT_URLS: "http://0.0.0.0:2379"
     ETCD_LISTEN_CLIENT_URLS: "http://0.0.0.0:2379"
     ETCDCTL_API: "3"
   volumes: 
     - ./default.etcd:/default.etcd
   ports:
   - 2379:2379
   - 2380:2380
   - 4001:4001
  etcd-manage:
    # build: .
    image: "shiguanghuxian/etcd-manage:1"
    volumes:
      - ./bin/lang/cfg.toml:/app/lang/cfg.toml
      - ./bin/logs:/app/logs
    ports:
      - "10280:10280"
    depends_on:
     - etcd
```

### more bin/config/cfg.toml
```yaml
# debug模式
debug = false
# 日志文件路径
log_path = ""

# http 监听端口
[http]
# 监听地址
address = "0.0.0.0"
# 监听端口
port = 10280

# 使用 Let's Encrypt 证书 - tls_enable为true优先使用本地证书模式
tls_encrypt_enable = false
# 域名列表
tls_encrypt_domain_names = ["shiguanghuxian.com"]

# 是否启用tls
tls_enable = false
# tls证书文件
[http.tls_config]
cert_file = "cert_file"
key_file = "key_file"


## 一下每一个server为一个etcd服务 ##
[[server]]
# 显示名称
title = "pixso_etcd"
# 标识名 - 只能是字母数字或下划线
name = "pixso_etcd"
# etcd连接地址 如果为集群请填写全部地址
#address = ["etcd0:2379","etcd1:2379","etcd2:2379"]
address = ["etcd-1:2379","etcd-2:2379","etcd-3:2379"]
# 查看的key前缀
key_prefix = "/"
# 简述信息
#desc = "docker方式etcd集群方式"
desc = "pixso_etcd"
username="root"
password="pixsoetcd6IfFedsoXY8HGYL"
# 可访问服务器角色列表 - 不写则为所有用户可访问
roles = ["admin"]
# 是否启用tls连接
tls_enable = false
# tls证书配置
[server.tls_config]
cert_file = "/etc/etcd/etcdSSL/etcd.pem"
key_file = "/etc/etcd/etcdSSL/etcd-key.pem"
ca_file = "/etc/etcd/etcdSSL/etcd-root-ca.pem"

#[[server]]
#title = "make docker_run"
#name = "docker_run"
#address = ["etcd:2379"]
#key_prefix = "/"
#desc = "docker方式etcd非集群方式"
#roles = ["admin","dev"]

#[[server]]
#title = "本地etcd"
#name = "local"
#address = ["127.0.0.1:2379"]
#key_prefix = "/"
#desc = "本机环境"
#roles = ["admin","dev"]

## 以下为用户列表 ##
[[user]]
username = "admin"
password = "123456"
role = "admin"

[[user]]
username = "dev_user"
password = "devuser@2021"
role = "dev"
```