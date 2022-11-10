package jsonfield

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tiger-guo/benchmark/mysql/orm"
	"github.com/tiger-guo/benchmark/mysql/types"
	"github.com/tiger-guo/benchmark/pkg/errf"
)

const (
	// 以下字段是各压测场景的索引字段，可以通过 select * from bare_metal_server limit 500000,1;
	// 获取生成的100w数据中的中间数据的一下字段值用于索引，
	instanceID         = "inst-907ba054-60c8-11ed-bcbd-acde48001122"
	alias              = "name-907ba054-60c8-11ed-bcbd-acde48001122"
	extMaintainMessage = "msg-907ba2e8-60c8-11ed-bcbd-acde48001122"
	maintainStatus     = "state-907ba234-60c8-11ed-bcbd-acde48001122"
	size               = 283930890
)

func BenchmarkSelectNormalIndexField(b *testing.B) {
	orm, err := orm.NewORM()
	errf.CheckErr(err)
	defer orm.Close()

	sql := fmt.Sprintf("select * from bare_metal_server where instanceID = '%s'", instanceID)
	bm := new(types.BareMetalServer)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err = orm.Get(context.Background(), bm, sql)
		errf.CheckErr(err)

		// printBMS(bm)
	}
	b.StopTimer()
}

func BenchmarkSelectNormalField(b *testing.B) {
	orm, err := orm.NewORM()
	errf.CheckErr(err)
	defer orm.Close()

	sql := fmt.Sprintf("select * from bare_metal_server where alias = '%s'", alias)
	bm := new(types.BareMetalServer)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err = orm.Get(context.Background(), bm, sql)
		errf.CheckErr(err)

		// printBMS(bm)
	}
	b.StopTimer()
}

func BenchmarkSelectJsonIndexField(b *testing.B) {
	orm, err := orm.NewORM()
	errf.CheckErr(err)
	defer orm.Close()

	sql := fmt.Sprintf("select * from bare_metal_server where extMaintainMessage = '%s'", extMaintainMessage)
	bm := new(types.BareMetalServer)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err = orm.Get(context.Background(), bm, sql)
		errf.CheckErr(err)

		// printBMS(bm)
	}
	b.StopTimer()
}

func BenchmarkSelectJsonUnIndexField(b *testing.B) {
	orm, err := orm.NewORM()
	errf.CheckErr(err)
	defer orm.Close()

	sql := fmt.Sprintf("select * from bare_metal_server where bmsExtension->'$.MaintainStatus' = '%s'", maintainStatus)
	bm := new(types.BareMetalServer)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err = orm.Get(context.Background(), bm, sql)
		errf.CheckErr(err)

		// printBMS(bm)
	}
	b.StopTimer()
}

func BenchmarkSelectArrayIndexField(b *testing.B) {
	orm, err := orm.NewORM()
	errf.CheckErr(err)
	defer orm.Close()

	sql := fmt.Sprintf("SELECT * FROM bare_metal_server WHERE JSON_CONTAINS(bmsExtension->'$.Size',CAST('[%d]' AS JSON))", size)
	bm := new(types.BareMetalServer)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err = orm.Get(context.Background(), bm, sql)
		errf.CheckErr(err)

		// printBMS(bm)
	}
	b.StopTimer()
}

func printBMS(bms *types.BareMetalServer) {
	marshal, err := json.Marshal(bms)
	errf.CheckErr(err)

	fmt.Printf("[ ---- BareMetalServer ---- ]: %s \n", marshal)
}
