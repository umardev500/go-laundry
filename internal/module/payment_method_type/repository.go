package paymentmethodtype

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	paymentMethodTypeEntity "github.com/umardev500/go-laundry/ent/paymentmethodtype"
	"github.com/umardev500/go-laundry/internal/db"
	paymentmethodtype "github.com/umardev500/go-laundry/internal/domain/payment_method_type"
)

type repositoryImpl struct {
	client *db.Client
}

// Create implements paymentmethodtype.Repository.
func (r *repositoryImpl) Create(ctx context.Context, payload *paymentmethodtype.Create) (*paymentmethodtype.PaymentMethodType, error) {
	conn := r.client.GetConn(ctx)

	entObj, err := conn.PaymentMethodType.
		Create().
		SetNillableStatus((*paymentMethodTypeEntity.Status)(payload.Status)).
		SetName(payload.Name).
		SetDisplayName(payload.DisplayName).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(entObj), nil
}

// Delete implements paymentmethodtype.Repository.
func (r *repositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	_, err := conn.PaymentMethodType.
		Delete().
		Where(paymentMethodTypeEntity.IDEQ(id)).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetByID implements paymentmethodtype.Repository.
func (r *repositoryImpl) GetByID(ctx context.Context, id uuid.UUID, filter *paymentmethodtype.Filter) (*paymentmethodtype.PaymentMethodType, error) {
	conn := r.client.GetConn(ctx)

	q := conn.PaymentMethodType.
		Query().
		Where(paymentMethodTypeEntity.IDEQ(id))

	r.applyFilter(&q, filter)

	entObj, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(entObj), nil
}

// List implements paymentmethodtype.Repository.
func (r *repositoryImpl) List(ctx context.Context, filter *paymentmethodtype.Filter) ([]*paymentmethodtype.PaymentMethodType, error) {
	conn := r.client.GetConn(ctx)

	q := conn.PaymentMethodType.Query()

	r.applyFilter(&q, filter)

	entObjs, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnts(entObjs), nil
}

// Update implements paymentmethodtype.Repository.
func (r *repositoryImpl) Update(ctx context.Context, id uuid.UUID, payload *paymentmethodtype.Update) (*paymentmethodtype.PaymentMethodType, error) {
	conn := r.client.GetConn(ctx)

	builder := conn.PaymentMethodType.
		UpdateOneID(id).
		SetNillableStatus((*paymentMethodTypeEntity.Status)(payload.Status)).
		SetNillableName(payload.Name).
		SetNillableDisplayName(payload.DisplayName)

	entObj, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(entObj), nil
}

func (r *repositoryImpl) applyFilter(q **ent.PaymentMethodTypeQuery, filter *paymentmethodtype.Filter) {
	if filter == nil {
		return
	}

	if filter.Status != nil {
		*q = (*q).Where(paymentMethodTypeEntity.StatusEQ(paymentMethodTypeEntity.Status(*filter.Status)))
	}
}

func (r *repositoryImpl) mapFromEnts(es []*ent.PaymentMethodType) []*paymentmethodtype.PaymentMethodType {
	var result []*paymentmethodtype.PaymentMethodType
	for _, e := range es {
		result = append(result, r.mapFromEnt(e))
	}
	return result
}

func (r *repositoryImpl) mapFromEnt(e *ent.PaymentMethodType) *paymentmethodtype.PaymentMethodType {
	return &paymentmethodtype.PaymentMethodType{
		ID:          e.ID,
		Name:        *e.Name,
		DisplayName: *e.DisplayName,
		Status:      paymentmethodtype.Status(*e.Status),
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func NewRepository(client *db.Client) paymentmethodtype.Repository {
	return &repositoryImpl{client: client}
}
