#!/usr/bin/env python

import os
import io
import sqlite3
import zipfile
import csv

dbFile = "sqlite.db"

if os.path.exists(dbFile):
    os.remove(dbFile)
    os.mknod(dbFile)
else:
    os.mknod(dbFile)

with open('schema.sql', 'r') as sql_file:
    sql_script = sql_file.read()

con = sqlite3.connect(dbFile)
cur = con.cursor()
cur.executescript(sql_script)
con.commit()

with zipfile.ZipFile("archive.zip") as zf:
    for filename in zf.namelist():
        csv_file = zf.open(filename, 'r')
        dict = csv.DictReader(io.TextIOWrapper(
            csv_file, "utf-8"))

        if "country_data" in filename:
            db_data = [(row["country"], row["region"]) for row in dict]
            cur.executemany(
                "INSERT INTO countries (country, region) VALUES (?, ?);", db_data)
            con.commit()
        elif "song_data" in filename:
            db_data = []
            for row in dict:
                year = row["year"]
                semi_draw_position = row["semi_draw_position"]
                final_draw_position = None
                country = row["country"]
                artist_name = row["artist_name"]
                song_name = row["song_name"]
                language = row["language"]
                style = row["style"]
                direct_qualifier_10 = None
                gender = row["gender"]
                main_singers = row["main_singers"]
                age = row["age"]
                selection = row["selection"]
                key = row["key"]
                BPM = row["BPM"]
                energy = row["energy"]
                danceability = row["danceability"]
                happiness = row["happiness"]
                loudness = row["loudness"]
                acousticness = row["acousticness"]
                instrumentalness = row["instrumentalness"]
                liveness = row["liveness"]
                speechiness = row["speechiness"]
                release_date = row["release_date"]
                key_change_10 = row["key_change_10"]
                backing_dancers = row["backing_dancers"]
                backing_singers = row["backing_singers"]
                backing_instruments = row["backing_instruments"]
                instrument_10 = row["instrument_10"]
                qualified = None
                final_televote_points = None
                final_jury_points = None
                final_televote_votes = None
                final_jury_votes = None
                final_place = None
                final_total_points = None
                semi_place = None
                semi_total_points = None
                favourite_10 = row["favourite_10"]
                race = row["race"]
                host_10 = row["host_10"]
                if row["final_draw_position"] != "-":
                    final_draw_position = row["final_draw_position"]
                if row["direct_qualifier_10"] != "-":
                    direct_qualifier_10 = row["direct_qualifier_10"]
                if row["qualified"] != "-":
                    qualified = row["qualified"]
                if row["final_televote_points"] != "-":
                    final_televote_points = row["final_televote_points"]
                if row["final_jury_points"] != "-":
                    final_jury_points = row["final_jury_points"]
                if row["final_televote_votes"] != "-":
                    final_televote_votes = row["final_televote_votes"]
                if row["final_jury_votes"] != "-":
                    final_jury_votes = row["final_jury_votes"]
                if row["final_place"] != "-":
                    final_place = row["final_place"]
                if row["final_total_points"] != "-":
                    final_total_points = row["final_total_points"]
                if row["semi_place"] != "-":
                    semi_place = row["semi_place"]
                if row["semi_total_points"] != "-":
                    semi_total_points = row["semi_total_points"]
                db_data.append((year, semi_draw_position, final_draw_position, country, artist_name, song_name, language, style, direct_qualifier_10, gender, main_singers, age, selection, key, BPM, energy, danceability, happiness, loudness, acousticness, instrumentalness, liveness, speechiness,
                               release_date, key_change_10, backing_dancers, backing_singers, backing_instruments, instrument_10, qualified, final_televote_points, final_jury_points, final_televote_votes, final_jury_votes, final_place, final_total_points, semi_place, semi_total_points, favourite_10, race, host_10))
            cur.executemany(
                """INSERT INTO songs 
                (year, semi_draw_position, final_draw_position, country, artist_name, song_name, language, style, direct_qualifier_10, gender, main_singers, age, selection, key, BPM, energy, danceability, happiness, loudness, acousticness, instrumentalness, liveness, speechiness,
                               release_date, key_change_10, backing_dancers, backing_singers, backing_instruments, instrument_10, qualified, final_televote_points, final_jury_points, final_televote_votes, final_jury_votes, final_place, final_total_points, semi_place, semi_total_points, favourite_10, race, host_10) 
                VALUES 
                (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
                ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);""", db_data)
            con.commit()
        elif "contest_data" in filename:
            db_data = [(row["year"], row["host"], row["date"], row["semi_countries"], row["final_countries"],
                        row["jury_countries_voting_final"], row["televote_countries_voting_final"] or None) for row in dict]
            cur.executemany(
                """INSERT INTO contests 
                (year, host, date, semi_countries, final_countries, 
                jury_countries_voting_final, televote_countries_voting_final) 
                VALUES 
                (?, ?, ?, ?, ?, ?, ?);""", db_data)
            con.commit()
        elif "jury_results" in filename or "televote_results" in filename:
            db_data = []
            year = filename.split("_")[0].split("/")[-1]
            for row in dict:
                contestant = ""
                keys = row.keys()
                for key in keys:
                    country = ""
                    score = 0
                    if key == "Contestant":
                        contestant = row[key]
                    elif key == "Total score" or key == "Jury score" or key == "Televoting score":
                        continue
                    else:
                        country = key
                        score_str = row[key]
                        if score_str != "":
                            score = int(score_str)
                    db_data.append((year, contestant, country, score))

            table = "jury_results"
            if "televote_results" in filename:
                table = "televote_results"
            cur.executemany(
                "INSERT INTO {table} (year, contestant, country, score) VALUES (?, ?, ?, ?);".format(table=table), db_data)
            con.commit()
        elif "eurovisionworld_results" in filename:
            year = filename.split("_")[0].split("/")[-1]
            db_data = [(year, row["Contestant"], row["Votes"]) for row in dict]
            cur.executemany(
                "INSERT INTO eurovision_world_poll (year, contestant, votes) VALUES (?, ?, ?);", db_data)
            con.commit()
        elif "ogae_results" in filename:
            year = filename.split("_")[0].split("/")[-1]
            db_data = [(year, row["Contestant"], row["Average Points"])
                       for row in dict]
            cur.executemany(
                "INSERT INTO ogae_poll (year, contestant, average_points) VALUES (?, ?, ?);", db_data)
            con.commit()

con.close()
