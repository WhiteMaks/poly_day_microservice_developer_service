package microservice

import (
	"context"
	"fmt"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_logger/logger"
	"poly_day_microservice_developer_service/autogenerate/database"
)

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

	createEntityArgs := database.CreateEntityParams{
		UserID:         message.Meta.UserId,
		Name:           requestBody.Name,
		Type:           string(requestBody.Type),
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
