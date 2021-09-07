
drop  STREAM stream.STREAM_STG_TDH_HARDWARE_MEM_INFO_step0;
CREATE STREAM stream.STREAM_STG_TDH_HARDWARE_MEM_INFO_step0 (
KEY String
,HOSTNAME STRING
,HARDYINFO String
,Oiginal_Unixtime String
,Metrics_DAY String
,Metrics_TIME String
,Oiginal_Value String
,Metrics_Value_Ratio STRING)
ROW FORMAT SERDE 'org.apache.hadoop.hive.contrib.serde2.MultiDelimitSerDe'
WITH SERDEPROPERTIES ('input.delimited'='|+|') 
TBLPROPERTIES(
"topic"="tdh-hardwareinfo",
"kafka.zookeeper"="XX.XX.XX.XX:2181/kafak1",
"kafka.broker.list"="XX.XX.XX.XX:9192",
"transwarp.consumer.group.id"="MEM_INFO");

drop  STREAM stream.STREAM_STG_TDH_HARDWARE_MEM_INFO_step1;
CREATE STREAM stream.STREAM_STG_TDH_HARDWARE_MEM_INFO_step1 AS 
SELECT 
KEY,HOSTNAME,HARDYINFO,Oiginal_Unixtime,Metrics_DAY,Metrics_TIME,Oiginal_Value,Metrics_Value_Ratio
FROM stream.STREAM_STG_TDH_HARDWARE_MEM_INFO_step0
WHERE HARDYINFO='meminfo'
;

DROP  streamjob STREAMJOB_STG_TDH_HARDWARE_MEM_INFO ;
create streamjob STREAMJOB_STG_TDH_HARDWARE_MEM_INFO as ('
INSERT INTO STG.T_STG_TDH_HARDWARE_MEM_INFO SELECT 
KEY
,HOSTNAME
,HARDYINFO
,Oiginal_Unixtime
,Metrics_DAY
,Metrics_TIME
,Oiginal_Value
,Metrics_Value_Ratio
FROM stream.STREAM_STG_TDH_HARDWARE_MEM_INFO_step1')
jobproperties
('morphling.result.auto.flush'='true');

