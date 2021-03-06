# Testlink Go API

Based on: https://github.com/kinow/testlink-java-api

This library is the abstraction of Testlink XmlRPC API for access from applications written with Go.

## Installation

Add testlink-go-api dependency using:

```sh
go get -u github.com/ninteen19/testlink-go-api
```

Then, import it using:

```go
import (
    "github.com/ninteen19/testlink-go-api"
    "github.com/ninteen19/testlink-go-api/testcase"
    "github.com/ninteen19/testlink-go-api/testproject"
)
```

## Usage

```go
//Add testlink config
testlink.Conf.Key = devKey
testlink.Conf.Url = "https://testlink.gdn-app.com/lib/api/xmlrpc/v1/xmlrpc.php"

//Create testcase
err := testcase.Create(testcase.CreateRequest)

//Get testproject by projectName
resp, err := testproject.Get(testproject.GetRequest)
```


## Example

https://github.com/ninteen19/confluence-to-test-link.git

## TODO

- tests
- add other testlink API