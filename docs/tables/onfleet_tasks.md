# Table: onfleet_tasks

This table shows data for Onfleet Tasks.

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
|additional_quantities|JSON|
|appearance|JSON|
|barcodes|JSON|
|complete_after|Timestamp|
|complete_before|Timestamp|
|completion_details|JSON|
|container|JSON|
|creator|String|
|delay_time|Float|
|dependencies|StringArray|
|destination|JSON|
|estimated_arrival_time|Timestamp|
|estimated_completion_time|Timestamp|
|eta|Int|
|executor|String|
|feedback|JSON|
|id (PK)|String|
|identity|JSON|
|merchant|String|
|metadata|JSON|
|notes|String|
|organization|String|
|overrides|JSON|
|pickup_task|Bool|
|quantity|Float|
|recipients|JSON|
|scan_only_required_barcodes|Bool|
|service_time|Float|
|short_id|String|
|source_task_id|String|
|state|Int|
|time_created|Timestamp|
|time_last_modified|Timestamp|
|tracking_url|String|
|tracking_viewed|Bool|
|worker|String|