CREATE TYPE roles AS ENUM ('admin', 'user', 'employee');
CREATE TYPE asset_type as ENUM ('laptop', 'dev_equipment', 'iot_dev_board', 'test_equipment', 'sensor');
CREATE TYPE approval_status as ENUM ('pending', 'approved', 'rejected', 'cancelled');
CREATE TYPE asset_status as ENUM ('available', 'maintenance', 'damaged', 'retired');

CREATE TABLE IF NOT EXISTS "departments" (
    "id" serial primary key,
    "name" text unique not null
);

CREATE TABLE IF NOT EXISTS "users" (
    "id" serial primary key,
    "name" text not null,
    "email" text not null,
    "pw_hash" text not null,
    "role" roles not null,
    "department_id" integer not null references departments('id'),
    "is_active" boolean default true,
    "created_at" timestamp default now(),
    "updated_at" timestamp default now()
);

CREATE TABLE IF NOT EXISTS "assets" (
    "id" integer,
    "name" text not null,
    "asset_type" asset_type not null,
    "status" asset_status default 'available',
    "detail" jsonb,
    "purchase_date" timestamp not null,
    "updated_at" timestamp default now()  
);

CREATE TABLE IF NOT EXISTS "asset_assignments" (
    "id" serial primary key,
    "asset_id" integer not null references assets('id'),
    "user_id" integer not null references users('id'),
    "assigned_at" timestamp default now(),
    "returned_at" timestamp,
    "assigned_by" integer not null references users('id')
);

CREATE TABLE IF NOT EXISTS "requests" (
    "id" serial primary key,
    "user_id" integer not null references users('id'),
    "asset_type" asset_type,
    "reason" text not null,
    "status" approval_status default 'pending',
    "approved_by" integer not null references users('id'), 
    "created_at" timestamp default now(),
    "reviewed_at" timestamp default now()
);