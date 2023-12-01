package microservice

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_logger/logger"
	"poly_day_microservice_developer_service/autogenerate/database"
)

func (ms *Microservice) getModelSource10(message schema.Message) schema.Message {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of receiving model source has begun",
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	err := ms.checkUser("the user was denied permission to receive model source", message.Meta, &responseMessage)
	if err != nil {
		return responseMessage
	}

	requestBody, err := schema.GetGetModelSource10RequestFromMessage(message)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid request body for get model source",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)

		return responseMessage
	}

	dbModelSource, err := ms.store.GetModelSourceByEntityId(context.Background(), requestBody.EntityId)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when searching for model source in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)

		return responseMessage
	}

	err = schema.FillBodyGetModelSource10ResponseForMessage(dbModelSource.ID, dbModelSource.SourceBase64.String, &responseMessage)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when filling out the body of the response to a request for model source",
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
			ShortMessage: fmt.Sprintf("the model source receive process is complete. Sent model source [ %v ]", dbModelSource.ID),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)

	return responseMessage
}

func (ms *Microservice) saveModelSource10(message schema.Message) schema.Message {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of saving model source has begun",
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	err := ms.checkUser("the user was denied permission to saving model source", message.Meta, &responseMessage)
	if err != nil {
		return responseMessage
	}

	requestBody, err := schema.GetSaveModelSource10RequestFromMessage(message)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid request body for save model source",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)

		return responseMessage
	}

	dbEntity, err := ms.store.GetEntityById(context.Background(), requestBody.EntityId)
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

	if schema.EntityType(dbEntity.Type) != schema.MODEL {
		ms.log.Warn(
			logger.Message{
				ShortMessage: fmt.Sprintf("entity by id [ %v ] is not a model", dbEntity.ID),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)

		return responseMessage
	}

	var sourceBase64 sql.NullString

	if len(requestBody.SourceBase64) > 0 {
		sourceBase64 = sql.NullString{
			String: requestBody.SourceBase64,
			Valid:  true,
		}
	} else {
		sourceBase64 = sql.NullString{}
	}

	updateModelSourceByEntityIdArgs := database.UpdateModelSourceByEntityIdParams{
		EntityID:     requestBody.EntityId,
		SourceBase64: sourceBase64,
	}

	err = ms.store.UpdateModelSourceByEntityId(context.Background(), updateModelSourceByEntityIdArgs)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when updating a model source in the database",
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
			ShortMessage: fmt.Sprintf("the model source saving process is complete. Updated model source by entity id [ %v ]", requestBody.EntityId),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)

	return responseMessage
}

func (ms *Microservice) getShaderSource10(message schema.Message) {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of receiving shader source has begun",
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
				ShortMessage: "the user was denied permission to receive shader source",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.NOT_AUTHORIZED)
		ms.bus.SendMessage(responseMessage)
		return
	}

	requestBody, err := schema.GetGetShaderSource10RequestFromMessage(message)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid request body for get shader source",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)
		ms.bus.SendMessage(responseMessage)
		return
	}

	dbShaderSource, err := ms.store.GetShaderSourceByEntityId(context.Background(), requestBody.EntityId)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when searching for shader source in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)
		ms.bus.SendMessage(responseMessage)
		return
	}

	err = schema.FillBodyGetShaderSource10ResponseForMessage(dbShaderSource.ID, dbShaderSource.VertexShader.String, dbShaderSource.FragmentShader.String, &responseMessage)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when filling out the body of the response to a request for shader source",
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
			ShortMessage: fmt.Sprintf("the shader source receive process is complete. Sent shader source [ %v ]", dbShaderSource.ID),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)
	ms.bus.SendMessage(responseMessage)
}

func (ms *Microservice) saveShaderSource10(message schema.Message) {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of saving shader source has begun",
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	_, err := ms.store.GetDeveloperById(context.Background(), message.Meta.UserId)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "the user was denied permission to saving shader source",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.NOT_AUTHORIZED)
		ms.bus.SendMessage(responseMessage)
		return
	}

	requestBody, err := schema.GetSaveShaderSource10RequestFromMessage(message)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid request body for save shader source",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)
		ms.bus.SendMessage(responseMessage)
		return
	}

	dbEntity, err := ms.store.GetEntityById(context.Background(), requestBody.EntityId)
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
		ms.bus.SendMessage(responseMessage)
		return
	}

	if schema.EntityType(dbEntity.Type) != schema.SHADER {
		ms.log.Warn(
			logger.Message{
				ShortMessage: fmt.Sprintf("entity by id [ %v ] is not a shader", dbEntity.ID),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)
		ms.bus.SendMessage(responseMessage)
		return
	}

	var vertexShader sql.NullString
	var geometryShader sql.NullString
	var fragmentShader sql.NullString

	if len(requestBody.VertexShader) > 0 {
		vertexShader = sql.NullString{
			String: requestBody.VertexShader,
			Valid:  true,
		}
	} else {
		vertexShader = sql.NullString{}
	}

	geometryShader = sql.NullString{}

	if len(requestBody.FragmentShader) > 0 {
		fragmentShader = sql.NullString{
			String: requestBody.FragmentShader,
			Valid:  true,
		}
	} else {
		fragmentShader = sql.NullString{}
	}

	updateShaderSourceByEntityIdArgs := database.UpdateShaderSourceByEntityIdParams{
		EntityID:       requestBody.EntityId,
		VertexShader:   vertexShader,
		GeometryShader: geometryShader,
		FragmentShader: fragmentShader,
	}

	err = ms.store.UpdateShaderSourceByEntityId(context.Background(), updateShaderSourceByEntityIdArgs)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when updating a shader source in the database",
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
			ShortMessage: fmt.Sprintf("the shader source saving process is complete. Updated shader source by entity id [ %v ]", requestBody.EntityId),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)
	ms.bus.SendMessage(responseMessage)
}
