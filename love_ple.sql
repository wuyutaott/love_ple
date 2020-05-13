/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : love_ple

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2020-05-13 08:13:22
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `card_id` int(20) unsigned zerofill DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `gender` bit(1) DEFAULT NULL,
  `register_date` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of member
-- ----------------------------
INSERT INTO `member` VALUES ('1', '00000000000000001001', '18600756843', '伍玉涛', '1985-10-01', '', '2020-05-12');
INSERT INTO `member` VALUES ('2', '00000000000000001002', '18701089701', '李富婷', '1985-08-04', '\0', '2020-05-12');
INSERT INTO `member` VALUES ('3', '00000000000000001003', null, '杨志强', null, '', null);

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `member_id` int(11) unsigned zerofill DEFAULT NULL,
  `order_date` datetime NOT NULL,
  `money` float(10,1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES ('1', '00000000001', '2020-05-12 22:51:22', '100.5');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `password` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'yangzhiqiang', '123456');
