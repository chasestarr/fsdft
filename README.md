## FSDFT (File system depth-first traversal)

[Link to GoDocs](https://godoc.org/github.com/chasestarr/fsdft)

```go
import (
  "fmt"

  "github.com/chasestarr/fsdft"
)

func print(root string, file os.FileInfo) {
  fmt.Println(file.Name(), root)
}

func main() {
  fsdft.DFT(".", print)
}
```
