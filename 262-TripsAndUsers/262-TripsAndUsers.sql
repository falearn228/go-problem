-- Last updated: 11/7/2025, 2:48:09 PM
-- Write your PostgreSQL query statement below

-- completed
-- banned = No
-- /

select t.request_at as "Day", round(count(case when t.status in ('cancelled_by_driver', 'cancelled_by_client') then 1 else null end) * 1.0 / count(*)::NUMERIC, 2) as "Cancellation Rate"
from Trips as t
join Users as u1 on t.client_id = u1.users_id and u1.banned = 'No'
join Users as u2 on t.driver_id = u2.users_id and u2.banned = 'No'
WHERE t.request_at BETWEEN '2013-10-01' AND '2013-10-03'
group by t.request_at