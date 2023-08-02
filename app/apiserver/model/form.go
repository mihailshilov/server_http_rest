package model

type Form struct {
	TestMod          bool   `json:"testmode" extensions:"x-order=01" example:"true"`
	RequestId        string `json:"request_id" extensions:"x-order=02" example:"1019157100531"`
	SubdivisionsId   string `json:"request_subdivisions_id" extensions:"x-order=03" example:"6286424721659083063"`
	SubdivisionsName string `json:"request_subdivisions_name" extensions:"x-order=04" example:"Авторитейл М"`
	Area             string `json:"request_area" extensions:"x-order=05" example:"dealer"`
	TimeRequest      string `json:"request_datetime" extensions:"x-order=06" example:"2022-02-02T12:12:12"`
	Comment          string `json:"request_comment" extensions:"x-order=07" example:"Комментарий..."`
	FormName         string `json:"form_name" extensions:"x-order=08" example:"Получить счёт"`
	FormId           string `json:"id_form" extensions:"x-order=09" example:"635fc4247e0aed5155614771"`
	Division         string `json:"car_division" extensions:"x-order=10" example:"lcv/mcv"`
	BrandName        string `json:"car_brand_name" extensions:"x-order=11" example:"ГАЗ"`
	CarModel         string `json:"car_model" extensions:"x-order=12" example:"ГАЗон NEXT Бортовая платформа (3 места) ГАЗ-С41R33-6В"`
	Modification     string `json:"car_modification" extensions:"x-order=13" example:"ГАЗ-С41R33-6В"`
	ModFamily        string `json:"car_family" extensions:"x-order=14" example:"ГАЗель NN"`
	ModBodyType      string `json:"car_body_type" extensions:"x-order=15" example:"Комби+"`
	Clientid         string `json:"client_id" extensions:"x-order=16" example:"1646307578617354501"` //Yandex Metrics cookies
	СlientIP         string `json:"client_ip" extensions:"x-order=17" example:"83.220.238.137"`
	TypeClient       string `json:"client_type" extensions:"x-order=18" example:"personal"`
	CompanyName      string `json:"company_name" extensions:"x-order=19" example:"Авто-НН"`
	Name             string `json:"client_name" extensions:"x-order=20" example:"Иван"`
	Email            string `json:"client_email" extensions:"x-order=21" example:"i.ivanov@mail.test"`
	PhoneNumber      string `json:"client_phone_number" extensions:"x-order=22" example:"+79991234567"`
	Consentmailing   string `json:"consent_to_mailing" extensions:"x-order=23" example:"yes"`
	HostName         string `json:"host_name" extensions:"x-order=24" example:"azgaz.ru"`
}
