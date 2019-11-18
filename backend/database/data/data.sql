/****************************** TENANTS *******************************/
INSERT INTO tenants
(
    name, 
    geo
) VALUES
(
    'Schmilbligheim',
    '7.7475 ; 48.5827'
) ON CONFLICT DO NOTHING;

/****************************** TENANTS *******************************/


/****************************** ZONES *******************************/
INSERT INTO zones
(
    tenant_id, 
    name,
    type,
    color,
    geo
) VALUES
(
    (SELECT MIN(tenant_id) FROM tenants),
    'centre',
    'paid',
    'EB3434',
    '7.7475 ; 48.5827'
) ON CONFLICT DO NOTHING;
/****************************** ZONES *******************************/


/****************************** DEVICES *******************************/
INSERT INTO devices
(
    battery, 
    state
) VALUES
(
    70,
    DEFAULT
) ON CONFLICT DO NOTHING;

INSERT INTO devices
(
    battery, 
    state
) VALUES
(
    7,
    'occupied'
) ON CONFLICT DO NOTHING;

INSERT INTO devices
(
    battery, 
    state
) VALUES
(
    40,
    'free'
) ON CONFLICT DO NOTHING;
/****************************** DEVICES *******************************/



/****************************** PLACES *******************************/
INSERT INTO places
(
    zone_id, 
    type,
    geo,
    device_id
) VALUES
(
    (SELECT MIN(zone_id) FROM zones),
    'car',
    '7.739396 ; 48.579816',
    (SELECT MIN(device_id) FROM devices)
) ON CONFLICT DO NOTHING;

INSERT INTO places
(
    zone_id, 
    type,
    geo,
    device_id
) VALUES
(
    (SELECT MIN(zone_id) FROM zones),
    'car',
    '7.744117 ; 48.579134',
    (SELECT MIN(device_id) FROM devices)
) ON CONFLICT DO NOTHING;

INSERT INTO places
(
    zone_id, 
    type,
    geo,
    device_id
) VALUES
(
    (SELECT MIN(zone_id) FROM zones),
    'car',
    '7.742014 ; 48.579957',
    (SELECT MIN(device_id) FROM devices)
) ON CONFLICT DO NOTHING;
/****************************** PLACES *******************************/



/****************************** USERS *******************************/
INSERT INTO users
(
    tenant_id,
    username,
    password,
    email
) VALUES
(
    (SELECT MIN(tenant_id) FROM tenants),
    'constantin',
    'sydney',
    'divriotis.constantin@gmail.com'
) ON CONFLICT DO NOTHING;

INSERT INTO users
(
    tenant_id,
    username,
    password,
    email
) VALUES
(
    (SELECT MIN(tenant_id) FROM tenants),
    'quentin',
    'oslo',
    'gliech.quentin@gmail.com'
) ON CONFLICT DO NOTHING;

INSERT INTO users
(
    tenant_id,
    username,
    password,
    email
) VALUES
(
    (SELECT MIN(tenant_id) FROM tenants),
    'lionel',
    'moscou',
    'jung.lionel@gmail.com'
) ON CONFLICT DO NOTHING;
/****************************** USERS *******************************/