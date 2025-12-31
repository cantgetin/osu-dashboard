-- Insert 150 random users
INSERT INTO users (id, username, avatar_url, graveyard_beatmapset_count, unranked_beatmapset_count, user_stats, created_at, updated_at, map_counts)
WITH random_users AS (
    SELECT
        -- Generate sequential IDs starting from 7192130
        7192129 + row_number() OVER () as id,
        -- Generate random usernames
        'User_' || (7192129 + row_number() OVER ()) as username,
        -- Random avatar URL with different timestamps
        'https://a.ppy.sh/' || (7192129 + row_number() OVER ()) || '?' || (1700000000 + floor(random() * 100000000))::text || '.jpeg' as avatar_url,
        -- Random graveyard count between 0 and 100
        floor(random() * 101)::integer as graveyard_beatmapset_count,
        -- Random unranked count between 0 and 50
        floor(random() * 51)::integer as unranked_beatmapset_count,
        -- Current timestamp for created_at and updated_at
        NOW() as created_at,
        NOW() as updated_at
    FROM generate_series(1, 150)
)
SELECT
    id,
    username,
    avatar_url,
    graveyard_beatmapset_count,
    unranked_beatmapset_count,
    -- Generate random user_stats JSON similar to your example
    json_build_object(
            to_char(NOW(), 'YYYY-MM-DD"T"HH24:MI:SS.MS"Z"'),
            json_build_object(
                    'map_count', graveyard_beatmapset_count + unranked_beatmapset_count,
                    'play_count', floor(random() * 200000)::integer,
                    'comments_count', floor(random() * 200)::integer,
                    'favorite_count', floor(random() * 200)::integer
            )
    ) as user_stats,
    created_at,
    updated_at,
    -- Generate random map_counts JSON
    json_build_object(
            'wip', floor(random() * 20)::integer,
            'loved', floor(random() * 10)::integer,
            'ranked', floor(random() * 5)::integer,
            'pending', floor(random() * 15)::integer,
            'approved', floor(random() * 8)::integer,
            'graveyard', graveyard_beatmapset_count,
            'qualified', floor(random() * 3)::integer
    ) as map_counts
FROM random_users;