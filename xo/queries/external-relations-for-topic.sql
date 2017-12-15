SELECT wv.vocabulary, w.url, wr.relationship

FROM topics t INNER JOIN topics_weblinks tw ON t.tct_id = tw.topic_id
  INNER JOIN weblinks w ON w.tct_id = tw.weblink_id
  INNER JOIN weblinks_relationship wr ON wr.id = w.weblinks_relationship_id
  INNER JOIN weblinks_vocabulary wv ON wv.id = w.weblinks_vocabulary_id

  WHERE t.tct_id = %%topic_id int%%

ORDER BY wv.vocabulary, w.url, wr.relationship
