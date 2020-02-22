-- MySQL dump 10.13  Distrib 8.0.19, for Linux (x86_64)
--
-- Host: localhost    Database: gameaholic
-- ------------------------------------------------------
-- Server version	8.0.19

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
-- Table structure for table `achievements`
--

DROP TABLE IF EXISTS `achievements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `achievements` (
  `aid` int NOT NULL AUTO_INCREMENT,
  `aname` varchar(255) NOT NULL,
  `alevel` int unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `id` varchar(255) NOT NULL,
  `aavatar` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`aid`),
  KEY `alevel` (`alevel`),
  CONSTRAINT `achievements_ibfk_1` FOREIGN KEY (`alevel`) REFERENCES `levels` (`l_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `achievements`
--

LOCK TABLES `achievements` WRITE;
/*!40000 ALTER TABLE `achievements` DISABLE KEYS */;
INSERT INTO `achievements` VALUES (1,'Cadet',1,'2020-02-06 13:29:54','2020-02-19 14:23:08','51227445-08ca-47c1-bc04-aa56fb074b7f','https://image.flaticon.com/icons/svg/1492/1492443.svg'),(2,'Private',4,'2020-02-08 06:31:37','2020-02-19 14:23:43','3ee068ba-f1fd-46b9-ba5e-421adb72e7a3','https://image.flaticon.com/icons/svg/1492/1492444.svg'),(3,'Corporal',7,'2020-02-19 14:24:14','2020-02-19 14:24:14','66451754-ce13-4d77-ac62-28ef2d7ae1f0','https://image.flaticon.com/icons/svg/1492/1492453.svg'),(4,'Sergeant',10,'2020-02-19 14:24:45','2020-02-19 14:24:45','02e18079-5b19-4b83-b6a6-3c8bb63a9b32','https://image.flaticon.com/icons/svg/1492/1492460.svg'),(5,'Captain',13,'2020-02-19 14:25:14','2020-02-19 14:25:14','73d50687-adcd-4ea9-ab97-71559cac26bf','https://image.flaticon.com/icons/svg/1492/1492436.svg'),(6,'Major',16,'2020-02-19 14:25:39','2020-02-19 14:25:39','bc77cecd-ddf2-41a2-9d8d-7ee22e426fd6','https://image.flaticon.com/icons/svg/1492/1492433.svg');
/*!40000 ALTER TABLE `achievements` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `levels`
--

DROP TABLE IF EXISTS `levels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `levels` (
  `l_id` int unsigned NOT NULL,
  `level` int NOT NULL,
  `score` int NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `id` varchar(255) NOT NULL,
  `aavatar` varchar(255) DEFAULT NULL,
  `level_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`l_id`),
  KEY `l_id` (`l_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `levels`
--

LOCK TABLES `levels` WRITE;
/*!40000 ALTER TABLE `levels` DISABLE KEYS */;
INSERT INTO `levels` VALUES (1,1,1,'2020-02-06 13:29:36','2020-02-19 14:14:34','fe5c3283-e488-4ff4-aef7-35981f48ec2a','https://image.flaticon.com/icons/svg/340/340333.svg','Noob'),(2,2,10,'2020-02-08 06:25:54','2020-02-19 14:15:08','6e3261ed-b21d-4123-b6d2-ea54d3846d17','https://image.flaticon.com/icons/png/512/474/474230.png','Dreamy Eyes!'),(3,3,20,'2020-02-08 06:26:24','2020-02-19 14:16:38','a9a6a940-fc28-4678-9a93-8e3cdb8d2a64','https://image.flaticon.com/icons/png/512/474/474206.png','Good Job!'),(4,4,30,'2020-02-08 06:27:48','2020-02-19 14:16:59','16814858-281b-4eb9-a8e8-f46aee1a6264','https://image.flaticon.com/icons/png/512/474/474233.png','Great Expectations!'),(5,5,40,'2020-02-19 14:17:36','2020-02-19 14:17:36','c44816fa-501d-4f87-864c-a6983de821ac','https://image.flaticon.com/icons/svg/474/474232.svg','Power Packed!'),(6,6,50,'2020-02-19 14:18:03','2020-02-19 14:18:03','8f1a18b9-0174-45da-a6c5-5bedae85771a','https://image.flaticon.com/icons/svg/340/340334.svg','Sweet Treat!'),(7,7,60,'2020-02-19 14:18:23','2020-02-19 14:18:23','7dcda766-53f4-43e3-acc0-af64df02f299','https://image.flaticon.com/icons/png/512/474/474207.png','Unstoppable!'),(8,8,70,'2020-02-19 14:18:42','2020-02-19 14:18:42','c84d9cfc-e74e-427b-80aa-30fce8f515cc','https://image.flaticon.com/icons/png/512/474/474214.png','Shooting High!'),(9,9,80,'2020-02-19 14:19:05','2020-02-19 14:19:05','1536684d-74ae-4acb-8ec3-c0c3f4cbb1a5','https://image.flaticon.com/icons/svg/340/340336.svg','Arrived!'),(10,10,90,'2020-02-19 14:19:26','2020-02-19 14:19:26','10a9dc6b-6fcb-4834-900d-5d61fe2d32df','https://image.flaticon.com/icons/svg/340/340337.svg','Sharp Shooter!'),(11,11,100,'2020-02-19 14:19:53','2020-02-19 14:19:53','cbbdba56-4b49-49e3-ba8c-29c1e96faea6','https://image.flaticon.com/icons/svg/340/340338.svg','Boss!'),(12,12,110,'2020-02-19 14:20:16','2020-02-19 14:20:16','08494ad7-7e74-430a-aba6-4ad2ebbb89b1','https://image.flaticon.com/icons/svg/340/340328.svg','Onwards and Upwards!'),(13,13,120,'2020-02-19 14:20:40','2020-02-19 14:20:40','123b9486-1b3a-428b-98ec-0edfd1c84813','https://image.flaticon.com/icons/svg/340/340331.svg','Aye Aye Captain!'),(14,14,130,'2020-02-19 14:21:03','2020-02-19 14:21:03','7042312c-b676-4017-b2d8-fa2edc9129be','https://image.flaticon.com/icons/png/512/474/474211.png','Trusted Aide!'),(15,15,140,'2020-02-19 14:21:25','2020-02-19 14:21:25','4d54839d-3f6a-419c-95d7-916011cacd0d','https://image.flaticon.com/icons/svg/340/340340.svg','Speed Alert!'),(16,16,150,'2020-02-19 14:22:06','2020-02-19 14:22:06','6684d59f-bbed-4930-856d-1ae5698a39b3','https://image.flaticon.com/icons/svg/340/340339.svg','Juggernaut!');
/*!40000 ALTER TABLE `levels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `personal_achievements`
--

DROP TABLE IF EXISTS `personal_achievements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `personal_achievements` (
  `pa_id` int NOT NULL AUTO_INCREMENT,
  `player_id` int unsigned NOT NULL,
  `achiev_id` int NOT NULL,
  `achiev_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `id` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`pa_id`),
  KEY `idx_player_id` (`player_id`),
  KEY `achiev_id` (`achiev_id`),
  CONSTRAINT `personal_achievements_ibfk_1` FOREIGN KEY (`player_id`) REFERENCES `players` (`pid`),
  CONSTRAINT `personal_achievements_ibfk_2` FOREIGN KEY (`achiev_id`) REFERENCES `achievements` (`aid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `personal_achievements`
--

LOCK TABLES `personal_achievements` WRITE;
/*!40000 ALTER TABLE `personal_achievements` DISABLE KEYS */;
INSERT INTO `personal_achievements` VALUES (1,1,1,'2020-02-08 00:00:00','5f90859b-6890-4e05-b9e2-9f2901b4e27c','2020-02-08 21:04:54','2020-02-08 21:04:54');
/*!40000 ALTER TABLE `personal_achievements` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `players`
--

DROP TABLE IF EXISTS `players`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `players` (
  `pid` int unsigned NOT NULL AUTO_INCREMENT,
  `pname` varchar(255) NOT NULL,
  `pavatar` varchar(255) DEFAULT NULL,
  `pemailid` varchar(255) NOT NULL,
  `pcur_achive` varchar(255) NOT NULL,
  `pscore` int NOT NULL,
  `plevel` int unsigned NOT NULL,
  `bitbucket_uname` varchar(255) NOT NULL,
  `jira_uname` varchar(255) NOT NULL,
  `confluence_uname` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`pid`),
  KEY `plevel` (`plevel`),
  CONSTRAINT `players_ibfk_1` FOREIGN KEY (`plevel`) REFERENCES `levels` (`l_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `players`
--

LOCK TABLES `players` WRITE;
/*!40000 ALTER TABLE `players` DISABLE KEYS */;
INSERT INTO `players` VALUES (1,'Narayanan Krishnamurthy','https://image.flaticon.com/icons/svg/1492/1492437.svg','narayanan.kmu@gmail.com','Relentless',28,3,'krishnna','narayanan.krishnamurthy','narayanan.krishnamurthy','2020-02-06 13:53:37','2020-02-09 00:35:34','b59ca5aa-de08-4750-b52b-05899b416e44');
/*!40000 ALTER TABLE `players` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-02-19 14:28:16