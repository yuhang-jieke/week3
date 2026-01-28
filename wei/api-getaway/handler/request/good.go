package request

type GoodsAdd struct {
	Name  string  `form:"name"   binding:"required"`
	Price float32 `form:"price"  binding:"required"`
	Stock int     `form:"stock"  binding:"required"`
}
type GoodsShow struct {
	Id int `form:"id"  binding:"required"`
}
type Es struct {
	KeyWord string `form:"keyword"  binding:"required"`
	Page    int    `form:"page"  binding:"required"`
	Size    int    `form:"size"  binding:"required"`
}
