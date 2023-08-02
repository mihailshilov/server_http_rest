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
	QueryInsertMssqlBookingF(model.DataBookingF) (string, error)
	QueryInsertMssqlBookingU(model.DataBookingU) (string, error)
	//sites methods
	QueryInsertBookingPostgres(model.DataBooking) error
	QueryInsertBookingPostgresF(model.DataBookingF) error
	QueryInsertBookingPostgresU(model.DataBookingU) error
	QueryInsertFormsPostgres(model.DataForms) error
	QueryFormsPostgres(model.Form, bool) error
	//gaz crm
	RequestGazCrmApiBooking(model.DataBooking, *model.Service) (*model.ResponseGazCrm, error)
	RequestGazCrmApiBookingF(model.DataBookingF, *model.Service) (*model.ResponseGazCrm, error)
	RequestGazCrmApiBookingU(model.DataBookingU, *model.Service) (*model.ResponseGazCrm, error)
	RequestGazCrmApiForms(model.DataForms, *model.Service) (*model.ResponseGazCrm, error)
	RequestForms(model.Form, *model.Service) (*model.ResponseGazCrm, error)
	// Личный кабинет
	RequestLkOrder(model.DataBooking, *model.Service) (*http.Response, error)
	RequestLkOrderF(model.DataBookingF, *model.Service) (*http.Response, error)
	RequestLkOrderU(model.DataBookingU, *model.Service) (*http.Response, error)
	RequestLkProfile(model.DataBooking, *model.Service) (*http.Response, error)
	RequestLkProfileF(model.DataBookingF, *model.Service) (*http.Response, error)
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
	QueryModels() ([]model.DataModels, error)
	QuerySpravModels() ([]model.DataSpravModels, error)
	QueryOptionsData() ([]model.DataOptions, error)
	QueryOptions() ([]model.Options, error) ///new
	QuerySpecial() ([]model.Special, error) ///new
	QueryOptionsDataSprav() ([]model.DataOptionsSprav, error)
	QueryPacketsData() ([]model.DataPackets, error)
	QueryPackets() ([]model.DataPackets_l, error)
	QueryColorsData() ([]model.DataColors, error)
	QueryStatusesLkData() ([]model.DataStatusesLk, error)
	QueryTechData() (*[]model.TechDataObj, error)

	QueryGreyINN() ([]string, error)

	//mailing call method
	CallMSMailing(model.DataBooking, *model.Service) (string, error)

	//Rabbitmq

	AddToRabbit([]byte, string, string) error
}
