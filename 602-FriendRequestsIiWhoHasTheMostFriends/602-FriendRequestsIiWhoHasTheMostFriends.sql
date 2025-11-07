-- Last updated: 11/7/2025, 2:47:37 PM
-- Write your PostgreSQL query statement below
with requestors as (
    select requester_id as id
    from RequestAccepted

    union all

    select accepter_id as id
    from RequestAccepted
)

select id, count(*) as num
from requestors
group by id
order by num desc
limit 1
