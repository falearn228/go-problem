-- Last updated: 11/7/2025, 2:46:22 PM
-- Write your PostgreSQL query statement below
select user_id, count(follower_id) as followers_count
from Followers
group by user_id
order by 1;