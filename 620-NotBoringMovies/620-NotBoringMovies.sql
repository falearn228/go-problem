-- Last updated: 11/7/2025, 2:47:31 PM
-- Write your PostgreSQL query statement below
select id, movie, description, rating
from Cinema
where description NOT LIKE 'boring' and id % 2 = 1
order by rating desc