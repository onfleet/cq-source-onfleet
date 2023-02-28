# Table: onfleet_tasks

https://docs.onfleet.com/reference/list-tasks

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
|time_created|Timestamp|
|time_last_modified|Timestamp|
|organization|String|
|short_id|String|
|tracking_url|String|
|worker|String|
|merchant|String|
|executor|String|
|creator|String|
|dependencies|JSON|
|state|Int|
|complete_after|Timestamp|
|complete_before|Timestamp|
|service_time|Int|
|pickup_task|Bool|
|notes|String|
|completion_details|JSON|
|feedback|JSON|
|metadata|JSON|
|overrides|JSON|
|container|JSON|
|recipients|JSON|
|estimated_completion_time|Timestamp|
|estimated_arrival_time|Timestamp|
|destination|JSON|
|did_auto_assign|Bool|
|tracking_viewed|Bool|
|quantity|Int|