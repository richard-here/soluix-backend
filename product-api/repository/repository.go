package repository

import (
	"fmt"
	"richard-here/soluix/product-api/database"

	localmodel "richard-here/soluix/product-api/model/local"
)

type Repo struct {
	Db database.DBInstance
}

func CreateRepository(db database.DBInstance) Repo {
	return Repo{Db: db}
}

func (r *Repo) GetProducts(pagination Pagination, name, subcategory, brand string, minPrice, maxPrice int, status string) (*Pagination, error) {
	var products []localmodel.Product

	tx := r.Db.Db
	if name != "" {
		tx = tx.Where("LOWER(name) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", name))
	}
	if subcategory != "" {
		tx = tx.Where("LOWER(subcategory) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", subcategory))
	}
	if brand != "" {
		tx = tx.Where("LOWER(brand) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", brand))
	}
	if minPrice != 0 {
		tx = tx.Where("retail_price >= ?", minPrice)
	}
	if maxPrice != 0 {
		tx = tx.Where("retail_price <= ?", maxPrice)
	}
	if status != "" {
		if status == "1" {
			tx = tx.Where("status = true")
		} else if status == "0" {
			tx = tx.Where("status = false")
		}
	}
	err := tx.Scopes(paginate(products, &pagination, tx)).Find(&products).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = products
	return &pagination, nil
}

func (r *Repo) AddProduct(p *localmodel.Product) error {
	err := r.Db.Db.Create(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateProduct(p *localmodel.Product, status bool, retailPrice int) error {
	err := r.Db.Db.First(&p).Error
	if err != nil {
		return err
	}

	p.Status = status
	p.RetailPrice = retailPrice
	err = r.Db.Db.Save(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetProduct(id string) (*localmodel.Product, error) {
	p := new(localmodel.Product)
	err := r.Db.Db.First(&p, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}
