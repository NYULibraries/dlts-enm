SELECT topic_id, COUNT(*) AS number_of_occurrences
FROM occurrences
WHERE topic_id = %%topic_id int%%
GROUP BY topic_id
