-- Last updated: 11/7/2025, 2:46:50 PM
-- Write your PostgreSQL query statement below
select p.product_id, round(
    coalesce(sum(p.price * u.units) * 1.0 / sum(u.units), 0), 2) as average_price
from Prices as p
left join UnitsSold as u
on p.product_id = u.product_id
and u.purchase_date >= p.start_date
and u.purchase_date <= p.end_date
group by p.product_id