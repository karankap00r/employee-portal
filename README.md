# Workday like Employee Portal

## Problem 1: Employee Database Management System

Develop a simple Employee Database Management System using Object-Oriented
Programming

```text
Class Employee
    Attributes (all should be immutable and strongly typed):
        employee_id: Final[UUID]
        name: Final[str]
        position: Final[str]
        email: Final[str]
        salary: Final[float]
        created_at: Final[datetime]
        modified_at: Final[datetime]
        
    Methods:
       - A constructor to initialize all attributes.      
       - __str__ method to return a string representation of the employee.
```

```text
Class EmployeeDatabase
    Attributes:
        A private attribute to store employee objects.
    
    Methods:
       - add_employee: Adds a new Employee to the database.
       - update_employee: Updates an existing Employee to the database.
       - get_employee: Retrieves an Employee's information by ID.
       - get_all: Returns all the employees.
       - __str__: Returns a string representation of all employees in the database.
```

### Complexities
1. Although the database functionality will work locally due to the
   scope of this exercise, it must simulate is a remote server with a different timezone,
   Therefore, create a utility class/function to supply the datetime as if there was a different timezone.
2. Whenever returning any employee, the created/modified fields should be
   returned but transformed in local datetime, considering the mocked server
   timezone.
3. The email format must be checked, and raise an exception if invalid.

### Notes:
1. Results should be dumped into the console.
2. For the database you can use an embedded one such as H2 or SQLite.

## Problem 2: Time-off module

In our Time Off module we have Time Off Categories that you can use to
request the right time for those categories as a team member.<br>

We added a rule that you can’t request two categories at the same time with
the same date range. However, there is a case that you can have Work Remotely
applicable, and you want to request Annual leave.

Right now the constraint of two categories can’t be overlapping is present
since you can’t request annual leave while we have another category within
the same dates.

![Timeoff Module](resources/timeoff.png)
<p style="text-align: center;">Figure 1: Timeoff Module Representation</p>

### Database Tables

```text
    - Time_off_request
        - Id
        - request_category_id
        - employee_id
        - start_date
        - end_date
    
    - Request_category
        - id
        - Name
```

## Problem 3: Public Holiday Service

**Goal**: Create a service that will provide the public holidays from the employees that you have from exercise 1 (create fake employees).

Requirements:
* Service should return the current 7 days public holidays for the employees' residence location
* Create an email alert with the current public holidays in the next 7 days upcoming
* Store from a 3rd party API or you parse it online to get the source of truth of the public holidays

## Postman Collection:
Please refer `resources/Employee Management System.postman_collection.json` for the API documentation.

## Possible Improvements:
1. Fetching configuration from a file
2. Addition of monitoring and alerting
3. Load Testing the endpoints
4. Adding more features like leave history, etc.
5. Making storage persistent
6. Adding more validations
7. Adding more test cases
8. Writing cron for sending email alerts for next 7 days public holidays
9. Make the "7" day configuration driven.