package services

import (
	"Walker/app/http/model"
	"Walker/global"
	"errors"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type TestService struct {
}

func (service *TestService) Oversold() error {
	conn := global.ConfigRedis.Get()
	//购买商品
	goodsNum, err := redis.Int(conn.Do("Decr", "goods_num"))
	if err != nil {
		return err
	}
	if goodsNum >= 0 {
		//查询商品数量
		var goodsBuyLogModel model.GoodsBuyLog
		goodsBuyLogModel.UserId = 100
		model.UseCon().Create(&goodsBuyLogModel)
		model.UseCon().Model(model.Goods{}).Where("id=1").Update("goods_num", gorm.Expr("goods_num-?", 1))
	} else {
		return errors.New("买完了")
	}
	err = conn.Close()
	return nil
}

func (service *TestService) OversoldQueue() error {
	//获得一个redis客户端
	conn := global.ConfigRedis.Get()
	//减少库存（如果购买失败在取消顶顶那时候返回库存）
	reply, err := redis.Strings(conn.Do("lpop", "goods_stock", 1))
	if len(reply) == 0 || err != nil {
		return errors.New("买完了")
	}
	var goodsBuyLogModel model.GoodsBuyLog
	goodsBuyLogModel.UserId = 100
	model.UseCon().Create(&goodsBuyLogModel)
	model.UseCon().Model(model.Goods{}).Where("id=1").Update("goods_num", gorm.Expr("goods_num-?", 1))
	return nil
}

func (service *TestService) OversoldSQueue() error {
	userId := 100
	//获得一个redis客户端
	conn := global.ConfigRedis.Get()

	//增加设计一个人只能购买一次
	if reply, err := redis.Int(conn.Do("Sadd", "user_buy", userId)); err != nil || reply == 0 {
		return errors.New("每人仅限购买一次")
	}
	//增加设计一个人只能购买一次

	//减少库存（如果购买失败在取消顶顶那时候返回库存）
	reply, err := redis.Int(conn.Do("Spop", "goods_stock"))
	if reply == 0 || err != nil {
		return errors.New("买完了")
	}
	var goodsBuyLogModel model.GoodsBuyLog
	goodsBuyLogModel.UserId = userId
	model.UseCon().Create(&goodsBuyLogModel)
	model.UseCon().Model(model.Goods{}).Where("id=1").Update("goods_num", gorm.Expr("goods_num-?", 1))
	return nil
}

func (service *TestService) OversoldStrQueue() error {
	userId := 100
	num := 5
	//获得一个redis客户端
	conn := global.ConfigRedis.Get()
	//增加设计一个人只能购买一次
	if reply, err := redis.Int(conn.Do("Sadd", "user_buy", userId)); err != nil || reply == 0 {
		return errors.New("每人仅限购买一次")
	}
	//增加设计一个人只能购买一次

	//减少库存可以购买多个（如果购买失败在取消顶顶那时候返回库存）
	reply, err := redis.Int(conn.Do("Decrby", "goods_stock", num))
	if reply < 0 || err != nil {
		//返回库存
		conn.Do("Incrby", "goods_stock", num)
		return errors.New("库存不足")
	}
	var goodsBuyLogModel model.GoodsBuyLog
	goodsBuyLogModel.UserId = userId
	goodsBuyLogModel.Num = num
	model.UseCon().Create(&goodsBuyLogModel)
	model.UseCon().Model(model.Goods{}).Where("id=1").Update("goods_num", gorm.Expr("goods_num-?", num))
	return nil
}
