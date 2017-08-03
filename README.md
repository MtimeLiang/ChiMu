# ChiMu

```
Error 1136: Column count doesn't match value count at row 1

插入的数据的个数跟列的个数不匹配，或者Sql语句中'?'的个数不对
```

```
关于获取数据库中内容并转换成model时可能有的字段不能正常解析，
https://beego.me/docs/mvc/model/models.md

默认的命名规则，使用驼峰转蛇形：

AuthUser -> auth_user
Auth_User -> auth__user
DB_AuthUser -> d_b__auth_user
除了开头的大写字母以外，遇到大写会增加 _，原名称中的下划线保留。
```

```
column

为字段设置 db 字段的名称
Name string `orm:"column(user_name)"`
```


#### beego中文文档
[https://beego.me/docs/mvc/model/object.md](https://beego.me/docs/mvc/model/object.md)

#### jwt、beego三方
[https://github.com/juusechec/jwt-beego#ejemplo-de-uso](https://github.com/juusechec/jwt-beego#ejemplo-de-uso)