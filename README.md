# drone-oss

drone plugin of oss

```yaml
- name: Upload OSS
  image: guoxudongdocker/drone-oss
  settings:
    dist: dist                              # dist package
    path: kk-k8s-oss/devops                 # bucket/object
    endpoint: oss-cn-shanghai.aliyuncs.com  # oss endpoint
    access_key_id: 
      from_secret: ACCESS_KEY_ID
    access_key_secret: 
      from_secret: ACCESS_KEY_SECRET
```