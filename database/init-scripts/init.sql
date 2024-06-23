-- Insert a company
INSERT INTO public.companies (created_at, updated_at, name, logo_url, website, email, phone, address, city, state, country, postal_code, company_size)
VALUES (now(), now(), 'Sample Company', 'http://example.com/logo.png', 'http://example.com', 'info@example.com', '1234567890', '123 Example St', 'Example City', 'Example State', 'Example Country', '12345', 'Small');

-- Insert departments
INSERT INTO public.departments (created_at, updated_at, company_id, name, description)
VALUES (now(), now(), 1, 'Engineering', 'Engineering Department'),
       (now(), now(), 1, 'HR', 'Human Resources Department'),
       (now(), now(), 1, 'Sales', 'Sales Department');

-- Insert locations
INSERT INTO public.locations (created_at, updated_at, company_id, name, is_head_office, phone, address, city, state, country, postal_code)
VALUES (now(), now(), 1, 'Head Office', true, '1234567890', '123 Example St', 'Example City', 'Example State', 'Example Country', '12345'),
       (now(), now(), 1, 'Branch Office', false, '0987654321', '456 Another St', 'Another City', 'Another State', 'Another Country', '67890');

-- Insert users
-- Use the same bcrypt-hashed password for all users: 'testtesttest'
INSERT INTO public.users (created_at, updated_at, is_active, email, password, avatar_url, last_login, is_archived, first_name, middle_name, last_name, nickname)
VALUES (now(), now(), true, 'phyokyawsoe89@gmail.com', '$2a$10$58nfSXSDyfbc6kWEut3dweQmYKeRFkSi46/QN6Ay88vYgIe/w3E7i', 'http://example.com/avatar.png', now(), false, 'Root', '', 'User', 'RootUser'),
       (now(), now(), true, 'john.doe@mailnesia.com', '$2a$10$58nfSXSDyfbc6kWEut3dweQmYKeRFkSi46/QN6Ay88vYgIe/w3E7i', 'http://example.com/avatar2.png', now(), false, 'John', 'A', 'Doe', 'JohnD'),
       (now(), now(), true, 'jane.smith@mailnesia.com', '$2a$10$58nfSXSDyfbc6kWEut3dweQmYKeRFkSi46/QN6Ay88vYgIe/w3E7i', 'http://example.com/avatar3.png', now(), false, 'Jane', 'B', 'Smith', 'JaneS'),
       (now(), now(), true, 'michael.brown@mailnesia.com', '$2a$10$58nfSXSDyfbc6kWEut3dweQmYKeRFkSi46/QN6Ay88vYgIe/w3E7i', 'http://example.com/avatar4.png', now(), false, 'Michael', 'C', 'Brown', 'MikeB');

-- Insert roles
INSERT INTO public.roles (created_at, updated_at, company_id, name, description)
VALUES (now(), now(), 1, 'root', 'Root user. Full access to all features.'),
       (now(), now(), 1, 'admin', 'Admin role. Full access to manage company.'),
       (now(), now(), 1, 'manager', 'Manager role. Limited access to manage specific department.'),
       (now(), now(), 1, 'employee', 'Employee role. Basic access to personal information.');

-- Insert user roles
INSERT INTO public.user_roles (created_at, updated_at, user_id, role_id, company_id, location_id)
VALUES (now(), now(), 1, 1, 1, null),
       (now(), now(), 2, 2, 1, null),
       (now(), now(), 3, 3, 1, 1),
       (now(), now(), 4, 3, 1, 2);

-- Insert contact info
INSERT INTO public.contact_infos (created_at, updated_at, company_id, user_id, address, city, state, postal_code, country, mobile, email)
VALUES (now(), now(), 1, 1, '123 Example St', 'Example City', 'Example State', '12345', 'Example Country', '1234567890', 'contact1@example.com'),
       (now(), now(), 1, 2, '456 Another St', 'Another City', 'Another State', '67890', 'Another Country', '0987654321', 'contact2@example.com'),
       (now(), now(), 1, 3, '789 Third St', 'Third City', 'Third State', '34567', 'Third Country', '1122334455', 'contact3@example.com'),
       (now(), now(), 1, 4, '101 Fourth St', 'Fourth City', 'Fourth State', '45678', 'Fourth Country', '5566778899', 'contact4@example.com');

-- Insert emergency contacts
INSERT INTO public.emergency_contacts (created_at, updated_at, company_id, user_id, first_name, middle_name, last_name, relation, mobile, email)
VALUES (now(), now(), 1, 1, 'Jane', 'M', 'Doe', 'Spouse', '0987654321', 'emergency1@example.com'),
       (now(), now(), 1, 2, 'Emma', 'N', 'Doe', 'Mother', '1122334455', 'emergency2@example.com'),
       (now(), now(), 1, 3, 'Jack', 'O', 'Smith', 'Brother', '5566778899', 'emergency3@example.com'),
       (now(), now(), 1, 4, 'Lucy', 'P', 'Brown', 'Sister', '6677889900', 'emergency4@example.com');

-- Insert positions
INSERT INTO public.positions (created_at, updated_at, company_id, name, description, duties, qualifications, experience, min_salary, max_salary, department_id, manager_id)
VALUES (now(), now(), 1, 'Software Engineer', 'Software Engineer Position', 'Develop software', 'Bachelor degree', '2 years', 60000, 80000, 1, null),
       (now(), now(), 1, 'HR Manager', 'HR Manager Position', 'Manage HR department', 'Bachelor degree', '3 years', 50000, 70000, 2, null),
       (now(), now(), 1, 'Sales Representative', 'Sales Representative Position', 'Handle sales', 'Bachelor degree', '2 years', 40000, 60000, 3, null);

-- Insert user positions
INSERT INTO public.user_positions (created_at, updated_at, company_id, user_id, position_id, location_id, employment_status, start_date)
VALUES (now(), now(), 1, 1, 1, 1, 'Full-time', now()),
       (now(), now(), 1, 2, 2, 2, 'Full-time', now()),
       (now(), now(), 1, 3, 3, 1, 'Part-time', now()),
       (now(), now(), 1, 4, 3, 2, 'Full-time', now());

-- Insert leaves
INSERT INTO public.leaves (created_at, updated_at, company_id, user_id, type, start_date, end_date, reason, approval_status)
VALUES (now(), now(), 1, 1, 'Sick Leave', '2024-06-01', '2024-06-05', 'Medical reasons', 'Approved'),
       (now(), now(), 1, 2, 'Vacation', '2024-07-01', '2024-07-10', 'Family trip', 'Approved'),
       (now(), now(), 1, 3, 'Sick Leave', '2024-06-10', '2024-06-15', 'Flu', 'Approved'),
       (now(), now(), 1, 4, 'Personal Leave', '2024-08-01', '2024-08-05', 'Personal reasons', 'Pending');

-- Insert attendances
INSERT INTO public.attendances (created_at, updated_at, company_id, user_id, date, clock_in, clock_out, notes, approval_status)
VALUES (now(), now(), 1, 1, '2024-06-23', '2024-06-23 09:00:00', '2024-06-23 17:00:00', 'Regular workday', 'Approved'),
       (now(), now(), 1, 2, '2024-06-23', '2024-06-23 09:00:00', '2024-06-23 17:00:00', 'Regular workday', 'Approved'),
       (now(), now(), 1, 3, '2024-06-23', '2024-06-23 09:00:00', '2024-06-23 17:00:00', 'Regular workday', 'Approved'),
       (now(), now(), 1, 4, '2024-06-23', '2024-06-23 09:00:00', '2024-06-23 17:00:00', 'Regular workday', 'Approved');

-- Insert salaries
INSERT INTO public.salaries (created_at, updated_at, company_id, user_id, amount, currency, payment_interval, effective_date, approval_status)
VALUES (now(), now(), 1, 1, 70000, 'USD', 'Annually', '2024-01-01', 'Approved'),
       (now(), now(), 1, 2, 60000, 'USD', 'Annually', '2024-01-01', 'Approved'),
       (now(), now(), 1, 3, 50000, 'USD', 'Annually', '2024-01-01', 'Approved'),
       (now(), now(), 1, 4, 45000, 'USD', 'Annually', '2024-01-01', 'Approved');

-- Insert payments
INSERT INTO public.payments (created_at, updated_at, company_id, user_id, salary_id, payment_date, amount, payment_method, period_start, period_end, notes, approval_status)
VALUES (now(), now(), 1, 1, 1, '2024-06-23', 5833.33, 'Bank Transfer', '2024-06-01', '2024-06-23', 'June Salary', 'Approved'),
       (now(), now(), 1, 2, 2, '2024-06-23', 5000.00, 'Bank Transfer', '2024-06-01', '2024-06-23', 'June Salary', 'Approved'),
       (now(), now(), 1, 3, 3, '2024-06-23', 4166.67, 'Bank Transfer', '2024-06-01', '2024-06-23', 'June Salary', 'Approved'),
       (now(), now(), 1, 4, 4, '2024-06-23', 3750.00, 'Bank Transfer', '2024-06-01', '2024-06-23', 'June Salary', 'Approved');

-- Insert adjustments
INSERT INTO public.adjustments (created_at, updated_at, company_id, payment_id, amount, adjustment_type, notes)
VALUES (now(), now(), 1, 1, 100, 'Bonus', 'Performance bonus'),
       (now(), now(), 1, 2, 50, 'Penalty', 'Late submission'),
       (now(), now(), 1, 3, 150, 'Bonus', 'Overtime work'),
       (now(), now(), 1, 4, 75, 'Bonus', 'Employee of the month');

-- Insert documents
INSERT INTO public.documents (created_at, updated_at, company_id, user_id, url, notes)
VALUES (now(), now(), 1, 1, 'http://example.com/document1.pdf', 'Employee contract'),
       (now(), now(), 1, 2, 'http://example.com/document2.pdf', 'Performance review'),
       (now(), now(), 1, 3, 'http://example.com/document3.pdf', 'Training completion certificate'),
       (now(), now(), 1, 4, 'http://example.com/document4.pdf', 'Promotion letter');
