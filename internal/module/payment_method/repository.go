package paymentmethod

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	paymentMethodEntity "github.com/umardev500/go-laundry/ent/paymentmethod"
	"github.com/umardev500/go-laundry/ent/tenant"
	"github.com/umardev500/go-laundry/internal/db"
	paymentmethod "github.com/umardev500/go-laundry/internal/domain/payment_method"
	paymentmethodtype "github.com/umardev500/go-laundry/internal/domain/payment_method_type"
)

type repositoryImpl struct {
	db *db.Client
}

// Create implements paymentmethod.Repository.
func (r *repositoryImpl) Create(ctx context.Context, payload *paymentmethod.Create) (*paymentmethod.PaymentMethod, error) {
	conn := r.db.GetConn(ctx)

	entobj, err := conn.PaymentMethod.
		Create().
		SetTenantID(payload.TenantID).
		SetPaymentMethodTypeID(payload.TypeID).
		SetMetadata(payload.Metadata).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(entobj), nil
}

// Delete implements paymentmethod.Repository.
func (r *repositoryImpl) Delete(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) error {
	conn := r.db.GetConn(ctx)

	q := conn.PaymentMethod.DeleteOneID(id)

	// Tenant scope check
	if tenantID != nil {
		q = q.Where(paymentMethodEntity.HasTenantWith(tenant.IDEQ(*tenantID)))
	}

	if err := q.Exec(ctx); err != nil {
		return err
	}

	return nil
}

// GetByID implements paymentmethod.Repository.
func (r *repositoryImpl) GetByID(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, filter *paymentmethod.Filter) (*paymentmethod.PaymentMethod, error) {
	conn := r.db.GetConn(ctx)

	q := conn.PaymentMethod.Query().Where(paymentMethodEntity.IDEQ(id))

	r.applyFilter(&q, tenantID, filter)

	entObj, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(entObj), nil
}

// List implements paymentmethod.Repository.
func (r *repositoryImpl) List(ctx context.Context, tenantID *uuid.UUID, filter *paymentmethod.Filter) ([]*paymentmethod.PaymentMethod, error) {
	conn := r.db.GetConn(ctx)

	q := conn.PaymentMethod.Query()

	r.applyFilter(&q, tenantID, filter)

	// TODO: apply filter when extended
	entObjs, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnts(entObjs), nil
}

// Update implements paymentmethod.Repository.
func (r *repositoryImpl) Update(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, payload *paymentmethod.Update) (*paymentmethod.PaymentMethod, error) {
	conn := r.db.GetConn(ctx)

	builder := conn.PaymentMethod.
		UpdateOneID(id)

	if payload.Metadata != nil {
		builder = builder.SetMetadata(*payload.Metadata)
	}

	// Tenant scope check
	if tenantID != nil {
		builder = builder.Where(paymentMethodEntity.HasTenantWith(tenant.IDEQ(*tenantID)))
	}

	entObj, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(entObj), nil
}

func (r *repositoryImpl) applyFilter(q **ent.PaymentMethodQuery, tenantID *uuid.UUID, filter *paymentmethod.Filter) {
	if tenantID != nil {
		*q = (*q).Where(paymentMethodEntity.HasTenantWith(tenant.IDEQ(*tenantID)))
	}

	if filter == nil {
		return
	}

	if filter.IncludeType {
		*q = (*q).WithPaymentMethodType()
	}

}

func (r *repositoryImpl) mapFromEnts(ents []*ent.PaymentMethod) []*paymentmethod.PaymentMethod {
	var paymentMethods []*paymentmethod.PaymentMethod
	for _, entobj := range ents {
		paymentMethods = append(paymentMethods, r.mapFromEnt(entobj))
	}
	return paymentMethods
}

func (r *repositoryImpl) mapFromEnt(entobj *ent.PaymentMethod) *paymentmethod.PaymentMethod {
	var methodType *paymentmethodtype.PaymentMethodType

	entMethodType := entobj.Edges.PaymentMethodType
	if entMethodType != nil {
		methodType = &paymentmethodtype.PaymentMethodType{
			ID:          entMethodType.ID,
			Name:        *entMethodType.Name,
			DisplayName: *entMethodType.DisplayName,
			Status:      paymentmethodtype.Status(*entMethodType.Status),
			CreatedAt:   entMethodType.CreatedAt,
			UpdatedAt:   entMethodType.UpdatedAt,
		}
	}

	return &paymentmethod.PaymentMethod{
		ID:        entobj.ID,
		TenantID:  entobj.TenantID,
		TypeID:    *entobj.PaymentMethodTypeID,
		Type:      methodType,
		Metadata:  entobj.Metadata,
		CreatedAt: entobj.CreatedAt,
		UpdatedAt: entobj.UpdatedAt,
	}
}

func NewRepository(client *db.Client) paymentmethod.Repository {
	return &repositoryImpl{
		db: client,
	}
}
