SELECT od.title, od.author, oee.publisher, LEFT(RIGHT(oee.source, 18), 13) AS isbn,
       COUNT(oo.*) number_of_occurrences

FROM occurrence_document od
  INNER JOIN otx_epub_epub oee ON oee.document_ptr_id = od.id
  INNER JOIN occurrence_location ol on od.id = ol.document_id
  INNER JOIN occurrence_occurrence oo ON oo.location_id = ol.id
  INNER JOIN hit_basket hb ON hb.id = oo.basket_id

WHERE hb.id = %%topic_id int%%

GROUP BY od.title, od.author, oee.publisher, isbn

ORDER BY number_of_occurrences, od.title DESC

