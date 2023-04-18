# Table: onfleet_workers

This table shows data for Onfleet Workers.

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
|account_status|String|
|active_task|String|
|additional_capacities|JSON|
|addresses|JSON|
|analytics|JSON|
|capacity|Float|
|delay_time|Float|
|display_name|String|
|has_recently_used_spoofed_locations|Bool|
|id (PK)|String|
|image_url|String|
|location|JSON|
|metadata|JSON|
|name|String|
|on_duty|Bool|
|organization|String|
|phone|String|
|tasks|StringArray|
|teams|StringArray|
|time_created|Timestamp|
|time_last_modified|Timestamp|
|time_last_seen|Timestamp|
|user_data|JSON|
|timezone|String|
|vehicle|JSON|