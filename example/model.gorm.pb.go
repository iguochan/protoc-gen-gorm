// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// - protoc-gen-gorm v
// - protoc          v4.25.1
// source: model.proto

package model

import (
	driver "database/sql/driver"
	json "encoding/json"
	fmt "fmt"
	mysqlx "github.com/IguoChan/go-pkg/mysqlx"
	ptypes "github.com/golang/protobuf/ptypes"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	datatypes "gorm.io/datatypes"
	gorm "gorm.io/gorm"
	reflect "reflect"
	time "time"
)

// GormModel is an option for User_Author
type User_AuthorModel struct {
	Name        string         `gorm:"column:name" json:"name,omitempty"`
	Email       string         `gorm:"column:email" json:"email,omitempty"`
	DeletedTime gorm.DeletedAt `gorm:"column:deleted_time;type:datetime;index:idx_deleted_time" json:"deleted_time,omitempty"`
}

func (x *User_AuthorModel) User_Author() (*User_Author, error) {
	deletedTime, err := ptypes.TimestampProto(time.Time(x.DeletedTime.Time))
	if err != nil {
		return nil, err
	}

	return &User_Author{
		Name:        x.Name,
		Email:       x.Email,
		DeletedTime: deletedTime,
	}, nil
}

func (x *User_Author) User_AuthorModel() (*User_AuthorModel, error) {
	deletedTimeTime, err := ptypes.Timestamp(x.DeletedTime)
	if err != nil {
		return nil, err
	}
	deletedTime := gorm.DeletedAt{Time: deletedTimeTime}

	return &User_AuthorModel{
		Name:        x.Name,
		Email:       x.Email,
		DeletedTime: deletedTime,
	}, nil
}

var _ *timestamp.Timestamp

// GormModel is an option for User
type UserModel struct {
	Id          int32             `gorm:"column:id;type:bigint;primaryKey" json:"id,omitempty"`
	Name        string            `gorm:"column:name;index:idx_name_addr,unique" json:"name,omitempty"`
	Email       string            `gorm:"column:email" json:"email,omitempty"`
	Address     string            `gorm:"column:address;index:idx_name_addr,unique" json:"address,omitempty"`
	Phone       string            `gorm:"column:phone" json:"phone,omitempty"`
	Score       float32           `gorm:"column:score;type:decimal(10,2) unsigned;not null" json:"score,omitempty"`
	Balance     float64           `gorm:"column:balance" json:"balance,omitempty"`
	CreateTime  time.Time         `gorm:"column:create_time;type:datetime;autoCreateTime;not null" json:"create_time,omitempty"`
	UpdateTime  time.Time         `gorm:"column:update_time;type:datetime;autoUpdateTime;not null" json:"update_time,omitempty"`
	DeletedTime time.Time         `gorm:"column:deleted_time;type:datetime;not null" json:"deleted_time,omitempty"`
	Date        datatypes.Date    `gorm:"column:date;type:date" json:"date,omitempty"`
	Extra       []byte            `gorm:"column:extra;type:tinyblob;default:'abc'" json:"extra,omitempty"`
	Author      *User_AuthorModel `gorm:"-" json:"author,omitempty"`
	Status      string            `gorm:"column:status;type:enum('PENDING','UNKNOWN');default:'UNKNOWN'" json:"status,omitempty"`
	Status1     string            `gorm:"column:status1;default:'UNKNOWN'" json:"status1,omitempty"`
	Status2     Status            `gorm:"column:status2;type:int;default:0" json:"status2,omitempty"`
	T1          time.Time         `gorm:"column:t1;type:datetime;autoCreateTime;not null" json:"t1,omitempty"`
	T2          time.Time         `gorm:"column:t2;type:datetime;autoUpdateTime;not null" json:"t2,omitempty"`
	T3          time.Time         `gorm:"column:t3;type:datetime;not null" json:"t3,omitempty"`
	D1          datatypes.Date    `gorm:"column:d1;type:date" json:"d1,omitempty"`
	Ids         UserModelIdsList  `gorm:"column:ids;type:longtext" json:"ids,omitempty"`
	M1          UserModelM1Map    `gorm:"column:m1;type:longtext" json:"m1,omitempty"`
}

func (x *UserModel) TableName() string {
	return "user"
}

type UserModelIdsList []int32

func (x UserModelIdsList) Value() (driver.Value, error) {
	return Value(x)
}

func (x *UserModelIdsList) Scan(value any) error {
	return Scan(x, value)
}

type UserModelM1Map map[string]string

func (x UserModelM1Map) Value() (driver.Value, error) {
	return Value(x)
}

func (x *UserModelM1Map) Scan(value any) error {
	return Scan(x, value)
}

// Value for valuer helper
func Value(data any) (any, error) {
	v := reflect.ValueOf(data)
	if v.IsZero() {
		return nil, nil
	}
	return json.Marshal(data)
}

// Scan for scanner helper
func Scan(data any, value any) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, data)
	case string:
		return json.Unmarshal([]byte(v), data)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into Go value type %T", v, data)
	}
}

func (x *UserModel) User() (*User, error) {
	createTime, err := ptypes.TimestampProto(x.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(x.UpdateTime)
	if err != nil {
		return nil, err
	}

	deletedTime, err := ptypes.TimestampProto(x.DeletedTime)
	if err != nil {
		return nil, err
	}

	date, err := ptypes.TimestampProto(time.Time(x.Date))
	if err != nil {
		return nil, err
	}

	author, err := x.Author.User_Author()
	if err != nil {
		return nil, err
	}

	status := Status(Status_value[x.Status])

	status1 := Status(Status_value[x.Status1])

	t1 := int64(x.T1.Unix())

	t2 := uint64(x.T2.Unix())

	t3 := x.T3.Format("2006-01-02 15:04:05")

	d1 := int64(time.Time(x.D1).Unix())

	ids := []int32(x.Ids)

	m1 := map[string]string(x.M1)

	return &User{
		Id:          x.Id,
		Name:        x.Name,
		Email:       x.Email,
		Address:     x.Address,
		Phone:       x.Phone,
		Score:       x.Score,
		Balance:     x.Balance,
		CreateTime:  createTime,
		UpdateTime:  updateTime,
		DeletedTime: deletedTime,
		Date:        date,
		Extra:       x.Extra,
		Author:      author,
		Status:      status,
		Status1:     status1,
		Status2:     x.Status2,
		T1:          t1,
		T2:          t2,
		T3:          t3,
		D1:          d1,
		Ids:         ids,
		M1:          m1,
	}, nil
}

func (x *User) UserModel() (*UserModel, error) {
	createTime, err := ptypes.Timestamp(x.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.Timestamp(x.UpdateTime)
	if err != nil {
		return nil, err
	}

	deletedTime, err := ptypes.Timestamp(x.DeletedTime)
	if err != nil {
		return nil, err
	}

	dateTime, err := ptypes.Timestamp(x.Date)
	if err != nil {
		return nil, err
	}
	date := datatypes.Date(dateTime)

	author, err := x.Author.User_AuthorModel()
	if err != nil {
		return nil, err
	}

	status := x.Status.String()

	status1 := x.Status1.String()

	t1 := time.Unix(int64(x.T1), 0)

	t2 := time.Unix(int64(x.T2), 0)

	t3, err := time.Parse("2006-01-02 15:04:05", x.T3)
	if err != nil {
		return nil, err
	}

	d1 := datatypes.Date(time.Unix(int64(x.D1), 0))

	ids := UserModelIdsList(x.Ids)

	m1 := UserModelM1Map(x.M1)

	return &UserModel{
		Id:          x.Id,
		Name:        x.Name,
		Email:       x.Email,
		Address:     x.Address,
		Phone:       x.Phone,
		Score:       x.Score,
		Balance:     x.Balance,
		CreateTime:  createTime,
		UpdateTime:  updateTime,
		DeletedTime: deletedTime,
		Date:        date,
		Extra:       x.Extra,
		Author:      author,
		Status:      status,
		Status1:     status1,
		Status2:     x.Status2,
		T1:          t1,
		T2:          t2,
		T3:          t3,
		D1:          d1,
		Ids:         ids,
		M1:          m1,
	}, nil
}

func UserModelId(x int32) mysqlx.Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id= ?", x)
	}
}

func UserModelName(x string) mysqlx.Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name= ?", x)
	}
}

func UserModelAddress(x string) mysqlx.Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("address= ?", x)
	}
}

var _ *timestamp.Timestamp
