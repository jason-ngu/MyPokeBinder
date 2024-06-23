-- Database Schema for MyPokeBinder

CREATE TABLE types(
    type_id         SERIAL PRIMARY KEY,
    type_name       VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE subtypes(
    subtype_id      SERIAL PRIMARY KEY,
    subtype_name    VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE supertypes(
    supertype_id        SERIAL PRIMARY KEY,
    supertype_name    VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE rarities(
    rarity_id       SERIAL PRIMARY KEY,
    rarity_name     VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE series(
    series_id       SERIAL PRIMARY KEY,
    series_name     VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE pricetypes(
    pricetype_id       SERIAL PRIMARY KEY,
    pricetype_name     VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE users(
    user_id       SERIAL PRIMARY KEY,
    user_name     text UNIQUE NOT NULL
);

CREATE TABLE collections(
    collection_id       SERIAL PRIMARY KEY,
    collection_name     VARCHAR(50) UNIQUE NOT NULL,
    user_id             INT NOT NULL REFERENCES users(user_id)
                                        ON UPDATE CASCADE
                                        ON DELETE RESTRICT
);

CREATE TABLE sets(
    set_id              SERIAL PRIMARY KEY,
    api_id              INTEGER NOT NULL,
    set_name            text NOT NULL,
    series_id           INTEGER NOT NULL REFERENCES series(series_id),
    ptcgo_code          VARCHAR(5),
    card_total          INTEGER NOT NULL,
    extended_card_total INTEGER NOT NULL,
    set_release_Date    TIMESTAMPTZ,
    symbol_image        TEXT,
    logo_image          TEXT,
    sync_date_created   TIMESTAMPTZ,
    sync_date_updated   TIMESTAMPTZ
);

CREATE TABLE cards(
    card_id             SERIAL PRIMARY KEY,
    api_id              INTEGER NOT NULL,
    card_name           text NOT NULL,
    set_id              INTEGER NOT NULL REFERENCES sets(set_id),
    supertype_id        INTEGER NOT NULL REFERENCES supertypes(supertype_id),
    rarity_id           INTEGER NOT NULL REFERENCES rarities(rarity_id),
    market_price        DECIMAL(12,2),
    pricetype_id        INTEGER NOT NULL REFERENCES pricetypes(pricetype_id),
    image               TEXT,
    sync_date_created   TIMESTAMPTZ,
    sync_date_updated   TIMESTAMPTZ 
);

-- Bridge tables

CREATE TABLE cardtypes(
    card_id    INT NOT NULL REFERENCES cards(card_id),
    type_id    INT NOT NULL REFERENCES types(type_id)
);

CREATE TABLE cardsubtypes(
    card_id     INT NOT NULL REFERENCES cards(card_id),
    subtype_id  INT NOT NULL REFERENCES subtypes(subtype_id)
);

CREATE TABLE collectioncards(
    collection_id     INT NOT NULL REFERENCES collections(collection_id),
    card_id           INT NOT NULL REFERENCES cards(card_id)
);