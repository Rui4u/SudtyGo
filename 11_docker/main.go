package main

/*
安装docker & docker-compose

安装docker
curl -fsSL https://get.docker.com | bash -s docker -mirror Aliyun

设置开机启动docker
systemctl enable docker
systemctl start docker
docker ps -a
ps -ef |grep docker

配置阿里云镜像
https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors

	sudo mkdir -p /etc/docker
	sudo tee /etc/docker/daemon.json <<-'EOF'
	{
  	"registry-mirrors": ["https://ovc8iq55.mirror.aliyuncs.com"]
	}
	EOF
	sudo systemctl daemon-reload
	sudo systemctl restart docker

docker ps -a
docker start 8b230cd2da64

docker-compose
https://docs.docker.com/compose/install/standalone/
 curl -SL https://github.com/docker/compose/releases/download/v2.17.3/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose

If the command docker-compose fails after installation, check your path. You can also create a symbolic link to /usr/bin or any other directory in your path. For example:
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

-bash: /usr/local/bin/docker-compose: Permission denied
chmod +x /usr/local/bin/docker-compose


安装mysql
docker pull mysql:5.7

docker images 查看安装镜像

镜像启动
docker run -p 3306:3306 --name mymysql -v $PWD/conf:/etc/mysql/conf.d -v $PWD/log:/logs -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7
-p 3306:3306  将容器的3306端口映射到主机的3306端口
-v $PWD/conf:/etc/mysql/conf.d 将主机当前目录的conf/conf.cnf 挂在到容器的/etc/mysql/conf.d
-v $PWD/log:/logs  主机单签目录下的logs目录挂载到容器的logs目录
-v $PWD/data:/var/lib/mysql  将主机当前牡蛎下的data目录挂载到容器的/var/li/mysql
-e MYSQL_ROOT_PASSWORD=123456 初始化root用户的密码

docker ps -a  查看ID

docker logs 8b230cd2da64 查看启动日志

进入容器
docker exec -it 8b230cd2da64 /bin/bash
登录mysql
mysql -uroot -p123456
建用户并授权
	密码更改为root
 	GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'root' WITH GRANT OPTION;
 	GRANT ALL PRIVILEGES ON *.* TO 'root'@'127.0.0.1' IDENTIFIED BY 'root' WITH GRANT OPTION;
 	GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' IDENTIFIED BY 'root' WITH GRANT OPTION;
	FLUSH PRIVILEGES;
退出
	exit
*/
