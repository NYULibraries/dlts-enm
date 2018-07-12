SELECT DISTINCT
  ol.id,
  od.title,
  od.author AS authors,
  oee.publisher,
  LEFT(RIGHT(oee.source, 18), 13) AS isbn,
  ol.localid AS page_localid,
  ol.sequence_number AS page_sequence,
  oc.text AS page_text

FROM occurrence_location ol
  INNER JOIN otx_epub_epub oee ON ol.document_id = oee.document_ptr_id
  INNER JOIN occurrence_document od ON od.id = oee.document_ptr_id
  INNER JOIN occurrence_content oc ON oc.id = ol.content_id

ORDER BY ol.id
