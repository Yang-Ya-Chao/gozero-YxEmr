Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
#dtm
Dtm: {{.host}}:{{.port}}
#rpc服务
XXX:
  #使用nacos发现服务
  #Target: nacos://127.0.0.1:8848/add.rpc?namespaceid=public&timeout=5000s
  NonBlock: True
  Endpoints:
    - {{.host}}:{{.port}}
