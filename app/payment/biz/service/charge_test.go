package service

import (
	"context"
	payment "gomall_demo/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()

	initPaymentClient()
	resp, err := paymentClient.Charge(ctx, &payment.ChargeReq{
		Amount: 1,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "1234567890123456",
			CreditCardCvv:             123,
			CreditCardExpirationYear:  2022,
			CreditCardExpirationMonth: 6,
		},
		OrderId: "3",
		UserId:  3,
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
