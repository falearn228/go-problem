-- Last updated: 11/7/2025, 2:46:43 PM
-- Write your PostgreSQL query statement below


select p.product_name, sum(o.unit) as unit
from Products p
join (
    select product_id, order_date, unit
    from Orders
    where order_date > '2020-1-31' and order_date < '2020-03-01'
) as o using (product_id)
group by p.product_name
having sum(o.unit) >= 100