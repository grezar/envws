# envws

## What
This is a CLI tool for set AWS credentials as environment varialbe.

## Install
```
$ go get github.com/grezar/envws
```

## Usage

```
$ eval $(envws set)
```

If you want to specify profile, you can use `--profile` option.

```
$ eval $(envws set --profile <profile name>)
```

I recommend you define shell function wrapped envws in ~/.bashrc like below.
```
function awsenv() {
  eval $(envws $@)
}
```
