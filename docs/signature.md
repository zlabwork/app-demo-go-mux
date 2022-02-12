## Authorization Header

```
Authorization: Credential=AKIZ9SIKFWLQ0J8M,Signature=a6bb707d2903b84dc3de6de1f3b49997872176f7b0da51ac2124c6e09d15a413
```

## 格式

```
<HTTPMethod>\n
<CanonicalURI>\n
<CanonicalQueryString>\n
<CanonicalHeaders>\n
<HashedPayload>
```

## 签名介绍

1. HTTPMethod  
   e.g. GET, PUT, HEAD, and DELETE.


2. CanonicalURI  
   请求地址 URI , 不包括查询字符串(即 "?" 以及后面的字符串).  
   e.g. `/api/users`


3. CanonicalQueryString  
   查询字符串，需要按字母升序排列, 其中 key 与 name 均需要 URLEncode  
   e.g. `age=34&name=Joe`


4. CanonicalHeaders  
   所有请求头 name 需要小写，并且按字母升序排列   
   name 与 value 分隔使用 `:`, 去空格  
   必须包含的请求头名称:  
   a. Host  
   b. Content-Type  
   c. 所有 X-Lab-* 的请求头 ( 其中必须包含 `X-Lab-Date` 建议包含 `X-Lab-Content-Sha256` `X-Lab-Nonce`)  

```
Lowercase(<HeaderName1>)+":"+Trim(<value>)+"\n"  
Lowercase(<HeaderName2>)+":"+Trim(<value>)+"\n"  
```

e.g.

```
content-type:text/html
host:api.zlab.dev
x-lab-content-sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
x-lab-date:20170628T011555Z
x-lab-nonce:Y0fleXNf
```

5. HashedPayload  
   `Hex(SHA256Hash(<payload>))` 其中 payload 可为空字符串  
   e.g. `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`


### 参考 signedBody

```
GET
/api/users
age=34&name=Joe
content-type:text/html
host:api.zlab.dev
x-lab-content-sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
x-lab-date:20170628T011555Z
x-lab-nonce:Y0fleXNf
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
```


### 签名 hmac-sha256
   Signature = `HMAC-SHA256(<signedBody>)`  
   ```language
   # Demo
   # key = "ImXgsvndC6roCIY91exhIaOsR8UQcm09"
   # signedBody = "<above-signedBody-data>"
   signature = a6bb707d2903b84dc3de6de1f3b49997872176f7b0da51ac2124c6e09d15a413
   ```

### 拼接 Authorization  
```
# Credential=AKIZ9SIKFWLQ0J8M  
# Signature=a6bb707d2903b84dc3de6de1f3b49997872176f7b0da51ac2124c6e09d15a413  
Authorization: Credential=AKIZ9SIKFWLQ0J8M,Signature=a6bb707d2903b84dc3de6de1f3b49997872176f7b0da51ac2124c6e09d15a413  
```


## 参考文档

https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-header-based-auth.html  
https://help.aliyun.com/document_detail/31951.html  
