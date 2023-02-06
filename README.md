# CloudQuery Onfleet Source Plugin

[![test](https://github.com/cloudquery/cq-source-onfleet/actions/workflows/test.yaml/badge.svg)](https://github.com/cloudquery/cq-source-onfleet/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudquery/cq-source-onfleet)](https://goreportcard.com/report/github.com/cloudquery/cq-source-onfleet)

An Onfleet source plugin for CloudQuery that loads data from Onfleet to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more. Based on a fork of (MIT LIcense https://github.com/keplr-team/go-onfleet). 

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)


## Authentication

To run `cloudquery sync` with the source plugin, you will need to use an [Onfleet API KEY](https://docs.onfleet.com/reference/authentication).
The API key needs to be provided in the `api_key` configuration parameter (see below).

## Configuration

The following source configuration file will sync to a PostgreSQL database. It uses an API key provided in the `ONFLEET_API_KEY` environment variable.
See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "onfleet"
  path: "onfleet/onfleet"
  version: "v1.0.0"
  destinations:
    - "postgresql"
  spec:
    api_key: "${ONFLEET_API_KEY}"
```

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```