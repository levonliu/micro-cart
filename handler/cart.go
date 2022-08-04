package handler

import (
	"context"
	"go.micro.service.cart/common"
	"go.micro.service.cart/domain/model"
	"go.micro.service.cart/domain/service"
	"go.micro.service.cart/proto/cart"
)

type Cart struct {
	CartDataService service.ICartDataService
}

//
//
//DeleteItemByID(context.Context, *CarID, *Response) error
//GetAll(context.Context, *CartFindAll, *CartAll) error

func (c *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) error {
	cartInfo := &model.Cart{}
	if err := common.SwapTo(request, cartInfo); err != nil {
		return err
	}

	id, err := c.CartDataService.AddCart(cartInfo)
	if err != nil {
		return err
	}

	response.Id = id
	response.Message = "新增购物车成功"

	return nil
}

func (c *Cart) CleanCart(ctx context.Context, request *cart.Clean, response *cart.Response) error {
	err := c.CartDataService.CleanCart(request.UserId)
	if err != nil {
		return err
	}

	response.Message = "清空购物车成功"

	return nil
}

func (c *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	err := c.CartDataService.IncrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}

	response.Message = "添加购物车成功"

	return nil
}

func (c *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	err := c.CartDataService.DecrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}

	response.Message = "删除数量成功"

	return nil
}

func (c *Cart) DeleteItemByID(ctx context.Context, request *cart.CarID, response *cart.Response) error {
	err := c.CartDataService.DeleteCart(request.Id)
	if err != nil {
		return err
	}

	response.Message = "删除购物车成功"

	return nil
}

func (c *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error {
	carts, err := c.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}

	for _, ca := range carts {
		cr := &cart.CartInfo{}
		err := common.SwapTo(ca, cr)
		if err != nil {
			return err
		}

		response.CartInfo = append(response.CartInfo, cr)
	}
	return nil
}
