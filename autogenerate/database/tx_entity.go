package database

import (
	"context"
	"database/sql"
)

func (store *SQLStore) CreateModelEntity(ctx context.Context, createEntityArgs CreateEntityParams) (RootEntity, error) {
	var result RootEntity

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result, err = q.CreateEntity(ctx, createEntityArgs)
		if err != nil {
			return err
		}

		createModelSourceArgs := CreateModelSourceParams{
			result.ID,
			sql.NullString{},
		}

		_, err = q.CreateModelSource(ctx, createModelSourceArgs)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

func (store *SQLStore) CreateShaderEntity(ctx context.Context, createEntityArgs CreateEntityParams) (RootEntity, error) {
	var result RootEntity

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result, err = q.CreateEntity(ctx, createEntityArgs)
		if err != nil {
			return err
		}

		createShaderSourceArgs := CreateShaderSourceParams{
			result.ID,
			sql.NullString{},
			sql.NullString{},
			sql.NullString{},
		}

		_, err = q.CreateShaderSource(ctx, createShaderSourceArgs)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

func (store *SQLStore) DeleteModelEntity(ctx context.Context, entityId int64) error {
	err := store.execTx(ctx, func(q *Queries) error {
		err := q.DeleteModelSourceByEntityId(ctx, entityId)
		if err != nil {
			return err
		}

		err = q.DeleteEntityById(ctx, entityId)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (store *SQLStore) DeleteShaderEntity(ctx context.Context, entityId int64) error {
	err := store.execTx(ctx, func(q *Queries) error {
		err := q.DeleteShaderSourceByEntityId(ctx, entityId)
		if err != nil {
			return err
		}

		err = q.DeleteEntityById(ctx, entityId)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
