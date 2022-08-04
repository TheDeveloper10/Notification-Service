package repository

import (
	"log"
	"strconv"

	"notification-service.com/packages/internal/clients"
	"notification-service.com/packages/internal/dto"
)

type TemplateRepository interface {
	Insert(req *dto.CreateTemplateRequest) bool
	Get(req *dto.TemplateIdRequest) (*dto.TemplateRecord, int)
	Update(req *dto.UpdateTemplateRequest) int
	Delete(req *dto.TemplateIdRequest) bool
}

type basicTemplateRepository struct { }

func NewTemplateRepository() TemplateRepository {
	return &basicTemplateRepository{}
}

func (btr *basicTemplateRepository) Insert(req *dto.CreateTemplateRequest) bool {
	stmt, err1 := clients.SQLClient.Prepare("insert into Templates(ContactType, Template) values(?, ?)")
	if err1 != nil {
		log.Print(err1.Error())
		return false
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(*req.ContactType, *req.Template)
	if err2 != nil {
		log.Print(err2.Error())
		return false
	}
	id, err3 := res.LastInsertId()
	if err3 != nil {
		log.Print(err3.Error())
		return false
	}
	log.Print("Inserted template with id " + strconv.FormatInt(id, 10))
	return true
}

func (btr *basicTemplateRepository) Get(req *dto.TemplateIdRequest) (*dto.TemplateRecord, int) {
	stmt, err1 := clients.SQLClient.Prepare("select * from Templates where Id=?")
	if err1 != nil {
		log.Print(err1.Error())
		return nil, 1
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(*req.Id)
	if err2 != nil {
		log.Print(err2.Error())
		return nil, 1
	}
	defer rows.Close()

	if rows.Next() {
		record := dto.TemplateRecord{ }
		if err3 := rows.Scan(&record.Id, &record.ContactType, &record.Template); err3 != nil {
			log.Print(err3.Error())
			return nil, 2
		}
		log.Print("Fetched template with id " + strconv.Itoa(*req.Id))
		return &record, 0
	} else {
		log.Print("Template with id " + strconv.Itoa(*req.Id) + " was not found")
		return nil, 2
	}
}

func (btr *basicTemplateRepository) Update(req *dto.UpdateTemplateRequest) int {
	stmt, err1 := clients.SQLClient.Prepare("update Templates set Template=? where Id=?")
	if err1 != nil {
		log.Print(err1.Error())
		return 1
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(req.Template, req.Id)
	if err2 != nil {
		log.Print(err2.Error())
		return 1
	}
	affectedRows, err3 := res.RowsAffected() 
	if err3 != nil {
		log.Print(err3.Error())
		return 1
	}
	if affectedRows <= 0 {
		log.Print("No template was found!")
		return 2
	}

	log.Print("Updated template with id " + strconv.Itoa(*req.Id))
	return 0
}

func (btr *basicTemplateRepository) Delete(req *dto.TemplateIdRequest) bool {
	stmt, err1 := clients.SQLClient.Prepare("delete from Templates where Id=?")
	if err1 != nil {
		log.Print(err1.Error())
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.Id)
	if err2 != nil {
		log.Print(err2.Error())
		return false
	}
	log.Print("Deleted template with id " + strconv.Itoa(*req.Id))
	return true
}
