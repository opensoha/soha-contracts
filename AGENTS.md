# Soha Contracts 仓库入口

- 本仓是公开 OpenAPI、JSON Schema、生成 SDK 和兼容性基线的真实源；先改源契约再运行生成入口，不手改 `gen/` 下的 DTO。
- 在 OpenSoha 多仓工作区中读取 `../AGENTS.md` 一次；独立克隆时使用本仓规则，不要求自动初始化相邻仓库或规划工具。
- 契约实现或审查按需使用 [soha-contracts](.agents/skills/soha-contracts/SKILL.md)。公开行为变更按依赖方向同步受影响 consumers。
- 主入口为 `npm test`；公开契约改动必须执行兼容性和完整 consumer 矩阵，使用 [CI](.github/workflows/ci.yml) 与包脚本中的真实参数。缺少 consumer 或发布依赖时报告验证缺口，不用本地替换掩盖缺失。
- 文档和技能改动只检查内容、链接与差异；相关代码和环境未变化时复用成功验证，保留用户未提交改动。
