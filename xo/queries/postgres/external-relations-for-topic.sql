SELECT
  SUBSTRING(oww.content FROM '^(.*) *\([^()]+\)$') AS vocabulary,
  oww.url,
  SUBSTRING(oww.content FROM '\(([^()]+)\)$') AS relationship

FROM otx_weblink_weblink oww INNER JOIN otx_weblink_weblink_baskets owwb on oww.id = owwb.weblink_id

WHERE owwb.basket_id = %%topic_id int%%

ORDER BY vocabulary, url, relationship

