CREATE  TABLE `T_STG_TDH_HARDWARE_MEM_INFO`(
  `key` string DEFAULT NULL, 
  `hostname` string DEFAULT NULL, 
  `hardyinfo` string DEFAULT NULL, 
  `oiginal_unixtime` string DEFAULT NULL, 
  `metrics_day` string DEFAULT NULL, 
  `metrics_time` string DEFAULT NULL, 
  `oiginal_value` string DEFAULT NULL, 
  `metrics_value_ratio` string DEFAULT NULL
)
ROW FORMAT SERDE 
  'io.transwarp.hyperdrive.serde.HyperdriveSerDe' 
STORED BY 
  'io.transwarp.hyperdrive.HyperdriveStorageHandler' 
WITH SERDEPROPERTIES ( 
  'hbase.columns.mapping'=':key,f:HOSTNAME,f:HARDYINFO,f:Oiginal_Unixtime,f:Metrics_DAY,f:Metrics_TIME,f:HOSTNAME,f:HOSTNAME', 
  'serialization.format'='1')
LOCATION
  'hdfs://nameservice1/inceptor8/user/hive/warehouse/stg.db/stg/stg.t_stg_tdh_hardware_mem_info@hyperdrive.stargate'
TBLPROPERTIES (
  'transient_lastDdlTime'='1629947225', 
  'hbase.table.name'='T_STG_TDH_HARDWARE_MEM_INFO', 
  'hyperdrive.virtual.column'='_vc', 
  'hyperdrive.virtual.family'='f')
