# golang-seed

> :warning: **Still under construction**

[![Build Status](https://travis-ci.org/alobaton/golang-seed.svg?branch=master)](https://travis-ci.org/alobaton/golang-seed)

Provides fast, reliable and extensible starter for the development of Golang projects.

`golang-seed` provides the following features:

- Modularized project.
- i18n support.
- Production and development builds.
- Provides full Docker support for both development and production environment.

## How to start?

Before start update the /etc/hosts file:
```bash
# tools
127.0.0.1       db.dev.local
127.0.0.1       db.prod.local
127.0.0.1       api.dev.local
127.0.0.1       api.prod.local
```
Clone the repository:
```bash
$ git clone --depth 1 https://github.com/alobaton/golang-seed.git
$ cd golang-seed
```

Start the database:
```bash
$ docker-compose -f docker-compose.dev.yml up -d db
```

Start authentication and authorization server:
```bash
$ go run apps/auth/com/auth/main.go
```

