package microservice

import (
	"context"
	"fmt"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_logger/logger"
	"poly_day_microservice_developer_service/autogenerate/database"
)

func (ms *Microservice) getEntities10(message schema.Message) {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of receiving entities has begun",
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	_, err := ms.store.GetDeveloperById(context.Background(), message.Meta.UserId)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "the user was denied permission to receive entities",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.NOT_AUTHORIZED)
		ms.bus.SendMessage(responseMessage)
		return
	}

	dbEntities, err := ms.store.GetEntitiesByParentEntityId(context.Background(), 0)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when searching for entities in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)
		ms.bus.SendMessage(responseMessage)
		return
	}

	entitiesBody := make([]schema.Entity[any], len(dbEntities))

	for i, value := range dbEntities {
		entityBody := schema.GenerateEntityStruct[any](value.ID, value.Name, schema.EntityType(value.Type), nil)

		dbEntityChildList, err := ms.store.GetEntitiesByParentEntityId(context.Background(), entityBody.ID)

		if err != nil {
			ms.log.Error(
				logger.Message{
					ShortMessage: fmt.Sprintf("error when searching for entities in the database by parent entity id [ %v ]", entityBody.ID),
					FullMessage:  err.Error(),
					UserId:       message.Meta.UserId,
					WsId:         message.Meta.WsId,
				},
			)
			continue
		}

		err = ms.generateEntitiesForParent(&entityBody, dbEntityChildList)
		if err != nil {
			ms.log.Error(
				logger.Message{
					ShortMessage: fmt.Sprintf("error when generating child list for entity id [ %v ]", entityBody.ID),
					FullMessage:  err.Error(),
					UserId:       message.Meta.UserId,
					WsId:         message.Meta.WsId,
				},
			)
			continue
		}

		entitiesBody[i] = entityBody
	}

	err = schema.FillBodyGetEntities10ResponseForMessage(entitiesBody, &responseMessage)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when filling out the body of the response to a request for entities",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)
		ms.bus.SendMessage(responseMessage)
		return
	}

	ms.log.Info(
		logger.Message{
			ShortMessage: fmt.Sprintf("the entity receive process is complete. Sent entities [ %v ]", len(dbEntities)),
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)
	ms.bus.SendMessage(responseMessage)
}

func (ms *Microservice) generateEntitiesForParent(parentEntity *schema.Entity[any], dbEntities []database.RootEntity) error {
	entitiesBody := make([]schema.Entity[any], len(dbEntities))

	for i, value := range dbEntities {
		entityBody := schema.GenerateEntityStruct[any](value.ID, value.Name, schema.EntityType(value.Type), nil)

		dbEntitiesChild, err := ms.store.GetEntitiesByParentEntityId(context.Background(), entityBody.ID)
		if err != nil {
			return err
		}

		err = ms.generateEntitiesForParent(&entityBody, dbEntitiesChild)
		if err != nil {
			return err
		}

		entitiesBody[i] = entityBody
	}

	parentEntity.Entities = entitiesBody

	return nil
}

func (ms *Microservice) createEntity10(message schema.Message) {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of creating a entity has begun",
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	_, err := ms.store.GetDeveloperById(context.Background(), message.Meta.UserId)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "the user was denied permission to create an entity",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.NOT_AUTHORIZED)
		ms.bus.SendMessage(responseMessage)
		return
	}

	requestBody, err := schema.GetCreateEntity10RequestFromMessage(message)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid request body for create entity",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)
		ms.bus.SendMessage(responseMessage)
		return
	}

	entityType := requestBody.Type.String()
	if entityType == nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid entity type in body for create entity",
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)
		ms.bus.SendMessage(responseMessage)
		return
	}

	if requestBody.ParentEntityId > 0 {
		dbParentEntity, err := ms.store.GetEntityById(context.Background(), requestBody.ParentEntityId)
		if err != nil {
			ms.log.Warn(
				logger.Message{
					ShortMessage: fmt.Sprintf("parent entity id [ %v ] in body for create entity not found in databse", requestBody.ParentEntityId),
					FullMessage:  err.Error(),
					UserId:       message.Meta.UserId,
					WsId:         message.Meta.WsId,
				},
			)

			schema.FillStatusForMessage(&responseMessage, schema.NOT_FOUND)
			ms.bus.SendMessage(responseMessage)
			return
		}

		if schema.EntityType(dbParentEntity.Type) != schema.FOLDER {
			ms.log.Warn(
				logger.Message{
					ShortMessage: fmt.Sprintf("parent entity type is [ %v ] not equal [ %v ]", requestBody.Type, string(schema.FOLDER)),
					UserId:       message.Meta.UserId,
					WsId:         message.Meta.WsId,
				},
			)

			schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)
			ms.bus.SendMessage(responseMessage)
			return
		}
	}

	createEntityArgs := database.CreateEntityParams{
		UserID:         message.Meta.UserId,
		Name:           requestBody.Name,
		Type:           *entityType,
		ParentEntityID: requestBody.ParentEntityId,
	}

	dbEntity, err := ms.store.CreateEntity(context.Background(), createEntityArgs)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when creating an entity in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)
		ms.bus.SendMessage(responseMessage)
		return
	}

	err = schema.FillBodyCreateEntity10ResponseForMessage(dbEntity.ID, &responseMessage)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when filling out the body of the response to a request to create an entity",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)
		ms.bus.SendMessage(responseMessage)
		return
	}

	ms.log.Info(
		logger.Message{
			ShortMessage: fmt.Sprintf("the entity creation process is complete. Added entity [ %v ]", dbEntity.ID),
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)
	ms.bus.SendMessage(responseMessage)
}
