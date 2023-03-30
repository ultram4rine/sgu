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
                semi_draw_position = None
                final_draw_position = None
                country = row["country"]
                artist_name = row["artist_name"]
                song_name = row["song_name"]
                language = row["language"]
                style = row["style"]
                direct_qualifier_10 = None
                gender = row["gender"]
                main_singers = None
                age = row["age"]
                selection = row["selection"]
                key = row["key"]
                BPM = None
                energy = None
                danceability = None
                happiness = None
                loudness = row["loudness"]
                acousticness = None
                instrumentalness = None
                liveness = None
                speechiness = row["speechiness"]
                release_date = row["release_date"]
                key_change_10 = row["key_change_10"]
                backing_dancers = None
                backing_singers = None
                backing_instruments = None
                instrument_10 = None
                qualified = None
                final_televote_points = None
                final_jury_points = None
                final_televote_votes = None
                final_jury_votes = None
                final_place = None
                final_total_points = None
                semi_place = None
                semi_total_points = None
                favourite_10 = None
                race = row["race"]
                host_10 = None

                if row["semi_draw_position"] != "-" and row["semi_draw_position"] != "":
                    semi_draw_position = row["semi_draw_position"]
                if row["final_draw_position"] != "-" and row["final_draw_position"] != "":
                    final_draw_position = row["final_draw_position"]
                if row["direct_qualifier_10"] != "-" and row["direct_qualifier_10"] != "":
                    direct_qualifier_10 = row["direct_qualifier_10"]
                if row["main_singers"] != "-" and row["main_singers"] != "":
                    main_singers = row["main_singers"]
                if row["qualified"] != "-" and row["qualified"] != "":
                    qualified = row["qualified"]
                if row["BPM"] != "-" and row["BPM"] != "":
                    BPM = row["BPM"]
                if row["energy"] != "-" and row["energy"] != "":
                    energy = row["energy"]
                if row["danceability"] != "-" and row["danceability"] != "":
                    danceability = row["danceability"]
                if row["happiness"] != "-" and row["happiness"] != "":
                    happiness = row["happiness"]
                if row["acousticness"] != "-" and row["acousticness"] != "":
                    acousticness = row["acousticness"]
                if row["instrumentalness"] != "-" and row["instrumentalness"] != "":
                    instrumentalness = row["instrumentalness"]
                if row["liveness"] != "-" and row["liveness"] != "":
                    liveness = row["liveness"]
                if row["backing_dancers"] != "-" and row["backing_dancers"] != "":
                    backing_dancers = row["backing_dancers"]
                if row["backing_singers"] != "-" and row["backing_singers"] != "":
                    backing_singers = row["backing_singers"]
                if row["backing_instruments"] != "-" and row["backing_instruments"] != "":
                    backing_instruments = row["backing_instruments"]
                if row["instrument_10"] != "-" and row["instrument_10"] != "":
                    instrument_10 = row["instrument_10"]
                if row["final_televote_points"] != "-" and row["final_televote_points"] != "":
                    final_televote_points = row["final_televote_points"]
                if row["final_jury_points"] != "-" and row["final_jury_points"] != "":
                    final_jury_points = row["final_jury_points"]
                if row["final_televote_votes"] != "-" and row["final_televote_votes"] != "":
                    final_televote_votes = row["final_televote_votes"]
                if row["final_jury_votes"] != "-" and row["final_jury_votes"] != "":
                    final_jury_votes = row["final_jury_votes"]
                if row["final_place"] != "-" and row["final_place"] != "":
                    final_place = row["final_place"]
                if row["final_total_points"] != "-" and row["final_total_points"] != "":
                    final_total_points = row["final_total_points"]
                if row["semi_place"] != "-" and row["semi_place"] != "":
                    semi_place = row["semi_place"]
                if row["semi_total_points"] != "-" and row["semi_total_points"] != "":
                    semi_total_points = row["semi_total_points"]
                if row["favourite_10"] != "-" and row["favourite_10"] != "":
                    favourite_10 = row["favourite_10"]
                if row["host_10"] != "-" and row["host_10"] != "":
                    host_10 = row["host_10"]

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
                        row["jury_countries_voting_final"] or 0, row["televote_countries_voting_final"] or 0) for row in dict]
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
