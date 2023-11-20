# Table: onfleet_teams

https://docs.onfleet.com/reference/list-teams

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|enable_self_assignment|`bool`|
|hub|`utf8`|
|id (PK)|`utf8`|
|managers|`list<item: utf8, nullable>`|
|name|`utf8`|
|tasks|`list<item: utf8, nullable>`|
|time_created|`timestamp[us, tz=UTC]`|
|time_last_modified|`timestamp[us, tz=UTC]`|
|workers|`list<item: utf8, nullable>`|