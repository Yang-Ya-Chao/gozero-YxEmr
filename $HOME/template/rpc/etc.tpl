Name: {{.serviceName}}.rpc
ListenOn: 0.0.0.0:9000
DataSourceName: server=192.168.200.66\cx1;user id=sa;password=123qwe,.;database=yxhis;encrypt=disable
Log:
  Mode: file
DBLog : false
Endpoints:
    Key: {{.serviceName}}.rpc # 直连服务注册key
Cache:
  Hosts:
  - 127.0.0.1:6379