SELECT hb.id AS tct_id, hb.display_name, hh.name
FROM hit_basket hb INNER JOIN hit_hit hh ON hb.id = hh.basket_id
ORDER BY display_name, name
