select Employee.name as Employee
from Employee
         inner join Employee as Manager on Employee.managerId = Manager.id
Where Manager.salary < Employee.salary