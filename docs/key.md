## Create Key Pair
```bash
# use for RSA encrypt & decrypt
# -----BEGIN RSA PRIVATE KEY-----
openssl genrsa 2048 > private.key
openssl rsa -in private.key -pubout > public.key
```

```bash
# use for ssh
# -----BEGIN OPENSSH PRIVATE KEY-----
ssh-keygen
ssh-keygen -t rsa -b 4096
```
