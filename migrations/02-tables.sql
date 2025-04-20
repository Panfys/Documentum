SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `doc`;
CREATE TABLE `doc` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `fnum` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `fdate` date DEFAULT NULL,
  `lnum` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `ldate` date DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `sender` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `ispolnitel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `familiar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `count` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `copy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `width` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `file` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `funcs`;
CREATE TABLE `funcs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `fullname_f` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `funcs` (`id`, `name`, `fullname_f`) VALUES
(1,	'Начальник Центра',	'Начальник Центра'),
(2,	'Заместитель НЦ (по ОРДС)',	'Заместитель начальника Центра <br> (по организации работы <br> дежурных смен)'),
(3,	'Заместитель НЦ - НН',	'Заместитель начальника Центра - начальник направления'),
(4,	'Начальник отдела',	'Начальник отдела'),
(5,	'Помощник НЦ (по тылу)',	'Помощник начальника Центра (по тылу)'),
(6,	'Заместитель начальника направления',	'Заместитель начальника направления'),
(7,	'Начальник дежурной смены',	'Начальник дежурной смены'),
(8,	'Начальник группы',	'Начальник группы'),
(9,	'Начальник службы - ПНЦ по ЗГТ',	'Начальник службы - помощник начальника Центра по защите государственной тайны'),
(10,	'Старший офицер-оператор',	'Старший офицер-оператор'),
(11,	'Начальник отделения',	'Начальник отделения'),
(12,	'Офицер по ОБИ',	'Офицер (по обеспечению безопасности информации)');

DROP TABLE IF EXISTS `funcs_groups`;
CREATE TABLE `funcs_groups` (
  `funcs_id` int DEFAULT NULL,
  `units_id` int DEFAULT NULL,
  `groups_id` int DEFAULT NULL,
  KEY `funcs_id` (`funcs_id`),
  KEY `groups_id` (`groups_id`),
  KEY `units_id` (`units_id`),
  CONSTRAINT `funcs_groups_ibfk_1` FOREIGN KEY (`funcs_id`) REFERENCES `funcs` (`id`) ON DELETE SET NULL,
  CONSTRAINT `funcs_groups_ibfk_2` FOREIGN KEY (`groups_id`) REFERENCES `groups` (`id`) ON DELETE SET NULL,
  CONSTRAINT `funcs_groups_ibfk_3` FOREIGN KEY (`units_id`) REFERENCES `units` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `funcs_groups` (`funcs_id`, `units_id`, `groups_id`) VALUES
(8,	5,	2),
(8,	5,	3),
(8,	5,	4),
(8,	5,	5),
(8,	6,	6),
(8,	6,	7),
(8,	6,	8),
(8,	6,	9),
(8,	7,	10),
(8,	7,	11),
(8,	8,	12),
(8,	8,	13),
(8,	9,	14),
(8,	9,	15),
(10,	5,	2),
(10,	5,	3),
(10,	5,	4),
(10,	5,	5),
(10,	6,	6),
(10,	6,	7),
(10,	6,	8),
(10,	6,	9),
(10,	7,	10),
(10,	7,	11),
(10,	8,	12),
(10,	8,	13),
(10,	9,	14),
(10,	9,	15),
(11,	4,	16);

DROP TABLE IF EXISTS `funcs_units`;
CREATE TABLE `funcs_units` (
  `funcs_id` int DEFAULT NULL,
  `units_id` int DEFAULT NULL,
  KEY `funcs_id` (`funcs_id`),
  KEY `units_id` (`units_id`),
  CONSTRAINT `funcs_units_ibfk_1` FOREIGN KEY (`funcs_id`) REFERENCES `funcs` (`id`) ON DELETE SET NULL,
  CONSTRAINT `funcs_units_ibfk_2` FOREIGN KEY (`units_id`) REFERENCES `units` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `funcs_units` (`funcs_id`, `units_id`) VALUES
(1,	1),
(2,	2),
(3,	5),
(3,	6),
(4,	7),
(4,	8),
(4,	9),
(5,	1),
(6,	5),
(6,	6),
(7,	2),
(8,	3),
(8,	5),
(8,	6),
(8,	7),
(8,	8),
(8,	9),
(9,	4),
(10,	2),
(10,	3),
(10,	5),
(10,	6),
(10,	7),
(10,	8),
(10,	9),
(11,	4),
(12,	4);

DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `fullname_g` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `groups` (`id`, `name`, `fullname_g`) VALUES
(1,	'',	''),
(2,	'Группа (СА, С и Р КСБ ВС РФ)',	'Группа (системного анализа, строительства и развития комплексной системы безопасности ВС РФ)'),
(3,	'Группа (ОП КСБ ВС РФ)',	'Группа (организации применения комплексной системы безопасности ВС РФ)'),
(4,	'Группа (С, ОП и Р СТ)',	'Группа (создания, организации применения и развития специальной техники)'),
(5,	'Группа (ИСБ)',	'Группа (интеллектуальных систем безопасности)'),
(6,	'Группа (СО)',	'Группа (специального обеспечения)'),
(7,	'Группа (ФЗ)',	'Группа (физической защиты)'),
(8,	'Группа (ИБ)',	'Группа (информационной безопасности)'),
(9,	'Группа (ИТБ)',	'Группа (инженерно-технической безопасности)'),
(10,	'Группа (С и ОП КСБ)',	'Группа (создания и организации применения комплексной системы безопасности)'),
(11,	'Группа (П и П МКОБ)',	'Группа (подготовки и проведения мероприятий комплексного обеспечения безопасности)'),
(12,	'Группа (ТЗИ)',	'Группа (технической защиты информации)'),
(13,	'Группа (СП)',	'Группа (специального контроля)'),
(14,	'Группа (О В(С) СТ)',	'Группа (оснащения войск (сил) специальной техникой)'),
(15,	'Группа (О и К ЭСТ)',	'Группа (организации и контроля применения специальной техники)'),
(16,	'Секретное отделение',	'Секретное отделение');

DROP TABLE IF EXISTS `res`;
CREATE TABLE `res` (
  `id` int NOT NULL AUTO_INCREMENT,
  `doc_id` int,
  `ispolnitel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `text` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `time` date DEFAULT NULL,
  `date` date DEFAULT NULL,
  `user` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`),
  KEY `doc_id` (`doc_id`),
  CONSTRAINT `res_ibfk_1` FOREIGN KEY (`doc_id`) REFERENCES `doc` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `units`;
CREATE TABLE `units` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `fullname_u` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `units` (`id`, `name`, `fullname_u`) VALUES
(1,	'Командование',	'Командование Центра'),
(2,	'Дежурная смена',	'Дежурная смена (мониторинга и управления)'),
(3,	'Организационно-плановая группа',	'Организационно-плановая группа'),
(4,	'Служба ЗГТ',	'Служба защиты государственной тайны'),
(5,	'Направление (КБ ВС РФ)',	'Направление (комплексной безопасности ВС РФ)'),
(6,	'Направление (КБ 3 Дома МО РФ)',	'Направление (комплексной безопасности 3 Дома МО РФ)'),
(7,	'Отдел (КБ АЗМО и 1 Дома)',	'Отдел (комплексной безопасности Административного здания МО РФ и 1 Дома МО РФ)'),
(8,	'Отдел (КБ РС МО РФ)',	'Отдел (комплексной безопасности руководящего состава МО РФ)'),
(9,	'Отдел (ТО КСБ ВС РФ)',	'Отдел (технического обеспечению комплексной системы безопасности ВС РФ)');

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `login` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `func_id` int,
  `unit_id` int,
  `group_id` int,
  `pass` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`),
  KEY `func_id` (`func_id`),
  KEY `unit_id` (`unit_id`),
  KEY `users_ibfk_2` (`group_id`),
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`func_id`) REFERENCES `funcs` (`id`) ON DELETE SET NULL,
  CONSTRAINT `users_ibfk_2` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`),
  CONSTRAINT `users_ibfk_3` FOREIGN KEY (`unit_id`) REFERENCES `units` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `users` (`id`, `login`, `name`, `func_id`, `unit_id`, `group_id`, `pass`, `status`, `icon`) VALUES
(1,	'test',	'Иванов И.И.',	12,	4,	1, 	'$2y$10$3FVwCogIztiFc4FWUkhaAu7dvn05IumU6EoDgNyRY6yMjaJtOQVha',	'Администратор','');