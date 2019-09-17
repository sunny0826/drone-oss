# drone-oss
![](https://img.shields.io/docker/cloud/automated/guoxudongdocker/drone-oss.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/guoxudongdocker/drone-oss.svg)

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