# Table: onfleet_hubs

https://docs.onfleet.com/reference/list-hubs

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|organization_id (PK)|String|
|id (PK)|String|
|name|String|
|location|JSON|
|address|JSON|
|teams|StringArray|