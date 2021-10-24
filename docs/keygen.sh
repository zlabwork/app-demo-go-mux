#!/bin/bash

work_path='../config/'
cd ${work_path}
openssl genrsa 2048 > private.dev.pem
openssl rsa -in private.dev.pem -pubout > public.dev.pem

openssl genrsa 2048 > private.prod.pem
openssl rsa -in private.prod.pem -pubout > public.prod.pem
