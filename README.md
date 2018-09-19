# monobuild

... is a tool to build software within a mono repository. Though you might use monobuild to run the build it self it is best combined with another build tool like [task](https://github.com/go-task/task) or make.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

monobuild is build with Go. So you need to have a Go environment up and running. Support for Go with modules is planned but not in place. See Go's [Getting Started](https://golang.org/doc/install) to setup your Go environment.

### Installing

To get the code of monobuild you can run go get:

    go get -u github.com/monobuild/monobuild

Within `$GOPATH/src/github.com/monobuild/monobuild` you should be able to run a test:

    go run cmd/monobuild/main.go

## Running the tests

TODO

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
* [Viper](https://github.com/spf13/viper) - Go configuration with fangs

## Contributing

Please read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/monobuild/monobuild/tags).

## Authors

* **Sascha Andres** - *Initial work* - [sascha-andres](https://github.com/sascha-andres)

See also the list of [contributors](https://github.com/monobuild/monobuild/contributors) who participated in this project.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [Contributor Covenant](https://www.contributor-covenant.org/) as the source for the code of conduct
* [PurpleBooth](https://github.com/PurpleBooth) for the [README blueprint](https://gist.githubusercontent.com/PurpleBooth/109311bb0361f32d87a2/raw/8254b53ab8dcb18afc64287aaddd9e5b6059f880/README-Template.md)