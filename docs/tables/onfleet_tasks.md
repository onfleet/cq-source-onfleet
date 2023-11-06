# Table: onfleet_tasks

This table shows data for Onfleet Tasks.

https://docs.onfleet.com/reference/list-tasks

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|additional_quantities|`json`|
|appearance|`json`|
|barcodes|`json`|
|complete_after|`time64[us]`|
|complete_before|`time64[us]`|
|completion_details|`json`|
|container|`json`|
|creator|`utf8`|
|delay_time|`float64`|
|dependencies|`list<item: utf8, nullable>`|
|destination|`json`|
|estimated_arrival_time|`time64[us]`|
|estimated_completion_time|`time64[us]`|
|eta|`int64`|
|executor|`utf8`|
|feedback|`json`|
|id (PK)|`utf8`|
|identity|`json`|
|merchant|`utf8`|
|metadata|`json`|
|notes|`utf8`|
|organization|`utf8`|
|overrides|`json`|
|pickup_task|`bool`|
|quantity|`float64`|
|recipients|`json`|
|scan_only_required_barcodes|`bool`|
|service_time|`float64`|
|short_id|`utf8`|
|source_task_id|`utf8`|
|state|`int64`|
|time_created|`time64[us]`|
|time_last_modified|`time64[us]`|
|tracking_url|`utf8`|
|tracking_viewed|`bool`|
|worker|`utf8`|