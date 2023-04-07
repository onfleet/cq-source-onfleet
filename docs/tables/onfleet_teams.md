# Table: onfleet_teams

This table shows data for Onfleet Teams.

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
|enable_self_assignment|Bool|
|hub|String|
|id (PK)|String|
|managers|StringArray|
|name|String|
|tasks|StringArray|
|time_created|Timestamp|
|time_last_modified|Timestamp|
|workers|StringArray|