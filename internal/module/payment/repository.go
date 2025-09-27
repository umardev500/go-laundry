package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/payment"

	paymentEntity "github.com/umardev500/go-laundry/ent/payment"
	"github.com/umardev500/go-laundry/ent/tenant"
)

type repositoryImpl struct {
	client *db.Client
}

// Update implements payment.Repository.
func (r *repositoryImpl) Update(ctx context.Context, payload *payment.PaymentUpdate, id uuid.UUID, TenantID *uuid.UUID) (*payment.Payment, error) {
	conn := r.client.GetConn(ctx)

	builder := conn.Payment.
		UpdateOneID(id).
		SetNillableAmount(payload.Amount).
		SetNillableStatus((*paymentEntity.Status)(payload.Status)).
		SetNillablePaidAt(payload.PaidAt)

	pymnt, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(pymnt), nil
}

// GetByID implements payment.Repository.
func (r *repositoryImpl) GetByID(ctx context.Context, id uuid.UUID, filter *payment.PaymentFilter, tenantID *uuid.UUID) (*payment.Payment, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Payment.
		Query().
		Where(paymentEntity.IDEQ(id))

	q = r.applyPaymentFilter(q, filter, tenantID)

	pymnt, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(pymnt), nil
}

// List implements payment.Repository.
func (r *repositoryImpl) List(ctx context.Context, filter *payment.PaymentFilter, tenantID *uuid.UUID) ([]*payment.Payment, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Payment.
		Query()

	q = r.applyPaymentFilter(q, filter, tenantID)

	pymnts, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnts(pymnts), nil
}

// Create implements payment.Repository.
func (r *repositoryImpl) Create(ctx context.Context, payload *payment.PaymentCreate) (*payment.Payment, error) {
	conn := r.client.GetConn(ctx)

	builder := conn.Payment.
		Create().
		SetUserID(payload.UserID).
		SetNillableTenantID(payload.TenantID).
		SetReferenceID(payload.ReferenceID).
		SetReferenceType(paymentEntity.ReferenceType(payload.ReferenceType)).
		SetAmount(payload.Amount).
		SetCurrency(paymentEntity.Currency(payload.Currency)).
		SetStatus(paymentEntity.Status(payload.Status)).
		SetNillablePaidAt(payload.PaidAt)

	pymnt, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(pymnt), nil
}

func (r *repositoryImpl) applyPaymentFilter(q *ent.PaymentQuery, filter *payment.PaymentFilter, tenantID *uuid.UUID) *ent.PaymentQuery {
	if filter.Status != nil {
		q = q.Where(paymentEntity.StatusEQ(paymentEntity.Status(*filter.Status)))
	}

	if filter.Type != nil {
		q = q.Where(paymentEntity.ReferenceTypeEQ(paymentEntity.ReferenceType(*filter.Type)))
	}

	if tenantID != nil {
		q = q.Where(paymentEntity.HasTenantWith(tenant.IDEQ(*tenantID)))
	}

	return q
}

func (r *repositoryImpl) mapFromEnts(es []*ent.Payment) []*payment.Payment {
	var result []*payment.Payment
	for _, e := range es {
		result = append(result, r.mapFromEnt(e))
	}
	return result
}

func (r *repositoryImpl) mapFromEnt(e *ent.Payment) *payment.Payment {

	return &payment.Payment{
		ID:            e.ID,
		UserID:        *e.UserID,
		TenantID:      e.TenantID,
		ReferenceID:   *e.ReferenceID,
		ReferenceType: payment.ReferenceType(*e.ReferenceType),
		Amount:        *e.Amount,
		Currency:      payment.Currency(*e.Currency),
		Status:        payment.Status(*e.Status),
		PaidAt:        e.PaidAt,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}

func NewRepository(client *db.Client) payment.Repository {
	return &repositoryImpl{
		client: client,
	}
}
