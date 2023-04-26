package repository

import (
	"certifisafe-back/model"
	"certifisafe-back/utils"
	"database/sql"
	"errors"
	"math/big"
)

const store = "keystore.jsk"

var (
	ErrCertificateNotFound = errors.New("FromRepository - certificate not found")
)

type ICertificateRepository interface {
	GetCertificate(id big.Int) (model.Certificate, error)
	DeleteCertificate(id big.Int) error
	CreateCertificate(certificate model.Certificate) (model.Certificate, error)
	GetCertificates() ([]model.Certificate, error)
}

type InmemoryCertificateRepository struct {
	DB *sql.DB
}

func NewInMemoryCertificateRepository(db *sql.DB) *InmemoryCertificateRepository {
	return &InmemoryCertificateRepository{
		DB: db,
	}
}

func (i *InmemoryCertificateRepository) GetCertificate(id big.Int) (model.Certificate, error) {
	//TODO add subject and issuer
	stmt, err := i.DB.Prepare("SELECT id, name, valid_from, valid_to, type, status  FROM certificates WHERE id=$1")
	utils.CheckError(err)

	var certificate model.Certificate
	err = stmt.QueryRow(id.String()).Scan(&certificate.Id, &certificate.Name, &certificate.ValidFrom, &certificate.ValidTo,
		&certificate.Type, &certificate.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case of no rows returned.
		}
		return model.Certificate{}, err

	}
	return certificate, nil
}

func (i *InmemoryCertificateRepository) GetCertificates() ([]model.Certificate, error) {
	var result []model.Certificate
	//TODO add subject and issuer
	rows, err := i.DB.Query("SELECT id, name, valid_from, valid_to, type, status  FROM certificates")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var certificate model.Certificate
		rows.Scan(&certificate.Id, &certificate.Name, &certificate.ValidFrom, &certificate.ValidTo,
			&certificate.Type, &certificate.Status)
		result = append(result, certificate)
	}
	utils.CheckError(err)

	return result, nil
}

func (i *InmemoryCertificateRepository) DeleteCertificate(id big.Int) error {
	return nil
	//return ErrMovieNotFound
}

func (i *InmemoryCertificateRepository) CreateCertificate(certificate model.Certificate) (model.Certificate, error) {
	subject := 1
	if certificate.Subject != nil {
		subject = certificate.Subject.Id
	}
	issuer := 1

	if certificate.Issuer != nil {
		issuer = certificate.Issuer.Id
	}
	t := model.INTERMEDIATE
	if certificate.Issuer == nil {
		t = model.ROOT
	}

	err := i.DB.QueryRow(
		"INSERT INTO certificates(name, valid_from, valid_to, subject_id, issuer_id, type, status) VALUES($1, $2, $3, $4, $5, $6, $7)", certificate.Name, certificate.ValidFrom, certificate.ValidTo, subject, issuer, t, model.NOT_ACTIVE)
	if err.Err() != nil {
		return model.Certificate{}, err.Err()
	}

	return certificate, nil

}
