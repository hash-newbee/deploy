#!/bin/bash
sysbench --config-file=sysbench.conf oltp_insert --tables=32 --threads=128 run;
