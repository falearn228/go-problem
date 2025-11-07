-- Last updated: 11/7/2025, 2:46:25 PM
-- Write your PostgreSQL query statement below
select user_id, concat(upper(substring(name, 1, 1)), lower(substring(name, 2))) as name
from Users
order by user_id