-- Last updated: 11/7/2025, 2:47:07 PM
-- Write your PostgreSQL query statement below
select customer_id
from Customer
group by customer_id
having count(distinct product_key) = (
    select count(*)
    from Product
)