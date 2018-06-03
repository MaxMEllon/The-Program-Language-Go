|項目|意味|
|---|---|
|real|プログラムの呼び出しから終了までにかかった実時間（秒）|
|user|プログラム自体の処理時間（秒）（ユーザCPU時間）|
|sys |プログラムを処理するために、OSが処理をした時間（秒）（システム時間）|

|引数の数|bad|good|
|---|---|---|
|5000|real 0.37<br />user 0.27<br />sys 0.20<br /> |real 0.32<br /> user 0.25<br /> sys 0.18<br />|
|10000|real 0.39<br /> user 0.30<br /> sys 0.21<br />| real 0.33<br /> user 0.27<br /> sys 0.18<br />|
|15000|real 0.42<br /> user 0.36<br /> sys 0.20<br />| real 0.35<br /> user 0.29<br /> sys 0.19<br />|
|20000| real 0.46<br /> user 0.45<br /> sys 0.20<br />| real 0.36<br /> user 0.29<br /> sys 0.18<br />|
|25000| real 0.57<br /> user 0.53<br /> sys 0.22<br />| real 0.36<br /> user 0.30<br /> sys 0.19<br />|

![](./graph.png)
