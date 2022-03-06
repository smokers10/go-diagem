-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 06, 2022 at 10:51 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 7.4.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `diagem`
--

-- --------------------------------------------------------

--
-- Table structure for table `admins`
--

CREATE TABLE `admins` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `nama` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `roles` enum('super admin','marketing','admin','') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'admin',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `admins`
--

INSERT INTO `admins` (`id`, `nama`, `username`, `email`, `password`, `roles`, `created_at`, `updated_at`, `email_verified_at`) VALUES
(1, 'super admin', 'superadmin', 'superadmin@diagem.id', '$2a$12$9luAOV0UKNv/e3KNsE5Keu1EQGomXyuWNTG.oQehwKz1UHI7o.IgO', 'super admin', '2022-02-27 14:27:46', '2022-02-27 14:27:47', '2022-02-27 14:27:47'),
(2, 'marketing1', 'marketing1', 'marketing1@diagem.id', '$2a$12$x/rXWYkDdiX2wzasK2z14uIeDAIv4D5kWLT9k4LB4cuK/htFVCnuS', 'marketing', '2022-02-27 14:27:47', '2022-02-27 14:27:47', '2022-02-27 14:27:47'),
(3, 'admin', 'admin', 'admin@diagem.id', '$2a$12$sVC2GTj.SXuWkAyM67csOu5w7FuIjHKjbU2X17K8FETuUR2iOFqAW', 'admin', '2022-02-27 14:30:23', '2022-02-27 14:30:23', '2022-02-27 14:30:23');

-- --------------------------------------------------------

--
-- Table structure for table `alamat`
--

CREATE TABLE `alamat` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `alamat` text NOT NULL,
  `nama` varchar(191) NOT NULL,
  `penerima` varchar(191) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `provinsi_id` varchar(30) NOT NULL,
  `kota_id` varchar(30) NOT NULL,
  `nama_provinsi` varchar(30) NOT NULL,
  `nama_kota` varchar(30) NOT NULL,
  `kd_pos` char(8) NOT NULL,
  `is_utama` tinyint(4) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `alamat_origin`
--

CREATE TABLE `alamat_origin` (
  `id` int(11) NOT NULL,
  `kontak_admin` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `alamat_lengkap` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `provinsi_id` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kota_id` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_provinsi` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_kota` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kd_pos` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `alamat_origin`
--

INSERT INTO `alamat_origin` (`id`, `kontak_admin`, `alamat_lengkap`, `provinsi_id`, `kota_id`, `nama_provinsi`, `nama_kota`, `kd_pos`) VALUES
(1, '081-123-123-123', 'Jl. Ir. H. Juanda No.8A Coblong', '6', '153', 'DKI Jakarta', 'Jakarta Selatan', '12230');

-- --------------------------------------------------------

--
-- Table structure for table `bank`
--

CREATE TABLE `bank` (
  `id` bigint(20) NOT NULL,
  `no_va` varchar(50) NOT NULL,
  `vendor` enum('bca_va','bni_va','bri_va','') NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `bank`
--

INSERT INTO `bank` (`id`, `no_va`, `vendor`, `status`) VALUES
(1, '800088812442312', 'bca_va', 1),
(2, '800088812442334', 'bni_va', 1),
(3, '800066512423223', 'bri_va', 1);

-- --------------------------------------------------------

--
-- Table structure for table `blogs`
--

CREATE TABLE `blogs` (
  `id` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `judul` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `isi` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `image` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_publish` tinyint(1) NOT NULL,
  `seo_keyword` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `seo_tags` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `seo_deskripsi` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `tgl_publish` timestamp DEFAULT current_timestamp(),
  `tgl_update` timestamp DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `blogs`
--

INSERT INTO `blogs` (`id`, `judul`, `slug`, `isi`, `image`, `is_publish`, `seo_keyword`, `seo_tags`, `seo_deskripsi`, `tgl_publish`, `tgl_update`) VALUES
('e7d4dcdc-576f-451f-98a9-124b876c519a', 'Mutiara', 'mutiara', '<p><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">Sebuah mutiara alami terbentuk tanpa intervensi&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Manusia\" title=\"Manusia\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">manusia</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;sama sekali, di alam liar, dan sangat jarang terjadi. Kira-kira ratusan&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Bivalvia\" title=\"Bivalvia\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">kerang</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;mutiara harus dikumpulkan dan dibuka, dan dengan demikian dibunuh, hanya untuk menemukan satu mutiara liar, dan selama berabad-abad itulah satu-satunya cara untuk memperoleh mutiara. Ini adalah alasan utama mengapa mutiara termasuk sangat berharga pada masa lalu. Mutiara budidaya, di sisi lain, merupakan salah jenis mutiara yang dibentuk dengan melibatkan manusia, di sebuah peternakan mutiara.</span><br></p><p style=\"margin: 0.5em 0px; color: rgb(32, 33, 34); font-family: sans-serif;\">Hampir semua moluska bercangkang bisa menghasilkan beberapa jenis mutiara, melalui proses alami, ketika suatu objek mikroskopis terperangkap di dalam mantel lipatan moluska, tetapi sebagian besar dari mutiara tidak dihargai sebagai batu permata.</p><p style=\"margin: 0.5em 0px; color: rgb(32, 33, 34); font-family: sans-serif;\">Mutiara yang dianggap berkualitas hampir selalu berwarna-warni dan menyerupai&nbsp;<i>mother of pearl</i>, seperti interior kulit yang memproduksi mereka. Namun, hampir semua jenis moluska bercangkang mampu menghasilkan mutiara yang sedikit kurang bersinar atau berbentuk kurang bulat seperti bola. Meskipun mereka mungkin juga sah disebut sebagai \"mutiara\" oleh&nbsp;<a href=\"https://id.wikipedia.org/wiki/Laboratorium\" title=\"Laboratorium\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">laboratorium</a>&nbsp;<a href=\"https://id.wikipedia.org/wiki/Mineralogi#Gemologi\" title=\"Mineralogi\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">gemologi</a>&nbsp;dan juga di bawah aturan&nbsp;<a href=\"https://id.wikipedia.org/wiki/Securities_and_Exchange_Commission\" class=\"mw-redirect\" title=\"Securities and Exchange Commission\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">US Federal Trade Commission</a>&nbsp;dan terbentuk dengan cara yang sama, kebanyakan dari mereka tidak bernilai, kecuali sebagai barang antik.</p><p style=\"margin: 0.5em 0px; color: rgb(32, 33, 34); font-family: sans-serif;\">Mutiara sangat mahal harganya kalau dijual dan dijadikan perhiasan sangat bagus untuk dunia fashion. Mutiara berharga terdapat di alam liar, tetapi dalam kuantitas yang sangat jarang. Mutiara budidaya atau mutiara yang berasal dari&nbsp;<a href=\"https://id.wikipedia.org/wiki/Tiram\" title=\"Tiram\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">tiram</a>&nbsp;merupakan mayoritas dari mutiara-mutiara yang dijual di pasaran. Mutiara&nbsp;<a href=\"https://id.wikipedia.org/wiki/Air_asin\" title=\"Air asin\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">laut</a>&nbsp;dihargai lebih tinggi dari mutiara&nbsp;<a href=\"https://id.wikipedia.org/wiki/Air_tawar\" title=\"Air tawar\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">air tawar</a>. Yang banyak dijual dengan harga murah adalah mutiara&nbsp;<a href=\"https://id.wikipedia.org/wiki/Imitasi\" title=\"Imitasi\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">imitasi</a>, tetapi kualitasnya biasanya jelek. Secara umum, mutiara imitasi dapat dengan mudah dibedakan dari mutiara asli. Mutiara banyak dibudidaya untuk digunakan sebagai&nbsp;<a href=\"https://id.wikipedia.org/wiki/Perhiasan\" title=\"Perhiasan\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">perhiasan</a>. Namun pada masa lalu, mutiara juga digunakan sebagai hiasan pada&nbsp;<a href=\"https://id.wikipedia.org/wiki/Pakaian\" title=\"Pakaian\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">pakaian-pakaian</a>&nbsp;mewah. Mutiara juga bisa dihancurkan dan digunakan dalam&nbsp;<a href=\"https://id.wikipedia.org/wiki/Kosmetik\" title=\"Kosmetik\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">kosmetik</a>,&nbsp;<a href=\"https://id.wikipedia.org/wiki/Obat\" title=\"Obat\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">obat-obatan</a>, atau dalam formula&nbsp;<a href=\"https://id.wikipedia.org/wiki/Cat\" title=\"Cat\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">cat</a>.</p><p><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">Mutiara pertama kali di gunakan oleh raja&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Yunani\" title=\"Yunani\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">Yunani</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Xerxes\" class=\"mw-disambig\" title=\"Xerxes\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">Xerxes</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">. Pada saat itu harga mutiara sama dengan harga kepala manusia. Pada saat itu mutiara dipercaya sebagai jimat yang membawa keberuntungan di saat perang dan juga obat yang dapat menyembuhkan penyakit.</span><sup class=\"noprint Inline-Template\" style=\"color: rgb(32, 33, 34); font-family: sans-serif; line-height: 1; font-size: 11.2px;\"><span title=\"Kalimat yang diikuti tag ini membutuhkan rujukan.\" style=\"white-space: nowrap;\">[<i><a href=\"https://id.wikipedia.org/wiki/Wikipedia:Kutip_sumber_tulisan\" title=\"Wikipedia:Kutip sumber tulisan\" style=\"color: rgb(6, 69, 173); background-image: none; background-position: initial; background-size: initial; background-repeat: initial; background-attachment: initial; background-origin: initial; background-clip: initial;\">butuh rujukan</a></i>]</span></sup></p><p><b style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">Mutiara</b><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;adalah suatu benda keras yang&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Sekresi\" title=\"Sekresi\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">diproduksi</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;di dalam&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Jaringan_lunak\" title=\"Jaringan lunak\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">jaringan lunak</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;(khususnya mantel) dari&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Moluska\" title=\"Moluska\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">moluska</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;hidup. Sama seperti&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Periostrakum\" title=\"Periostrakum\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">cangkang</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">-nya, mutiara terdiri dari&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Kalsium_karbonat\" title=\"Kalsium karbonat\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">kalsium karbonat</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;dalam bentuk&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Kristal\" title=\"Kristal\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">kristal</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;yang telah disimpan dalam lapisan-lapisan konsentris. Mutiara yang ideal adalah yang berbentuk sempurna bulat dan halus, tetapi ada juga berbagai macam bentuk lain. Mutiara alami berkualitas terbaik telah sangat dihargai sebagai&nbsp;</span><a href=\"https://id.wikipedia.org/wiki/Permata\" class=\"mw-redirect\" title=\"Permata\" style=\"font-family: sans-serif; background: none rgb(255, 255, 255); color: rgb(6, 69, 173);\">batu permata</a><span style=\"color: rgb(32, 33, 34); font-family: sans-serif;\">&nbsp;dan objek keindahan selama berabad-abad, dan oleh karena itu, kata \"mutiara\" telah menjadi metafora untuk sesuatu yang sangat langka, baik, mengagumkan, dan berharga.</span></p>', 'public/uploads/blog/e7d4dcdc-576f-451f-98a9-124b876c519a.png', 1, 'apa itu mutiara, mutiaran adalah perhiasan, mutiara cantik', 'mutiara, perhiasan', 'Mutiara adalah suatu benda keras yang diproduksi di dalam jaringan lunak (khususnya mantel) dari moluska hidup. Sam', '2022-02-27 14:41:24', '2022-02-27 14:54:05');

-- --------------------------------------------------------

--
-- Table structure for table `carts`
--

CREATE TABLE `carts` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `produk_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `variasi_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `feedback`
--

CREATE TABLE `feedback` (
  `id` bigint(20) NOT NULL,
  `comment` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `rating` int(1) NOT NULL,
  `produk_id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `kategori`
--

CREATE TABLE `kategori` (
  `id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cover` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `kategori`
--

INSERT INTO `kategori` (`id`, `nama`, `slug`, `cover`, `deleted`, `created_at`, `updated_at`) VALUES
('13fb3640-08d8-4eb1-ad11-487b04e16dbc', 'Rings', 'rings', 'public/uploads/kategori/cover-13fb3640-08d8-4eb1-ad11-487b04e16dbc.', 0, '2022-02-27 14:57:26', '2022-02-27 14:57:26'),
('67775374-b45a-408b-b985-1c7704c0459e', 'Necklae', 'necklae', 'public/uploads/kategori/cover-67775374-b45a-408b-b985-1c7704c0459e.', 0, '2022-02-27 14:57:34', '2022-02-27 14:57:34'),
('6bab247b-1209-437c-b082-61ed5eb818b6', 'Pendant', 'pendant', 'public/uploads/kategori/cover-6bab247b-1209-437c-b082-61ed5eb818b6.', 0, '2022-02-27 14:57:42', '2022-02-27 14:57:42');

-- --------------------------------------------------------

--
-- Table structure for table `mitra`
--

CREATE TABLE `mitra` (
  `id` int(11) NOT NULL,
  `seller_id` varchar(6) COLLATE utf8_unicode_ci NOT NULL,
  `nama` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(60) COLLATE utf8_unicode_ci NOT NULL,
  `alamat` varchar(60) COLLATE utf8_unicode_ci NOT NULL,
  `kota` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  `kota_id` int(11) NOT NULL,
  `kontak` varchar(30) COLLATE utf8_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `mitra`
--

INSERT INTO `mitra` (`id`, `seller_id`, `nama`, `email`, `alamat`, `kota`, `kota_id`, `kontak`) VALUES
(1, 'dg-001', 'John Doe', 'johndoe@gmail.com', 'Jl. Ir. H Juanda No.9B', 'Bandung', 23, '081-111-111-111');

-- --------------------------------------------------------

--
-- Table structure for table `order_bayar`
--

CREATE TABLE `order_bayar` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `order_checkout_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `jumlah` decimal(12,2) NOT NULL,
  `status` enum('belum lunas','pending','capture','settlement','deny','cancel','expire','failure','chargeback','refund','partial_refund','partial_chargeback','authorize') DEFAULT 'pending',
  `token` varchar(100) NOT NULL,
  `redirect_url` text NOT NULL,
  `tgl_bayar` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `order_checkout`
--

CREATE TABLE `order_checkout` (
  `id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('dipesan','dikonfirmasi','dikirim','selesai','batal','dikemas') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'dipesan',
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `alamat_id` bigint(20) NOT NULL,
  `kurir` enum('jne','pos','tiki','') COLLATE utf8mb4_unicode_ci NOT NULL,
  `paket_kurir` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ongkir` int(11) NOT NULL,
  `no_resi` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `invoice_no` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tgl_order` timestamp DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `order_item`
--

CREATE TABLE `order_item` (
  `id` bigint(20) NOT NULL,
  `order_checkout_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `produk_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `variasi_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `password_resets`
--

CREATE TABLE `password_resets` (
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `produk`
--

CREATE TABLE `produk` (
  `id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deskripsi` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `spesifikasi` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`spesifikasi`)),
  `berat` int(11) DEFAULT NULL,
  `satuan_berat` enum('g','kg') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lebar` float DEFAULT NULL,
  `panjang` float DEFAULT NULL,
  `tinggi` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `harga` int(11) DEFAULT NULL,
  `kode` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `stok` int(11) DEFAULT NULL,
  `jumlah_penjualan` int(11) NOT NULL DEFAULT 0,
  `dilihat` bigint(20) DEFAULT 0,
  `is_has_variant` tinyint(4) DEFAULT 0,
  `deleted` tinyint(1) NOT NULL DEFAULT 0,
  `discount` int(3) NOT NULL,
  `kategori_id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `produk`
--

INSERT INTO `produk` (`id`, `nama`, `slug`, `deskripsi`, `spesifikasi`, `berat`, `satuan_berat`, `lebar`, `panjang`, `tinggi`, `harga`, `kode`, `stok`, `jumlah_penjualan`, `dilihat`, `is_has_variant`, `deleted`, `discount`, `kategori_id`, `created_at`, `updated_at`) VALUES
('d6585406-1260-430c-8017-8b7e4d2495d1', 'Cincin Mutiara Lombok', 'cincin-mutiara-lombok', '<p><span style=\"color: rgba(49, 53, 59, 0.96); font-family: &quot;open sans&quot;, tahoma, sans-serif;\">Cincin perak s925 celup emas putih dan mutiara laut asli. Mutiara: sputh sea pearl/ mutiara laut Besar mutiara: 9mm Berat mutiara: 1gr Warna mutiara: putih Rangka: solid s925 silver lapis white gold Cincin sangat berkilau dan tidak menghitam karna dilapisi dengan yellow gold asli. Semua product kami special design jadi stock barang sangat limited. Harga termasuk box. Warna, size, dan grade mutiara bisa di ganti sesuai permintaan. *kami adalah pearl wholesaler berbasis di lombok, NTB. Cek website kami di www•nhpearl•com atau bisa email kami di info@nhpearl•com</span><br></p>', '[{\"nama\":\"Kondisi\",\"value\":\"Baru\"},{\"nama\":\"Berat\",\"value\":\"75gr\"},{\"nama\":\"Bahan Cincin\",\"value\":\"Emas 24carat\"},{\"nama\":\"Mutiara\",\"value\":\"Mutiara Lombok\"}]', 75, 'g', 0, 0, '0', 450000, 'CM-01-L', 18, 2, 0, 0, 0, 20, '13fb3640-08d8-4eb1-ad11-487b04e16dbc', '2022-02-27 15:00:20', '2022-02-27 15:00:20');

-- --------------------------------------------------------

--
-- Table structure for table `produk_foto`
--

CREATE TABLE `produk_foto` (
  `id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `produk_id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `path` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_utama` tinyint(4) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `produk_foto`
--

INSERT INTO `produk_foto` (`id`, `produk_id`, `path`, `is_utama`, `created_at`, `updated_at`) VALUES
('b0fd99f5-dbce-473e-acdc-8ada02971410', 'd6585406-1260-430c-8017-8b7e4d2495d1', '/uploads/produk/d6585406-1260-430c-8017-8b7e4d2495d1/foto-produk-2.png', 0, '2022-02-27 15:00:20', '2022-02-27 15:00:20'),
('f8505ce4-a898-4881-96e0-b121b1cdfdbb', 'd6585406-1260-430c-8017-8b7e4d2495d1', '/uploads/produk/d6585406-1260-430c-8017-8b7e4d2495d1/foto-produk-1.png', 1, '2022-02-27 15:00:20', '2022-02-27 15:00:20');

-- --------------------------------------------------------

--
-- Table structure for table `produk_variasi`
--

CREATE TABLE `produk_variasi` (
  `id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `variant` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sku` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `harga` decimal(10,2) NOT NULL,
  `produk_id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `stok` int(11) DEFAULT 0,
  `deleted` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `promo`
--

CREATE TABLE `promo` (
  `id` varchar(191) NOT NULL,
  `judul` varchar(191) NOT NULL,
  `slug` varchar(191) NOT NULL,
  `deskripsi` longtext DEFAULT NULL,
  `image` varchar(191) DEFAULT NULL,
  `is_publish` tinyint(4) NOT NULL DEFAULT 0,
  `is_featured` tinyint(4) NOT NULL DEFAULT 0,
  `seo_keyword` text NOT NULL,
  `seo_tags` text NOT NULL,
  `seo_deskripsi` text NOT NULL,
  `tgl_mulai` timestamp DEFAULT current_timestamp(),
  `tgl_selesai` timestamp DEFAULT current_timestamp(),
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `rekening`
--

CREATE TABLE `rekening` (
  `id` bigint(20) NOT NULL,
  `bank_id` bigint(20) NOT NULL,
  `rekening_no` varchar(30) NOT NULL,
  `nama` varchar(191) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `session`
--

CREATE TABLE `session` (
  `k` varchar(64) NOT NULL DEFAULT '',
  `v` blob NOT NULL,
  `e` bigint(20) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `session`
--

INSERT INTO `session` (`k`, `v`, `e`) VALUES
('be6bc4e1-461e-421c-8529-a4eae6418121', 0x0eff81040102ff8200010c01100000fe010dff82000105746f6b656e06737472696e670cfff900fff64265617265722065794a68624763694f694a49557a49314e694973496e523563434936496b705856434a392e65794a6c6257467062434936496e4e766247467961584e736157646f644467775147647459576c734c6d4e7662534973496d563463434936496a49774d6a49744d4451744d4446554d6a41364d5455364d7a51754d5445794f5445354e7973774e7a6f774d434973496d6c6b496a6f774c434a70633139325a584a705a6d6c6c5a43493664484a315a53776964486c775a534936496e567a5a58496966512e35427656477065376b7356616c557155766b64514e3164767244342d623165344a30376c6b306f62587859, 1646831734);

-- --------------------------------------------------------

--
-- Table structure for table `slider`
--

CREATE TABLE `slider` (
  `id` varchar(100) NOT NULL,
  `name` varchar(191) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `image` varchar(191) NOT NULL,
  `is_publish` tinyint(1) NOT NULL,
  `url` varchar(191) NOT NULL,
  `type` int(11) DEFAULT NULL,
  `created_at` timestamp DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `slider`
--

INSERT INTO `slider` (`id`, `name`, `description`, `image`, `is_publish`, `url`, `type`, `created_at`, `updated_at`) VALUES
('00bca9c4-dfcf-4de2-8d18-8008c8bc4e21', 'Software Engineer Mindmap', '<div class=\"slide-block\" style=\"font-family: SourceSansPro, sans-serif; color: rgb(17, 25, 64); padding: 0px 32.2969px;\"><div class=\"slide-des\" style=\"margin-bottom: 24px; margin-top: 16px; line-height: 24px;\"><div style=\"margin: 0px;\"><p align=\"left\" style=\"margin-right: 0px; margin-bottom: 15px; margin-left: 0px;\">The sad reality about new IT technologies is that security is often an afterthought. Cloud computing came along and people jumped onboard before a robust and well-planned security roadmap could be established. Because of this, many early cloud adopters are scrambling to re-architect their cloud services with advanced security. While cloud security essentially uses the same tools found in traditional infrastructure security, there are more things to consider. Security considerations ranging from third-party data storage, data access, and even&nbsp;<a href=\"http://searchcloudsecurity.techtarget.com/tip/Securing-a-multi-tenant-environment\" target=\"_blank\" style=\"color: rgb(17, 25, 64); margin: 0px;\">multi-tenancy</a>&nbsp;issues are new skills you can acquire.</p><p align=\"left\" style=\"margin-right: 0px; margin-bottom: 15px; margin-left: 0px;\">(Image:&nbsp;<a href=\"https://pixabay.com/en/castle-privacy-policy-security-538722/\" target=\"_blank\" style=\"color: rgb(17, 25, 64); margin: 0px;\">Geralt</a>&nbsp;via Pixabay)</p></div></div></div><div class=\"article-content slide-content\" style=\"font-family: SourceSansPro, sans-serif; color: rgb(17, 25, 64);\"><div class=\"\" data-name=\"Article Video Ad - video_v\" data-size=\"1*1\" data-targeting=\"pos=video_v\" id=\"7_8uno8w4h4ia000008uno8w4h4ia000008un\" style=\"font-size: 16px; line-height: 27px; word-break: break-word;\"></div><div style=\"font-size: 16px; line-height: 27px; word-break: break-word;\"><p align=\"left\" style=\"margin-bottom: 1rem; line-height: 27px; word-break: break-word; padding: 0px 32.2969px;\">Generally speaking, IT specializations are changing at a pace that\'s faster than ever before. Everyone -- from desktop support to developers and database administrators -- has had to adjust and to acquire new skills to keep up with business demands. But one area of IT that has historically been slower to adopt new technologies and skills was IT infrastructure. It used to be that for applications, data, and processes to change, they had to change within the rigid constraints of the IT infrastructure. Infrastructures such as private data centers were big, expensive, and incredibly inflexible. All of that is changing now.</p><div class=\"rectangle-ads\" style=\"line-height: 27px; word-break: break-word; float: right; margin: 1rem 0px 1rem 1rem;\"><div class=\"\" data-name=\"Rectangle_300_1v_article\" data-size=\"300*250,300*600\" data-targeting=\"pos=300_1v_article\" id=\"8_hl0gm0hisi800000hl0gm0hisi800000hl0\" style=\"line-height: 27px; word-break: break-word;\"></div></div><p align=\"left\" style=\"margin-bottom: 1rem; line-height: 27px; word-break: break-word; padding: 0px 32.2969px;\">Server and storage virtualization was the first true shift in IT infrastructure over the past decade that caused a dramatic shift in IT skillsets. The virtualization movement affected not only server and storage administrators, but also network engineers. Networks, server hardware, and virtualization platforms were suddenly designed and deployed to be more flexible than ever before. Brand-new server instances could be spun up in a matter of minutes -- as opposed to the traditional days or weeks.</p></div></div>', 'public/uploads/slider/cover-00bca9c4-dfcf-4de2-8d18-8008c8bc4e21.png', 1, '', 2, '2022-02-27 16:31:57', '2022-02-27 16:31:57');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `hp` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tgl_lahir` timestamp DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_verified` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `hp`, `email`, `password`, `tgl_lahir`, `is_verified`, `created_at`, `updated_at`) VALUES
(3, 'user01', '081-111-111-22', 'solarislight80@gmail.com', '$2a$10$sUtf7m3Tobq585wkl1AUauwCScD0JE2lsRg53M5nhS.CzF72PYdB6', '2022-03-02 13:15:34', 1, '2022-03-02 13:14:25', '2022-03-02 13:14:25');

-- --------------------------------------------------------

--
-- Table structure for table `verifikasi`
--

CREATE TABLE `verifikasi` (
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `kode` varchar(200) COLLATE utf8_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `admins_username_unique` (`username`),
  ADD UNIQUE KEY `admins_email_unique` (`email`);

--
-- Indexes for table `alamat`
--
ALTER TABLE `alamat`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `alamat_origin`
--
ALTER TABLE `alamat_origin`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `bank`
--
ALTER TABLE `bank`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `vendor` (`vendor`);

--
-- Indexes for table `blogs`
--
ALTER TABLE `blogs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `judul` (`judul`);

--
-- Indexes for table `carts`
--
ALTER TABLE `carts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_chart_foreign` (`user_id`),
  ADD KEY `produk_id` (`produk_id`),
  ADD KEY `variasi_id` (`variasi_id`);

--
-- Indexes for table `feedback`
--
ALTER TABLE `feedback`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `produk_id` (`produk_id`);

--
-- Indexes for table `kategori`
--
ALTER TABLE `kategori`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `mitra`
--
ALTER TABLE `mitra`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `seller_id` (`seller_id`);

--
-- Indexes for table `order_bayar`
--
ALTER TABLE `order_bayar`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `order_checkout_id` (`order_checkout_id`),
  ADD UNIQUE KEY `token` (`token`);

--
-- Indexes for table `order_checkout`
--
ALTER TABLE `order_checkout`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `invoice_no` (`invoice_no`),
  ADD KEY `alamat_id` (`alamat_id`);

--
-- Indexes for table `order_item`
--
ALTER TABLE `order_item`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order_detail_foreign` (`order_checkout_id`);

--
-- Indexes for table `password_resets`
--
ALTER TABLE `password_resets`
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `password_resets_email_index` (`email`);

--
-- Indexes for table `produk`
--
ALTER TABLE `produk`
  ADD PRIMARY KEY (`id`),
  ADD KEY `kategori_id` (`kategori_id`);

--
-- Indexes for table `produk_foto`
--
ALTER TABLE `produk_foto`
  ADD PRIMARY KEY (`id`),
  ADD KEY `produk_foto-foreign` (`produk_id`);

--
-- Indexes for table `produk_variasi`
--
ALTER TABLE `produk_variasi`
  ADD PRIMARY KEY (`id`),
  ADD KEY `produk_variasi_produk_id_foreign` (`produk_id`);

--
-- Indexes for table `promo`
--
ALTER TABLE `promo`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `rekening`
--
ALTER TABLE `rekening`
  ADD PRIMARY KEY (`id`),
  ADD KEY `bank_id` (`bank_id`);

--
-- Indexes for table `session`
--
ALTER TABLE `session`
  ADD PRIMARY KEY (`k`);

--
-- Indexes for table `slider`
--
ALTER TABLE `slider`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_email_unique` (`email`),
  ADD KEY `id` (`id`);

--
-- Indexes for table `verifikasi`
--
ALTER TABLE `verifikasi`
  ADD UNIQUE KEY `user_id` (`user_id`),
  ADD KEY `user_id_2` (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admins`
--
ALTER TABLE `admins`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `alamat`
--
ALTER TABLE `alamat`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `alamat_origin`
--
ALTER TABLE `alamat_origin`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `bank`
--
ALTER TABLE `bank`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `carts`
--
ALTER TABLE `carts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `feedback`
--
ALTER TABLE `feedback`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `mitra`
--
ALTER TABLE `mitra`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `order_bayar`
--
ALTER TABLE `order_bayar`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `order_item`
--
ALTER TABLE `order_item`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `rekening`
--
ALTER TABLE `rekening`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `alamat`
--
ALTER TABLE `alamat`
  ADD CONSTRAINT `alamat_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `carts`
--
ALTER TABLE `carts`
  ADD CONSTRAINT `carts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `carts_ibfk_2` FOREIGN KEY (`produk_id`) REFERENCES `produk` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `feedback`
--
ALTER TABLE `feedback`
  ADD CONSTRAINT `feedback_ibfk_1` FOREIGN KEY (`produk_id`) REFERENCES `produk` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `feedback_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `order_bayar`
--
ALTER TABLE `order_bayar`
  ADD CONSTRAINT `order_bayar_ibfk_1` FOREIGN KEY (`order_checkout_id`) REFERENCES `order_checkout` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `order_checkout`
--
ALTER TABLE `order_checkout`
  ADD CONSTRAINT `order_checkout_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `order_checkout_ibfk_2` FOREIGN KEY (`alamat_id`) REFERENCES `alamat` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `order_item`
--
ALTER TABLE `order_item`
  ADD CONSTRAINT `order_item_ibfk_1` FOREIGN KEY (`order_checkout_id`) REFERENCES `order_checkout` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `password_resets`
--
ALTER TABLE `password_resets`
  ADD CONSTRAINT `password_resets_ibfk_1` FOREIGN KEY (`email`) REFERENCES `users` (`email`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `produk`
--
ALTER TABLE `produk`
  ADD CONSTRAINT `produk_ibfk_1` FOREIGN KEY (`kategori_id`) REFERENCES `kategori` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `produk_foto`
--
ALTER TABLE `produk_foto`
  ADD CONSTRAINT `produk_foto_foreign` FOREIGN KEY (`produk_id`) REFERENCES `produk` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `produk_variasi`
--
ALTER TABLE `produk_variasi`
  ADD CONSTRAINT `variasi_produk` FOREIGN KEY (`produk_id`) REFERENCES `produk` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `rekening`
--
ALTER TABLE `rekening`
  ADD CONSTRAINT `rekening_ibfk_1` FOREIGN KEY (`bank_id`) REFERENCES `bank` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `verifikasi`
--
ALTER TABLE `verifikasi`
  ADD CONSTRAINT `verifikasi_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
