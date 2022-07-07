## Random string or bytes
https://xie.infoq.cn/article/f274571178f1bbe6ff8d974f3
```go
// Random string - method 1
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// need rand.Seed
func main() {
	rand.Seed(time.Now().UnixNano())
	randomString(10)
}
```

```go
// Random string - method 2
var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func randomString(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
```


```go
// Random bytes
// set rand.Seed(time.Now().UnixNano()) if use math/rand
import (
	//"math/rand"
	"crypto/rand"
	"fmt"
)

func randomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
```


## 切片排序 []byte
```go
type SortBytes [][]byte

func (a SortBytes) Len() int           { return len(a) }
func (a SortBytes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBytes) Less(i, j int) bool { return bytes.Compare(a[i], a[j]) == -1 }

func main() {
	data := [][]byte{
		{1, 1, 1, 1},
		{3, 3, 3, 3},
		{2, 2, 2, 2},
	}
	sort.Sort(SortBytes(data))
	fmt.Println(data) // [[1 1 1 1] [2 2 2 2] [3 3 3 3]]
}
```


## 切片去重
```go
func main() {

	data := [][]byte{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
	}
	fmt.Println(uniq(data)) // [[1 1 1 1] [2 2 2 2]]
}

func uniq(data [][]byte) [][]byte {

	l := len(data)
	if l == 0 {
		return nil
	}
	idx := 0
	for i := 1; i < l; i++ {
		if bytes.Compare(data[i], data[i-1]) == 0 {
			continue
		} else {
			idx += 1
			data[idx] = data[i]
		}
	}

	return data[:idx+1]
}
```
