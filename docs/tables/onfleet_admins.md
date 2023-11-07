# Table: onfleet_admins

https://docs.onfleet.com/reference/list-administrators

The composite primary key for this table is (**id**, **organization**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|email|`utf8`|
|id (PK)|`utf8`|
|is_account_owner|`bool`|
|is_active|`bool`|
|is_read_only|`bool`|
|metadata|`json`|
|name|`utf8`|
|organization (PK)|`utf8`|
|phone|`utf8`|
|teams|`list<item: utf8, nullable>`|
|time_created|`timestamp[us, tz=UTC]`|
|time_last_modified|`timestamp[us, tz=UTC]`|
|type|`utf8`|