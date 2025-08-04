# LDAP ADM

这是一个轻量管理ldap账号的信息的cli命令行工具，无需编辑复杂的ldif文件，即可快速管理user，group等信息。

### 特点：
1、使用简单  
2、无需复杂配置  
3、快速部署  

### 安装方式
1、下载并解压文件  
```shell
wget https://github.com/ELSTHub/ldapadm/releases/download/v0.1.0/linux-v0.1.0.tar.gz
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

### 配置LDAP ADM
```shell
# 编辑配置文件 vim /etc/ldapadm/ldapadm.yaml ，可根据实际情况填写
ldap_server_conf:
  host: 127.0.0.1 # LDAP服务器地址
  port: 389 # LDAP服务端口号
  base: dc=elst,dc=dev # LDAP Base DN
  login_dn: cn=admin,dc=elst,dc=dev # LDAP管理员用户
  password: Elst@2025 # LDAP管理员密码
  user_dn: ou=People,dc=elst,dc=dev # 用户所属OU
  group_dn: ou=Group,dc=elst,dc=dev # 用户组所属OU
  default_home_path: /home # 用户默认HOME目录
  default_bash: /bin/bash # 用户默认bash
  password_encryption: SHA1 # 用户密码加密算法（MD5、SHA1）

ldap_adm:
  uid: /etc/ldapadm/UID # 下一个UID 编号
  gid: /etc/ldapadm/GID # 下一个GID 编号
  uid_lock_file: /tmp/ldapadm_uid.lock # UID生成锁
  gid_lock_file: /tmp/ldapadm_gid.lock # GID 生成锁
  min_gid: 1000 # 最小GID
  max_gid: 65535 # 最大GID
  min_uid: 1000 # 最小UID
  max_uid: 65535 # 最大UID
```
