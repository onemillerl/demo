package product

import (
	"context"
	product "gomall_demo/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResp, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (resp *product.SearchProductsResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateProduct(ctx context.Context, req *product.CreateProductReq, callOptions ...callopt.Option) (resp *product.CreateProductResp, err error) {
	resp, err = defaultClient.CreateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProduct(ctx context.Context, req *product.UpdateProductReq, callOptions ...callopt.Option) (resp *product.UpdateProductResp, err error) {
	resp, err = defaultClient.UpdateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *product.DeleteProductReq, callOptions ...callopt.Option) (resp *product.DeleteProductResp, err error) {
	resp, err = defaultClient.DeleteProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProductcenter(ctx context.Context, req *product.GetProductcenterReq, callOptions ...callopt.Option) (resp *product.GetProductcenterResp, err error) {
	resp, err = defaultClient.GetProductcenter(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductcenter call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Searchproductcenters(ctx context.Context, req *product.SearchproductcentersReq, callOptions ...callopt.Option) (resp *product.SearchproductcentersResp, err error) {
	resp, err = defaultClient.Searchproductcenters(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Searchproductcenters call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
