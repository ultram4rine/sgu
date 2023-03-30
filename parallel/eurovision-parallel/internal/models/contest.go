package models

type Contest struct {
	Year                            string `db:"year"`
	Host                            string `db:"host"`
	Date                            string `db:"date"`
	Semi_countries                  int    `db:"semi_countries"`
	Final_countries                 int    `db:"final_countries"`
	Jury_countries_voting_final     int    `db:"jury_countries_voting_final"`
	Televote_countries_voting_final int    `db:"televote_countries_voting_final"`
}
