delete p1.*
from Person p1
where p1.id not in (
    select id
    from (
             select min(id) as id
             from Person
             group by email
         ) as ps
)
