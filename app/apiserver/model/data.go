package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type NullByte string
type NullString string
type NullInt string

func (s *NullByte) Scan(value interface{}) error {

	if value == nil {
		*s = "0"
		return nil
	}
	val, ok := value.([]uint8)
	if !ok {
		*s = "0"
		return nil

	}
	*s = NullByte(string(val))

	return nil
}

func (s *NullString) Scan(value interface{}) error {

	if value == nil || value == "NULL" {
		*s = "nil"
		return nil
	}

	val, ok := value.(string)
	if !ok {
		*s = "nil"
		return nil

	}
	*s = NullString(val)

	return nil
}

func (s *NullInt) Scan(value interface{}) error {

	if value == nil || value == "NULL" {
		*s = "nil"
		return nil
	}

	val, ok := value.(int64)
	if !ok {
		*s = "nil"
		return nil
	}
	*s = NullInt(string(val))

	return nil
}

// Data booking
type DataBooking struct {
	TestMod                   bool   `json:"testmode"` //true - test, false - prod
	RequestId                 string `json:"request_id"`
	SubdivisionsId            string `json:"request_subdivisions_id"` //gazcrm
	Area                      string `json:"request_area"`            //gazcrm
	ActionType                string `json:"request_type"`
	BillNumber                string `json:"request_bill_namber"`
	TimeRequest               string `json:"request_datetime"`
	Comment                   string `json:"request_comment"`
	Consentmailing            string `json:"consent_to_mailing"`
	Vin                       string `json:"car_vin"`
	UniqModCode               int    `json:"car_uniq_code"`
	Modification              string `json:"car_modification"`
	ModFamily                 string `json:"car_family"`
	ModBodyType               string `json:"car_body_type"`
	ModEngine                 string `json:"car_engine"`
	ModBase                   string `json:"car_base"`
	ModTuning                 string `json:"car_tuning"`
	PriceWithNds              int    `json:"car_price"`
	Division                  string `json:"car_division"`
	BrandName                 string `json:"car_brand_name"`
	CarModel                  string `json:"car_model"`
	PreviewUrl                string `json:"car_preview_url"`
	СlientToken               string `json:"client_token"`
	TypeClient                string `json:"client_type"`
	Surname                   string `json:"client_surname"`
	Name                      string `json:"client_name"`
	Patronymic                string `json:"client_patronymic"`
	Email                     string `json:"client_email"`
	PhoneNumber               string `json:"client_phone_number"`
	DateOfBirth               string `json:"client_date_of_birth"`
	PassportSer               string `json:"passport_ser"`
	PassportNumber            string `json:"passport_number"`
	PassportDate              string `json:"passport_date"`       //new
	PassportOrgan             string `json:"passport_organ"`      //new
	PassportOrganCode         string `json:"passport_organ_code"` //new
	Snils                     string `json:"snils"`
	YurAddress                string `json:"client_reg_address"`
	PostAddress               string `json:"post_address"`
	DeliveryAddress           string `json:"delivery_address"`
	DeliveryAddressCode       string `json:"delivery_address_code"`
	Inn                       string `json:"company_inn"`
	Kpp                       string `json:"company_kpp"`
	Ogrn                      string `json:"company_ogrn"`
	CompanyName               string `json:"company_name"`
	CompanyAdress             string `json:"company_adress"`
	BankBik                   string `json:"bank_bik"`
	BankName                  string `json:"bank_name"`
	BankRS                    string `json:"bank_raschetniy_schet"`
	RepresentativeName        string `json:"representative_name"`
	RepresentativeSurname     string `json:"representative_surname"`
	RepresentativePhoneNumber string `json:"representative_phone_number"`
	RepresentativeEmail       string `json:"representative_email"`
	UrlMod                    string `json:"car_model_url"`
	Clientid                  string `json:"clientid_google"` //Google Analytics cookies
	Ymuid                     string `json:"ClientID"`        //Yandex Metrics cookies
	MetricsType               string `json:"metrics_type"`
	СlientIP                  string `json:"client_IP"`
	FormName                  string `json:"form_name"`
	FormId                    string `json:"id_form"`
	HostName                  string `json:"host_name"`
	Hid                       string `json:"company_dadata_id"`
	File                      string `json:"file"`
	SubdivisionsName          string `json:"request_subdivisions_name"`
}

// Validation data booking
func (d *DataBooking) ValidateDataBooking() error {
	return validation.ValidateStruct(
		d,
		validation.Field(&d.RequestId, validation.Required),
		//validation.Field(&d.UniqModCode, validation.Required),
		validation.Field(&d.Modification, validation.Required),
		validation.Field(&d.ModFamily, validation.Required),
		validation.Field(&d.ModBodyType, validation.Required),
		validation.Field(&d.ModEngine, validation.Required),
		validation.Field(&d.ModBase, validation.Required),
		// validation.Field(&d.ModTuning, validation.Required),
		validation.Field(&d.Vin, validation.Required),
		validation.Field(&d.PriceWithNds, validation.Required),
		// validation.Field(&d.TypeClient, validation.In("company", "personal")),
		// validation.Field(&d.Inn, validation.Required),
		// validation.Field(&d.Kpp, validation.Required),
		// validation.Field(&d.Ogrn, validation.Required),
		//validation.Field(&d.YurAddress, validation.Required),
		//validation.Field(&d.DeliveryAddress, validation.Required),
		// validation.Field(&d.Hid, validation.Required),
		// validation.Field(&d.CompanyName, validation.Required),
		//
		// validation.Field(&d.Surname, validation.Required),
		// validation.Field(&d.Name, validation.Required),
		// validation.Field(&d.Patronymic, validation.Required),
		//validation.Field(&d.DateOfBirth, validation.Date("2006-01-02")),
		//validation.Field(&d.PassportSer, validation.Required),
		//validation.Field(&d.PassportNumber, validation.Required),
		// validation.Field(&d.Email, validation.Required),
		// validation.Field(&d.PhoneNumber, validation.Required),
		validation.Field(&d.TimeRequest, validation.Date("2006-01-02T15:04:05")),
		validation.Field(&d.Consentmailing, validation.In("yes", "no")),
		//validation fo gaz crm fields
		validation.Field(&d.Division, validation.In("lcv/mcv", "bus")),
		validation.Field(&d.Area, validation.In("dealer", "distrib")),
		// validation.Field(&d.MetricsType, validation.In("yandex")),
		//validation.Field(&d.TypeClient, validation.In("Юрлицо", "Физлицо")),
		validation.Field(&d.TimeRequest, validation.Date("2006-01-02T15:04:05")),
		validation.Field(&d.ActionType, validation.Required, validation.In("онлайн-заказ", "бронирование", "эквайринг")),
	)
}

// Data get
type DataStocks struct {
	VIN                            string
	Площадка                       string
	Наименование_номенклатуры      string
	Номер_согласно_КД              string
	Дивизион                       string
	Доработчик_Подрядчик           NullString
	Test_truck                     bool
	Телематика                     string
	Номер_шасси                    string
	Номер_двигателя                NullString
	Грузоподъемность_кг            string
	Цвет                           string
	ЦветИд                         string
	ЦветRGB                        NullString
	Вариант_сборки                 string
	Расшифровка_варианта_сборки    string
	Вариант_сборки_свернутый       NullString
	Год_VIN                        string
	Дата_сборки                    NullString
	Справочная_стоимость_по_прайсу string
	Дата_отгрузка                  NullString
	Дата_прихода                   NullString
	Страна                         NullString
	Контрагент_получателя          string
	Стоянка                        string
	Город_стоянки                  NullString
	Площадка_получателя_Ид         string
	Контрагент_получателя_Ид       string
	Город_стоянки_Ид               NullString
	Номер_заявки                   NullString
	Для_доработки                  NullString
	Номерной_товар                 string
}

// data price basic models
type DataBasicModelsPrice struct {
	Товар          string
	НачалоДействия string
	Цена           string
	НДС            string
	СтавкаНДС      string
}

// data price options
type DataOptionsPrice struct {
	ЕНСП_Модификация_Ид NullString
	ТоварИд             string
	Товар               string
	ЗначениеОпцииИд     string
	ЗначениеОпции       string
	ОбозначениеОпции    string
	Цена                NullByte
	СтавкаНДС_Ид        string
	НДС                 NullByte
	НачалоДействия      string
	СоставПакета        NullString
}

// data price general
type DataGeneralPrice struct {
	Товар                    string
	ВариантСборки            string
	ВариантСборкиРазвернутый string
	Цена                     string
	СтавкаНДС                string
	НДС                      string
	НачалоДействия           string
}

// data sprav models
type DataSpravModels struct {
	Ид                         string
	Наименование               string
	Дивизион                   string
	СтатусМоделиВПроизводстве  string
	БазовыйТовар               NullString
	ХарактеристикиНоменклатуры string
	Цена                       NullString
	СтавкаНДС                  NullString
	НДС                        NullString
}

// data sprav
type DataSprav struct {
	Наименование                               string
	НомерСогласноКД                            string
	Дивизион                                   string
	СтатусМоделиВПроизводстве                  string
	МассаСнагрузкой                            NullString
	МассаБезНагрузки                           NullString
	ОписаниеДляПрайса                          string
	База                                       string
	БазаАвтомобиляДлина                        string
	ТипКузова                                  string
	ТипФургона                                 string
	ОбозначениеДвигателя                       string
	ОбъемДвигателя                             string
	ВидТоплива                                 string
	СтабилизаторЗаднейПодвески                 string
	ГорныйТормоз                               string
	ТормознаяСистемаТип                        string
	ЦветаДопустимыеВЭтомМесяце                 string
	ОпцииДопустимыеВЭтомМесяце                 string
	ОпцииПоУмолчанию                           string
	ЧислоПосадочныхМест                        string
	ЭкКласс                                    string
	Привод                                     string
	Семейство                                  string
	Лебедка                                    string
	КПП                                        string
	ГБО                                        string
	Надстройка                                 string
	ОсобенностьНадстройки                      string
	БазовыйТовар                               NullString
	ОпцииАЗ                                    string
	ХарактеристикиНоменклатуры                 string
	ИзЭПТС_ДопустимаяМаксимальнаяМассаСтандарт *string
	ИзЭПТС_ДопустимаяМаксимальнаяМасса9РА      *string
	ИзЭПТС_ДопустимаяМаксимальнаяМасса9РВ      *string
	ИзЭПТС_СнаряженнаяМасса                    *string
	ДоступностьКЗаказу                         string
}

// options data
type DataOptions struct {
	НоменклатураИд                 string
	НоменклатураНаименование       string
	ИдГруппыОпций                  NullString
	КраткоеНаименованиеГруппыОпций NullString
	ПолноеНаименованиеГруппыОпций  NullString
	ОпцияИд                        string
	КраткоеНаименованиеОпции       string
	ПолноеНаименованиеОпции        string
	ОписаниеОпции                  NullString
	ЦенаОпции                      NullByte
	Обязательная                   string
	ВыбранаПоУмолчанию             string
	ЭтоПакет                       string
	ЦенаНулл                       string
}

// options sprav data
type DataOptionsSprav struct {
	НоменклатураИд           string `json:"НоменклатураИд" extensions:"x-order=a" example:"865436"`
	НоменклатураНаименование string `json:"НоменклатураНаименование" extensions:"x-order=b" example:"ГАЗ-А21R33-10"`
	ЗначениеОпции1           string `json:"ЗначениеОпции1" extensions:"x-order=c" example:"560"`
	ЗначениеОпции2           string `json:"ЗначениеОпции2" extensions:"x-order=d" example:"691"`
	КодОпции1                string `json:"КодОпции1" extensions:"x-order=e" example:"8BA"`
	КодОпции2                string `json:"КодОпции2" extensions:"x-order=f" example:"8LB"`
	ВидСочетания             string `json:"ВидСочетания" extensions:"x-order=g" example:"недопустимое"`
}

// packets data
type DataPackets struct {
	НоменклатураИд                 string `json:"НоменклатураИд" extensions:"x-order=a" example:"865436"`
	НоменклатураНаименование       string `json:"НоменклатураНаименование" extensions:"x-order=b" example:"ГАЗ-А21R33-10"`
	ИдПакета                       string `json:"ИдПакета" extensions:"x-order=c" example:"975"`
	КраткоеНаименованиеПакета      string `json:"КраткоеНаименованиеПакета" extensions:"x-order=d" example:"ST(N)"`
	ПолноеНаименованиеПакета       string `json:"ПолноеНаименованиеПакета" extensions:"x-order=e" example:"ПолноеНаименованиеПакета": "Пакет \"Стандарт\" для 1-ряд а/м Газель Next борт, шасси и САТ, дизель, бензин"`
	ИдГруппыОпций                  string `json:"ИдГруппыОпций" extensions:"x-order=f" example:"141"`
	КраткоеНаименованиеГруппыОпций string `json:"КраткоеНаименованиеГруппыОпций" extensions:"x-order=g" example:"2K"`
	ПолноеНаименованиеГруппыОпций  string `json:"ПолноеНаименованиеГруппыОпций" extensions:"x-order=h" example:"Сиденье водителя"`
	ИдОпции                        string `json:"ИдОпции" extensions:"x-order=i" example:"684"`
	ПолноеНаименованиеОпции        string `json:"ПолноеНаименованиеОпции" extensions:"x-order=j" example:"Сиденье водителя с подлокотником"`
	КраткоеНаименованиеОпции       string `json:"КраткоеНаименованиеОпции" extensions:"x-order=k" example:"2KB"`
}

// colors data
type DataColors struct {
	НоменклатураИд           string  `json:"НоменклатураИд" extensions:"x-order=a" example:"865436"`
	НоменклатураНаименование string  `json:"НоменклатураНаименование" extensions:"x-order=b" example:"ГАЗ-А21R33-10"`
	ЦветИд                   string  `json:"ЦветИд" extensions:"x-order=c" example:"996"`
	Наименование             string  `json:"Наименование" extensions:"x-order=d" example:"СИЛЬВЕР ЛАЙТ"`
	ПолноеНаименование       string  `json:"ПолноеНаименование" extensions:"x-order=e" example:"СВЕТЛО-СЕРЫЙ"`
	ЦветRGB                  string  `json:"ЦветRGB" extensions:"x-order=f" example:"157,163,166"`
	Слойность                *string `json:"Слойность" extensions:"x-order=g" example:"2"`
}

// data options
type OptionsData struct {
}

// Data forms
type DataForms struct {
	//gaz crm fields
	TimeRequest      string `json:"event_datetime"` //general field with booking
	RequestId        string `json:"request_id"`     //general field with booking
	SubdivisionsId   string `json:"subdivisions_id"`
	SubdivisionsName string `json:"subdivisions_name"`
	FormName         string `json:"form_name"`
	FormId           string `json:"id_form"`
	HostName         string `json:"host_name"`
	Division         string `json:"division"`
	Area             string `json:"area"`
	BrandName        string `json:"brand_name"`
	CarModel         string `json:"car_model"`
	Clientid         string `json:"ClientID"` //general field with booking
	MetricsType      string `json:"metrics_type"`
	СlientIP         string `json:"client_IP"`
	TypeClient       string `json:"client_type"`         //general field with booking
	CompanyName      string `json:"client_company_name"` //general field with booking
	Name             string `json:"client_name"`         //general field with booking
	Email            string `json:"client_email"`        //general field with booking
	PhoneNumber      string `json:"client_phone_number"` //general field with booking
	Comment          string `json:"commentary"`          //general field with booking
	Consentmailing   string `json:"agreement_mailing"`   //general field with booking
	//additional fields
	ActionType   string `json:"action_type"`
	Modification string `json:"modification"`
	ModFamily    string `json:"mod_family"`
	ModBodyType  string `json:"mod_body_type"`
	ModEngine    string `json:"mod_engine"`
	ModBase      string `json:"mod_base"`
	ModTuning    string `json:"mod_tuning"`
	Vin          string `json:"vin"`
	PriceWithNds int    `json:"price"`
	UrlMod       string `json:"url_mod"`
}

// Validation data fiz
func (d *DataForms) ValidateDataForms() error {
	return validation.ValidateStruct(
		d,
		validation.Field(&d.RequestId, validation.Required),
		validation.Field(&d.SubdivisionsId, validation.Required),
		validation.Field(&d.SubdivisionsName, validation.Required),
		validation.Field(&d.FormName, validation.Required),
		validation.Field(&d.FormId, validation.Required),
		validation.Field(&d.HostName, validation.Required),
		validation.Field(&d.BrandName, validation.Required),
		validation.Field(&d.CarModel, validation.Required),
		validation.Field(&d.Clientid, validation.Required),
		validation.Field(&d.MetricsType, validation.Required),
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Email, validation.Required),
		validation.Field(&d.PhoneNumber, validation.Required),
		validation.Field(&d.UrlMod, validation.Required),
		validation.Field(&d.Modification, validation.Required),
		validation.Field(&d.ModFamily, validation.Required),
		validation.Field(&d.ModBodyType, validation.Required),
		validation.Field(&d.ModEngine, validation.Required),
		validation.Field(&d.ModBase, validation.Required),
		validation.Field(&d.ModTuning, validation.Required),
		validation.Field(&d.Vin, validation.Required),
		validation.Field(&d.PriceWithNds, validation.Required),
		validation.Field(&d.Division, validation.In("lcv/mcv", "bus")),
		validation.Field(&d.Area, validation.In("dealer", "distrib")),
		validation.Field(&d.MetricsType, validation.In("yandex")),
		validation.Field(&d.TypeClient, validation.In("company", "personal")),
		validation.Field(&d.TimeRequest, validation.Date("2006-01-02T15:04:05")),
		validation.Field(&d.ActionType, validation.Required, validation.In("form")),
		validation.Field(&d.Consentmailing, validation.In("yes", "no")),
	)
}

// gaz crm
// data struct for call gaz crm api method
type DataGazCrm struct {
	Data []*DataGazCrmReq `json:"Data"`
}

// data struct for call gaz crm api method
type DataGazCrmReq struct {
	//gaz crm fields
	TimeRequest       string `json:"event_datetime,omitempty"` //general field with booking
	RequestId         string `json:"request_id,omitempty"`     //general field with booking
	SubdivisionsId    string `json:"subdivisions_id,omitempty"`
	SubdivisionsName  string `json:"subdivisions_name,omitempty"`
	FormName          string `json:"form_name,omitempty"`
	FormId            string `json:"id_form,omitempty"`
	HostName          string `json:"host_name,omitempty"`
	Division          string `json:"division,omitempty"`
	Area              string `json:"area,omitempty"`
	BrandName         string `json:"brand_name,omitempty"`
	CarModel          string `json:"car_model,omitempty"`
	ClientID          string `json:"ClientID,omitempty"` //general field with booking
	MetricsType       string `json:"metrics_type,omitempty"`
	СlientIP          string `json:"client_IP,omitempty"`
	TypeClient        string `json:"client_type,omitempty"`         //general field with booking
	CompanyName       string `json:"client_company_name,omitempty"` //general field with booking
	СlientName        string `json:"client_name,omitempty"`         //general field with booking
	ClientEmail       string `json:"client_email,omitempty"`        //general field with booking
	ClientPhoneNumber string `json:"client_phone_number,omitempty"` //general field with booking
	Commentary        string `json:"commentary,omitempty"`          //general field with booking
	AgreementMailing  string `json:"agreement_mailing,omitempty"`   //general field with booking
}

// data struct for call gaz crm api method
// lead_get gaz crm
type DataLeadGet struct {
	Data []DataLeadGet_Gazcrm `json:"Data"`
}

// lead_get gaz crm
type DataLeadGet_Gazcrm struct {
	TimeRequest      string `json:"event_datetime,omitempty"`
	EventName        string `json:"event_name,omitempty"`
	RequestId        string `json:"request_id,omitempty"`
	SubdivisionsId   string `json:"subdivisions_id,omitempty"`
	SubdivisionsName string `json:"subdivisions_name,omitempty"`
	FormName         string `json:"form_name,omitempty"`
	HostName         string `json:"host_name,omitempty"`
	Division         string `json:"division,omitempty"`
	Area             string `json:"area,omitempty"`
	BrandName        string `json:"brand_name,omitempty"`
	ClientID         string `json:"ClientID,omitempty"`
	MetricsType      string `json:"metrics_type,omitempty"`
}

// work_list gaz crm
type DataWorkList struct {
	Data []DataWorkList_Gazcrm `json:"Data"`
}

// work_list gaz crm
type DataWorkList_Gazcrm struct {
	TimeRequest      string `json:"event_datetime,omitempty"`
	EventName        string `json:"event_name,omitempty"`
	GazcrmClientId   string `json:"gazcrm_client_id,omitempty"`
	GazCrmWorkListId string `json:"gazcrm_worklist_id,omitempty"`
}

// status gaz crm
type DataStatuses struct {
	Data []DataStatuses_Gazcrm `json:"Data"`
}

// status gaz crm
type DataStatuses_Gazcrm struct {
	TimeRequest      string `json:"event_datetime,omitempty"`
	EventName        string `json:"event_name,omitempty"`
	RequestId        string `json:"request_id,omitempty"`
	GazcrmClientId   string `json:"gazcrm_client_id,omitempty"`
	GazCrmWorkListId string `json:"gazcrm_worklist_id,omitempty"`
	ClientID         string `json:"ClientID,omitempty"`
	MetricsType      string `json:"metrics_type,omitempty"`
}

// status lk
type DataStatusesLk struct {
	ИдЗаказа      string `json:"order_id,omitempty" extensions:"x-order=a" example:"123"`
	СтатусЗаказа  string `json:"order_status,omitempty" extensions:"x-order=b" example:"В процессе доставки"`
	НомернойТовар string `json:"id_isk,omitempty" extensions:"x-order=c" example:"2281063"`
	ВИН           string `json:"vin,omitempty" extensions:"x-order=d" example:"X96A21R32N2856911"`
}

// resp struct api gaz crm
type ResponseGazCrm struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type DataLkOrder struct {
	Id          string `json:"id"`
	StageCode   string `json:"stageCode"`
	Title       string `json:"title"`
	Vin         string `json:"vin"`
	Cost        string `json:"cost"`
	ModelFamily string `json:"modelFamily"`
	PreviewUrl  string `json:"previewUrl"`
}

type DataLkProfile struct {
	User         DataLkProfileUser         `json:"user"`
	PersonalData DataLkProfilePersonalData `json:"personalData"`
	Address      []DataLkProfileAddress    `json:"address"`
}

type DataLkProfileUser struct {
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	BirthDay   string `json:"birthDay"`
}

type DataLkProfilePersonalData struct {
	PassportSeries     string `json:"passportSeries"`
	PassportNumber     string `json:"passportNumber"`
	PassportIssuedBy   string `json:"passportIssuedBy"`
	PassportIssuerCode string `json:"passportIssuerCode"`
	PassportIssueDate  string `json:"passportIssueDate"`
	Snils              string `json:"snils"`
}

type DataLkProfileAddress struct {
	TypeCode          string `json:"typeCode"`
	UnrestrictedValue string `json:"unrestrictedValue"`
}

type TechData struct {
	Data TechDataObj `json:"Data"`
}

type TechDataObj struct {
	Ид                   int    `json:"Ид" extensions:"x-order=a" example:"866508"`
	Модель               string `json:"Модель" extensions:"x-order=b" example:"ГАЗ-А21R25-20"`
	ИдКатегории          int    `json:"ИдКатегории" extensions:"x-order=c" example:"415"`
	Категория            string `json:"Категория" extensions:"x-order=d" example:"Характеристики"`
	ИдРодителя           int    `json:"ИдРодителя" extensions:"x-order=e" example:"422"`
	НаименованиеРодителя string `json:"НаименованиеРодителя" extensions:"x-order=f" example:"Шины"`
	ИдСвойства           int    `json:"ИдСвойства" extensions:"x-order=g" example:"462"`
	НаименованиеСвойства string `json:"НаименованиеСвойства" extensions:"x-order=h" example:"Размерность"`
	ИдЗначенияСвойства   int    `json:"ИдЗначенияСвойства" extensions:"x-order=i" example:"5809"`
	ЗначениеСвойства     string `json:"ЗначениеСвойства" extensions:"x-order=j" example:"185/75R16C"`
}
