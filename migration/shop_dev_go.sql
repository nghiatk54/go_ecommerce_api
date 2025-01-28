mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 8.0.40, for Linux (aarch64)
--
-- Host: localhost    Database: shopDevGo
-- ------------------------------------------------------
-- Server version	8.0.40

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `go_crm_user`
--

DROP TABLE IF EXISTS `go_crm_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `go_crm_user` (
  `usr_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account Id',
  `usr_email` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
  `usr_phone` varchar(15) NOT NULL DEFAULT '' COMMENT 'Phone number',
  `usr_username` varchar(30) NOT NULL DEFAULT '' COMMENT 'User name',
  `usr_password` varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
  `usr_create_at` int NOT NULL DEFAULT '0' COMMENT 'Creation time',
  `usr_update_at` int NOT NULL DEFAULT '0' COMMENT 'Update time',
  `usr_create_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Creation IP',
  `usr_last_login_at` int NOT NULL DEFAULT '0' COMMENT 'Last login time',
  `usr_last_login_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Last login IP',
  `usr_login_times` int NOT NULL DEFAULT '0' COMMENT 'Login times',
  `usr_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Status: 1 - enabled, 0 - disabled, -1 - deleted',
  PRIMARY KEY (`usr_id`),
  KEY `idx_email` (`usr_email`),
  KEY `idx_phone` (`usr_phone`),
  KEY `idx_username` (`usr_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Account';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `go_crm_user`
--

LOCK TABLES `go_crm_user` WRITE;
/*!40000 ALTER TABLE `go_crm_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `go_crm_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `go_crm_user_v2`
--

DROP TABLE IF EXISTS `go_crm_user_v2`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `go_crm_user_v2` (
  `usr_id` bigint NOT NULL AUTO_INCREMENT COMMENT 'Account Id',
  `usr_email` varchar(255) NOT NULL COMMENT 'Email',
  `usr_phone` longtext NOT NULL COMMENT 'Phone number',
  `usr_username` longtext NOT NULL COMMENT 'User name',
  `usr_password` longtext NOT NULL COMMENT 'Password',
  `usr_create_at` bigint NOT NULL COMMENT 'Creation time',
  `usr_update_at` bigint NOT NULL COMMENT 'Update time',
  `usr_create_ip_at` longtext NOT NULL COMMENT 'Creation IP',
  `usr_last_login_at` bigint NOT NULL COMMENT 'Last login time',
  `usr_last_login_ip_at` longtext NOT NULL COMMENT 'Last login IP',
  `usr_login_times` bigint NOT NULL COMMENT 'Login times',
  `usr_status` tinyint(1) NOT NULL COMMENT 'Status: 1 - enabled, 0 - disabled, -1 - deleted',
  PRIMARY KEY (`usr_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `go_crm_user_v2`
--

LOCK TABLES `go_crm_user_v2` WRITE;
/*!40000 ALTER TABLE `go_crm_user_v2` DISABLE KEYS */;
/*!40000 ALTER TABLE `go_crm_user_v2` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `go_db_role`
--

DROP TABLE IF EXISTS `go_db_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `go_db_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '''Primary key is Id''',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_name` longtext,
  `role_note` text,
  PRIMARY KEY (`id`),
  KEY `idx_go_db_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `go_db_role`
--

LOCK TABLES `go_db_role` WRITE;
/*!40000 ALTER TABLE `go_db_role` DISABLE KEYS */;
/*!40000 ALTER TABLE `go_db_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `go_db_user`
--

DROP TABLE IF EXISTS `go_db_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `go_db_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(255) NOT NULL,
  `user_name` longtext,
  `is_active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_go_db_user_uuid` (`uuid`),
  KEY `idx_go_db_user_deleted_at` (`deleted_at`),
  KEY `idx_uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `go_db_user`
--

LOCK TABLES `go_db_user` WRITE;
/*!40000 ALTER TABLE `go_db_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `go_db_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `go_user_role`
--

DROP TABLE IF EXISTS `go_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `go_user_role` (
  `user_id` bigint unsigned NOT NULL,
  `role_id` bigint NOT NULL COMMENT '''Primary key is Id''',
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `fk_go_user_role_role` (`role_id`),
  CONSTRAINT `fk_go_user_role_role` FOREIGN KEY (`role_id`) REFERENCES `go_db_role` (`id`),
  CONSTRAINT `fk_go_user_role_user` FOREIGN KEY (`user_id`) REFERENCES `go_db_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `go_user_role`
--

LOCK TABLES `go_user_role` WRITE;
/*!40000 ALTER TABLE `go_user_role` DISABLE KEYS */;
/*!40000 ALTER TABLE `go_user_role` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-01-28 10:06:19
