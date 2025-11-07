-- Last updated: 11/7/2025, 2:46:33 PM
-- Write your PostgreSQL query statement below

select sell_date, count(distinct product) as num_sold,
string_agg(distinct product, ',' order by product) as products
from Activities
group by sell_date
order by sell_date
