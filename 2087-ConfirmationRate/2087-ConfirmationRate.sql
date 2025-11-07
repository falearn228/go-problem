-- Last updated: 11/7/2025, 2:46:15 PM
-- Write your PostgreSQL query statement below
select s.user_id,
    Round(
    case
        when count(c.user_id) = 0 then 0.00
        else sum(case when c.action = 'confirmed' then 1 else 0 end) * 1.0 / count(c.user_id)
    end, 2) as confirmation_rate
from Signups as s
left join Confirmations as c
on c.user_id = s.user_id
group by s.user_id