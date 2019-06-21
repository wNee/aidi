# aidi

REST API testing framework inspired by frisby, written in Go
### 前言

最近在写go api测试，之前用过cucumber godog这些，用起来感觉很笨重，很不爽，不符合码农习惯


### 亮点

实现了json equal 和 contains功能

### examples


```go
package main

import (
	"fmt"
	"strings"

	"github.com/wNee/aidi"
)

func main() {
	fmt.Println("Aidi!\n")

	aidi.CreateCase("Test GET Go homepage").
		Get("http://golang.org").
		Send().
		ExpectStatus(200)

	body := strings.NewReader(`{"test_key":"test_value"}`)
	aidi.CreateCase("Test POST").
		Post("http://golang.org").
		SetBody(body).
		Send().
		ExpectStatus(200)

	aidi.CreateCase("Test GET Go homepage").
		Get("http://golang.org").
		Send().
		ExpectStatus(400)
	aidi.Global.PrintReport()
}

```

Sample Output

```
Aidi!

.......
For 3 requests made
  FAILED  [1/3]
      [Test GET Go homepage]
        -  Expected Status 400, but got 200: "200 OK"
        
```

json equal and contrains
```
aidi.CreateCase("Test GET Go homepage").
		Get("url").
		Send().
		ExpectStatus(200).
		ExpectBodyJson(`{
                         "id": 1,
                         "name": "testName",
                         "email": "test@gmail.com"
                       }`)
如果email无法确实是什么可以用来确定json包含关系
aidi.CreateCase("Test GET Go homepage").
		Get("url").
		Send().
		ExpectStatus(200).
		ExpectBodyContainJson()(`{
                        "id": 1,
                        "name": "testName",
                      
                      }`)