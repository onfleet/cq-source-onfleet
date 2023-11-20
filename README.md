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
### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).

### Publish a new version to the Cloudquery Hub

After tagging a release, you can build and publish a new version to the [Cloudquery Hub](https://hub.cloudquery.io/) by running the following commands.
Replace `v1.0.0` with the new version number.

```bash
# "make dist" uses the README as main documentation and adds a generic release note. Output is created in dist/
VERSION=v1.0.0 make dist

# Login to cloudquery hub and publish the new version
cloudquery login
cloudquery plugin publish --finalize
```

After publishing the new version, it will show up in the [hub](https://hub.cloudquery.io/).

For more information please refer to the official [Publishing a Plugin to the Hub](https://www.cloudquery.io/docs/developers/publishing-a-plugin-to-the-hub) guide.