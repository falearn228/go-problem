-- Last updated: 11/7/2025, 2:46:52 PM
-- Write your PostgreSQL query statement below
with ranked_turn as (
    select
    person_name,
    sum(weight) over (order by turn) as total,
    turn
    from Queue
) 

select person_name
from ranked_turn
where total <= 1000
order by total desc
limit 1;
