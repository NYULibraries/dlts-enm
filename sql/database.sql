-- MySQL dump 10.13  Distrib 5.5.42, for osx10.6 (i386)
--
-- Host: localhost    Database: enm
-- ------------------------------------------------------
-- Server version	5.5.42

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `epubs`
--

DROP TABLE IF EXISTS `epubs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `epubs` (
  `tct_id` int(11) NOT NULL,
  `title` varchar(3000) NOT NULL,
  `author` varchar(3000) NOT NULL,
  `publisher` varchar(3000) NOT NULL,
  `isbn` varchar(13) NOT NULL,
  `indexpattern_id` int(11) NOT NULL,
  PRIMARY KEY (`tct_id`),
  UNIQUE KEY `isbn` (`isbn`),
  KEY `indexpattern_id` (`indexpattern_id`),
  CONSTRAINT `fk__epubs__indexpatterns` FOREIGN KEY (`indexpattern_id`) REFERENCES `indexpatterns` (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `indexpatterns`
--

DROP TABLE IF EXISTS `indexpatterns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `indexpatterns` (
  `tct_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(3000) NOT NULL,
  `pagenumber_pre_strings` varchar(512) NOT NULL,
  `pagenumber_css_selector_pattern` varchar(512) NOT NULL,
  `pagenumber_xpath_pattern` varchar(512) NOT NULL,
  `xpath_entry` varchar(512) NOT NULL,
  `see_split_strings` varchar(512) NOT NULL,
  `see_also_split_strings` varchar(512) NOT NULL,
  `xpath_see` varchar(512) NOT NULL,
  `xpath_see_also` varchar(512) NOT NULL,
  `separator_between_sees` varchar(512) NOT NULL,
  `separator_between_seealsos` varchar(512) NOT NULL,
  `separator_see_subentry` varchar(512) NOT NULL,
  `inline_see_start` varchar(512) NOT NULL,
  `inline_see_also_start` varchar(512) NOT NULL,
  `inline_see_end` varchar(512) NOT NULL,
  `inline_see_also_end` varchar(512) NOT NULL,
  `subentry_classes` varchar(512) NOT NULL,
  `separator_between_subentries` varchar(512) NOT NULL,
  `separator_between_entry_and_occurrences` varchar(512) NOT NULL,
  `separator_before_first_subentry` varchar(512) NOT NULL,
  `xpath_occurrence_link` varchar(512) NOT NULL,
  `indicators_of_occurrence_range` varchar(512) NOT NULL,
  PRIMARY KEY (`tct_id`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `locations`
--

DROP TABLE IF EXISTS `locations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `locations` (
  `tct_id` int(11) NOT NULL,
  `epub_id` int(11) NOT NULL,
  `localid` int(11) NOT NULL,
  `sequence_number` int(11) NOT NULL,
  `content_unique_descriptor` varchar(255) NOT NULL,
  `content_descriptor` varchar(255) NOT NULL,
  `content_text` text NOT NULL,
  `context` int(11) NOT NULL,
  `pagenumber_filepath` varchar(1024) NOT NULL,
  `pagenumber_tag` varchar(255) NOT NULL,
  `pagenumber_css_selector` varchar(255) NOT NULL,
  `pagenumber_xpath` varchar(255) NOT NULL,
  `next_location_id` int(11) NOT NULL,
  `previous_location_id` int(11) NOT NULL,
  PRIMARY KEY (`tct_id`),
  UNIQUE KEY `next_location_id` (`next_location_id`),
  UNIQUE KEY `previous_location_id` (`previous_location_id`) USING BTREE,
  KEY `epub_id` (`epub_id`),
  CONSTRAINT `fk__locations__epubs` FOREIGN KEY (`epub_id`) REFERENCES `epubs` (`tct_id`),
  CONSTRAINT `fk__locations__next_location_id__locations__tct_id` FOREIGN KEY (`next_location_id`) REFERENCES `locations` (`tct_id`),
  CONSTRAINT `fk__locations__previous_location_id__locations__tct_id` FOREIGN KEY (`previous_location_id`) REFERENCES `locations` (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `names`
--

DROP TABLE IF EXISTS `names`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `names` (
  `tct_id` int(11) NOT NULL,
  `topic_id` int(11) NOT NULL,
  `name` varchar(3000) NOT NULL,
  `scope_id` int(11) NOT NULL,
  `bypass` tinyint(1) NOT NULL,
  `hidden` tinyint(1) NOT NULL,
  `preferred` tinyint(1) NOT NULL,
  PRIMARY KEY (`tct_id`),
  KEY `name` (`name`(255)),
  KEY `topic_id` (`topic_id`),
  KEY `scope_id` (`scope_id`),
  CONSTRAINT `fk__names__scopes` FOREIGN KEY (`scope_id`) REFERENCES `scopes` (`tct_id`),
  CONSTRAINT `fk__names__topics` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `occurrences`
--

DROP TABLE IF EXISTS `occurrences`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `occurrences` (
  `tct_id` int(11) NOT NULL,
  `topic_id` int(11) NOT NULL,
  `ring_next` int(11) DEFAULT NULL,
  `ring_prev` int(11) NOT NULL,
  PRIMARY KEY (`tct_id`),
  KEY `topic_id` (`topic_id`),
  KEY `ring_next` (`ring_next`),
  KEY `ring_prev` (`ring_prev`),
  CONSTRAINT `fk__occurrences__ring_next__locations__tct_id` FOREIGN KEY (`ring_next`) REFERENCES `locations` (`tct_id`),
  CONSTRAINT `fk__occurrences__ring_prev__locations__tct_id` FOREIGN KEY (`ring_prev`) REFERENCES `topics` (`tct_id`),
  CONSTRAINT `fk__occurrences__topics` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `relation_direction`
--

DROP TABLE IF EXISTS `relation_direction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `relation_direction` (
  `id` int(11) NOT NULL,
  `direction` varchar(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `relation_type`
--

DROP TABLE IF EXISTS `relation_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `relation_type` (
  `tct_id` int(11) NOT NULL,
  `rtype` varchar(255) NOT NULL,
  `role_from` varchar(255) NOT NULL,
  `role_to` varchar(255) NOT NULL,
  `symmetrical` tinyint(1) NOT NULL,
  PRIMARY KEY (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `relations`
--

DROP TABLE IF EXISTS `relations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `relations` (
  `tct_id` int(11) NOT NULL,
  `relation_type_id` int(11) NOT NULL,
  `topic_id` int(11) NOT NULL,
  `relation_direction_id` int(11) NOT NULL,
  PRIMARY KEY (`tct_id`),
  KEY `relation_type_id` (`relation_type_id`),
  KEY `topic_id` (`topic_id`),
  KEY `relation_direction_id` (`relation_direction_id`),
  CONSTRAINT `fk__relations__relation_direction` FOREIGN KEY (`relation_direction_id`) REFERENCES `relation_direction` (`id`),
  CONSTRAINT `fk__relations__relation_type` FOREIGN KEY (`relation_type_id`) REFERENCES `relation_type` (`tct_id`),
  CONSTRAINT `fk__relations__topics` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `scopes`
--

DROP TABLE IF EXISTS `scopes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `scopes` (
  `tct_id` int(11) NOT NULL,
  `scope` varchar(255) NOT NULL,
  PRIMARY KEY (`tct_id`),
  KEY `scope` (`scope`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `topics`
--

DROP TABLE IF EXISTS `topics`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `topics` (
  `tct_id` int(11) NOT NULL,
  `display_name_do_not_use` varchar(3000) NOT NULL COMMENT 'Workaround for apparent knq/xo bug that prevents creation of the full set of CRUD methods  when table has only one column.',
  PRIMARY KEY (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-04-28 22:19:41
