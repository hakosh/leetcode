select c1.seat_id
from Cinema c1
         left join Cinema c2 on c1.seat_id + 1 = c2.seat_id
         left join Cinema c3 on c1.seat_id - 1 = c3.seat_id
where (
          case
              when c1.free
                  then c1.free + ifnull(c2.free, 0) + ifnull(c3.free, 0)
              else 0
              end
          ) > 1
order by seat_id;

select distinct a.seat_id
from Cinema a
         join Cinema b
              on abs(a.seat_id - b.seat_id) = 1
where a.free
  and b.free
order by a.seat_id;