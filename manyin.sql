-- MySQL dump 10.13  Distrib 8.0.25, for Linux (x86_64)
--
-- Host: localhost    Database: manyin
-- ------------------------------------------------------
-- Server version	8.0.25

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
-- Table structure for table `address`
--

DROP TABLE IF EXISTS `address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `address` (
  `uuid` char(50) DEFAULT NULL,
  `id` char(50) NOT NULL,
  `name` char(20) NOT NULL,
  `addressee` char(30) NOT NULL,
  `address` text NOT NULL,
  `phone` char(20) NOT NULL,
  `is_default` tinyint(1) NOT NULL,
  `ava` tinyint(1) NOT NULL,
  KEY `uuid` (`uuid`),
  CONSTRAINT `address_ibfk_1` FOREIGN KEY (`uuid`) REFERENCES `user` (`uuid`),
  CONSTRAINT `address_chk_1` CHECK ((`ava` = 1)),
  CONSTRAINT `address_chk_2` CHECK (((`is_default` = 0) or (`is_default` = 1)))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `address`
--

LOCK TABLES `address` WRITE;
/*!40000 ALTER TABLE `address` DISABLE KEYS */;
/*!40000 ALTER TABLE `address` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `enterprise_info`
--

DROP TABLE IF EXISTS `enterprise_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `enterprise_info` (
  `user_id` char(50) NOT NULL,
  `full_name` varchar(50) NOT NULL,
  `social_code` varchar(60) DEFAULT NULL,
  `registry_money` int DEFAULT NULL,
  `enter_address` varchar(50) DEFAULT NULL,
  `operator_name` char(20) DEFAULT NULL,
  `operator_email` varchar(100) DEFAULT NULL,
  `operator_phone` char(16) DEFAULT NULL,
  `create_date` timestamp NOT NULL,
  `status` tinyint DEFAULT NULL COMMENT '1:审核中, 2:审核通过, 3:不通过',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `social_code` (`social_code`),
  CONSTRAINT `enterprise_info_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `enterprise_info`
--

LOCK TABLES `enterprise_info` WRITE;
/*!40000 ALTER TABLE `enterprise_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `enterprise_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `paid_order`
--

DROP TABLE IF EXISTS `paid_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `paid_order` (
  `order_id` bigint NOT NULL,
  `user_id` char(50) DEFAULT NULL,
  `prod_id` char(40) DEFAULT NULL,
  `prod_name` varchar(30) NOT NULL,
  `prod_avatar` text NOT NULL,
  `total_price` int NOT NULL,
  `buy_num` int NOT NULL,
  `addressee` varchar(30) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `addressee_phone` varchar(20) DEFAULT NULL,
  `create_time` timestamp NOT NULL,
  `status` int NOT NULL,
  PRIMARY KEY (`order_id`),
  KEY `user_id` (`user_id`),
  KEY `prod_id` (`prod_id`),
  CONSTRAINT `paid_order_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`uuid`),
  CONSTRAINT `paid_order_ibfk_2` FOREIGN KEY (`prod_id`) REFERENCES `product` (`prod_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paid_order`
--

LOCK TABLES `paid_order` WRITE;
/*!40000 ALTER TABLE `paid_order` DISABLE KEYS */;
/*!40000 ALTER TABLE `paid_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `prod_id` char(40) NOT NULL,
  `prod_name` varchar(30) NOT NULL,
  `prod_desc` text NOT NULL,
  `prod_type_id` tinyint NOT NULL,
  `prod_avatar` text NOT NULL,
  `prod_digital` text NOT NULL,
  `prod_price` int NOT NULL,
  `prod_unit` char(5) NOT NULL,
  PRIMARY KEY (`prod_id`),
  KEY `prod_type_id` (`prod_type_id`),
  CONSTRAINT `product_ibfk_1` FOREIGN KEY (`prod_type_id`) REFERENCES `product_types` (`prod_t_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES ('19aa5ac90aaf4e2eb96c980e743e0b7a','柱','通过对建筑垃圾的再生利用生产预制柱，不仅可以有效地减少施工成本，同时采用预制件的建筑工地现场作业量明显减少，能有效降低粉尘污染、噪音污染',1,'/assets/img/item-avatar-19aa5ac90aaf4e2eb96c980e743e0b7a.png','/assets/img/item-digital-19aa5ac90aaf4e2eb96c980e743e0b7a.png',12500,'个'),('1f69974afd414a46b0e7bc9e1407f2d3','沥青冷补料','沥青冷补料，沥青材料完成分选分离后可以循环利用，制成用于铺筑路面面层或基层的材料。',2,'/assets/img/item-avatar-1f69974afd414a46b0e7bc9e1407f2d3.png','/assets/img/item-digital-1f69974afd414a46b0e7bc9e1407f2d3.png',72000,'吨'),('5f84edbeaf924b89aa62fdd47ce2862a','建筑垃圾再生骨料透水砖','透气透水，改善气候，海绵城市，缓解热岛，铺装简易，方便经济，吸收噪音，舒适安全，规格多样，色彩丰富，自然朴实，经济实惠的特点，适用于市政广场，人行步道，停车场，自行车道，景观道路等透水性要求高的场所。',3,'/assets/img/item-avatar-5f84edbeaf924b89aa62fdd47ce2862a.png','/assets/img/item-digital-5f84edbeaf924b89aa62fdd47ce2862a.png',35,'块'),('7ae1f3f4c5714f33958e4a2de699e317','可爱的动物玩具','在你的心中是否希望有一群动物可以陪着你但是又因为家中的空间不足而烦恼呢?这些动物们是你的不二选择，你可以挑选可爱的长颈鹿，亦或是呆萌的兔子，有了他们的陪伴，你将不会在觉得孤独。',4,'/assets/img/item-avatar-7ae1f3f4c5714f33958e4a2de699e317.png','/assets/img/item-digital-7ae1f3f4c5714f33958e4a2de699e317.png',1000,'个'),('81462f72bb1249a59944f5526a1855d4','再生彩色海绵','无论是理发时擦去身上的头发屑，抑或是洗碗时洗去碗筷上的油污，这一抹彩色必将使你心情愉悦，向着心中更好的生活努力进发，',4,'/assets/img/item-avatar-81462f72bb1249a59944f5526a1855d4.png','/assets/img/item-digital-81462f72bb1249a59944f5526a1855d4.png',1000,'个'),('85b967b2a9584a84821326b1dca06011','再生植草砖','通过废弃的建筑垃圾再生后的植草砖，适用于别墅花园后的幽静小路，也适用于公园草地中的绿茵小道，有了它，会给绿意盎然的草地增添不一样的美。',4,'/assets/img/item-avatar-85b967b2a9584a84821326b1dca06011.png','/assets/img/item-digital-85b967b2a9584a84821326b1dca06011.png',20000,'平米'),('8af6b5f5fd56463cb5e133f298d52ebe','再生混凝土骨料','应用场景大部分是替代或部分替代天然砂，用于砂石垫层、道路铺设、制砖，拌制混凝土、水稳、二灰石喷砂、路基填筑、极配砂石等，适用于工业与民用建筑工程，市政工程、园林绿化、装修装饰、道路工程，海绵城市建筑工程等。',3,'/assets/img/item-avatar-8af6b5f5fd56463cb5e133f298d52ebe.png','/assets/img/item-digital-8af6b5f5fd56463cb5e133f298d52ebe.png',3000,'吨'),('a5c281bebfec41e4aa63a9671773bd12','板式楼梯','利用再生材料、再生骨料等制成预制板式楼梯，安装后可做施工通道，避免工人踩坏楼梯模板和钢筋。',1,'/assets/img/item-avatar-a5c281bebfec41e4aa63a9671773bd12.png','/assets/img/item-digital-a5c281bebfec41e4aa63a9671773bd12.png',45000,'个'),('ae0b4d1dab484303af5e9a82984ce4d8','过梁','利用再生材料制作的预制过梁事前制备好再运往现场安装，减少了80%的现场施工劳动力，还大大减轻现场工人的劳动强度，保证施工人员的安全作业。在确保安全和质量的同时，也显著提高了工效。',1,'/assets/img/item-avatar-ae0b4d1dab484303af5e9a82984ce4d8.png','/assets/img/item-digital-ae0b4d1dab484303af5e9a82984ce4d8.png',12500,'个'),('be4afdb4cfad435b9ace96d161a67f9f','建筑垃圾再生骨料实心砖','强度高，不怕水抗风化，耐腐蚀，粉层薄，造价低，施工快，破损低，尺寸标准，离散小，再生资源，循环利用的特点，可依据需要用于各种建筑内外墙，承重墙，填充墙，围护墙等部位，适用于工业与民用建筑园林绿化，市政交通，装饰装修，安装工程。',3,'/assets/img/item-avatar-be4afdb4cfad435b9ace96d161a67f9f.png','/assets/img/item-digital-be4afdb4cfad435b9ace96d161a67f9f.png',20,'块'),('dd7df47af1ea49ffb9494e23abe12682','观赏型混凝土茶艺再生工艺品','可以用混凝土做成各式各样的茶艺混凝土工艺品，满足消费者对观赏品的一切颜色以及形状要求，表面光滑，形态含蓄内敛，非常华丽。',4,'/assets/img/item-avatar-dd7df47af1ea49ffb9494e23abe12682.png','/assets/img/item-digital-dd7df47af1ea49ffb9494e23abe12682.png',5000,'个'),('e2b4fa0260da4cf780e29a9db9cb138a','复合板材','复合板材，旧木材，木屑作为比较常见的建筑垃圾，通过改进加工可以做成复合板材，使用广泛，可用于家装工装店面装饰等场所。',2,'/assets/img/item-avatar-e2b4fa0260da4cf780e29a9db9cb138a.png','/assets/img/item-digital-e2b4fa0260da4cf780e29a9db9cb138a.png',4300,'元'),('f5d65177a675456d92ef7b577d588d0d','内隔墙','本产品利用建筑垃圾再生材料，适当掺加粉煤灰、矿渣和外加剂，复合而成的轻质墙板材料本产品与轻钢结构镶嵌安装的新型节能建房具有快速施工，防火抗震，隔音降噪，防水防潮等特点，开创建筑节能新时代。',1,'/assets/img/item-avatar-f5d65177a675456d92ef7b577d588d0d.png','/assets/img/item-digital-f5d65177a675456d92ef7b577d588d0d.png',40000,'个');
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_types`
--

DROP TABLE IF EXISTS `product_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_types` (
  `prod_t_id` tinyint NOT NULL AUTO_INCREMENT,
  `t_name` varchar(50) NOT NULL,
  `sub_text` varchar(50) NOT NULL,
  `description` text NOT NULL,
  PRIMARY KEY (`prod_t_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_types`
--

LOCK TABLES `product_types` WRITE;
/*!40000 ALTER TABLE `product_types` DISABLE KEYS */;
INSERT INTO `product_types` VALUES (1,'装配式建筑','用于装配式建筑的一体化建筑设计和施工，配件可多次使用，减少垃圾产生','此产品可用于装配式建筑的一体化建筑设计和施工，在工厂进行加工后，直接运送至工地进行装配，建筑时无需搭建外架、手脚板等，可以有效节约人力和物力。装配式构配件尚可多次使用，减少建筑垃圾的产生。'),(2,'再生建材','蔓茵科技的再生材料含杂率低，原材料符合绿色回收标准，可节约工程成本','蔓茵科技的再生建材含杂率低，生产时进行了颗粒大小、生产成份种类的精准细分，在抗水、抗压等方面与普通建材一致，能够达到工程建设所用的标准。此外，再生建材的原材料符合绿色回收标准，能够节约部分建筑公司的工程成本。'),(3,'再生骨料','根据颗粒大小对再生骨料进行了精准分类，继续针对性生产，提高了利用率，间接提高了再生建材的整体性能。','再生骨料表面粗糙且母体混凝土块在经过本公司的设备进行破碎处理后，表面砂浆内部存在大量微裂纹，从而使其吸水率远高于天然骨料，且本公司对再生骨料进行了精准分类，根据不同的颗粒大小，继续针对性生产，提高了再生骨料的利用率，也间接提高了再生建材的整体性能。'),(4,'再生工艺品','将不符合标准的再生原材料进行多重利用，将其生产为再生工艺品','蔓茵回收将不符合建筑工程施工标准的再生原材料进行多重利用，将其生产为再生工艺品，可放在公共场所或店面装饰，起到美化环境、提高资源利用率的作用。');
/*!40000 ALTER TABLE `product_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `uuid` char(50) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone` char(16) DEFAULT NULL,
  `nike_name` char(50) NOT NULL,
  `password` char(65) NOT NULL,
  `salt` char(65) NOT NULL,
  `reg_date` timestamp NOT NULL,
  PRIMARY KEY (`uuid`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-21 10:22:49
