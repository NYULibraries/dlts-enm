SELECT t.tct_id, t.display_name_do_not_use, n.name

FROM topics t INNER JOIN names n ON t.tct_id = n.topic_id

ORDER BY t.display_name_do_not_use, n.name
