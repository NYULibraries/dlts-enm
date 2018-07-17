SELECT hb.id, hb.display_name, COUNT(o.*) AS number_of_occurrences
FROM
  (
    SELECT rrb.destination_id AS related_topic_id
    FROM relation_relatedbasket rrb
    WHERE rrb.source_id = %%topic_id_1 int%%
    AND rrb.forbidden = FALSE

    UNION

    SELECT rrb.source_id AS related_topic_id
    FROM relation_relatedbasket rrb
    WHERE rrb.destination_id = %%topic_id_2 int%%
    AND rrb.forbidden = FALSE
  ) AS r INNER JOIN hit_basket hb ON r.related_topic_id = hb.id
         INNER JOIN occurrence_occurrence o on hb.id = o.basket_id
GROUP BY hb.id, hb.display_name

UNION

SELECT hb.id, hb.display_name, 0 AS number_of_occurrences
FROM
  (
    SELECT rrb.destination_id AS related_topic_id
    FROM relation_relatedbasket rrb
    WHERE rrb.source_id = %%topic_id_3 int%%
          AND rrb.forbidden = FALSE

    UNION

    SELECT rrb.source_id AS related_topic_id
    FROM relation_relatedbasket rrb
    WHERE rrb.destination_id = %%topic_id_4 int%%
          AND rrb.forbidden = FALSE
  ) AS r INNER JOIN hit_basket hb ON r.related_topic_id = hb.id
  LEFT OUTER JOIN occurrence_occurrence o on hb.id = o.basket_id
WHERE o.basket_id IS NULL
GROUP BY hb.id, hb.display_name

ORDER BY display_name

