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

- 顺序写极限大约130MBps(通过dd指令测试观测得到)
- 随机写极限大约48MBps(通过fio指令测试观测得到)

详细参看 [fio.md](./fio.md)

## 集群拓扑

|IP                 |角色            |vCPU       |内存       |磁盘       |端口|
|---                |---            |---        |---        |---        |---|
|10.13.132.27       |TiDB           |4          |8GB        |200GB      |4000|
|10.13.178.195      |TiKV           |4          |8GB        |200GB      |20160|
|10.13.88.248       |PD             |4          |8GB        |200GB      |2380|
|10.13.88.248       |Prometheus     |4          |8GB        |200GB      |9090|
|10.13.88.248       |Granfana       |4          |8GB        |200GB      |3000|
|10.13.88.248       |TiDB Dashboard |4          |8GB        |200GB      |2379|

### TiDB/MySQL Proxy

```bash
# 106.75.175.149是外网IP，这意味着可以在外网机器上做
mysql -h 106.75.175.149 -u root -P 4000 Sysbench测试
# `sb_oltp_insert_test`库目前有32 * 100k的数据记录(sysbench oltp_common)
show databases;
```

### TiDB Dashboard

可以利用TiDB Dashboard进行慢SQL的查询等操作

- http://106.75.174.70:2379/dashboard/

### Granfana地址: 

可以利用Granfana观测集群结点的各项监控指标

- http://106.75.174.70:3000 
- usr:admin pwd:admin

详细集群拓扑配置参看 [topo.yaml](./topo.yaml)

## 手动编译

### tips

- 106.75.174.70已经安装了所有的编译工具
- 106.75.174.70:/opt 目录下有tidb、tikv、pd项目的源代码
- 编译tikv的时候用make dev
- 编译完成得到的二进制可执行文件记得打包，可以用git commit log的hash前7位命名

### 指令

```bash
# 编译TiDB，结果文件在/opt/tidb/bin下
make /opt/tidb
# 编译TiKV，结果文件在/opt/tikv/target/release下
make dev /opt/tikv
# 编译PD，结果文件在/opt/pd/bin下
make /opt/pd
```

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

### Insert SQL模板

```mysql
# TODO
INSERT INTO sbtest14 (id, k, c, pad) VALUES (...)
```

### 数据库/表模式

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
-- 一共有(sbtest1,sbtest2...,sbtest32)张表，且数据表模式一致
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

#### 可以利用的工具有

- [TiDB Dashboard](http://106.75.174.70:2379/dashboard/)
- [Granfana](http://106.75.174.70:3000 )
- [TiDB Trace](https://docs.pingcap.com/zh/tidb/stable/sql-statement-trace)
- CPU profile

### 如何获取测试结果

#### Granfana截图

#### TiDB Tace查询

```sql
trace format="row" INSERT INTO
  sbtest19 (id, k, c, pad)
VALUES
  (
    0,
    5004,
    '41783842345-83708116042-69812290394-97528821899-23702363349-99445178113-25538720927-56652535101-32509923548-28575012141',
    '98852473800-29909224398-15665687870-23077232324-01373855493'
  );
```

### 如何分析测试结果

TODOs

### 如何共享分析结果

TODOs

## 开发规范 (WIP)

1. fork repos
2. check new branch
3. modify
4. push || pull request

---

# 竞赛计划 (WIP)

## 策略

- 性能优化比例高的issue和PR（少）
- 性能优化比例低的issue和PR（多）

## 待办列表

- TiDB trace
- 根据trace内容阅读源码，找到开销较大的模块进行CPU profile
- TODOs

## 任务板

TODOs
