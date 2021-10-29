
DROP STREAM STREAM.STREAM_STG_TDH_DBASERVICE_TASK_INFO_step0;
CREATE STREAM STREAM.STREAM_STG_TDH_DBASERVICE_TASK_INFO_step0(
	TYPE                STRING,
    KEY                 STRING,
    SERVERKEY           STRING,
    TASKID              STRING,
    SQLID               STRING,
    STAGEID             STRING,
    TASKHOST            STRING,
    TASKSTATUS          STRING,
    TASKMESSAGE         STRING,
    metrics_time 		STRING,
    TASKSUBMISSIONTIME  STRING,
    TASKCOMPLETIONTIME  STRING)
    ROW FORMAT SERDE 'org.apache.hadoop.hive.contrib.serde2.MultiDelimitSerDe'
WITH SERDEPROPERTIES ('input.delimited'='|+|')
TBLPROPERTIES(
"topic"="tdh-dbaservice",
"kafka.broker.list"="192.168.210.128:9192",
"kafka.zookeeper"="192.168.210.128:2281/kafak1",
--"kafka.zookeeper"="10.16.32.54:2181/kafak1",
--"kafka.broker.list"="none-datacenter.kafka.chinner.com:9192",
"transwarp.consumer.group.id"="TASK_INFO1");




drop  STREAM stream.STREAM_STG_TDH_DBASERVICE_TASK_INFO_step1;
CREATE STREAM stream.STREAM_STG_TDH_DBASERVICE_TASK_INFO_step1 AS
SELECT
	KEY,
	SERVERKEY,
	TASKID,
	SQLID,
	STAGEID,
	TASKHOST,
	TASKSTATUS,
	TASKMESSAGE,
    metrics_time,
	TASKSUBMISSIONTIME,
	TASKCOMPLETIONTIME
FROM stream.STREAM_STG_TDH_DBASERVICE_TASK_INFO_step0
WHERE type='task';



DROP  streamjob STREAMJOB_STG_TDH_DBASERVICE_TASK_INFO ;
create streamjob STREAMJOB_STG_TDH_DBASERVICE_TASK_INFO as ('
INSERT INTO STG.T_STG_TDH_DBASERVICE_TASK_INFO SELECT *
FROM stream.STREAM_STG_TDH_DBASERVICE_TASK_INFO_step1')
jobproperties
('morphling.result.auto.flush'='true');




START STREAMJOB STREAMJOB_STG_TDH_DBASERVICE_TASK_INFO ;
