-- Last updated: 11/7/2025, 2:48:27 PM
-- Write your PostgreSQL query statement below
delete from Person a
    using Person b
    where a.id > b.id and a.email = b.email