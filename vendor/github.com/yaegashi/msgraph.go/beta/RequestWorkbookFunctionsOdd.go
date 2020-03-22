// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsOddRequestBuilder struct{ BaseRequestBuilder }

// Odd action undocumented
func (b *WorkbookFunctionsRequestBuilder) Odd(reqObj *WorkbookFunctionsOddRequestParameter) *WorkbookFunctionsOddRequestBuilder {
	bb := &WorkbookFunctionsOddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/odd"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsOddRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsOddRequestBuilder) Request() *WorkbookFunctionsOddRequest {
	return &WorkbookFunctionsOddRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsOddRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsOddFPriceRequestBuilder struct{ BaseRequestBuilder }

// OddFPrice action undocumented
func (b *WorkbookFunctionsRequestBuilder) OddFPrice(reqObj *WorkbookFunctionsOddFPriceRequestParameter) *WorkbookFunctionsOddFPriceRequestBuilder {
	bb := &WorkbookFunctionsOddFPriceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/oddFPrice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsOddFPriceRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsOddFPriceRequestBuilder) Request() *WorkbookFunctionsOddFPriceRequest {
	return &WorkbookFunctionsOddFPriceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsOddFPriceRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsOddFYieldRequestBuilder struct{ BaseRequestBuilder }

// OddFYield action undocumented
func (b *WorkbookFunctionsRequestBuilder) OddFYield(reqObj *WorkbookFunctionsOddFYieldRequestParameter) *WorkbookFunctionsOddFYieldRequestBuilder {
	bb := &WorkbookFunctionsOddFYieldRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/oddFYield"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsOddFYieldRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsOddFYieldRequestBuilder) Request() *WorkbookFunctionsOddFYieldRequest {
	return &WorkbookFunctionsOddFYieldRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsOddFYieldRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsOddLPriceRequestBuilder struct{ BaseRequestBuilder }

// OddLPrice action undocumented
func (b *WorkbookFunctionsRequestBuilder) OddLPrice(reqObj *WorkbookFunctionsOddLPriceRequestParameter) *WorkbookFunctionsOddLPriceRequestBuilder {
	bb := &WorkbookFunctionsOddLPriceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/oddLPrice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsOddLPriceRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsOddLPriceRequestBuilder) Request() *WorkbookFunctionsOddLPriceRequest {
	return &WorkbookFunctionsOddLPriceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsOddLPriceRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsOddLYieldRequestBuilder struct{ BaseRequestBuilder }

// OddLYield action undocumented
func (b *WorkbookFunctionsRequestBuilder) OddLYield(reqObj *WorkbookFunctionsOddLYieldRequestParameter) *WorkbookFunctionsOddLYieldRequestBuilder {
	bb := &WorkbookFunctionsOddLYieldRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/oddLYield"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsOddLYieldRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsOddLYieldRequestBuilder) Request() *WorkbookFunctionsOddLYieldRequest {
	return &WorkbookFunctionsOddLYieldRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsOddLYieldRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}