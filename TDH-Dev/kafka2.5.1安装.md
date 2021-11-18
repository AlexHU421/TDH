# 1 简介

# 1.1 简介

Kafka is a distributed,partitioned,replicated commit logservice。它提供了类似于JMS的特性，但是在设计实现上完全不同，此外它并不是JMS规范的实现。kafka对消息保存时根据Topic进行归类，发送消息者成为Producer,消息接受者成为Consumer,此外kafka集群有多个kafka实例组成，每个实例(server)成为broker。无论是kafka集群，还是producer和consumer都依赖于zookeeper来保证系统可用性集群保存一些meta信息。

官网： [http://kafka.apache.org/](https://link.segmentfault.com/?url=http%3A%2F%2Fkafka.apache.org%2F)
中文网： [https://www.orchome.com/5](https://link.segmentfault.com/?url=https%3A%2F%2Fwww.orchome.com%2F5)

在其中我们知道如果要kafka正常运行，必须配置zookeeper，否则无论是kafka集群还是客户端的生存者和消费者都无法正常的工作的，使用需要先安装zookeeper集群：

Zookeeper下载地址：[http://zookeeper.apache.org/r...](https://link.segmentfault.com/?url=http%3A%2F%2Fzookeeper.apache.org%2Freleases.html)

## 1.2 参考资料

[http://kafka.apache.org/intro...](https://link.segmentfault.com/?url=http%3A%2F%2Fkafka.apache.org%2Fintro.html)
[https://www.cnblogs.com/likeh...](https://link.segmentfault.com/?url=https%3A%2F%2Fwww.cnblogs.com%2Flikehua%2Fp%2F3999538.html)

# 2 集群规划

## 2.1 机器规划

ip1 : zookeeper + kafka
ip2 : zookeeper + kafka
...
ipn : zookeeper + kafka

鉴于zab原理需求，zookeeper集群建议部署于3-5台节点。

# 2.2 目录规划

```perl
mkdir /tmp/kakfainstall   #临时目录，用于上传压缩包
mkdir /home/kafka  # 安装目录
```

# 3 zookeeper集群

## 3.1 部署环境

操作系统：CentOS7 +
服务器：3-5台，本文档按3台zk集群为前提部署  ip1,ip2,ip3
安装包： kafka_2.12-2.5.1.tgz

前置条件： jdk1.8+

# 3.2 安装zookeeper

### 3.2.1 上传zookeeper安装包

在其中一台服务器上，上传安装包：

```shell
cd  /tmp/kakfainstall   # 跳转到临时目录
```

### 3.2.2 安装

1、解压

```apache
tar  -zxvf  apache-zookeeper-3.5.8-bin.tar.gz  # 解压
mv  apache-zookeeper-3.5.8-bin/  zookeeper-3.5.8  #重命名
mv  zookeeper-3.5.8/  /home/kafka/  # 将解压包拷贝到指定目录
```

2、进入zookeeper的目录

```awk
cd  /home/kafka251/zookeeper-3.5.8/  #进入目录
ll  # 列出目录的文件列表
```

3、修改配置文件
进入conf目录，复制zoo_sample.cfg，并重命名为zoo.cfg

```avrasm
cp  zoo_sample.cfg   zoo.cfg
```

4、打开并修改zoo.cfg

```nginx
vi  zoo.cfg
```

修改dataDir的存放目录， clientPort 默认保持不变，添加zookeeper的节点信息。

```ini
# dataDir 数据存放目录
dataDir=/u01/zkData
# dataLogDir 日志存放目录
dataLogDir=/u01/zkLog
# the port at which the clients will connect
clientPort=2181
# zookeeper 集群的ip:lead和folw的通信端口:选举端口
# server.index 为服务的myid
server.0=ip1:2888:3888
server.1=ip2:2888:3888
server.2=ip3:2888:3888
#tickTime为kafka访问zk元数据的凭证超时时间，对于高吞吐集群、虚拟机集群建议该值调大至9000毫秒，避免引起zk自动超时kill session导致的kafka进程假死
tickTime=9000
```

wq!退出保存后， 创建目录

```shell
#zookeeper
mkdir -p /u01/zkData  #数据目录
mkdir -p  /u01/zkLog  #日志目录
```

在目录中创建myid文件（内容为0）

```shell
touch /u01/zkData/myid
echo 0 > /u01/zkData/myid 
```

5、把 zookeeper文件夹复制到另外两个服务器上，并创建相同的dataDir目录和myid文件。

```shell
scp -r /home/kafka/zookeeper-3.5.8  root@ip2:/home/kafka/

scp -r /home/kafka/zookeeper-3.5.8  root@ip3:/home/kafka/
```

注意：ip2的 myid 文件内容为 1，ip3的 myid 文件内容为 2，需要与zoo.cfg中的保持一致。

```shell
# ip2执行
mkdir -p /u01/zkData/data
touch /u01/zkData/myid
echo 1 > /u01/zkData/myid  

# ip3执行
mkdir -p /u01/zkData/data
touch /u01/zkData/myid
echo 2 > /u01/zkData/myid  
```

6、关闭防火墙

```shell
# 关闭防火墙
service iptables stop 
# 查看防火墙状态
service iptables status
```

### 3.2.3 设置环境变量

设置环境变量

```shell
# 打开profile
vim /etc/profile
#加入以下内容
export ZK_HOME=/home/kafka/zookeeper-3.5.8
export PATH=$ZK_HOME/bin:$PATH
#保持退出后，执行以下命令，立即生效
source /etc/profile
```

### 3.2.4 启动

分别在三台机器上启动zookeeper：

```shell
/home/kafka/zookeeper-3.5.8../bin/zkServer.sh  start
```

启动成功后：
ZooKeeper JMX enabled by default
Using config: /home/kafka/zookeeper-3.5.8/bin/../conf/zoo.cfg
Starting zookeeper ... STARTED

### 3.2.5 验证

在三台服务器上执行命令，查看启动状态：

```shell
/home/kafka/zookeeper-3.5.8/bin/zkServer.sh status
```

可以查看服务器的角色：leader / follower
如：
ZooKeeper JMX enabled by default
Using config: /home/kafka/zookeeper-3.5.8/bin/../conf/zoo.cfg
Mode: follower

# 4 kafka集群

## 4.1 部署环境

操作系统：CentOS7 +
服务器：3-5台，本文档按3台zk集群为前提部署  ip1,ip2,ip3
安装包： kafka_2.12-2.5.1.tgz

## 4.2 安装kafka

### 4.2.1 上传kafka安装包

在其中一台服务器上，上传安装包：

```shell
cd  /tmp/kakfainstall/   # 跳转到临时目录
```

4.2.2 安装
1、解压

```shell
tar  -zxvf  kafka_2.12-2.5.1.tgz  # 解压
mv  kafka_2.12-2.5.1  kafka251  #重命名
mv  kafka251/  /home/kafka/  # 将解压包拷贝到指定目录

```

2、修改server.properties
进入kafka的config目录

```shell
cd  /home/kafka251/kafka/config
vi  server.properties  #修改配置

```

修改内容如下：

```ini
# broker的id
broker.id=0
#指定本机ip
host.name=ip1
#增加网络及IP处理线程数
num.network.threads=10
num.io.threads=10


#设置单个topic的默认partition值，建议根据kafka集群节点数分配对应分区数
num.partitions=3


#增大网络、消息、日志单条消息体最大数据上限，防止无法处理前方生产者的超长单条数据
socket.send.buffer.bytes=6525000
message.max.bytes =102400000
log.flush.interval.messages=20000
#增加网络、消息、日志处理的超时时间阈值
log.flush.interval.ms=2000
replica.fetch.wait.max.ms=6000
queued.max.requests=2000

#禁止自动创建topic
auto.create.topics.enable=false

# 修改日志目录
log.dirs=/u03/kafkaLog_0.10

# 连接zookeeper集群，若zokeeper集群申请域名，此处可直接填写dns:2181
zookeeper.connect=ip1:2181,ip2:2181,ip3:2181
# 可删除topic
delete.topic.enable=true
```

3、创建对应的日志目录

```shell
mkdir  -p  /u03/kafkaLog_0.10
```

4、将kafka文件夹复制到另外的两个节点上

```shell
scp  -r  /home/kafka/kafka251/ root@ip2:/home/kafka/kafka251/  # 远程复制

scp  -r  /home/kafka/kafka251/ root@ip3:/home/kafka/kafka251/  # 远程复制
```

5、创建相同的logDir目录，且修改每个节点对应的server.properties文件的broker.id、hostname:
ip2 broker.id为 1
ip2 host.name为 ip2
ip3 broker.id为 2
ip3  host.name为 ip3

```shell
mkdir  -p /u03/kafkaLog_0.10 # 另外两台上创建目录
vim   /home/kafka/kafka251/config/server.properties 

#打开文件修改对应的broker.id、host.name
#ip2
broker.id=1
host.name=ip2

#ip3
broker.id=2
host.name=ip3
```

### 4.2.3 设置环境变量

设置环境变量

```ini
# 打开profile
vim /etc/profile
#加入以下内容
export KAFKA_HOME=/home/kafka251/kafka
export PATH=$KAFKA_HOME/bin:$PATH
#保持退出后，执行以下命令，立即生效
source /etc/profile
```

### 4.2.4 启动kafka

**注：启动时：先启动 zookeeper，后启动 kafka；关闭时：先关闭 kafka，后关闭zookeeper**

1、分别在每个节点上执行命令，启动zookeeper，如果已经启动或，则跳过

```shell
/home/kafka/zookeeper-3.5.8/bin/zkServer.sh  start
```

2、启动kafka

```shell
nohup /home/kafka251/kafka/bin/kafka-server-start.sh  /home/kafka251/kafka/config/server.properties  > /u02/kafka.log 2>&1 &
```

## 4.3 测试

### 4.3.1 创建topic

```shell
#无域名
/home/kafka251/kafka/bin/kafka-topics.sh --create  --bootstrap-server  kafka-test-206110:9092,kafka-test-206111:9092,kafka-test-206112:9092 --replication-factor 3 --partitions 3 --topic test

#有域名
/home/kafka/kafka251/bin/kafka-topics.sh --create --bootstrap-server dns:2181 --replication-factor 3 --partitions 3 --topic test
```

### 4.3.2 显示topic信息

```shell
#无域名
/home/kafka251/kafka/bin/kafka-topics.sh --describe --bootstrap-server  kafka-test-206110:9092,kafka-test-206111:9092,kafka-test-206112:9092 --topic test
 
 #有域名
/home/kafka/kafka251/bin/kafka-topics.sh --describe --bootstrap-server  dns:2181 --topic test
```

### 4.3.3 列出topic

```shell
#无域名
/home/kafka251/kafka/bin/kafka-topics.sh --list  --bootstrap-server kafka-test-206110:9092,kafka-test-206111:9092,kafka-test-206112:9092

#有域名
/home/kafka/kafka251/bin/kafka-topics.sh --list --bootstrap-server dns:2181
```

### 4.3.4 创建生产者

在master节点上 测试生产消息

```shell
#无域名
/home/kafka/kafka251/bin/kafka-console-producer.sh  --broker-list  ip1:9092,ip2:9092,ip3:9092  -topic  test

#有域名
/home/kafka/kafka251/bin/kafka-console-producer.sh  --broker-list  dns:9092   -topic  test
```

### 4.3.5 创建消费者

在worker节点上 测试消费

```shell
#无域名
/home/kafka/kafka251/bin/kafka-consumer-groups.sh  --bootstrap-server ip1:9092,ip2:9092,ip3:9092  -topic test --from-beginning

#有域名
/home/kafka/kafka251/bin/kafka-consumer-groups.sh  --bootstrap-server dns:9092  -topic test --from-beginning
```

### 4.3.6查看偏移量属性

```shell
#无域名
/home/kafka/kafka251/bin/kafka-consumer-groups.sh  --bootstrap-server ip1:9092,ip2:9092,ip3:9092 --describe --group  testgroup

#有域名
/home/kafka/kafka251/bin/kafka-consumer-groups.sh  --bootstrap-server dns:9092  --topic test --describe --group  testgroup
```



### 4.3.7 删除topic和关闭服务

删除topic

```shell
#无域名
/home/kafka/kafka251/bin/kafka-topics.sh --delete --bootstrap-server  ip1:9092,ip2:9092,ip3:9092  --topic test

#有域名
/home/kafka/kafka251/bin/kafka-topics.sh --delete --bootstrap-server  dns:9092 --topic test
```

关闭kafka服务，在三台机器上执行kafka-server-stop.sh命令：

```shell
/home/kafka/kafka251/bin/kafka-server-stop.sh
```

关闭zookeeper:

```shell
/home/kafka/zookeeper-3.5.8/bin/zkServer.sh  stop
```
