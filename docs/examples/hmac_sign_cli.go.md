```go
package remote

import (
    "bytes"
    "crypto/hmac"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "sort"
    "strings"
    "time"
)

const (
    dateFormat = "20060102T150405Z0700"
    reqTimeout = 10 * time.Second
)

type client struct {
    Debug           bool
    accessKeyId     string
    accessKeySecret string
    endpoint        string
}

// NewClient e.g. NewClient("ACCESS-KEY-ID", "ACCESS-KEY-SECRET", "http://localhost:8000")
func NewClient(accessKeyId, accessKeySecret, endpoint string) *client {

    return &client{
        Debug:           true,
        accessKeyId:     accessKeyId,
        accessKeySecret: accessKeySecret,
        endpoint:        endpoint,
    }
}

func (cli *client) Call(uri string, method string, reqBody []byte) (responseBody []byte, err error) {

    // Parse configure
    aki := cli.accessKeyId
    aks := cli.accessKeySecret
    host := cli.endpoint

    // 1. Nonce & Date
    bs := make([]byte, 8)
    _, err = rand.Read(bs)
    if err != nil {
        return nil, err
    }
    nonce := hex.EncodeToString(bs)
    date := time.Now().UTC().Format(dateFormat)

    // 2. Body hash
    s256 := sha256.New()
    s256.Write(reqBody)
    bodyHash := hex.EncodeToString(s256.Sum(nil))

    // 3. Headers
    headers := make(map[string]string)
    // headers["Host"] = host
    headers["Content-Type"] = "application/json" // text/plain, text/html, application/x-www-form-urlencoded
    headers["X-Lab-Nonce"] = nonce               // option
    headers["X-Lab-Date"] = date                 // option
    headers["X-Lab-Content-Sha256"] = bodyHash   // option

    // 4. Headers Authorization
    sign, err := cli.buildSignature(method, host+uri, bodyHash, headers, date, nonce, aks)
    if err != nil {
        return nil, err
    }
    headers["Authorization"] = "ZLAB " + fmt.Sprintf("Credential=%s, Date=%s, Nonce=%s, Signature=%s", aki, date, nonce, sign)

    // 5. Request
    req, err := http.NewRequest(method, host+uri, bytes.NewReader(reqBody))
    if err != nil {
        return nil, err
    }

    // set headers
    for k, v := range headers {
        req.Header.Set(k, v)
    }

    // debug
    if cli.Debug {
        cli.debug(req)
    }

    // 6. Do Request
    httpCli := http.Client{Timeout: reqTimeout}
    resp, err := httpCli.Do(req)
    if err != nil {
        return nil, fmt.Errorf("[RPC] " + err.Error())
    }
    defer resp.Body.Close()

    // 7. Response
    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("[RPC] error StatusCode %d", resp.StatusCode)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}

func (cli *client) buildSignature(method, uri string, bodyHash string, headers map[string]string, date, nonce, keySecret string) (string, error) {

    // 1. parse uri
    u, err := url.Parse(uri)
    if err != nil {
        return "", err
    }

    // 2. check request Host & Content-Type
    hostAddress := u.Host
    if len(hostAddress) == 0 {
        _, ok := headers["Host"]
        if !ok {
            return "", fmt.Errorf("invalid Host")
        }
        hostAddress = headers["Host"]
    }
    if _, ok := headers["Content-Type"]; !ok {
        return "", fmt.Errorf("the Content-Type header is missing")
    }

    // 3.
    signedBody := date + "\n" + nonce + "\n"
    signedBody += strings.ToUpper(method) + "\n" + u.Path + "\n"

    // query
    query := strings.Split(u.RawQuery, "&")
    sort.Strings(query)
    signedBody += strings.Join(query, "&") + "\n"

    // heads
    var heads []string
    heads = append(heads, "host:"+hostAddress)
    heads = append(heads, "content-type:"+headers["Content-Type"])
    for k, v := range headers {
        if !strings.HasPrefix(k, "X-Lab-") {
            continue
        }
        heads = append(heads, strings.ToLower(k)+":"+v)
    }
    sort.Strings(heads)
    signedBody += strings.Join(heads, "\n") + "\n"

    // HashedPayload
    signedBody += bodyHash

    // Hmac-sha256
    h := hmac.New(sha256.New, []byte(keySecret))
    h.Write([]byte(signedBody))
    signature := hex.EncodeToString(h.Sum(nil))

    return signature, nil
}

func (cli *client) debug(req *http.Request) {

    fmt.Println("~~~~~~~~ Http Request Debug ~~~~~~~~")
    fmt.Println(req.Method + " " + req.URL.String() + " " + req.Proto)
    for key, item := range req.Header {
        fmt.Println(key + ": " + strings.Join(item, ";"))
    }
    fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
```
