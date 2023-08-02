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
	TestMod                   bool   `json:"testmode" extensions:"x-order=01" example:"true"`
	RequestId                 string `json:"request_id" extensions:"x-order=02" example:"1019157100531"`
	SubdivisionsId            string `json:"request_subdivisions_id" extensions:"x-order=03" example:"6286424721659083063"`
	SubdivisionsName          string `json:"request_subdivisions_name" extensions:"x-order=04" example:"Авторитейл"`
	Area                      string `json:"request_area" extensions:"x-order=05" example:"distrib"`
	ActionType                string `json:"request_type" extensions:"x-order=06" example:"онлайн-заказ"`
	BillNumber                string `json:"request_bill_number" extensions:"x-order=07" example:"123456"`
	TimeRequest               string `json:"request_datetime" extensions:"x-order=08" example:"2022-02-02T12:12:12"`
	Comment                   string `json:"request_comment" extensions:"x-order=09" example:"Комментарий..."`
	Consentmailing            string `json:"consent_to_mailing" extensions:"x-order=10" example:"yes"`
	Vin                       string `json:"car_vin" extensions:"x-order=11" example:"X96A31S12P0962150"`
	UniqModCode               int    `json:"car_uniq_code" extensions:"x-order=12" example:"2372376"`
	Modification              string `json:"car_modification" extensions:"x-order=13" example:"ГАЗ-С41R33-6В"`
	ModFamily                 string `json:"car_family" extensions:"x-order=14" example:"ГАЗель NN"`
	ModBodyType               string `json:"car_body_type" extensions:"x-order=15" example:"Комби+"`
	ModEngine                 string `json:"car_engine" extensions:"x-order=16" example:"Дизель Cummins 2,8"`
	ModBase                   string `json:"car_base" extensions:"x-order=17" example:"3145"`
	ModTuning                 string `json:"car_tuning" extensions:"x-order=18" example:"Тюнинг"`
	PriceWithNds              int    `json:"car_price" extensions:"x-order=19" example:"1709800"`
	Division                  string `json:"car_division" extensions:"x-order=20" example:"lcv/mcv"`
	BrandName                 string `json:"car_brand_name" extensions:"x-order=21" example:"ГАЗ"`
	CarModel                  string `json:"car_model" extensions:"x-order=22" example:"ГАЗон NEXT Бортовая платформа (3 места) ГАЗ-С41R33-6В"`
	UrlMod                    string `json:"car_model_url" extensions:"x-order=23" example:"https://catalog.azgaz.ru/gazel-nn/kombi-gazel-nn-gaz-a32s22-420/641b1c2491cdddf86efceb38/"`
	PreviewUrl                string `json:"car_preview_url" extensions:"x-order=24" example:"https://fs.azgaz.dev.perx.ru/media/PRDiceqdYBoc6bwQEDVoHW_2"`
	СlientToken               string `json:"client_token" extensions:"x-order=25" example:"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkIjo0NywibmFtZSI6Ik1pa2UiLCJsYXN0X25hbWUiOiJvbGVnb3ZpY2giLCJzZWNvbmRfbmFtZSI6IlNoaWxvdiIsImVtYWlsIjoibUBzaGlsb3YucHJvIn0sInVpZCI6NDcsImlhdCI6MTY3OTQwNzgwOCwiZXhwIjoxNjc5NTY3ODA4fQ.b9v1PU7XybdQ9fUcFyPd22S-tENEujdiFt_KJCxyRw0"`
	TypeClient                string `json:"client_type" extensions:"x-order=26" example:"Физлицо"`
	Surname                   string `json:"client_surname" extensions:"x-order=27" example:"Иванов"`
	Name                      string `json:"client_name" extensions:"x-order=28" example:"Иван"`
	Patronymic                string `json:"client_patronymic" extensions:"x-order=29" example:"Иванович"`
	Email                     string `json:"client_email" extensions:"x-order=30" example:"i.ivanov@mail.test"`
	PhoneNumber               string `json:"client_phone_number" extensions:"x-order=31" example:"+79991234567"`
	DateOfBirth               string `json:"client_date_of_birth" extensions:"x-order=32" example:"1991-09-01"`
	PassportSer               string `json:"passport_ser" extensions:"x-order=33" example:"2288"`
	PassportNumber            string `json:"passport_number" extensions:"x-order=34" example:"199455"`
	PassportDate              string `json:"passport_date" extensions:"x-order=35" example:"2005-09-01"`                       //new
	PassportOrgan             string `json:"passport_organ" extensions:"x-order=36" example:"МВД по Автозаводском р-ну г. НН"` //new
	PassportOrganCode         string `json:"passport_organ_code" extensions:"x-order=37" example:"520-001"`                    //new
	Snils                     string `json:"snils" extensions:"x-order=38" example:"12345678910"`
	YurAddress                string `json:"client_reg_address" extensions:"x-order=9" example:"Н.Новгород, ул. Ленина, д. 1"`
	PostAddress               string `json:"post_address" extensions:"x-order=40" example:"Н.Новгород, ул. Ленина, д. 1"`
	DeliveryAddress           string `json:"delivery_address" extensions:"x-order=41" example:"Н.Новгород, ул. Ленина, д. 1"`
	DeliveryAddressCode       string `json:"delivery_address_code" extensions:"x-order=42" example:"987987987987987987987"`
	Inn                       string `json:"company_inn" extensions:"x-order=43" example:"821636535086"`
	Kpp                       string `json:"company_kpp" extensions:"x-order=44" example:"153243695"`
	Ogrn                      string `json:"company_ogrn" extensions:"x-order=45" example:"4052411237237"`
	CompanyName               string `json:"company_name" extensions:"x-order=46" example:"Авто-НН"`
	CompanyAdress             string `json:"company_adress" extensions:"x-order=47" example:"Н.Новгород, ул. Ленина, д. 1"`
	Hid                       string `json:"company_dadata_id" extensions:"x-order=48" example:"Н.Новгород, ул. Ленина, д. 1"`
	BankBik                   string `json:"bank_bik" extensions:"x-order=49" example:"044525593"`
	BankName                  string `json:"bank_name" extensions:"x-order=50" example:"АрфаБанк"`
	BankRS                    string `json:"bank_raschetniy_schet" extensions:"x-order=51" example:"40702810102130000132"`
	RepresentativeName        string `json:"representative_name" extensions:"x-order=52" example:"Иван"`
	RepresentativeSurname     string `json:"representative_surname" extensions:"x-order=53" example:"Иванов"`
	RepresentativePhoneNumber string `json:"representative_phone_number" extensions:"x-order=54" example:"+79991234567"`
	RepresentativeEmail       string `json:"representative_email" extensions:"x-order=55" example:"company@mail.ru"`
	Clientid                  string `json:"clientid_google" extensions:"x-order=56" example:"128613585.1680699283"` //Google Analytics cookies
	Ymuid                     string `json:"ClientID" extensions:"x-order=57" example:"1646307578617354501"`         //Yandex Metrics cookies
	MetricsType               string `json:"metrics_type" extensions:"x-order=58" example:"yandex"`
	СlientIP                  string `json:"client_IP" extensions:"x-order=59" example:"83.220.238.137"`
	FormName                  string `json:"form_name" extensions:"x-order=60" example:"Получить счёт"`
	FormId                    string `json:"id_form" extensions:"x-order=61" example:"635fc4247e0aed5155614771"`
	HostName                  string `json:"host_name" extensions:"x-order=62" example:"azgaz.ru"`
	File                      string `json:"file" extensions:"x-order=63" example:"ссылка на счет"`
}

type DataBookingDel struct {
	TestMod     bool   `json:"testmode"` //true - test, false - prod
	RequestId   string `json:"request_id"`
	TimeRequest string `json:"request_datetime"`
	Comment     string `json:"request_comment"`
	Vin         string `json:"car_vin"`
	UniqModCode int    `json:"car_uniq_code"`
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

type DataBookingF struct {
	TestMod           bool   `json:"testmode" extensions:"x-order=01" example:"true"`
	RequestId         string `json:"request_id" extensions:"x-order=02" example:"1019157100531"`
	SubdivisionsId    string `json:"request_subdivisions_id" extensions:"x-order=03" example:"6286424721659083063"`
	Area              string `json:"request_area" extensions:"x-order=05" example:"distrib"`
	BillNumber        string `json:"request_bill_number" extensions:"x-order=07" example:"12345678"`
	BillURL           string `json:"request_bill_link" extensions:"x-order=07" example:"https//st.tech/bill001"`
	TimeRequest       string `json:"request_datetime" extensions:"x-order=08" example:"2022-02-02T12:12:12"`
	Comment           string `json:"request_comment" extensions:"x-order=09" example:"Комментарий..."`
	Consentmailing    string `json:"consent_to_mailing" extensions:"x-order=10" example:"yes"`
	Vin               string `json:"car_vin" extensions:"x-order=11" example:"X96A31S12P0962150"`
	UniqModCode       int    `json:"car_uniq_code" extensions:"x-order=12" example:"2372376"`
	Modification      string `json:"car_modification" extensions:"x-order=13" example:"ГАЗ-С41R33-6В"`
	ModFamily         string `json:"car_family" extensions:"x-order=14" example:"ГАЗель NN"`
	ModBodyType       string `json:"car_body_type" extensions:"x-order=15" example:"Комби+"`
	PriceWithNds      int    `json:"car_price" extensions:"x-order=19" example:"1709800"`
	Division          string `json:"car_division" extensions:"x-order=20" example:"lcv/mcv"`
	BrandName         string `json:"car_brand_name" extensions:"x-order=21" example:"ГАЗ"`
	CarModel          string `json:"car_model" extensions:"x-order=22" example:"ГАЗон NEXT Бортовая платформа (3 места) ГАЗ-С41R33-6В"`
	PreviewUrl        string `json:"car_preview_url" extensions:"x-order=24" example:"https://fs.azgaz.dev.perx.ru/media/PRDiceqdYBoc6bwQEDVoHW_2"`
	СlientToken       string `json:"client_token" extensions:"x-order=25" example:"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkIjo0NywibmFtZSI6Ik1pa2UiLCJsYXN0X25hbWUiOiJvbGVnb3ZpY2giLCJzZWNvbmRfbmFtZSI6IlNoaWxvdiIsImVtYWlsIjoibUBzaGlsb3YucHJvIn0sInVpZCI6NDcsImlhdCI6MTY3OTQwNzgwOCwiZXhwIjoxNjc5NTY3ODA4fQ.b9v1PU7XybdQ9fUcFyPd22S-tENEujdiFt_KJCxyRw0"`
	Surname           string `json:"client_surname" extensions:"x-order=27" example:"Иванов"`
	Name              string `json:"client_name" extensions:"x-order=28" example:"Иван"`
	Patronymic        string `json:"client_patronymic" extensions:"x-order=29" example:"Иванович"`
	PhoneNumber       string `json:"client_phone_number" extensions:"x-order=31" example:"+79991234567"`
	Email             string `json:"client_email" extensions:"x-order=30" example:"i.ivanov@mail.test"`
	DateOfBirth       string `json:"client_date_of_birth" extensions:"x-order=32" example:"1991-09-01"`
	PassportSer       string `json:"passport_ser" extensions:"x-order=33" example:"2288"`
	PassportNumber    string `json:"passport_number" extensions:"x-order=34" example:"199455"`
	PassportDate      string `json:"passport_date" extensions:"x-order=35" example:"2005-09-01"`                       //new
	PassportOrgan     string `json:"passport_organ" extensions:"x-order=36" example:"МВД по Автозаводском р-ну г. НН"` //new
	PassportOrganCode string `json:"passport_organ_code" extensions:"x-order=37" example:"520-001"`                    //new
	Snils             string `json:"snils" extensions:"x-order=38" example:"12345678910"`
	YurAddress        string `json:"client_reg_address" extensions:"x-order=9" example:"Н.Новгород, ул. Ленина, д. 1"`
	PostAddress       string `json:"post_address" extensions:"x-order=40" example:"Н.Новгород, ул. Ленина, д. 1"`
	DeliveryAddress   string `json:"delivery_address" extensions:"x-order=41" example:"Н.Новгород, ул. Ленина, д. 1"`
	Clientid          string `json:"client_id" extensions:"x-order=56" example:"128613585.1680699283"` //Google Analytics cookies
	СlientIP          string `json:"client_ip" extensions:"x-order=59" example:"83.220.238.137"`
	FormName          string `json:"form_name" extensions:"x-order=60" example:"Получить счёт"`
	FormId            string `json:"id_form" extensions:"x-order=61" example:"635fc4247e0aed5155614771"`
	HostName          string `json:"host_name" extensions:"x-order=62" example:"azgaz.ru"`
}

type DataBookingU struct {
	TestMod                   bool   `json:"testmode" extensions:"x-order=01" example:"true"`
	RequestId                 string `json:"request_id" extensions:"x-order=02" example:"1019157100531"`
	SubdivisionsId            string `json:"request_subdivisions_id" extensions:"x-order=03" example:"6286424721659083063"`
	Area                      string `json:"request_area" extensions:"x-order=05" example:"distrib"`
	ActionType                string `json:"request_type" extensions:"x-order=06" example:"онлайн-заказ"`
	BillNumber                string `json:"request_bill_number" extensions:"x-order=07" example:"123456"`
	BillURL                   string `json:"request_bill_link" extensions:"x-order=07" example:"https//st.tech/bill001"`
	TimeRequest               string `json:"request_datetime" extensions:"x-order=08" example:"2022-02-02T12:12:12"`
	Comment                   string `json:"request_comment" extensions:"x-order=09" example:"Комментарий..."`
	Consentmailing            string `json:"consent_to_mailing" extensions:"x-order=10" example:"yes"`
	Vin                       string `json:"car_vin" extensions:"x-order=11" example:"X96A31S12P0962150"`
	UniqModCode               int    `json:"car_uniq_code" extensions:"x-order=12" example:"2372376"`
	Modification              string `json:"car_modification" extensions:"x-order=13" example:"ГАЗ-С41R33-6В"`
	ModFamily                 string `json:"car_family" extensions:"x-order=14" example:"ГАЗель NN"`
	ModBodyType               string `json:"car_body_type" extensions:"x-order=15" example:"Комби+"`
	PriceWithNds              int    `json:"car_price" extensions:"x-order=19" example:"1709800"`
	Division                  string `json:"car_division" extensions:"x-order=20" example:"lcv/mcv"`
	BrandName                 string `json:"car_brand_name" extensions:"x-order=21" example:"ГАЗ"`
	CarModel                  string `json:"car_model" extensions:"x-order=22" example:"ГАЗон NEXT Бортовая платформа (3 места) ГАЗ-С41R33-6В"`
	UrlMod                    string `json:"car_model_url" extensions:"x-order=23" example:"https://catalog.azgaz.ru/gazel-nn/kombi-gazel-nn-gaz-a32s22-420/641b1c2491cdddf86efceb38/"`
	PreviewUrl                string `json:"car_preview_url" extensions:"x-order=24" example:"https://fs.azgaz.dev.perx.ru/media/PRDiceqdYBoc6bwQEDVoHW_2"`
	СlientToken               string `json:"client_token" extensions:"x-order=25" example:"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkIjo0NywibmFtZSI6Ik1pa2UiLCJsYXN0X25hbWUiOiJvbGVnb3ZpY2giLCJzZWNvbmRfbmFtZSI6IlNoaWxvdiIsImVtYWlsIjoibUBzaGlsb3YucHJvIn0sInVpZCI6NDcsImlhdCI6MTY3OTQwNzgwOCwiZXhwIjoxNjc5NTY3ODA4fQ.b9v1PU7XybdQ9fUcFyPd22S-tENEujdiFt_KJCxyRw0"`
	Inn                       string `json:"company_inn" extensions:"x-order=43" example:"821636535086"`
	Kpp                       string `json:"company_kpp" extensions:"x-order=44" example:"153243695"`
	Ogrn                      string `json:"company_ogrn" extensions:"x-order=45" example:"4052411237237"`
	CompanyName               string `json:"company_name" extensions:"x-order=46" example:"Авто-НН"`
	CompanyAdress             string `json:"company_adress" extensions:"x-order=47" example:"Н.Новгород, ул. Ленина, д. 1"`
	PostAddress               string `json:"post_address" extensions:"x-order=40" example:"Н.Новгород, ул. Ленина, д. 1"`
	DeliveryAddress           string `json:"delivery_address" extensions:"x-order=41" example:"Н.Новгород, ул. Ленина, д. 1"`
	Hid                       string `json:"company_dadata_id" extensions:"x-order=48" example:"c7566f19aea28a24766b66ad9ef8f21c60cd8292a43338fbf9d705f88cec9a90"`
	BankBik                   string `json:"bank_bik" extensions:"x-order=49" example:"044525593"`
	BankName                  string `json:"bank_name" extensions:"x-order=50" example:"АрфаБанк"`
	BankRS                    string `json:"bank_raschetniy_schet" extensions:"x-order=51" example:"40702810102130000132"`
	RepresentativeName        string `json:"representative_name" extensions:"x-order=52" example:"Иван"`
	RepresentativeSurname     string `json:"representative_surname" extensions:"x-order=53" example:"Иванов"`
	RepresentativePhoneNumber string `json:"representative_phone_number" extensions:"x-order=54" example:"+79991234567"`
	RepresentativeEmail       string `json:"representative_email" extensions:"x-order=55" example:"company@mail.ru"`
	Clientid                  string `json:"client_id" extensions:"x-order=57" example:"1646307578617354501"` //Yandex Metrics cookies
	СlientIP                  string `json:"client_ip" extensions:"x-order=59" example:"83.220.238.137"`
	FormName                  string `json:"form_name" extensions:"x-order=60" example:"Получить счёт"`
	FormId                    string `json:"id_form" extensions:"x-order=61" example:"635fc4247e0aed5155614771"`
	HostName                  string `json:"host_name" extensions:"x-order=62" example:"azgaz.ru"`
}

// Data get
type DataStocks struct {
	VIN                            string  `json:"VIN" extensions:"x-order=a" example:"866508"`
	Площадка                       string  `json:"Площадка" extensions:"x-order=b" example:"866508"`
	Наименование_номенклатуры      string  `json:"Наименование_номенклатуры" extensions:"x-order=c" example:"866508"`
	Номер_согласно_КД              string  `json:"Номер_согласно_КД" extensions:"x-order=d" example:"866508"`
	Дивизион                       string  `json:"Дивизион" extensions:"x-order=e" example:"866508"`
	Доработчик_Подрядчик           *string `json:"Доработчик_Подрядчик" extensions:"x-order=f" example:"866508"`
	Test_truck                     bool    `json:"Test_truck" extensions:"x-order=g" example:"866508"`
	Телематика                     string  `json:"Телематика" extensions:"x-order=h" example:"866508"`
	Номер_шасси                    string  `json:"Номер_шасси" extensions:"x-order=i" example:"866508"`
	Номер_двигателя                *string `json:"Номер_двигателя" extensions:"x-order=j" example:"866508"`
	Грузоподъемность_кг            string  `json:"Грузоподъемность_кг" extensions:"x-order=k" example:"866508"`
	Цвет                           string  `json:"Цвет" extensions:"x-order=l" example:"866508"`
	ЦветИд                         string  `json:"ЦветИд" extensions:"x-order=m" example:"866508"`
	ЦветRGB                        *string `json:"ЦветRGB" extensions:"x-order=n" example:"866508"`
	Вариант_сборки                 string  `json:"Вариант_сборки" extensions:"x-order=o" example:"866508"`
	Расшифровка_варианта_сборки    string  `json:"Расшифровка_варианта_сборки" extensions:"x-order=p" example:"866508"`
	Вариант_сборки_свернутый       *string `json:"Вариант_сборки_свернутый" extensions:"x-order=q" example:"866508"`
	Год_VIN                        string  `json:"Год_VIN" extensions:"x-order=r" example:"866508"`
	Дата_сборки                    *string `json:"Дата_сборки" extensions:"x-order=s" example:"866508"`
	Справочная_стоимость_по_прайсу string  `json:"Справочная_стоимость_по_прайсу" extensions:"x-order=t" example:"866508"`
	Дата_отгрузка                  *string `json:"Дата_отгрузка" extensions:"x-order=u" example:"866508"`
	Дата_прихода                   *string `json:"Дата_прихода" extensions:"x-order=v" example:"866508"`
	Страна                         *string `json:"Страна" extensions:"x-order=w" example:"866508"`
	Контрагент_получателя          string  `json:"Контрагент_получателя" extensions:"x-order=wa" example:"866508"`
	Стоянка                        string  `json:"Стоянка" extensions:"x-order=wb" example:"866508"`
	Город_стоянки                  *string `json:"Город_стоянки" extensions:"x-order=wc" example:"866508"`
	Площадка_получателя_Ид         string  `json:"Площадка_получателя_Ид" extensions:"x-order=wd" example:"866508"`
	Контрагент_получателя_Ид       string  `json:"Контрагент_получателя_Ид" extensions:"x-order=we" example:"866508"`
	Город_стоянки_Ид               *string `json:"Город_стоянки_Ид" extensions:"x-order=wf" example:"866508"`
	Номер_заявки                   *string `json:"Номер_заявки" extensions:"x-order=wg" example:"866508"`
	Для_доработки                  *string `json:"Для_доработки" extensions:"x-order=wh" example:"866508"`
	Номерной_товар                 string  `json:"Номерной_товар" extensions:"x-order=wi" example:"866508"`
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

// data models
type DataModels struct {
	ИдМодели                 int     `json:"ИдМодели" extensions:"x-order=01" example:"866508"`
	Модель                   string  `json:"Модель" extensions:"x-order=02" example:"ГАЗ-А21R25-20"`
	ИдНабораОпций            int     `json:"ИдНабораОпций" extensions:"x-order=03" example:"123565"`
	НаборОпций               string  `json:"НаборОпций" extensions:"x-order=04" example:"1/ST(N)[2KB,5]GF,8KB,EA3/ПП"`
	ИдДивизиона              int     `json:"ИдДивизиона" extensions:"x-order=05" example:"7133"`
	Дивизион                 string  `json:"Дивизион" extensions:"x-order=06" example:"LCV"`
	ИдСемейства              *int    `json:"ИдСемейства" extensions:"x-order=07" example:"5725"`
	Семейство                *string `json:"Семейство" extensions:"x-order=08" example:"Газель Next"`
	ВидПродукции             *string `json:"ВидПродукции" extensions:"x-order=09" example:"Коммерческая техника"`
	Класс                    *string `json:"Класс" extensions:"x-order=10" example:"Бортовая платформа"`
	Кузов                    *string `json:"Кузов" extensions:"x-order=11" example:"Комби"`
	ИдНазначения             *int    `json:"ИдНазначения" extensions:"x-order=12" example:"20"`
	Назначение               *string `json:"Назначение" extensions:"x-order=13" example:"ПЕРЕВОЗКИ С ТЕМПЕРАТУРНЫМ РЕЖИМОМ"`
	ИдПроизводителя          int     `json:"ИдПроизводителя" extensions:"x-order=14" example:"60"`
	Производитель            string  `json:"Производитель" extensions:"x-order=15" example:"ГАЗ"`
	Цена                     *string `json:"Цена" extensions:"x-order=16" example:"2960000.00"`
	ПризнакЦеныОт            int     `json:"ПризнакЦеныОт" extensions:"x-order=17" example:"0"`
	СообщениеГдеНетЦен       string  `json:"СообщениеГдеНетЦен" extensions:"x-order=18" example:"Нет цены опций 1,7BA,EA3,5GF"`
	ДокументУстановившийЦену string  `json:"ДокументУстановившийЦену" extensions:"x-order=19" example:"Ввод цен (модели) № 400 от 20.02.2023 (15443251); 1 - Ввод цен (опции) № 183 от 01.09.2022 (15137904); 8KB - Ввод цен (опции) № 185 от 01.09.2022 (15137935); EA3 - Ввод цен (опции) № 185 от 01.09.2022 (15137935); 5GF - Ввод цен (опции) № 185 от 01.09.2022 (15137935); ST(N) - Ввод цен (опции) № 187 от 01.09.2022 (15137944);"`
	Сортировка               int     `json:"Сортировка" extensions:"x-order=20" example:"500"`
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
	ПолноеНаименованиеПакета       string `json:"ПолноеНаименованиеПакета" extensions:"x-order=e" example:"Пакет \"Стандарт\" для 1-ряд а/м Газель Next борт, шасси и САТ, дизель, бензин"`
	ИдГруппыОпций                  string `json:"ИдГруппыОпций" extensions:"x-order=f" example:"141"`
	КраткоеНаименованиеГруппыОпций string `json:"КраткоеНаименованиеГруппыОпций" extensions:"x-order=g" example:"2K"`
	ПолноеНаименованиеГруппыОпций  string `json:"ПолноеНаименованиеГруппыОпций" extensions:"x-order=h" example:"Сиденье водителя"`
	ИдОпции                        string `json:"ИдОпции" extensions:"x-order=i" example:"684"`
	ПолноеНаименованиеОпции        string `json:"ПолноеНаименованиеОпции" extensions:"x-order=j" example:"Сиденье водителя с подлокотником"`
	КраткоеНаименованиеОпции       string `json:"КраткоеНаименованиеОпции" extensions:"x-order=k" example:"2KB"`
}

// packets data new
type DataPackets_l struct {
	НоменклатураИд                 int    `json:"НоменклатураИд" extensions:"x-order=01" example:"865436"`
	НоменклатураНаименование       string `json:"НоменклатураНаименование" extensions:"x-order=02" example:"ГАЗ-А21R33-10"`
	ИдПакета                       int    `json:"ИдПакета" extensions:"x-order=03" example:"975"`
	КраткоеНаименованиеПакета      string `json:"КраткоеНаименованиеПакета" extensions:"x-order=04" example:"ST(N)"`
	ПолноеНаименованиеПакета       string `json:"ПолноеНаименованиеПакета" extensions:"x-order=05" example:"Пакет \"Стандарт\" для 1-ряд а/м Газель Next борт, шасси и САТ, дизель, бензин"`
	ИдГруппыОпций                  int    `json:"ИдГруппыОпций" extensions:"x-order=06" example:"141"`
	КраткоеНаименованиеГруппыОпций string `json:"КраткоеНаименованиеГруппыОпций" extensions:"x-order=07" example:"2K"`
	ПолноеНаименованиеГруппыОпций  string `json:"ПолноеНаименованиеГруппыОпций" extensions:"x-order=08" example:"Сиденье водителя"`
	ИдОпции                        int    `json:"ИдОпции" extensions:"x-order=09" example:"684"`
	ПолноеНаименованиеОпции        string `json:"ПолноеНаименованиеОпции" extensions:"x-order=10" example:"Сиденье водителя с подлокотником"`
	КраткоеНаименованиеОпции       string `json:"КраткоеНаименованиеОпции" extensions:"x-order=11" example:"2KB"`
	ЦенаПакета                     string `json:"ЦенаПакета" extensions:"x-order=12" example:"10000.00"`
}

// colors data
type DataColors struct {
	НоменклатураИд           string  `json:"НоменклатураИд" extensions:"x-order=01" example:"865436"`
	НоменклатураНаименование string  `json:"НоменклатураНаименование" extensions:"x-order=02" example:"ГАЗ-А21R33-10"`
	ЦветИд                   string  `json:"ЦветИд" extensions:"x-order=03" example:"996"`
	Наименование             string  `json:"Наименование" extensions:"x-order=04" example:"СИЛЬВЕР ЛАЙТ"`
	ПолноеНаименование       string  `json:"ПолноеНаименование" extensions:"x-order=05" example:"СВЕТЛО-СЕРЫЙ"`
	ЦветRGB                  string  `json:"ЦветRGB" extensions:"x-order=06" example:"157,163,166"`
	Слойность                *string `json:"Слойность" extensions:"x-order=07" example:"2"`
	Цена                     *string `json:"Цена" extensions:"x-order=08" example:"0"`
}

type Options struct {
	НомерСтроки                        int32   `json:"НомерСтроки" extensions:"x-order=01" example:"82"`
	Модель                             int32   `json:"Модель" extensions:"x-order=02" example:"936749"`
	БазоваяМодель                      *int32  `json:"БазоваяМодель" extensions:"x-order=02" example:"935649"`
	КодКатегории                       *string `json:"КодКатегории" extensions:"x-order=03" example:"2"`
	НаименованиеКатегории              *string `json:"НаименованиеКатегории" extensions:"x-order=04" example:"Комфорт (часть 1)"`
	ИдЗначенияОпции                    *int16  `json:"ИдЗначенияОпции" extensions:"x-order=05" example:"2413"`
	КодЗначенияОпции                   string  `json:"КодЗначенияОпции" extensions:"x-order=06" example:"2MD"`
	НаименованиеЗначенияОпции          string  `json:"НаименованиеЗначенияОпции" extensions:"x-order=07" example:"Без кондиционера"`
	ОписаниеЗначенияОпцииДляСайта      *string `json:"ОписаниеЗначенияОпцииДляСайта" extensions:"x-order=09" example:"Без системы кондиционирования для комплектации low-cost"`
	ВыбранаПоУмолчанию                 bool    `json:"ВыбранаПоУмолчанию" extensions:"x-order=11" example:"true"`
	Обязательная                       *bool   `json:"Обязательная" extensions:"x-order=12" example:"false"`
	ДопустимоТолькоОдноЗначениеВНаборе *bool   `json:"ДопустимоТолькоОдноЗначениеВНаборе" extensions:"x-order=13" example:"true"`
	ОбязательноеСочетаниеОпций         *string `json:"ОбязательноеСочетаниеОпций" extensions:"x-order=16" example:"2LF"`
	НедопустимоеСочетаниеОпций         *string `json:"НедопустимоеСочетаниеОпций" extensions:"x-order=17" example:"9PA"`
	Цена                               *string `json:"Цена" extensions:"x-order=19" example:"-55000.00"`
	ГруппаНаСайте                      *string `json:"ГруппаНаСайте" extensions:"x-order=20" example:"Оснащение"`
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

type Special struct {
	ИдМодели              int     `json:"ИдМодели" extensions:"x-order=01" example:"867560"`
	НаименованиеМодели    string  `json:"НаименованиеМодели" extensions:"x-order=02" example:"А23R25-0011-18-641-60-00-900"`
	ИдПараметра           int     `json:"ИдПараметра" extensions:"x-order=03" example:"10"`
	НаименованиеПараметра string  `json:"НаименованиеПараметра" extensions:"x-order=04" example:"Внешняя длина фургона, мм"`
	ИдЗначения            int     `json:"ИдЗначения" extensions:"x-order=05" example:"1180"`
	НаименованиеЗначения  string  `json:"НаименованиеЗначения" extensions:"x-order=06" example:"3060"`
	ПорядокСортировки     *int    `json:"ПорядокСортировки" extensions:"x-order=07" example:"7"`
	ИдГруппы              int     `json:"ИдГруппы" extensions:"x-order=08" example:"4"`
	НаименованиеГруппы    string  `json:"НаименованиеГруппы" extensions:"x-order=09" example:"Описание надстройки"`
	ИдКатегории           *int    `json:"ИдКатегории" extensions:"x-order=10" example:"6"`
	НаименованиеКатегории *string `json:"НаименованиеКатегории" extensions:"x-order=11" example:"Размеры фургона"`
}
