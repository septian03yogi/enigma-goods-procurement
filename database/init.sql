CREATE TABLE employees (
	id varchar(255) NOT NULL,
	employee_name varchar(255) NOT NULL,
	phone_number varchar(20) NOT NULL,
	department_id varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	is_delete bool NOT NULL DEFAULT false,
	CONSTRAINT employees_pkey PRIMARY KEY (id),
	CONSTRAINT employees_department_id_fkey FOREIGN KEY (department_id) REFERENCES departments(id)
);

CREATE TABLE departments (
	id varchar(255) NOT NULL,
	department_name varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT departments_pkey PRIMARY KEY (id)
);

CREATE TABLE role_users (
	id varchar(255) NOT NULL,
	role_name varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT role_users_pkey PRIMARY KEY (id)
);

CREATE TABLE users (
	id varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	role_id varchar(255) NOT NULL,
	employee_id varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES role_users(id)
);

CREATE TABLE items (
	id varchar(255) NOT NULL,
	item_name varchar(255) NOT NULL,
	stock int4 NOT NULL,
	uom_id varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT items_pkey PRIMARY KEY (id),
	CONSTRAINT items_uom_id_fkey FOREIGN KEY (uom_id) REFERENCES uoms(id)
);

CREATE TABLE uoms (
	id varchar(255) NOT NULL,
	uom_name varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT uoms_pkey PRIMARY KEY (id)
);

CREATE TABLE periods (
	id varchar(255) NOT NULL,
	start_date timestamp NOT NULL,
	end_date timestamp NOT NULL,
	period_name varchar(255) NOT NULL,
	is_active bool NOT NULL DEFAULT true,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT periods_pkey PRIMARY KEY (id)
);

CREATE TABLE submissions (
	id varchar(255) NOT NULL,
	submission_name varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT submissions_pkey PRIMARY KEY (id)
);

CREATE TABLE item_submissions (
	id varchar(255) NOT NULL,
	employee_id varchar(255) NOT NULL,
	period_id varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT item_submissions_pkey PRIMARY KEY (id),
	CONSTRAINT item_submissions_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES employees(id),
	CONSTRAINT item_submissions_period_id_fkey FOREIGN KEY (period_id) REFERENCES periods(id)
);

CREATE TABLE item_submission_details (
	id varchar(255) NOT NULL,
	submission_id varchar(255) NOT NULL,
	item_id varchar(255) NOT NULL,
	amount_submit int4 NOT NULL,
	status_details varchar(255) NOT NULL,
	amount_approve int4 NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL);
