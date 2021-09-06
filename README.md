# scenarigo-plugin-example

[scenarigo](https://github.com/zoncoen/scenarigo) plugin usage example

## Structure

Basically this is the same of [v7.0.0 usage example](https://github.com/zoncoen/scenarigo/tree/3a6173af80230d70c2a3bf264430382eec1caf75#usage).

Added plugin file `gen.go` to check it works correctly.

```shell
$ go build -buildmode=plugin -o gen.so gen.go
```

## What's problem?

I tried using [scenarigo_v0.7.0_Darwin_x86_64](https://github.com/zoncoen/scenarigo/releases/tag/v0.7.0) but it seems not to work plugin function well on `macOS Catalina v10.15.7`.

The error message is below.

``` shell
$ ./v0.7.0_Darwin_x86_64/scenarigo run
$ --- FAIL: github.yml (0.00s)
$     --- FAIL: github.yml/get_scenarigo_repository (0.00s)
$             failed to open plugin: plugin: not implemented
$ FAIL
$ FAIL	github.yml	0.001s
$ FAIL
```