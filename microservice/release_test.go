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

func TestMicroservice_GetReleases10(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbReleases := generateRandomDBReleases(100)

	testMicroserviceGetReleasesMockDB(store, dbDeveloper, dbReleases)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)

	actualMessage := ms.getReleases10(requestMessage)
	expectedMessage := schema.Message{}
	expectedReleases := make([]schema.Release, len(dbReleases))
	for i, value := range dbReleases {
		expectedReleases[i] = schema.GenerateReleaseStruct(value.ID, value.Name)
	}
	err := schema.FillBodyGetReleases10ResponseForMessage(expectedReleases, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_GetReleases10_NOT_AUTHORIZED(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()

	testMicroserviceUserNotExistMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)

	actualMessage := ms.getReleases10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_AUTHORIZED, actualMessage)
}

func TestMicroservice_GetReleases10_INTERNAL_ERROR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()

	testMicroserviceGetReleasesErrorMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)

	actualMessage := ms.getReleases10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.INTERNAL_ERROR, actualMessage)
}

func TestMicroservice_CreateRelease10(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()

	testMicroserviceCreateReleaseMockDB(store, dbDeveloper, dbRelease)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateRelease10Request{
		Name: dbRelease.Name,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createRelease10(requestMessage)
	expectedMessage := schema.Message{}
	err = schema.FillBodyCreateRelease10ResponseForMessage(dbRelease.ID, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_CreateRelease10_NOT_AUTHORIZED(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()

	testMicroserviceUserNotExistMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateRelease10Request{
		Name: dbRelease.Name,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createRelease10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_AUTHORIZED, actualMessage)
}

func TestMicroservice_CreateRelease10_INTERNAL_ERROR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()

	testMicroserviceCreateReleaseErrorMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateRelease10Request{
		Name: dbRelease.Name,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createRelease10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.INTERNAL_ERROR, actualMessage)
}

func TestMicroservice_CreateDelete10(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()

	testMicroserviceDeleteReleaseMockDB(store, dbDeveloper, dbRelease)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.DeleteRelease10Request{
		ID: dbRelease.ID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.deleteRelease10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponseWithBodyNil(t, actualMessage)
}

func TestMicroservice_DeleteRelease10_NOT_AUTHORIZED(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()

	testMicroserviceUserNotExistMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.DeleteRelease10Request{
		ID: dbRelease.ID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.deleteRelease10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_AUTHORIZED, actualMessage)
}

func TestMicroservice_DeleteRelease10_INTERNAL_ERROR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbRelease := generateRandomDBRelease()

	testMicroserviceDeleteReleaseErrorMockDB(store, dbDeveloper, dbRelease)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.DeleteRelease10Request{
		ID: dbRelease.ID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.deleteRelease10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.INTERNAL_ERROR, actualMessage)
}

func testMicroserviceGetReleasesErrorMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleases(gomock.Any()).
		Times(1).
		Return([]database.RootRelease{}, errors.New("unexpected error"))
}

func testMicroserviceGetReleasesMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbReleases []database.RootRelease) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetReleases(gomock.Any()).
		Times(1).
		Return(dbReleases, nil)
}

func testMicroserviceDeleteReleaseMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		DeleteReleaseById(
			gomock.Any(),
			dbRelease.ID,
		).Times(1).
		Return(nil)
}

func testMicroserviceDeleteReleaseErrorMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		DeleteReleaseById(
			gomock.Any(),
			dbRelease.ID,
		).Times(1).
		Return(errors.New("unexpected error"))
}

func testMicroserviceCreateReleaseMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbRelease database.RootRelease) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		CreateRelease(
			gomock.Any(),
			gomock.Any(),
		).Times(1).
		Return(dbRelease, nil)
}

func testMicroserviceCreateReleaseErrorMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		CreateRelease(
			gomock.Any(),
			gomock.Any(),
		).Times(1).
		Return(database.RootRelease{}, errors.New("unexpected error"))
}

func generateRandomDBReleases(count int) []database.RootRelease {
	dbReleases := make([]database.RootRelease, count)

	for i := 0; i < count; i++ {
		dbReleases[i] = generateRandomDBRelease()
	}

	return dbReleases
}

func generateRandomDBRelease() database.RootRelease {
	randomDuration := time.Duration(utils.RandomInt(0, 10000))
	randomCreateAt := time.Now().
		Add(-randomDuration)

	return database.RootRelease{
		ID:        utils.RandomInt(0, 100000),
		UserID:    utils.RandomInt(0, 100000),
		CreatedAt: randomCreateAt,
		UpdatedAt: time.Now(),
	}
}
