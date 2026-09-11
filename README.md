# cross-border-net-scheduler

钱潮杯 · AI+跨境出海。方向：**电商出海的网络智能调度**。

接入卖家已有宽带、Clash/代理或 SD-WAN，围绕后台操作、同步、直播等经营任务，完成发现异常、判断原因、受控处理和验证恢复。不卖节点和线路；多线路、双出口不是产品前提。

## 什么叫做完

完整项目不是「代码能跑」，而是七段能互相引用。定义见 [docs/PROJECT.md](docs/PROJECT.md)。

| 模块 | 目录 | 钱潮杯最低完成 |
|---|---|---|
| 市场调研 | [docs/research/](docs/research/) | 客群、竞品分层、痛点三分、证据台账 |
| 需求分析 | [docs/requirements/](docs/requirements/) | 带编号的可验收条目 |
| 方案设计 | [docs/design/](docs/design/) | 一页架构 + 调度策略 |
| 代码实现 | `src/` `cmd/` `web/` | 一个选定场景的运维闭环及失败分支能走通 |
| 测试报告 | [docs/test/](docs/test/) + `tests/` | 对照实验，引用需求编号 |
| 部署上线 | [docs/deploy/](docs/deploy/) + `deploy/` | 单机可复现，不是公网 SaaS |
| 比赛交付 | [docs/competition/](docs/competition/) | 一页纸 + 三问 |

范围锁： [docs/problem.md](docs/problem.md) · [docs/mvp.md](docs/mvp.md) · [docs/qianchao-cup.md](docs/qianchao-cup.md) · [AGENTS.md](AGENTS.md)

## 一句话

> 用卖家已有的网络资源，发现影响经营的问题，完成受控处理，并验证业务是否恢复。

## 不是什么

- 不是自建海缆 / IPLC / 再卖一条专线
- 不是「彻底解决跨境延迟」
- 不是「保证 Amazon / TikTok 不封号」
- 不是再套一层独立站 CDN

## 当前状态

正式市场调研报告：[docs/research/REPORT.md](docs/research/REPORT.md)。首版已锁定为“后台页面异常，先不切换店铺出口”，需求见 [docs/requirements/](docs/requirements/README.md)，设计见 [docs/design/](docs/design/README.md)。当前已有模拟状态机和测试，真实适配器及业务恢复仍未验证。

**9/8 官方复核：报名和材料均在 9/16 北京时间 24:00 截止。** 用户最新确认调整为 9/11 报名，先核实审核时效。第一天成果与待办见 [核查报告](docs/progress/2026-09-08.md)。

当前执行日程与每日 9 点复盘规则见 [docs/PLAN.md](docs/PLAN.md)：按任务推进调研、需求、原型、试用和比赛交付。

两条目标见 [docs/GOALS.md](docs/GOALS.md)。报名表官方截止 **2026-09-16**，提交后不可改。**计划 9 月 11 日核对后再交。** 官方快照 [docs/competition/official.md](docs/competition/official.md)，填写稿 [docs/competition/registration.md](docs/competition/registration.md)。正赛 BP/DEMO 过审后再传。

调研公开情报已进 `docs/research/`。一手核验未做。杯赛材料仍先按 `how-to-verify.md` 做 A/B，能访则访 C。
