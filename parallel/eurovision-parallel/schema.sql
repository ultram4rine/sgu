CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS countries (
    country TEXT NOT NULL,
    region TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS songs (
    year INTEGER NOT NULL,
    country TEXT NOT NULL,
    artist_name TEXT NOT NULL,
    song_name TEXT NOT NULL,
    semi_draw_position INTEGER NOT NULL,
    final_draw_position INTEGER NULL,
    language TEXT NOT NULL,
    style TEXT NOT NULL,
    direct_qualifier_10 INTEGER NULL,
    gender TEXT NOT NULL,
    main_singers INTEGER NOT NULL,
    age TEXT NOT NULL,
    selection TEXT NOT NULL,
    key TEXT NOT NULL,
    bpm INTEGER NOT NULL,
    energy INTEGER NOT NULL,
    danceability INTEGER NOT NULL,
    happiness INTEGER NOT NULL,
    loudness TEXT NOT NULL,
    acousticness INTEGER NOT NULL,
    instrumentalness INTEGER NOT NULL,
    liveness INTEGER NOT NULL,
    speechiness INTEGER NOT NULL,
    release_date TEXT NOT NULL,
    key_change_10 TEXT NOT NULL,
    backing_dancers INTEGER NOT NULL,
    backing_singers INTEGER NOT NULL,
    backing_instruments INTEGER NOT NULL,
    instrument_10 INTEGER NOT NULL,
    qualified INTEGER NULL,
    final_televote_points INTEGER NULL,
    final_jury_points INTEGER NULL,
    final_televote_votes INTEGER NULL,
    final_jury_votes INTEGER NULL,
    final_place INTEGER NULL,
    final_total_points INTEGER NULL,
    semi_place INTEGER NULL,
    semi_total_points INTEGER NULL,
    favourite_10 INTEGER NOT NULL,
    race TEXT NOT NULL,
    host_10 INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS contests (
    year INTEGER NOT NULL,
    host TEXT NOT NULL,
    date TEXT NOT NULL,
    semi_countries INTEGER NOT NULL,
    final_countries INTEGER NOT NULL,
    jury_countries_voting_final INTEGER NOT NULL,
    televote_countries_voting_final INTEGER NULL
);
CREATE TABLE IF NOT EXISTS jury_results (
    year INTEGER NOT NULL,
    contestant TEXT NOT NULL,
    country TEXT NOT NULL,
    score INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS televote_results (
    year INTEGER NOT NULL,
    contestant TEXT NOT NULL,
    country TEXT NOT NULL,
    score INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS eurovision_world_poll (
    year INTEGER NOT NULL,
    contestant TEXT NOT NULL,
    votes INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS ogae_poll (
    year INTEGER NOT NULL,
    contestant TEXT NOT NULL,
    average_points REAL NOT NULL
);