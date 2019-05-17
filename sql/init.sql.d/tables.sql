CREATE TABLE members
(
    id SERIAL,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE,
    salt VARCHAR,
    passhash VARCHAR,
    PRIMARY KEY (id)
);

CREATE TABLE customers
(
    id SERIAL,
    name TEXT UNIQUE,
    contact TEXT,
    department TEXT,
    street TEXT,
    city TEXT,
    postal TEXT,
    country TEXT,
    PRIMARY KEY (id)
);

CREATE TABLE documents
(
    id SERIAL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    reference TEXT,
    file_name TEXT,
    PRIMARY KEY (id)
);

CREATE TABLE tasks
(
    id SERIAL,
    assignee_id SERIAL,
    title TEXT,
    effort FLOAT,
    effort_unit TEXT,
    PRIMARY KEY (id),
    FOREIGN KEY (assignee_id) REFERENCES members(id)
);

CREATE TABLE projects
(
    id SERIAL,
    title TEXT,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    customer_id SERIAL,
    responsible_id SERIAL,
    accepted_at TIMESTAMP,
    paid_at TIMESTAMP,
    invoice_id SERIAL,
    PRIMARY KEY (id),
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (responsible_id) REFERENCES members(id),
    FOREIGN KEY (invoice_id) REFERENCES documents(id)
);

CREATE TABLE project_members
(
    project_id SERIAL,
    member_id SERIAL,
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (member_id) REFERENCES members(id)
);

CREATE TABLE project_offers
(
    project_id SERIAL,
    offer_id SERIAL,
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (offer_id) REFERENCES documents(id)
);
