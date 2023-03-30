package models

type Country struct {
	Country string `db:"country"`
	Region  string `db:"region"`
}
