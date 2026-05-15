# UOS Mgmt Exporter for Prometheus

基于统信UOS操作系统的管理资源监控导出器，用于收集和导出系统管理相关的Prometheus监控指标。

## 项目简介

UOS Mgmt Exporter 是一个专门为统信UOS操作系统开发的管理资源监控工具。它收集系统管理相关的指标数据并以Prometheus格式导出，包括资源状态、检查应用操作、失败统计等关键监控指标。

## 功能特性

- 🚀 **全面的管理指标收集**: 支持资源状态监控、检查应用操作统计、失败资源追踪
- 📊 **Prometheus 兼容**: 原生支持Prometheus监控体系，无缝集成
- 🎯 **UOS 优化**: 专为统信UOS操作系统环境优化
- ⚡ **高性能**: 低资源占用，高效稳定运行
- 🔧 **灵活配置**: 支持YAML配置文件和命令行参数
- 🛡️ **企业级特性**: 内置限流保护、健康检查、完善的日志管理

## 开发指南

### 项目结构
```
uos-mgmt-exporter/
├── main.go                    # 程序入口
├── exporter.go               # 主程序逻辑
├── config/                   # 配置文件
│   ├── config.go            # 配置解析
│   └── mgmt-exporter.yaml   # 默认配置
├── internal/                 # 内部实现
│   ├── exporter/            # 导出器核心
│   ├── metrics/              # 指标收集
│   └── server/               # HTTP服务器
├── Makefile                 # 构建脚本
└── uos-mgmt-exporter.service # Systemd服务文件
```

### 添加新指标

1. 在 `internal/metrics/` 目录下修改指标收集器
2. 实现 `exporter.Metric` 接口
3. 在初始化函数中注册新指标
4. 更新指标文档

## 参与贡献

欢迎提交 Issue 和 Pull Request 来改进项目。

### 提交规范

- 使用清晰的提交信息
- 添加适当的测试用例
- 更新相关文档
- 遵循 Go 代码规范

## 安装

### 从源码编译

```bash
git clone https://atomgit.com/deepin-community/uos-mgmt-exporter.git
cd uos-mgmt-exporter
go build
```

### 二进制安装

从 [发布页面](https://atomgit.com/deepin-community/uos-mgmt-exporter/releases) 下载适用于您系统的最新二进制文件。

## 使用方法

### 基本使用

```bash
./uos-mgmt-exporter 
```

### YAML 配置文件

也可以使用 YAML 配置文件：

```yaml

address: "0.0.0.0"
port: 9098
metricsPath: "/metrics"
log:
  level: "debug"
log_path: "/var/log/uos-exporter/mgmt-exporter.log"
```

## 监控指标

### 资源管理指标

| 指标名称 | 类型 | 标签 | 说明 |
|---------|------|------|------|
| `mgmt_resources` | Gauge | kind | 当前管理的资源数量，按资源类型分类 |
| `mgmt_checkapply_total` | Counter | kind, apply, eventful, errorful | 检查应用操作的总次数 |
| `mgmt_failures_total` | Counter | kind, failure | 失败资源的累计总数 |
| `mgmt_failures` | Gauge | kind, failure | 当前失败资源的数量 |
| `mgmt_graph_start_time_seconds` | Gauge | 无 | 图表启动的时间戳 |

### 资源状态分类

- **Soft Fail**: 软失败，可能恢复的错误状态
- **Hard Fail**: 硬失败，需要人工干预的严重错误
- **OK**: 正常状态

### 指标示例

```prometheus
# HELP mgmt_resources Number of managed resources.
# TYPE mgmt_resources gauge
mgmt_resources{kind="Pod"} 5
mgmt_resources{kind="Deployment"} 3

# HELP mgmt_failures Number of failing resources.
# TYPE mgmt_failures gauge
mgmt_failures{failure="hard",kind="Pod"} 1
mgmt_failures{failure="soft",kind="Deployment"} 2
```

## Prometheus 配置

在您的 `prometheus.yaml` 中添加以下配置：

```yaml
scrape_configs:
  - job_name: "uos-mgmt-exporter"
    static_configs:
      - targets: ["localhost:9098"]
```

## 系统服务管理

### 使用 systemd 管理

```bash
# 启动服务
sudo systemctl start uos-mgmt-exporter

# 设置开机自启
sudo systemctl enable uos-mgmt-exporter

# 查看服务状态
sudo systemctl status uos-mgmt-exporter

# 停止服务
sudo systemctl stop uos-mgmt-exporter

# 重启服务
sudo systemctl restart uos-mgmt-exporter
```

### 日志查看

```bash
# 查看服务日志
sudo journalctl -u uos-mgmt-exporter -f

# 查看应用日志
tail -f /var/log/uos-exporter/mgmt-exporter.log
```

## 参与贡献

1. Fork 本仓库
2. 新建 feature 分支 (`git checkout -b feature/AmazingFeature`)
3. 提交修改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## 许可证

// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
