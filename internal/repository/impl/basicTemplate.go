package impl

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"notification-service/internal/client"
	entity2 "notification-service/internal/data/entity"
	"notification-service/internal/helper"
	"notification-service/internal/util"
	"strconv"
)

type BasicTemplateRepository struct {
	cache map[int]*entity2.TemplateEntity
}

func (btr *BasicTemplateRepository) Init() {
	btr.cache = map[int]*entity2.TemplateEntity{}
}

func (btr *BasicTemplateRepository) Insert(entity *entity2.TemplateEntity) (int, util.RepoStatusCode) {
	res := client.SQLClient.Exec(
		"insert into Templates(EmailBody, SMSBody, PushBody, Language, Type) values(?, ?, ?, ?, ?)",
		entity.Body.Email, entity.Body.SMS, entity.Body.Push, entity.Language, entity.Type,
	)
	if res == nil {
		return -1, util.RepoStatusError
	}
	id, err3 := res.LastInsertId()
	if helper.IsError(err3) {
		return -1, util.RepoStatusError
	}

	logrus.Info("Inserted template into the database with id " + strconv.FormatInt(id, 10))
	btr.cache[int(id)] = entity
	return int(id), util.RepoStatusSuccess
}

func (btr *BasicTemplateRepository) Get(id int) (*entity2.TemplateEntity, util.RepoStatusCode) {
	if result, ok := btr.cache[id]; ok {
		return result, util.RepoStatusSuccess
	}

	rows := client.SQLClient.Query("select * from Templates where Id=?", id)
	if rows == nil {
		return nil, util.RepoStatusError
	}
	defer helper.HandledClose(rows)

	if rows.Next() {
		record := btr.GetTemplateEntityFromSQLRows(rows)
		if record == nil {
			return nil, util.RepoStatusNotFound
		}

		logrus.Info("Fetched template with id " + strconv.Itoa(id))

		go btr.clearCache()
		btr.cache[id] = record
		return record, util.RepoStatusSuccess
	} else {
		logrus.Warn("Template with id " + strconv.Itoa(id) + " was not found")
		return nil, util.RepoStatusNotFound
	}
}

func (btr *BasicTemplateRepository) GetBulk(filter *entity2.TemplateFilter) (*[]entity2.TemplateEntity, util.RepoStatusCode) {
	builder := util.NewQueryBuilder("select * from Templates")

	offset := (filter.Page - 1) * filter.Size
	query := builder.End(&filter.Size, &offset)

	rows := client.SQLClient.Query(*query)
	if rows == nil {
		return nil, util.RepoStatusError
	}
	defer helper.HandledClose(rows)

	var templates []entity2.TemplateEntity
	for rows.Next() {
		record := btr.GetTemplateEntityFromSQLRows(rows)
		if record == nil {
			return nil, util.RepoStatusError
		}

		templates = append(templates, *record)
	}

	logrus.Info("Fetched " + strconv.Itoa(len(templates)) + " template(s)")
	return &templates, util.RepoStatusSuccess
}

func (btr *BasicTemplateRepository) GetTemplateEntityFromSQLRows(rows *sql.Rows)  *entity2.TemplateEntity {
	record := entity2.TemplateEntity{}
	record.Body = entity2.TemplateBody{}
	email := ""
	sms := ""
	push := ""
	err3 := rows.Scan(&record.Id, &email, &sms, &push, &record.Language, &record.Type)
	if helper.IsError(err3) {
		return nil
	}
	if email != "" {
		record.Body.Email = &email
	}
	if sms != "" {
		record.Body.SMS = &sms
	}
	if push != "" {
		record.Body.Push = &push
	}
	return &record
}

func (btr *BasicTemplateRepository) Update(entity *entity2.TemplateEntity) util.RepoStatusCode {
	res := client.SQLClient.Exec(
		"update Templates set EmailBody=?, SMSBody=?, PushBody=?, Language=?, Type=? where Id=?",
		entity.Body.Email, entity.Body.SMS, entity.Body.Push, entity.Language, entity.Type,
	)
	if res == nil {
		return util.RepoStatusError
	}

	logrus.Info("Updated template in the database with id " + strconv.Itoa(entity.Id))
	btr.cache[entity.Id] = entity
	return util.RepoStatusSuccess
}

func (btr *BasicTemplateRepository) Delete(id int) util.RepoStatusCode {
	res := client.SQLClient.Exec("delete from Templates where Id=?", id)
	if res == nil {
		return util.RepoStatusError
	}

	logrus.Info("Deleted template from the database with id " + strconv.Itoa(id))
	delete(btr.cache, id)
	return util.RepoStatusSuccess
}

func (btr *BasicTemplateRepository) clearCache() {
	cacheSize := len(btr.cache)
	maxCacheSize := helper.Config.Service.TemplateCacheSize
	if cacheSize < maxCacheSize {
		return
	}

	for key := range btr.cache {
		if cacheSize < maxCacheSize {
			return
		}

		delete(btr.cache, key)
		cacheSize--
		break
	}
}