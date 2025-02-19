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

| Language     | files | blank | comment | code  |
| ------------ | ----- | ----- | ------- | ----- |
| Go           | 181   | 7355  | 421     | 53981 |
| Thrift       | 8     | 79    | 0       | 341   |
| Bourne Shell | 12    | 52    | 0       | 198   |
| XML          | 4     | 0     | 0       | 192   |
| YAML         | 7     | 7     | 0       | 86    |
| Markdown     | 1     | 26    | 0       | 75    |
| make         | 2     | 5     | 1       | 10    |
| JSON         | 1     | 0     | 0       | 1     |
| Text         | 1     | 0     | 0       | 1     |
| SUM:         | 217   | 7524  | 422     | 54885 |

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

## API requests

user register

```powershell
curl --location --request POST '/douyin/user/register/?username=&password=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "user_id": 0,
    "token": "string"
}
```

user login

```powershell
curl --location --request POST '/douyin/user/login/?username=&password=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "user_id": 0,
    "token": "string"
}
```

infomation of user

```powershell
curl --location --request GET '/douyin/user/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "user": {
        "id": 0,
        "name": "string",
        "follow_count": 0,
        "follower_count": 0,
        "is_follow": true,
        "avatar": "string",
        "background_image": "string",
        "signature": "string",
        "total_favorited": "string",
        "work_count": 0,
        "favorite_count": 0
    }
}
```

get video stream

```powershell
curl --location --request GET '/douyin/feed/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "next_time": 0,
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
```

logged in user selects video to upload

```powershell
curl --location --request POST '/douyin/publish/action/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--form 'data=@""' \
--form 'token=""' \
--form 'title=""'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

list all videos contributed by the user

```powershell
curl --location --request GET '/douyin/publish/list/?token=&user_id=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
```

all liked videos by user

```powershell
curl --location --request GET '/douyin/favorite/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
```

logged in user to comment on video

```powershell
curl --location --request POST '/douyin/comment/action/?token=&video_id=&action_type=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment": {
        "id": 0,
        "user": {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        },
        "content": "string",
        "create_date": "string"
    }
}
```

logged in user to comment on video

```powershell
curl --location --request POST '/douyin/comment/action/?token=&video_id=&action_type=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment": {
        "id": 0,
        "user": {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        },
        "content": "string",
        "create_date": "string"
    }
}
```

view all comments on video

```powershell
curl --location --request GET '/douyin/comment/list/?token=&video_id=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment_list": [
        {
            "id": 0,
            "user": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "content": "string",
            "create_date": "string"
        }
    ]
}
```

follow

```powershell
curl --location --request POST '/douyin/relation/action/?token=&to_user_id=&action_type=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

follow list

```powershell
curl --location --request GET '/douyin/relation/follow/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```

follower list

```powershell
curl --location --request GET '/douyin/relation/follower/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```

friend list

```powershell
curl --location --request GET '/douyin/relation/friend/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```

send message

```powershell
curl --location --request POST '/douyin/message/action/?token=&to_user_id=&action_type=&content=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

message list

```powershell
curl --location --request GET '/douyin/message/chat/?token=&to_user_id=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "message_list": [
        {
            "id": 0,
            "to_user_id": 0,
            "from_user_id": 0,
            "content": "string",
            "create_time": 0
        }
    ]
}
```

## Screenshot

<img src="https://github.com/T4t4KAU/douyin/blob/main/image/image1.png?raw=true" alt="image1.png" style="width:30%; height:auto;">
<br>
<img src="https://github.com/T4t4KAU/douyin/blob/main/image/image2.png?raw=true" alt="image2.png" style="width:30%; height:auto;">

## Give a star! ⭐

If you think this project is interesting, or helpful to you, please give a star!