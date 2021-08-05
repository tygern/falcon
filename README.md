# Falcon

![falcon](readme-images/falcon.png)

A weather CLI

## Test

Run tests in all packages with

```shell
go test ./...
```

## Build

Use the Go CLI's `build` and `install` commands to build or install the app.
Set the `OWM_KEY` environment variable to your OpenWeather API key and use the `run` command to run the app.

## Usage

Pass a zipcode as an argument to the CLI to find the forecast for that zipcode.

```shell
falcon -zip 67530
```
