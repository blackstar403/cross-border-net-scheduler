# cross-border-net-scheduler

钱潮杯 · AI+跨境出海。方向：**电商出海的网络智能调度**。

中小卖家不是没有网，是同一条不够好的跨境链路上，直播、广告、后台、ERP、多店铺互相踩。现有服务商卖管子（加速器 / SD-WAN / 云 GA / 专线），我们做控制面：按店铺、按场次，把当时最好的那一跳分给最值钱的流量。

## 什么叫做完

完整项目不是「代码能跑」，而是七段能互相引用。定义见 [docs/PROJECT.md](docs/PROJECT.md)。

| 模块 | 目录 | 钱潮杯最低完成 |
|---|---|---|
| 市场调研 | [docs/research/](docs/research/) | 客群、竞品分层、痛点三分、证据台账 |
| 需求分析 | [docs/requirements/](docs/requirements/) | 带编号的可验收条目 |
| 方案设计 | [docs/design/](docs/design/) | 一页架构 + 调度策略 |
| 代码实现 | `src/` `cmd/` `web/` | 演示脚本三步能走通 |
| 测试报告 | [docs/test/](docs/test/) + `tests/` | 对照实验，引用需求编号 |
| 部署上线 | [docs/deploy/](docs/deploy/) + `deploy/` | 单机可复现，不是公网 SaaS |
| 比赛交付 | [docs/competition/](docs/competition/) | 一页纸 + 三问 |

范围锁： [docs/problem.md](docs/problem.md) · [docs/mvp.md](docs/mvp.md) · [docs/qianchao-cup.md](docs/qianchao-cup.md) · [AGENTS.md](AGENTS.md)

## 一句话

> 在卖家已经买得起的那几条烂管子上，把稀缺的「还算能用的路径」分给此刻最值钱的流量，并让结果可看见。

## 不是什么

- 不是自建海缆 / IPLC / 再卖一条专线
- 不是「彻底解决跨境延迟」
- 不是「保证 Amazon / TikTok 不封号」
- 不是再套一层独立站 CDN

## 当前状态

调研公开情报已进 `docs/research/`（证据台账）。**一手核验未做**，痛点仍是假设。先按 `docs/research/how-to-verify.md` 做晚高峰测量、同出口检查、3 个卖家访谈，再拆需求。
