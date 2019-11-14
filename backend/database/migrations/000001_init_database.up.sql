CREATE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE tenants
(
	tenant_id  serial      PRIMARY KEY,
	name       text        NOT NULL UNIQUE,
	geo        text,       --GEOGRAPHY(point)
	created_at timestamptz NOT NULL DEFAULT NOW(),
	updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON tenants
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TYPE zonetype AS ENUM ('free', 'paid', 'blue');
CREATE TYPE devicestate AS ENUM ('free', 'occupied');

CREATE TABLE zones
(
	zone_id    serial      NOT NULL PRIMARY KEY,
	tenant_id  integer     NOT NULL REFERENCES tenants(tenant_id),
	name       text        NOT NULL,
	type       zonetype    NOT NULL,
	color      char(6)     CHECK (color ~ '^[0-9A-F]{6}$'),    --to convert to hex
	geo        text,       --GEOGRAPHY(linestring),
	created_at timestamptz NOT NULL DEFAULT NOW(),
	updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON zones
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE devices
(
	device_id  serial      PRIMARY KEY,
	battery    integer     NOT NULL CHECK(battery <= 100),
	state      devicestate DEFAULT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON devices
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE places
(
	place_id   serial      PRIMARY KEY,
	zone_id    integer     REFERENCES zones(zone_id),
	type       varchar(50) UNIQUE NOT NULL,
	geo        text,       --GEOGRAPHY(point)
	device_id  integer     REFERENCES devices(device_id),
	created_at timestamptz NOT NULL DEFAULT NOW(),
	updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON places
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE users
(
	user_id    serial      PRIMARY KEY,
	tenant_id  integer     NOT NULL REFERENCES tenants(tenant_id),
	username   varchar(50) UNIQUE CHECK(username ~ '^[a-z0-9_]+$'),
	password   char(60)    NOT NULL,
	email      varchar(50) NOT NULL UNIQUE,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	updated_at timestamptz NOT NULL DEFAULT NOW(),
    last_login timestamptz
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

COMMENT ON COLUMN users.password IS 'Bcrypt hash (Blowfish 2a)';
