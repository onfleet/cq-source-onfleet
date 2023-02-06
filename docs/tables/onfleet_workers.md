# Table: onfleet_workers

https://docs.onfleet.com/reference/list-workers

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
|name|String|
|display_name|String|
|phone|String|
|active_task|String|
|tasks|StringArray|
|on_duty|Bool|
|is_responding|Bool|
|time_last_seen|Timestamp|
|capacity|Int|
|account_status|String|
|image_url|String|
|location|JSON|
|delay_time|Float|
|timezone|String|
|teams|StringArray|
|user_data|JSON|
|vehicle|JSON|
|metadata|JSON|
|has_recently_used_spoofed_locations|Bool|