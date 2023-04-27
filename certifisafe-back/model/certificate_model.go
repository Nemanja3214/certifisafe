package model

import (
	"gorm.io/gorm"
	"time"
)

type Certificate struct {
	gorm.Model
	Id        *int64 `gorm:"autoIncrement;PRIMARY_KEY"`
	Name      string
	Issuer    User `gorm:"foreignKey:IssuerID;references:ID"`
	Subject   User `gorm:"foreignKey:SubjectID;references:ID"`
	ValidFrom time.Time
	ValidTo   time.Time
	Status    CertificateStatus
	Type      CertificateType
	IssuerID  *int64
	SubjectID *int64
}

type CertificateType int64
type CertificateStatus int64

const (
	ROOT CertificateType = iota
	INTERMEDIATE
	END
)
const (
	ACTIVE CertificateStatus = iota
	EXPIRED
	WITHDRAWN
	NOT_ACTIVE
)
