# Golang日志工具集

`logger`主要封装了主流的golang日志库，方便开箱即用。

- [logrus](logrus): https://github.com/sirupsen/logrus
- [zap](zap): https://github.com/uber-go/zap
- [glog](glog): https://github.com/golang/glog
- [klog](): https://github.com/kubernetes/klog

# 日志包的基本功能

### 日志级别

以下日志级别由低到高，高于指定级别的日志不输出。

- Panic：记录日志，然后panic。
- Fatal：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；
- Error：错误日志，需要查看原因；
- Warn：警告信息，提醒程序员注意；
- Info：关键操作，核心流程的日志；
- Debug：一般程序中输出的调试信息；

### 日志字段

基础字段

- time: 日志时间
- level: 日志级别
- msg: 日志信息
- file: 文件及行数
- func: 调用函数
- field: 自定义字段

HTTP请求字段

- path: 请求路径
- query: 请求参数
- method: 请求方法[GET,POST,PUT,DELETE]
- code: 状态码[200,400,500]
- latency: 请求耗时[ms]
- ip: 请求客户端IP
- req_id: 请求随机ID

### 日志格式

- json
- text

### 日志输出

- stdout: 控制台标准输出
- file: 日志文件输出[access.log,error.log]

### 日志切割

- 按日期切割
- 按文件大小切割

### 日志性能

- zap > glog > logrus > log

参考：[performance](https://github.com/uber-go/zap#performance)

| Package         | Time        | Time % to zap | Objects Allocated |
| --------------- | ----------- | ------------- | ----------------- |
| ⚡ zap           | 862 ns/op   | +0%           | 5 allocs/op       |
| ⚡ zap (sugared) | 1250 ns/op  | +45%          | 11 allocs/op      |
| zerolog         | 4021 ns/op  | +366%         | 76 allocs/op      |
| go-kit          | 4542 ns/op  | +427%         | 105 allocs/op     |
| apex/log        | 26785 ns/op | +3007%        | 115 allocs/op     |
| logrus          | 29501 ns/op | +3322%        | 125 allocs/op     |
| log15           | 29906 ns/op | +3369%        | 122 allocs/op     |
