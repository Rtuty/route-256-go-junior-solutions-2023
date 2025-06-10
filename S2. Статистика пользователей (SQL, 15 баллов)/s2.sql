select u.id, u.name,
        (select count(distinct c.id) from submissions s
                    inner join problems p on (p.id = s.problem_id)
                    inner join contests c on (c.id = p.contest_id)
               where s.user_id = u.id and s.success
        )        as solved_at_least_one_contest_count,
        (select count(distinct c.id) from submissions s
                    inner join problems p on (p.id = s.problem_id)
                    inner join contests c on (c.id = p.contest_id)
               where s.user_id = u.id
        )        as take_part_contest_count
from users u
order by solved_at_least_one_contest_count desc, take_part_contest_count desc, u.id
