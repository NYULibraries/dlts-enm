SELECT basket_id, COUNT(*) AS number_of_occurrences
FROM occurrence_occurrence
WHERE basket_id = %%topic_id int%%
GROUP BY basket_id
