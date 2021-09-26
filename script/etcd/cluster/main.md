> 注意替换ip地址！！！
### 节点一
```yaml
version: "3"
services:
  etcd:
    user: "0"
    network_mode: "host"
    privileged: "true"
    image: "aimart.tencentcloudcr.com/base/etcd:v3.3.23"
    container_name: "etcd-01"
    restart: always
    stdin_open: true
    environment:
      TZ: Asia/Shanghai
    tty: true
    command: etcd --name etcd-srv1 --data-dir=/var/lib/etcd/ --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://172.30.0.53:2379 --initial-advertise-peer-urls http://172.30.0.53:2380 --listen-peer-urls http://0.0.0.0:2380 --initial-cluster-token etcd-cluster --initial-cluster "etcd-srv1=http://172.30.0.53:2380,etcd-srv2=http://172.30.0.91:2380,etcd-srv3=http://172.30.0.112:2380" --initial-cluster-state new    
    ports:
      - "2379:2379"
      - "2380:2380"
    volumes:
      - /data/etcd1:/var/lib/etcd
```
 ### 节点二
```yaml
version: "3"
services:
  etcd:
    user: "0"
    network_mode: "host"
    privileged: "true"
    image: "aimart.tencentcloudcr.com/base/etcd:v3.3.23"
    container_name: "etcd-02"
    restart: always
    stdin_open: true
    environment:
      TZ: Asia/Shanghai
    tty: true
    command: etcd --name etcd-srv2 --data-dir=/var/lib/etcd/ --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://172.30.0.91:2379 --initial-advertise-peer-urls http://172.30.0.91:2380 --listen-peer-urls http://0.0.0.0:2380 --initial-cluster-token etcd-cluster --initial-cluster "etcd-srv1=http://172.30.0.53:2380,etcd-srv2=http://172.30.0.91:2380,etcd-srv3=http://172.30.0.112:2380" --initial-cluster-state new    
    ports:
      - "2379:2379"
      - "2380:2380"
    volumes:
      - /data/etcd2:/var/lib/etcd
```
### 节点三
```yaml
version: "3"
services:
  etcd:
    user: "0"
    network_mode: "host"
    privileged: "true"
    image: "aimart.tencentcloudcr.com/base/etcd:v3.3.23"
    container_name: "etcd-03"
    restart: always
    stdin_open: true
    environment:
      TZ: Asia/Shanghai
    tty: true
    command: etcd --name etcd-srv3 --data-dir=/var/lib/etcd/ --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://172.30.0.112:2379 --initial-advertise-peer-urls http://172.30.0.112:2380 --listen-peer-urls http://0.0.0.0:2380 --initial-cluster-token etcd-cluster --initial-cluster "etcd-srv1=http://172.30.0.53:2380,etcd-srv2=http://172.30.0.91:2380,etcd-srv3=http://172.30.0.112:2380" --initial-cluster-state new    
    ports:
      - "2379:2379"
      - "2380:2380"
    volumes:
      - /data/etcd3:/var/lib/etcd
```