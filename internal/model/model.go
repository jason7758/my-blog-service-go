package model

import (
	"fmt"
	otgorm "github.com/eddycjy/opentracing-gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"my-blog-service-go/global"
	"my-blog-service-go/pkg/setting"
	"time"
)

const (
	STATE_OPEN = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn   uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}
// 创建数据库连接
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error)  {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	otgorm.AddGormCallbacks(db) //add数据链路
	return db, nil
}

// 增加数据新增时回调，增加修改时默认创建时间或者修改时间戳的值
func updateTimeStampForCreateCallback(scope *gorm.Scope)  {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreateOn"); ok {
			if createTimeField.IsBlank {
				_= createTimeField.Set(nowTime)
			}
		}

		if modifyTimeFields, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeFields.IsBlank {
				_ = modifyTimeFields.Set(nowTime)
			}
		}
	}
}
//增加数据记录修改回调，增加默认「修改时间」戳
func updateTimeStampForUpdateCallback(scope *gorm.Scope)  {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

//增加数据删除回掉，
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}


func addExtraSpaceIfExist(str string) string  {
	if str != "" {
		return " " + str
	}
	return  ""
}
