-- Last updated: 11/7/2025, 2:46:26 PM
-- Write your PostgreSQL query statement below
with data as (
    select
    machine_id,
    process_id,
    max(case when activity_type = 'end' then timestamp end) -
    max(case when activity_type = 'start' then timestamp end) as processing_time
    from activity
    group by machine_id, process_id
)

select machine_id, Round(avg(processing_time)::numeric, 3) as processing_time
from data
group by machine_id
order by machine_id