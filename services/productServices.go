package services

import (
	"github.com/dragonyui/crud_go_rest_sample/models"
	"github.com/dragonyui/config"
	"github.com/dragonyui/crud_go_rest_sample/constants"
	"database/sql"
)

type ProductDao struct {
	models.Product
}
func (p *ProductDao) getDb()(*sql.DB){
	return config.GetDefaultDb().GetDb()
}
func (p *ProductDao) Insert() (error) {
	return ValidateOperation(p.getDb().Exec(constants.QUERY_PRODUCT_INSERT, p.Sku, p.Name, p.Price, p.Qty))
}

func (p *ProductDao) Update() (error){
	return ValidateOperation(p.getDb().Exec(constants.QUERY_PRODUCT_UPDATE,p.Sku,p.Name,p.Price,p.Qty,p.Id))
}
func (p *ProductDao) Delete() (error){
	return ValidateOperation(p.getDb().Exec(constants.QUERY_PRODUCT_DELETE,p.Id))
}
func (p *ProductDao) Get(id int64 ) (error){
	return constants.ErrorMessageSwitcher(p.getDb().QueryRow(constants.QUERY_PRODUCT_GET_BY_ID,p.Id).Scan(p.GetScanParams()...))
}
func (p *ProductDao) GetAll()(*[]interface{},error){
	rows,err:=p.getDb().Query(constants.QUERY_PRODUCT_GET_ALL)
	return LoadAllProduct(rows,err,p)
}

func (p *ProductDao) GetScanParams() ([]interface{}){
	return []interface{}{&p.Id,&p.Sku,&p.Name,&p.Price,&p.Qty}
}
func (p *ProductDao) GetModelObject()(interface{})  {
	return p.Product
}