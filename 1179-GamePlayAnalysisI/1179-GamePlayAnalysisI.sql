-- Last updated: 11/7/2025, 2:47:02 PM
-- Write your PostgreSQL query statement below
select player_id , min(event_date) as first_login
from Activity
group by 1