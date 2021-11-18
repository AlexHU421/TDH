# Kafka 2.5.0及zookeeper 安装

## 系统环境检查及安装

##### 主机名制作

```shell
hostnamectl set-hostname kafka3250
hostnamectl set-hostname kafka3251
hostnamectl set-hostname kafka3252
hostnamectl set-hostname kafka3253
hostnamectl set-hostname kafka3254
hostnamectl set-hostname kafka3255

```

##### 所有节点制作host文件

```shell
echo "10.16.32.50 kafka3250" >> /etc/hosts
echo "10.16.32.51 kafka3251" >> /etc/hosts
echo "10.16.32.52 kafka3252" >> /etc/hosts
echo "10.16.32.53 kafka3253" >> /etc/hosts
echo "10.16.32.54 kafka3254" >> /etc/hosts
echo "10.16.32.55 kafka3255" >> /etc/hosts

```

##### 临时节点制作免密登录 ES3259

```shell
ssh-keygen -t rsa

Generating public/private rsa key pair.
Enter file in which to save the key (/root/.ssh/id_rsa):
/root/.ssh/id_rsa already exists.
Overwrite (y/n)? 

y

Enter passphrase (empty for no passphrase):

Enter same passphrase again:

Your identification has been saved in /root/.ssh/id_rsa.
....

```

##### 所有节点添加临时节点公钥

```shell
 vi /root/.ssh/authorized_keys

o

ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCzU4f/XP7HoI4iykslPlfQIr2vM01R04LptCWUc7GY8OAnBK6S15nSuXY1qsaB1lSb9NY66gFbBE5ZSWrxLl7ORCfZk7wTjE/nCaeZUhw6zjQChwnTN1ZTq240G1oBQ1MK9yxkYOyVltIbXbeTyzb4+z3+bUACxUFc+BEugT4ir/PFkwBG4ikFoX6BE5MaL1T/DZF+/kDPorcFzKO/941bivrVmm8eSSfouC+0dlyWvoR2wdrMU20Cg7OZYgjgwH+LvRJIV0a35KEFtW68SE9YpIPMQLmY383CoYdjtfdsIY4WO4wXTIpbSHxKnbMZ1Q4pp17DsbL8wv0sZlrToLw5 root@kafka3250



:wq


```

##### JDK 1.8安装

```shell
tar -zxvf jdk-8u241-linux-x64.tar.gz
mkdir -p /usr/java
mv jdk1.8.0_241 /usr/java/

yum -y install vim

vim /etc/profile

JAVA_HOME=/usr/java/jdk1.8.0_241
PATH=$PATH:$JAVA_HOME/bin
CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar

export JAVA_HOME
export PATH
export CLASSPATH



source  /etc/profile


java -version
```

pssh批量操作工具安装

```shell
yum -y install wget
 wget ftp://ftp.pbone.net/mirror/ftp5.gwdg.de/pub/opensuse/repositories/home:/matthewdva:/build:/EPEL:/el7/CentOS_7/noarch/pssh-2.3.1-5.el7.noarch.rpm
rpm -ivh pssh-2.3.1-5.el7.noarch.rpm

```



192.168.1.101:2181,192.168.1.102:2181,192.168.1.103:2181

kafka1:zookeeper.connect=192.168.1.101:2181,192.168.1.102:2181,192.168.1.103:2181/kafka_namespace1

kafka1:zookeeper.connect=192.168.1.101:2181,192.168.1.102:2181,192.168.1.103:2181/kafka_namespace1

