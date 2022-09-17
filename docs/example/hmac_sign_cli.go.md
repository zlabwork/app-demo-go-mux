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
    timeout    = 10 * time.Second
)

type client struct {
    Debug           bool
    accessKeyId     string
    accessKeySecret string
    endpoint        string
}

func newClient(keyId, keySecret, endpoint string) *client {

    return &client{
        Debug:           false,
        accessKeyId:     keyId,
        accessKeySecret: keySecret,
        endpoint:        endpoint,
    }
}

func (cli *client) buildSignature(method, uri string, bodyHash string, headers map[string]string, date, nonce, keySecret string) string {

    u, _ := url.Parse(uri)

    signedBody := date + "\n" + nonce + "\n"
    signedBody += strings.ToUpper(method) + "\n" + u.Path + "\n"

    // query
    query := strings.Split(u.RawQuery, "&")
    sort.Strings(query)
    signedBody += strings.Join(query, "&") + "\n"

    // heads
    var heads []string
    heads = append(heads, "host:"+headers["Host"])
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

    return signature
}

func (cli *client) call(uri string, method string, reqBody []byte) (statusCode int, responseBody []byte, err error) {

    // Parse configure
    aki := cli.accessKeyId
    aks := cli.accessKeySecret
    host := cli.endpoint

    // Parse host
    u, err := url.Parse(host)
    if err != nil {
        return statusCode, nil, err
    }

    // 1. Nonce & Date
    bs := make([]byte, 8)
    _, err = rand.Read(bs)
    if err != nil {
        return statusCode, nil, err
    }
    nonce := hex.EncodeToString(bs)
    date := time.Now().UTC().Format(dateFormat)

    // 2. Body hash
    s256 := sha256.New()
    s256.Write(reqBody)
    bodyHash := hex.EncodeToString(s256.Sum(nil))

    // 3. Headers
    headers := make(map[string]string)
    headers["Host"] = u.Host
    headers["Content-Type"] = "text/html"      // todo
    headers["X-Lab-Nonce"] = nonce             // option
    headers["X-Lab-Date"] = date               // option
    headers["X-Lab-Content-Sha256"] = bodyHash // option

    // 4. Headers Authorization
    sign := cli.buildSignature(method, host+uri, bodyHash, headers, date, nonce, aks)
    headers["Authorization"] = "ZLAB " + fmt.Sprintf("Credential=%s, Date=%s, Nonce=%s, Signature=%s", aki, date, nonce, sign)

    // 5. Request
    req, err := http.NewRequest(method, host+uri, bytes.NewReader(reqBody))
    if err != nil {
        return statusCode, nil, err
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
    httpCli := http.Client{Timeout: timeout}
    resp, err := httpCli.Do(req)
    if err != nil {
        return statusCode, nil, err
    }
    defer resp.Body.Close()

    // 7. Response
    if resp.StatusCode != 200 {
        return statusCode, nil, fmt.Errorf("error StatusCode")
    }
    body, err := ioutil.ReadAll(resp.Body)

    return resp.StatusCode, body, nil
}

func (cli *client) debug(req *http.Request) {

    fmt.Println("~~~~~~~~ Http Request Debug ~~~~~~~~")
    fmt.Println(req.Method + " " + req.Host + req.URL.EscapedPath() + " " + req.Proto)
    for key, item := range req.Header {
        fmt.Println(key + ": " + strings.Join(item, ";"))
    }
    fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
```
