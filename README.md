# Simple Douyin

## Introduction

A minimalist tiktok, using a distributed microservice architecture based on **Kitex** and **Hertz**

### Use Basic Features

- Middleware、Rate Limiting、Request Retry、Timeout Control、Connection Multiplexing
- Message Queue
  - use **RabbitMQ** for asynchronous communication and module decoupling
- Memory Cache
  - use **Redis** to cache hot data

- Tracing
  - use **jaeger** to tracing
- Customized BoundHandler
  - achieve CPU utilization rate customized bound handler
- Service Discovery and Register
  - use [registry-etcd](https://github.com/kitex-contrib/registry-etcd) to discovery and register service

### catalog introduce

| catalog        | introduce                |
| :------------- | :----------------------- |
| pkg/constants  | constant                 |
| pkg/bound      | customized bound handler |
| pkg/errno      | customized error number  |
| pkg/middleware | RPC middleware           |
| pkg/tracer     | init jaeger              |
| dal            | db operation             |
| pkg            | data pack                |
| service        | business logic           |

### code count
| Language      | files | blank | comment | code  |
|---------------|-------|-------|---------|-------|
| Go            | 181   | 7355  | 421     | 53981 |
| Thrift        | 8     | 79    | 0       | 341   |
| Bourne Shell  | 12    | 52    | 0       | 198   |
| XML           | 4     | 0     | 0       | 192   |
| YAML          | 7     | 7     | 0       | 86    |
| Markdown      | 1     | 26    | 0       | 75    |
| make          | 2     | 5     | 1       | 10    |
| JSON          | 1     | 0     | 0       | 1     |
| Text          | 1     | 0     | 0       | 1     |
| SUM:          | 217   | 7524  | 422     | 54885 |

## Quick Start

### 1. Setup Basic Dependence

```shell
docker-compose up
```

### 2. Run User RPC Server

```shell
cd cmd/user
sh build.sh
go run output/bin
```

### 3. Run Publish RPC Server

```shell
cd cmd/publish
sh build.sh
go run output/bin
```

### 4. Run Comment RPC Server

```shell
cd cmd/comment
sh build.sh
go run output/bin
```

### 5. Run Favorite RPC Server

```shell
cd cmd/favorite
sh build.sh
go run output/bin
```

### 6. Run Message RPC Server

```shell
cd cmd/message
sh build.sh
go run output/bin
```

### 7. Run Relation RPC Server

```shell
cd cmd/relation
sh build.sh
go run output/bin
```

### 8. Run API Server

```shell
cd cmd/api
chmod +x run.sh
go run api
```

### 9. Jaeger

visit `http://127.0.0.1:16686/` on browser.

## screenshot

<img src="https://github.com/T4t4KAU/douyin/blob/main/image/image1.png?raw=true" alt="image1.png" style="width:30%; height:auto;">
<br>
<img src="https://github.com/T4t4KAU/douyin/blob/main/image/image2.png?raw=true" alt="image2.png" style="width:30%; height:auto;">

## Give a star! ⭐
If you think this project is interesting, or helpful to you, please give a star!