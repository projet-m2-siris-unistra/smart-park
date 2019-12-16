-- ZONES (fk : tenant_id)
ALTER TABLE zones
DROP CONSTRAINT zones_tenant_id_fkey;

ALTER TABLE zones
ADD CONSTRAINT zones_tenant_id_fkey
FOREIGN KEY (tenant_id)
REFERENCES tenants (tenant_id);


-- PLACES (fk : device_id)
ALTER TABLE places
DROP CONSTRAINT places_device_id_fkey;

ALTER TABLE places
ADD CONSTRAINT places_device_id_fkey
FOREIGN KEY (device_id)
REFERENCES devices (device_id);


-- PLACES (fk : zone_id)
ALTER TABLE places
DROP CONSTRAINT places_zone_id_fkey;

ALTER TABLE places
ADD CONSTRAINT places_zone_id_fkey
FOREIGN KEY (zone_id)
REFERENCES zones (zone_id);


-- USERS (fk : tenant_id)
ALTER TABLE users
DROP CONSTRAINT users_tenant_id_fkey;

ALTER TABLE users
ADD CONSTRAINT users_tenant_id_fkey
FOREIGN KEY (tenant_id)
REFERENCES tenants (tenant_id);


-- DEVICES (fk : tenant_id)
ALTER TABLE devices
DROP CONSTRAINT devices_tenant_id_fkey;

ALTER TABLE devices
ADD CONSTRAINT devices_tenant_id_fkey
FOREIGN KEY (tenant_id)
REFERENCES tenants (tenant_id);