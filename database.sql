SELECT id, name,
CASE
    WHEN parent_id = 2 THEN 'Ilham'
    WHEN parent_id = 3 THEN 'Irwan'
    ELSE NULL
END AS parent_name
FROM table;