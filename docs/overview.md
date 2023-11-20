# CloudQuery Onfleet Source Plugin

An Onfleet source plugin for CloudQuery that loads data from Onfleet to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more. Based on the official [gonfleet library](https://github.com/onfleet/gonfleet/blob/main/LICENSE).

## Links

- [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
- [Supported Tables](tables/README.md)

## Authentication

To run `cloudquery sync` with the source plugin, you will need to use an [Onfleet API KEY](https://docs.onfleet.com/reference/authentication).
The API key needs to be provided in the `api_key` configuration parameter (see below).

## Configuration

The following source configuration file will sync to a PostgreSQL database. It uses an API key provided in the `ONFLEET_API_KEY` environment variable.
See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml copy
kind: source
spec:
  name: "onfleet"
  path: "onfleet/onfleet"
  registry: "cloudquery"
  tables: ["*"]
  version: "v2.0.0"
  destinations:
    - "postgresql"
  spec:
    api_key: "${ONFLEET_API_KEY}"
    # optional: timestamp to sync tasks from (by default, will only sync last 3 months of tasks)
    # list_tasks_from: "2023-04-01T01:00:00Z"
    # optional: concurrency setting
    # concurrency: 1000
```
### Plugin Spec

- `api_key` (string, required): The API key
- `concurrency` (int, optional, default: `1000`): Best effort maximum number of Go routines to use. Lower this number to reduce memory usage.
- `list_tasks_from` (string timestamp formatted as RFC3339, optional): Timestamp to sync tasks from