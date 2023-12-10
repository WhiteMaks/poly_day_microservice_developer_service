package microservice

import (
	"context"
	"fmt"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_logger/logger"
	"poly_day_microservice_developer_service/autogenerate/database"
)

func (ms *Microservice) getReleases10(message schema.Message) schema.Message {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of receiving releases has begun",
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	err := ms.checkUser("the user was denied permission to receive releases", message.Meta, &responseMessage)
	if err != nil {
		return responseMessage
	}

	dbReleases, err := ms.store.GetReleases(context.Background())
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when searching for releases in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)

		return responseMessage
	}

	releasesBody := make([]schema.Release, len(dbReleases))

	for i, value := range dbReleases {
		releasesBody[i] = schema.GenerateReleaseStruct(value.ID, value.Name)
	}

	err = schema.FillBodyGetReleases10ResponseForMessage(releasesBody, &responseMessage)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when filling out the body of the response to a request for releases",
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
			ShortMessage: fmt.Sprintf("the release receive process is complete. Sent releases [ %v ]", len(dbReleases)),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)

	return responseMessage
}

func (ms *Microservice) createRelease10(message schema.Message) schema.Message {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of creating a release has begun",
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	err := ms.checkUser("the user was denied permission to create a release", message.Meta, &responseMessage)
	if err != nil {
		return responseMessage
	}

	requestBody, err := schema.GetCreateRelease10RequestFromMessage(message)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid request body for create release",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)

		return responseMessage
	}

	createReleaseArgs := database.CreateReleaseParams{
		UserID: message.Meta.UserId,
		Name:   requestBody.Name,
	}

	dbRelease, err := ms.store.CreateRelease(context.Background(), createReleaseArgs)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when creating a release in the database",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.INTERNAL_ERROR)

		return responseMessage
	}

	err = schema.FillBodyCreateRelease10ResponseForMessage(dbRelease.ID, &responseMessage)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when filling out the body of the response to a request to create a release",
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
			ShortMessage: fmt.Sprintf("the release creation process is complete. Added release [ %v ]", dbRelease.ID),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)

	return responseMessage
}

func (ms *Microservice) deleteRelease10(message schema.Message) schema.Message {
	ms.log.Info(
		logger.Message{
			ShortMessage: "the process of deleting a release has begun",
			FullMessage:  string(message.Body),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	responseMessage := schema.GenerateResponseMessage(message)

	err := ms.checkUser("the user was denied permission to delete a release", message.Meta, &responseMessage)
	if err != nil {
		return responseMessage
	}

	requestBody, err := schema.GetDeleteRelease10RequestFromMessage(message)
	if err != nil {
		ms.log.Warn(
			logger.Message{
				ShortMessage: "invalid request body for delete release",
				FullMessage:  err.Error(),
				UserId:       message.Meta.UserId,
				WsId:         message.Meta.WsId,
			},
		)

		schema.FillStatusForMessage(&responseMessage, schema.BAD_REQUEST)

		return responseMessage
	}

	err = ms.store.DeleteReleaseById(context.Background(), requestBody.ID)
	if err != nil {
		ms.log.Error(
			logger.Message{
				ShortMessage: "error when deleting a release in the database",
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
			ShortMessage: fmt.Sprintf("the release deletion process is complete. Deleted release [ %v ]", requestBody.ID),
			UserId:       message.Meta.UserId,
			WsId:         message.Meta.WsId,
		},
	)

	schema.FillStatusForMessage(&responseMessage, schema.OK)

	return responseMessage
}
