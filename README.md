
## Welcome to 'byteconv'
*byteconv is a fast []byte converter package*
### Supported types:
* string
* bool
* float32/64
* int8/16/32/64
* uint8/16/32/64

***

### Conver []byte with only one function `ToByte()`
This example uses type `int` but you can use any other supported types from above.

Use `byteconv.ToByte()` to convert `[]byte` .
```go
package main

import (
    "github.com/ecoshub/byteconv"
    "fmt"
)

func main() {
    // test value as int
    val := 1289
    
    // []byte convertion with
    arr, err := byteconv.ToBytes(val)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(arr)
}
```
This is array representation of `int` 1289
```
Output:[9 5 0 0 0 0 0 0]
```
To convert back to `Int` use `ToInt()` function
```go
package main

import (
    "byteconv"
    "fmt"
)

func main() {
    // I created an array from other examples output
    // you can directly use 'arr' variable if you merge them
    arr := []byte{9, 5, 0, 0, 0, 0, 0, 0}

    val, err := byteconv.ToInt(arr)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(val)
}
```
```
Output:1289
```
