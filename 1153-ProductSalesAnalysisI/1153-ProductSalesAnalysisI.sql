-- Last updated: 11/7/2025, 2:47:05 PM
-- Write your PostgreSQL query statement below
select p.product_name, s.year, s.price
from Sales as s join Product as p
using(product_id) 