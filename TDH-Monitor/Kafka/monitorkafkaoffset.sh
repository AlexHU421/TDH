#!/bin/bash
###################
#使用说明
# 
# 
# shell名称： MonitorkafkaTopic.sh 
# 参数： null
# 依赖： /home/TDH-Client/  必须保存tdh客户端
# 创建记录：
#   制作kafka监控脚本，判断值确定kafka是否处于数据积压异常状态
#
###################
#######程序初始化部分########
#source /var/lib/workflow/TDH-Client/init.sh


#######函数声明部分########
#######调度 部分###################


#######函数变量声明部分########


function  report()
{
da=`date +"%Y-%m-%d"`
ts=`date +"%H:%m:%S"`
##待确认正确发送连接
serverclient="http://192.168.210.201:8780/bell/message/send"
##待确认正确连接登录项目名称，用户名密码
#projectCode='TDHMonitorMemory'
projectCode='TDHMonitorFlinkKafkaOffset'
username='TDH'
password='TDHPassword1'
subject=$1
content=`echo $2 |sed 's/ /\\\\r/g'`"\r:<br>"$3
mail="huyicai@inner.sss.com"

curl -X POST -H "Content-Type: application/json" -d '{"auth": {"projectCode": "'${projectCode}'","username": "'${username}'","password": "'${password}'"},"messages": [{"ower": "","subject": "'${subject}'","details": [{"type": "E","content":"'${content}'","fromUser": "","toUsers": "'${mail}'","sendMode": "I","sendTime": "'${da}" "${ts}'"}]}]}' ${serverclient}


echo "\r"

echo $3

echo $1 $2

}

#######变量声明部分########
bootstrapserver="10.16.32.92:9092,10.16.32.93:9092,10.16.32.94:9092"          ##线上
#bootstrapserver="192.168.210.127:2181,192.168.210.128:2181,192.168.210.129:2181"          ##测试
group="flink_price_monitor_group"      ##线上
#group="flink_u_ogg_group"               ##测试
tmpdir=`date +"%Y%m%d"`
tmpdir2=`date +"%H%M%S"`
tmpfile="/tmp/monitorkafkaoffset/${tmpdir}/${tmpdir2}/offset.tmp"
monitorcount=$1
allmonitorcount=$2
#10.16.32.92:9091,10.16.32.93:9091,10.16.32.94:9092



######修改修改连接脚本方式

mkdir -p /tmp/monitorkafkaoffset/${tmpdir}/${tmpdir2}
sh /var/lib/workflow/TDH-Client/kafka/bin/kafka-consumer-groups.sh --new-consumer --bootstrap-server ${bootstrapserver} --describe --group ${group} |sed '/^$/d' > ${tmpfile}
#sh /home/TDH-Client/kafka/bin/kafka-consumer-groups.sh --zookeeper ${bootstrapserver} --describe --group ${group} |sed '/^$/d' > ${tmpfile}
row=`wc -l ${tmpfile} |awk -F ' ' '{print $1}'`
list=(`cat ${tmpfile}`)
onerow=`cat ${tmpfile}  |sed ':label;N;s/\n/<br>/;b label'|sed 's/[ ][ ]*/\\\\r/g'`



#######调度 部分###################


if [[ x${list[0]} = x"TOPIC" && x${list[1]} = x"PARTITION"  && x${list[2]} = x"CURRENT-OFFSET"  && x${list[3]} = x"LOG-END-OFFSET"  && x${list[4]} = x"LAG"  && x${list[5]} = x"CONSUMER-ID" ]];then
  
    if [[ $row -gt 2 ]];then
        
        sed -n '1!P;N;$q;D' ${tmpfile}
##        cat  /tmp/monitorkafkaoffset/lizi | while read LINE
        while read LINE
        do
            tmpcount=`echo $LINE|awk -F ' ' '{print $5}'`       #真实使用$5
            #tmpcount=`echo $LINE|awk -F ' ' '{print $4}'`        #模拟测试使用$4
            if [[ ${tmpcount} -gt ${monitorcount} ]];then
                #report "OnePartitionHeigt:" "monitorKafkaOffsetOnePartitionIsHeight" ${onerow}
                echo "monitor kafka offset Warn: One Partition is Height"
                rm -rf /tmp/monitorkafkaoffset/${tmpdir}
                echo ${onerow}
                exit 1
            fi
            let count+=tmpcount
            #echo $count
            
        done < ${tmpfile}


    else
        echo "monitor kafka offset is wround : data err"
        rm -rf /tmp/monitorkafkaoffset/${tmpdir}
        exit 1
    fi

else
    echo "monitor kafka offset is wround : from err"
    rm -rf /tmp/monitorkafkaoffset/${tmpdir}
    exit 1
fi

echo "offsetCount:"${count}

if [[ ${count} -gt ${allmonitorcount} ]];then
#report "TotalPartitionsHeigt:" "monitor Kafka Offset  Warn:  Total Partitions Height" ${onerow}
echo "monitor Kafka Offset  Warn:  Total Partitions Height"
rm -rf /tmp/monitorkafkaoffset/${tmpdir}
echo ${onerow}
exit 1
else
echo "succeeded"
echo ${onerow}
rm -rf /tmp/monitorkafkaoffset/${tmpdir}
fi

