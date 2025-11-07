-- Last updated: 11/7/2025, 2:46:29 PM
-- Write your PostgreSQL query statement below

select r.contest_id, round(count(r.user_id)  * 100.0  / (select count(*) from Users) ,2) as percentage
from Register as r
group by (r.contest_id)
order by percentage desc, r.contest_id asc
-- 