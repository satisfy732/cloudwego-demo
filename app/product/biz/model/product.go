package model

import (
	"context"
	"time"
	"fmt"

	"github.com/redis/go-redis/v9"
	"encoding/json"
	
	"gorm.io/gorm"
)
type Product struct{
	Base
	Name string `json:"name"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	Price float32 `json:"price"`

	Categories [] Category`json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string{
	return "product"
}

type ProductQuery struct{
	ctx context.Context
	db *gorm.DB
}
func (p ProductQuery)GetById(productId int)(product Product,err error){
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product,productId).Error
    return product,err
}

func (p ProductQuery)SearchProducts(q string)(products []*Product,err error){
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products,"name like ? or description like ?",
	"%"+q+"%","%"+q+"%",
	).Error
    return products,err
}

func NewProductQuery(ctx context.Context,db *gorm.DB) *ProductQuery{
	return &ProductQuery{
		ctx:ctx,
		db:db,
	}
}

type CachedProductQuery struct{
	ProductQuery ProductQuery
	cacheClient *redis.Client
	prefix string 
}
func (c CachedProductQuery)GetById(productId int)(product Product,err error){
	cachedKey :=fmt.Sprintf("%s_%s_%d",c.prefix,"product_by_id",productId)
	cachedResult :=c.cacheClient.Get(c.ProductQuery.ctx,cachedKey)
	err = func() error{
		if err :=cachedResult.Err();err!=nil{
			return err
	}
	cachedResultBytes,err :=cachedResult.Bytes()
	if err !=nil{
		return err 
	}

	err = json.Unmarshal(cachedResultBytes,&product)
	if err !=nil{
		return err
	}
	return nil
	}()
	if err !=nil{
		product,err = c.ProductQuery.GetById(productId)
		if err !=nil{
			return Product{},err
		}
		encoded,err :=json.Marshal(product)
		if err !=nil{
			return product,nil
		}
		_ = c.cacheClient.Set(c.ProductQuery.ctx,cachedKey,encoded,time.Hour)

	}
	return 
}

func (c CachedProductQuery)SearchProducts(q string)(products []*Product,err error){
	return c.ProductQuery.SearchProducts(q)
}

func NewCachedProductQuery(ctx context.Context,db *gorm.DB,cacheClient *redis.Client) *CachedProductQuery{
	return &CachedProductQuery{
		ProductQuery:ProductQuery{
			ctx:ctx,
			db:db,
		},
		cacheClient:cacheClient,
		prefix:"shop",
	}
}

type ProductMutation struct{
	ctx context.Context
	db *gorm.DB
}
