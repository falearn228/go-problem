-- Last updated: 11/7/2025, 2:47:42 PM
with data as (
    select *
    from Insurance as i1
    where exists (
        select 1
        from Insurance i2
        where i1.pid <> i2.pid
        and i1.tiv_2015 = i2.tiv_2015
    )
)

SELECT ROUND(SUM(i.tiv_2016)::numeric, 2) AS tiv_2016
FROM data i
INNER JOIN (
    SELECT lat, lon 
    FROM Insurance 
    GROUP BY lat, lon 
    HAVING COUNT(*) = 1
) u ON i.lat = u.lat AND i.lon = u.lon