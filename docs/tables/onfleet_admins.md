# Table: onfleet_admins

This table shows data for Onfleet Admins.

https://docs.onfleet.com/reference/list-administrators

The composite primary key for this table is (**id**, **organization**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|email|String|
|id (PK)|String|
|is_account_owner|Bool|
|is_active|Bool|
|is_read_only|Bool|
|metadata|JSON|
|name|String|
|organization (PK)|String|
|phone|String|
|teams|StringArray|
|time_created|Timestamp|
|time_last_modified|Timestamp|
|type|String|