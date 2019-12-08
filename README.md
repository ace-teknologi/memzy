# Memzy

[![Travis build status](https://api.travis-ci.org/ace-teknologi/memzy.png)](https://travis-ci.org/ace-teknologi/memzy)
[![Maintainability](https://api.codeclimate.com/v1/badges/28282cdb245406093d59/maintainability)](https://codeclimate.com/github/ace-teknologi/memzy/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/28282cdb245406093d59/test_coverage)](https://codeclimate.com/github/ace-teknologi/memzy/test_coverage)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Face-teknologi%2Fmemzy.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Face-teknologi%2Fmemzy?ref=badge_shield)

A simple object persistance inferface for golang.

## Usage

Prequisite: add json tags to any objects you want to store.
```golang
type Bob {
    Name string `json:"name"` // in our examples this will be the primary key
    Height int  `json:"height"`
}
```

### DynamoDB client

```golang
import "github.com/ace-teknologi/memzy/dynamodb"

...

c := dynamodb.New("BOB_STORAGE")
```

### Memory client

The memory client is good for testing.

```golang
import "github.com/ace-teknologi/memzy/memory"

...

c := memory.New("name")
```

### Use the interface

Generally I don't use the above clients directly. Instead I use the interface
which enables me to switch implementations in testing.

```golang
import (
    "os"

    "github.com/ace-teknologi/memzy"
    "github.com/ace-teknologi/memzy/dynamodb"
    "github.com/ace-teknologi/memzy/memory"
)

var memzyClient memzy.memzy

func init() {
    if os.Getenv == "PRODUCTION" {
        memzyClient = dynamodb.New("BOB_STORAGE")
    } else {
        memzyClient = memory.New("name")
    }
}

```

### GetItem

```golang
var rdj Bob
memzyClient.GetItem(rdj, map[string]interface{}{"Name": "Robert Downey Jr."})
fmt.Printf("Robert Downey Jr is %d cm tall", rdj.Height) // Robery Downey Jr. is 173 cm tall
```

### PutItem

```golang
var rnm = &Bob{
    Name: "Robert Nesta Marley, OM",
    Height: 170,
}
memzyClient.PutItem(rnm)
```

### NewIter

This has basically no features, just the ability to iterate over everything you have stored.

```golang
iter := memzyClient.NewIter()
// do stuff
```

__Warning__: the API of this pre-1.0 library is unstable. It is recommended to use
dependency management.
