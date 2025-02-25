package service

import (
	"context"
	checkout "gomall_demo/rpc_gen/kitex_gen/checkout"
	"gomall_demo/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCheckout_Run(t *testing.T) {
	ctx := context.Background()

	initCheckoutClient()
	resp, err := checkoutClient.Checkout(ctx, &checkout.CheckoutReq{
		UserId:    3,
		Firstname: "John",
		Lastname:  "Smith",
		Email:     "1234@qq.com",
		Address: &checkout.Address{
			StreetAddress: "123 Main St",
			City:          "Anytown",
			State:         "CA",
			Country:       "China",
			ZipCode:       "12345",
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "1234567890123456",
			CreditCardCvv:             123,
			CreditCardExpirationYear:  2022,
			CreditCardExpirationMonth: 6,
		},
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
