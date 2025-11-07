-- Last updated: 11/7/2025, 2:47:43 PM
-- Write your PostgreSQL query statement below
SELECT name
FROM Customer
WHERE referee_id <> 2 OR referee_id IS NULL