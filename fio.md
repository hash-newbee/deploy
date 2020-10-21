```bash
fio -filename=/dev/vdb -direct=1 -iodepth 1 -thread -rw=randwrite -ioengine=psync -bs=4k -size=1G -numjobs=64 -runtime=10 -group_reporting -name=filels
filels: (g=0): rw=randwrite, bs=(R) 4096B-4096B, (W) 4096B-4096B, (T) 4096B-4096B, ioengine=psync, iodepth=1
...
fio-3.7
Starting 64 threads
Jobs: 64 (f=64): [w(64)][100.0%][r=0KiB/s,w=46.1MiB/s][r=0,w=11.8k IOPS][eta 00m:00s]
filels: (groupid=0, jobs=64): err= 0: pid=9195: Fri Oct 16 12:08:10 2020
  write: IOPS=11.8k, BW=46.1MiB/s (48.3MB/s)(461MiB/10006msec)
    clat (usec): min=102, max=9153, avg=5419.74, stdev=507.91
     lat (usec): min=102, max=9153, avg=5420.35, stdev=507.92
    clat percentiles (usec):
     |  1.00th=[ 4883],  5.00th=[ 4948], 10.00th=[ 5014], 20.00th=[ 5014],
     | 30.00th=[ 5014], 40.00th=[ 5014], 50.00th=[ 5014], 60.00th=[ 5932],
     | 70.00th=[ 5997], 80.00th=[ 5997], 90.00th=[ 5997], 95.00th=[ 5997],
     | 99.00th=[ 6063], 99.50th=[ 6128], 99.90th=[ 6652], 99.95th=[ 7046],
     | 99.99th=[ 8029]
   bw (  KiB/s): min=  656, max=  800, per=1.56%, avg=737.46, stdev=27.18, samples=1265
   iops        : min=  164, max=  200, avg=184.33, stdev= 6.80, samples=1265
  lat (usec)   : 250=0.01%, 500=0.01%, 1000=0.01%
  lat (msec)   : 2=0.02%, 4=0.16%, 10=99.81%
  cpu          : usr=0.03%, sys=0.05%, ctx=118084, majf=0, minf=0
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=0,118073,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
  WRITE: bw=46.1MiB/s (48.3MB/s), 46.1MiB/s-46.1MiB/s (48.3MB/s-48.3MB/s), io=461MiB (484MB), run=10006-10006msec

Disk stats (read/write):
  vdb: ios=42/116714, merge=0/1, ticks=8/627327, in_queue=518306, util=99.08%
```

```bash
fio -filename=/dev/vdb-direct=1 -iodepth 1 -thread -rw=randread -ioengine=psync -bs=4k -size=1G -numjobs=64 -runtime=10 -group_reporting -name=file
file: (g=0): rw=randread, bs=(R) 4096B-4096B, (W) 4096B-4096B, (T) 4096B-4096B, ioengine=psync, iodepth=1
...
fio-3.7
Starting 64 threads
file: Laying out IO file (1 file / 1024MiB)
Jobs: 47 (f=47): [r(1),_(1),r(4),_(2),r(1),_(1),r(1),_(1),r(12),_(5),r(1),_(1),r(1),_(3),r(1),_(2),r(6),_(1),r(19)][81.8%][r=7055MiB/s,w=0KiB/s][r=1806k,w=0 IOPS][eta 00m:02s]
file: (groupid=0, jobs=64): err= 0: pid=37018: Fri Oct 16 12:12:25 2020
   read: IOPS=1784k, BW=6970MiB/s (7308MB/s)(64.0GiB/9403msec)
    clat (nsec): min=520, max=569801k, avg=19121.35, stdev=1854345.65
     lat (nsec): min=700, max=569801k, avg=22013.85, stdev=1991530.44
    clat percentiles (nsec):
     |  1.00th=[     764],  5.00th=[     812], 10.00th=[     948],
     | 20.00th=[    1048], 30.00th=[    1064], 40.00th=[    1144],
     | 50.00th=[    1272], 60.00th=[    1304], 70.00th=[    1336],
     | 80.00th=[    1448], 90.00th=[    1576], 95.00th=[    1656],
     | 99.00th=[    1928], 99.50th=[    2096], 99.90th=[    9024],
     | 99.95th=[   10560], 99.99th=[54788096]
   bw (  KiB/s): min=35679, max=280151, per=1.60%, avg=113882.70, stdev=27965.13, samples=1083
   iops        : min= 8919, max=70037, avg=28470.40, stdev=6991.27, samples=1083
  lat (nsec)   : 750=0.49%, 1000=11.36%
  lat (usec)   : 2=87.43%, 4=0.36%, 10=0.30%, 20=0.05%, 50=0.01%
  lat (usec)   : 100=0.01%, 250=0.01%, 500=0.01%
  lat (msec)   : 2=0.01%, 4=0.01%, 10=0.01%, 20=0.01%, 50=0.01%
  lat (msec)   : 100=0.01%, 250=0.01%, 500=0.01%, 750=0.01%
  cpu          : usr=1.77%, sys=4.90%, ctx=3699, majf=0, minf=0
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=16777216,0,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
   READ: bw=6970MiB/s (7308MB/s), 6970MiB/s-6970MiB/s (7308MB/s-7308MB/s), io=64.0GiB (68.7GB), run=9403-9403msec
```

```bash
fio -filename=/dev/vdb-direct=1 -iodepth 1 -thread -rw=read -ioengine=psync -bs=4k -size=1G -numjobs=64 -runtime=10 -group_reporting -name=file
file: (g=0): rw=read, bs=(R) 4096B-4096B, (W) 4096B-4096B, (T) 4096B-4096B, ioengine=psync, iodepth=1
...
fio-3.7
Starting 64 threads
Jobs: 48 (f=48): [R(10),_(1),R(15),_(2),R(1),_(2),R(3),_(1),R(1),_(3),R(1),_(2),R(4),_(1),R(1),_(1),R(1),_(1),R(5),_(2),R(6)][88.9%][r=8825MiB/s,w=0KiB/s][r=2259k,w=0 IOPS][eta 00m:01s]
file: (groupid=0, jobs=64): err= 0: pid=37085: Fri Oct 16 12:14:59 2020
   read: IOPS=2208k, BW=8625MiB/s (9044MB/s)(64.0GiB/7598msec)
    clat (nsec): min=500, max=390089k, avg=13768.37, stdev=1571686.57
     lat (nsec): min=680, max=390089k, avg=16489.92, stdev=1723796.91
    clat percentiles (nsec):
     |  1.00th=[   652],  5.00th=[   660], 10.00th=[   668], 20.00th=[   692],
     | 30.00th=[   724], 40.00th=[   772], 50.00th=[   868], 60.00th=[   932],
     | 70.00th=[  1012], 80.00th=[  1080], 90.00th=[  1208], 95.00th=[  1288],
     | 99.00th=[  1496], 99.50th=[  1624], 99.90th=[  8640], 99.95th=[ 12480],
     | 99.99th=[187392]
   bw (  KiB/s): min=62674, max=281808, per=1.59%, avg=140627.64, stdev=34686.91, samples=855
   iops        : min=15668, max=70452, avg=35156.63, stdev=8671.71, samples=855
  lat (nsec)   : 750=34.59%, 1000=32.90%
  lat (usec)   : 2=32.27%, 4=0.05%, 10=0.12%, 20=0.06%, 50=0.01%
  lat (usec)   : 100=0.01%, 250=0.01%
  lat (msec)   : 2=0.01%, 4=0.01%, 10=0.01%, 20=0.01%, 50=0.01%
  lat (msec)   : 100=0.01%, 250=0.01%, 500=0.01%
  cpu          : usr=1.85%, sys=4.84%, ctx=3033, majf=0, minf=0
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=16777216,0,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
   READ: bw=8625MiB/s (9044MB/s), 8625MiB/s-8625MiB/s (9044MB/s-9044MB/s), io=64.0GiB (68.7GB), run=7598-7598msec
```

```bash
fio -filename=/dev/vdb -direct=1 -iodepth 1 -thread -rw=write -ioengine=psync -bs=4k -size=1G -numjobs=64 -runtime=10 -group_reporting -name=file
file: (g=0): rw=write, bs=(R) 4096B-4096B, (W) 4096B-4096B, (T) 4096B-4096B, ioengine=psync, iodepth=1
...
fio-3.7
Starting 64 threads
Jobs: 64 (f=64): [W(64)][100.0%][r=0KiB/s,w=46.1MiB/s][r=0,w=11.8k IOPS][eta 00m:00s]
file: (groupid=0, jobs=64): err= 0: pid=37151: Fri Oct 16 12:15:53 2020
  write: IOPS=11.8k, BW=46.1MiB/s (48.3MB/s)(461MiB/10006msec)
    clat (usec): min=118, max=7293, avg=5419.48, stdev=491.81
     lat (usec): min=119, max=7293, avg=5420.07, stdev=491.81
    clat percentiles (usec):
     |  1.00th=[ 4948],  5.00th=[ 5014], 10.00th=[ 5014], 20.00th=[ 5014],
     | 30.00th=[ 5014], 40.00th=[ 5014], 50.00th=[ 5014], 60.00th=[ 5932],
     | 70.00th=[ 5997], 80.00th=[ 5997], 90.00th=[ 5997], 95.00th=[ 5997],
     | 99.00th=[ 6063], 99.50th=[ 6063], 99.90th=[ 6128], 99.95th=[ 6194],
     | 99.99th=[ 6587]
   bw (  KiB/s): min=  680, max=  800, per=1.56%, avg=737.53, stdev=23.35, samples=1280
   iops        : min=  170, max=  200, avg=184.35, stdev= 5.84, samples=1280
  lat (usec)   : 250=0.01%, 500=0.01%, 1000=0.01%
  lat (msec)   : 2=0.01%, 4=0.07%, 10=99.90%
  cpu          : usr=0.02%, sys=0.09%, ctx=118096, majf=0, minf=0
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued rwts: total=0,118089,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=1

Run status group 0 (all jobs):
  WRITE: bw=46.1MiB/s (48.3MB/s), 46.1MiB/s-46.1MiB/s (48.3MB/s-48.3MB/s), io=461MiB (484MB), run=10006-10006msec

Disk stats (read/write):
  vdb: ios=84/116738, merge=0/13, ticks=16/630633, in_queue=537147, util=99.11%
```
