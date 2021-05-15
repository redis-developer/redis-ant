

![RedisAnt Logo](https://raw.githubusercontent.com/ramantehlan/redis-ant/main/assets/image/redisAnt.png?token=AG5RGAE6SVY3MTACL6YJGQTAVAOF2)

![](https://goreportcard.com/badge/github.com/ramantehlan/redisant)
![](https://img.shields.io/badge/godoc-reference-green)
![](https://img.shields.io/badge/license-MIT-blue)

# RedisAnt

RedisAnt maintains your client-side cache end-to-end at the speed of light. It connects with your MongoDB database and automatically ingests the requested data in your local cache and keeps the cache fresh as a daisy.


# Index

- [About](#about)
  - [RedisAnt Server](#redisant-server)
  - [RedisAnt Client](#redisant-client)
- [Usage](#usage)
  - [Pre-Requisites](#pre-requisites)
  - [Development Environment](#development-environment)
  - [Build](#build)
  - [Deploy](#deploy)
- [FAQ](#faq)
- [Acknowledgment](#acknowledgment)
- [License](#license)

# About

There are 2 main services:

1. RedisAnt-Server
2. RedisAnt-Client


It helps you maintain your client side cache end-to-end an
d automatically. You just define the MongoDB `database`, `collection`, `fields` and RedisAnt sync and keep your cache fresh as a daisy.




You just define your source(Database), and RedisAnt will a
utomatically sync and update your local cache
Automatic client-side caching using RedisJSON and MongoDB.

### RedisAnt Server

### RedisAnt Client

# Usage

### Pre-Requisites

This pre-requisites are not necessarily for running the project, but if you plan to use or contribute to this project or play with the source code, knowledge of following things is recommended.

- [Golang](https://golang.org/)
- [Redis](https://redis.io/)

### Development Environment

To develop or build this project, make sure you have the following environment setup:

- Install, and setup Go environment.
- Install python.
- Install NodeJS and yarn.
- Install Make.
- Clone this project in your workspace.

Once you have set up the above environment, we will use make to install dependencies. Go to the root of the project and run the following command; it installs all the node modules, Python and Go dependencies.

```sh
$ make setup
```

### Build

### Deploy


# FAQ


# Acknowledgement


- Logo is from [Flaticon](https://www.flaticon.com/free-icon/ant_809140?term=ants&page=1&position=6&page=1&position=6&related_id=809140&origin=search).
- RedisLab for free access to Redis Enterprise Cloud.
- All the open-source contributors whose code was used in this project. Thank you. : )



# License

**MIT License**

Copyright (c) 2021 Raman Tehlan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
