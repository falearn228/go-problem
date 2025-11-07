-- Last updated: 11/7/2025, 2:47:33 PM
-- Write your PostgreSQL query statement below
select x, y, z,(case when sum(x + y) > z and sum(x + z) > y and sum(y + z) > x then 'Yes' else 'No' end) as triangle
from Triangle
group by x , y, z