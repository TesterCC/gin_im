# MongoDB Collection Design

## user_basic collection

```text
{
    "account": "账号",  // admin
    "password": "密码",  // admin password md5 hash
    "nickname": "昵称",
    "gender": "性别",   // 0-unknown, 1-male, 2-female  // video use sex
    "email": "邮箱",
    "avatar": "头像",
    "created_at": 1,   // create time
    "updated_at": 1    // update time
}
```

## message_basic collection

```text
{
    "user_id":"用户唯一标识",   // video use user_identity
    "room_id":"房间唯一标识",   // video use room_identity
    "data":"发送的数据",
    "created_at":1,    
    "updated_at":1  
}
```

## room_basic collection
```text
{
    "number":"房间号",
    "name":"房间名称",
    "info":"房间简介",
    "user_id":"房间创建者的唯一标识",
    "created_at":1,     // room create time
    "updated_at":1      // room update time
}
```

## user_room collection
```text
{
    "user_id":"用户唯一标识",  // video use user_identity
    "room_id":"房间唯一标识",  // video use room_identity
    "message_id":"消息唯一标识",   // video use message_identity
    "created_at":1,          // create time
    "updated_at":1           // update time
}
```