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
    year TEXT NOT NULL,
    country TEXT NOT NULL,
    artist_name TEXT NOT NULL,
    song_name TEXT NOT NULL,
    semi_draw_position INTEGER NULL,
    final_draw_position INTEGER NULL,
    language TEXT NOT NULL,
    style TEXT NOT NULL,
    direct_qualifier_10 INTEGER NULL,
    gender TEXT NOT NULL,
    main_singers INTEGER NULL,
    age TEXT NOT NULL,
    selection TEXT NOT NULL,
    key TEXT NOT NULL,
    bpm INTEGER NULL,
    energy INTEGER NULL,
    danceability INTEGER NULL,
    happiness INTEGER NULL,
    loudness TEXT NOT NULL,
    acousticness INTEGER NULL,
    instrumentalness INTEGER NULL,
    liveness INTEGER NULL,
    speechiness TEXT NOT NULL,
    release_date TEXT NOT NULL,
    key_change_10 TEXT NOT NULL,
    backing_dancers INTEGER NULL,
    backing_singers INTEGER NULL,
    backing_instruments INTEGER NULL,
    instrument_10 INTEGER NULL,
    qualified INTEGER NULL,
    final_televote_points INTEGER NULL,
    final_jury_points INTEGER NULL,
    final_televote_votes INTEGER NULL,
    final_jury_votes INTEGER NULL,
    final_place INTEGER NULL,
    final_total_points INTEGER NULL,
    semi_place INTEGER NULL,
    semi_total_points INTEGER NULL,
    favourite_10 INTEGER NULL,
    race TEXT NOT NULL,
    host_10 INTEGER NULL
);
CREATE TABLE IF NOT EXISTS contests (
    year TEXT NOT NULL,
    host TEXT NOT NULL,
    date TEXT NOT NULL,
    semi_countries INTEGER NOT NULL,
    final_countries INTEGER NOT NULL,
    jury_countries_voting_final INTEGER NULL,
    televote_countries_voting_final INTEGER NULL
);
CREATE TABLE IF NOT EXISTS jury_results (
    year TEXT NOT NULL,
    contestant TEXT NOT NULL,
    country TEXT NOT NULL,
    score INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS televote_results (
    year TEXT NOT NULL,
    contestant TEXT NOT NULL,
    country TEXT NOT NULL,
    score INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS eurovision_world_poll (
    year TEXT NOT NULL,
    contestant TEXT NOT NULL,
    votes INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS ogae_poll (
    year TEXT NOT NULL,
    contestant TEXT NOT NULL,
    average_points REAL NOT NULL
);