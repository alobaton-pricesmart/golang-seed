-- clients
INSERT INTO  clients (code, secret, `domain`, user_id, created_at, updated_at)
VALUES('client-id', 'kP++3#8Q', NULL, NULL,  NOW(), NULL);
-- users
ALTER TABLE users ADD UNIQUE (email);

INSERT INTO users (created_at, updated_at, id, nickname, name, last_name, email, password, locked, enabled)
VALUES( NOW(), NULL, 'f7528acd-b825-4147-a258-53de545f782d', 'admin', 'Administrator', NULL, 'admin@domain.com', '$2a$10$yvACdxM6acaYts.lyK6NBOyI5woXICQ9IMkmjZHlAtmj3Pm3nQzsy', 0, 1);

-- roles
INSERT INTO roles (created_at, updated_at, code, group_id, name, description)
VALUES(NOW(), NULL, 'ADMIN', NULL, 'Administrator', NULL);

-- user_role
INSERT INTO user_role (created_at, updated_at, user_id, role_code)
VALUES(NOW(), NULL, 'f7528acd-b825-4147-a258-53de545f782d', 'ADMIN');

-- permissions
INSERT INTO permissions ( code, name , created_at)
VALUES ( 'read:user', NULL, NOW());

INSERT INTO permissions( code, name, created_at)
VALUES( 'read:users', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'create:user', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'update:user', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'delete:user', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'read:role', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'read:roles', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'create:role', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'update:role', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'delete:role', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'read:permission', NULL, NOW());

INSERT INTO permissions( code, name , created_at)
VALUES( 'read:permissions', NULL, NOW());

INSERT INTO permissions ( code, name , created_at)
VALUES ( 'create:permission', NULL, NOW());

INSERT INTO permissions ( code, name , created_at)
VALUES ( 'update:permission', NULL, NOW());

INSERT INTO permissions ( code, name , created_at)
VALUES ( 'delete:permission', NULL, NOW());

INSERT INTO permissions ( code, name , created_at)
VALUES ( 'read:client', NULL, NOW());

INSERT INTO permissions ( code, name , created_at)
VALUES ( 'read:clients', NULL, NOW());

INSERT INTO permissions ( code, name , created_at)
VALUES ( 'create:client', NULL, NOW());

INSERT INTO permissions ( code, name , created_at)
VALUES ( 'update:client', NULL, NOW());

INSERT INTO permissions ( code, name , created_at)
VALUES ( 'delete:client', NULL, NOW());

-- role_permission
INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'read:user', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'read:users', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'create:user', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'update:user', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'delete:user', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'read:role', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'read:roles', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'create:role', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'update:role', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'delete:role', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'read:permission', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'read:permissions', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'create:permission', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'update:permission', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'delete:permission', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'read:client', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'read:clients', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'create:client', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'update:client', 'ADMIN', NOW() );

INSERT INTO role_permission ( permission_code, role_code, created_at)
VALUES ( 'delete:client', 'ADMIN', NOW() );



