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
-- Table structure for table `editorial_review_status_state`
--

DROP TABLE IF EXISTS `editorial_review_status_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `editorial_review_status_state` (
  `id` int(11) NOT NULL,
  `state` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
-- Temporary table structure for view `epubs_number_of_pages`
--

DROP TABLE IF EXISTS `epubs_number_of_pages`;
/*!50001 DROP VIEW IF EXISTS `epubs_number_of_pages`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `epubs_number_of_pages` (
  `ISBN` tinyint NOT NULL,
  `number_of_pages` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

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
  `localid` varchar(255) NOT NULL,
  `sequence_number` int(11) NOT NULL,
  `content_unique_indicator` varchar(255) NOT NULL,
  `content_descriptor` varchar(255) NOT NULL,
  `content_text` text NOT NULL,
  `pagenumber_filepath` varchar(1024) NOT NULL,
  `pagenumber_tag` varchar(255) NOT NULL,
  `pagenumber_css_selector` varchar(255) NOT NULL,
  `pagenumber_xpath` varchar(255) NOT NULL,
  `next_location_id` int(11) DEFAULT NULL,
  `previous_location_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`tct_id`),
  UNIQUE KEY `previous_location_id` (`previous_location_id`) USING BTREE,
  UNIQUE KEY `next_location_id` (`next_location_id`),
  KEY `epub_id` (`epub_id`),
  CONSTRAINT `fk__locations__epubs` FOREIGN KEY (`epub_id`) REFERENCES `epubs` (`tct_id`)
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
  `location_id` int(11) NOT NULL,
  `topic_id` int(11) NOT NULL,
  `ring_next` int(11) DEFAULT NULL,
  `ring_prev` int(11) DEFAULT NULL,
  PRIMARY KEY (`tct_id`),
  KEY `topic_id` (`topic_id`),
  KEY `ring_next` (`ring_next`),
  KEY `ring_prev` (`ring_prev`),
  KEY `location_id` (`location_id`),
  CONSTRAINT `fk__occurrences__locations` FOREIGN KEY (`location_id`) REFERENCES `locations` (`tct_id`),
  CONSTRAINT `fk__occurrences__topics` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Temporary table structure for view `page_topic_names`
--

DROP TABLE IF EXISTS `page_topic_names`;
/*!50001 DROP VIEW IF EXISTS `page_topic_names`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `page_topic_names` (
  `page_id` tinyint NOT NULL,
  `topic_id` tinyint NOT NULL,
  `topic_display_name` tinyint NOT NULL,
  `topic_name` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Temporary table structure for view `pages`
--

DROP TABLE IF EXISTS `pages`;
/*!50001 DROP VIEW IF EXISTS `pages`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `pages` (
  `id` tinyint NOT NULL,
  `title` tinyint NOT NULL,
  `authors` tinyint NOT NULL,
  `publisher` tinyint NOT NULL,
  `isbn` tinyint NOT NULL,
  `page_pattern` tinyint NOT NULL,
  `page_localid` tinyint NOT NULL,
  `page_sequence` tinyint NOT NULL,
  `page_text` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `readium_goto_urls`
--

DROP TABLE IF EXISTS `readium_goto_urls`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `readium_goto_urls` (
  `location_id` int(11) NOT NULL,
  `readium_goto_value` varchar(1024) NOT NULL,
  `readium_goto_url` varchar(1024) NOT NULL,
  `readium_goto_value_encoded` varchar(1024) NOT NULL,
  `readium_goto_url_encoded` varchar(1024) NOT NULL,
  `stage_readium_goto_value` varchar(1024) NOT NULL,
  `stage_readium_goto_url` varchar(1024) NOT NULL,
  `stage_readium_goto_value_encoded` varchar(1024) NOT NULL,
  `stage_readium_goto_url_encoded` varchar(1024) NOT NULL,
  `dev_readium_goto_value` varchar(1024) NOT NULL,
  `dev_readium_goto_url` varchar(1024) NOT NULL,
  `dev_readium_goto_value_encoded` varchar(1024) NOT NULL,
  `dev_readium_goto_url_encoded` varchar(1024) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `readium_goto_urls_chrome_bug_workaround`
--

DROP TABLE IF EXISTS `readium_goto_urls_chrome_bug_workaround`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `readium_goto_urls_chrome_bug_workaround` (
  `location_id` int(11) NOT NULL,
  `readium_goto_value` varchar(1024) NOT NULL,
  `readium_goto_url` varchar(1024) NOT NULL,
  `readium_goto_value_encoded` varchar(1024) NOT NULL,
  `readium_goto_url_encoded` varchar(1024) NOT NULL,
  `stage_readium_goto_value` varchar(1024) NOT NULL,
  `stage_readium_goto_url` varchar(1024) NOT NULL,
  `stage_readium_goto_value_encoded` varchar(1024) NOT NULL,
  `stage_readium_goto_url_encoded` varchar(1024) NOT NULL,
  `dev_readium_goto_value` varchar(1024) NOT NULL,
  `dev_readium_goto_url` varchar(1024) NOT NULL,
  `dev_readium_goto_value_encoded` varchar(1024) NOT NULL,
  `dev_readium_goto_url_encoded` varchar(1024) NOT NULL
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
  `relation_direction_id` int(11) NOT NULL,
  `role_from_topic_id` int(11) NOT NULL,
  `role_to_topic_id` int(11) NOT NULL,
  PRIMARY KEY (`tct_id`),
  KEY `relation_type_id` (`relation_type_id`),
  KEY `relation_direction_id` (`relation_direction_id`),
  KEY `role_from_topic_id` (`role_from_topic_id`) USING BTREE,
  KEY `role_to_topic_id` (`role_to_topic_id`) USING BTREE,
  CONSTRAINT `fk__relations__relation_direction` FOREIGN KEY (`relation_direction_id`) REFERENCES `relation_direction` (`id`),
  CONSTRAINT `fk__relations__relation_type` FOREIGN KEY (`relation_type_id`) REFERENCES `relation_type` (`tct_id`),
  CONSTRAINT `fk__relations__topics__role_from` FOREIGN KEY (`role_from_topic_id`) REFERENCES `topics` (`tct_id`),
  CONSTRAINT `fk__relations__topics__role_to` FOREIGN KEY (`role_to_topic_id`) REFERENCES `topics` (`tct_id`)
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
-- Temporary table structure for view `topic_relations_simple`
--

DROP TABLE IF EXISTS `topic_relations_simple`;
/*!50001 DROP VIEW IF EXISTS `topic_relations_simple`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `topic_relations_simple` (
  `topic1_id` tinyint NOT NULL,
  `topic2_id` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `topic_type`
--

DROP TABLE IF EXISTS `topic_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `topic_type` (
  `tct_id` int(11) NOT NULL,
  `ttype` varchar(3000) NOT NULL,
  PRIMARY KEY (`tct_id`)
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
  `editorial_review_status_reviewer` varchar(255) DEFAULT NULL,
  `editorial_review_status_time` varchar(28) DEFAULT NULL,
  `editorial_review_status_state_id` int(11) NOT NULL,
  PRIMARY KEY (`tct_id`),
  KEY `editorial_review_status_state_id` (`editorial_review_status_state_id`) USING BTREE,
  CONSTRAINT `fk__topics__editorial_review_status_state` FOREIGN KEY (`editorial_review_status_state_id`) REFERENCES `editorial_review_status_state` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `topics_topic_type`
--

DROP TABLE IF EXISTS `topics_topic_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `topics_topic_type` (
  `topic_id` int(11) NOT NULL,
  `topic_type_id` int(11) NOT NULL,
  KEY `topic_id` (`topic_id`),
  KEY `topic_type_id` (`topic_type_id`),
  CONSTRAINT `fk__topics_topic_type__topics` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`tct_id`),
  CONSTRAINT `fk__topics_topic_type__topic_type` FOREIGN KEY (`topic_type_id`) REFERENCES `topic_type` (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `topics_weblinks`
--

DROP TABLE IF EXISTS `topics_weblinks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `topics_weblinks` (
  `topic_id` int(11) NOT NULL,
  `weblink_id` int(11) NOT NULL,
  UNIQUE KEY `idx_unique_topics_weblinks` (`topic_id`,`weblink_id`),
  KEY `topic_id` (`topic_id`),
  KEY `weblink_id` (`weblink_id`),
  CONSTRAINT `fk__topics_weblinks__topics` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`tct_id`),
  CONSTRAINT `fk__topics_weblinks__weblinks` FOREIGN KEY (`weblink_id`) REFERENCES `weblinks` (`tct_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `weblinks`
--

DROP TABLE IF EXISTS `weblinks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `weblinks` (
  `tct_id` int(11) NOT NULL,
  `url` varchar(3000) NOT NULL,
  `weblinks_relationship_id` int(11) NOT NULL,
  `weblinks_vocabulary_id` int(11) NOT NULL,
  PRIMARY KEY (`tct_id`),
  KEY `weblinks_relationship_id` (`weblinks_relationship_id`),
  KEY `weblinks_vocabulary_id` (`weblinks_vocabulary_id`),
  CONSTRAINT `fk__weblinks__weblinks_relationship` FOREIGN KEY (`weblinks_relationship_id`) REFERENCES `weblinks_relationship` (`id`),
  CONSTRAINT `fk__weblinks__weblinks_vocabulary` FOREIGN KEY (`weblinks_vocabulary_id`) REFERENCES `weblinks_vocabulary` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `weblinks_relationship`
--

DROP TABLE IF EXISTS `weblinks_relationship`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `weblinks_relationship` (
  `id` int(11) NOT NULL,
  `relationship` varchar(3000) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `weblinks_vocabulary`
--

DROP TABLE IF EXISTS `weblinks_vocabulary`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `weblinks_vocabulary` (
  `id` int(11) NOT NULL,
  `vocabulary` varchar(3000) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Final view structure for view `epubs_number_of_pages`
--

/*!50001 DROP TABLE IF EXISTS `epubs_number_of_pages`*/;
/*!50001 DROP VIEW IF EXISTS `epubs_number_of_pages`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_unicode_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `epubs_number_of_pages` AS select `epubs`.`isbn` AS `ISBN`,count(0) AS `number_of_pages` from (`epubs` join `locations` on((`epubs`.`tct_id` = `locations`.`epub_id`))) group by `epubs`.`isbn` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `page_topic_names`
--

/*!50001 DROP TABLE IF EXISTS `page_topic_names`*/;
/*!50001 DROP VIEW IF EXISTS `page_topic_names`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_unicode_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `page_topic_names` AS select `p`.`id` AS `page_id`,`t`.`tct_id` AS `topic_id`,`t`.`display_name_do_not_use` AS `topic_display_name`,`n`.`name` AS `topic_name` from (((`pages` `p` left join `occurrences` `o` on((`p`.`id` = `o`.`location_id`))) left join `topics` `t` on((`o`.`topic_id` = `t`.`tct_id`))) join `names` `n` on((`n`.`topic_id` = `t`.`tct_id`))) where (`t`.`tct_id` is not null) order by `p`.`id`,`t`.`tct_id`,`n`.`name` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `pages`
--

/*!50001 DROP TABLE IF EXISTS `pages`*/;
/*!50001 DROP VIEW IF EXISTS `pages`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `pages` AS select `l`.`tct_id` AS `id`,`e`.`title` AS `title`,`e`.`author` AS `authors`,`e`.`publisher` AS `publisher`,`e`.`isbn` AS `isbn`,`i`.`pagenumber_css_selector_pattern` AS `page_pattern`,`l`.`localid` AS `page_localid`,`l`.`sequence_number` AS `page_sequence`,`l`.`content_text` AS `page_text` from ((`locations` `l` left join `epubs` `e` on((`l`.`epub_id` = `e`.`tct_id`))) left join `indexpatterns` `i` on((`e`.`indexpattern_id` = `i`.`tct_id`))) order by `l`.`tct_id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `topic_relations_simple`
--

/*!50001 DROP TABLE IF EXISTS `topic_relations_simple`*/;
/*!50001 DROP VIEW IF EXISTS `topic_relations_simple`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_unicode_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `topic_relations_simple` AS select `r`.`role_from_topic_id` AS `topic1_id`,`r`.`role_to_topic_id` AS `topic2_id` from `relations` `r` union select `r`.`role_to_topic_id` AS `topic1_id`,`r`.`role_from_topic_id` AS `topic2_id` from `relations` `r` order by `topic1_id`,`topic2_id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-15 17:10:44
