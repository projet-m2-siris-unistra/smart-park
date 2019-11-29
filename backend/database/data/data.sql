/****************************** TENANTS *******************************/
INSERT INTO tenants
(
    name, 
    geo
) VALUES
(
    'Schmilbligheim',
    '[7.7475,48.5827]'
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
    '[[7.739396,48.579816],[7.742014,48.579957],[7.744117,48.579134],[7.747464,48.578623],[7.74888,48.57885],[7.751756,48.579929],[7.755189,48.581831],[7.756906,48.583251],[7.754288,48.58555],[7.753558,48.586061],[7.751455,48.586743],[7.748537,48.58714],[7.746906,48.586828],[7.744503,48.585834],[7.740769,48.584244],[7.73901,48.582967],[7.738409,48.581973],[7.738495,48.580781],[7.739396,48.579816]]'
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
    '[7.746680,48.580402]',
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
    '[7.7475,48.5827]',
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
    '[7.742014,48.579957]',
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