package microservice

import (
	"database/sql"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/WhiteMaks/poly_day_bus_schema/wbus"
	"github.com/WhiteMaks/poly_day_logger/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"poly_day_microservice_developer_service/autogenerate/database"
	mockdatabase "poly_day_microservice_developer_service/autogenerate/database/mock"
	"poly_day_microservice_developer_service/utils"
	"testing"
	"time"
)

const (
	rabbitConStr = "amqp://user:password@localhost:5672/"
)

var bus *wbus.Rabbit

func TestMain(m *testing.M) {

	bus = wbus.NewRabbit(schema.PathMsAuthenticate)
	err := bus.Start(rabbitConStr)
	if err != nil {
		log.Fatalf("error start bus [ %v ]: %v", bus.Name, err)
	}

	os.Exit(m.Run())
}

func newMs(store *mockdatabase.MockStore) *Microservice {
	return NewMs(store, bus, logger.NewPolyDayLogger(logger.NewDefaultLogger(logger.DEBUG), "Test"))
}

func generateRandomDBDeveloper() database.RootDeveloper {
	return database.RootDeveloper{
		UserID:    utils.RandomInt(0, 10000),
		CreatedAt: time.Now(),
	}
}

func prepareRequestMessage(userId int64) schema.Message {
	meta := schema.Meta{
		UserId: userId,
		WsId:   utils.RandomString(10),
	}

	return schema.Message{
		FromMicroservice: utils.RandomString(15),
		ToMicroservice:   utils.RandomString(15),
		RequestId:        utils.RandomString(15),
		Method:           utils.RandomString(15),
		Version:          utils.RandomString(15),
		Meta:             meta,
	}
}

func testMicroserviceUserNotExistMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(database.RootDeveloper{}, sql.ErrNoRows)
}

func checkResponseWithoutBody(t *testing.T, requestMessage schema.Message, responseMessage schema.Message) {
	assert.NotNil(t, requestMessage, "Response message must be")
	assert.Equal(t, requestMessage.ToMicroservice, responseMessage.FromMicroservice, "From microservice")
	assert.Equal(t, requestMessage.FromMicroservice, responseMessage.ToMicroservice, "To microservice")
	assert.Equal(t, requestMessage.Method, responseMessage.Method, "Method")
	assert.Equal(t, requestMessage.Version, responseMessage.Version, "Version")
	assert.Equal(t, requestMessage.RequestId, responseMessage.ResponseId, "Response id")
	assert.NotNil(t, responseMessage.Status, "Status must be")
	assert.NotEmpty(t, responseMessage.Status, "Status must be not empty")
}

func checkSuccessResponseWithBodyNil(t *testing.T, actualMessage schema.Message) {
	assert.Equal(t, schema.OK, actualMessage.Status, "Status")
	assert.Nil(t, actualMessage.Body, "Body must be nil")
}

func checkSuccessResponse(t *testing.T, expectedMessage schema.Message, actualMessage schema.Message) {
	assert.Equal(t, schema.OK, actualMessage.Status, "Status")
	assert.NotNil(t, actualMessage.Body, "Body must not be nil")
	assert.Equal(t, expectedMessage.Body, actualMessage.Body, "Body")
}

func checkErrorResponse(t *testing.T, statusCode schema.Status, responseMessage schema.Message) {
	assert.Equal(t, statusCode, responseMessage.Status, "Status")
	assert.Nil(t, responseMessage.Body, "Body must be nil")
}
