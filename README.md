# go-utils

golang的一些常用方法

[go-utils](https://github.com/lizongying/go-utils)
[document](https://pkg.go.dev/github.com/lizongying/go-utils)

## Usage

* slice 排序

```go
package main

import (
  "fmt"
  "github.com/lizongying/go-utils"
)

func main() {
  a := []string{"1", "3", "2"}
  utils.SliceAsc(a)

  //[1 2 3]
  fmt.Println(a)

  b := []uint64{1, 3, 2}
  utils.SliceDesc(b)

  //[3 2 1]
  fmt.Println(b)
}

```

* number 大小

```go
package main

import (
  "fmt"
  "github.com/lizongying/go-utils"
)

func main() {
  a := []int8{1, 3, 2}

  //3
  fmt.Println(utils.Max(a...))

  b := []uint64{1, 3, 2}

  //1
  fmt.Println(utils.Min(b...))
}

 ```

## TODO



