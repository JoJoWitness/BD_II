USE classicmodels;

--Total cobrado por cliente

SELECT c.customerName AS Customer, SUM(p.amount) as Total_Amount
FROM customers AS c
JOIN payments AS p ON  c.customerNumber = p.customerNumber
GROUP BY c.customerName;

--Empleado gestionador de una oficiana especifica
SELECT e.firstName AS name, e.lastName as Last_Name, e.email as Email, e.extension AS Extension, e.jobTitle AS Role
FROM employees AS e
WHERE e.officeCode = '7'; 

-- clientes que mas han gastado en los ultimos 3 meses

SELECT c.customerName AS Customer, SUM(p.amount) AS Total_Amount 
FROM customers AS c
JOIN payments AS p ON  c.customerNumber = p.customerNumber
WHERE p.paymentDate >= (
    SELECT MAX(p.paymentDate) - INTERVAL 3 MONTH
    FROM payments AS p)
GROUP BY c.customerName
ORDER BY SUM(p.amount) DESC;

--top 3 de productos mas vendidos por linea de producto
WITH CTE AS (
    SELECT  pl.productLine AS Product_Line, p.productName AS Product_Name, SUM(od.quantityOrdered) AS Quantity_Ordered,
    ROW_NUMBER() OVER (
        PARTITION BY pl.productLine
        ORDER BY SUM(od.quantityOrdered) DESC) AS row_num
    FROM productlines AS pl
    JOIN products AS p ON pl.productLine = p.productLine
    JOIN orderdetails AS od ON p.productCode = od.productCode
    GROUP BY pl.productLine, p.productName
    )
SELECT  pl.productLine AS Product_Line, p.productName AS Product_Name, SUM(od.quantityOrdered) AS Quantity_Ordered, FROM CTE WHERE row_num <= 3;

--Dia que se realizo mas ventas en un mes dado por teclado
DELIMITER //
    CREATE PROCEDURE getTopSellingDayOfAMonth(IN v_month INT)
    BEGIN
        SELECT p.paymentDate AS Date, SUM(p.amount) AS Gross_Profits, COUNT(DISTINCT o.orderNumber) AS Total_Orders
        FROM payments AS p
        JOIN customers AS c ON p.customerNumber = c.customerNumber
        JOIN orders AS o ON c.customerNumber = o.customerNumber
        WHERE EXTRACT(MONTH FROM p.paymentDate) = v_month
        GROUP BY p.paymentDate
        ORDER BY COUNT(o.orderNumber) DESC
        LIMIT 1;
    END //
DELIMITER ;

DROP PROCEDURE getTopSellingDayOfAMonth;

--Cantidad de personas que compran en una misma ciudad por mes
SELECT o.city AS city, DATE_FORMAT(p.paymentDate, "%y-%m") as Sales_Month, COUNT(DISTINCT c.customerNumber) AS Total_Clients
FROM offices AS o
JOIN employees AS e ON o.officeCode = e.officeCode
JOIN customers AS c on e.employeeNumber = c.salesRepEmployeeNumber
JOIN payments AS p on c.customerNumber = p.customerNumber
GROUP BY city, Sales_Month
Order BY city, Sales_MOnth DESC;


