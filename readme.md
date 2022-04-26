## 什么值得买文章推送器
-----

### 目的
:smile:​ **主要是自己没啥钱，又比较懒，爬取自己想买的商品看看有没有打折，这样就不会错过了！！**
**顺便练习go语法**

### 已实现
- [x] 自定义文章提取规则
- [x] 推送文章
- [x]  去重文章
- [x]  定时推送
- [x] 设定关键字，爬取含关键词的商品
- [x] 利用github Action 自动编译，部署到个人服务器
  
### 待实现


- [ ] 每天定时打卡
- [ ] 配置server酱

### 使用步骤
下载整个代码 window平台直接运行`smzdm.exe`，切勿挪动exe文件，会导致读不到配置
如果想用关键字或者推送自己的钉钉，可以修改配置信息
**配置式：**
修改以下配置，保存配置，再运行`smzdm.exe`即可
```yml
# 搜索关键词
keyWord: 
- 信小兔
- 零食

# 最低评论数
lowCommentNum: 0
# 最低值率
lowWorthyNum: 0
# 满意商品数量
satisfyNum: 10
# 过滤词
filterWords: 
- "榴莲"
- "唯品会"
- "牛奶"
- "电脑"

# 定时任务多长执行一次 单位秒 默认 12个小时
tickTime: 43200
# 钉钉token
dingdingToken: "xxxxx"
```

如果觉得麻烦可以进群，每天都会推送消息哦（钉钉二维码在最下方！！）


### 效果
![image-20220419205742369](https://img.ggball.top/picGo/image-20220419205742369.png)

![image-20220419205914792](https://img.ggball.top/picGo/image-20220419205914792.png)





### 钉钉二维码

![image-20220420200534357](https://img.ggball.top/picGo/image-20220420200534357.png)

