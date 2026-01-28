package server

import (
	"context"
	"log"
	"reflect"
	"week3/wei/user-server/basic/config"
	"week3/wei/user-server/basic/inits"
	__ "week3/wei/user-server/handler/proto"
	"week3/wei/user-server/model"

	"github.com/olivere/elastic/v7"
)

type Server struct {
	__.UnimplementedGoodsServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodsAdd(_ context.Context, in *__.GoodsAddReq) (*__.GoodsAddResp, error) {
	log.Printf("Received: %v", in.GetName())
	goods := model.Goodss{
		Name:  in.Name,
		Price: in.Price,
		Stock: int(in.Stock),
	}
	if err := goods.GoodsAdd(config.DB); err != nil {
		panic("商品上架失败")
	}
	return &__.GoodsAddResp{
		Greet: "商品上架成功",
	}, nil
}
func (s *Server) GoodsList(_ context.Context, in *__.GoodsListReq) (*__.GoodsListResp, error) {
	var list []*__.List
	var goods model.Goodss
	var err error
	if list, err = goods.GoodsShow(config.DB, int(in.Id)); err != nil {
		panic("查询失败")
	}
	return &__.GoodsListResp{
		Greet: "商品查询成功",
		List:  list,
	}, nil
}
func (s *Server) Es(_ context.Context, in *__.EsReq) (*__.EsResp, error) {
	boolQuery := elastic.NewBoolQuery()
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 || in.Size > 3 {
		in.Size = 3
	}
	if in.Keyword != "" {
		keyQuery := elastic.NewMultiMatchQuery(in.Keyword, "name")
		boolQuery.Must(keyQuery)
	}
	Query := inits.ElasticClient.Search().Index("goodsses").Query(boolQuery)
	offset := (in.Page - 1) * in.Size
	res, err := Query.From(int(offset)).Size(int(in.Size)).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var list []*__.Es
	var show *__.Es
	for _, item := range res.Each(reflect.TypeOf(show)) { //从搜索结果中取数据的方法
		t := item.(*__.Es)
		list = append(list, t)
	}
	return &__.EsResp{
		Greet: "es查询成功",
		Es:    list,
	}, nil
}
