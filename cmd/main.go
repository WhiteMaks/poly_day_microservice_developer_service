package main

import (
	"database/sql"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_bus_schema/wbus"
	"github.com/WhiteMaks/poly_day_logger/logger"
	"os"
	"poly_day_microservice_developer_service/autogenerate/database"
	"poly_day_microservice_developer_service/microservice"
)

var (
	psqlConStr   = os.Getenv("POSTGRES_CONNECTING_STRING")
	rabbitConStr = os.Getenv("RABBIT_CONNECTING_STRING")
	gelfConStr   = os.Getenv("GELF_CONNECTING_STRING")
)

func main() {
	msName := schema.PathMsAuthenticate

	logDefault := logger.NewDefaultLogger(logger.DEBUG)
	logGelf := logger.NewGelf(logDefault, gelfConStr)
	log := logger.NewPolyDayLogger(logGelf, msName)

	bus := wbus.NewRabbit(msName)
	err := bus.Start(rabbitConStr)
	if err != nil {
		log.Error(
			logger.Message{
				ShortMessage: "error connecting to bus",
				FullMessage:  err.Error(),
			},
		)
		return
	}

	connection, err := sql.Open("postgres", psqlConStr)
	if err != nil {
		log.Error(
			logger.Message{
				ShortMessage: "error connecting to database",
				FullMessage:  err.Error(),
			},
		)
		return
	}

	store := database.NewStore(connection)

	ms := microservice.NewMs(store, bus, log)
	err = ms.Start()
	if err != nil {
		log.Error(
			logger.Message{
				ShortMessage: "microservice startup error",
				FullMessage:  err.Error(),
			},
		)
		return
	}
}
