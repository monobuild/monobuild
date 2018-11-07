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

This is a vgo enabled repository

## Running the tests

You can run the tests by calling `go test ./...`

## Deployment

You can download the binary from the releases page or use the deb package to install it on a Debian system.

## Usage

### How does monobuild work

monobuild scans the current directory and all subdirectories for a marker file (default name: .MONOBUILD). Each marker can have a dependency on another marker. Based on the dependencies so called stages are calculated and executed.

### What is in a marker file

A marker file contains a build configuration with the following fields:

|Field|Required|Description|
|---|---|---|
|Commands|false|Commands are shell commands to be executed|
|Environment|false|A list of environment variables passed to the commands|
|Label |true|Label is the name of the build configuration|
|Dependencies|false|A list of dependencies to other build configurations|
|Parallel|false|Build configuration may be run in parallel with other build configurations|     

A valid sample:

    ---

    commands:
      - echo new marker
    environment:
      MONOBUILD_VERSION: develop
    label: main

### Customization

On the root level there are the following commandline parameters possible:

|Parameter|Description|
|---|---|
|--config|Set an alternative config file|
|--log-level|Set log level to debug, info or warn (fallback) (default "warn")|
|--marker|name of marker file (default ".MONOBUILD")|
|--no-parallelism|disable parallel execution of steps|
|--quit|Do not show header (version info and name)|

### Sample layout

`.MONOBUILD` in root directory

    ---
    
    commands:
      - echo root dir
    dependencies:
      - other component
    label: main

`.MONOBUILD` in sub directory:

    ---
    
    commands:
      - echo other component
    label: other component
    parallel: true

Above sample creates a run with two stages to run first `other component` and than `main`.

Let's add another `.MONOBUILD` in another directory:

    ---
    
    commands:
      - echo yet other component
    label: yet other component
    parallel: true

This will be added to the first stage (no dependency) and executed in parallel as both configurations allow for parallel execution.

### Parallelism

Unless `--no-parallelism` is passed first all configurations that are not allowed to run in parallel in a stage are executed and afterward the one that are allowed are executed. As such you _could_ introduce sub stages but this is highly discouraged. It is better to clearly communicate dependencies and adding additional stages. 

### Usage as an API

monobuild was designed to be usable as a library.

To use monobuild create an instance by calling `NewMonobuild` with the base path to work on.

    cfg := monobuild.NewMonoBuild(dir)

    // cfg.DisableParallelism = true to disable any parallel tasks
    // cfg.MarkerName = "somestring" to use somestring as markerfile

A four step process follows:

1. Load the build configurations from disk using `cfg.LoadConfigurations()`
2. Add configurations from code using `cfg.AddBuildConfiguration()`
3. Run `cfg.Setup()` to build the plan
4. Execute the build using `cfg.Run()`

See the monobuild utility for an example how to consume this library.

#### AddBuildConfiguration

## History

|Version|Description|
|---|---|
|1.0.0|Allow adding buildconfigurations from code|
||Use go modules|
||Allow limiting to one build configuration including dependencies|
|0.9.0|initial working copy|

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
