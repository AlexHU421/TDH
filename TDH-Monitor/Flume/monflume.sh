#!/bin/bash
###################
#使用说明
# shell名称： flumemonitor.sh 
# 参数： null
# 依赖： node HostKey
# 依赖规则：
#       1、使用workflow节点私钥（记录地址为容器内：/var/lib/workflow/monitorshell/key/keys）
#       2、需要秘钥登录节点/root/.ssh/authorized_keys       记载workflow节点公钥具体内容如下：
#       ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC4fpJuHWS1jXQdrKtybhQqfzFKCYmIoWdHmnzsRsFe6Tvi2/TXub+Emhjnalhokfuu+yYxfVHnwLcamaVsoBd7qy5gpIffSqdyPNvQlW/SukLpEprWO4USyt0/cbqMm3nEpSJ7sNHjJXoc6KdHkbjneijc+Gsh3RTLeChL1zcW2uyNgUFR3qXmV8Gv2QBHyO3aZqzClnj5xzjs27qWkZBqVl1C4MBivAjdTY0umNo0vPRIaNUBNI+bHTZYCpjyf1pkIDqB7ABPe42OrADhMI95QwDRGDdHD1hW+La0p1NDxEPPiWS6DQbujeNxU3a/v9Jmlb2yrMgeADwQ/AzuWpEb root@node11
#       ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCh0IadLvrA5m7IbAgMRItbUNCjoMH2wVCD6ygohnY7fNPUyWII2Fu0dn7w1QwfRuT+Udax5B18Qes13ghJ/jWs2qVKuJxaJwYrlhMF4p5dQHnM+skwI/zNpXA3w8P5Nul8o2B2cn19vzq4zw0PXRAOKqPtS3aqckK/5/ywtBsfMKO6j1MfD79jbceLbCHe3sSNrb3sOEqppl1vwA726TQJ1QEwEVVk9cAzMZoe8z2ZUViXZNSXcqHWiEdFMA3+RcIz+m0YNeU//rFON9FM9/irG7gYMpbUCWcGXqhqYgKIiD73oNKZzPQk5QHeYANYQGwsCZw4Er6HUV7JneuzjsxN root@node10
#       3、登录节点/root/.ssh/authorized_keys文件权限必须为600
#       4、程序本身使用ssh -i 秘钥形式访问目标节点，由于目标节点已经记载正确公钥信息，可以通过私钥以密钥方式访问目标节点
#       5、测试环境已经记载公密钥，并且测试节点保存完整信息可实现免密登录
# 创建记录：
#   制作flume进程监控脚本，用于监控指定列表中flume进程信息（指定变量，多个节点多个进程，创建多个workflow实现）
#
###################
#######程序初始化部分########
node="node38"
thread="flume_airdatacollect_after"
######程序调度

result=`ssh -o "StrictHostKeyChecking no" -o "PasswordAuthentication no" -i /var/lib/workflow/monitorshell/key/keys -t -t root@${node}  ps aux |awk -F ' ' '{print$12}' |grep -v grep |grep ${thread}|wc -l`

if [[  ${result} -gt 0 ]];then
    ssh -o "StrictHostKeyChecking no" -o "PasswordAuthentication no" -i /var/lib/workflow/monitorshell/key/keys -t -t root@${node}  ps aux |grep -v grep |grep ${thread}
    echo "secceeded"
else
    ssh -o "StrictHostKeyChecking no" -o "PasswordAuthentication no" -i /var/lib/workflow/monitorshell/key/keys -t -t root@${node}  ps aux |grep -v grep |grep ${thread}
    echo "ERR"
    exit 1
    
fi



