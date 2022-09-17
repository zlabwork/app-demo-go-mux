## Authorization Header

```
Authorization: ZLAB Credential=AKIZ9SIKFWLQ0J8M, Date=20220917T165312Z, Nonce=a15a389f9e2f1582, Signature=a6bb707d2903b84dc3de6de1f3b49997872176f7b0da51ac2124c6e09d15a413
```

## 格式

```
<Date>\n
<Nonce>\n
<HTTPMethod>\n
<CanonicalURI>\n
<CanonicalQueryString>\n
<CanonicalHeaders>\n
<HashedPayload>
```

e.g. 1
```text
20220917T165350Z
20f8e8834a6e68f0
GET
/users/123456
name=joe&type=1
content-type:text/html
host:zlab.dev
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
```

e.g. 2
```text
20220917T165312Z
a15a389f9e2f1582
GET
/users/123456
name=joe&type=1
content-type:text/html
host:zlab.dev
x-lab-content-sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
x-lab-date:20220917T165312Z
x-lab-nonce:a15a389f9e2f1582
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
```

## 签名介绍

1. Date  
   UTC 时间，格式 20060102T150405Z0700  
   e.g. 20220917T165312Z


2. Nonce  
   随机字符串 A-Za-z0-9


3. HTTPMethod  
   e.g. GET, PUT, HEAD, and DELETE.


4. CanonicalURI  
   请求地址 URI , 不包括查询字符串(即 "?" 以及后面的字符串).  
   e.g. `/api/users`


5. CanonicalQueryString  
   查询字符串，需要按字母升序排列, 其中 key 与 name 均需要 URLEncode  
   e.g. `age=34&name=Joe`


6. CanonicalHeaders  
   所有请求头 name 需要小写，并且按字母升序排列   
   name 与 value 分隔使用 `:`, 去空格  
   必须包含的请求头名称:  
   a. Host  
   b. Content-Type  
   c. 所有 X-Lab-* 的请求头 (建议包含 `X-Lab-Content-Sha256` `X-Lab-Date` `X-Lab-Nonce`)  

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
20220917T171905Z
ee20793474e82dbf
GET
/api/users
age=34&name=Joe
content-type:text/html
host:zlab.dev
x-lab-content-sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
x-lab-date:20220917T171905Z
x-lab-nonce:ee20793474e82dbf
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
```


### 签名 hmac-sha256
   Signature = `HMAC-SHA256(<signedBody>)`  
   ```language
   # Demo
   # key = "ImXgsvndC6roCIY91exhIaOsR8UQcm09"
   # signedBody = "<above-signedBody-data>"
   signature = 707732d6a997df65d73dfea193a9b7d66162b1754afb2419b0dd31c9bbda328a
   ```

### 拼接 Authorization  
```
# Credential=AKIZ9SIKFWLQ0J8M  
# Signature=707732d6a997df65d73dfea193a9b7d66162b1754afb2419b0dd31c9bbda328a  
Authorization: ZLAB Credential=AKIZ9SIKFWLQ0J8M, Date=20220917T171905Z, Nonce=ee20793474e82dbf, Signature=707732d6a997df65d73dfea193a9b7d66162b1754afb2419b0dd31c9bbda328a  
```


## 参考文档

https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-header-based-auth.html  
https://help.aliyun.com/document_detail/31951.html  
