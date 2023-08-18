CREATE TABLE department (
	id varchar(255) NOT NULL,
	department_name varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT departments_pkey PRIMARY KEY (id)
);

CREATE TABLE employee (
	id varchar(255) NOT NULL,
	employee_name varchar(255) NOT NULL,
	phone_number varchar(20) NOT NULL,
	department_id varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	is_delete bool NOT NULL DEFAULT false,
	CONSTRAINT employees_pkey PRIMARY KEY (id),
	CONSTRAINT employees_department_id_fkey FOREIGN KEY (department_id) REFERENCES department(id)
);

CREATE TABLE role_user (
	id varchar(255) NOT NULL,
	role_name varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT role_users_pkey PRIMARY KEY (id)
);

CREATE TABLE user_credential (
	id varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	role_id varchar(255) NOT NULL,
	employee_id varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES role_user(id)
);

CREATE TABLE uom (
	id varchar(255) NOT NULL,
	uom_name varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT uoms_pkey PRIMARY KEY (id)
);

CREATE TABLE item (
	id varchar(255) NOT NULL,
	item_name varchar(255) NOT NULL,
	stock int4 NOT NULL,
	uom_id varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT items_pkey PRIMARY KEY (id),
	CONSTRAINT items_uom_id_fkey FOREIGN KEY (uom_id) REFERENCES uom(id)
);

CREATE TABLE period (
	id varchar(255) NOT NULL,
	period_name varchar(255) NOT NULL,
	start_date timestamp NOT NULL,
	end_date timestamp NOT NULL,
	is_active bool NOT NULL DEFAULT true,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT periods_pkey PRIMARY KEY (id)
);

CREATE TABLE submission (
	id varchar(255) NOT NULL,
	employee_id varchar(255) NOT NULL,
	period_id varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT item_submissions_pkey PRIMARY KEY (id),
	CONSTRAINT item_submissions_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES employee(id),
	CONSTRAINT item_submissions_period_id_fkey FOREIGN KEY (period_id) REFERENCES period(id)
);

CEATE TABLE submission_status (
	id varchar(255),
	status_detail varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT submissions_status_pkey PRIMARY KEY (id)
)

CREATE TABLE submission_detail (
	id varchar(255) NOT NULL,
	submission_id varchar(255) NOT NULL REFERENCES submission(id),
	item_id varchar(255) NOT NULL REFERENCES item(id),
	amount_submit int4 NOT NULL,
	status_id varchar(255) NOT NULL REFERENCES submission_status(id),
	amount_approve int4 NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
);
