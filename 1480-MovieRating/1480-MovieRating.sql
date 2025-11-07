-- Last updated: 11/7/2025, 2:46:41 PM
-- Write your PostgreSQL query statement below
(
    select u.name as results
    from Users as u
    join (
        select user_id , count(*) as return_count
        from MovieRating
        group by user_id
    ) as feb on u.user_id = feb.user_id
    order by feb.return_count desc, u.name asc
    limit 1
)
UNION ALL

(
    select m.title as results
    from Movies as m
    join (
        select movie_id , avg(rating) as average_rating
        from MovieRating
        WHERE created_at BETWEEN '2020-02-01' AND '2020-02-29'
        group by movie_id
    ) as feb on m.movie_id = feb.movie_id
    order by feb.average_rating desc, m.title asc
    limit 1
)