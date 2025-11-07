-- Last updated: 11/7/2025, 2:48:30 PM
-- Write your PostgreSQL query statement below
select c.name as Customers
from Customers c
where c.id not in (select customerId from Orders)