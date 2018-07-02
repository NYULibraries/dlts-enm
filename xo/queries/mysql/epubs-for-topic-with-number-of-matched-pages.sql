SELECT e.title, e.author, e.publisher, e.isbn, COUNT( o.tct_id ) AS number_of_occurrences

FROM epubs e INNER JOIN locations l ON e.tct_id = l.epub_id
  INNER JOIN occurrences o ON o.location_id = l.tct_id
  INNER JOIN topics t ON t.tct_id = o.topic_id

  WHERE t.tct_id = %%topic_id int%%

GROUP BY e.title, e.author, e.publisher, e.isbn

ORDER BY number_of_occurrences DESC