# 参与指南
- [x]  fio测试
- [x]  机器配置查询
- [x]  tiup部署起集群
- [x]  sysbench数据prepare
- [ ]  学习tiup如何替换binary
- [ ]  替换集群binary
    - [ ]  指派编译工作
    - [ ]  编译打包和上传
    - [ ]  执行patch
- [ ]  修改集群拓扑扩容
    - [ ]  通过测试的反馈来设置一个合理的集群拓扑
    - [ ]  尽可能的利用系统资源
- [ ]  文档跟进与更新
    - [ ]  如何使用sysbench
    - [ ]  有何种工作负载
    - [ ]  如何得出理论预期
    - [ ]  如何设计测试
    - [ ]  如何获取测试结果
    - [ ]  如何分析测试结果
    - [ ]  如何共享测试结果
- [ ]  讨论与确立issues
    - [ ]  提交issue
    - [ ]  认领issue
    - [ ]  提交PR
    - [ ]  pre-review PR

[TOC]

## 磁盘IO速度测试

### fio

参看 [fio.md](./fio.md)

## 集群拓扑

|IP                 |角色           |vCPU       |内存       |磁盘       |
|---                |---            |---        |---        |---        |
|10.13.132.27       |TiDB           |4          |8GB        |200GB      |
|10.13.178.195      |TiKV           |4          |8GB        |200GB      |
|10.13.88.248       |PD             |4          |8GB        |200GB      |
|10.13.88.248       |Prometheus     |4          |8GB        |200GB      |
|10.13.88.248       |TiDB Dashboard |4          |8GB        |200GB      |


详细配置参看 [topo.yaml](./topo.yaml)

## TiUp

### 集群启动与停止

```bash
# 启动
tiup cluster start <cluster-name>
# 停止
tiup cluster stop <cluster-name>
```

### 集群拓扑更新

```bash
# 更新配置
tiup cluster edit-config <cluster-name>
# 重载配置
tiup cluster reload <cluster-name>

# 扩容
tiup cluster scale-out <cluster-name> <topology.yaml>
```

### 集群组件版本替换

```bash
# 替换TiKV结点的服务器二进制可执行文件，并在未来的扩容中保持一致
tiup cluster patch <cluster-name> <package-path> -R tikv --overwrite
```

## Sysbench

### Insert SQLs

```bash
# TODO
```

### 数据库

```mysql
-- Data Schema
CREATE TABLE `sbtest1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `k` int(11) NOT NULL DEFAULT 0,
  `c` char(120) NOT NULL DEFAULT '',
  `pad` char(60) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `k_1` (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin

-- Count 32 * 100k
use sb_oltp_insert_test;
show tables;
```

### 压测指令

```bash
# TODO
sysbench oltp_insert help
sysbench --config-file=sysbench.conf oltp_insert run
```

### 获取测试结果

```bash
# TODO
```

### 分析测试结果

```bash
# TODO
```

### 共享分析结果

```bash
# TODO
```
