# LDAP ADM

这是一个轻量管理ldap账号的信息的cli命令行工具，无需编辑复杂的ldif文件，即可快速管理user，group，ou等信息。

### 特点：
1、使用简单  
2、无需复杂配置  
3、快速部署  

### 安装方式
1、下载并解压文件  
```shell
wget 
```
2、执行安装脚本
```shell
bash install.sh
```

### 编译构建
```shell
git clone https://github.com/ELSTHub/ldapadm.git
cd ldapadm
go get .

# Windows
./scripts/build.bat

# Linux 
# bash ./scripts/build.sh
```