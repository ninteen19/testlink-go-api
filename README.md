# Testlink Go API

Based on: https://github.com/kinow/testlink-java-api

This library is the abstraction of Testlink XmlRPC API for access from applications written with Go.

## Installation

Install testlink-go-api with:

```sh
go get -u github.com/ninteen19/testlink-go-api
```

Then, import it using:

```go
import (
    "github.com/ninteen19/testlink-go-api"
    "github.com/ninteen19/testlink-go-api/testcase"
)
```

## Usage

```go
//Add testlink config
testlink.Conf.Key = devKey
testlink.Conf.Url = "https://testlink.gdn-app.com/lib/api/xmlrpc/v1/xmlrpc.php"

//Create testcase
err := testcase.Create(testcase.CreateRequest)
```

## TODO

- tests
- add another API