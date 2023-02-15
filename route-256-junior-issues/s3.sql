select rank() over(order by inf.pcnt desc, inf.sbm asc) as rank,
       u.id as user_id,
       u.name as user_name,
       inf.pcnt as problem_count,
       inf.sbm as latest_successful_submitted_at
from users u, (select id from contests c order by id desc limit 1)  cmax
inner join lateral  ( --start lateral
    select count(distinct q.pcnt) pcnt, max(q.sbm) sbm from
    (select s.problem_id pcnt, min(s.submitted_at) over (partition by s.problem_id) sbm from submissions s
        inner join problems p on (p.id = s.problem_id and p.contest_id = cmax.id) where s.user_id = u.id and s.success
    ) q having exists       (
        select * from submissions s2
            inner join problems p2 on (p2.id=s2.problem_id and p2.contest_id = cmax.id)
    where s2.user_id = u.id )
                    ) inf on true --end lateral
order by inf.pcnt desc, inf.sbm asc, u.id asc