

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identifier` varchar(100) DEFAULT NULL,
  `created` date DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  `country` varchar(100) DEFAULT NULL,
  `type_identifier` varchar(100) DEFAULT NULL,
  `full_name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `class` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name_class` varchar(100) DEFAULT NULL,
  `status` bit(1) DEFAULT NULL,
  `id_program` varchar(100) DEFAULT NULL,
  `id_teacher` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `class_FK` (`id_teacher`),
  CONSTRAINT `class_FK` FOREIGN KEY (`id_teacher`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `student_class` (
  `id` int NOT NULL AUTO_INCREMENT,
  `year` int DEFAULT NULL,
  `semester` int DEFAULT NULL,
  `id_student` int DEFAULT NULL,
  `id_class` int DEFAULT NULL,
  `nota` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `student_class_FK` (`id_student`),
  KEY `student_class_FK_1` (`id_class`),
  CONSTRAINT `student_class_FK` FOREIGN KEY (`id_student`) REFERENCES `users` (`id`),
  CONSTRAINT `student_class_FK_1` FOREIGN KEY (`id_class`) REFERENCES `class` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;