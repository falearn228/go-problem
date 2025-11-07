-- Last updated: 11/7/2025, 2:48:31 PM
-- Write your PostgreSQL query statement below
select email
from Person
group by email
having count(*) > 1