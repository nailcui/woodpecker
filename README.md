woodpecker - 服务健康检查
---

### 功能

- [x] `http` 健康检查
  - [x] `https 自动忽略证书
  - [x] 支持登录认证接口
  - [x] 支持超时时间配置
  - [x] 支持状态码校验
  - [x] 支持`response` 校验, 必须包含 or 必须不包含
- [x] `tcp` 健康检查
  - [x] 支持超时时间配置
- [ ] 策略
  - [x] 支持`yaml` 文件配置各个接口
  - [x] 支持`cron` 配置健康检查触发时机
  - [ ] 失败、成功告警阈值配置
- [x] 告警渠道
  - [x] 钉钉
  - [ ] 企业微信
- [ ] 告警消息
  - [ ] 支持在配置中指定消息模版
  - [ ] 支持在配置中指定模版所需变量

### 配置

> - 默认加载路径: ./resource.d

```yaml
# cat check_http_example.yaml
apiVersion: v1alpha1
kind: checker
metadata:
  # name 在所有 checker 中唯一
  name: "http-dive"
  labels:
    a: b
spec:
  # http 健康检查
  kind: http
  spec:
    # 指向下方的钉钉通知配置
    notifier: dingtalk-base
    # true: 禁用 false: 启用
    disable: false
    # 失败时，连续告警次数；超过3次不再告警；有恢复通知
    reportTimes: 3
    # url的请求超时时间，默认5000，单位: ms
    timeout: 5000
    url: "http://httpbin.org/get"
    cron: "0/5 * * * * ?"
    # url的返回值必须包含的字符串, 数字类型的必须用引号
    mustContain:
      - "green"
      - "10010"
    # url的返回值一定不能包含的字符串, 数字类型的必须用引号
    mustNotContain:
      - "red"
    # http 成功的状态码
    successCode:
      - "200"

---
apiVersion: v1alpha1
kind: notifier
metadata:
  # name 在所有 notifier 中唯一
  name: dingtalk-base
  labels:
    a: b
spec:
  # 钉钉群告警
  kind: dingtalk
  spec:
    # 钉钉群配置
    accessToken: 30a3f2336exxxxxx
    secret: SEC51866cxxxxx

```
