package middleware

import (
	"app"
	"app/msg"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

const requestTimeout = 300

const dateFormat = "20060102T150405Z0700"

type authorization struct {
	Credential string
	Signature  string
}

func SignatureMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1. Check date time
		if checkDatetime(r) != nil {
			app.ResponseData(w, msg.ErrTimeout, struct {
				ServerDate string
			}{ServerDate: time.Now().UTC().Format(dateFormat)})
			return
		}

		// 2. Authorization
		token, err := parseAuthorization(r)
		if err != nil {
			app.ResponseData(w, msg.ErrAccess, struct {
				Error string
			}{Error: err.Error()})
			return
		}
		aks := os.Getenv("AK_" + token.Credential)

		// 3. SignedBody
		signedBody, err := createSignedBody(r)
		if err != nil {
			app.ResponseData(w, msg.ErrSignature, struct {
				Error string
			}{err.Error()})
			return
		}

		// 4. Calculating sign hmac-sha256
		h := hmac.New(sha256.New, []byte(aks))
		h.Write([]byte(signedBody))
		signature := hex.EncodeToString(h.Sum(nil))

		// 5. Check
		if signature != token.Signature {
			app.ResponseData(w, msg.ErrSignature, struct {
				Error      string
				SignedBody string
				Refer      string
			}{SignedBody: signedBody, Error: "signature is not matched"})
			return
		}

		// 5. Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func parseAuthorization(r *http.Request) (*authorization, error) {

	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("missing Authorization")
	}
	data := make(map[string]string)
	for _, tk := range strings.Split(token, ",") {
		if !strings.Contains(tk, "=") {
			continue
		}
		s := strings.Split(tk, "=")
		data[s[0]] = s[1]
	}
	sign, ok := data["Signature"]
	if ok != true {
		return nil, fmt.Errorf("missing Authorization Signature")
	}
	aki, ok := data["Credential"]
	if ok != true {
		return nil, fmt.Errorf("missing Authorization Credential")
	}

	return &authorization{
		Credential: aki,
		Signature:  sign,
	}, nil
}

func createSignedBody(r *http.Request) (string, error) {

	// 1. HTTPMethod CanonicalURI
	signedBody := r.Method + "\n" + r.URL.Path + "\n"

	// 2. CanonicalQueryString
	query := strings.Split(r.URL.RawQuery, "&")
	sort.Strings(query)
	signedBody += strings.Join(query, "&") + "\n"

	// 3. CanonicalHeaders
	if r.Header.Get("Content-Type") == "" {
		return "", fmt.Errorf("missing Content-Type")
	}
	var headers []string
	headers = append(headers, "host:"+r.Host)
	headers = append(headers, "content-type:"+r.Header.Get("Content-Type"))
	for k, v := range r.Header {
		if !strings.HasPrefix(k, "X-Lab-") {
			continue
		}
		headers = append(headers, strings.ToLower(k)+":"+strings.Join(v, ","))
	}
	sort.Strings(headers)
	signedHeaders := strings.Join(headers, "\n")
	signedBody += signedHeaders + "\n"

	// 4. HashedPayload
	var reader io.Reader = r.Body
	b, _ := ioutil.ReadAll(reader)
	s256 := sha256.New()
	s256.Write(b)
	contentHash := hex.EncodeToString(s256.Sum(nil))
	signedBody += contentHash

	// Check X-Lab-Content-Sha256
	hash := r.Header.Get("X-Lab-Content-Sha256")
	if len(hash) > 0 && (hash != contentHash) {
		return "", fmt.Errorf("invalid HashedPayload")
	}

	return signedBody, nil
}

func checkDatetime(r *http.Request) error {

	dateTime, err := time.Parse(dateFormat, r.Header.Get("X-Lab-Date"))
	if err != nil {
		return err
	}
	if math.Abs(float64(time.Now().Unix()-dateTime.Unix())) > requestTimeout {
		return fmt.Errorf("timeout")
	}
	return nil
}
