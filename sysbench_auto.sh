#!/bin/bash
sysbench --config-file=sysbench.conf oltp_insert --tables=32 --threads=1 run;
sysbench --config-file=sysbench.conf oltp_insert --tables=32 --threads=4 run;
sysbench --config-file=sysbench.conf oltp_insert --tables=32 --threads=8 run;
sysbench --config-file=sysbench.conf oltp_insert --tables=32 --threads=16 run;
sysbench --config-file=sysbench.conf oltp_insert --tables=32 --threads=32 run;
sysbench --config-file=sysbench.conf oltp_insert --tables=32 --threads=64 run;
sysbench --config-file=sysbench.conf oltp_insert --tables=32 --threads=128 run;
sysbench --config-file=sysbench.conf oltp_insert --tables=32 --threads=256 run;