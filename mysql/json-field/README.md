# Mysql Json字段性能压测 
## 压测背景
对比Mysql Json字段和常规字段通用查询场景下的查询性能差距。因为MySQL 8.0.17支持了 Mysql Json字段索引和Json数组字段索引。

## 压测数据
```
{
	"InstanceId": "inst-907ba054-60c8-11ed-bcbd-acde48001122",
	"VpcId": "127.0.0.1",
	"SubnetId": "127.10.0.1",
	"DeviceStatus": 1,
	"OperateStatus": 2,
	"OsTypeId": 3,
	"RaidId": 4,
	"Alias": "name-ajnb",
	"AppId": 5,
	"Zone": "ap-guangzhou",
	"WanIp": "127.10.125.1",
	"LanIp": "127.10.125.1",
	"DeliverTime": "2022-11-10T15:17:19+08:00",
	"Deadline": "unUsed",
	"AutoRenewFlag": 6,
	"DeviceClassCode": "s3",
	"CpmPayMode": 7,
	"DhcpIp": "127.10.125.2",
	"VpcName": "vpc-cldm",
	"SubnetName": "subnet-rxdq",
	"SubnetCidrBlock": "127.10.125.1",
	"IsLuckyDevice": 8,
	"ExtMaintainMessage": "msg-907ba2e8-60c8-11ed-bcbd-acde48001122",
	"BMSExtension": {
		"MaintainStatus": "state-907ba234-60c8-11ed-bcbd-acde48001122",
		"MaintainMessage": "msg-907ba2e8-60c8-11ed-bcbd-acde48001122",
		"VpcCidrBlock": ["nrpgk", "etipe"],
		"Size": [283930890, 63856000],
		"Tags": [{
			"TagKey": "name",
			"TagValues": ["zfgkcch"]
		}]
	}
}
```
随机生成100w以上数据结构数据。
索引：
-  `InstanceId` 主键索引
-  `MaintainMessage` Json虚拟键索引
-  `Size` Json数组索引

## 压测场景
1. Mysql 非Json索引字段查询 （`InstanceId`）
2. Mysql 非Json非索引字段查询（`Alias`）
3. Mysql Json索引字段查询（`MaintainMessage`）
4. Mysql Json非索引字段查询（`MaintainStatus`）
5. Mysql Json数组索引查询（`Size`）

`注`：以上压测场景命中数据均是一条。

## 压测环境
因为只是对比Mysql Json字段和非Json字段的性能差距，所以使用本机mac进行压测。
Mac配置：处理器（2.6 GHz 六核Intel Core i7） 内存（16 GB 2667 MHz DDR4）

## 压测结果
```
goos: darwin
goarch: amd64
pkg: github.com/tiger-guo/benchmark/mysql/json-field
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSelectNormalIndexField-12        234063            143614 ns/op            5336 B/op        105 allocs/op      (Mysql 非Json索引字段查询)
BenchmarkSelectNormalField-12                 19        1804771724 ns/op            5557 B/op        105 allocs/op      (Mysql 非Json非索引字段查询)
BenchmarkSelectJsonIndexField-12          215239            177479 ns/op            5336 B/op        105 allocs/op	    (Mysql Json索引字段查询)
BenchmarkSelectJsonUnIndexField-12            19        1744164361 ns/op            5557 B/op        105 allocs/op	    (Mysql Json非索引字段查询)
BenchmarkSelectArrayIndexField-12         181699            202587 ns/op            5336 B/op        105 allocs/op	    (Mysql Json数组索引查询)
PASS
ok      github.com/tiger-guo/benchmark/mysql/json-field 185.024s
```

### 总结：
Json索引字段的QPS是 `5634`，非Json索引字段的QPS是`6963`，有`1/5`的QPS性能差距。

## 附件：
生成数据代码：https://github.com/tiger-guo/benchmark/blob/main/mysql/gener/json_field.go
压测代码：https://github.com/tiger-guo/benchmark/blob/main/mysql/json-field/select_test.go
