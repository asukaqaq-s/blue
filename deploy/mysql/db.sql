-- MySQL dump 10.13  Distrib 5.7.26, for Win64 (x86_64)
--
-- Host: localhost    Database: fim_server_db
-- ------------------------------------------------------
-- Server version	5.7.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `chat_models`
--

DROP TABLE IF EXISTS `chat_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chat_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `send_user_id` bigint(20) unsigned DEFAULT NULL,
  `rev_user_id` bigint(20) unsigned DEFAULT NULL,
  `msg_type` tinyint(4) DEFAULT NULL,
  `msg_preview` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `msg` longtext COLLATE utf8mb4_unicode_ci,
  `system_msg` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chat_models`
--

LOCK TABLES `chat_models` WRITE;
/*!40000 ALTER TABLE `chat_models` DISABLE KEYS */;
INSERT INTO `chat_models` VALUES (1,'2024-03-10 13:43:03.000','2024-03-10 13:43:04.000',1,2,NULL,'你好',NULL,NULL),(2,'2024-03-10 13:46:56.000','2024-03-10 13:46:56.000',2,1,NULL,'什么事情',NULL,NULL),(3,'2024-03-11 00:15:12.000','2024-03-11 00:15:13.000',1,2,NULL,'哦',NULL,NULL),(7,'2024-03-12 21:52:09.098','2024-03-12 21:52:09.098',3,1,1,'我们已经是好友了，开始聊天吧！','{\"type\":1,\"textMsg\":{\"content\":\"我们已经是好友了，开始聊天吧！\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null}',NULL),(8,'2024-03-12 21:52:53.000','2024-03-12 21:52:54.000',1,3,1,'你好',NULL,NULL),(9,'2024-03-14 21:52:53.000','2024-03-14 21:52:54.000',1,1,1,'哈哈哈',NULL,NULL),(10,'2024-03-14 22:56:50.216','2024-03-14 22:56:50.216',3,1,2,'[图片消息] - xxx.jpg','{\"type\":2,\"textMsg\":null,\"imageMsg\":{\"title\":\"xxx.jpg\",\"src\":\"/uploads/xx.jpg\"},\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(11,'2024-03-17 23:36:16.630','2024-03-17 23:36:16.630',1,3,4,'[文件消息] - ','{\"type\":4,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":{\"title\":\"\",\"src\":\"/api/file/fddee72a-3976-4dc5-98e2-5c1f7ea57765\",\"size\":0,\"type\":\"\"},\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(12,'2024-03-17 23:37:35.823','2024-03-17 23:37:35.823',1,3,4,'[文件消息] - c3.jpg','{\"type\":4,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":{\"title\":\"c3.jpg\",\"src\":\"/api/file/fddee72a-3976-4dc5-98e2-5c1f7ea57765\",\"size\":244201,\"type\":\"jpg\"},\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(13,'2024-03-18 23:23:42.028','2024-03-18 23:23:42.028',3,1,2,'[图片消息] - xxx.jpg','{\"type\":2,\"textMsg\":null,\"imageMsg\":{\"title\":\"xxx.jpg\",\"src\":\"/uploads/xx.jpg\"},\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(14,'2024-03-18 23:27:57.805','2024-03-18 23:27:57.805',3,1,2,'[图片消息] - xxx.jpg','{\"type\":2,\"textMsg\":null,\"imageMsg\":{\"title\":\"xxx.jpg\",\"src\":\"/uploads/xx.jpg\"},\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(15,'2024-03-18 23:54:38.522','2024-03-18 23:55:03.042',3,1,1,'你好啊','{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"content\":\"xx 撤回了一条消息\",\"msgID\":15},\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(16,'2024-03-19 00:04:22.498','2024-03-19 00:04:22.498',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(17,'2024-03-19 00:04:30.785','2024-03-19 00:04:30.785',3,3,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(18,'2024-03-19 00:04:39.297','2024-03-19 00:04:39.297',3,3,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(19,'2024-03-19 00:04:43.482','2024-03-19 00:04:43.482',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(20,'2024-03-19 00:06:13.652','2024-03-19 00:06:13.652',3,3,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(21,'2024-03-19 00:06:22.037','2024-03-19 00:06:22.037',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(22,'2024-03-19 00:11:04.669','2024-03-19 00:11:04.669',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(23,'2024-03-19 00:11:11.924','2024-03-19 00:11:11.924',3,3,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(24,'2024-03-19 00:12:48.680','2024-03-19 00:12:48.680',3,3,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(25,'2024-03-19 00:13:02.226','2024-03-19 00:13:02.226',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(26,'2024-03-19 00:19:51.835','2024-03-19 00:19:51.835',3,3,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(27,'2024-03-19 00:19:58.565','2024-03-19 00:19:58.565',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(28,'2024-03-19 00:42:36.790','2024-03-19 00:42:36.790',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(29,'2024-03-19 00:43:05.566','2024-03-19 00:43:05.566',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(30,'2024-03-19 00:43:18.591','2024-03-19 00:43:18.591',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(31,'2024-03-19 00:44:30.585','2024-03-19 00:44:30.585',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(32,'2024-03-19 00:44:36.751','2024-03-19 00:44:36.751',3,3,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(33,'2024-03-19 00:44:45.241','2024-03-19 00:44:45.241',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(34,'2024-03-19 00:44:54.979','2024-03-19 00:44:54.979',1,3,4,'[文件消息] - c3.jpg','{\"type\":4,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":{\"title\":\"c3.jpg\",\"src\":\"/api/file/fddee72a-3976-4dc5-98e2-5c1f7ea57765\",\"size\":244201,\"type\":\"jpg\"},\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(35,'2024-03-19 19:03:29.821','2024-03-19 19:03:29.821',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(36,'2024-03-19 19:06:32.340','2024-03-19 19:06:32.340',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(37,'2024-03-19 19:06:42.138','2024-03-19 19:06:42.138',3,3,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(38,'2024-03-19 19:06:50.812','2024-03-19 19:06:50.812',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(39,'2024-03-19 19:10:51.115','2024-03-19 19:10:51.115',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(40,'2024-03-19 19:11:44.046','2024-03-19 19:11:44.046',3,1,1,'xxx','{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(41,'2024-03-19 19:11:55.206','2024-03-19 19:11:55.206',1,3,4,'[文件消息] - c3.jpg','{\"type\":4,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":{\"title\":\"c3.jpg\",\"src\":\"/api/file/fddee72a-3976-4dc5-98e2-5c1f7ea57765\",\"size\":244201,\"type\":\"jpg\"},\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(42,'2024-03-19 21:20:21.947','2024-03-19 21:20:21.947',3,1,9,'[回复消息] - 这是回复','{\"type\":9,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":{\"msgID\":40,\"content\":\"这是回复\",\"msg\":{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null},\"userID\":3,\"userNickName\":\"zhangsan\",\"originMsgDate\":\"2024-03-19T19:11:44.046+08:00\"},\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(43,'2024-03-19 21:20:45.051','2024-03-19 21:20:45.051',3,1,9,'[回复消息] - 这是回复','{\"type\":9,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":{\"msgID\":34,\"content\":\"这是回复\",\"msg\":{\"type\":4,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":{\"title\":\"c3.jpg\",\"src\":\"/api/file/fddee72a-3976-4dc5-98e2-5c1f7ea57765\",\"size\":244201,\"type\":\"jpg\"},\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null},\"userID\":1,\"userNickName\":\"fengfeng\",\"originMsgDate\":\"2024-03-19T00:44:54.979+08:00\"},\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(44,'2024-03-19 21:24:36.442','2024-03-19 21:24:36.442',3,1,9,'[回复消息] - 这是回复','{\"type\":9,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":{\"msgID\":2,\"content\":\"这是回复\",\"msg\":{\"type\":0,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null},\"userID\":2,\"userNickName\":\"枫枫\",\"originMsgDate\":\"2024-03-10T13:46:56+08:00\"},\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(45,'2024-03-19 21:38:12.376','2024-03-19 21:38:12.376',3,1,44,'[未知消息]','{\"type\":44,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"content\":\"\",\"msgID\":1},\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(46,'2024-03-19 21:40:57.574','2024-03-19 21:41:06.733',3,1,9,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":46},\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(47,'2024-03-19 21:46:45.104','2024-03-19 21:46:45.104',3,1,9,'[回复消息] - 这是回复','{\"type\":9,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":{\"msgID\":40,\"content\":\"这是回复\",\"msg\":{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null},\"userID\":3,\"userNickName\":\"zhangsan\",\"originMsgDate\":\"2024-03-19T19:11:44.046+08:00\"},\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(48,'2024-03-19 21:53:07.915','2024-03-19 21:54:01.148',3,1,9,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":48},\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(49,'2024-03-19 21:57:20.866','2024-03-19 21:57:20.866',3,1,9,'[回复消息] - 这是回复','{\"type\":9,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":{\"msgID\":40,\"content\":\"这是回复\",\"msg\":{\"type\":1,\"textMsg\":{\"content\":\"xxx\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null},\"userID\":3,\"userNickName\":\"zhangsan\",\"originMsgDate\":\"2024-03-19T19:11:44.046+08:00\"},\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(50,'2024-03-19 22:00:55.271','2024-03-19 22:00:55.271',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(51,'2024-03-19 22:04:35.855','2024-03-19 22:04:42.078',3,1,1,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":51,\"originMsg\":{\"type\":1,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}},\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(52,'2024-03-19 22:08:44.213','2024-03-19 22:08:44.213',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(53,'2024-03-19 22:10:49.281','2024-03-19 22:10:55.368',3,1,1,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":53,\"originMsg\":{\"type\":1,\"textMsg\":{\"content\":\"嘻嘻嘻\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}},\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(54,'2024-03-19 22:12:12.346','2024-03-19 22:12:19.076',3,1,1,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":54,\"originMsg\":{\"type\":1,\"textMsg\":{\"content\":\"嘻嘻嘻\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}},\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(55,'2024-03-19 22:13:47.031','2024-03-19 22:13:52.109',3,1,1,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"textMsg\":null,\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":55,\"originMsg\":{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"},\"imageMsg\":null,\"videoMsg\":null,\"fileMsg\":null,\"voiceMsg\":null,\"voiceCallMsg\":null,\"videoCallMsg\":null,\"withdrawMsg\":null,\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}},\"replyMsg\":null,\"quoteMsg\":null,\"atMsg\":null,\"tipMsg\":null}',NULL),(56,'2024-03-19 22:14:58.919','2024-03-19 22:14:58.919',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(57,'2024-03-19 22:15:02.038','2024-03-19 22:16:49.658',3,1,1,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":57,\"originMsg\":{\"type\":8}}}',NULL),(58,'2024-03-19 22:23:48.630','2024-03-19 22:23:55.190',3,1,1,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":58,\"originMsg\":{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}}}',NULL),(59,'2024-03-19 22:29:58.715','2024-03-19 22:29:58.715',3,1,10,'[引用消息] - 这是引用回复','{\"type\":10,\"quoteMsg\":{\"msgID\":56,\"content\":\"这是引用回复\",\"msg\":{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}},\"userID\":3,\"userNickName\":\"zhangsan\",\"originMsgDate\":\"2024-03-19T22:14:58.919+08:00\"}}',NULL),(60,'2024-03-19 22:31:32.740','2024-03-19 22:31:32.740',3,1,10,'[引用消息] - 这是引用回复','{\"type\":10,\"quoteMsg\":{\"msgID\":58,\"content\":\"这是引用回复\",\"msg\":{\"type\":8,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":58}},\"userID\":3,\"userNickName\":\"zhangsan\",\"originMsgDate\":\"2024-03-19T22:23:48.63+08:00\"}}',NULL),(61,'2024-03-19 22:32:55.918','2024-03-19 22:33:07.782',3,1,8,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":61,\"originMsg\":{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}}}',NULL),(62,'2024-03-20 00:10:21.467','2024-03-20 00:10:21.467',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(63,'2024-03-20 00:11:17.781','2024-03-20 00:11:17.781',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(64,'2024-03-20 00:11:26.670','2024-03-20 00:11:26.670',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(65,'2024-03-20 00:11:46.170','2024-03-20 00:11:46.170',1,3,4,'[文件消息] - c3.jpg','{\"type\":4,\"fileMsg\":{\"title\":\"c3.jpg\",\"src\":\"/api/file/fddee72a-3976-4dc5-98e2-5c1f7ea57765\",\"size\":244201,\"type\":\"jpg\"}}',NULL),(66,'2024-03-20 21:21:35.967','2024-03-20 21:21:35.967',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(67,'2024-03-20 21:22:09.328','2024-03-20 21:22:09.328',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(68,'2024-03-20 21:22:27.664','2024-03-20 21:22:27.664',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(69,'2024-03-20 21:22:51.154','2024-03-20 21:22:51.154',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(70,'2024-03-20 21:37:14.002','2024-03-20 21:37:14.002',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(71,'2024-03-20 21:37:22.077','2024-03-20 21:37:22.077',3,1,9,'[回复消息] - 这是回复','{\"type\":9,\"replyMsg\":{\"msgID\":40,\"content\":\"这是回复\",\"msg\":{\"type\":1,\"textMsg\":{\"content\":\"xxx\"}},\"userID\":3,\"userNickName\":\"zhangsan\",\"originMsgDate\":\"2024-03-19T19:11:44.046+08:00\"}}',NULL),(72,'2024-03-20 21:38:32.812','2024-03-20 21:38:32.812',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(73,'2024-03-20 21:38:41.013','2024-03-20 21:38:41.013',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(74,'2024-03-20 21:40:20.687','2024-03-20 21:40:20.687',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(75,'2024-03-20 21:40:36.008','2024-03-20 21:40:36.008',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(76,'2024-03-20 21:55:42.634','2024-03-20 21:55:42.634',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(77,'2024-03-20 21:55:48.292','2024-03-20 21:55:48.292',3,1,9,'[回复消息] - 这是回复','{\"type\":9,\"replyMsg\":{\"msgID\":40,\"content\":\"这是回复\",\"msg\":{\"type\":1,\"textMsg\":{\"content\":\"xxx\"}},\"userID\":3,\"userNickName\":\"zhangsan\",\"originMsgDate\":\"2024-03-19T19:11:44.046+08:00\"}}',NULL),(78,'2024-03-20 21:56:23.939','2024-03-20 21:56:23.939',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(79,'2024-03-20 21:56:38.691','2024-03-20 21:56:38.691',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(80,'2024-03-20 21:56:42.444','2024-03-20 21:56:42.444',3,1,9,'[回复消息] - 这是回复','{\"type\":9,\"replyMsg\":{\"msgID\":40,\"content\":\"这是回复\",\"msg\":{\"type\":1,\"textMsg\":{\"content\":\"xxx\"}},\"userID\":3,\"userNickName\":\"zhangsan\",\"originMsgDate\":\"2024-03-19T19:11:44.046+08:00\"}}',NULL),(81,'2024-03-20 22:01:35.396','2024-03-20 22:01:35.396',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(82,'2024-03-20 22:01:40.748','2024-03-20 22:01:40.748',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(83,'2024-03-27 22:25:21.584','2024-03-27 22:25:21.584',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL),(84,'2024-03-27 22:25:31.744','2024-03-27 22:25:31.744',3,1,1,'这是文本','{\"type\":1,\"textMsg\":{\"content\":\"这是文本\"}}',NULL);
/*!40000 ALTER TABLE `chat_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `file_models`
--

DROP TABLE IF EXISTS `file_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `file_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `file_name` longtext COLLATE utf8mb4_unicode_ci,
  `size` bigint(20) DEFAULT NULL,
  `path` longtext COLLATE utf8mb4_unicode_ci,
  `hash` longtext COLLATE utf8mb4_unicode_ci,
  `uid` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `file_models`
--

LOCK TABLES `file_models` WRITE;
/*!40000 ALTER TABLE `file_models` DISABLE KEYS */;
INSERT INTO `file_models` VALUES (3,'2024-03-17 09:40:46.490','2024-03-17 09:40:46.490',1,'bg-05e3c1f.jpg',162882,'uploads/group_avatar/bg-05e3c1f.jpg','2f466794e52b79092bfc7ba663761a37','d46ce2be-dcd5-476c-ae90-3fa94e87307f');
/*!40000 ALTER TABLE `file_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `friend_models`
--

DROP TABLE IF EXISTS `friend_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `friend_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `send_user_id` bigint(20) unsigned DEFAULT NULL,
  `rev_user_id` bigint(20) unsigned DEFAULT NULL,
  `notice` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sen_user_notice` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `rev_user_notice` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_friend_models_send_user_model` (`send_user_id`),
  KEY `fk_friend_models_rev_user_model` (`rev_user_id`),
  CONSTRAINT `fk_friend_models_rev_user_model` FOREIGN KEY (`rev_user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_friend_models_send_user_model` FOREIGN KEY (`send_user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `friend_models`
--

LOCK TABLES `friend_models` WRITE;
/*!40000 ALTER TABLE `friend_models` DISABLE KEYS */;
INSERT INTO `friend_models` VALUES (2,'2024-03-06 23:21:17.000','2024-03-08 23:10:25.974',1,2,'张三','张四丰',NULL),(6,'2024-03-12 21:52:09.095','2024-03-12 21:52:09.095',3,1,NULL,'','');
/*!40000 ALTER TABLE `friend_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `friend_verify_models`
--

DROP TABLE IF EXISTS `friend_verify_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `friend_verify_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `send_user_id` bigint(20) unsigned DEFAULT NULL,
  `rev_user_id` bigint(20) unsigned DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `additional_messages` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `verification_question` longtext COLLATE utf8mb4_unicode_ci,
  `send_status` tinyint(4) DEFAULT NULL,
  `rev_status` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_friend_verify_models_send_user_model` (`send_user_id`),
  KEY `fk_friend_verify_models_rev_user_model` (`rev_user_id`),
  CONSTRAINT `fk_friend_verify_models_rev_user_model` FOREIGN KEY (`rev_user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_friend_verify_models_send_user_model` FOREIGN KEY (`send_user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `friend_verify_models`
--

LOCK TABLES `friend_verify_models` WRITE;
/*!40000 ALTER TABLE `friend_verify_models` DISABLE KEYS */;
INSERT INTO `friend_verify_models` VALUES (7,'2024-03-09 13:01:48.963','2024-03-12 21:52:09.101',3,1,0,'我是枫枫','{\"problem1\":\"你的童年是什么1？\",\"problem2\":\"你的童年是什么1？\",\"problem3\":null,\"answer1\":\"上学\",\"answer2\":\"上学\",\"answer3\":null}',0,1);
/*!40000 ALTER TABLE `friend_verify_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_member_models`
--

DROP TABLE IF EXISTS `group_member_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_member_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `group_id` bigint(20) unsigned DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `member_nickname` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `role` tinyint(4) DEFAULT NULL,
  `prohibition_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_group_models_member_list` (`group_id`),
  CONSTRAINT `fk_group_member_models_group_model` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`),
  CONSTRAINT `fk_group_models_member_list` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_member_models`
--

LOCK TABLES `group_member_models` WRITE;
/*!40000 ALTER TABLE `group_member_models` DISABLE KEYS */;
INSERT INTO `group_member_models` VALUES (4,'2024-03-20 23:34:18.995','2024-03-22 00:42:49.137',3,1,'',1,NULL),(5,'2024-03-21 21:34:18.995','2024-03-23 00:55:49.616',3,2,'',3,NULL),(8,'2024-03-22 22:15:01.569','2024-03-22 22:15:01.569',4,1,'',1,NULL);
/*!40000 ALTER TABLE `group_member_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_models`
--

DROP TABLE IF EXISTS `group_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` longtext COLLATE utf8mb4_unicode_ci,
  `abstract` longtext COLLATE utf8mb4_unicode_ci,
  `avatar` longtext COLLATE utf8mb4_unicode_ci,
  `creator` bigint(20) unsigned DEFAULT NULL,
  `is_search` tinyint(1) DEFAULT NULL,
  `verification` tinyint(4) DEFAULT NULL,
  `verification_question` longtext COLLATE utf8mb4_unicode_ci,
  `is_invite` tinyint(1) DEFAULT NULL,
  `is_temporary_session` tinyint(1) DEFAULT NULL,
  `is_prohibition` tinyint(1) DEFAULT NULL,
  `size` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_models`
--

LOCK TABLES `group_models` WRITE;
/*!40000 ALTER TABLE `group_models` DISABLE KEYS */;
INSERT INTO `group_models` VALUES (3,'2024-03-20 23:34:18.991','2024-03-20 23:34:18.991','我的群聊','本群创建于2024-03-20:  群主很懒,什么都没有留下','我',1,1,2,NULL,0,0,0,10),(4,'2024-03-22 22:15:01.567','2024-03-22 22:57:16.425','枫枫博客交流群','本群创建于2024-03-22:  群主很懒,什么都没有留下','/uploads/xxx',1,0,1,'{\"problem1\":\"你是谁\",\"problem2\":null,\"problem3\":null,\"answer1\":\"我是海绵宝宝\",\"answer2\":null,\"answer3\":null}',0,0,0,10);
/*!40000 ALTER TABLE `group_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_msg_models`
--

DROP TABLE IF EXISTS `group_msg_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_msg_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `group_id` bigint(20) unsigned DEFAULT NULL,
  `send_user_id` bigint(20) unsigned DEFAULT NULL,
  `msg_type` tinyint(4) DEFAULT NULL,
  `msg_preview` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `msg` longtext COLLATE utf8mb4_unicode_ci,
  `system_msg` longtext COLLATE utf8mb4_unicode_ci,
  `member_id` bigint(20) unsigned DEFAULT NULL,
  `group_member_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_group_msg_models_member_model` (`member_id`),
  KEY `fk_group_msg_models_group_member_model` (`group_member_id`),
  KEY `fk_group_models_group_msg_list` (`group_id`),
  CONSTRAINT `fk_group_models_group_msg_list` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`),
  CONSTRAINT `fk_group_msg_models_group_member_model` FOREIGN KEY (`group_member_id`) REFERENCES `group_member_models` (`id`),
  CONSTRAINT `fk_group_msg_models_group_model` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`),
  CONSTRAINT `fk_group_msg_models_member_model` FOREIGN KEY (`member_id`) REFERENCES `group_member_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_msg_models`
--

LOCK TABLES `group_msg_models` WRITE;
/*!40000 ALTER TABLE `group_msg_models` DISABLE KEYS */;
INSERT INTO `group_msg_models` VALUES (2,'2024-03-23 10:23:37.000','2024-03-23 10:23:38.000',3,1,NULL,'你会好好的',NULL,NULL,NULL,NULL),(3,'2024-03-23 11:33:13.000','2024-03-23 11:33:11.000',4,1,NULL,'xxx',NULL,NULL,NULL,NULL),(6,'2024-03-23 22:04:20.828','2024-03-23 22:04:20.828',3,1,1,'大家好','{\"type\":1,\"textMsg\":{\"content\":\"大家好\"}}',NULL,NULL,NULL),(7,'2024-03-23 22:05:08.433','2024-03-23 22:05:08.433',3,2,1,'大家好，我是枫枫','{\"type\":1,\"textMsg\":{\"content\":\"大家好，我是枫枫\"}}',NULL,NULL,NULL),(10,'2024-03-23 22:16:17.884','2024-03-23 22:16:17.884',3,1,1,'大家好','{\"type\":1,\"textMsg\":{\"content\":\"大家好\"}}',NULL,NULL,NULL),(11,'2024-03-23 22:16:46.500','2024-03-23 22:16:46.500',3,2,1,'大家好，我是枫枫','{\"type\":1,\"textMsg\":{\"content\":\"大家好，我是枫枫\"}}',NULL,NULL,NULL),(12,'2024-03-23 22:17:55.744','2024-03-23 22:17:55.744',3,1,1,'大家好','{\"type\":1,\"textMsg\":{\"content\":\"大家好\"}}',NULL,NULL,NULL),(13,'2024-03-23 22:18:00.910','2024-03-23 22:55:57.177',3,1,8,'[撤回消息] - 你撤回了一条消息','{\"type\":8,\"withdrawMsg\":{\"content\":\"你撤回了一条消息\",\"msgID\":13,\"originMsg\":{\"type\":1,\"textMsg\":{\"content\":\"大家好\"}}}}',NULL,NULL,NULL),(14,'2024-03-23 22:21:08.662','2024-03-23 22:21:08.662',3,1,1,'大家好','{\"type\":1,\"textMsg\":{\"content\":\"大家好\"}}',NULL,NULL,NULL),(15,'2024-03-23 22:21:17.892','2024-03-23 22:21:17.892',3,2,1,'大家好，我是枫枫','{\"type\":1,\"textMsg\":{\"content\":\"大家好，我是枫枫\"}}',NULL,NULL,NULL),(19,'2024-03-23 22:22:59.407','2024-03-23 22:22:59.407',3,1,1,'大家好','{\"type\":1,\"textMsg\":{\"content\":\"大家好\"}}',NULL,NULL,NULL);
/*!40000 ALTER TABLE `group_msg_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_user_msg_delete_models`
--

DROP TABLE IF EXISTS `group_user_msg_delete_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_user_msg_delete_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `msg_id` bigint(20) unsigned DEFAULT NULL,
  `group_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_user_msg_delete_models`
--

LOCK TABLES `group_user_msg_delete_models` WRITE;
/*!40000 ALTER TABLE `group_user_msg_delete_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_user_msg_delete_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_user_top_models`
--

DROP TABLE IF EXISTS `group_user_top_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_user_top_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `group_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_user_top_models`
--

LOCK TABLES `group_user_top_models` WRITE;
/*!40000 ALTER TABLE `group_user_top_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_user_top_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_verify_models`
--

DROP TABLE IF EXISTS `group_verify_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_verify_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `group_id` bigint(20) unsigned DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `additional_messages` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `verification_question` longtext COLLATE utf8mb4_unicode_ci,
  `type` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_group_verify_models_group_model` (`group_id`),
  CONSTRAINT `fk_group_verify_models_group_model` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_verify_models`
--

LOCK TABLES `group_verify_models` WRITE;
/*!40000 ALTER TABLE `group_verify_models` DISABLE KEYS */;
INSERT INTO `group_verify_models` VALUES (3,'2024-03-22 23:38:30.492','2024-03-22 23:38:30.492',4,3,1,'我想进群',NULL,1),(4,'2024-03-23 00:37:19.098','2024-03-23 00:37:19.098',4,3,0,'',NULL,2);
/*!40000 ALTER TABLE `group_verify_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_models`
--

DROP TABLE IF EXISTS `log_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `log_type` tinyint(4) DEFAULT NULL,
  `ip` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `addr` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `user_nickname` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_avatar` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `level` varchar(12) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `title` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `content` longtext COLLATE utf8mb4_unicode_ci,
  `service` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_read` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_models`
--

LOCK TABLES `log_models` WRITE;
/*!40000 ALTER TABLE `log_models` DISABLE KEYS */;
INSERT INTO `log_models` VALUES (46,'2024-03-26 19:26:16.361','2024-03-26 19:57:23.794',3,'192.168.0.108','内网地址',0,'','','info','','<div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div>\n<div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">你好123</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div><div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div>\n<div class=\"log_item info\"><div class=\"log_item_label\">xxx</div> <div class=\"log_item_content\">注销了</div></div>','auth',NULL),(47,'2024-03-26 19:26:16.362','2024-03-26 19:26:16.362',2,'192.168.0.108','内网地址',2,'枫枫','http://thirdqq.qlogo.cn/ek_qqapp/AQVLjicwG14hrSLsMWRFF5BEHktJiaY39GqqYDoG350dwn8Mtz7ep8fiakFnw2QOKj07soBebmw/0','info','用户登录成功','<div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0MzQ2MzkwfQ.sO5n4Ded6zoEd7L1AW0mAvSzIRiQujqRJWeexNfzUFs\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEwNTY3fQ.u2s6ubAEw2mcBODYTT1vT6W-SfamPpqSwxp2sV_2u8U\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEwNzI0fQ.mURcs8qnNtqUL6k4CFoyO_voJHrPv6c2kimGxhl85Nc\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEwNzYzfQ.oyILwzSEpBAEXmJrwktjJq1XHcfraPlaGHiRSifRaUI\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEwNzY1fQ.RuP5ElDSPIyzEzWv1nJ9BLMDw4lkYunfXA7iz4RkvrU\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEyMjE5fQ.GI14akLZIAPKLuxmwdayDKL2dvbs9GhdYSZ3QC-aBBY\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEyMjIwfQ.up6PR2lOe2bEyiN54M674Q_6ujP8xLk93kxD9yUuJOA\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div>','auth',NULL),(48,'2024-03-26 19:26:16.363','2024-03-26 19:26:16.363',2,'192.168.0.108','内网地址',2,'枫枫','http://thirdqq.qlogo.cn/ek_qqapp/AQVLjicwG14hrSLsMWRFF5BEHktJiaY39GqqYDoG350dwn8Mtz7ep8fiakFnw2QOKj07soBebmw/0','info','用户登录成功','<div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0MzQ2MzkwfQ.sO5n4Ded6zoEd7L1AW0mAvSzIRiQujqRJWeexNfzUFs\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEwNTY3fQ.u2s6ubAEw2mcBODYTT1vT6W-SfamPpqSwxp2sV_2u8U\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEwNzI0fQ.mURcs8qnNtqUL6k4CFoyO_voJHrPv6c2kimGxhl85Nc\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEwNzYzfQ.oyILwzSEpBAEXmJrwktjJq1XHcfraPlaGHiRSifRaUI\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEwNzY1fQ.RuP5ElDSPIyzEzWv1nJ9BLMDw4lkYunfXA7iz4RkvrU\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEyMjE5fQ.GI14akLZIAPKLuxmwdayDKL2dvbs9GhdYSZ3QC-aBBY\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEyMjIwfQ.up6PR2lOe2bEyiN54M674Q_6ujP8xLk93kxD9yUuJOA\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEyMzU4fQ.e-TVo6Gx4Cugm1sFwLFU_wq0R1zqiiiXNriALQ4ahaY\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div>','auth',NULL),(49,'2024-03-26 19:54:31.324','2024-03-26 19:54:31.324',2,'192.168.0.108','内网地址',2,'枫枫','http://thirdqq.qlogo.cn/ek_qqapp/AQVLjicwG14hrSLsMWRFF5BEHktJiaY39GqqYDoG350dwn8Mtz7ep8fiakFnw2QOKj07soBebmw/0','info','用户登录成功','<div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDEyMzYwfQ.MJsEKHrzuIoR6cI8IOzLspnurib6-UsItKzOTqWeqfk\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div>','auth',NULL),(50,'2024-03-26 19:56:42.743','2024-03-26 19:56:42.743',2,'192.168.0.108','内网地址',2,'枫枫','http://thirdqq.qlogo.cn/ek_qqapp/AQVLjicwG14hrSLsMWRFF5BEHktJiaY39GqqYDoG350dwn8Mtz7ep8fiakFnw2QOKj07soBebmw/0','info','用户登录成功','<div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDE0MDY4fQ.gG-6nO7rBJDTPnwzBOHWqU44T94jF-l1pjaVS2GIl2Y\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div>','auth',NULL),(51,'2024-03-26 19:57:25.825','2024-03-26 19:57:25.825',2,'192.168.0.108','内网地址',2,'枫枫','http://thirdqq.qlogo.cn/ek_qqapp/AQVLjicwG14hrSLsMWRFF5BEHktJiaY39GqqYDoG350dwn8Mtz7ep8fiakFnw2QOKj07soBebmw/0','info','用户登录成功','<div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDE0MjAwfQ.i-N7QDSi5gS6O3SLGA6WVQ0eRBMe0DrS0iT0EoEDufc\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div>','auth',NULL),(52,'2024-03-26 20:07:52.366','2024-03-26 20:07:52.366',2,'192.168.0.108','内网地址',2,'枫枫','http://thirdqq.qlogo.cn/ek_qqapp/AQVLjicwG14hrSLsMWRFF5BEHktJiaY39GqqYDoG350dwn8Mtz7ep8fiakFnw2QOKj07soBebmw/0','info','用户登录成功','<div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDE0MjAwfQ.i-N7QDSi5gS6O3SLGA6WVQ0eRBMe0DrS0iT0EoEDufc\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDE0MjQzfQ.urDX8zZeVFb22S97-vgUCulnb7_18zWYkRFrYcahi40\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div>','auth',NULL),(53,'2024-03-26 20:08:40.388','2024-03-26 20:08:40.388',2,'192.168.0.108','内部地址',2,'枫枫','http://thirdqq.qlogo.cn/ek_qqapp/AQVLjicwG14hrSLsMWRFF5BEHktJiaY39GqqYDoG350dwn8Mtz7ep8fiakFnw2QOKj07soBebmw/0','info','用户登录成功','<div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDE0MjAwfQ.i-N7QDSi5gS6O3SLGA6WVQ0eRBMe0DrS0iT0EoEDufc\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDE0MjQzfQ.urDX8zZeVFb22S97-vgUCulnb7_18zWYkRFrYcahi40\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div><div class=\"log_request\">\n  <div class=\"log_request_head\">\n    <span class=\"log_request_method post\">POST</span>\n    <span class=\"log_request_path\">/api/auth/login</span>\n  </div>\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\r\n    \"userName\": \"2\",\r\n    \"password\": \"1234\"\r\n}</pre>\n  </div>\n</div><div class=\"log_request_header\">\n  <div class=\"log_request_body\">\n    <pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Content-Length\":[\"50\"],\"Content-Type\":[\"application/json\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjIsIm5pY2tuYW1lIjoi5p6r5p6rIiwicm9sZSI6MiwiZXhwIjoxNzI0NDE0ODUzfQ.8moCvS5Gab_hFShKlYge4UxrSaF0ZZcSQ8LidF669p0\"],\"User-Agent\":[\"Apifox/1.0.0 (https://apifox.com)\"],\"Validpath\":[\"/api/auth/login\"],\"X-Forwarded-For\":[\"192.168.0.108\"]}</pre>\n  </div>\n</div><div class=\"log_item info\"><div class=\"log_item_label\">nickName</div> <div class=\"log_item_content\">2</div></div>','auth',NULL),(55,'2024-03-26 22:45:45.689','2024-03-26 22:45:45.689',2,'192.168.0.108','内部地址',1,'fengfeng','','info','删除日志操作','<div class=\"log_item info\"><div class=\"log_item_label\">删除日志条数</div> <div class=\"log_item_content\">1</div></div>','logs',0);
/*!40000 ALTER TABLE `log_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `settings_models`
--

DROP TABLE IF EXISTS `settings_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `settings_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `site` longtext COLLATE utf8mb4_unicode_ci,
  `qq` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `settings_models`
--

LOCK TABLES `settings_models` WRITE;
/*!40000 ALTER TABLE `settings_models` DISABLE KEYS */;
INSERT INTO `settings_models` VALUES (1,'2024-03-28 02:06:26.606','2024-03-28 19:47:03.518','{\"created_at\":\"2024-03-28\",\"bei_an\":\"xxxx\",\"version\":\"1.0.1\",\"qq_image\":\"\",\"wechat_image\":\"\",\"bilibili_url\":\"\",\"gitee_url\":\"\",\"github_url\":\"\"}','{\"enable\":true,\"app_id\":\"101974593\",\"key\":\"db7dff0d0159b078ec225b1648b8ef50\",\"redirect\":\"http://www.fengfengzhidao.com/login?flag=qq\",\"webPath\":\"\"}');
/*!40000 ALTER TABLE `settings_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `top_user_models`
--

DROP TABLE IF EXISTS `top_user_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `top_user_models` (
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `top_user_id` bigint(20) unsigned DEFAULT NULL,
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `top_user_models`
--

LOCK TABLES `top_user_models` WRITE;
/*!40000 ALTER TABLE `top_user_models` DISABLE KEYS */;
INSERT INTO `top_user_models` VALUES (1,2,3,'2024-03-14 21:10:06.604','2024-03-14 21:10:06.604');
/*!40000 ALTER TABLE `top_user_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_chat_delete_models`
--

DROP TABLE IF EXISTS `user_chat_delete_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_chat_delete_models` (
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `chat_id` bigint(20) unsigned DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_chat_delete_models`
--

LOCK TABLES `user_chat_delete_models` WRITE;
/*!40000 ALTER TABLE `user_chat_delete_models` DISABLE KEYS */;
INSERT INTO `user_chat_delete_models` VALUES (1,3),(1,8),(1,7),(1,1);
/*!40000 ALTER TABLE `user_chat_delete_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_conf_models`
--

DROP TABLE IF EXISTS `user_conf_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_conf_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `recall_message` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `friend_online` tinyint(1) DEFAULT NULL,
  `sound` tinyint(1) DEFAULT NULL,
  `secure_link` tinyint(1) DEFAULT NULL,
  `save_pwd` tinyint(1) DEFAULT NULL,
  `search_user` tinyint(4) DEFAULT NULL,
  `verification` tinyint(4) DEFAULT NULL,
  `verification_question` longtext COLLATE utf8mb4_unicode_ci,
  `online` tinyint(1) DEFAULT NULL,
  `curtail_chat` tinyint(1) DEFAULT NULL,
  `curtail_add_user` tinyint(1) DEFAULT NULL,
  `curtail_create_group` tinyint(1) DEFAULT NULL,
  `curtail_in_group_chat` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_models_user_conf_model` (`user_id`),
  CONSTRAINT `fk_user_conf_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_user_models_user_conf_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_conf_models`
--

LOCK TABLES `user_conf_models` WRITE;
/*!40000 ALTER TABLE `user_conf_models` DISABLE KEYS */;
INSERT INTO `user_conf_models` VALUES (1,'2024-03-06 00:17:06.000','2024-03-06 23:06:33.206',1,NULL,1,1,1,0,1,NULL,'{\"problem1\":\"你的童年是什么1？\",\"problem2\":\"你的童年是什么1？\",\"problem3\":null,\"answer1\":\"上学\",\"answer2\":\"上学\",\"answer3\":null}',NULL,NULL,NULL,NULL,NULL),(2,'2024-03-06 23:22:06.000','2024-03-27 22:20:04.750',2,NULL,NULL,NULL,NULL,NULL,0,NULL,NULL,NULL,0,0,1,0),(3,'2024-03-08 22:28:06.000','2024-03-27 22:33:03.026',3,NULL,NULL,NULL,NULL,NULL,1,4,'{\"problem1\":\"你的童年是什么1？\",\"problem2\":\"你的童年是什么1？\",\"problem3\":null,\"answer1\":\"上学\",\"answer2\":\"上学\",\"answer3\":null}',NULL,0,0,1,0),(4,'2024-03-28 19:06:14.085','2024-03-28 19:06:14.085',4,NULL,0,1,0,0,2,2,NULL,1,0,0,0,0);
/*!40000 ALTER TABLE `user_conf_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_models`
--

DROP TABLE IF EXISTS `user_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `pwd` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nickname` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `abstract` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `addr` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `role` tinyint(4) DEFAULT NULL,
  `open_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `register_source` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_models`
--

LOCK TABLES `user_models` WRITE;
/*!40000 ALTER TABLE `user_models` DISABLE KEYS */;
INSERT INTO `user_models` VALUES (1,'2024-03-01 20:52:02.000','2024-03-06 23:06:33.202','$2a$04$BDhAbcEq67Q/PgOqvSb3/uXAdvLAZK1ZBpLded0DH5H8oQmVW/Nxi','fengfeng','我就是我，是不一样的花火',NULL,'127.0.0.1',NULL,1,NULL,'email'),(2,'2024-03-01 23:52:07.340','2024-03-01 23:52:07.340','$2a$04$BDhAbcEq67Q/PgOqvSb3/uXAdvLAZK1ZBpLded0DH5H8oQmVW/Nxi','枫枫','','http://thirdqq.qlogo.cn/ek_qqapp/AQVLjicwG14hrSLsMWRFF5BEHktJiaY39GqqYDoG350dwn8Mtz7ep8fiakFnw2QOKj07soBebmw/0','127.0.0.1','',2,'E98D5D2EED06C2313CDD331E8BFA2421','qq'),(3,'2024-03-01 23:52:07.340','2024-03-01 23:52:07.340','$2a$04$BDhAbcEq67Q/PgOqvSb3/uXAdvLAZK1ZBpLded0DH5H8oQmVW/Nxi','zhangsan','','http://thirdqq.qlogo.cn/ek_qqapp/AQVLjicwG14hrSLsMWRFF5BEHktJiaY39GqqYDoG350dwn8Mtz7ep8fiakFnw2QOKj07soBebmw/0','127.0.0.1','',2,'123','email'),(4,'2024-03-28 19:06:14.084','2024-03-28 19:06:14.084','','枫枫知道','','http://thirdqq.qlogo.cn/ek_qqapp/AQSMPI0aiay9hnolHxPhib1woxMdsLHWFkh13P5CFgCdYfUMOnHd5gXEdibBIMU0iaicg2LLsQdww/0','','',2,'C52BC72A1507F6367539E992CCACCB2A','qq');
/*!40000 ALTER TABLE `user_models` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-03-30 10:01:36
