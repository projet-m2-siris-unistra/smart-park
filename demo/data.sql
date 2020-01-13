DELETE FROM users;
DELETE FROM places;
DELETE FROM devices;
DELETE FROM zones;
DELETE FROM tenants;
ALTER SEQUENCE tenants_tenant_id_seq RESTART;
ALTER SEQUENCE zones_zone_id_seq RESTART;
ALTER SEQUENCE devices_device_id_seq RESTART;
ALTER SEQUENCE places_place_id_seq RESTART;
ALTER SEQUENCE users_user_id_seq RESTART;

/****************************** TENANTS *******************************/
INSERT INTO tenants (name, geo) VALUES
('Schmilbligheim','[7.7475,48.5827]') ON CONFLICT DO NOTHING;

/****************************** TENANTS *******************************/

/****************************** ZONES *******************************/
INSERT INTO zones (tenant_id, name,type,color,geo) VALUES
((SELECT MIN(tenant_id) FROM tenants),'Krutenau','paid','FF00E2','[[7.751412,48.579191],[7.753987,48.575301],[7.763557,48.574989],[7.757721,48.582939],[7.751412,48.579191]]') ON CONFLICT DO NOTHING;
INSERT INTO zones (tenant_id,name,type,color,geo) VALUES
((SELECT MIN(tenant_id) FROM tenants),'Centre-ville','paid','001CF2','[[7.740726,48.584472],[7.748709,48.578907],[7.75712,48.582939],[7.748795,48.587481],[7.740726,48.584472]]') ON CONFLICT DO NOTHING;
INSERT INTO zones (tenant_id,name,type,color,geo) VALUES
((SELECT MIN(tenant_id) FROM tenants),'Esplanade','blue','41FF00','[[7.77626,48.574449],[7.781754,48.581434],[7.762613,48.586317],[7.758923,48.583989],[7.761154,48.575187],[7.77626,48.574449]]') ON CONFLICT DO NOTHING;
/****************************** ZONES *******************************/


/****************************** DEVICES *******************************/
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('eb228ed2043331ce',1,70,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('az228ed2063371cd',1,7,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('228e67a2063380ab',1,40,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('ba21d430e4a14f8c',1,98,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('f2a6fa67abea7ca1',1,58,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('afa5fa32a02a8da7',1,38,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('3ea75a54a03aaba1',1,28,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('2fadea8ea3da44ad',1,72,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('9aa56ae5af6aa3a5',1,98,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('9ba32ad3ab7a16a8',1,62,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('ffa0fa95a87a5fa6',1,83,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('07ab7a16a14ad8a7',1,43,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('3ea42a72a18a24ac',1,45,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('1eabeac7ab5ac2a2',1,97,'free') ON CONFLICT DO NOTHING;
INSERT INTO devices (device_eui, tenant_id, battery, state) VALUES
('8ea25a35afda98ac',1,31,'free') ON CONFLICT DO NOTHING;
/****************************** DEVICES *******************************/



/****************************** PLACES *******************************/
INSERT INTO places (zone_id,type,geo,device_id) VALUES
(1,'car','[7.7562118148803165,48.58000286426977]',1) ON CONFLICT DO NOTHING;
INSERT INTO places (zone_id,type,geo,device_id) VALUES
(1,'car','[7.756855545044573,48.58042873737395]',2) ON CONFLICT DO NOTHING;
INSERT INTO places (zone_id,type,geo,device_id) VALUES
(1,'car','[7.7565122223070375,48.57600562405872]',3) ON CONFLICT DO NOTHING;

INSERT INTO places (zone_id,type,geo,device_id) VALUES
(2,'car','[7.748572883605334,48.58111012687789]',1) ON CONFLICT DO NOTHING;
INSERT INTO places (zone_id,type,geo,device_id) VALUES
(2,'car','[7.742736396789951,48.584431769173165]',2) ON CONFLICT DO NOTHING;
INSERT INTO places (zone_id,type,geo,device_id) VALUES
(2,'car','[7.755696830749002,48.58275678025768]',3) ON CONFLICT DO NOTHING;

INSERT INTO places (zone_id,type,geo,device_id) VALUES
(3,'car','[7.770037391200276,48.58193255586923]',1) ON CONFLICT DO NOTHING;
INSERT INTO places (zone_id,type,geo,device_id) VALUES
(3,'car','[7.779245939261955,48.580371954495746]',2) ON CONFLICT DO NOTHING;
INSERT INTO places (zone_id,type,geo,device_id) VALUES
(3,'car','[7.775411827451535,48.577329805544565]',3) ON CONFLICT DO NOTHING;
/****************************** PLACES *******************************/
