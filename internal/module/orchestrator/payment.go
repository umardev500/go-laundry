package orchestrator

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/payment"
	"github.com/umardev500/go-laundry/internal/domain/subscription"
)

type PaymentService struct {
	subSrv     subscription.Service
	paymentSrv payment.Service
	client     *db.Client
}

func NewPaymentService(subSrv subscription.Service, paymentSrv payment.Service, client *db.Client) *PaymentService {
	return &PaymentService{
		subSrv:     subSrv,
		paymentSrv: paymentSrv,
		client:     client,
	}
}

func (p *PaymentService) ProcessPayment(ctx context.Context, id uuid.UUID) (*payment.Payment, error) {
	var updatedPayment *payment.Payment

	err := p.client.WithTransaction(ctx, func(ctx context.Context) error {

		// Set payment to completed
		pymnt, err := p.paymentSrv.Update(ctx, &payment.PaymentUpdate{
			Status: func() *payment.Status {
				status := payment.Completed
				return &status
			}(),
		}, id, nil)
		if err != nil {
			return err
		}

		refID := pymnt.ReferenceID
		refType := pymnt.ReferenceType

		switch refType {
		case payment.Subscription:
			_, err := p.subSrv.Activate(ctx, refID)
			if err != nil {
				return err
			}
		}

		updatedPayment, err = p.paymentSrv.GetByID(ctx, id, &payment.PaymentFilter{
			IncludeMethod:     true,
			IncludeMethodType: true,
			IncludeTenant:     true,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return updatedPayment, err
}
