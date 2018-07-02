SELECT trs.topic2_id, t.display_name_do_not_use, COUNT( * ) AS number_of_occurrences
FROM topic_relations_simple trs INNER JOIN topics t ON trs.topic2_id = t.tct_id
INNER JOIN occurrences o ON o.topic_id = trs.topic2_id
WHERE trs.topic1_id = %%topic_id int%%
GROUP BY trs.topic2_id, t.display_name_do_not_use

UNION

SELECT trs.topic2_id, t.display_name_do_not_use, 0 AS number_of_occurrences
FROM topic_relations_simple trs INNER JOIN topics t ON trs.topic2_id = t.tct_id
LEFT OUTER JOIN occurrences o ON o.topic_id = trs.topic2_id
WHERE trs.topic1_id = %%topic_id int%%
AND o.topic_id IS NULL
GROUP BY trs.topic2_id, t.display_name_do_not_use

ORDER BY display_name_do_not_use