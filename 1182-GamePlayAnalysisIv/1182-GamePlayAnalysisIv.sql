-- Last updated: 11/7/2025, 2:47:01 PM
-- # Write your MySQL query statement below
-- # Event_date + 1 = Event_date
-- # games_played <> 0
-- # player_id = player_id 
with first_login as (
    select player_id,
    min(event_date) as first_login_date
    from Activity
    group by player_id
)


select round(count(distinct a.player_id) * 1.0 / (select count(distinct player_id) from Activity), 2) as fraction
from Activity as a
join first_login as ac on a.player_id = ac.player_id
and a.event_date - interval '1 day' = ac.first_login_date
and a.games_played <> 0