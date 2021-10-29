

DROP TABLE STG.T_STG_TDH_DBASERVICE_TASK_INFO;
CREATE TABLE STG.T_STG_TDH_DBASERVICE_TASK_INFO(
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
STORED BY
  'io.transwarp.hyperdrive.HyperdriveStorageHandler'
WITH SERDEPROPERTIES (
  'hbase.columns.mapping'=':key,f:SERVERKEY,f:TASKID,f:SQLID,f:STAGEID,f:TASKHOST,f:TASKSTATUS,f:TASKMESSAGE,f:metrics_time,f:TASKSUBMISSIONTIME,f:TASKCOMPLETIONTIME',
  'serialization.format'='1')
TBLPROPERTIES ('hbase.table.name'='T_STG_TDH_DBASERVICE_TASK_INFO');


DROP TABLE ODS.T_ODS_TDH_DBASERVICE_TASK_INFO;
CREATE TABLE ODS.T_ODS_TDH_DBASERVICE_TASK_INFO(
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
PARTITIONED BY RANGE (metrics_time) (
  PARTITION before202201 VALUES LESS THAN ('2022-01-01'),
  PARTITION before202207 VALUES LESS THAN ('2022-07-01'),
  PARTITION before202301 VALUES LESS THAN ('2023-01-01'),
  PARTITION before202307 VALUES LESS THAN ('2023-07-01'),
  PARTITION before202401 VALUES LESS THAN ('2024-01-01'),
  PARTITION before202407 VALUES LESS THAN ('2024-07-01'),
  PARTITION before202501 VALUES LESS THAN ('2025-01-01'),
  PARTITION before202507 VALUES LESS THAN ('2025-07-01'),
  PARTITION before202601 VALUES LESS THAN ('2026-01-01'),
  PARTITION before202607 VALUES LESS THAN ('2026-07-01'),
  PARTITION before202701 VALUES LESS THAN ('2027-01-01'),
  PARTITION before202707 VALUES LESS THAN ('2027-07-01'),
  PARTITION before202801 VALUES LESS THAN ('2028-01-01'),
  PARTITION before202807 VALUES LESS THAN ('2028-07-01'),
  PARTITION before202901 VALUES LESS THAN ('2029-01-01'),
  PARTITION beforemax VALUES LESS THAN (MAXVALUE))
CLUSTERED BY (SQLID) INTO 83 BUCKETS
STORED AS ORC
TBLPROPERTIES ('transactional'='true');


