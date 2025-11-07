-- Last updated: 11/7/2025, 2:46:59 PM
select activity_date as day,
count(distinct user_id) as active_users
from activity
where activity_date between '2019-06-28' AND DATE '2019-07-27'
GROUP BY 
    activity_date
ORDER BY 
    activity_date;