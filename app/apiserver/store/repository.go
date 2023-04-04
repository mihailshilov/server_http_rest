package store

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mihailshilov/server_http_rest/app/apiserver/model"
)

// user repository
type UserRepository interface {
	//auth methods
	FindUser(string, string) (*model.User1, error)
	FindUserid(uint64) error
	//jwt methods
	CreateToken(uint64, *model.Service) (string, time.Time, error)
	ExtractTokenMetadata(*http.Request, *model.Service) (*model.AccessDetails, error)
	VerifyToken(*http.Request, *model.Service) (*jwt.Token, error)
	ExtractToken(*http.Request) string
}

// data repository
type DataRepository interface {
	QueryInsertMssql(model.DataBooking) (string, error)
	//sites methods
	QueryInsertBookingPostgres(model.DataBooking) error
	QueryInsertFormsPostgres(model.DataForms) error
	//gaz crm
	RequestGazCrmApiBooking(model.DataBooking, *model.Service) (*model.ResponseGazCrm, error)
	RequestGazCrmApiForms(model.DataForms, *model.Service) (*model.ResponseGazCrm, error)
	// Личный кабинет
	RequestLkOrder(model.DataBooking, *model.Service) (*http.Response, error)
	//PG
	QueryInsertLeadGetPostgres(model.DataLeadGet) error
	QueryInsertWorkListsPostgres(model.DataWorkList) error
	QueryInsertStatusesPostgres(model.DataStatuses) error
	//get methods
	QueryStocksMssql() ([]model.DataStocks, error)
	QueryBasicModelsPriceMssql() ([]model.DataBasicModelsPrice, error)
	QueryBasicModelsPriceMssql2() ([]model.DataBasicModelsPrice, error)
	QueryOptionsPriceMssql() ([]model.DataOptionsPrice, error)
	QueryGeneralPriceMssql() ([]model.DataGeneralPrice, error)
	QuerySprav() ([]model.DataSprav, error)
	QuerySpravModels() ([]model.DataSpravModels, error)
	QueryOptionsData() ([]model.DataOptions, error)
	QueryOptionsDataSprav() ([]model.DataOptionsSprav, error)
	QueryPacketsData() ([]model.DataPackets, error)
	QueryColorsData() ([]model.DataColors, error)
	QueryStatusesLkData() ([]model.DataStatusesLk, error)
	QueryTechData() (*[]model.TechDataObj, error)

	//mailing call method
	CallMSMailing(model.DataBooking, *model.Service) (string, error)
}
