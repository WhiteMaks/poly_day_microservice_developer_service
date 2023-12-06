package microservice

import (
	"context"
	"fmt"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_logger/logger"
	"poly_day_microservice_developer_service/autogenerate/database"
)

func (ms *Microservice) createShaderProgram10(message schema.Message) schema.Message {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of creating a shader program has begun",
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	err := ms.checkUser("the user was denied permission to create a shader program", message.Meta, &responseMessage)
	if err != nil {
		return responseMessage
	}

	requestBody, err := schema.GetCreateShaderProgram10RequestFromMessage(message)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid request body for create shader program",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)

		return responseMessage
	}

	dbRelease, err := ms.store.GetReleaseById(context.Background(), requestBody.ReleaseId)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "error when searching release in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.NOT_FOUND)

		return responseMessage
	}

	dbEntity, err := ms.store.GetEntityById(context.Background(), requestBody.ShaderEntityId)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "error when searching entity in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.NOT_FOUND)

		return responseMessage
	}

	if !ms.dbEntityTypeIs(dbEntity, schema.SHADER, message.Meta, &responseMessage) {
		return responseMessage
	}

	createShaderProgramArgs := database.CreateShaderProgramParams{
		Name:           requestBody.Name,
		ReleaseID:      dbRelease.ID,
		UserID:         message.Meta.UserId,
		ShaderEntityID: requestBody.ShaderEntityId,
	}

	dbShaderProgram, err := ms.store.CreateShaderProgram(context.Background(), createShaderProgramArgs)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when creating a shader program in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)

		return responseMessage
	}

	err = schema.FillBodyCreateShaderProgram10ResponseForMessage(dbShaderProgram.ID, &responseMessage)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when filling out the body of the response to a request to create a shader program",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)

		return responseMessage
	}

	ms.log.Info(
		logger.Message{
			ShortMessage: fmt.Sprintf("the shader program creation process is complete. Added shader program [ %v ]", dbEntity.ID),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)

	return responseMessage
}
