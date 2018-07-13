SELECT
  oo.location_id AS page_id,
  hb.id AS topic_id,
  hb.display_name AS topic_display_name,
  hh.name AS topic_name,

  CASE WHEN LEFT( hb.display_name, 1 ) = '"'
    THEN LOWER( SUBSTR( hb.display_name, 2 ) )
  ELSE LOWER( hb.display_name )
  END AS topic_display_name_sort_key,

  CASE WHEN LEFT( hh.name, 1 ) = '"'
    THEN LOWER( SUBSTR( hh.name, 2 ) )
  ELSE LOWER( hh.name )
  END AS topic_name_sort_key

FROM occurrence_occurrence oo INNER JOIN hit_basket hb ON oo.basket_id = hb.id
  INNER JOIN hit_hit hh ON hh.basket_id = hb.id

WHERE oo.location_id = %%page_id int%%

ORDER BY page_id, topic_display_name_sort_key, topic_name_sort_key
