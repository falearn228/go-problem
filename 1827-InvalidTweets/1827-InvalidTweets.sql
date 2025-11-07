-- Last updated: 11/7/2025, 2:46:24 PM
-- Write your PostgreSQL query statement below
SELECT tweet_id
FROM Tweets
WHERE LENGTH(content) > 15