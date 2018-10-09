# DuckTracker
Sample application to track duck activities

## Development

### Requirements
- Docker 1.12.3-rc1 or above
- Docker Composer 1.8.1 or above
- GNU Make

### Setup
Once you've cloned the repo run:

```
make
```

Sets up the Docker Compose development environment, setup the vendor directory, and compile the go app.

Then run:

```
make run ARGS="make testv"
```

If you see `PASS` then everything is fine.

### Developing

```
make run-dev ARGS="make testx"
```

You run this command after making a code change. This runs the development toolchain and the tests.

### Testing
```
make imaget
```

Makes the ci test image

```
make run-dev ARGS="make test"
```

Runs the tests

```
make run-dev ARGS="make testx"
```

Runs the development toolchain and then the tests

```
make run-dev ARGS="make cover"
```

Run coverage report and compile into html in the reports directory.

```
make run-dev ARGS="make race"
```

Race condition check.

### Go

```
make run-dev ARGS="make compile"
```

Compile the go app.

```
make run-dev ARGS="make dep"
```

Pull packages into vendor folder from remote repository with revisions from vendor.json file.

In order to pass dep authentication, add a file to the root of this repository called `.netrc` with the following content:

```
machine github.com
    login GITHUB_TOKEN
```

```
make run-dev ARGS="make depinfo"
```

Show information about the status of the dependencies.

```
make run-dev ARGS="make pre"
```

Run pre-processors like `goimports` on the code base

```
make run-dev ARGS="make check"
```

Static code analysis tools like `go lint` on the code base.

### Clean up

```
make clean
```

Remove all artifacts created from the project. Does not remove production images.

### Release

```
BUILD_NUMBER=number make image && make push
```

Run to assign a build number to the image and push to image repo.
Nothing needs to be done for Jenkins as it will automatically assign a build no.
