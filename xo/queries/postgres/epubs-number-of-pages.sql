SELECT LEFT(RIGHT(oee.source, 18), 13) AS isbn, COUNT(ol.*) as number_of_pages
FROM otx_epub_epub oee INNER JOIN occurrence_location ol on oee.document_ptr_id = ol.document_id
GROUP BY isbn
ORDER BY isbn
