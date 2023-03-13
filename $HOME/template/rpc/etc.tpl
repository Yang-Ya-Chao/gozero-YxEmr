Name: {{.serviceName}}.rpc
ListenOn: 0.0.0.0:9000
DataSourceName: server=192.168.200.66\cx1;user id=sa;password=123qwe,.;database=yxhis;encrypt=disable
Log:
  Mode: file
Endpoints:
    Key: {{.serviceName}}.rpc # 直连服务注册key
#redis缓存
#Cache:
#  - Host: 127.0.0.1:6379