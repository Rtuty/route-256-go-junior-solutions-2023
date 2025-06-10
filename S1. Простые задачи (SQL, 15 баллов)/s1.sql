select * from problems p
where exists (select 1 from (select distinct s.user_id from submissions s
                                    inner join users u on (u.id = s.user_id)
                             where s.problem_id = p.id and s.success) q
having count(q.user_id) >= 2)
order by p.id