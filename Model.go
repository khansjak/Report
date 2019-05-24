package pck

import "time"

type AppointmentOutput struct {
	ID                  string    `bson:"dealerID" json:"dealerID"`
	Dealer              string    `bson:"Dealer" json:"Dealer,omitempty"`
	ServiceAdvisor      string    `bson:"serviceAdvisorName" json:"serviceAdvisorName"`
	AppointmentDateTime time.Time `bson:"appointmentDateTime" json:"appointmentDateTime"`
	CustomerName        string    `bson:"CustomerName" json:"CustomerName,omitempty"`
	LastUpdatedBy       string    `bson:"lastUpdatedByDisplay" json:"lastUpdatedByDisplay"`
	AppointmentStatus   int    `bson:"status" json:"status"`
	LastName			string    `bson:"lastName" json:"lastName"`
	FirstName			string    `bson:"firstName" json:"firstName"`
}

type Dealer struct {
	Id string `bson:"_id" json:"_id"`
	Name string `bson:"dealerName" json:"dealerName"`
}

type Report struct {
	ID                  string    `bson:"dealerID" json:"dealerID" `
	Dealer              string    `bson:"Dealer" json:"Dealer,omitempty"`
	ServiceAdvisor      string    `bson:"serviceAdvisorName" json:"serviceAdvisorName"`
	AppointmentDateTime time.Time `bson:"appointmentDateTime" json:"appointmentDateTime"`
	CustomerName        string    `bson:"CustomerName" json:"CustomerName,omitempty"`
	LastUpdatedBy       string    `bson:"lastUpdatedByDisplay" json:"lastUpdatedByDisplay"`
	AppointmentStatus   int    `bson:"status" json:"status"`
}
