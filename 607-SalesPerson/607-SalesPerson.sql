-- Last updated: 11/7/2025, 2:47:34 PM
-- Write your PostgreSQL query statement below

select s.name
from SalesPerson as s
where s.sales_id not in (select o.sales_id from Orders o join Company c on o.com_id = c.com_id
where c.name = 'RED')
 