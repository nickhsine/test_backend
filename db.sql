--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `camera_id` varchar(100) NOT NULL,
  `prediction` varchar(250) DEFAULT NULL,
  `starting_timestamp` int(10) unsigned DEFAULT NULL,
  `thumbnail` varchar(512) DEFAULT NULL,
  `is_viewed` tinyint(0) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_events_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

