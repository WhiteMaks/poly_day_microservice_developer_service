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

func TestMicroservice_GetShaderPrograms10(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)
	dbShaderPrograms := []database.RootShaderProgram{dbShaderProgram}

	testMicroserviceGetShaderProgramsMockDB(store, dbDeveloper, dbRelease, dbShaderPrograms)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.GetShaderPrograms10Request{
		ReleaseId: dbShaderProgram.ReleaseID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.getShaderPrograms10(requestMessage)
	expectedMessage := schema.Message{}
	expectedShaderPrograms := make([]schema.ShaderProgram, len(dbShaderPrograms))
	for i, value := range dbShaderPrograms {
		expectedShaderPrograms[i] = schema.GenerateShaderProgramStruct(value.ID, value.Name, value.ShaderEntityID)
	}
	err = schema.FillBodyGetShaderPrograms10ResponseForMessage(expectedShaderPrograms, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_GetShaderPrograms10_NOT_AUTHORIZED(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceUserNotExistMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.GetShaderPrograms10Request{
		ReleaseId: dbShaderProgram.ReleaseID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.getShaderPrograms10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_AUTHORIZED, actualMessage)
}

func TestMicroservice_GetShaderPrograms10_NOT_FOUND_ReleaseNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceGetShaderProgramsErrorReleaseMockDB(store, dbDeveloper, dbRelease)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.GetShaderPrograms10Request{
		ReleaseId: dbShaderProgram.ReleaseID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.getShaderPrograms10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_FOUND, actualMessage)
}

func TestMicroservice_GetShaderPrograms10_INTERNAL_ERROR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceGetShaderProgramsErrorMockDB(store, dbDeveloper, dbRelease)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.GetShaderPrograms10Request{
		ReleaseId: dbShaderProgram.ReleaseID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.getShaderPrograms10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.INTERNAL_ERROR, actualMessage)
}

func TestMicroservice_CreateShaderProgram10(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceCreateShaderProgramMockDB(store, dbDeveloper, dbRelease, dbShaderEntity, dbShaderProgram)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgram10Request{
		Name:           dbShaderProgram.Name,
		ReleaseId:      dbShaderProgram.ReleaseID,
		ShaderEntityId: dbShaderProgram.ShaderEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgram10(requestMessage)
	expectedMessage := schema.Message{}
	err = schema.FillBodyCreateShaderProgram10ResponseForMessage(dbShaderProgram.ID, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_CreateShaderProgram10_NOT_AUTHORIZED(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceUserNotExistMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgram10Request{
		Name:           dbShaderProgram.Name,
		ReleaseId:      dbShaderProgram.ReleaseID,
		ShaderEntityId: dbShaderProgram.ShaderEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgram10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_AUTHORIZED, actualMessage)
}

func TestMicroservice_CreateShaderProgram10_NOT_FOUND_ReleaseNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceCreateShaderProgramErrorReleaseMockDB(store, dbDeveloper, dbRelease)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgram10Request{
		Name:           dbShaderProgram.Name,
		ReleaseId:      dbShaderProgram.ReleaseID,
		ShaderEntityId: dbShaderProgram.ShaderEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgram10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_FOUND, actualMessage)
}

func TestMicroservice_CreateShaderProgram10_NOT_FOUND_EntityNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceCreateShaderProgramErrorEntityMockDB(store, dbDeveloper, dbRelease, dbShaderEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgram10Request{
		Name:           dbShaderProgram.Name,
		ReleaseId:      dbShaderProgram.ReleaseID,
		ShaderEntityId: dbShaderProgram.ShaderEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgram10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_FOUND, actualMessage)
}

func TestMicroservice_CreateShaderProgram10_BAD_REQUES_EntityTypeNotShader(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbModelEntity := generateRandomDBModelEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbModelEntity.ID)

	testMicroserviceCreateShaderProgramErrorShaderTypeMockDB(store, dbDeveloper, dbRelease, dbModelEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgram10Request{
		Name:           dbShaderProgram.Name,
		ReleaseId:      dbShaderProgram.ReleaseID,
		ShaderEntityId: dbShaderProgram.ShaderEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgram10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.BAD_REQUEST, actualMessage)
}

func TestMicroservice_CreateShaderProgram10_INTERNAL_ERROR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceCreateShaderProgramErrorMockDB(store, dbDeveloper, dbRelease, dbShaderEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateShaderProgram10Request{
		Name:           dbShaderProgram.Name,
		ReleaseId:      dbShaderProgram.ReleaseID,
		ShaderEntityId: dbShaderProgram.ShaderEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createShaderProgram10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.INTERNAL_ERROR, actualMessage)
}

func TestMicroservice_DeleteShaderProgram10(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceDeleteShaderProgramMockDB(store, dbDeveloper, dbShaderProgram)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.DeleteShaderProgram10Request{
		ID: dbShaderProgram.ID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.deleteShaderProgram10(requestMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponseWithBodyNil(t, actualMessage)
}

func TestMicroservice_DeleteShaderProgram10_NOT_AUTHORIZED(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceUserNotExistMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.DeleteShaderProgram10Request{
		ID: dbShaderProgram.ID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.deleteShaderProgram10(requestMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_AUTHORIZED, actualMessage)
}

func TestMicroservice_DeleteShaderProgram10_NOT_FOUND(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceDeleteShaderProgramErrorShaderProgramMockDB(store, dbDeveloper, dbShaderProgram)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.DeleteShaderProgram10Request{
		ID: dbShaderProgram.ID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.deleteShaderProgram10(requestMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_FOUND, actualMessage)
}

func TestMicroservice_DeleteShaderProgram10_INTERNAL_ERROR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()
	dbShaderEntity := generateRandomDBShaderEntity()
	dbShaderProgram := generateRandomDBShaderProgram(dbRelease.ID, dbShaderEntity.ID)

	testMicroserviceDeleteShaderProgramErrorMockDB(store, dbDeveloper, dbShaderProgram)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.DeleteShaderProgram10Request{
		ID: dbShaderProgram.ID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.deleteShaderProgram10(requestMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.INTERNAL_ERROR, actualMessage)
}

func testMicroserviceDeleteShaderProgramErrorMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbShaderProgram database.RootShaderProgram) {
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
		DeleteShaderProgramById(
			gomock.Any(),
			gomock.Eq(dbShaderProgram.ID),
		).Times(1).
		Return(errors.New("unexpected error"))
}

func testMicroserviceDeleteShaderProgramErrorShaderProgramMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbShaderProgram database.RootShaderProgram) {
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

func testMicroserviceDeleteShaderProgramMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbShaderProgram database.RootShaderProgram) {
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
		DeleteShaderProgramById(
			gomock.Any(),
			gomock.Eq(dbShaderProgram.ID),
		).Times(1).
		Return(nil)
}

func testMicroserviceGetShaderProgramsErrorMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleaseById(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return(dbRelease, nil)

	store.EXPECT().
		GetShaderProgramsByReleaseId(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return([]database.RootShaderProgram{}, errors.New("unexpected error"))
}

func testMicroserviceGetShaderProgramsErrorReleaseMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleaseById(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return(database.RootRelease{}, errors.New("unexpected error"))
}

func testMicroserviceGetShaderProgramsMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease, dbShaderPrograms []database.RootShaderProgram) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleaseById(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return(dbRelease, nil)

	store.EXPECT().
		GetShaderProgramsByReleaseId(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return(dbShaderPrograms, nil)
}

func testMicroserviceCreateShaderProgramErrorMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease, dbShaderEntity database.RootEntity) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleaseById(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return(dbRelease, nil)

	store.EXPECT().
		GetEntityById(
			gomock.Any(),
			gomock.Eq(dbShaderEntity.ID),
		).Times(1).
		Return(dbShaderEntity, nil)

	store.EXPECT().
		CreateShaderProgram(
			gomock.Any(),
			gomock.Any(),
		).Times(1).
		Return(database.RootShaderProgram{}, errors.New("unexpected error"))
}

func testMicroserviceCreateShaderProgramErrorShaderTypeMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease, dbShaderEntity database.RootEntity) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleaseById(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return(dbRelease, nil)

	store.EXPECT().
		GetEntityById(
			gomock.Any(),
			gomock.Eq(dbShaderEntity.ID),
		).Times(1).
		Return(dbShaderEntity, nil)
}

func testMicroserviceCreateShaderProgramErrorEntityMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease, dbShaderEntity database.RootEntity) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleaseById(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return(dbRelease, nil)

	store.EXPECT().
		GetEntityById(
			gomock.Any(),
			gomock.Eq(dbShaderEntity.ID),
		).Times(1).
		Return(database.RootEntity{}, errors.New("unexpected error"))
}

func testMicroserviceCreateShaderProgramErrorReleaseMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleaseById(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return(database.RootRelease{}, errors.New("unexpected error"))
}

func testMicroserviceCreateShaderProgramMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease, dbShaderEntity database.RootEntity, dbShaderProgram database.RootShaderProgram) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleaseById(
			gomock.Any(),
			gomock.Eq(dbRelease.ID),
		).Times(1).
		Return(dbRelease, nil)

	store.EXPECT().
		GetEntityById(
			gomock.Any(),
			gomock.Eq(dbShaderEntity.ID),
		).Times(1).
		Return(dbShaderEntity, nil)

	store.EXPECT().
		CreateShaderProgram(
			gomock.Any(),
			gomock.Any(),
		).Times(1).
		Return(dbShaderProgram, nil)
}

func generateRandomDBShaderProgram(releaseId int64, shaderEntityId int64) database.RootShaderProgram {
	randomDuration := time.Duration(utils.RandomInt(0, 10000))
	randomCreateAt := time.Now().
		Add(-randomDuration)

	return database.RootShaderProgram{
		ID:             utils.RandomInt(0, 100000),
		Name:           utils.RandomString(56),
		UserID:         utils.RandomInt(0, 100000),
		CreatedAt:      randomCreateAt,
		UpdatedAt:      time.Now(),
		ReleaseID:      releaseId,
		ShaderEntityID: shaderEntityId,
	}
}
