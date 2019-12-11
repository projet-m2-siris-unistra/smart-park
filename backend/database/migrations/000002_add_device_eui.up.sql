ALTER TABLE devices
  ADD COLUMN device_eui char(16)
    NOT NULL DEFAULT substring(md5(random()::text), 0, 16), -- let's assume we won't have conflicts
  ADD COLUMN tenant_id INTEGER
    REFERENCES tenants (tenant_id);

UPDATE devices SET tenant_id = (SELECT tenant_id FROM tenants LIMIT 1);

ALTER TABLE devices
  ALTER COLUMN tenant_id SET NOT NULL,
  ALTER COLUMN device_eui DROP DEFAULT;
