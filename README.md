# lark

通过**飞书多维表格(Bitable)**管理**Gitlab Issue**

## 配置说明

### 飞书配置

1. 新建飞书机器人应用，获取 `LARK_APP_ID` & `LARK_APP_SECRET` [飞书官方文档](https://open.feishu.cn/document/home/index)
![feishu-app](https://raw.githubusercontent.com/geeklubcn/lark/master/.design/feishu-app.png)
2. 机器人应用添加bitable相关权限 注：需要发布后生效
![feishu-permission](https://raw.githubusercontent.com/geeklubcn/lark/master/.design/feishu-permission.png)
3. 新建多维表格，获取 `LARK_BITABLE_APP_TOKEN`。可在浏览器url中获取
![feishu-bitable-app-token](https://raw.githubusercontent.com/geeklubcn/lark/master/.design/feishu-bitable-app-token.png)
4. 多维表格中添加机器人应用 注： 配合第2步中的权限，两者缺一不可。 https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN
![feishu-bitable-add-app](https://raw.githubusercontent.com/geeklubcn/lark/master/.design/feishu-bitable-add-app.png)

### Gitlab配置

1. 在profile->Access Tokens生成个人token，作为 `LARK_GITLAB_TOKEN` 注： 必需勾选api scope，推荐全部勾选
![gitlab-token](https://raw.githubusercontent.com/geeklubcn/lark/master/.design/gitlab-token.png)
2. `LARK_GITLAB_DOMAIN` 为 gitlab 域名，实际访问API会添加API版本前缀，如 https://gitlab.com/api/v4/issues

## 使用方式

### 方式一: docker compose

```yaml
version: '3'
services:
  lark:
    image: geeklubcn/lark:v20220408
    environment:
      - LARK_APP_ID=cli_aaa
      - LARK_APP_SECRET=6nbbb
      - LARK_BITABLE_APP_TOKEN=bascnE9ccc
      - LARK_GITLAB_DOMAIN=https://gitlab.com
      - LARK_GITLAB_TOKEN=glpat-RGA_ddd
      # filter issue by label. SYNC_ALL will fetch all issue
      - LARK_GITLAB_ISSUE_LABEL=SYNC_ALL
      # LARK_SYNC_PERIOD default 60s
      - LARK_SYNC_PERIOD=10s
      # LARK_LOG_LEVEL default info
      - LARK_LOG_LEVEL=debug
```

编辑好参数后命令行执行

```shell
docker-compose up 
```

