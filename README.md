

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



### RedisAnt Server

### RedisAnt Client

![RedisAnt Diagram](https://raw.githubusercontent.com/ramantehlan/redis-ant/main/assets/image/technical-diagram.png?token=AG5RGAEZCOCQAJOBF7RHRM3AVB3WK)

# Usage

### Pre-Requisites

This pre-requisites are not necessarily for running the project, but if you plan to use or contribute to this project or play with the source code, some knowledge of following things is recommended.

- [Golang](https://golang.org/)
- [Redis](https://redis.io/)
- [Docker](https://www.docker.com/)
- Make

### Development Environment

To develop or build this project, make sure you have the following environment setup:

- Install, and setup Go environment.
- Install Docker.
- Install Make.
- Clone this project in your **go workspace**.

Once you have set up the above environment, we will use make to install dependencies. Go to the root of the project and run the following command; it installs all the Go dependencies.

```sh
$ make setup
```

### Build

### Deploy


# FAQ

#### What are some use cases for a tool like this?
- For user service to store user details and serve them fast.
- For storing device_id in notification service to send notifications
- Fewer updates comparatively. No transformation/operation required.
- It will automate and reduce the time it takes to set up the syncing.

#### What is the future scopes of this project?

This project should be considered a pre-alpha and a lot can be improved in it. Here are something I feel can be added or improved:

- We can enable users to select sub-fields they want to sync from Mongo.
- Increase support for multiple collections and databases.
- The same can be applied to the global cache.

#### Why the name Redis-Ant?
Ants are known for working in teams in sync. If one ant finds food, it very quickly transfers the location data to all the other ants. Then all the ants work together to take food from source to destination, bit by bit. I think this tool also works(at least that's what I wanted) in the same way,  hence the name RedisAnt.

#### Why RedisJSON and MongoDB?
MongoDB is a no-SQL database and supports JSON document structure. It is a really popular choice for startups(enterprises too) to create their ever-evolving data structure DB.

RedisJSON adds native support for storing and retrieving JSON documents at the speed of Redis.

RedisAnt combines the power of MongoDB and RedisJSON to automate the local caching with zero overhead and high performance.

#### How is it different from Client Tracking feature?

The client tracking feature of Redis Enterprise helps companies manage client-side cache assisted by the server-side cache. It required some initial configuration and the server needs to be programmed accordingly. Also, it is obnoxious about the database.

RedisAnt is aware of the database, and it communicates with it on your behalf. It manages client-side caching end-to-end without any assistance from the server.

#### When was this project created?

**This project was created during Redis Hackathons 2021**. I had this idea in my head for quite some time and this hackathon provided the perfect space to play with it.

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
