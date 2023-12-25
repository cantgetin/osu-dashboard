-- +migrate Up
CREATE TABLE users
(
    id                         integer not null primary key,
    username                   text    not null,
    avatar_url                 text    not null,
    graveyard_beatmapset_count integer not null,
    unranked_beatmapset_count  integer not null,
    created_at                 timestamp default NOW(),
    updated_at                 timestamp default NOW()
);

CREATE TABLE mapsets
(
    id           integer not null primary key,
    artist       text    not null,
    title        text    not null,
    covers       jsonb    not null,
    status       text    not null,
    last_updated timestamp,
    user_id      integer not null,
    constraint user_id_fk foreign key (user_id) references users (id),
    creator      text    not null,
    preview_url  text    not null,
    tags         text    not null,
    bpm          real    not null,
    mapset_stats jsonb,
    created_at   timestamp default NOW(),
    updated_at   timestamp default NOW()
);

CREATE TABLE beatmaps
(
    id                integer not null primary key,
    mapset_id         integer not null,
    constraint mapset_id_fk foreign key (mapset_id) references mapsets (id),
    difficulty_rating float   not null,
    version           text    not null,
    accuracy          float   not null,
    ar                float   not null,
    bpm               float   not null,
    cs                float   not null,
    status            text    not null,
    url               text    not null,
    total_length      integer not null,
    user_id           integer not null,
    constraint user_id_fk foreign key (user_id) references users (id),
    last_updated      timestamp,
    beatmap_stats     jsonb,
    created_at        timestamp default NOW(),
    updated_at        timestamp default NOW()
);

-- +migrate Down
DROP TABLE beatmaps;
DROP TABLE mapsets;
DROP TABLE users;