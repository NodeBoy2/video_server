# 数据库表设计

## 用户表
- TABLE : users
- id UNSIGNED INT, PRIMARY KEY, AUTO_INCREMENT
- login_name VARCHAR(65), UNIQUE KEY
- pwd TEXT

    create table users (
        id INT PRIMARY KEY AUTO_INCREMENT, 
        login_name VARCHAR(65) UNIQUE KEY, 
        pwd TEXT);

## 视频资源表
- TABLE : video_info
- id VARCHAR(64), PRIMARY KEY, NOT NULL
- author_id UNSIGNED INT
- name TEXT
- display_ctime TEXT
- create_time DATETIME

    create table video_info(
        id VARCHAR(64) PRIMARY KEY NOT NULL,
        author_id INT,
        name TEXT,
        display_ctime TEXT,
        create_time DATETIME);

## 评论表
- TABLE : comments
- id VARCHAR(64), PRIMARY KEY, NOT NULL
- video_id VARCHAR(64)
- author_id UNSIGNED INT
- content TEXT
- time DATETIME

    create table comments(
        id VARCHAR(64) PRIMARY KEY NOT NULL,
        video_id VARCHAR(64),
        author_id INT,
        content TEXT,
        time DATETIME DEFAULT CURRENT_TIMESTAMP
    );

## 会话表
- TABLE : sessions
- session_id VARCHAR(256), PRIMARY KEY, NOT NULL
- TTL TINYTEXT
- login_name VARCHAR(64)

    create table sessions(
        session_id VARCHAR(256) PRIMARY KEY NOT NULL,
        TTL TINYTEXT,
        login_name VARCHAR(64)
    );