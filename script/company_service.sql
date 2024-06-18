CREATE TABLE pos_companies (
    company_id UUID PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_store_branches (
    branch_id UUID PRIMARY KEY,
    branch_name VARCHAR(255) NOT NULL,
    company_id UUID REFERENCES pos_companies(company_id) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_stores (
    store_id UUID PRIMARY KEY,
    store_name VARCHAR(255) NOT NULL,
    branch_id UUID REFERENCES pos_store_branches(branch_id) NOT NULL,
    location VARCHAR(255),
    company_id UUID REFERENCES pos_companies(company_id) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_roles (
    role_id UUID PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_users (
    user_id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role_id UUID REFERENCES pos_roles(role_id) NOT NULL,
    company_id UUID REFERENCES pos_companies(company_id),
    branch_id UUID REFERENCES pos_store_branches(branch_id),
    store_id UUID REFERENCES pos_stores(store_id),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone_number VARCHAR(20),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);
