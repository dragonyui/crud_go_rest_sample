package constants

const QUERY_PRODUCT_INSERT  = `INSERT INTO product (sku, name, price, qty) VALUES (?, ?, ?, ?)`
const QUERY_PRODUCT_UPDATE  = `UPDATE product set sku=?, name=?, price=?, qty=? where id=?`
const QUERY_PRODUCT_DELETE  = `DELETE FROM product  where id=?`
const QUERY_PRODUCT_GET_BY_ID=`SELECT id, sku, name, price, qty from product where id=?`
const QUERY_PRODUCT_GET_ALL =`SELECT id, sku, name, price, qty from product`

