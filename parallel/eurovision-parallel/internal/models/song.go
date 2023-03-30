package models

import "database/sql"

type Song struct {
	Year                  string        `db:"year"`
	Country               string        `db:"country"`
	Artist_name           string        `db:"artist_name"`
	Song_name             string        `db:"song_name"`
	Semi_draw_position    sql.NullInt64 `db:"semi_draw_position"`
	Final_draw_position   sql.NullInt64 `db:"final_draw_position"`
	Language              string        `db:"language"`
	Style                 string        `db:"style"`
	Direct_qualifier_10   sql.NullInt64 `db:"direct_qualifier_10"`
	Gender                string        `db:"gender"`
	Main_singers          sql.NullInt64 `db:"main_singers"`
	Age                   string        `db:"age"`
	Selection             string        `db:"selection"`
	Key                   string        `db:"key"`
	Bpm                   sql.NullInt64 `db:"bpm"`
	Energy                sql.NullInt64 `db:"energy"`
	Danceability          sql.NullInt64 `db:"danceability"`
	Happiness             sql.NullInt64 `db:"happiness"`
	Loudness              string        `db:"loudness"`
	Acousticness          sql.NullInt64 `db:"acousticness"`
	Instrumentalness      sql.NullInt64 `db:"instrumentalness"`
	Liveness              sql.NullInt64 `db:"liveness"`
	Speechiness           string        `db:"speechiness"`
	Release_date          string        `db:"release_date"`
	Key_change_10         string        `db:"key_change_10"`
	Backing_dancers       sql.NullInt64 `db:"backing_dancers"`
	Backing_singers       sql.NullInt64 `db:"backing_singers"`
	Backing_instruments   sql.NullInt64 `db:"backing_instruments"`
	Instrument_10         sql.NullInt64 `db:"instrument_10"`
	Qualified             sql.NullInt64 `db:"qualified"`
	Final_televote_points sql.NullInt64 `db:"final_televote_points"`
	Final_jury_points     sql.NullInt64 `db:"final_jury_points"`
	Final_televote_votes  sql.NullInt64 `db:"final_televote_votes"`
	Final_jury_votes      sql.NullInt64 `db:"final_jury_votes"`
	Final_place           sql.NullInt64 `db:"final_place"`
	Final_total_points    sql.NullInt64 `db:"final_total_points"`
	Semi_place            sql.NullInt64 `db:"semi_place"`
	Semi_total_points     sql.NullInt64 `db:"semi_total_points"`
	Favourite_10          sql.NullInt64 `db:"favourite_10"`
	Race                  string        `db:"race"`
	Host_10               sql.NullInt64 `db:"host_10"`
}
