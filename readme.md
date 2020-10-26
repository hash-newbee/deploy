# 参与指南

- [x]  fio测试
- [x]  机器配置查询
- [x]  tiup部署起集群
- [x]  sysbench数据prepare
- [x]  学习tiup如何替换binary
- [x]  替换集群binary
    - [x]  指派编译工作
    - [x]  编译打包和上传
    - [x]  执行patch
- [ ]  修改集群拓扑扩容
    - [x] 如何修改集群拓扑
- [ ]  文档跟进与更新
    - [x]  如何使用sysbench
    - [x]  有何种工作负载
    - [ ]  如何设计测试
    - [ ]  如何获取测试结果
    - [ ]  如何分析测试结果
    - [ ]  如何共享测试结果
- [ ]  讨论与确立issues
    - [ ]  提交issue
    - [ ]  认领issue
    - [ ]  提交PR
    - [ ]  pre-review PR

---

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

## 手动编译

### tips

- 106.75.174.70已经安装了所有的编译工具
- 106.75.174.70:/opt 目录下有tidb、tikv、pd项目的源代码
- 编译tikv的时候用make dev
- 编译完成得到的二进制可执行文件记得打包，可以用git commit log的hash前7位命名？

## TiUp

### tips

- 106.75.174.70下已经安装了tiup及部分组件

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

```mysql
# TODO
INSERT INTO sbtest14 (id, k, c, pad) VALUES (...)
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
sysbench oltp_insert help
# OLTP Insert
# tables固定是32
# table-size可以不设置
# threads可以设置为1 4 8 32 64 128 256来观测TPS的稳定性
sysbench --config-file=sysbench.conf --threads=N --tables=32 --table-size=S oltp_insert run
```

### 如何设计测试

#### 整体思路是

0. TiDB执行SQL可以被拆分成很多个进程
![TiDB执行框架流程图](https://download.pingcap.com/images/blog-cn/tidb-source-code-reading-2/2.png)
1. 观测流水线上的进程
2. 找到延迟较长的进程
3. 理论分析
4. 确定优化进程
5. 利用CPU profile对某些进程模块进行具体观测

#### 有帮助的知识有

- [TiDB 源码阅读系列文章（二）初识 TiDB 源码](https://pingcap.com/blog-cn/tidb-source-code-reading-2/)
- [TiDB 源码阅读系列文章（三）SQL 的一生](https://pingcap.com/blog-cn/tidb-source-code-reading-3/)
- [TiDB 源码阅读系列文章（四）Insert 语句概览](https://pingcap.com/blog-cn/tidb-source-code-reading-4/)
- [TiDB 源码阅读系列文章（十六）INSERT 语句详解](https://pingcap.com/blog-cn/tidb-source-code-reading-16/)
- [TiDB 源码阅读系列文章（十）Chunk 和执行框架简介](https://pingcap.com/blog-cn/tidb-source-code-reading-10/)
- [TiDB 源码阅读系列文章（十八）tikv-client（上）](https://pingcap.com/blog-cn/tidb-source-code-reading-18/)
- [TiDB 源码阅读系列文章（十九）tikv-client（下）](https://pingcap.com/blog-cn/tidb-source-code-reading-19/)
- [TiKV 源码解析系列文章（九）Service 层处理流程解析](https://pingcap.com/blog-cn/tikv-source-code-reading-9/)
- [TiKV 源码解析系列文章（十一）Storage - 事务控制层](https://pingcap.com/blog-cn/tikv-source-code-reading-11/)
- [TiKV 源码解析系列文章（十二）分布式事务](https://pingcap.com/blog-cn/tikv-source-code-reading-12/)

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

## 开发规范

TODOs

---

# 竞赛计划

## 待办列表

## 任务板

TODOs
