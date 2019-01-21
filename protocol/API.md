# API接口文档

## 用户协议概述 
| 操作              | URL       | Method    |SC |
| :----------------| :----------|:-------:|:------:|
|创建(注册)用户     |/user         | POST    | 201,400,500|
|用户登录           |/user/:username| POST   | 200,400,500|
|获取用户基本信息    |/user/:username|GET     | 200,400,401,403,500|
|用户注销           |/user/:username|DELETE | 204,400,401,403,500|

## 用户资源协议概述

| 操作              | URL                           | Method    |SC |
| :----------------| :------------------------------|:--------:|:------:|
|用户视频列表       |/user/:username/videos          |GET       |200,400,500|
|获取一个视频       |/user/:username/videos/:vid-id  |GET      |200,400,500|
|删除一个视频       |/user/:username/videos/:vid-id  |DELETE   |204,400,401,403,500|

## 评论协议概述
| 操作              | URL                          | Method    |SC |
| :-----------------| :----------------------------|:--------:|:------:|
|显示评论            |/videos/:vid-id/comments     |GET        |200,400,500|
|发表一个评论        |/videos/:vid-id/comments      |POST       |201,400,500|
|删除一个评论        |/videos/:vid-id/comment/:comment-id|DELETE|204,400,401,403,500|
