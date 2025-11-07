-- Last updated: 11/7/2025, 2:47:32 PM
-- Write your PostgreSQL query statement below
select max(num) as num
from (
    select num
    from MyNumbers
    group by num
    having (count(*) = 1)
)