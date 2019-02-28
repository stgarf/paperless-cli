# paperless-cli

A CLI tool written in Go to interface with a [Paperless](https://github.com/the-paperless-project/paperless) instance.

## Development

### Prerequisites

You should have a working Go environment and have `$GOPATH/bin` in your `$PATH`.

### Get the code

To download the source code, run:
```shell
$ go get -du github.com/stgarf/paperless-cli
```

The source code will be located in `$GOPATH/src/github.com/stgarf/paperless-cli`.

## Installation and usage

### Precompiled binary

You can get a precompiled binary from the releases page.

### Self-compiled binary

To download source, compile, and install, run:
```shell
$ go get -u github.com/stgarf/paperless-cli
```

The source code will be located in `$GOPATH/src/github.com/stgarf/paperless-cli`.

`$ which paperless-cli` should return the path to the newly installed binary.

### Usage

#### Setting up a config

You can set up a basic YAML-based configuration to be read by placing it in
`$HOME/.paperless-cli.yaml`. Here's an example configuration:
```yaml
# A basic paperless-cli configuration file.

# The hostname of the Paperless instance.
hostname: localhost
# Connect via HTTP or HTTPS.
use_https: false
# Port the Paperless instance is listening on.
port: 8000
# Path to the API root.
root: /api
```

## Running the tests

`$ go test`

## Built With

* [cobra](https://github.com/spf13/cobra/) - A Commander for modern Go CLI interactions

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on the code of conduct, and the process for submitting pull requests to the project.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/stgarf/paperless-cli/tags). 

## Authors

* **Steve Garf** - *Initial CLI work* - [stgarf](https://github.com/stgarf)

See also the list of [contributors](https://github.com/stgarf/paperless-cli/contributors) who participated in this project.

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

* Hat tip to anyone whose code was used
* The awesome community of people maintaining [Paperless](https://github.com/the-paperless-project/paperless)
