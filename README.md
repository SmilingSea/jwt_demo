# hertz中使用jwt鉴权以及接口防刷功能实现demo

目前是使用的Hertz官方提供的jwt工具`github.com/hertz-contrib/jwt`进行鉴权，这个demo没有使用数据库，仅仅用来学习和测试可行性，同时使用Golang官方提供的限流算法`golang.org/x/time/rate`来实现接口防刷。全过程基本实在使用现成的轮子。

接口文档：https://github.com/SmilingSea/jwt_demo/blob/master/jwt_demo.md