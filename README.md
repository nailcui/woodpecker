woodpecker
---

健康检查

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
