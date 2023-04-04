package apiserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	logger "github.com/mihailshilov/server_http_rest/app/apiserver/logger"
	"github.com/mihailshilov/server_http_rest/app/apiserver/model"
	"github.com/mihailshilov/server_http_rest/app/apiserver/store"
	_ "github.com/mihailshilov/server_http_rest/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API СТТ
// @version 1.0
// @oas 3
// @description API-сервер СТТ
// @contact.name API Support
// @contact.email shilovmo@st.tech
// @host onsales.st.tech
// @BasePath /
// @query.collection.format multi
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// errors
var (
	errIncorrectEmailOrPassword = errors.New("incorrect auth")
	errReg                      = errors.New("service registration error")
	errJwt                      = errors.New("token error")
	errFindUser                 = errors.New("user not found")
	errMssql                    = errors.New("mssql error")
)

// responses
var (
	respGazCrmWorkList = "data work_list recieved"
	respGazCrmLeadGet  = "data lead_get recieved"
	respGazCrmStatuses = "data statuses recieved"
	respBooking        = "data booking sent to gazcrm"
	respForm           = "data form sent to gazcrm"
	errPg              = "error postgres storing"
)

// server configure
type server struct {
	router *mux.Router
	store  store.Store
	config *model.Service
	client *http.Client
}

func newServer(store store.Store, config *model.Service, client *http.Client) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
		config: config,
		client: client,
	}
	s.configureRouter()
	return s
}

// write new token struct
func newToken(token string, exp time.Time) *model.Token_exp {
	return &model.Token_exp{
		Token: token,
		Exp:   exp,
	}
}

// write response struct
func newResponse(status string, response string) *model.Response {
	return &model.Response{
		Status:   status,
		Response: response,
	}
}

// write response struct booking
func newResponseBooking(statusms string, responsems string, statuslk string, responsegcrm string) *model.ResponseBooking {
	return &model.ResponseBooking{
		StatusMs:       statusms,
		ResponseMs:     responsems,
		StatusLK:       statuslk,
		ResponseGazCrm: responsegcrm,
	}
}

// write http error
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})

}

// write http response
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		//httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
		httpSwagger.UIConfig(map[string]string{
			"showExtensions":        "true",
			"onComplete":            `() => { window.ui.setBasePath('v3'); }`,
			"defaultModelRendering": `"model"`,
		}),
		//httpSwagger.Plugins([]string),
	)).Methods(http.MethodGet)

	//open
	s.router.HandleFunc("/authentication", s.handleAuth()).Methods("POST")
	//private
	auth := s.router.PathPrefix("/auth").Subrouter()
	auth.Use(s.middleWare)
	//booking, forms submit
	auth.HandleFunc("/requestbooking", s.handleRequestBooking()).Methods("POST")
	auth.HandleFunc("/requestform", s.handleRequestForm()).Methods("POST")
	//gaz crm
	auth.HandleFunc("/requestleadget", s.handleRequestLeadGetGazCrm()).Methods("POST")
	auth.HandleFunc("/requestworklist", s.handleRequestWorkListGazCrm()).Methods("POST")
	auth.HandleFunc("/requeststatus", s.handleRequestStatusGazCrm()).Methods("POST")
	//stock
	auth.HandleFunc("/getdatastocks", s.handleGetDataStocks()).Methods("GET")
	//prices
	auth.HandleFunc("/getbasicmodelsprice", s.handleBasicModelsPrice()).Methods("GET")
	auth.HandleFunc("/getbasicmodelsprice2", s.handleBasicModelsPrice2()).Methods("GET")
	auth.HandleFunc("/getoptionsprice", s.handleOptionsPrice()).Methods("GET")
	auth.HandleFunc("/getgeneralprice", s.handleGeneralPrice()).Methods("GET")
	//sprav models
	auth.HandleFunc("/getsprav", s.handleSprav()).Methods("GET")
	auth.HandleFunc("/getspravmodels", s.handleSpravModels()).Methods("GET")
	auth.HandleFunc("/techdata", s.handleTechData()).Methods("GET")
	//options
	auth.HandleFunc("/getoptionsdata", s.handleOptionsData()).Methods("GET")
	auth.HandleFunc("/getoptionsdatasprav", s.handleOptionsDataSprav()).Methods("GET")
	auth.HandleFunc("/getpacketsdata", s.handlePacketsData()).Methods("GET")
	//colors
	auth.HandleFunc("/getcolorsdata", s.handleColorsData()).Methods("GET")
	//statuses
	auth.HandleFunc("/getstatusesdata", s.handleStatusesData()).Methods("GET")
}

// HandleAuth godoc
// @Summary Авторизация
// @Description Auth Login
// @Tags Авторизация
// @ID auth-login
// @Accept  json
// @Produce  json
// @Param input body model.User1 true "user info"
// @Success 200 {object} model.Token_exp "OK"
// @Failure 400	{object} model.HTTPerrReg
// @Failure 401	{object} model.HTTPerrIncorrectEmailOrPassword
// @Router /authentication [post]
func (s *server) handleAuth() http.HandlerFunc {

	var req model.User1

	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, errReg)
			logger.ErrorLogger.Println(err)
			return
		}

		u, err := s.store.User().FindUser(req.Email, req.Password)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			logger.ErrorLogger.Println(err)
			return
		}

		token, datetime_exp, err := s.store.User().CreateToken(uint64(u.ID), s.config)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errJwt)
			logger.ErrorLogger.Println(err)
			return
		}
		token_data := newToken(token, datetime_exp)
		s.respond(w, r, http.StatusOK, token_data)
		logger.InfoLogger.Println("token issued success")

	}

}

// Middleware
func (s *server) middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")

		//extract user_id
		user_id, err := s.store.User().ExtractTokenMetadata(r, s.config)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errJwt)
			logger.ErrorLogger.Println(err)
			return
		}

		if err := s.store.User().FindUserid(user_id.UserId); err != nil {
			s.error(w, r, http.StatusUnauthorized, errFindUser)
			logger.ErrorLogger.Println(err)
			return
		}

		next.ServeHTTP(w, r)

	})

}

// handle Client Data
func (s *server) handleRequestBooking() http.HandlerFunc {

	var errMs string

	return func(w http.ResponseWriter, r *http.Request) {

		req := model.DataBooking{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logger.ErrorLogger.Println(err)
			return
		}

		logger.InfoLogger.Println("Запрос от Перкса:\n")
		bodyBytesReq, err := json.Marshal(req)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		logger.InfoLogger.Println(bytes.NewBuffer(bodyBytesReq))
		logger.InfoLogger.Println("\n Конец запроса от Перкса")

		resp, err := s.store.Data().QueryInsertMssql(req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
			logger.ErrorLogger.Println(resp)
			return
		}

		if resp != "Обработка данных прошла успешно" {
			errMs = "Error"
			logger.ErrorLogger.Println(resp)
			return
		} else {
			errMs = "Ok"
			logger.InfoLogger.Println("data booking stored in mssql")

			//respm, err := s.store.Data().CallMSMailing(req, s.config)
			//if err != nil {
			//ErrorLogger.Println(err)
			//ErrorLogger.Println(respm)
			//}
			//InfoLogger.Println("email=" + respm)

		}

		StatusLK := ""

		if req.СlientToken != "" {

			//send order data to lk
			res_lk, err := s.store.Data().RequestLkOrder(req, s.config)
			if err != nil {
				logger.ErrorLogger.Println(err)
			} else {
				logger.InfoLogger.Println(res_lk)
				StatusLK = res_lk.Status
			}

		}

		//request gazcrm api
		respg, err := s.store.Data().RequestGazCrmApiBooking(req, s.config)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		if respg.Status != "OK" {
			logger.ErrorLogger.Println(respg)
			s.respond(w, r, http.StatusBadRequest, newResponseBooking(errMs, resp, StatusLK, respg.Message))
		} else {
			logger.InfoLogger.Println("gazcrm booking data transfer success")
			s.respond(w, r, http.StatusOK, newResponseBooking(errMs, resp, StatusLK, respBooking))
		}

		//insert data in postgres
		if err := s.store.Data().QueryInsertBookingPostgres(req); err != nil {
			logger.ErrorLogger.Println(err)
		} else {
			logger.InfoLogger.Println("sites booking data stored")
		}

	}

}

// handle request forms
func (s *server) handleRequestForm() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		req := model.DataForms{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logger.ErrorLogger.Println(err)
			return
		}

		//request gazcrm api
		respg, err := s.store.Data().RequestGazCrmApiForms(req, s.config)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		if respg.Status != "OK" {
			logger.ErrorLogger.Println(respg)
			s.respond(w, r, http.StatusBadRequest, newResponse("Error", respg.Message))
		} else {
			logger.InfoLogger.Println("gazcrm form data transfer success")
			s.respond(w, r, http.StatusOK, newResponse("Ok", respForm))
		}

		//insert data in postgres
		if err := s.store.Data().QueryInsertFormsPostgres(req); err != nil {
			logger.ErrorLogger.Println(err)
		} else {
			logger.InfoLogger.Println("sites form data stored")
		}

	}

}

// gaz crm
// handle request lead get from gaz crm
func (s *server) handleRequestLeadGetGazCrm() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		req := model.DataLeadGet{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logger.ErrorLogger.Println(err)
			return
		}

		//insert data in postgres
		if err := s.store.Data().QueryInsertLeadGetPostgres(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logger.ErrorLogger.Println(err)
			s.respond(w, r, http.StatusBadRequest, newResponse("Error", errPg))
		} else {
			logger.InfoLogger.Println("gazcrm lead_get inserted in postgres")
			s.respond(w, r, http.StatusOK, newResponse("Ok", respGazCrmLeadGet))
		}

	}

}

// handle request work list from gaz crm
func (s *server) handleRequestWorkListGazCrm() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		req := model.DataWorkList{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logger.ErrorLogger.Println(err)
			return
		}

		//insert data in postgres
		if err := s.store.Data().QueryInsertWorkListsPostgres(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logger.ErrorLogger.Println(err)
			s.respond(w, r, http.StatusBadRequest, newResponse("Error", errPg))
		} else {
			logger.InfoLogger.Println("gazcrm work_list inserted in postgres")
			s.respond(w, r, http.StatusOK, newResponse("Ok", respGazCrmWorkList))
		}

	}

}

// handle request status from gaz crm
func (s *server) handleRequestStatusGazCrm() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		req := model.DataStatuses{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logger.ErrorLogger.Println(err)
			return
		}

		//insert data in postgres
		if err := s.store.Data().QueryInsertStatusesPostgres(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logger.ErrorLogger.Println(err)
			s.respond(w, r, http.StatusBadRequest, newResponse("Error", errPg))
		} else {
			logger.InfoLogger.Println("gazcrm statuses inserted in postgres")
			s.respond(w, r, http.StatusOK, newResponse("Ok", respGazCrmStatuses))
		}

	}

}

// handle request stocks
func (s *server) handleGetDataStocks() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryStocksMssql()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data stocks sent")

	}

}

// handle request basic model price
func (s *server) handleBasicModelsPrice() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryBasicModelsPriceMssql()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data price basic models sent")

	}

}

// handle request basic model price 2
func (s *server) handleBasicModelsPrice2() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryBasicModelsPriceMssql2()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data price basic 2 models sent")

	}

}

// handle request options price
func (s *server) handleOptionsPrice() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryOptionsPriceMssql()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data price options sent")

	}

}

// handle request general price
func (s *server) handleGeneralPrice() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryGeneralPriceMssql()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data price general sent")

	}

}

// handle request sprav
func (s *server) handleSprav() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QuerySprav()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data sprav sent")

	}

}

// handle request sprav
func (s *server) handleSpravModels() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QuerySpravModels()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data sprav sent")

	}

}

// handle request options data
func (s *server) handleOptionsData() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryOptionsData()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data options sent")

	}

}

// handlePacketsDataSprav godoc
// @Summary Получить список недопустимых и обязательных опций
// @Tags Данные по автомобилям для заказа
// @Description Получить список недопустимых и обязательных опций
// @ID get-packetsdatasprav
// @Accept  json
// @Produce  json
// @Success 200 {array} model.DataOptionsSprav "OK"
// @Router /auth/getoptionsdatasprav [get]
// @Security ApiKeyAuth
func (s *server) handleOptionsDataSprav() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryOptionsDataSprav()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data options sprav sent")

	}

}

// handlePacketsData godoc
// @Summary Получить список пакетов опций
// @Tags Данные по автомобилям для заказа
// @Description Получить список пакетов опций
// @ID get-packetsdata
// @Accept  json
// @Produce  json
// @Success 200 {array} model.DataPackets "OK"
// @Router /auth/getpacketsdata [get]
// @Security ApiKeyAuth
func (s *server) handlePacketsData() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryPacketsData()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data packets sent")

	}
}

// handleColorsData godoc
// @Summary Получить цвета автомобилей
// @Tags Данные по автомобилям для заказа
// @Description Получить цвета автомобилей
// @ID get-colorsdata
// @Accept  json
// @Produce  json
// @Success 200 {array} model.DataColors
// @Router /auth/getcolorsdata [get]
// @Security ApiKeyAuth
func (s *server) handleColorsData() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryColorsData()
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data colors sent")

	}

}

// handleStatusesData godoc
// @Summary Получить статусы заказов для личного кабинета
// @Tags Данные для личного кабинета
// @Description Получить статусы заказов для личного кабинета
// @ID get-statusesdata
// @Accept  json
// @Produce  json
// @Success 200 {array} model.DataStatusesLk "OK"
// @Router /auth/getstatusesdata [get]
// @Security ApiKeyAuth
func (s *server) handleStatusesData() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryStatusesLkData()
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		/*
			data := &[]model.DataStatusesLk{
				{"753159", "Заказ подтверждён", "2344504", "X96330200P2882559"},
				{"753160", "Формирование документов", "2344505", "X96330200P2882560"},
				{"753161", "Подготовка к отгрузке", "2344506", "X96330200P2882561"},
				{"753162", "На складе перевозчика", "2344507", "X96330200P2882562"},
				{"753162", "В процессе доставки", "2344508", "X96330200P2882563"},
				{"753162", "Доставлено до дилера", "2344509", "X96330200P2882564"},
			}
		*/

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("data statuses sent")

	}

}

// handleTechData godoc
// @Summary Получить технические характеристика автомобилей
// @Tags Данные по автомобилям для заказа
// @Description Получить технические характеристики
// @ID get-techdata
// @Accept  json
// @Produce  json
// @Success 200 {array} model.TechDataObj "OK"
// @Router /auth/techdata [get]
// @Security ApiKeyAuth
func (s *server) handleTechData() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Data().QueryTechData()

		if err != nil {
			s.error(w, r, http.StatusBadRequest, errMssql)
			logger.ErrorLogger.Println(err)
		}

		s.respond(w, r, http.StatusOK, data)
		logger.InfoLogger.Println("tech data sent")

	}

}
