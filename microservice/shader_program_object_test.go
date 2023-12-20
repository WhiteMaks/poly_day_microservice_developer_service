package microservice

import (
	"encoding/json"
	"errors"
	"github.com/WhiteMaks/poly_day_bus_schema/schema"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"poly_day_microservice_developer_service/autogenerate/database"
	mockdatabase "poly_day_microservice_developer_service/autogenerate/database/mock"
	"poly_day_microservice_developer_service/utils"
	"testing"
	"time"
)

func TestMicroservice_CreateShaderProgramObject10(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbModelEntity := generateRandomDBModelEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)
	dbShaderProgramObject := generateRandomDBShaderProgramObject(dbShaderProgram.ID, dbModelEntity.ID)

	testMicroserviceCreateShaderProgramObjectMockDB(store, dbDeveloper, dbShaderProgram, dbModelEntity, dbShaderProgramObject)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgramObject10Request{
		ShaderProgramId: dbShaderProgramObject.ShaderProgramID,
		ModelEntityId:   dbShaderProgramObject.ModelEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgramObject10(requestMessage)
	expectedMessage := schema.Message{}
	err = schema.FillBodyCreateShaderProgramObject10ResponseForMessage(dbShaderProgramObject.ID, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_CreateShaderProgramObject10_NOT_AUTHORIZED(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbModelEntity := generateRandomDBModelEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)
	dbShaderProgramObject := generateRandomDBShaderProgramObject(dbShaderProgram.ID, dbModelEntity.ID)

	testMicroserviceUserNotExistMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgramObject10Request{
		ShaderProgramId: dbShaderProgramObject.ShaderProgramID,
		ModelEntityId:   dbShaderProgramObject.ModelEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgramObject10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_AUTHORIZED, actualMessage)
}

func TestMicroservice_CreateShaderProgramObject10_NOT_FOUND_ShaderProgramNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbModelEntity := generateRandomDBModelEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)
	dbShaderProgramObject := generateRandomDBShaderProgramObject(dbShaderProgram.ID, dbModelEntity.ID)

	testMicroserviceCreateShaderProgramObjectErrorShaderProgramMockDB(store, dbDeveloper, dbShaderProgram)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgramObject10Request{
		ShaderProgramId: dbShaderProgramObject.ShaderProgramID,
		ModelEntityId:   dbShaderProgramObject.ModelEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgramObject10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_FOUND, actualMessage)
}

func TestMicroservice_CreateShaderProgramObject10_NOT_FOUND_EntityNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbModelEntity := generateRandomDBModelEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)
	dbShaderProgramObject := generateRandomDBShaderProgramObject(dbShaderProgram.ID, dbModelEntity.ID)

	testMicroserviceCreateShaderProgramObjectErrorEntityMockDB(store, dbDeveloper, dbShaderProgram, dbModelEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgramObject10Request{
		ShaderProgramId: dbShaderProgramObject.ShaderProgramID,
		ModelEntityId:   dbShaderProgramObject.ModelEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgramObject10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_FOUND, actualMessage)
}

func TestMicroservice_CreateShaderProgramObject10_BAD_REQUEST_EntityTypeNotModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbModelEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)
	dbShaderProgramObject := generateRandomDBShaderProgramObject(dbShaderProgram.ID, dbModelEntity.ID)

	testMicroserviceCreateShaderProgramObjectErrorModelTypeMockDB(store, dbDeveloper, dbShaderProgram, dbModelEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgramObject10Request{
		ShaderProgramId: dbShaderProgramObject.ShaderProgramID,
		ModelEntityId:   dbShaderProgramObject.ModelEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgramObject10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.BAD_REQUEST, actualMessage)
}

func TestMicroservice_CreateShaderProgramObject10_INTERNAL_ERROR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbModelEntity := generateRandomDBModelEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)
	dbShaderProgramObject := generateRandomDBShaderProgramObject(dbShaderProgram.ID, dbModelEntity.ID)

	testMicroserviceCreateShaderProgramObjectErrorMockDB(store, dbDeveloper, dbShaderProgram, dbModelEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgramObject10Request{
		ShaderProgramId: dbShaderProgramObject.ShaderProgramID,
		ModelEntityId:   dbShaderProgramObject.ModelEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgramObject10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.INTERNAL_ERROR, actualMessage)
}

func generateRandomDBShaderProgramObject(shaderProgramId int64, modelEntityId int64) database.RootShaderProgramObject {
	randomDuration := time.Duration(utils.RandomInt(0, 10000))
	randomCreateAt := time.Now().
		Add(-randomDuration)

	return database.RootShaderProgramObject{
		ID:              utils.RandomInt(0, 100000),
		UserID:          utils.RandomInt(0, 100000),
		CreatedAt:       randomCreateAt,
		UpdatedAt:       time.Now(),
		ShaderProgramID: shaderProgramId,
		ModelEntityID:   modelEntityId,
	}
}

func testMicroserviceCreateShaderProgramObjectMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbShaderProgram database.RootShaderProgram, dbModelEntity database.RootEntity, dbShaderProgramObject database.RootShaderProgramObject) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetShaderProgramById(
			gomock.Any(),
			gomock.Eq(dbShaderProgram.ID),
		).Times(1).
		Return(dbShaderProgram, nil)

	store.EXPECT().
		GetEntityById(
			gomock.Any(),
			gomock.Eq(dbModelEntity.ID),
		).Times(1).
		Return(dbModelEntity, nil)

	store.EXPECT().
		CreateShaderProgramObject(
			gomock.Any(),
			gomock.Any(),
		).Times(1).
		Return(dbShaderProgramObject, nil)
}

func testMicroserviceCreateShaderProgramObjectErrorMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbShaderProgram database.RootShaderProgram, dbModelEntity database.RootEntity) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetShaderProgramById(
			gomock.Any(),
			gomock.Eq(dbShaderProgram.ID),
		).Times(1).
		Return(dbShaderProgram, nil)

	store.EXPECT().
		GetEntityById(
			gomock.Any(),
			gomock.Eq(dbModelEntity.ID),
		).Times(1).
		Return(dbModelEntity, nil)

	store.EXPECT().
		CreateShaderProgramObject(
			gomock.Any(),
			gomock.Any(),
		).Times(1).
		Return(database.RootShaderProgramObject{}, errors.New("unexpected error"))
}

func testMicroserviceCreateShaderProgramObjectErrorEntityMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbShaderProgram database.RootShaderProgram, dbModelEntity database.RootEntity) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetShaderProgramById(
			gomock.Any(),
			gomock.Eq(dbShaderProgram.ID),
		).Times(1).
		Return(dbShaderProgram, nil)

	store.EXPECT().
		GetEntityById(
			gomock.Any(),
			gomock.Eq(dbModelEntity.ID),
		).Times(1).
		Return(database.RootEntity{}, errors.New("unexpected error"))
}

func testMicroserviceCreateShaderProgramObjectErrorModelTypeMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbShaderProgram database.RootShaderProgram, dbModelEntity database.RootEntity) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetShaderProgramById(
			gomock.Any(),
			gomock.Eq(dbShaderProgram.ID),
		).Times(1).
		Return(dbShaderProgram, nil)

	store.EXPECT().
		GetEntityById(
			gomock.Any(),
			gomock.Eq(dbModelEntity.ID),
		).Times(1).
		Return(dbModelEntity, nil)
}

func testMicroserviceCreateShaderProgramObjectErrorShaderProgramMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbShaderProgram database.RootShaderProgram) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetShaderProgramById(
			gomock.Any(),
			gomock.Eq(dbShaderProgram.ID),
		).Times(1).
		Return(database.RootShaderProgram{}, errors.New("unexpected error"))
}
