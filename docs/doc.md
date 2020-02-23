# 设备登录过程方式1
1. 通过http api 或者扫码登录方式获得token.
2. 上传设备信息并获取到设备id.
3. 用设备id登录tcp server.

# 设备登录过程方式2
1. 通过http api 或者扫码登录方式获得带user id 的 JWT token
2. 客户端上传设备信息.
3. 客户端用token登录tcp server.
