# TENTATIVE ROUTES

## Admin (Company Management)

### Company Setup and Management

#### Company Profile

FE: /admin/company/profile
BE: /api/v1/admin/company

#### Locations

FE: /admin/locations
BE: /api/v1/admin/locations

#### Departments

FE: /admin/departments
BE: /api/v1/admin/departments

#### Positions

FE: /admin/positions
BE: /api/v1/admin/positions

#### User Accounts

FE: /admin/users
BE: /api/v1/admin/users

### Attendance Management

#### Attendance Policies

FE: /admin/attendance/policies
BE: /api/v1/admin/attendance/policies

#### Integration Setup

FE: /admin/attendance/integration
BE: /api/v1/admin/attendance/integration

### Leave Management

#### Leave Policies

FE: /admin/leave/policies
BE: /api/v1/admin/leave/policies

#### Leave Reports

FE: /admin/leave/reports
BE: /api/v1/admin/leave/reports

### Payroll Management

#### Salary Configuration

FE: /admin/payroll/salary-configuration
BE: /api/v1/admin/payroll/salary-configuration

#### Tax Setup

FE: /admin/payroll/tax-setup
BE: /api/v1/admin/payroll/tax-setup

#### Payslip Templates

FE: /admin/payroll/payslip-templates
BE: /api/v1/admin/payroll/payslip-templates

### Compliance Management

#### Document Repository

FE: /admin/compliance/documents
BE: /api/v1/admin/compliance/documents

#### Legal Alerts

FE: /admin/compliance/alerts
BE: /api/v1/admin/compliance/alerts

### Reports and Analytics

#### Comprehensive Reports

FE: /admin/reports
BE: /api/v1/admin/reports

#### Data Export

FE: /admin/data/export
BE: /api/v1/admin/data/export

## Manager (Location/Branch Management)

### Employee Management

#### Employee Profiles

FE: /manager/employees
BE: /api/v1/manager/employees

#### Position Assignment

FE: /manager/positions
BE: /api/v1/manager/positions

### Attendance Management

#### Attendance Tracking

FE: /manager/attendance
BE: /api/v1/manager/attendance

#### Attendance Reports

FE: /manager/attendance/reports
BE: /api/v1/manager/attendance/reports

### Leave Management

#### Leave Requests

FE: /manager/leave/requests
BE: /api/v1/manager/leave/requests

#### Leave Tracking

FE: /manager/leave/tracking
BE: /api/v1/manager/leave/tracking

### Claims Management

#### Claims Requests

FE: /manager/claims/requests
BE: /api/v1/manager/claims/requests

#### Claims Tracking

FE: /manager/claims/tracking
BE: /api/v1/manager/claims/tracking

### Documents Management

#### Document Submission

FE: /manager/documents
BE: /api/v1/manager/documents

### Compliance Management

#### Compliance Monitoring

FE: /manager/compliance/monitoring
BE: /api/v1/manager/compliance/monitoring

#### Local Compliance Alerts

FE: /manager/compliance/alerts
BE: /api/v1/manager/compliance/alerts

## User (Employee)

### Profile Management

#### Personal Profile

FE: /user/profile
BE: /api/v1/user/profile

#### Document Upload

FE: /user/documents
BE: /api/v1/user/documents

### Attendance Management

#### View Attendance

FE: /user/attendance
BE: /api/v1/user/attendance

#### Attendance Correction Requests

FE: /user/attendance/correction
BE: /api/v1/user/attendance/correction

### Leave Management

#### Leave Requests

FE: /user/leave/requests
BE: /api/v1/user/leave/requests

#### Leave Status

FE: /user/leave/status
BE: /api/v1/user/leave/status

### Payroll Management

#### View Payslips

FE: /user/payslips
BE: /api/v1/user/payslips

#### Salary Details

FE: /user/salary/details
BE: /api/v1/user/salary/details

### Claims Management

#### Submit Claims

FE: /user/claims
BE: /api/v1/user/claims

#### View Claim Status

FE: /user/claims/status
BE: /api/v1/user/claims/status

### Compliance Management

#### Compliance Updates

FE: /user/compliance/updates
BE: /api/v1/user/compliance/updates

### Communication and Feedback

#### Internal Messaging

FE: /user/messaging
BE: /api/v1/user/messaging

#### Feedback Submission

FE: /user/feedback
BE: /api/v1/user/feedback

src/
├── assets/
├── components/
├── views/
│ ├── admin/
│ │ ├── CompanyProfile.vue
│ │ ├── Locations.vue
│ │ ├── Departments.vue
│ │ ├── Positions.vue
│ │ ├── Users.vue
│ │ ├── AttendancePolicies.vue
│ │ ├── AttendanceIntegration.vue
│ │ ├── LeavePolicies.vue
│ │ ├── LeaveReports.vue
│ │ ├── SalaryConfiguration.vue
│ │ ├── TaxSetup.vue
│ │ ├── PayslipTemplates.vue
│ │ ├── ComplianceDocuments.vue
│ │ ├── LegalAlerts.vue
│ │ ├── Reports.vue
│ │ ├── DataExport.vue
│ ├── manager/
│ │ ├── Employees.vue
│ │ ├── Positions.vue
│ │ ├── Attendance.vue
│ │ ├── AttendanceReports.vue
│ │ ├── LeaveRequests.vue
│ │ ├── LeaveTracking.vue
│ │ ├── ClaimsRequests.vue
│ │ ├── ClaimsTracking.vue
│ │ ├── Documents.vue
│ │ ├── ComplianceMonitoring.vue
│ │ ├── ComplianceAlerts.vue
│ ├── user/
│ │ ├── Profile.vue
│ │ ├── Documents.vue
│ │ ├── Attendance.vue
│ │ ├── AttendanceCorrection.vue
│ │ ├── LeaveRequests.vue
│ │ ├── LeaveStatus.vue
│ │ ├── Payslips.vue
│ │ ├── SalaryDetails.vue
│ │ ├── Claims.vue
│ │ ├── ClaimsStatus.vue
│ │ ├── ComplianceUpdates.vue
│ │ ├── Messaging.vue
│ │ ├── Feedback.vue
├── router/
│ ├── index.js
├── store/
├── App.vue
├── main.js
