SELECT trs.topic2_id, t.display_name_do_not_use, COUNT( o.tct_id )

FROM topic_relations_simple trs INNER JOIN topics t ON trs.topic2_id = t.tct_id
  INNER JOIN occurrences o ON o.topic_id = trs.topic2_id

WHERE topic1_id = %%topic_id int%%

GROUP BY trs.topic2_id, t.display_name_do_not_use

ORDER BY t.display_name_do_not_use