package middleware

import (
	"app/msg"
	"app/response"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	requestTimeout      = 300
	dateFormat          = "20060102T150405Z0700"
	authorizationPrefix = "ZLAB"
)

func SignatureMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ignore
		prefix := []string{"/keys", "/token", "/assets"}
		for _, p := range prefix {
			if strings.HasPrefix(r.RequestURI, p) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// 1. Authorization
		auth, err := parseAuthorization(r)
		if err != nil {
			response.Message(r.Context(), w, msg.ErrAccess, err.Error())
			return
		}

		// 2. Check Authorization format & Check date time
		date, ok := auth["Date"]
		if !ok || checkDatetime(date) != nil {
			response.Data(r.Context(), w, msg.ErrTimeout, struct {
				Error      string `json:"error"`
				ServerDate string `json:"server_date"`
			}{ServerDate: time.Now().UTC().Format(dateFormat), Error: "request Date timeout"})
			return
		}
		nonce, ok := auth["Nonce"]
		if !ok {
			response.Message(r.Context(), w, msg.ErrAccess, "missing Authorization Nonce")
			return
		}
		sign, ok := auth["Signature"]
		if !ok {
			response.Message(r.Context(), w, msg.ErrAccess, "missing Authorization Signature")
			return
		}
		aki, ok := auth["Credential"]
		if !ok {
			response.Message(r.Context(), w, msg.ErrAccess, "missing Authorization Credential")
			return
		}
		// FIXME: get keySecret
		aks := os.Getenv("AK_" + aki)

		// 3. SignedBody
		signedBody, err := createSignedBody(date, nonce, r)
		if err != nil {
			response.Message(r.Context(), w, msg.ErrSignature, err.Error())
			return
		}

		// 4. Calculating sign hmac-sha256
		h := hmac.New(sha256.New, []byte(aks))
		h.Write([]byte(signedBody))
		signature := hex.EncodeToString(h.Sum(nil))

		// 5. Check
		if signature != sign {
			response.Data(r.Context(), w, msg.ErrSignature, struct {
				Error      string `json:"error"`
				SignedBody string `json:"signed_body"`
				Refer      string `json:"refer"`
			}{SignedBody: signedBody, Error: "signature is not matched"})
			return
		}

		// 6. Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func parseAuthorization(r *http.Request) (map[string]string, error) {

	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil, fmt.Errorf("missing Authorization")
	}
	l := len(authorizationPrefix)
	if auth[:l] != authorizationPrefix {
		return nil, fmt.Errorf("error Authorization prefix")
	}

	data := make(map[string]string)
	for _, tk := range strings.Split(auth[l+1:], ",") {
		if !strings.Contains(tk, "=") {
			continue
		}
		s := strings.Split(strings.TrimSpace(tk), "=")
		data[s[0]] = s[1]
	}

	return data, nil
}

func checkDatetime(date string) error {

	dt, err := time.Parse(dateFormat, date)
	if err != nil {
		return err
	}
	if math.Abs(float64(time.Now().Unix()-dt.Unix())) > requestTimeout {
		return fmt.Errorf("timeout")
	}
	return nil
}

func createSignedBody(date, nonce string, r *http.Request) (string, error) {

	// 1. HTTPMethod CanonicalURI
	signedBody := date + "\n" + nonce + "\n"
	signedBody += r.Method + "\n" + r.URL.EscapedPath() + "\n"

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
	b, _ := io.ReadAll(reader)
	r.Body = io.NopCloser(bytes.NewReader(b)) // reuse body
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
