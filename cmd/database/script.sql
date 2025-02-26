-- Roles

INSERT INTO roles (id, description) 
VALUES (gen_random_uuid(), 'Administrador');

INSERT INTO roles (id, description) 
VALUES (gen_random_uuid(), 'Usuario');

--Deparments

INSERT INTO public.departments(id, name)
VALUES
	(5,'Antioquia'),
	(8,'Atlantico'),
	(11,'Bogota'),
	(13,'Bolivar'),
	(15,'Boyaca'),
	(17,'Caldas'),
	(18,'Caqueta'),
	(19,'Cauca'),
	(20,'Cesar'),
	(23,'Cordoba'),
	(25,'Cundinamarca'),
	(27,'Choco'),
	(41,'Huila'),
	(44,'La Guajira'),
	(47,'Magdalena'),
	(50,'Meta'),
	(52,'Nariño'),
	(54,'N. De Santander'),
	(63,'Quindio'),
	(66,'Risaralda'),
	(68,'Santander'),
	(70,'Sucre'),
	(73,'Tolima'),
	(76,'Valle Del Cauca'),
	(81,'Arauca'),
	(85,'Casanare'),
	(86,'Putumayo'),
	(88,'San Andres'),
	(91,'Amazonas'),
	(94,'Guainia'),
	(95,'Guaviare'),
	(97,'Vaupes'),
	(99,'Vichada');
