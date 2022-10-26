-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 26 Okt 2022 pada 04.55
-- Versi server: 10.4.24-MariaDB
-- Versi PHP: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_fistx_harga_udang`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `harga_udang`
--

CREATE TABLE `harga_udang` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tanggal` datetime(3) DEFAULT NULL,
  `nama_penambak` longtext DEFAULT NULL,
  `harga` bigint(20) DEFAULT NULL,
  `ukuran` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `harga_udang`
--

INSERT INTO `harga_udang` (`id`, `created_at`, `updated_at`, `deleted_at`, `tanggal`, `nama_penambak`, `harga`, `ukuran`) VALUES
(1, NULL, NULL, NULL, '2022-10-25 10:00:18.000', 'asdsada', 121, 121),
(2, '2022-10-25 10:07:48.185', '2022-10-25 10:07:48.185', NULL, '2021-08-25 10:00:18.000', 'asdsada', 20, 20),
(3, '2022-10-25 10:08:21.348', '2022-10-25 10:08:21.348', NULL, '2021-08-25 10:00:18.000', 'babi', 18, 5);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `harga_udang`
--
ALTER TABLE `harga_udang`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_harga_udang_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `harga_udang`
--
ALTER TABLE `harga_udang`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
