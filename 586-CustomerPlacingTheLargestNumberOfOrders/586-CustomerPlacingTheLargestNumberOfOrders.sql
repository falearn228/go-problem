-- Last updated: 11/7/2025, 2:47:41 PM
-- Write your PostgreSQL query statement below

select customer_number as customer_number
from (
    select customer_number, count(customer_number) as counts
    from Orders
    group by customer_number
)
order by counts desc
limit 1