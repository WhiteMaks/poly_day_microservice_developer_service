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

func TestMicroservice_GetEntities10(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbEntities := generateRandomDBEntities()

	testMicroserviceGetEntitiesMockDB(store, dbDeveloper, dbEntities)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)

	actualMessage := ms.getEntities10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	assert.Equal(t, schema.OK, actualMessage.Status, "status")
	assert.NotNil(t, actualMessage.Body, "body")
}

func TestMicroservice_GetEntities10_NOT_AUTHORIZED(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()

	testMicroserviceUserNotExistMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)

	actualMessage := ms.getEntities10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_AUTHORIZED, actualMessage)
}

func TestMicroservice_GetEntities10_INTERNAL_ERROR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()

	testMicroserviceGetEntitiesErrorMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)

	actualMessage := ms.getEntities10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.INTERNAL_ERROR, actualMessage)
}

func TestMicroservice_CreateEntity10_FOLDER(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbEntity := generateRandomDBFolderEntity()

	testMicroserviceCreateEntityMockDB(store, dbDeveloper, dbEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateEntity10Request{
		Name:           dbEntity.Name,
		Type:           schema.EntityType(dbEntity.Type),
		ParentEntityId: dbEntity.ParentEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createEntity10(requestMessage)
	expectedMessage := schema.Message{}
	err = schema.FillBodyCreateEntity10ResponseForMessage(dbEntity.ID, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_CreateEntity10_TEXTURE(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbEntity := generateRandomDBTextureEntity()

	testMicroserviceCreateEntityMockDB(store, dbDeveloper, dbEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateEntity10Request{
		Name:           dbEntity.Name,
		Type:           schema.EntityType(dbEntity.Type),
		ParentEntityId: dbEntity.ParentEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createEntity10(requestMessage)
	expectedMessage := schema.Message{}
	err = schema.FillBodyCreateEntity10ResponseForMessage(dbEntity.ID, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_CreateEntity10_SOUND(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbEntity := generateRandomDBSoundEntity()

	testMicroserviceCreateEntityMockDB(store, dbDeveloper, dbEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateEntity10Request{
		Name:           dbEntity.Name,
		Type:           schema.EntityType(dbEntity.Type),
		ParentEntityId: dbEntity.ParentEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createEntity10(requestMessage)
	expectedMessage := schema.Message{}
	err = schema.FillBodyCreateEntity10ResponseForMessage(dbEntity.ID, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_CreateEntity10_SCRIPT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbEntity := generateRandomDBScriptEntity()

	testMicroserviceCreateEntityMockDB(store, dbDeveloper, dbEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateEntity10Request{
		Name:           dbEntity.Name,
		Type:           schema.EntityType(dbEntity.Type),
		ParentEntityId: dbEntity.ParentEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createEntity10(requestMessage)
	expectedMessage := schema.Message{}
	err = schema.FillBodyCreateEntity10ResponseForMessage(dbEntity.ID, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_CreateEntity10_MUSIC(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbEntity := generateRandomDBMusicEntity()

	testMicroserviceCreateEntityMockDB(store, dbDeveloper, dbEntity)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateEntity10Request{
		Name:           dbEntity.Name,
		Type:           schema.EntityType(dbEntity.Type),
		ParentEntityId: dbEntity.ParentEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createEntity10(requestMessage)
	expectedMessage := schema.Message{}
	err = schema.FillBodyCreateEntity10ResponseForMessage(dbEntity.ID, &expectedMessage)
	assert.NoError(t, err, "Error must be nil")

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkSuccessResponse(t, expectedMessage, actualMessage)
}

func TestMicroservice_CreateEntity10_NOT_AUTHORIZED(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbEntity := generateRandomDBMusicEntity()

	testMicroserviceUserNotExistMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateEntity10Request{
		Name:           dbEntity.Name,
		Type:           schema.EntityType(dbEntity.Type),
		ParentEntityId: dbEntity.ParentEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createEntity10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.NOT_AUTHORIZED, actualMessage)
}

func TestMicroservice_CreateEntity10_BAD_REQUEST(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbEntity := generateRandomDBMusicEntity()

	testMicroserviceCreateEntityBadRequestMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateEntity10Request{
		Name:           dbEntity.Name,
		Type:           "NOT EXIST ENTITY TYPE",
		ParentEntityId: dbEntity.ParentEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createEntity10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.BAD_REQUEST, actualMessage)
}

func TestMicroservice_CreateEntity10_INTERNAL_ERROR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockStore(ctrl)
	ms := newMs(store)

	dbDeveloper := generateRandomDBDeveloper()
	dbEntity := generateRandomDBMusicEntity()

	testMicroserviceCreateEntityErrorMockDB(store, dbDeveloper)

	requestMessage := prepareRequestMessage(dbDeveloper.UserID)
	body := schema.CreateEntity10Request{
		Name:           dbEntity.Name,
		Type:           schema.EntityType(dbEntity.Type),
		ParentEntityId: dbEntity.ParentEntityID,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err, "Error must be nil")
	requestMessage.Body = bodyBytes

	actualMessage := ms.createEntity10(requestMessage)

	checkResponseWithoutBody(t, requestMessage, actualMessage)
	checkErrorResponse(t, schema.INTERNAL_ERROR, actualMessage)
}

func testMicroserviceCreateEntityErrorMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		CreateEntity(
			gomock.Any(),
			gomock.Any(),
		).
		Times(1).
		Return(database.RootEntity{}, errors.New("unexpected error"))
}

func testMicroserviceCreateEntityBadRequestMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)
}

func testMicroserviceCreateEntityMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbEntity database.RootEntity) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		CreateEntity(
			gomock.Any(),
			gomock.Any(),
		).
		Times(1).
		Return(dbEntity, nil)
}

func testMicroserviceGetEntitiesMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper, dbEntities []database.RootEntity) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetEntitiesByParentEntityId(
			gomock.Any(),
			gomock.Eq(int64(0)),
		).
		Times(1).
		Return(dbEntities, nil)

	store.EXPECT().
		GetEntitiesByParentEntityId(
			gomock.Any(),
			gomock.Not(int64(0)),
		).
		Times(len(dbEntities)).
		Return([]database.RootEntity{}, nil)
}

func testMicroserviceGetEntitiesErrorMockDB(store *mockdatabase.MockStore, dbDeveloper database.RootDeveloper) {
	store.EXPECT().
		GetDeveloperById(
			gomock.Any(),
			gomock.Eq(dbDeveloper.UserID),
		).Times(1).
		Return(dbDeveloper, nil)

	store.EXPECT().
		GetEntitiesByParentEntityId(
			gomock.Any(),
			gomock.Eq(int64(0)),
		).
		Times(1).
		Return([]database.RootEntity{}, errors.New("unexpected error"))
}

func generateRandomDBEntities() []database.RootEntity {
	result := []database.RootEntity{
		{
			ID:             utils.RandomInt(0, 100000),
			UserID:         utils.RandomInt(0, 100000),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			ParentEntityID: 0,
			Name:           "Folder",
			Type:           *schema.FOLDER.String(),
		},
		generateRandomDBModelEntity(),
		generateRandomDBShaderEntity(),
	}

	return result
}

func generateRandomDBMusicEntity() database.RootEntity {
	return database.RootEntity{
		ID:             utils.RandomInt(0, 100000),
		UserID:         utils.RandomInt(0, 100000),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ParentEntityID: 0,
		Name:           utils.RandomString(15),
		Type:           *schema.MUSIC.String(),
	}
}

func generateRandomDBSoundEntity() database.RootEntity {
	return database.RootEntity{
		ID:             utils.RandomInt(0, 100000),
		UserID:         utils.RandomInt(0, 100000),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ParentEntityID: 0,
		Name:           utils.RandomString(15),
		Type:           *schema.SOUND.String(),
	}
}

func generateRandomDBScriptEntity() database.RootEntity {
	return database.RootEntity{
		ID:             utils.RandomInt(0, 100000),
		UserID:         utils.RandomInt(0, 100000),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ParentEntityID: 0,
		Name:           utils.RandomString(15),
		Type:           *schema.SCRIPT.String(),
	}
}

func generateRandomDBTextureEntity() database.RootEntity {
	return database.RootEntity{
		ID:             utils.RandomInt(0, 100000),
		UserID:         utils.RandomInt(0, 100000),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ParentEntityID: 0,
		Name:           utils.RandomString(15),
		Type:           *schema.TEXTURE.String(),
	}
}

func generateRandomDBFolderEntity() database.RootEntity {
	return database.RootEntity{
		ID:             utils.RandomInt(0, 100000),
		UserID:         utils.RandomInt(0, 100000),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ParentEntityID: 0,
		Name:           utils.RandomString(15),
		Type:           *schema.FOLDER.String(),
	}
}

func generateRandomDBShaderEntity() database.RootEntity {
	return database.RootEntity{
		ID:             utils.RandomInt(0, 100000),
		UserID:         utils.RandomInt(0, 100000),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ParentEntityID: 0,
		Name:           utils.RandomString(15),
		Type:           *schema.SHADER.String(),
	}
}

func generateRandomDBModelEntity() database.RootEntity {
	return database.RootEntity{
		ID:             utils.RandomInt(0, 100000),
		UserID:         utils.RandomInt(0, 100000),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ParentEntityID: 0,
		Name:           utils.RandomString(15),
		Type:           *schema.MODEL.String(),
	}
}
