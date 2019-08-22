package services

import (
	"database/sql/driver"
	"github.com/dragonyui/crud_go_rest_sample/constants"
	"database/sql"
	"github.com/dragonyui/config"
	"errors"
)


type Dao interface {
	GetScanParams() ([]interface{})
	GetModelObject()(interface{})
	Insert()(error)
	Update() (error)
	Delete() (error)
	Get(id int64 ) (error)
	GetAll()(*[]interface{},error)

}

func  ValidateOperation(result driver.Result,err error) error{
	if config.FancyHandleError(err){
		return constants.ErrorMessageSwitcher(err)
	}
	if rowsAffected,err:=result.RowsAffected();err!=nil || rowsAffected==0{
		if config.FancyHandleError(err){
			return err
		}else {
			//insert / update but there is no row affected
			return errors.New("Product Hasn't been changed/inserted")
		}
	}
	return nil
}
func  LoadAllProduct(rows *sql.Rows,err error, dao Dao)(*[]interface{},error){
	if config.FancyHandleError(err){
		return nil,constants.ErrorMessageSwitcher(err)
	}
	defer rows.Close()
	objectsArray:=make([]interface{},0)
	for rows.Next(){
		err:=arrayRowLoader(&objectsArray,rows,dao)
		if config.FancyHandleError(err){
			return nil,err
		}

	}
	return &objectsArray,nil
}
func arrayRowLoader(arraysModelObject *[]interface{},rows *sql.Rows, dao Dao) error {
	err:=rows.Scan(dao.GetScanParams()...)
	if config.FancyHandleError(err){
		return err
	}
	(*arraysModelObject)=append((*arraysModelObject),dao.GetModelObject())
	return nil
}

