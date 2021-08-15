CREATE VIEW WordListsSummaries AS (
    SELECT
        wl.id,
        wl.user_id,
        wl.word_list_title,
        wl.explanation,
        count(DISTINCT w.id) AS word_count,
        count(DISTINCT s.id) AS play_count,
        max(s.played_at) AS played_at,
        wl.created_at,
        wl.updated_at
    FROM wordlists wl
    LEFT JOIN words w 
        ON w.word_list_id = wl.id 
    LEFT JOIN scores s 
        ON s.word_list_id = wl.id
    GROUP BY wl.id
);