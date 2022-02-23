select wc.id
from Weather wc
         join Weather wp on date_sub(wc.recordDate, interval 1 day) = wp.recordDate
where wc.temperature > wp.temperature