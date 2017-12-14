-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 14, 2017 at 03:12 AM
-- Server version: 10.1.28-MariaDB
-- PHP Version: 7.1.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `scheduler`
--

-- --------------------------------------------------------

--
-- Table structure for table `jadwal`
--

CREATE TABLE `jadwal` (
  `ID` int(11) NOT NULL,
  `Hari` int(11) NOT NULL,
  `Bulan` int(11) NOT NULL,
  `Tahun` int(11) NOT NULL,
  `Jam` int(11) NOT NULL,
  `Tempat` text NOT NULL,
  `Kegiatan` text NOT NULL,
  `Keterangan` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `jadwal`
--

INSERT INTO `jadwal` (`ID`, `Hari`, `Bulan`, `Tahun`, `Jam`, `Tempat`, `Kegiatan`, `Keterangan`) VALUES
(1, 13, 12, 17, 18, 'Lab SI', 'Ngambis Kapsel', 'Bareng Winaldo'),
(2, 1, 1, 18, 0, '-', 'Tahun baru', '-'),
(3, 10, 4, 18, 0, '-', 'ultah', '-'),
(4, 14, 12, 17, 9, '7602', 'UAS Progif', 'Mantab'),
(5, 15, 12, 17, 9, '7601', 'UAS Kapsel', 'Last!!'),
(6, 17, 12, 17, 13, 'Baranangsiang', 'Natalan', 'Berangkat cepat'),
(7, 19, 12, 17, 8, 'TVST', 'TFT SSDK', 'Jangan telat'),
(8, 18, 12, 17, 9, 'Elcavana', 'Nge-gym', 'Gausah bayar pendaftaran'),
(9, 14, 12, 17, 18, 'ITB', 'UAS Sismul', '-'),
(10, 19, 3, 18, 0, '-', 'ultah papa', '-');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `jadwal`
--
ALTER TABLE `jadwal`
  ADD PRIMARY KEY (`ID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `jadwal`
--
ALTER TABLE `jadwal`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
