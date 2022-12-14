package models

type GroupServices struct {
	Id     int
	Code   string
	Name   string
	Hidden bool
}

type MeasureUnit struct {
	Id        int
	Name      string
	ShortName string
}

type Product struct {
	Id          int
	Code        string
	Name        string
	Basecost    float64
	Hidden      bool
	MeasureUnit int
}

type Service struct {
	Id       int
	Code     string
	Name     string
	Basecost float64
	Hidden   bool
	GroupId  int
}

type Attacher struct {
	Product int
	Service int
	Count   int
}

type Price struct {
	ServiceName                        string
	Price, ServicePrice, MaterialPrice float64
}

type MaterialToService struct {
	ServiceName   string
	ProductName   string
	CountProduct  int
	MeasureUnit   string
	MaterialPrice float64
}
