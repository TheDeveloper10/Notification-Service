package repository

import (
	"notification-service/internal/helper"
	"strconv"

	log "github.com/sirupsen/logrus"

	"notification-service/internal/clients"
	"notification-service/internal/entity"
)

type TemplateRepository interface {
	Insert(entity *entity.TemplateEntity) int64
	Get(id int) 						  (*entity.TemplateEntity, int)
	Update(entity *entity.TemplateEntity) int
	Delete(id int) 						  bool
}

type basicTemplateRepository struct { }

func NewTemplateRepository() TemplateRepository {
	return &basicTemplateRepository{}
}

func (btr *basicTemplateRepository) Insert(entity *entity.TemplateEntity) int64 {
	stmt, err1 := clients.SQLClient.Prepare("insert into Templates(ContactType, Template, Language, Type) values(?, ?, ?, ?)")
	if helper.IsError(err1) {
		return -1
	}
	defer helper.HandledClose(stmt)

	res, err2 := stmt.Exec(entity.ContactType, entity.Template, entity.Language, entity.Type)
	if helper.IsError(err2) {
		return -1
	}
	id, err3 := res.LastInsertId()
	if helper.IsError(err3) {
		return -1
	}
	log.Info("Inserted template with id " + strconv.FormatInt(id, 10))
	return id
}

func (btr *basicTemplateRepository) Get(id int) (*entity.TemplateEntity, int) {
	stmt, err1 := clients.SQLClient.Prepare("select * from Templates where Id=?")
	if helper.IsError(err1) {
		return nil, 1
	}
	defer helper.HandledClose(stmt)

	rows, err2 := stmt.Query(id)
	if helper.IsError(err2) {
		return nil, 1
	}
	defer helper.HandledClose(rows)

	if rows.Next() {
		record := entity.TemplateEntity{}
		err3 := rows.Scan(&record.Id, &record.ContactType, &record.Template, &record.Language, &record.Type)
		if helper.IsError(err3) {
			return nil, 2
		}
		log.Info("Fetched template with id " + strconv.Itoa(id))
		return &record, 0
	} else {
		log.Warn("Template with id " + strconv.Itoa(id) + " was not found")
		return nil, 2
	}
}

func (btr *basicTemplateRepository) Update(entity *entity.TemplateEntity) int {
	stmt, err1 := clients.SQLClient.Prepare("update Templates set Template=?, ContactType=?, Language=?, Type=? where Id=?")
	if helper.IsError(err1) {
		return 1
	}
	defer helper.HandledClose(stmt)

	res, err2 := stmt.Exec(entity.Template, entity.ContactType, entity.Language, entity.Type, entity.Id)
	if helper.IsError(err2) {
		return 1
	}
	affectedRows, err3 := res.RowsAffected() 
	if helper.IsError(err3) {
		return 1
	}

	// TODO: it's zero also when template is found but the value you set is the same** FIX IT
	if affectedRows <= 0 {
		log.Warn("No template was found with id " + strconv.Itoa(entity.Id))
		return 2
	}

	log.Info("Updated template with id " + strconv.Itoa(entity.Id))
	return 0
}

func (btr *basicTemplateRepository) Delete(id int) bool {
	stmt, err1 := clients.SQLClient.Prepare("delete from Templates where Id=?")
	if helper.IsError(err1) {
		return false
	}
	defer helper.HandledClose(stmt)

	_, err2 := stmt.Exec(id)
	if helper.IsError(err2) {
		return false
	}
	log.Info("Deleted template with id " + strconv.Itoa(id))
	return true
}
