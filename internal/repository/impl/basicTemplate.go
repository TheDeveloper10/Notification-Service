package impl

import (
	"github.com/sirupsen/logrus"
	"notification-service/internal/client"
	"notification-service/internal/entity"
	"notification-service/internal/helper"
	"notification-service/internal/util"
	"strconv"
)

type BasicTemplateRepository struct {
	cache map[int]*entity.TemplateEntity
}

func (btr *BasicTemplateRepository) Init() {
	btr.cache = map[int]*entity.TemplateEntity{}
}

func (btr *BasicTemplateRepository) Insert(entity *entity.TemplateEntity) int {
	res := client.SQLClient.Exec(
		"insert into Templates(ContactType, Template, Language, Type) values(?, ?, ?, ?)",
		entity.ContactType, entity.Template, entity.Language, entity.Type,
	)
	if res == nil {
		return -1
	}
	id, err3 := res.LastInsertId()
	if helper.IsError(err3) {
		return -1
	}

	logrus.Info("Inserted template into the database with id " + strconv.FormatInt(id, 10))
	btr.cache[int(id)] = entity
	return int(id)
}

func (btr *BasicTemplateRepository) Get(id int) (*entity.TemplateEntity, int) {
	if result, ok := btr.cache[id]; ok {
		return result, 0
	}

	rows := client.SQLClient.Query("select * from Templates where Id=?", id)
	if rows == nil {
		return nil, 1
	}
	defer helper.HandledClose(rows)

	if rows.Next() {
		record := entity.TemplateEntity{}
		err3 := rows.Scan(&record.Id, &record.ContactType, &record.Template, &record.Language, &record.Type)
		if helper.IsError(err3) {
			return nil, 2
		}

		logrus.Info("Fetched template with id " + strconv.Itoa(id))

		go btr.clearCache()
		btr.cache[id] = &record
		return &record, 0
	} else {
		logrus.Warn("Template with id " + strconv.Itoa(id) + " was not found")
		return nil, 2
	}
}

func (btr *BasicTemplateRepository) GetBulk(filter *entity.TemplateFilter) *[]entity.TemplateEntity {
	builder := util.NewQueryBuilder("select * from Templates")

	offset := (filter.Page - 1) * filter.Size
	query := builder.End(&filter.Size, &offset)

	rows := client.SQLClient.Query(*query)
	if rows == nil {
		return nil
	}
	defer helper.HandledClose(rows)

	var templates []entity.TemplateEntity
	for rows.Next() {
		record := entity.TemplateEntity{}
		err3 := rows.Scan(&record.Id, &record.ContactType,
			&record.Template, &record.Language, &record.Type)
		if helper.IsError(err3) {
			return nil
		}

		btr.cache[record.Id] = &record
		templates = append(templates, record)
	}

	logrus.Info("Fetched " + strconv.Itoa(len(templates)) + " template(s)")
	return &templates
}

func (btr *BasicTemplateRepository) Update(entity *entity.TemplateEntity) int {
	res := client.SQLClient.Exec(
		"update Templates set Template=?, ContactType=?, Language=?, Type=? where Id=?",
		entity.Template, entity.ContactType, entity.Language, entity.Type, entity.Id,
	)
	if res == nil {
		return 1
	}

	affectedRows, err := res.RowsAffected()
	if helper.IsError(err) {
		return 1
	}

	// TODO: it's zero also when template is found but the value you set is the same** FIX IT
	if affectedRows <= 0 {
		logrus.Warn("No template was found with id " + strconv.Itoa(entity.Id))
		return 2
	}

	logrus.Info("Updated template in the database with id " + strconv.Itoa(entity.Id))
	btr.cache[entity.Id] = entity
	return 0
}

func (btr *BasicTemplateRepository) Delete(id int) bool {
	res := client.SQLClient.Exec("delete from Templates where Id=?", id)
	if res == nil {
		return false
	}

	logrus.Info("Deleted template from the database with id " + strconv.Itoa(id))
	delete(btr.cache, id)
	return true
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