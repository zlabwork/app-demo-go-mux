```go
package remote

import (
	"bytes"
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

type reqConfig struct {
	accessKeyId     string
	accessKeySecret string
	endpoint        string
}

type client struct {
	Debug bool
}

func newClient() *client {

	return &client{Debug: false}
}

func (cli *client) buildSignature(method, uri string, body []byte, headers map[string]string, aks string) string {

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
	s256.Write(body)
	hash := hex.EncodeToString(s256.Sum(nil))
	signedBody += hash

	// Hmac-sha256
	h := hmac.New(sha256.New, []byte(aks))
	h.Write([]byte(signedBody))
	signature := hex.EncodeToString(h.Sum(nil))

	return signature
}

func (cli *client) callRemote(cfg *reqConfig, method string, uri string, reqBody []byte) ([]byte, error) {

	// Parse configure
	aki := cfg.accessKeyId
	aks := cfg.accessKeySecret
	host := cfg.endpoint

	// Parse host
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	// 1. Headers
	bs := make([]byte, 8)
	rand.Read(bs)
	s256 := sha256.New()
	s256.Write(reqBody)
	headers := make(map[string]string)
	headers["Host"] = u.Host
	headers["Content-Type"] = "text/html"
	headers["X-Lab-Nonce"] = base64.StdEncoding.EncodeToString(bs)
	headers["X-Lab-Date"] = time.Now().UTC().Format(dateFormat)
	headers["X-Lab-Content-Sha256"] = hex.EncodeToString(s256.Sum(nil))
	// Authorization
	sign := cli.buildSignature(method, host+uri, reqBody, headers, aks)
	headers["Authorization"] = "Credential=" + aki + ",Signature=" + sign

	// 2. Request
	req, err := http.NewRequest(method, host+uri, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// debug
	if cli.Debug {
		cli.debug(req)
	}

	// 3. Do Request
	client := http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 4. Response
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error StatusCode")
	}
	body, err := ioutil.ReadAll(resp.Body)

	return body, nil
}

func (cli *client) debug(req *http.Request) {
	fmt.Println("======== Http Request Debug ========")
	fmt.Println(req.Method + " " + req.Host + req.URL.EscapedPath() + " " + req.Proto)
	for key, item := range req.Header {
		fmt.Println(key + ": " + strings.Join(item, ";"))
	}
	fmt.Println("====================================")
}

// TODO: custom request libraries
type requestLibs struct {
	cli *client
	cfg map[string]*reqConfig
}

func NewRequestLibs() *requestLibs {
	return &requestLibs{
		cli: newClient(),
		cfg: make(map[string]*reqConfig),
	}
}

func (lib *requestLibs) getConfig(tag string) *reqConfig {

	v, ok := lib.cfg[tag]
	if ok {
		return v
	}
	lib.cfg[tag] = &reqConfig{
		accessKeyId:     os.Getenv("RPC_" + tag + "_ID"),
		accessKeySecret: os.Getenv("RPC_" + tag + "_KEY"),
		endpoint:        os.Getenv("RPC_" + tag + "_URL"),
	}
	return lib.cfg[tag]
}

func (lib *requestLibs) CallUser(method string, uri string, reqBody []byte) (string, error) {

	cfg := lib.getConfig("USR")
	bs, err := lib.cli.callRemote(cfg, method, uri, reqBody)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
```
