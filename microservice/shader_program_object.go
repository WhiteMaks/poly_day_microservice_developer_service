package microservice

import (
	"context"
	"fmt"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_logger/logger"
	"poly_day_microservice_developer_service/autogenerate/database"
)

func (ms *Microservice) createShaderProgramObject10(message schema.Message) schema.Message {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of creating a shader program object has begun",
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	err := ms.checkUser("the user was denied permission to create a shader program object", message.Meta, &responseMessage)
	if err != nil {
		return responseMessage
	}

	requestBody, err := schema.GetCreateShaderProgramObject10RequestFromMessage(message)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid request body for create shader program object",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)

		return responseMessage
	}

	dbShaderProgram, err := ms.store.GetShaderProgramById(context.Background(), requestBody.ShaderProgramId)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "error when searching shader program in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.NOT_FOUND)

		return responseMessage
	}

	dbEntity, err := ms.store.GetEntityById(context.Background(), requestBody.ModelEntityId)
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

	if !ms.dbEntityTypeIs(dbEntity, schema.MODEL, message.Meta, &responseMessage) {
		return responseMessage
	}

	createShaderProgramObjectArgs := database.CreateShaderProgramObjectParams{
		ShaderProgramID: dbShaderProgram.ID,
		UserID:          message.Meta.UserId,
		ModelEntityID:   requestBody.ModelEntityId,
	}

	dbShaderProgramObject, err := ms.store.CreateShaderProgramObject(context.Background(), createShaderProgramObjectArgs)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when creating a shader program object in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)

		return responseMessage
	}

	err = schema.FillBodyCreateShaderProgramObject10ResponseForMessage(dbShaderProgramObject.ID, &responseMessage)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when filling out the body of the response to a request to create a shader program object",
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
			ShortMessage: fmt.Sprintf("the shader program object creation process is complete. Added shader program object [ %v ]", dbShaderProgramObject.ID),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)

	return responseMessage
}
