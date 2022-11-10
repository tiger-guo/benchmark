package gener

import (
	"context"
	"fmt"
	"time"

	"github.com/tiger-guo/benchmark/mysql/orm"
	"github.com/tiger-guo/benchmark/mysql/types"
	"github.com/tiger-guo/benchmark/pkg/errf"
	"github.com/tiger-guo/benchmark/pkg/rand"
	"github.com/tiger-guo/benchmark/pkg/uuid"
)

/*
############## Mysql DB命令 ##############

1. 创建表结构
create table if not exists bare_metal_server
(
    instanceId      varchar(64) default '',
    vpcId           varchar(64) default '',
    subnetId        varchar(64) default '',
    alias           varchar(64) default '',
    zone            varchar(64) default '',
    wanIp           varchar(64) default '',
    lanIp           varchar(64) default '',
    deliverTime     varchar(64) default '',
    deadline        varchar(64) default '',
    deviceClassCode varchar(64) default '',
    dhcpIp          varchar(64) default '',
    vpcName         varchar(64) default '',
    subnetName      varchar(64) default '',
    subnetCidrBlock varchar(64) default '',
    bmsExtension    json        default null,
    deviceStatus    bigint(1)   default 0,
    operateStatus   bigint(1)   default 0,
    osTypeId        bigint(1)   default 0,
    raidId          bigint(1)   default 0,
    appId           bigint(1)   default 0,
    autoRenewFlag   bigint(1)   default 0,
    cpmPayMode      bigint(1)   default 0,
    isLuckyDevice   bigint(1)   default 0,

    primary key (instanceId)
) engine = innodb default charset = utf8mb4;

2. 构建json索引
alter table `bare_metal_server` add column `extMaintainMessage` varchar(64) as (`bmsExtension`->>"$.MaintainMessage");
alter table `bare_metal_server` add index `index_emm`(`extMaintainMessage`);

3. 构建JSON数组索引
ALTER TABLE `bare_metal_server` ADD INDEX size((CAST(bmsExtension->'$.Size' AS UNSIGNED ARRAY)));
*/

// BenchJsonData 生成Mysql Json字段压测数据
func BenchJsonData() {
	generDataTotal := 1000000

	orm, err := orm.NewORM()
	errf.CheckErr(err)
	defer orm.Close()

	sql := fmt.Sprintf(`insert into bare_metal_server
	(instanceId,vpcId,subnetId,alias,zone,wanIp,lanIp,deliverTime,deadline,deviceClassCode,dhcpIp,vpcName,subnetName,
	subnetCidrBlock,bmsExtension,deviceStatus,operateStatus,osTypeId,raidId,appId,autoRenewFlag,cpmPayMode,isLuckyDevice)
	Values(:instanceId,:vpcId,:subnetId,:alias,:zone,:wanIp,:lanIp,:deliverTime,:deadline,:deviceClassCode,:dhcpIp,
	:vpcName,:subnetName,:subnetCidrBlock,:bmsExtension,:deviceStatus,:operateStatus,:osTypeId,:raidId,:appId,
	:autoRenewFlag,:cpmPayMode,:isLuckyDevice)`)

	bm := &types.BareMetalServer{
		InstanceId:      rand.StringN("inst", 6),
		VpcId:           "127.0.0.1",
		SubnetId:        "127.10.0.1",
		DeviceStatus:    1,
		OperateStatus:   2,
		OsTypeId:        3,
		RaidId:          4,
		Alias:           rand.StringN("name", 4),
		AppId:           5,
		Zone:            "ap-guangzhou",
		WanIp:           "127.10.125.1",
		LanIp:           "127.10.125.1",
		DeliverTime:     time.Now().Format(time.RFC3339),
		Deadline:        "unUsed",
		AutoRenewFlag:   6,
		DeviceClassCode: "s3",
		CpmPayMode:      7,
		DhcpIp:          "127.10.125.2",
		VpcName:         rand.StringN("vpc", 4),
		SubnetName:      rand.StringN("subnet", 4),
		SubnetCidrBlock: "127.10.125.1",
		IsLuckyDevice:   8,
		BMSExtension: &types.BMSExtension{
			MaintainStatus:  "ok",
			MaintainMessage: rand.String(6),
			VpcCidrBlock:    []string{},
			Size:            []int{},
			Tags:            nil,
		},
	}

	for i := 0; i < generDataTotal; i++ {
		bm.InstanceId = "inst-" + uuid.UUID()
		bm.Alias = "name-" + uuid.UUID()
		bm.VpcName = rand.StringN("vpc", 4)
		bm.SubnetName = rand.StringN("subnet", 4)
		bm.BMSExtension.VpcCidrBlock = []string{rand.String(5), rand.String(5)}
		bm.BMSExtension.MaintainStatus = "state-" + uuid.UUID()
		bm.BMSExtension.Size = []int{rand.Number(), rand.Number()}
		bm.BMSExtension.MaintainMessage = "msg-" + uuid.UUID()
		bm.BMSExtension.Tags = []*types.Tags{
			{
				TagKey:    "name",
				TagValues: []string{rand.String(7)},
			},
		}

		err := orm.Insert(context.Background(), sql, bm)
		errf.CheckErr(err)
	}

}
