-- Last updated: 11/7/2025, 2:46:44 PM
-- Write your PostgreSQL query statement below
with daily_sales as (
    select visited_on, sum(amount) as total
    from Customer
    group by visited_on
),
moving_data as (
    select visited_on,
    sum(total) over (order by visited_on rows between 6 preceding and current row ) as amount ,
    avg(total) over (order by visited_on rows between 6 preceding and current row ) as avg_7d,
    count(*) over (order by visited_on rows between 6 preceding and current row ) as count_7d
    from daily_sales
)

select visited_on, amount , round(avg_7d, 2) as average_amount 
from moving_data
where count_7d >= 7
order by visited_on asc