# SCN-01 测试报告

## 本次范围

日期：2026-09-10。

本次运行的是 Go 单元测试和模拟 CLI，不访问真实店铺、代理凭据、Cookie 或外部后台。环境为 macOS arm64、Go 1.26.2；构建缓存定向到临时目录。

命令：

```sh
env GOCACHE=/private/tmp/cross-border-net-scheduler-go-cache go test ./...
go run ./cmd/scn01-demo
```

结果：`go test ./...` 通过；`scn01-demo` 输出模拟任务的动作状态 `succeeded` 与业务状态 `recovered`。

## 覆盖结果

| 代码用例 | 测试计划 | 覆盖需求 | 结果 | 证据等级 |
|---|---|---|---|---|
| `TestTaskRequiresSpecificInput` | TC-01 | REQ-OBS-01 | 通过：缺少输入时停止执行 | 模拟单元测试 |
| 未单列代码用例 | TC-02 | REQ-OBS-03、REQ-VERIFY-01 | 待实现 | — |
| 未单列代码用例 | TC-03 | REQ-OBS-02 | 待实现 | — |
| `TestNonNetworkEvidenceDoesNotExecute` | TC-05 | REQ-DIAG-01、REQ-ACT-03 | 通过：非网络证据不执行网络动作 | 模拟单元测试 |
| 未单列代码用例 | TC-06 | REQ-DIAG-02 | 待实现 | — |
| `TestIneligibleEgressDoesNotExecute` | TC-07 | REQ-ACT-01 | 通过：出口前后状态不一致时拒绝动作 | 模拟单元测试 |
| `TestCancellationDoesNotExecute` | TC-08 | REQ-ACT-02 | 通过：取消确认后零执行 | 模拟单元测试 |
| 模拟 CLI | TC-09 | REQ-ACT-01、REQ-ACT-02、REQ-VERIFY-01、REQ-VERIFY-02 | 通过：模拟动作与模拟业务恢复分别记录 | 模拟运行 |
| `TestActionSuccessDoesNotImplyRecovery` | TC-10 | REQ-VERIFY-02、REQ-ACT-03 | 通过：动作成功不等于业务恢复 | 模拟单元测试 |
| `TestActionFailureIsRecorded` | TC-11 | REQ-ACT-02、REQ-VERIFY-02 | 通过：动作失败和业务未知分别记录 | 模拟单元测试 |
| `TestRevisionRequiresReason` | TC-12 | REQ-VERIFY-03 | 通过：更正需理由并保留记录 | 模拟单元测试 |
| 未单列代码用例 | TC-13 | NFR-EVID-01、NFR-PLAT-01、NFR-REP-01、NFR-UX-01 | 待实现 | — |

## 未验证与限制

- 没有真实 Mihomo/Clash 或其他适配器接入，`reconnect_current_egress` 只是模拟动作。
- 没有获取真实出口的前后状态，不能证明保持出口。
- 没有真实后台目标、登录态或卖家任务，不能证明业务恢复。
- 真实适配器验证前，产品只能称为“SCN-01 状态机与安全闸门原型”。

下次测试应先实现模拟 TC-02、TC-03、TC-04、TC-06、TC-13，再在脱敏授权环境中完成适配器能力探测和前后出口比较。
