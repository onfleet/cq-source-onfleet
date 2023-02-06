# Table: onfleet_admins

https://docs.onfleet.com/reference/list-administrators

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|time_created|Int|
|time_last_modified|Int|
|organization|String|
|email|String|
|type|String|
|name|String|
|is_active|Bool|
|is_read_only|Bool|
|is_account_owner|Bool|
|phone|String|
|teams|StringArray|
|metadata|JSON|