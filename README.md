# asap

## 安装Redis
brew install redis
打开 redis.conf 文件，然后按 command + f 进行搜索：#requirepass foobared
修改为：
requirepass 你的密码

/usr/local/bin/redis-server /usr/local/etc/redis.conf
