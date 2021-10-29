
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
"topic"="**-******",
"kafka.broker.list"="***.***.***.***:****",
"kafka.zookeeper"="***.***.***.***:****/****",
--"kafka.zookeeper"="***.***.***.***:****/****",
--"kafka.broker.list"=""="***.***.***.***:****",",
"transwarp.consumer.group.id"="******");



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
