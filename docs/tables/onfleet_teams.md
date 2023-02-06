# Table: onfleet_teams

https://docs.onfleet.com/reference/list-teams

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
|enable_self_assignment|Bool|
|name|String|
|hub|String|
|workers|StringArray|
|managers|StringArray|
|tasks|StringArray|