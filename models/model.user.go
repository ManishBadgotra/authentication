package models

import "time"

type TimeStamp struct {
	LoggedInBefore bool        `json:"firstlogin,omitempty"`
	Date_time      string      `json:"Date_time,omitempty" bson:"-"`
	LastLogin      []time.Time `json:"lastLogin,omitempty"`
	Created        time.Time   `json:"createdAt,omitempty"`
	Edited         time.Time   `json:"editedAt,omitempty"`
	Deleted        time.Time   `json:"deletedAt,omitempty"`
	IsDeleted      bool        `json:"isDeleted,omitempty"`
}

type EmailUser struct {
	Id                   string               `json:"id,omitempty" bson:"_id,omitempty"`
	Title                string               `json:"title,omitempty"`
	Firstname            string               `json:"firstname,omitempty"`
	Lastname             string               `json:"lastname,omitempty"`
	Birthdate            time.Time            `json:"birthdate,omitempty"`
	Gender               string               `json:"gender,omitempty"`
	ProfilePicture       string               `json:"profilepicture,omitempty"`
	MaritalStatus        string               `json:"maritalstatus,omitempty"`
	Address              Address              `json:"address,omitempty"`
	Bio                  string               `json:"bio,omitempty"`
	Preferences          Preferences          `json:"preferences,omitempty"`
	SecurityQuestions    SecurityQuestions    `json:"security,omitempty"`
	SocialMediaProfiles  SocialProfiles       `json:"socialmediaprofiles,omitempty"`
	PaymentInfo          PaymentInfo          `json:"paymentinfo,omitempty"`
	LegalConsent         LegalConsent         `json:"legalconsent,omitempty"`
	IdentityVerification IdentityVerification `json:"identityverification,omitempty"`
	Email                string               `json:"email" binding:"required"`
	Password             string               `json:"password" binding:"required"`
	PhoneCode            string               `json:"phonecode,omitempty"`
	PhoneNumber          string               `json:"phonenumber,omitempty"`
	Time                 TimeStamp            `json:"time"`
}

type PhoneUser struct {
	Id                   string               `json:"id,omitempty" bson:"_id,omitempty"`
	Title                string               `json:"title,omitempty"`
	Firstname            string               `json:"firstname,omitempty"`
	Lastname             string               `json:"lastname,omitempty"`
	Birthdate            time.Time            `json:"birthdate,omitempty"`
	Gender               string               `json:"gender,omitempty"`
	ProfilePicture       string               `json:"profilepicture,omitempty"`
	MaritalStatus        string               `json:"maritalstatus,omitempty"`
	Address              Address              `json:"address,omitempty"`
	Bio                  string               `json:"bio,omitempty"`
	Preferences          Preferences          `json:"preferences,omitempty"`
	SecurityQuestions    SecurityQuestions    `json:"security,omitempty"`
	SocialMediaProfiles  SocialProfiles       `json:"socialmediaprofiles,omitempty"`
	PaymentInfo          PaymentInfo          `json:"paymentinfo,omitempty"`
	LegalConsent         LegalConsent         `json:"legalconsent,omitempty"`
	IdentityVerification IdentityVerification `json:"identityverification,omitempty"`
	Email                string               `json:"email,omitempty"`
	Password             string               `json:"password,omitempty"`
	PhoneCode            string               `json:"phonecode" binding:"required"`
	PhoneNumber          string               `json:"phonenumber" binding:"required"`
	Time                 TimeStamp            `json:"time"`
}

type Otp struct {
	Otp string `json:"otp" binding:"required"`
}

type Address struct {
	CurrentAddress   string
	PermanentAddress string
	TemporaryStay    string
}

type Preferences struct {
	Preference []map[string]string
}

type SecurityQuestions struct {
	Question []map[string]string
}

type SocialProfiles struct {
	SocialProfiles []string
}

type PaymentInfo struct {
	PaymentInfo map[string]string
}

type LegalConsent struct {
	Consent map[string][]string
}

type IdentityVerification struct {
	IdType map[string]string
}
