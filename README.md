# 星辰工作室2021招新题- 第四周实践题

- 使用gin提供http server
- 使用者需要在本地localhost有一个mysql数据块，用户root，密码123456。且有一个名为xc的数据库。

# API

共六个API

## add

url: 

http://localhost:8080/add

参数：

| key     | value描述      | 例子           |
| ------- | -------------- | -------------- |
| title   | 帖子标题       | 今天是个好日子 |
| content | 帖子内容       | 真不错         |
| user_id | 发帖者的用户id | 1              |

## getByTitle

url:

http://localhost:8080/getByTitle

参数：

| key   | value描述 | 例子           |
| ----- | --------- | -------------- |
| title | 帖子标题  | 今天是个好日子 |

## getByID

url:

http://localhost:8080/getByID

参数：

| key  | value描述 | 例子 |
| ---- | --------- | ---- |
| key  | 帖子的ID  | 1    |

## getAll

url:

http://localhost:8080/getAll

参数：

无

## delete

url（参数在url里面）：

http://localhost:8080/delete/1

## update

url：

http://localhost:8080/put

参数：

| key     | value描述        | 例子  |
| ------- | ---------------- | ----- |
| key     | 帖子的ID         | 1     |
| content | 帖子更新后的内容 | 更新! |