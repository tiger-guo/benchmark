package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// BareMetalServer defines bare metal server struct.
type BareMetalServer struct {
	InstanceId         string        `db:"instanceId" json:"InstanceId"`
	VpcId              string        `db:"vpcId" json:"VpcId"`
	SubnetId           string        `db:"subnetId" json:"SubnetId"`
	DeviceStatus       int           `db:"deviceStatus" json:"DeviceStatus"`
	OperateStatus      int           `db:"operateStatus" json:"OperateStatus"`
	OsTypeId           int           `db:"osTypeId" json:"OsTypeId"`
	RaidId             int           `db:"raidId" json:"RaidId"`
	Alias              string        `db:"alias" json:"Alias"`
	AppId              int           `db:"appId" json:"AppId"`
	Zone               string        `db:"zone" json:"Zone"`
	WanIp              string        `db:"wanIp" json:"WanIp"`
	LanIp              string        `db:"lanIp" json:"LanIp"`
	DeliverTime        string        `db:"deliverTime" json:"DeliverTime"`
	Deadline           string        `db:"deadline" json:"Deadline"`
	AutoRenewFlag      int           `db:"autoRenewFlag" json:"AutoRenewFlag"`
	DeviceClassCode    string        `db:"deviceClassCode" json:"DeviceClassCode"`
	CpmPayMode         int           `db:"cpmPayMode" json:"CpmPayMode"`
	DhcpIp             string        `db:"dhcpIp" json:"DhcpIp"`
	VpcName            string        `db:"vpcName" json:"VpcName"`
	SubnetName         string        `db:"subnetName" json:"SubnetName"`
	SubnetCidrBlock    string        `db:"subnetCidrBlock" json:"SubnetCidrBlock"`
	IsLuckyDevice      int           `db:"isLuckyDevice" json:"IsLuckyDevice"`
	ExtMaintainMessage string        `db:"extMaintainMessage" json:"ExtMaintainMessage"`
	BMSExtension       *BMSExtension `db:"bmsExtension" json:"BMSExtension"`
}

// BMSExtension define bare metal server extension strcut.
type BMSExtension struct {
	MaintainStatus  string   `db:"maintainStatus" json:"MaintainStatus"`
	MaintainMessage string   `db:"maintainMessage" json:"MaintainMessage"`
	VpcCidrBlock    []string `db:"vpcCidrBlock" json:"VpcCidrBlock"`
	Size            []int    `db:"size" json:"Size"`
	Tags            []*Tags  `db:"Tags" json:"Tags"`
}

// Tags defines tags struct.
type Tags struct {
	TagKey    string   `db:"tagKey" json:"TagKey"`
	TagValues []string `db:"tagValues" json:"TagValues"`
}

// Scan is used to decode raw message which is read from db into a structured
// ScopeSelector instance.
func (bms *BMSExtension) Scan(raw interface{}) error {
	if bms == nil {
		return errors.New("bms extension is not initialized")
	}

	if raw == nil {
		return errors.New("raw is nil, can not be decoded")
	}

	switch v := raw.(type) {
	case []byte:
		if err := json.Unmarshal(v, &bms); err != nil {
			return fmt.Errorf("decode into bms extension failed, err: %v", err)

		}
		return nil
	case string:
		if err := json.Unmarshal([]byte(v), &bms); err != nil {
			return fmt.Errorf("decode into bms extension failed, err: %v", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported bms extension raw type: %T", v)
	}
}

// Value encode the scope selector to a json raw, so that it can be stored to db with
// json raw.
func (bms *BMSExtension) Value() (driver.Value, error) {
	if bms == nil {
		return nil, errors.New("scope selector is not initialized, can not be encoded")
	}

	return json.Marshal(bms)
}
