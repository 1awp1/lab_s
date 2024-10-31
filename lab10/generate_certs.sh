#!/bin/bash

# Генерация сертификата сервера
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/C=RU/ST=Moscow/L=Moscow/O=My Company/OU=IT Department/CN=localhost"
openssl x509 -req -in server.csr -signkey server.key -out server.crt -days 365

# Генерация сертификата CA
openssl genrsa -out ca.key 2048
openssl req -new -key ca.key -out ca.csr -subj "/C=RU/ST=Moscow/L=Moscow/O=My Company/OU=CA/CN=RootCA"

# Создание файла с расширениями для CA
cat > extensions.cnf <<EOF
[ v3_ca ]
subjectAltName = DNS:localhost
EOF

# Генерация ca.crt с использованием extensions.cnf
openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt -days 365 -extensions v3_ca -extfile extensions.cnf

# Преобразование сертификатов в формат PEM
openssl x509 -in server.crt -outform PEM -out server.crt 
openssl x509 -in ca.crt -outform PEM -out ca.crt
