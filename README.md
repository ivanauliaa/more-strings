# Add to Dependencies

```bash
go get github.com/ivanauliaa/more-strings
```

# Import

```go
import morestrings "github.com/ivanauliaa/more-strings"
```

# Usage

**Functions**

All of these functions are returning string

```go
- Intersection(str1, str2 string)
- RemoveDups(str string)
```

# Example
**[Hackerrank: Gemstones](https://www.hackerrank.com/challenges/gem-stones/problem?h_r=internal-search)**

```go
func gemstones(arr []string) int32 {
  result := removeDups(arr[0])

  for i := 1; i < len(arr); i++ {
    temp := removeDups(arr[i])
    result = intersection(result, temp)
  }

  return int32(len(result))
}
```
