# UOS Mgmt Exporter for Prometheus

A Prometheus exporter for management resources monitoring on UnionTech UOS operating system.

## Description

UOS Mgmt Exporter is a management resource monitoring tool specifically developed for UnionTech UOS operating system. It collects system management-related metric data and exports it in Prometheus format, including resource status, check-apply operations, failure statistics and other key monitoring metrics.

## Features

- **Comprehensive Management Metrics**: Resource status monitoring, check-apply operation statistics, failure resource tracking
- **Prometheus Compatible**: Native support for Prometheus monitoring ecosystem
- **UOS Optimized**: Optimized for UnionTech UOS operating system environment
- **High Performance**: Low resource usage, efficient and stable operation
- **Flexible Configuration**: Support YAML configuration file and command-line arguments
- **Enterprise Features**: Built-in rate limiting, health check, comprehensive log management

## Installation

### Build from Source

```bash
git clone https://atomgit.com/deepin-community/uos-mgmt-exporter.git
cd uos-mgmt-exporter
go build
```

### Binary Installation

Download the latest binary for your system from the [Releases](https://atomgit.com/deepin-community/uos-mgmt-exporter/releases) page.

## Usage

### Basic Usage

```bash
./uos-mgmt-exporter
```

### YAML Configuration

```yaml
address: "0.0.0.0"
port: 9098
metricsPath: "/metrics"
log:
  level: "debug"
log_path: "/var/log/uos-exporter/mgmt-exporter.log"
```

## Monitoring Metrics

| Metric Name | Type | Labels | Description |
|-------------|------|--------|-------------|
| `mgmt_resources` | Gauge | kind | Number of managed resources, classified by resource type |
| `mgmt_checkapply_total` | Counter | kind, apply, eventful, errorful | Total number of check-apply operations |
| `mgmt_failures_total` | Counter | kind, failure | Cumulative total of failed resources |
| `mgmt_failures` | Gauge | kind, failure | Current number of failing resources |
| `mgmt_graph_start_time_seconds` | Gauge | none | Timestamp of graph start time |

## Prometheus Configuration

Add the following configuration to your `prometheus.yaml`:

```yaml
scrape_configs:
  - job_name: "uos-mgmt-exporter"
    static_configs:
      - targets: ["localhost:9098"]
```

## System Service Management

### Using systemd

```bash
# Start service
sudo systemctl start uos-mgmt-exporter

# Enable on boot
sudo systemctl enable uos-mgmt-exporter

# Check service status
sudo systemctl status uos-mgmt-exporter

# Stop service
sudo systemctl stop uos-mgmt-exporter

# Restart service
sudo systemctl restart uos-mgmt-exporter
```

### View Logs

```bash
# View service logs
sudo journalctl -u uos-mgmt-exporter -f

# View application logs
tail -f /var/log/uos-exporter/mgmt-exporter.log
```

## Contribution

1. Fork the repository
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Create Pull Request

## License

// SPDX-FileCopyrightText: 2025 UnionTech Software Technology Co., Ltd.
// SPDX-License-Identifier: MIT
