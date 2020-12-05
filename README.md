# Gojira

[![Lint](https://github.com/leozhantw/gojira/workflows/Lint/badge.svg)](https://github.com/leozhantw/gojira/actions?query=workflow%3ALint)
[![Test](https://github.com/leozhantw/gojira/workflows/Test/badge.svg)](https://github.com/leozhantw/gojira/actions?query=workflow%3ATest)
[![Release](https://img.shields.io/github/v/release/leozhantw/gojira?label=Release)](https://github.com/leozhantw/gojira/releases)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/leozhantw/gojira/blob/master/LICENSE)

`gojira` can save your time on [Jira](https://www.atlassian.com/software/jira) operation through the command line.

## Installation

### Homebrew
```shell script
brew install leozhantw/tap/gojira
```

## Configuration
Before using the CLI, you need to configure your Jira information through the `gojira configure` command.
```shell script
$ gojira configure
Your website URL: https://gojira.atlassian.net
Your token: your-jira-token
```
