-- Last updated: 11/7/2025, 2:47:30 PM
-- Write your PostgreSQL query statement below
with swapped as (
    select id, student, lag(student, 1) over (order by id) as prev_data
    from Seat
)

select id, (
    case
        when id % 2 = 0 then prev_data
        else coalesce(lead(student, 1) over (order by id), student)
        end
) as student
from swapped