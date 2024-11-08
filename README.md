# Example usage

Writing data:

```go
import (
    "fmt"
    "github.com/bbfh-dev/go-bin"
)

func main() {
    var buffer bytes.Buffer
    err := bin.NewWriter(&buffer).
        Write(0, 1, 2, 3).
        WriteString("Hello World!").
        WriteUint16(69).
        WriteUint32(420).
        WriteUint64(69420).
        Error()

    if err != nil {
        fmt.Println(err.Error())
    }
}
```

Reading data:

```go
import (
    "github.com/bbfh-dev/go-bin"
)

func readData() error {
    var a byte
    var str string
    var x uint16
    var y uint32
    var z uint64
    err := bin.NewReader(&buffer).
        ReadByte(&a).
        ReadString(&str).
        ReadUint16(&x).
        ReadUint32(&y).
        ReadUint64(&z).
        Error()
    return err
}
```
