#!/bin/sh
###################
#使用说明
# 路径： /var/lib/workflow/monitorshell/
# shell名称： inceptor_health_check.sh
# 参数： $1 : inceptorServer
#        $2 : timeout
# 
# 输出：0 1
# 创建记录：
#               rc0: 2020年8月31日 监控inceptorServer变量值的inceptor服务启用心跳连接测试，限定timout变量值之内为服务正常，反之则异常脚本错误退出状态码1
#
# 修改记录：
#
###################



#######脚本初始化部分########
#######变量声明部分########
checkstatus=0
#获取并转换脚本执行传入参数
inceptorServer=$1
time=$2
#创建inceptor server list
##生产
serverlist=("node18" "node48" "node63" "node67")
##测试
#serverlist=("tdh2")

#######调度 部分###################
#循环inceptor server list判断需要监控的inceptor server是否在list中，若存在检查状态变量值变为1
for i in ${serverlist[@]};do
   if [[ "$i" == "${inceptorServer}" ]];then
     checkstatus=1
   fi
done
#检查上述循环检查结果，若检查状态变量值为1则脚本继续，否则报错退出 状态码 1
if [[ ${checkstatus} = 1 ]];then
echo "Server check is OK"
else
echo "Server check is ERROR This Server ${inceptorServer} is not in Serverlist"
exit 1
fi


#检查超时监控阈值是否为大于零数值，若非大于零或空或非数值，则报警告，并设置默认监控超时阈值为30秒
if [[ ${time} -gt 0 ]];then
echo "Check that the server timeout is within ${time} seconds of this connection"
else
time=30
echo "Time value small and 0 are set to the default value of 30s,Check that the server timeout is within ${30} seconds of this connection"
fi

#执行心跳访问 查询的default.aaa表，测试环境无数据，生产环境1条数据。实际后续若需检查executor节点数执行效率，可后续优化
timeout ${time} beeline -u jdbc:hive2://${inceptorServer}:10000 -n hive -p 123456 -e "SELECT * FROM default.aaa;"

#判断执行心跳结果返回值，若执行失败或执行超时 timout执行结果返回状态码124，若为124则报错退出。反之正常脚本结束
if [[ $? -eq '124' ]];then
echo "Timeout! Server ${inceptorServer} check this connection for more than ${time} seconds"
exit 1
else
echo "Server ${inceptorServer} check this connection is OK"
fi

