# 空闲
```
trace	13:56:21.027222	3.6906ms
  └─session.Execute	13:56:21.027227	3.675909ms
    ├─session.ParseSQL	13:56:21.027231	14.989µs
    └─session.ExecuteStmt	13:56:21.027266	3.632187ms
      ├─executor.Compile	13:56:21.027269	76.762µs
      └─session.runStmt	13:56:21.027354	3.528566ms
        ├─executor.handleNoDelayExecutor	13:56:21.027377	358.349µs
        │ └─*executor.InsertExec.Next	13:56:21.027380	341.458µs
        │   └─table.AllocBatchAutoIncrementValue	13:56:21.027393	3.949µs
        └─session.CommitTxn	13:56:21.027743	3.122685ms
          └─session.doCommitWitRetry	13:56:21.027745	3.115585ms
            └─tikvTxn.Commit	13:56:21.027750	3.094172ms
              ├─twoPhaseCommitter.prewriteMutations	13:56:21.027759	2.280513ms
              │ ├─regionReqauest.SendReqCtx	13:56:21.027791	1.939992ms
              │ │ └─rpcClient.SendRequest	13:56:21.027811	1.907223ms
              │ └─regionReqauest.SendReqCtx	13:56:21.027842	2.18759ms
              │   └─rpcClient.SendRequest	13:56:21.027850	2.171799ms
              ├─tikvStore.getTimestampWithRetry	13:56:21.030046	118.612µs
              └─twoPhaseCommitter.commitMutations	13:56:21.030175	658.047µs
                └─regionReqauest.SendReqCtx	13:56:21.030182	642.367µs
                  └─rpcClient.SendRequest	13:56:21.030186	627.836µs
```
---

# sysbench oltp_inert -threads=1 -tables=32
```
trace	17:28:48.101381	2.15208ms
  └─session.Execute	17:28:48.101386	2.137248ms
    ├─session.ParseSQL	17:28:48.101390	11.93µs
    └─session.ExecuteStmt	17:28:48.101423	2.096427ms
      ├─executor.Compile	17:28:48.101425	31.18µs
      └─session.runStmt	17:28:48.101462	2.035126ms
        ├─executor.handleNoDelayExecutor	17:28:48.101475	254.187µs
        │ └─*executor.InsertExec.Next	17:28:48.101478	237.697µs
        │   └─table.AllocBatchAutoIncrementValue	17:28:48.101483	1.41µs
        └─session.CommitTxn	17:28:48.101736	1.749108ms
          └─session.doCommitWitRetry	17:28:48.101748	1.732338ms
            └─tikvTxn.Commit	17:28:48.101751	1.715777ms
              ├─twoPhaseCommitter.prewriteMutations	17:28:48.101756	750.58µs
              │ ├─regionReqauest.SendReqCtx	17:28:48.101775	715.39µs
              │ │ └─rpcClient.SendRequest	17:28:48.101790	685.989µs
              │ └─regionReqauest.SendReqCtx	17:28:48.101786	695.409µs
              │   └─rpcClient.SendRequest	17:28:48.101798	666.878µs
              ├─tikvStore.getTimestampWithRetry	17:28:48.102513	279.777µs
              └─twoPhaseCommitter.commitMutations	17:28:48.102815	643.318µs
                └─regionReqauest.SendReqCtx	17:28:48.102823	621.567µs
                  └─rpcClient.SendRequest	17:28:48.102827	603.137µs
```
---

# sysbench oltp_inert -threads=4 -tables=32
```
trace	17:30:52.043047	3.056285ms
  └─session.Execute	17:30:52.043051	3.043964ms
    ├─session.ParseSQL	17:30:52.043054	14.032µs
    └─session.ExecuteStmt	17:30:52.043107	2.982653ms
      ├─executor.Compile	17:30:52.043111	30.49µs
      └─session.runStmt	17:30:52.043149	2.919712ms
        ├─executor.handleNoDelayExecutor	17:30:52.043164	170.094µs
        │ └─*executor.InsertExec.Next	17:30:52.043167	154.573µs
        │   └─table.AllocBatchAutoIncrementValue	17:30:52.043173	1.351µs
        └─session.CommitTxn	17:30:52.043342	2.715304ms
          └─session.doCommitWitRetry	17:30:52.043343	2.708686ms
            └─tikvTxn.Commit	17:30:52.043346	2.690486ms
              ├─twoPhaseCommitter.prewriteMutations	17:30:52.043350	1.439619ms
              │ ├─regionReqauest.SendReqCtx	17:30:52.043404	1.065201ms
              │ │ └─rpcClient.SendRequest	17:30:52.043417	1.042549ms
              │ └─regionReqauest.SendReqCtx	17:30:52.043431	1.347617ms
              │   └─rpcClient.SendRequest	17:30:52.043436	1.331457ms
              ├─tikvStore.getTimestampWithRetry	17:30:52.044795	252.887µs
              └─twoPhaseCommitter.commitMutations	17:30:52.045059	968.307µs
                └─regionReqauest.SendReqCtx	17:30:52.045064	953.946µs
                  └─rpcClient.SendRequest	17:30:52.045068	935.155µs
```
---

# sysbench oltp_inert -threads=8 -tables=32
```
trace	17:34:26.545181	4.276728ms
  └─session.Execute	17:34:26.545185	4.265268ms
    ├─session.ParseSQL	17:34:26.545188	23.691µs
    └─session.ExecuteStmt	17:34:26.545234	4.212576ms
      ├─executor.Compile	17:34:26.545237	43.181µs
      └─session.runStmt	17:34:26.545287	4.143864ms
        ├─executor.handleNoDelayExecutor	17:34:26.545302	172.815µs
        │ └─*executor.InsertExec.Next	17:34:26.545305	157.984µs
        │   └─table.AllocBatchAutoIncrementValue	17:34:26.545309	1.02µs
        └─session.CommitTxn	17:34:26.545480	3.942969ms
          └─session.doCommitWitRetry	17:34:26.545482	3.937489ms
            └─tikvTxn.Commit	17:34:26.545484	3.921979ms
              ├─twoPhaseCommitter.prewriteMutations	17:34:26.545489	1.852501ms
              │ ├─regionReqauest.SendReqCtx	17:34:26.545545	1.604294ms
              │ │ └─rpcClient.SendRequest	17:34:26.545560	1.576773ms
              │ └─regionReqauest.SendReqCtx	17:34:26.545552	1.779439ms
              │   └─rpcClient.SendRequest	17:34:26.545564	1.756048ms
              ├─tikvStore.getTimestampWithRetry	17:34:26.547348	292.708µs
              └─twoPhaseCommitter.commitMutations	17:34:26.547651	1.747708ms
                └─regionReqauest.SendReqCtx	17:34:26.547655	1.735718ms
                  └─rpcClient.SendRequest	17:34:26.547659	1.711828ms
```
---

# sysbench oltp_inert -threads=16 -tables=32
```
trace	17:39:47.177154	2.648282ms
  └─session.Execute	17:39:47.177161	2.630993ms
    ├─session.ParseSQL	17:39:47.177168	18.541µs
    └─session.ExecuteStmt	17:39:47.177229	2.55716ms
      ├─executor.Compile	17:39:47.177245	68.013µs
      └─session.runStmt	17:39:47.177323	2.428147ms
        ├─executor.handleNoDelayExecutor	17:39:47.177342	54.642µs
        │ └─*executor.InsertExec.Next	17:39:47.177346	36.232µs
        │   └─table.AllocBatchAutoIncrementValue	17:39:47.177353	1.251µs
        └─session.CommitTxn	17:39:47.177406	2.332764ms
          └─session.doCommitWitRetry	17:39:47.177408	2.324135ms
            └─tikvTxn.Commit	17:39:47.177411	2.295354ms
              ├─twoPhaseCommitter.prewriteMutations	17:39:47.177418	1.282816ms
              │ ├─regionReqauest.SendReqCtx	17:39:47.177440	1.174923ms
              │ │ └─rpcClient.SendRequest	17:39:47.177459	1.140712ms
              │ └─regionReqauest.SendReqCtx	17:39:47.177478	1.193404ms
              │   └─rpcClient.SendRequest	17:39:47.177486	1.150741ms
              ├─tikvStore.getTimestampWithRetry	17:39:47.178715	226.795µs
              └─twoPhaseCommitter.commitMutations	17:39:47.178967	726.371µs
                └─regionReqauest.SendReqCtx	17:39:47.178979	701.46µs
                  └─rpcClient.SendRequest	17:39:47.178988	672.978µs
```
---

# sysbench oltp_inert -threads=32 -tables=32
```
trace	17:43:47.700867	3.227497ms
  └─session.Execute	17:43:47.700873	3.207517ms
    ├─session.ParseSQL	17:43:47.700877	13.369µs
    └─session.ExecuteStmt	17:43:47.700914	3.157827ms
      ├─executor.Compile	17:43:47.700916	57.241µs
      └─session.runStmt	17:43:47.700981	3.054595ms
        ├─executor.handleNoDelayExecutor	17:43:47.700997	240.507µs
        │ └─*executor.InsertExec.Next	17:43:47.701000	221.395µs
        │   └─table.AllocBatchAutoIncrementValue	17:43:47.701005	1.46µs
        └─session.CommitTxn	17:43:47.701246	2.769796ms
          └─session.doCommitWitRetry	17:43:47.701248	2.760256ms
            └─tikvTxn.Commit	17:43:47.701251	2.729575ms
              ├─twoPhaseCommitter.prewriteMutations	17:43:47.701257	1.160343ms
              │ ├─regionReqauest.SendReqCtx	17:43:47.701311	980.707µs
              │ │ └─rpcClient.SendRequest	17:43:47.701326	948.546µs
              │ └─regionReqauest.SendReqCtx	17:43:47.701342	1.052789ms
              │   └─rpcClient.SendRequest	17:43:47.701348	996.987µs
              ├─tikvStore.getTimestampWithRetry	17:43:47.702431	429.931µs
              └─twoPhaseCommitter.commitMutations	17:43:47.702880	1.08409ms
                └─regionReqauest.SendReqCtx	17:43:47.702890	1.03726ms
                  └─rpcClient.SendRequest	17:43:47.702895	1.007739ms
```
---


# sysbench oltp_inert -threads=64 -tables=32
```
trace	17:49:50.793230	2.600152ms
  └─session.Execute	17:49:50.793235	2.588081ms
    ├─session.ParseSQL	17:49:50.793240	12.81µs
    └─session.ExecuteStmt	17:49:50.793274	2.545519ms
      ├─executor.Compile	17:49:50.793277	33.941µs
      └─session.runStmt	17:49:50.793318	2.483258ms
        ├─executor.handleNoDelayExecutor	17:49:50.793332	385.821µs
        │ └─*executor.InsertExec.Next	17:49:50.793341	361.6µs
        │   └─table.AllocBatchAutoIncrementValue	17:49:50.793346	1.99µs
        └─session.CommitTxn	17:49:50.793724	2.068396ms
          └─session.doCommitWitRetry	17:49:50.793726	2.062486ms
            └─tikvTxn.Commit	17:49:50.793729	2.045376ms
              ├─twoPhaseCommitter.prewriteMutations	17:49:50.793733	1.07856ms
              │ ├─regionReqauest.SendReqCtx	17:49:50.793800	672.638µs
              │ │ └─rpcClient.SendRequest	17:49:50.793815	643.668µs
              │ └─regionReqauest.SendReqCtx	17:49:50.793829	972.676µs
              │   └─rpcClient.SendRequest	17:49:50.793835	956.316µs
              ├─tikvStore.getTimestampWithRetry	17:49:50.794818	169.284µs
              └─twoPhaseCommitter.commitMutations	17:49:50.794999	765.911µs
                └─regionReqauest.SendReqCtx	17:49:50.795006	747.61µs
                  └─rpcClient.SendRequest	17:49:50.795010	729.15µs
```
---

# sysbench oltp_inert -threads=128 -tables=32
```
trace	17:58:28.797118	1.891721ms
  └─session.Execute	17:58:28.797122	1.880551ms
    ├─session.ParseSQL	17:58:28.797125	13.01µs
    └─session.ExecuteStmt	17:58:28.797157	1.84081ms
      ├─executor.Compile	17:58:28.797160	42.831µs
      └─session.runStmt	17:58:28.797209	1.772379ms
        ├─executor.handleNoDelayExecutor	17:58:28.797222	191.466µs
        │ └─*executor.InsertExec.Next	17:58:28.797225	173.794µs
        │   └─table.AllocBatchAutoIncrementValue	17:58:28.797230	1.18µs
        └─session.CommitTxn	17:58:28.797420	1.551052ms
          └─session.doCommitWitRetry	17:58:28.797422	1.544642ms
            └─tikvTxn.Commit	17:58:28.797424	1.527822ms
              ├─twoPhaseCommitter.prewriteMutations	17:58:28.797428	710.56µs
              │ ├─regionReqauest.SendReqCtx	17:58:28.797491	634.087µs
              │ │ └─rpcClient.SendRequest	17:58:28.797504	610.917µs
              │ └─regionReqauest.SendReqCtx	17:58:28.797510	614.257µs
              │   └─rpcClient.SendRequest	17:58:28.797516	591.117µs
              ├─tikvStore.getTimestampWithRetry	17:58:28.798145	132.943µs
              └─twoPhaseCommitter.commitMutations	17:58:28.798291	644.747µs
                └─regionReqauest.SendReqCtx	17:58:28.798300	623.827µs
                  └─rpcClient.SendRequest	17:58:28.798305	606.707µs
```
---

# sysbench oltp_inert -threads=256 -tables=32
```
trace	17:58:28.797118	1.891721ms
  └─session.Execute	17:58:28.797122	1.880551ms
    ├─session.ParseSQL	17:58:28.797125	13.01µs
    └─session.ExecuteStmt	17:58:28.797157	1.84081ms
      ├─executor.Compile	17:58:28.797160	42.831µs
      └─session.runStmt	17:58:28.797209	1.772379ms
        ├─executor.handleNoDelayExecutor	17:58:28.797222	191.466µs
        │ └─*executor.InsertExec.Next	17:58:28.797225	173.794µs
        │   └─table.AllocBatchAutoIncrementValue	17:58:28.797230	1.18µs
        └─session.CommitTxn	17:58:28.797420	1.551052ms
          └─session.doCommitWitRetry	17:58:28.797422	1.544642ms
            └─tikvTxn.Commit	17:58:28.797424	1.527822ms
              ├─twoPhaseCommitter.prewriteMutations	17:58:28.797428	710.56µs
              │ ├─regionReqauest.SendReqCtx	17:58:28.797491	634.087µs
              │ │ └─rpcClient.SendRequest	17:58:28.797504	610.917µs
              │ └─regionReqauest.SendReqCtx	17:58:28.797510	614.257µs
              │   └─rpcClient.SendRequest	17:58:28.797516	591.117µs
              ├─tikvStore.getTimestampWithRetry	17:58:28.798145	132.943µs
              └─twoPhaseCommitter.commitMutations	17:58:28.798291	644.747µs
                └─regionReqauest.SendReqCtx	17:58:28.798300	623.827µs
                  └─rpcClient.SendRequest	17:58:28.798305	606.707µs
```
---
