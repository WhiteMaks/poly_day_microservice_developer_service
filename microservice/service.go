package microservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_bus_schema/wbus"
	"github.com/WhiteMaks/poly_day_logger/logger"
	"poly_day_microservice_developer_service/autogenerate/database"
	"poly_day_microservice_developer_service/utils"
)

type Requests map[string]chan schema.Message

type Microservice struct {
	requests Requests
	store    database.Store
	bus      *wbus.Rabbit
	log      *logger.PolyDayLogger
}

func NewMs(store database.Store, bus *wbus.Rabbit, log *logger.PolyDayLogger) *Microservice {
	return &Microservice{
		requests: make(Requests),
		store:    store,
		bus:      bus,
		log:      log,
	}
}

func (ms *Microservice) Start() error {
	err := ms.bus.SubscribeOnQueue(ms.bus.Name, ms.rootMessageHandler)
	if err != nil {
		return err
	}

	ms.log.Info(
		logger.Message{
			ShortMessage: "the developer synchronization process has begun",
		},
	)

	dbDeveloper, err := ms.store.GetMaxUserIdFromDevelopers(context.Background())
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}

	developers10Response, err := ms.getDevelopers10(dbDeveloper.UserID)
	if err != nil {
		return err
	}

	developers := developers10Response.Developers
	for _, developer := range developers {
		_, err = ms.store.CreateDeveloper(context.Background(), developer.UserId)
		if err != nil {
			ms.log.Warn(
				logger.Message{
					ShortMessage: fmt.Sprintf("error when saving developer [ %v ] to database", developer.UserId),
					FullMessage:  err.Error(),
				},
			)
		}
	}

	ms.log.Info(
		logger.Message{
			ShortMessage: fmt.Sprintf("the developer synchronization process is complete. Added [ %v ] new entries", len(developers)),
		},
	)

	return nil
}

func (ms *Microservice) getDevelopers10(fromUserId int64) (*schema.GetDevelopers10Response, error) {
	requestBodyStruct := schema.GetDevelopers10Request{
		FromUserId: fromUserId,
	}

	requestBodyBates, err := json.Marshal(requestBodyStruct)
	if err != nil {
		return nil, err
	}

	requestMsg := schema.Message{
		FromMicroservice: ms.bus.Name,
		ToMicroservice:   schema.PathMsAuthenticate,
		Method:           schema.CommandGetDevelopers,
		Version:          schema.Version10,
		Body:             requestBodyBates,
		RequestId:        utils.RandomString(15),
	}

	ms.requests[requestMsg.RequestId] = make(chan schema.Message)

	ms.bus.SendMessage(requestMsg)

	responseMsg := <-ms.requests[requestMsg.RequestId]

	delete(ms.requests, responseMsg.ResponseId)

	var getDevelopers10Response schema.GetDevelopers10Response

	err = json.Unmarshal(responseMsg.Body, &getDevelopers10Response)
	if err != nil {
		return nil, err
	}

	return &getDevelopers10Response, nil
}

func (ms *Microservice) rootMessageHandler(message schema.Message) {
	if message.ResponseId != "" {
		ms.requests[message.ResponseId] <- message
		return
	}

	methodName := schema.GenerateFullMethodNameFromMessage(message)

	ms.log.Info(
		logger.Message{
			ShortMessage: fmt.Sprintf("received method name [ %v ] for further processing", methodName),
		},
	)

	switch methodName {

	case schema.CommandCreateRelease10:
		{
			go ms.createRelease10(message)
			break
		}

	case schema.CommandGetReleases10:
		{
			go ms.getReleases10(message)
			break
		}

	}
}
