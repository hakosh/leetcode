select C.customer_id, C.name
from Customers C
         inner join orders using (customer_id)
         inner join product using (product_id)
where year(order_date) = 2020
group by C.customer_id
having sum(if(month(order_date) = 6, price * quantity, 0)) >= 100
   and sum(if(month(order_date) = 7, price * quantity, 0)) >= 100
