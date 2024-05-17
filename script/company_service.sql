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
    company_id UUID REFERENCES pos_companies(company_id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by UUID,
    updated_at TIMESTAMP,
    updated_by UUID
);

CREATE TABLE pos_stores (
    store_id UUID PRIMARY KEY,
    store_name VARCHAR(255) NOT NULL,
    branch_id UUID REFERENCES pos_store_branches(branch_id),
    location VARCHAR(255),
    company_id UUID REFERENCES pos_companies(company_id),
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
    role_id UUID REFERENCES pos_roles(role_id),
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


-- Insert into pos_companies
INSERT INTO pos_companies (company_id, company_name)
VALUES ('550e8400-e29b-41d4-a716-446655440000', 'PT Trans Retail Indonesia');

-- Insert into pos_store_branches
INSERT INTO pos_store_branches (branch_id, branch_name, company_id)
VALUES 
('6ba7b810-9dad-11d1-80b4-00c04fd430c8', 'Carrefour', '550e8400-e29b-41d4-a716-446655440000'),
('6ba7b811-9dad-11d1-80b4-00c04fd430c8', 'Transmart Carrefour', '550e8400-e29b-41d4-a716-446655440000'),
('6ba7b812-9dad-11d1-80b4-00c04fd430c8', 'Metro', '550e8400-e29b-41d4-a716-446655440000'),
('6ba7b813-9dad-11d1-80b4-00c04fd430c8', 'Grosir Indonesia', '550e8400-e29b-41d4-a716-446655440000');

-- Insert into pos_stores
INSERT INTO pos_stores (store_id, store_name, branch_id, location, company_id)
VALUES 
('550e8400-e29b-41d4-a716-446655440001', 'Carrefour Store 1', '6ba7b810-9dad-11d1-80b4-00c04fd430c8', 'Jakarta, Indonesia', '550e8400-e29b-41d4-a716-446655440000'),
('550e8400-e29b-41d4-a716-446655440002', 'Transmart Carrefour Store 1', '6ba7b811-9dad-11d1-80b4-00c04fd430c8', 'Jakarta, Indonesia', '550e8400-e29b-41d4-a716-446655440000'),
('550e8400-e29b-41d4-a716-446655440003', 'Metro Store 1', '6ba7b812-9dad-11d1-80b4-00c04fd430c8', 'Jakarta, Indonesia', '550e8400-e29b-41d4-a716-446655440000'),
('550e8400-e29b-41d4-a716-446655440004', 'Grosir Indonesia Store 1', '6ba7b813-9dad-11d1-80b4-00c04fd430c8', 'Jakarta, Indonesia', '550e8400-e29b-41d4-a716-446655440000');

-- Insert into pos_roles
INSERT INTO pos_roles (role_id, role_name)
VALUES 
('550e8400-e29b-41d4-a716-446655440005', 'Company'),
('550e8400-e29b-41d4-a716-446655440006', 'Branch'),
('550e8400-e29b-41d4-a716-446655440007', 'Store');

-- Insert into pos_users
INSERT INTO pos_users (user_id, username, password_hash, role_id, company_id, branch_id, store_id, first_name, last_name, email, phone_number)
VALUES 
('550e8400-e29b-41d4-a716-446655440008', 'company_user', 'hashed_password', '550e8400-e29b-41d4-a716-446655440005', '550e8400-e29b-41d4-a716-446655440000', NULL, NULL, 'Ahmad', 'Fajar', 'ahmad.fajare@example.com', '1234567890'),
('550e8400-e29b-41d4-a716-446655440009', 'branch_user', 'hashed_password', '550e8400-e29b-41d4-a716-446655440006', '550e8400-e29b-41d4-a716-446655440000', '6ba7b810-9dad-11d1-80b4-00c04fd430c8', NULL, 'dafa', 'puja', 'dafa.puja@example.com', '0987654321'),
