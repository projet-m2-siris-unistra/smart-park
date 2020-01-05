-- PLACES (fk : place_id)
ALTER TABLE places
DROP CONSTRAINT places_device_id_fkey;

ALTER TABLE places
ADD CONSTRAINT places_device_id_fkey
FOREIGN KEY (device_id)
REFERENCES devices (device_id);


-- DEVICES (fk : tenant_id)
ALTER TABLE devices
DROP CONSTRAINT devices_tenant_id_fkey;

ALTER TABLE devices
ADD CONSTRAINT devices_tenant_id_fkey
FOREIGN KEY (tenant_id)
REFERENCES tenants (tenant_id);
