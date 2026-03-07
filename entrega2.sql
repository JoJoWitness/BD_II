USE classicmodels;

-- Total cobrado por cliente

SELECT c.customerName, SUM(p.amount) 
FROM customers AS c
JOIN payments AS p ON  c.customerNumber = p.customerNumber
GROUP BY c.customerName;

--Empleado gestionador de una oficiana especifica
SELECT e.firstName, e.lastName, e.email, e.extension, e.jobTitle
FROM employees AS e
WHERE e.officeCode = '1'; 

-- clientes que mas han gastado en los ultimos 3 meses

SELECT c.customerName, SUM(p.amount) 
FROM customers AS c
JOIN payments AS p ON  c.customerNumber = p.customerNumber
WHERE DATEDIFF(p.customerDate, 
    (SELECT EXTRACT(MONTH FROM p.paymentDate), EXTRACT(YEAR FROM p.paymentDate)
    FROM payments AS p
    ORDER BY p.paymentDate DESC
    LIMIT 1)
    ) <=3
GROUP BY c.customerName
ORDER BY SUM(p.amount)DESC;

SELECT 

SELECT * from payments;