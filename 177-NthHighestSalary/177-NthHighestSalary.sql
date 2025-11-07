-- Last updated: 11/7/2025, 2:48:35 PM
CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
BEGIN
  RETURN QUERY (
    with rank_salary as (
        select e.salary, dense_rank() over (order by e.salary desc) as ranked_salary
        from Employee e
    )
    
    select r.salary
    from rank_salary r
    where r.ranked_salary = N
    limit 1
  );
END;
$$ LANGUAGE plpgsql;