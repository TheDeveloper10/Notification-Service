package repository

import (
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
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(*req.ContactType, *req.Template)
	return err2 == nil
}

func (btr *basicTemplateRepository) Get(req *dto.TemplateIdRequest) (*dto.TemplateRecord, int) {
	stmt, err1 := clients.SQLClient.Prepare("select * from Templates where Id=?")
	if err1 != nil {
		return nil, 1
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(*req.Id)
	if err2 != nil {
		return nil, 1
	}
	defer rows.Close()

	if rows.Next() {
		record := dto.TemplateRecord{ }
		if err3 := rows.Scan(&record.Id, &record.ContactType, &record.Template); err3 != nil {
			return nil, 2
		}
		return &record, 0
	} else {
		return nil, 2
	}
}

func (btr *basicTemplateRepository) Update(req *dto.UpdateTemplateRequest) int {

	stmt, err1 := clients.SQLClient.Prepare("update Templates set Template=? where Id=?")
	if err1 != nil {
		return 1
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(req.Template, req.Id)
	if err2 != nil {
		return 1
	}
	affectedRows, err3 := res.RowsAffected() 
	if err3 != nil {
		return 1
	}
	if affectedRows <= 0 {
		return 2
	}

	return 0
}

func (btr *basicTemplateRepository) Delete(req *dto.TemplateIdRequest) bool {
	stmt, err1 := clients.SQLClient.Prepare("delete from Templates where Id=?")
	if err1 != nil {
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.Id)
	return err2 == nil
}
