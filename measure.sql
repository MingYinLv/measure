-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 192.168.3.1
-- Generation Time: 2018-02-12 12:15:46
-- 服务器版本： 5.6.37
-- PHP Version: 7.0.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `measure`
--

-- --------------------------------------------------------

--
-- 表的结构 `measure`
--

CREATE TABLE `measure` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL COMMENT '机型',
  `company` varchar(255) NOT NULL COMMENT '机型所属公司',
  `address` varchar(255) NOT NULL COMMENT '公司地址',
  `phone` varchar(13) NOT NULL COMMENT '联系人以及联系电话',
  `phoneType` varchar(100) NOT NULL COMMENT '联系人类型 1.技术主管 2.采购主管',
  `machineStruct` varchar(500) NOT NULL COMMENT '机床结构',
  `designUseSpeed` varchar(200) NOT NULL COMMENT '设计使用速度',
  `workTripX` varchar(100) NOT NULL COMMENT '工作行程X轴',
  `workTripY` varchar(100) NOT NULL COMMENT '工作行程Y轴',
  `workTripZ` varchar(100) NOT NULL COMMENT '工作行程Z轴',
  `shieldBracketLengthX` varchar(100) NOT NULL COMMENT '护罩支架长度X轴',
  `shieldBracketLengthY` varchar(100) NOT NULL COMMENT '护罩支架长度Y轴',
  `shieldBracketLengthZ` varchar(100) NOT NULL COMMENT '护罩支架长度Z轴',
  `shieldBracketLengthYH` varchar(100) NOT NULL COMMENT '护罩支架长度Y后',
  `shieldStruct` varchar(100) NOT NULL COMMENT '罩结构要求  1.无要求 2.要求如下',
  `shieldStructText` varchar(500) NOT NULL COMMENT '护罩结构要求内容',
  `workBenchWidth` varchar(100) NOT NULL COMMENT '工作台宽度',
  `workBenchLength` varchar(100) NOT NULL COMMENT '工作台长度',
  `trackSurfaceToBenchBottomHeight` varchar(100) NOT NULL COMMENT '轨道面到工作台底端高度',
  `trackSurfaceToBenchTwoHeight` varchar(100) NOT NULL COMMENT '轨道面到工作台第二台阶高度',
  `trackSurfaceBenchTopHeight` varchar(100) NOT NULL COMMENT '轨道面到工作台顶面高度',
  `workBenchMiddleDistanceBefore` varchar(100) NOT NULL COMMENT '工作台中分距离前',
  `workBenchMiddleDistanceAfter` varchar(100) NOT NULL COMMENT '工作台中分距离后',
  `workBenchProtrudingSaddleDistance` varchar(100) NOT NULL COMMENT '工作台凸出滑鞍距离',
  `xInTrackWidth` varchar(100) NOT NULL COMMENT '内轨宽',
  `xOutTrackWidth` varchar(100) NOT NULL COMMENT '外轨宽',
  `trackOutToTowlineSlotOutDistance` varchar(100) NOT NULL COMMENT '轨道外侧到拖链槽外边的距离',
  `TowlineHightLineTrackSurfaceHeight` varchar(100) NOT NULL COMMENT '拖链槽定边高出线轨面的高度',
  `trackOutToTowlineSlotInDistance` varchar(100) NOT NULL COMMENT '轨道外侧到拖链槽内边的距离',
  `xLineTrackSurfaceToMotorTopHeight` varchar(100) NOT NULL COMMENT '线轨面到电机座顶边的高度',
  `xMotorWidth` varchar(100) NOT NULL COMMENT '电机座宽度',
  `xLineTrackWidth` varchar(100) NOT NULL COMMENT '线轨的宽度',
  `xLineTrackHeight` varchar(100) NOT NULL COMMENT '线轨的高度',
  `xLineTrackOutToCastingDistance` varchar(100) NOT NULL COMMENT '线轨外侧到铸件边的距离',
  `xLineTrackInToCastingDistance` varchar(100) NOT NULL COMMENT '线轨内侧到铸件边的距离',
  `xLineTrackInToMotorSideDistance` varchar(100) NOT NULL COMMENT '线轨内侧到电机座侧边的距离',
  `yValidFixedSurfaceWidth` varchar(100) NOT NULL COMMENT '有效固定面的宽度',
  `yValidFixedSurfaceHeight` varchar(100) NOT NULL COMMENT '有效固定面的高度',
  `trackSurfaceToSaddleFixedSurfaceBottomHeight` varchar(100) NOT NULL COMMENT '轨道面到滑鞍固定面底端的高度',
  `trackSurfaceToSaddleFixedSurfaceTopHeight` varchar(100) NOT NULL COMMENT '轨道面到滑鞍固定面顶端的高度',
  `yInTrackWidth` varchar(100) NOT NULL COMMENT '内轨宽',
  `yOutTrackWidth` varchar(100) NOT NULL COMMENT '外轨宽',
  `insertsSideProtrudingDistance` varchar(100) NOT NULL COMMENT '镶条侧向凸出距离',
  `insertsStretchSaddleDistance` varchar(100) NOT NULL COMMENT '镶条伸出滑鞍距离',
  `nozzleProtrudingSaddleDistance` varchar(100) NOT NULL COMMENT '油嘴凸出滑鞍距离',
  `nozzleHightTrackSurfaceDistance` varchar(100) NOT NULL COMMENT '油嘴高出轨道面的距离',
  `yMotorPlace` varchar(100) NOT NULL COMMENT '电机放置 1 前置电机 2后置电机',
  `yLineTrackSurfaceToMotorTopHeight` varchar(100) NOT NULL COMMENT '线轨面到电机座顶边的高度',
  `yMotorWidth` varchar(100) NOT NULL COMMENT '电机座宽度',
  `yLineTrackWidth` varchar(100) NOT NULL COMMENT '线轨的宽度',
  `yLineTrackHeight` varchar(100) NOT NULL COMMENT '线轨的高度',
  `yLineTrackOutToCastingDistance` varchar(100) NOT NULL COMMENT '线轨外侧到铸件边的距离',
  `yLineTrackInToCastingDistance` varchar(100) NOT NULL COMMENT '线轨内侧到铸件边的距离',
  `yLineTrackInToMotorSideDistance` varchar(100) NOT NULL COMMENT '线轨内侧到电机座侧边的距离',
  `yHStruct` varchar(100) NOT NULL COMMENT 'Y后结构 1一片式 2多片式',
  `fixedSurfaceWidth` varchar(100) NOT NULL COMMENT '固定面的宽度',
  `fixedSurfaceTopToTrackHeight` varchar(100) NOT NULL COMMENT '固定面顶边到线轨面的高度',
  `bottomSideToInBorderDistance` varchar(100) NOT NULL COMMENT '底座侧边到立柱内侧边缘的距离',
  `bottomTrackWidth` varchar(100) NOT NULL COMMENT '底座线轨支座的宽度',
  `columnOpenSize` varchar(100) NOT NULL COMMENT '立柱开裆尺寸',
  `trackToOpenTopHeight` varchar(100) NOT NULL COMMENT '轨道面到开裆顶点高度',
  `openTopLength` varchar(100) NOT NULL COMMENT '开裆顶边长度',
  `yHLineTrackWidth` varchar(100) NOT NULL COMMENT '线轨的宽度',
  `yHLineTrackHeight` varchar(100) NOT NULL COMMENT '线轨的高度',
  `yHLineTrackOutToCastingDistance` varchar(100) NOT NULL COMMENT '线轨外侧到铸件边的距离',
  `yHLineTrackInToCastingDistance` varchar(100) NOT NULL COMMENT '线轨内侧到铸件边的距离',
  `yHLineTrackInToMotorSideDistance` varchar(100) NOT NULL COMMENT '线轨内侧到电机座侧边的距离',
  `yHLineTrackSurfaceToBottomHeight` varchar(100) NOT NULL COMMENT '轨道面到立柱最低点的高度',
  `zValidFixedSurfaceWidth` varchar(100) NOT NULL COMMENT '有效固定面的宽度',
  `zValidFixedSurfaceHeight` varchar(100) NOT NULL COMMENT '有效固定面的高度	',
  `trackToFixedTopHeight` varchar(100) NOT NULL COMMENT '轨道面到固定面的顶端的高度',
  `zInTrackWidth` varchar(100) NOT NULL COMMENT '内轨宽',
  `zOutTrackWidth` varchar(100) NOT NULL COMMENT '外轨宽',
  `zLineTrackWidth` varchar(100) NOT NULL COMMENT '线轨的宽度',
  `zLineTrackHeight` varchar(100) NOT NULL COMMENT '线轨的高度',
  `zLineTrackOutToCastingDistance` varchar(100) NOT NULL COMMENT '线轨外侧到铸件边的距离',
  `zLineTrackInToCastingDistance` varchar(100) NOT NULL COMMENT '线轨内侧到铸件边的距离',
  `zLineTrackInToBearingSideDistance` varchar(100) NOT NULL COMMENT '线轨内侧到轴承座侧边的距离',
  `zLineTrackToBearingTopHeight` varchar(100) NOT NULL COMMENT '线轨面到轴承座顶边的高度'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `measure`
--

INSERT INTO `measure` (`id`, `name`, `company`, `address`, `phone`, `phoneType`, `machineStruct`, `designUseSpeed`, `workTripX`, `workTripY`, `workTripZ`, `shieldBracketLengthX`, `shieldBracketLengthY`, `shieldBracketLengthZ`, `shieldBracketLengthYH`, `shieldStruct`, `shieldStructText`, `workBenchWidth`, `workBenchLength`, `trackSurfaceToBenchBottomHeight`, `trackSurfaceToBenchTwoHeight`, `trackSurfaceBenchTopHeight`, `workBenchMiddleDistanceBefore`, `workBenchMiddleDistanceAfter`, `workBenchProtrudingSaddleDistance`, `xInTrackWidth`, `xOutTrackWidth`, `trackOutToTowlineSlotOutDistance`, `TowlineHightLineTrackSurfaceHeight`, `trackOutToTowlineSlotInDistance`, `xLineTrackSurfaceToMotorTopHeight`, `xMotorWidth`, `xLineTrackWidth`, `xLineTrackHeight`, `xLineTrackOutToCastingDistance`, `xLineTrackInToCastingDistance`, `xLineTrackInToMotorSideDistance`, `yValidFixedSurfaceWidth`, `yValidFixedSurfaceHeight`, `trackSurfaceToSaddleFixedSurfaceBottomHeight`, `trackSurfaceToSaddleFixedSurfaceTopHeight`, `yInTrackWidth`, `yOutTrackWidth`, `insertsSideProtrudingDistance`, `insertsStretchSaddleDistance`, `nozzleProtrudingSaddleDistance`, `nozzleHightTrackSurfaceDistance`, `yMotorPlace`, `yLineTrackSurfaceToMotorTopHeight`, `yMotorWidth`, `yLineTrackWidth`, `yLineTrackHeight`, `yLineTrackOutToCastingDistance`, `yLineTrackInToCastingDistance`, `yLineTrackInToMotorSideDistance`, `yHStruct`, `fixedSurfaceWidth`, `fixedSurfaceTopToTrackHeight`, `bottomSideToInBorderDistance`, `bottomTrackWidth`, `columnOpenSize`, `trackToOpenTopHeight`, `openTopLength`, `yHLineTrackWidth`, `yHLineTrackHeight`, `yHLineTrackOutToCastingDistance`, `yHLineTrackInToCastingDistance`, `yHLineTrackInToMotorSideDistance`, `yHLineTrackSurfaceToBottomHeight`, `zValidFixedSurfaceWidth`, `zValidFixedSurfaceHeight`, `trackToFixedTopHeight`, `zInTrackWidth`, `zOutTrackWidth`, `zLineTrackWidth`, `zLineTrackHeight`, `zLineTrackOutToCastingDistance`, `zLineTrackInToCastingDistance`, `zLineTrackInToBearingSideDistance`, `zLineTrackToBearingTopHeight`) VALUES
(12, '测试型号', '小米', '北京', '12345678910', '技术主管', '三线轨', '48M/min以内', '10m', '20m', '10m', '10m', '20m', '20m', '30m', '要求如下', '没有要求', '200mm', '100mm', '150mm', '160mm', '170mm', '20m', '20m', '150mm', '150cm', '150mm', '40cm', '20cm', '15cm', '18cm', '20cm', '100cm', '90cm', '123m', '22m', '20cm', '18cm', '34cm', '230cm', '39cm', '23m', '23cm', '23cm', '2312mm', '12cm', '30cm', '后置电机', '20m', '23cm', '123cm', '12m', '23m', '20m', '12cm', '多片式', '23cm', '12cm', '12m', '32cm', '98m', '35m', '12m', '2m', '12m', '2cm', '8cm', '10m', '12cm', '12m', '90m', '39cm', '100mm', '100mm', '100mm', '100mm', '100mm', '100mm', '100mm', '100mm');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `measure`
--
ALTER TABLE `measure`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `measure`
--
ALTER TABLE `measure`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
