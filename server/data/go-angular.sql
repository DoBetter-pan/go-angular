-- --------------------------------------------------------
-- 主机:                           10.10.22.70
-- 服务器版本:                        5.5.49-0ubuntu0.14.04.1 - (Ubuntu)
-- 服务器操作系统:                      debian-linux-gnu
-- HeidiSQL 版本:                  8.3.0.4694
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 go-angular 的数据库结构
DROP DATABASE IF EXISTS `go-angular`;
CREATE DATABASE IF NOT EXISTS `go-angular` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `go-angular`;


-- 导出  表 go-angular.formula 结构
DROP TABLE IF EXISTS `formula`;
CREATE TABLE IF NOT EXISTS `formula` (
  `recipeid` int(11) NOT NULL,
  `ingredientid` int(11) NOT NULL,
  `amount` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`recipeid`,`ingredientid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  go-angular.formula 的数据：~7 rows (大约)
DELETE FROM `formula`;
/*!40000 ALTER TABLE `formula` DISABLE KEYS */;
INSERT INTO `formula` (`recipeid`, `ingredientid`, `amount`) VALUES
	(1, 1, 1),
	(2, 2, 1),
	(2, 3, 2),
	(5, 8, 10),
	(5, 9, 50),
	(5, 10, 2);
/*!40000 ALTER TABLE `formula` ENABLE KEYS */;


-- 导出  表 go-angular.ingredient 结构
DROP TABLE IF EXISTS `ingredient`;
CREATE TABLE IF NOT EXISTS `ingredient` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '""',
  `unit` varchar(32) NOT NULL DEFAULT '""',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- 正在导出表  go-angular.ingredient 的数据：~7 rows (大约)
DELETE FROM `ingredient`;
/*!40000 ALTER TABLE `ingredient` DISABLE KEYS */;
INSERT INTO `ingredient` (`id`, `name`, `unit`) VALUES
	(1, 'Chips Ahoy', 'packet'),
	(2, '鱼', '条'),
	(3, '茄子', '个'),
	(7, '消息', '分'),
	(8, '鸡块', '块'),
	(9, '花生米', '粒'),
	(10, '辣椒', '只');
/*!40000 ALTER TABLE `ingredient` ENABLE KEYS */;


-- 导出  表 go-angular.recipe 结构
DROP TABLE IF EXISTS `recipe`;
CREATE TABLE IF NOT EXISTS `recipe` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(48) NOT NULL DEFAULT '""',
  `description` varchar(256) NOT NULL DEFAULT '""',
  `instructions` varchar(256) NOT NULL DEFAULT '""',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- 正在导出表  go-angular.recipe 的数据：~6 rows (大约)
DELETE FROM `recipe`;
/*!40000 ALTER TABLE `recipe` DISABLE KEYS */;
INSERT INTO `recipe` (`id`, `title`, `description`, `instructions`) VALUES
	(1, 'Cookie', 'Delicious, crisp on the outside, chewy on the outside, oozing with chocolatey goodness', '1. Go buy a paket of Chips Ahoy\\n2. Heat it up in an oven\\n3. Enjoy warm cookies\\n4. '),
	(2, '鱼香茄子', '非常好吃的鱼与茄子', '1.放入茄子，2.放入鱼，3.爆炒'),
	(5, '宫保鸡丁', '非常好吃的鸡块~~', '1、炒香辣椒\n2、放入鸡块\n3、放入花生米\n4、炒');
/*!40000 ALTER TABLE `recipe` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
