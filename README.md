# 字节跳动青训营 - douyin [项目文档](https://ljxltr3g7w.feishu.cn/docs/doccnLberlBxkQjylBal5I6Tg6g)

# 业务架构
![](https://cover-1312359504.cos.ap-guangzhou.myqcloud.com/How%20We%20Built%20Whimsical%401.100000023841858x.webp)

# common
 存放公共代码：主要有对象存储的上传、下载、删除，token的生成与解析，Kafka所需的消息结构体，公共错误代码等。

# deploy
通过docker-compose.yml部署Kafka、ZooKeeper，Dockerfile是服务的镜像生成文件，具体使用在deploy下的README文件

# service
## api
暴露RESTful API，主要是对接口的访问控制、接口的参数校验、接口的返回结构体等。
## asynqJob
存放异步任务，client-scheduler 调度server执行任务（设置的每10s一次）
## mq
存放消息队列的消费者，将api放在redis中的数据取出后写入mysql中
## rpc-user-info
存放用户信息的rpc接口，完成用户注册、登陆、查看用户信息，主要是对user表的CURD
## rpc-user-operate
存放用户操作的rpc接口，完成点赞、评论、关注操作，主要是对user_favorite_list、user_follow_list、user_comment_list表的CURD
## rpc-video-service
存放对视频的操作的rpc接口，完成视频的发布、拉取、视频流操作，主要是对video表进行CURD

# tpl
是goctl脚手架工具所需的模板文件，因为对原本的模板文件进行了修改，所以在生成时需要指定为该目录下的tpl文件

# 写在最后
感谢以下项目让我学到了很多东西： <br>
[go-zero](https://github.com/zeromicro/go-zero) <br>
向万总和安前松大佬学到了很多东西！谢谢~！每次深夜去找他们请教问题，两位大佬也会很快回复我的消息，真的超感动！

[go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook) [B站up主~](https://space.bilibili.com/389552232) <br>
Mikaelemmmm大佬人超级棒~！ 谢谢~！正如他的签名：“做人但求无愧于心，做事但求无愧于人”，很佩服！

[cos-go-sdk-v5](https://github.com/tencentyun/cos-go-sdk-v5) <br>
腾讯云的技术客服超级耐心~！谢谢~！

```
最后的最后：
    感谢字节跳动的所有工作人员无私的辛苦付出，祝青训营越办越好~
    在这里接触到了很多的大佬！为你们的开源精神致敬！salute~
    我会继续加油，努力追赶你们的脚步
                                                                        2022.7.3
```