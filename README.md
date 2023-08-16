# Simple Douyin

## Introduction

simple douyin

### Use Basic Features

- Middleware、Rate Limiting、Request Retry、Timeout Control、Connection Multiplexing
- Message Queue
  - use **RabbitMQ** for asynchronous communication and module decoupling.
- Memory Cache
  - use **Redis** to cache hot data.

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

## Quick Start

### 1. Setup Basic Dependence

```shell
docker-compose up
```

### 2. Run User RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 3. Run Publish RPC Server

```shell
cd cmd/publish
sh build.sh
sh output/bootstrap.sh
```

### 4. Run Comment RPC Server

```shell
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```

### 5. Run Favorite RPC Server

```shell
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```

### 6. Run Message RPC Server

```shell
cd cmd/message
sh build.sh
sh output/bootstrap.sh
```

### 7. Run Relation RPC Server

```shell
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```

### 8. Run API Server

```shell
cd cmd/api
chmod +x run.sh
./run.sh
```

### 9. Jaeger

visit `http://127.0.0.1:16686/` on browser.