package sqlstore

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mihailshilov/server_http_rest/app/apiserver/model"

	logger "github.com/mihailshilov/server_http_rest/app/apiserver/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

// request GAZ CRM forms
func (r *DataRepository) RequestForms(data model.Form, config *model.Service) (*model.ResponseGazCrm, error) {

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
		MetricsType: "Yandex",
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

	var CrmApiUrl = config.Spec.Client.UrlGazCrm
	if data.TestMod == true {
		CrmApiUrl = config.Spec.Client.UrlGazCrmTest
	}

	resp, err := http.Post(CrmApiUrl, "application/json", bytes.NewBuffer(bodyBytesReq))
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

// insert forms in postgres
func (r *DataRepository) QueryFormsPostgres(data model.Form, CrmDone bool) error {

	query := `
	insert into forms_new ("testmode", "request_id", "subdivisions_id", "subdivisions_name", "request_area", "request_datetime", "request_comment", "form_name", "form_id", "car_division", "car_brand_name", "car_model", "car_modification", "car_family", "car_body_type", "client_id", "client_ip", "client_type", "company_name", "client_name", "client_email", "client_phone_number", "consent_to_mailing", "host_name", "crm_done")
	values($1, $2, $3, $4, $5, $6, $7, $8, $9,
		$10, $11, $12, $13, $14, $15, $16, $17, $18,
		$19, $20, $21, $22, $23, $24, $25)`

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	tx, err := r.store.dbPostgres.Begin(context.Background())
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	_, err = tx.Exec(ctx, query,
		data.TestMod,
		data.RequestId,
		data.SubdivisionsId,
		data.SubdivisionsName,
		data.Area,
		data.TimeRequest,
		data.Comment,
		data.FormName,
		data.FormId,
		data.Division,
		data.BrandName,
		data.CarModel,
		data.Modification,
		data.ModFamily,
		data.ModBodyType,
		data.Clientid,
		data.СlientIP,
		data.TypeClient,
		data.CompanyName,
		data.Name,
		data.Email,
		data.PhoneNumber,
		data.Consentmailing,
		data.HostName,
		CrmDone,
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

func (r *DataRepository) AddToRabbit(data []byte, queue string, exchange string) error {

	conn, err := amqp.Dial("amqp://rmuser:Gjcnvjlthybpv$@onsales.st.tech:5672/")
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	defer ch.Close()

	// q, err := ch.QueueDeclare(
	// 	queue, // name
	// 	true,  // durable
	// 	false, // delete when unused
	// 	false, // exclusive
	// 	false, // no-wait
	// 	args,   // arguments
	// )
	// if err != nil {
	// 	logger.ErrorLogger.Println(err)
	// 	return err
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		exchange, // exchange
		queue,    // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})

	return nil
}
