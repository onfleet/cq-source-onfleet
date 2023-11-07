# Table: onfleet_workers

https://docs.onfleet.com/reference/list-workers

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|account_status|`utf8`|
|active_task|`utf8`|
|additional_capacities|`json`|
|addresses|`json`|
|analytics|`json`|
|capacity|`float64`|
|delay_time|`float64`|
|display_name|`utf8`|
|has_recently_used_spoofed_locations|`bool`|
|id (PK)|`utf8`|
|image_url|`utf8`|
|location|`list<item: float64, nullable>`|
|metadata|`json`|
|name|`utf8`|
|on_duty|`bool`|
|organization|`utf8`|
|phone|`utf8`|
|tasks|`list<item: utf8, nullable>`|
|teams|`list<item: utf8, nullable>`|
|time_created|`timestamp[us, tz=UTC]`|
|time_last_modified|`timestamp[us, tz=UTC]`|
|time_last_seen|`timestamp[us, tz=UTC]`|
|user_data|`json`|
|timezone|`utf8`|
|vehicle|`json`|