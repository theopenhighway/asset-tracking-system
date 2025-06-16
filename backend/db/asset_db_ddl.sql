CREATE TYPE roles AS ENUM ('admin', 'user', 'employee');
CREATE TYPE asset_type as ENUM ('laptop', 'asdfa', 'asdfasdfa');
CREATE TYPE approval_status as ENUM ('pending', 'approved', 'rejected', 'cancelled');
CREATE TYPE asset_status as ENUM ('available', 'maintenance', 'damaged', 'retired');

CREATE TABLE "departments" (
    "id" serial primary key,
    "name" text unique not null
);

CREATE TABLE "users" (
    "id" serial primary key,
    "name" text not null,
    "email" text not null,
    "pw_hash" text not null,
    "role" roles not null,
    "department_id" integer not null references departments('id'),
    "is_active" boolean default true,
    "created_at" timestamp default now(),
    "updated_at" timestamp
);

CREATE TABLE "assets" (
    "id" integer,
    "name" text not null,
    "asset_type" asset_type not null,
    "status" asset_status default 'available',
    "description" varchar(100) not null,
    "purchase_date" timestamp not null,
    "updated_at" timestamp default now()  
);

CREATE TABLE "asset_assignments" (
    "id" serial primary key,
    "asset_id" integer not null references assets('id'),
    "user_id" integer not null references users('id'),
    "assigned_at" timestamp default now(),
    "returned_at" timestamp,
    "assigned_by" integer not null
);

CREATE TABLE "requests" (
    "id" serial primary key,
    "user_id" integer not null references users('id'),
    "asset_type" asset_type,
    "reason" text not null,
    "status" approval_status default 'pending',
    "approved_by" integer not null, 
    "created_at" timestamp default now(),
    "reviewed_at" timestamp default now()
);