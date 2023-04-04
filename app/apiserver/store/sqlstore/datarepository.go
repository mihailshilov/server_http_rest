package sqlstore

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/mihailshilov/server_http_rest/app/apiserver/model"

	logger "github.com/mihailshilov/server_http_rest/app/apiserver/logger"
)

// Data repository
type DataRepository struct {
	store *Store
}

// query insert mssql
func (r *DataRepository) QueryInsertMssql(data model.DataBooking) (string, error) {

	//validation
	if err := data.ValidateDataBooking(); err != nil {
		logger.ErrorLogger.Println(err)
		return "", err
	}

	//request mssql
	var mssql_respond string

	if data.TypeClient == "Юрлицо" || data.TypeClient == "юрлицо" {
		data.Email = data.RepresentativeEmail
		data.PhoneNumber = data.RepresentativePhoneNumber
		data.TypeClient = "company"
	} else {
		data.TypeClient = "personal"
	}

	_, err := r.store.dbMssql.Exec(r.store.config.Spec.Queryies.Booking,
		sql.Named("ИдентификаторОбращения", data.RequestId),
		sql.Named("Действие", data.ActionType),
		sql.Named("НомернойТовар", data.UniqModCode),
		sql.Named("ВИН", data.Vin),
		sql.Named("ЦенаСНДС", data.PriceWithNds),
		sql.Named("ТипКонтрагента", data.TypeClient),
		sql.Named("ИНН", data.Inn),
		sql.Named("КПП", data.Kpp),
		sql.Named("ОГРН", data.Ogrn),
		sql.Named("АдресЮридический", data.YurAddress),
		sql.Named("АдресПочтовый", data.PostAddress),
		sql.Named("АдресДоставки", data.DeliveryAddress),
		sql.Named("Hid", data.Hid),
		sql.Named("Наименование", data.CompanyName),
		sql.Named("Фамилия", data.Surname),
		sql.Named("Имя", data.Name),
		sql.Named("ДатаРождения", data.DateOfBirth),
		sql.Named("Отчество", data.Patronymic),
		sql.Named("СерияПаспорта", data.PassportSer),
		sql.Named("НомерПаспорта", data.PassportNumber),
		sql.Named("СНИЛС", data.Snils),
		sql.Named("ЭлектроннаяПочта", data.Email),
		sql.Named("Телефоны", data.PhoneNumber),
		sql.Named("МоментОбращения", data.TimeRequest),
		sql.Named("НомерСчета", data.BillNumber),
		sql.Named("Ошибка", sql.Out{Dest: &mssql_respond}),
		sql.Named("ВыполнитьТестовыйВызов", data.TestMod),
	)
	if err != nil {
		return "", err
	}

	return mssql_respond, nil
}

// insert booking in postgres
func (r *DataRepository) QueryInsertBookingPostgres(data model.DataBooking) error {

	query := `
	insert into booking
	("requestid", "actiontype", "uniqmodcode", "modification", "modfamily", "modbodytype", "modengine", "modbase", "modtuning", "vin", "pricewithnds", "typeclient", "inn", "kpp", "ogrn", "yuraddress", "postaddress", "deliveryaddress", "hid", "companyname", "representativename", "representativesurname", "surname", "name", "patronymic", "passportser", "passportnumber", "snils", "dateofbirth", "email", "phonenumber", "comment", "consentmailing", "timerequest", "file", "billnumber", "urlmod", "clientid", "ymuid", "testmod")
	values($1, $2, $3, $4, $5, $6, $7, $8, $9,
		$10, $11, $12, $13, $14, $15, $16, $17, $18,
		$19, $20, $21, $22, $23, $24, $25, $26, $27,
		$28, $29, $30, $31, $32, $33, $34, $35, $36,
		$37, $38, $39, $40)`

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	tx, err := r.store.dbPostgres.Begin(context.Background())
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	_, err = tx.Exec(ctx, query,
		data.RequestId,
		data.ActionType,
		strconv.Itoa(data.UniqModCode),
		data.Modification,
		data.ModFamily,
		data.ModBodyType,
		data.ModEngine,
		data.ModBase,
		data.ModTuning,
		data.Vin,
		strconv.Itoa(data.PriceWithNds),
		data.TypeClient,
		data.Inn,
		data.Kpp,
		data.Ogrn,
		data.YurAddress,
		data.PostAddress,
		data.DeliveryAddress,
		data.Hid,
		data.CompanyName,
		data.RepresentativeName,
		data.RepresentativeSurname,
		data.Surname,
		data.Name,
		data.Patronymic,
		data.PassportSer,
		data.PassportNumber,
		data.Snils,
		data.DateOfBirth,
		data.Email,
		data.PhoneNumber,
		data.Comment,
		data.Consentmailing,
		data.TimeRequest,
		data.File,
		data.BillNumber,
		data.UrlMod,
		data.Clientid,
		data.Ymuid,
		strconv.FormatBool(data.TestMod),
	)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	return nil

}

//

// insert forms in postgres
func (r *DataRepository) QueryInsertFormsPostgres(data model.DataForms) error {

	query := `
	insert into forms
	values($1, $2, $3, $4, $5, $6, $7, $8, $9,
		$10, $11, $12, $13, $14, $15, $16, $17, $18,
		$19, $20, $21, $22, $23, $24, $25, $26, $27,
		$28, $29, $30, $31)`

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	tx, err := r.store.dbPostgres.Begin(context.Background())
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	_, err = tx.Exec(ctx, query,
		data.TimeRequest,
		data.RequestId,
		data.SubdivisionsId,
		data.SubdivisionsName,
		data.FormName,
		data.FormId,
		data.HostName,
		data.Division,
		data.Area,
		data.BrandName,
		data.CarModel,
		data.Clientid,
		data.MetricsType,
		data.СlientIP,
		data.TypeClient,
		data.CompanyName,
		data.Name,
		data.Email,
		data.PhoneNumber,
		data.Comment,
		data.Consentmailing,
		data.ActionType,
		data.Modification,
		data.ModFamily,
		data.ModBodyType,
		data.ModEngine,
		data.ModBase,
		data.ModTuning,
		data.Vin,
		data.PriceWithNds,
		data.UrlMod,
	)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	return nil

}

// insert lead_get in postgres
func (r *DataRepository) QueryInsertLeadGetPostgres(data model.DataLeadGet) error {

	valSlice := reflect.ValueOf(data).FieldByName("Data").Interface().([]model.DataLeadGet_Gazcrm)

	query := `
	insert into gazcrm_lead_get
	values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	tx, err := r.store.dbPostgres.Begin(context.Background())
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	_, err = tx.Exec(ctx, query,
		valSlice[0].TimeRequest,
		valSlice[1].EventName,
		valSlice[2].RequestId,
		valSlice[3].SubdivisionsId,
		valSlice[4].SubdivisionsName,
		valSlice[5].FormName,
		valSlice[6].HostName,
		valSlice[7].Division,
		valSlice[8].Area,
		valSlice[9].BrandName,
		valSlice[10].ClientID,
		valSlice[11].MetricsType,
	)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	return nil
}

// insert work lists in postgres
func (r *DataRepository) QueryInsertWorkListsPostgres(data model.DataWorkList) error {

	valSlice := reflect.ValueOf(data).FieldByName("Data").Interface().([]model.DataWorkList_Gazcrm)

	query := `
	insert into gazcrm_work_list
	values($1, $2, $3, $4)`

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	tx, err := r.store.dbPostgres.Begin(context.Background())
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	_, err = tx.Exec(ctx, query,
		valSlice[0].TimeRequest,
		valSlice[1].EventName,
		valSlice[2].GazcrmClientId,
		valSlice[3].GazCrmWorkListId,
	)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	return nil
}

// insert work lists in postgres
func (r *DataRepository) QueryInsertStatusesPostgres(data model.DataStatuses) error {

	valSlice := reflect.ValueOf(data).FieldByName("Data").Interface().([]model.DataStatuses_Gazcrm)

	query := `
	insert into gazcrm_statuses
	values($1, $2, $3, $4, $5, $6, $7)`

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	tx, err := r.store.dbPostgres.Begin(context.Background())
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	_, err = tx.Exec(ctx, query,
		valSlice[0].TimeRequest,
		valSlice[1].EventName,
		valSlice[2].RequestId,
		valSlice[3].GazcrmClientId,
		valSlice[4].GazCrmWorkListId,
		valSlice[5].ClientID,
		valSlice[6].MetricsType,
	)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	return nil
}

// query stocks mssql
func (r *DataRepository) QueryStocksMssql() ([]model.DataStocks, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.Stocks)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataStocks{}

	for rows.Next() {

		data := &model.DataStocks{}

		err := rows.Scan(
			&data.VIN,
			&data.Площадка,
			&data.Наименование_номенклатуры,
			&data.Номер_согласно_КД,
			&data.Дивизион,
			&data.Доработчик_Подрядчик,
			&data.Test_truck,
			&data.Телематика,
			&data.Номер_шасси,
			&data.Номер_двигателя,
			&data.Грузоподъемность_кг,
			&data.Цвет,
			&data.ЦветИд,
			&data.ЦветRGB,
			&data.Вариант_сборки,
			&data.Расшифровка_варианта_сборки,
			&data.Вариант_сборки_свернутый,
			&data.Год_VIN,
			&data.Дата_сборки,
			&data.Справочная_стоимость_по_прайсу,
			&data.Дата_отгрузка,
			&data.Дата_прихода,
			&data.Страна,
			&data.Контрагент_получателя,
			&data.Стоянка,
			&data.Город_стоянки,
			&data.Площадка_получателя_Ид,
			&data.Контрагент_получателя_Ид,
			&data.Город_стоянки_Ид,
			&data.Номер_заявки,
			&data.Для_доработки,
			&data.Номерной_товар,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}

		//results = append(results, *data)

		// if data.Наименование_номенклатуры == "ГАЗ-А21S12-225" || data.Наименование_номенклатуры == "А23S12-1221-18-А83-60-00-900" {

		// 	logger.InfoLogger.Println("не выгружаем: " + data.Наименование_номенклатуры)

		// } else {
		// 	results = append(results, *data)
		// }

		switch data.Наименование_номенклатуры {
		case "ГАЗ-А21S12-225", "А23S12-1221-18-А83-60-00-900", "А31S12-0420-68-216-66-00-000", "А31S12-0420-68-218-66-00-000", "С41А23-1020-35-581-77-00-000":
			logger.InfoLogger.Println("не выгружаем: " + data.Наименование_номенклатуры)
		default:
			results = append(results, *data)
		}

	}

	return results, nil

}

// query mssql price basic models
func (r *DataRepository) QueryBasicModelsPriceMssql() ([]model.DataBasicModelsPrice, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.BasicModelsPrice)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataBasicModelsPrice{}

	for rows.Next() {

		data := &model.DataBasicModelsPrice{}

		err := rows.Scan(
			&data.Товар,
			&data.Цена,
			&data.СтавкаНДС,
			&data.НДС,
			&data.НачалоДействия,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// query mssql price2 basic models
func (r *DataRepository) QueryBasicModelsPriceMssql2() ([]model.DataBasicModelsPrice, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.BasicModelsPrice2)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataBasicModelsPrice{}

	for rows.Next() {

		data := &model.DataBasicModelsPrice{}

		err := rows.Scan(
			&data.Товар,
			&data.Цена,
			&data.СтавкаНДС,
			&data.НДС,
			&data.НачалоДействия,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// query mssql options price
func (r *DataRepository) QueryOptionsPriceMssql() ([]model.DataOptionsPrice, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.OptionsPrice)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataOptionsPrice{}

	for rows.Next() {

		data := &model.DataOptionsPrice{}

		err := rows.Scan(
			&data.ЕНСП_Модификация_Ид,
			&data.Товар,
			&data.ТоварИд,
			&data.ЗначениеОпции,
			&data.ЗначениеОпцииИд,
			&data.ОбозначениеОпции,
			&data.Цена,
			&data.СтавкаНДС_Ид,
			&data.НДС,
			&data.НачалоДействия,
			&data.СоставПакета,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// query mssql price general
func (r *DataRepository) QueryGeneralPriceMssql() ([]model.DataGeneralPrice, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.GeneralPrice)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataGeneralPrice{}

	for rows.Next() {

		data := &model.DataGeneralPrice{}

		err := rows.Scan(
			&data.Товар,
			&data.ВариантСборки,
			&data.ВариантСборкиРазвернутый,
			&data.Цена,
			&data.СтавкаНДС,
			&data.НДС,
			&data.НачалоДействия,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// query mssql sprav
func (r *DataRepository) QuerySprav() ([]model.DataSprav, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.Sprav)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataSprav{}

	for rows.Next() {

		data := &model.DataSprav{}

		err := rows.Scan(
			&data.Наименование,
			&data.НомерСогласноКД,
			&data.Дивизион,
			&data.СтатусМоделиВПроизводстве,
			&data.МассаСнагрузкой,
			&data.МассаБезНагрузки,
			&data.ОписаниеДляПрайса,
			&data.База,
			&data.БазаАвтомобиляДлина,
			&data.ТипКузова,
			&data.ТипФургона,
			&data.ОбозначениеДвигателя,
			&data.ОбъемДвигателя,
			&data.ВидТоплива,
			&data.СтабилизаторЗаднейПодвески,
			&data.ГорныйТормоз,
			&data.ТормознаяСистемаТип,
			&data.ЦветаДопустимыеВЭтомМесяце,
			&data.ОпцииДопустимыеВЭтомМесяце,
			&data.ОпцииПоУмолчанию,
			&data.ЧислоПосадочныхМест,
			&data.ЭкКласс,
			&data.Привод,
			&data.Семейство,
			&data.Лебедка,
			&data.КПП,
			&data.ГБО,
			&data.Надстройка,
			&data.ОсобенностьНадстройки,
			&data.БазовыйТовар,
			&data.ОпцииАЗ,
			&data.ХарактеристикиНоменклатуры,
			&data.ИзЭПТС_ДопустимаяМаксимальнаяМассаСтандарт,
			&data.ИзЭПТС_ДопустимаяМаксимальнаяМасса9РА,
			&data.ИзЭПТС_ДопустимаяМаксимальнаяМасса9РВ,
			&data.ИзЭПТС_СнаряженнаяМасса,
			&data.ДоступностьКЗаказу,
		)

		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}

		if data.МассаСнагрузкой == "nil" {
			data.МассаСнагрузкой = "nil"
		}

		// if data.Наименование == "ГАЗ-А31S22-420" || data.Наименование == "ГАЗ-А32S22-420" {
		// 	data.Семейство = "Соболь NN"
		// }

		if data.ИзЭПТС_ДопустимаяМаксимальнаяМассаСтандарт == nil {
			stringone := "nil" //тут можно добавить любое значение
			stringtwo := &stringone
			data.ИзЭПТС_ДопустимаяМаксимальнаяМассаСтандарт = stringtwo
		}

		//*&data
		/*
			blacklist := []string{
				"ГАЗ-А21S12-225",
				"А23S12-1221-18-А83-60-00-900",
			}

			for _, id_nom := range blacklist {

				if data.Наименование != id_nom {

					//results = results
					results = append(results, *data)

				} else {

					logger.InfoLogger.Println("не выгружаем: " + data.Наименование)

				}

			}
		*/

		//А31S12-0420-68-218-66-00-000

		switch data.Наименование {
		case "ГАЗ-А21S12-225", "А23S12-1221-18-А83-60-00-900", "А31S12-0420-68-216-66-00-000", "А31S12-0420-68-218-66-00-000", "С41А23-1020-35-581-77-00-000":
			logger.InfoLogger.Println("не выгружаем: " + data.Наименование)
		default:
			results = append(results, *data)
		}

		/*
			if data.Наименование == "ГАЗ-А21S12-225" || data.Наименование == "А23S12-1221-18-А83-60-00-900" {

				//results = results

				logger.InfoLogger.Println("не выгружаем: " + data.Наименование)

			} else {

			}
		*/

	}

	return results, nil

}

// query mssql options data
func (r *DataRepository) QueryOptionsData() ([]model.DataOptions, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.Options)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataOptions{}

	for rows.Next() {

		data := &model.DataOptions{}

		err := rows.Scan(
			&data.НоменклатураИд,
			&data.НоменклатураНаименование,
			&data.ИдГруппыОпций,
			&data.КраткоеНаименованиеГруппыОпций,
			&data.ПолноеНаименованиеГруппыОпций,
			&data.ОпцияИд,
			&data.КраткоеНаименованиеОпции,
			&data.ПолноеНаименованиеОпции,
			&data.ОписаниеОпции,
			&data.ЦенаОпции,
			&data.Обязательная,
			&data.ВыбранаПоУмолчанию,
			&data.ЭтоПакет,
			&data.ЦенаНулл,
		)

		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// query mssql options sprav data
func (r *DataRepository) QueryOptionsDataSprav() ([]model.DataOptionsSprav, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.OptionsSprav)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataOptionsSprav{} // creating empty slice

	for rows.Next() {

		data := &model.DataOptionsSprav{} // creating new struct for every row

		err := rows.Scan(
			&data.НоменклатураИд,
			&data.НоменклатураНаименование,
			&data.ЗначениеОпции1,
			&data.ЗначениеОпции2,
			&data.КодОпции1,
			&data.КодОпции2,
			&data.ВидСочетания,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// query mssql packets data
func (r *DataRepository) QueryPacketsData() ([]model.DataPackets, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.Packets)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataPackets{} // creating empty slice

	for rows.Next() {

		data := &model.DataPackets{} // creating new struct for every row

		err := rows.Scan(
			&data.НоменклатураИд,
			&data.НоменклатураНаименование,
			&data.ИдПакета,
			&data.КраткоеНаименованиеПакета,
			&data.ПолноеНаименованиеПакета,
			&data.ИдГруппыОпций,
			&data.КраткоеНаименованиеГруппыОпций,
			&data.ПолноеНаименованиеГруппыОпций,
			&data.ИдОпции,
			&data.ПолноеНаименованиеОпции,
			&data.КраткоеНаименованиеОпции,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// query mssql colors data
func (r *DataRepository) QueryColorsData() ([]model.DataColors, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.Colors)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataColors{} // creating empty slice

	for rows.Next() {

		data := &model.DataColors{} // creating new struct for every row

		err := rows.Scan(
			&data.НоменклатураИд,
			&data.НоменклатураНаименование,
			&data.ЦветИд,
			&data.Наименование,
			&data.ПолноеНаименование,
			&data.ЦветRGB,
			&data.Слойность,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// query mssql statuses data
func (r *DataRepository) QueryStatusesLkData() ([]model.DataStatusesLk, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.Statuses)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataStatusesLk{} // creating empty slice

	for rows.Next() {

		data := &model.DataStatusesLk{} // creating new struct for every row

		err := rows.Scan(
			&data.ИдЗаказа,
			&data.СтатусЗаказа,
			&data.НомернойТовар,
			&data.ВИН,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// query mssql sprav data
func (r *DataRepository) QuerySpravModels() ([]model.DataSpravModels, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.Sprav_new)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer rows.Close()

	results := []model.DataSpravModels{} // creating empty slice

	for rows.Next() {

		data := &model.DataSpravModels{} // creating new struct for every row

		err := rows.Scan(
			&data.Ид,
			&data.Наименование,
			&data.Дивизион,
			&data.СтатусМоделиВПроизводстве,
			&data.БазовыйТовар,
			&data.ХарактеристикиНоменклатуры,
			&data.Цена,
			&data.СтавкаНДС,
			&data.НДС,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		results = append(results, *data)
	}

	return results, nil

}

// call microservice mailing
func (r *DataRepository) CallMSMailing(data model.DataBooking, config *model.Service) (string, error) {

	bodyBytesReq, err := json.Marshal(data)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", err
	}

	resp, err := http.Post(config.Spec.Client.UrlMailingService, "application/json", bytes.NewBuffer(bodyBytesReq))
	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	bodyBytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytesResp), nil

}

// request GAZ CRM booking
func (r *DataRepository) RequestGazCrmApiBooking(data model.DataBooking, config *model.Service) (*model.ResponseGazCrm, error) {

	var dataset model.DataGazCrm
	var response *model.ResponseGazCrm

	if data.TypeClient == "Физлицо" || data.TypeClient == "физлицо" {
		data.TypeClient = "personal"
	} else {
		data.TypeClient = "company"
		data.Email = data.RepresentativeEmail
		data.PhoneNumber = data.RepresentativePhoneNumber
	}

	bodyJson0 := &model.DataGazCrmReq{
		TimeRequest: data.TimeRequest,
	}
	bodyJson1 := &model.DataGazCrmReq{
		RequestId: data.RequestId,
	}
	bodyJson2 := &model.DataGazCrmReq{
		SubdivisionsId: data.SubdivisionsId,
	}
	bodyJson3 := &model.DataGazCrmReq{
		SubdivisionsName: data.SubdivisionsName,
	}
	bodyJson4 := &model.DataGazCrmReq{
		FormName: data.FormName,
	}
	bodyJson5 := &model.DataGazCrmReq{
		FormId: data.FormId,
	}
	bodyJson6 := &model.DataGazCrmReq{
		HostName: data.HostName,
	}
	bodyJson7 := &model.DataGazCrmReq{
		Division: data.Division,
	}
	bodyJson8 := &model.DataGazCrmReq{
		Area: data.Area,
	}
	bodyJson9 := &model.DataGazCrmReq{
		BrandName: data.BrandName,
	}
	bodyJson10 := &model.DataGazCrmReq{
		CarModel: data.CarModel,
	}
	bodyJson11 := &model.DataGazCrmReq{
		ClientID: data.Clientid,
	}
	bodyJson12 := &model.DataGazCrmReq{
		MetricsType: data.MetricsType,
	}
	bodyJson13 := &model.DataGazCrmReq{
		СlientIP: data.СlientIP,
	}

	bodyJson14 := &model.DataGazCrmReq{
		TypeClient: data.TypeClient,
	}

	bodyJson15 := &model.DataGazCrmReq{
		CompanyName: data.CompanyName,
	}
	bodyJson16 := &model.DataGazCrmReq{
		СlientName: data.Name,
	}
	bodyJson17 := &model.DataGazCrmReq{
		ClientEmail: data.Email,
	}
	bodyJson18 := &model.DataGazCrmReq{
		ClientPhoneNumber: data.PhoneNumber,
	}
	bodyJson19 := &model.DataGazCrmReq{
		Commentary: data.Comment,
	}
	bodyJson20 := &model.DataGazCrmReq{
		AgreementMailing: data.Consentmailing,
	}

	dataset.Data = append(dataset.Data, bodyJson0, bodyJson1, bodyJson2, bodyJson3, bodyJson4, bodyJson5, bodyJson6,
		bodyJson7, bodyJson8, bodyJson9, bodyJson10, bodyJson11, bodyJson12, bodyJson13, bodyJson14,
		bodyJson15, bodyJson16, bodyJson17, bodyJson18, bodyJson19, bodyJson20)

	//d_spaces, err := json.MarshalIndent(dataset, "", "    ")
	//if err != nil {
	//return "", err
	//}

	bodyBytesReq, err := json.Marshal(dataset)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	resp, err := http.Post(config.Spec.Client.UrlGazCrmTest, "application/json", bytes.NewBuffer(bodyBytesReq))
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	defer resp.Body.Close()

	bodyBytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	if err := json.Unmarshal(bodyBytesResp, &response); err != nil {
		return nil, err
	}

	return response, nil

}

// request GAZ CRM forms
func (r *DataRepository) RequestGazCrmApiForms(data model.DataForms, config *model.Service) (*model.ResponseGazCrm, error) {

	var dataset model.DataGazCrm
	var response *model.ResponseGazCrm

	bodyJson0 := &model.DataGazCrmReq{
		TimeRequest: data.TimeRequest,
	}
	bodyJson1 := &model.DataGazCrmReq{
		RequestId: data.RequestId,
	}
	bodyJson2 := &model.DataGazCrmReq{
		SubdivisionsId: data.SubdivisionsId,
	}
	bodyJson3 := &model.DataGazCrmReq{
		SubdivisionsName: data.SubdivisionsName,
	}
	bodyJson4 := &model.DataGazCrmReq{
		FormName: data.FormName,
	}
	bodyJson5 := &model.DataGazCrmReq{
		FormId: data.FormId,
	}
	bodyJson6 := &model.DataGazCrmReq{
		HostName: data.HostName,
	}
	bodyJson7 := &model.DataGazCrmReq{
		Division: data.Division,
	}
	bodyJson8 := &model.DataGazCrmReq{
		Area: data.Area,
	}
	bodyJson9 := &model.DataGazCrmReq{
		BrandName: data.BrandName,
	}
	bodyJson10 := &model.DataGazCrmReq{
		CarModel: data.CarModel,
	}
	bodyJson11 := &model.DataGazCrmReq{
		ClientID: data.Clientid,
	}
	bodyJson12 := &model.DataGazCrmReq{
		MetricsType: data.MetricsType,
	}
	bodyJson13 := &model.DataGazCrmReq{
		СlientIP: data.СlientIP,
	}
	bodyJson14 := &model.DataGazCrmReq{
		TypeClient: data.TypeClient,
	}
	bodyJson15 := &model.DataGazCrmReq{
		CompanyName: data.CompanyName,
	}
	bodyJson16 := &model.DataGazCrmReq{
		СlientName: data.Name,
	}
	bodyJson17 := &model.DataGazCrmReq{
		ClientEmail: data.Email,
	}
	bodyJson18 := &model.DataGazCrmReq{
		ClientPhoneNumber: data.PhoneNumber,
	}
	bodyJson19 := &model.DataGazCrmReq{
		Commentary: data.Comment,
	}
	bodyJson20 := &model.DataGazCrmReq{
		AgreementMailing: data.Consentmailing,
	}

	dataset.Data = append(dataset.Data, bodyJson0, bodyJson1, bodyJson2, bodyJson3, bodyJson4, bodyJson5, bodyJson6,
		bodyJson7, bodyJson8, bodyJson9, bodyJson10, bodyJson11, bodyJson12, bodyJson13, bodyJson14,
		bodyJson15, bodyJson16, bodyJson17, bodyJson18, bodyJson19, bodyJson20)

	bodyBytesReq, err := json.Marshal(dataset)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(config.Spec.Client.UrlGazCrmTest, "application/json", bytes.NewBuffer(bodyBytesReq))
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	if err := json.Unmarshal(bodyBytesResp, &response); err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	return response, nil

}

// request LK order
func (r *DataRepository) RequestLkOrder(data model.DataBooking, config *model.Service) (*http.Response, error) {

	dataset := &model.DataLkOrder{
		Id:          data.RequestId,
		StageCode:   "new",
		Title:       data.ModFamily + " " + data.ModBodyType,
		Vin:         data.Vin,
		Cost:        strconv.Itoa(data.PriceWithNds),
		ModelFamily: data.ModFamily,
		PreviewUrl:  data.PreviewUrl,
	}

	var ApiUrl string

	if data.TestMod == true {
		ApiUrl = config.Spec.Client.UrlLkOrderTest
	} else {
		ApiUrl = config.Spec.Client.UrlLkOrder
	}

	client := &http.Client{}

	token := data.СlientToken

	bearer := "Bearer " + token

	bodyBytesReq, err := json.Marshal(dataset)
	if err != nil {
		return nil, err
	}

	logger.InfoLogger.Println(bytes.NewBuffer(bodyBytesReq))

	//resp, err := http.Post(config.Spec.Client.UrlLkOrder, "application/json", bytes.NewBuffer(bodyBytesReq))
	req, err := http.NewRequest(http.MethodPost, ApiUrl, bytes.NewBuffer(bodyBytesReq)) // URL-encoded payload
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	logger.ErrorLogger.Println(req)

	return response, nil

}

// query tech data
func (r *DataRepository) QueryTechData() (*[]model.TechDataObj, error) {

	rows, err := r.store.dbMssql.Query(r.store.config.Spec.Queryies.Techdata)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	//result := model.TechData{}

	var result string
	var response *[]model.TechDataObj

	for rows.Next() {

		data := "" // creating new struct for every row

		err := rows.Scan(
			&data,
		)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, err
		}
		result = result + data
	}

	//result_unq, err := json.Unmarshal([]byte(result), []model.TechDataObj)
	if err := json.Unmarshal([]byte(result), &response); err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	//result1 := model.TechData{Data: result_unq}

	defer rows.Close()

	return response, nil

}
