package audit

import (
	"context"

	"github.com/umardev500/go-laundry/ent/auditlog"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/audit"
)

type repositoryImpl struct {
	client *db.Client
}

func NewRepository(client *db.Client) audit.Repository {
	return &repositoryImpl{
		client: client,
	}
}

func (r *repositoryImpl) Create(ctx context.Context, payload *audit.Create) error {
	conn := r.client.GetConn(ctx)

	_, err := conn.AuditLog.
		Create().
		SetTableName(auditlog.TableName(payload.TableName)).
		SetRecordID(payload.RecordID).
		SetAction(auditlog.Action(payload.Action)).
		SetModifiedBy(payload.ModifiedBy).
		SetOldData(payload.OldData).
		SetNewData(payload.NewData).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
