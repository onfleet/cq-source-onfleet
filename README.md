# CloudQuery Onfleet Source Plugin

[![test](https://github.com/onfleet/cq-source-onfleet/actions/workflows/test.yaml/badge.svg)](https://github.com/onfleet/cq-source-onfleet/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/onfleet/cq-source-onfleet)](https://goreportcard.com/report/github.com/onfleet/cq-source-onfleet)

An Onfleet source plugin for CloudQuery that loads data from Onfleet to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more. Based on the official [gonfleet library](https://github.com/onfleet/gonfleet/blob/main/LICENSE).

## Using the plugin

See [overview.md](docs/overview.md)

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

### Compile
```bash
make build
```

### Release
```bash
VERSION=1.0.0 make dist
```