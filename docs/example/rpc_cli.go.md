```go
package rpc

import (
    "crypto/hmac"
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "sort"
    "strings"
    "time"
)

const (
    dateFormat = "20060102T150405Z0700"
    timeout    = 10 * time.Second
)

type accessConfig struct {
    AccessKeyId     string
    AccessKeySecret string
    Endpoint        string
}

func CallUser(method string, uri string, reqBody string) (string, error) {

    cfg := getConfig("USR")
    return callRemote(cfg, method, uri, reqBody)
}

func callRemote(cfg *accessConfig, method string, uri string, reqBody string) (string, error) {

    // Parse configure
    aki := cfg.AccessKeyId
    aks := cfg.AccessKeySecret
    host := cfg.Endpoint

    // Parse host
    u, err := url.Parse(host)
    if err != nil {
        return "", err
    }

    // 1. Headers
    bs := make([]byte, 8)
    rand.Read(bs)
    s256 := sha256.New()
    s256.Write([]byte(reqBody))
    headers := make(map[string]string)
    headers["Host"] = u.Host
    headers["Content-Type"] = "text/html"
    headers["X-Lab-Nonce"] = base64.StdEncoding.EncodeToString(bs)
    headers["X-Lab-Date"] = time.Now().Format(dateFormat)
    headers["X-Lab-Content-Sha256"] = hex.EncodeToString(s256.Sum(nil))
    // Authorization
    sign := createSignature(method, host+uri, reqBody, headers, aks)
    headers["Authorization"] = "Credential=" + aki + ",Signature=" + sign

    // 2. Request
    req, err := http.NewRequest(method, host+uri, strings.NewReader(reqBody))
    if err != nil {
        return "", err
    }

    // Set headers
    for k, v := range headers {
        req.Header.Set(k, v)
    }

    // 3. Do Request
    cli := http.Client{Timeout: timeout}
    resp, err := cli.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // 4. Response
    if resp.StatusCode != 200 {
        return "", fmt.Errorf("error StatusCode")
    }
    body, err := ioutil.ReadAll(resp.Body)

    return string(body), nil
}

func createSignature(method, uri, body string, headers map[string]string, aks string) string {

    u, _ := url.Parse(uri)
    signedBody := strings.ToUpper(method) + "\n" + u.Path + "\n"

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
    s256 := sha256.New()
    s256.Write([]byte(body))
    hash := hex.EncodeToString(s256.Sum(nil))
    signedBody += hash

    // Hmac-sha256
    h := hmac.New(sha256.New, []byte(aks))
    h.Write([]byte(signedBody))
    signature := hex.EncodeToString(h.Sum(nil))

    return signature
}

func getConfig(tag string) *accessConfig {

    aki := os.Getenv("RPC_" + tag + "_ID")
    aks := os.Getenv("RPC_" + tag + "_KEY")
    host := os.Getenv("RPC_" + tag + "_URL")
    return &accessConfig{
        AccessKeyId:     aki,
        AccessKeySecret: aks,
        Endpoint:        host,
    }
}

```
