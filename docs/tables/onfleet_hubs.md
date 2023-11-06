# Table: onfleet_hubs

This table shows data for Onfleet Hubs.

https://docs.onfleet.com/reference/list-hubs

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|address|`json`|
|id (PK)|`utf8`|
|location|`list<item: float64, nullable>`|
|name|`utf8`|
|teams|`list<item: utf8, nullable>`|