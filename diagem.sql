-- phpMyAdmin SQL Dump
-- version 5.1.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Feb 08, 2022 at 01:17 PM
-- Server version: 10.6.5-MariaDB
-- PHP Version: 8.1.2

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
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  `email_verified_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `admins`
--

INSERT INTO `admins` (`id`, `nama`, `username`, `email`, `password`, `roles`, `created_at`, `updated_at`, `email_verified_at`) VALUES
(1, 'Super Admin', 'super admin', 'superadmin@admin.com', '$2a$12$EEdZRYoYV.OmKsGeRTOmwOthsXrLqTcLwsvcGL0gsC/7yVjsxH7Ui\r\n', 'super admin', '2020-05-18 14:12:05', '2020-05-18 14:12:05', '2021-10-08 20:25:13'),
(9, 'shania yan', 'shaniayan', 'shaniayan@gmail.com', '$2a$10$LGqZRwVDN544SstqmmrbbunOCsrAF31EHDKEAP6hFemr355QI9PYa', 'marketing', '2022-02-05 15:07:29', '2022-02-05 15:08:18', '2022-02-05 15:07:29'),
(10, 'Bradley Hermeso', 'brandleyhermes', 'bradleyhermeso@gmail.com', '$2a$10$K0g2Eel1xTb0GfaQkNniAORia3EmxN5A/GkHmSd7J6CY9RK1N.Gte', 'admin', '2022-02-05 15:09:27', '2022-02-05 15:09:27', '2022-02-05 15:09:27');

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

--
-- Dumping data for table `alamat`
--

INSERT INTO `alamat` (`id`, `user_id`, `alamat`, `nama`, `penerima`, `phone`, `provinsi_id`, `kota_id`, `nama_provinsi`, `nama_kota`, `kd_pos`, `is_utama`, `created_at`, `updated_at`) VALUES
(26, 24, 'Jl. Tubagus Ismail Dalam Gg.1 No.153/A Coblong', 'Kostan', 'Johnny Walkers', '089-992-993-445', '9', '22', 'Jawa Barat', 'Bandung', '40311', 0, '2021-12-20 21:09:38', '2021-12-20 21:09:38'),
(27, 24, 'Jl.Otto Iskandar Dinata Gg.08 No.05', 'Rumah Istri', 'F. Walkers D', '081-121-032-556', '9', '55', 'Jawa Barat', 'Bekasi', '17121', 1, '2021-12-20 21:11:32', '2021-12-20 21:11:32'),
(33, 26, 'Jl.Otto Iskandar Dinata No.05', 'kantor', 'user 02', '089-123-432-555', '9', '107', 'Jawa Barat', 'Cimahi', '40512', 0, '2021-12-29 02:16:23', '2021-12-29 02:16:23'),
(34, 27, 'Jl. Tubagus Ismail Dalam Gg.1 No.153/A Coblong Bandung Jawa Barat', 'Kosan', 'Nadzar Mutaqin', '081-111-222-333', '9', '23', 'Jawa Barat', 'Bandung', '40111', 1, '2022-02-04 12:49:52', '2022-02-04 12:49:52'),
(35, 27, 'Jl. Godong Cluster SW No.A77 Garut', 'Rumah', 'Nadzar Mutaqin', '081-111-222-333', '9', '126', 'Jawa Barat', 'Garut', '44126', 0, '2022-02-04 12:51:08', '2022-02-04 12:51:08');

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
(136, '667770001', 'bca_va', 1),
(137, '44353444444', 'bni_va', 1),
(138, '556655355344', 'bri_va', 1);

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
  `tgl_publish` date NOT NULL DEFAULT current_timestamp(),
  `tgl_update` date NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `blogs`
--

INSERT INTO `blogs` (`id`, `judul`, `slug`, `isi`, `image`, `is_publish`, `seo_keyword`, `seo_tags`, `seo_deskripsi`, `tgl_publish`, `tgl_update`) VALUES
('b2aa058d-7d5d-4f2d-8bff-060af8568d1c', 'blog 2 tentang', 'blog-2-tentang', '<pre class=\"_3F2CP _3hukP\" style=\"box-sizing: inherit; font-family: &quot;Roboto Mono&quot;, &quot;Courier New&quot;, monospace; line-height: 1.4em; margin-bottom: 0px; background: rgb(248, 248, 248); color: rgb(33, 33, 33); font-size: 13px;\"><span class=\"_2jIGi\" style=\"box-sizing: inherit; display: block;\"><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">Sure it’s a calming notion, perpetual in motion\n</span></span><span class=\"_2jIGi\" style=\"box-sizing: inherit; display: block;\"><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">            <span class=\"_3PpPJ OrSDI\" data-name=\"Bm\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">Bm</span>                  <span class=\"_3PpPJ OrSDI\" data-name=\"C#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">C#m</span> <span class=\"_3PpPJ OrSDI\" data-name=\"F#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">F#m</span>\n</span><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">But I don’t need the comfort of any lies\n</span></span><span class=\"_2jIGi\" style=\"box-sizing: inherit; display: block;\"><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">           <span class=\"_3PpPJ OrSDI\" data-name=\"Bm\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">Bm</span>                  <span class=\"_3PpPJ OrSDI\" data-name=\"C#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">C#m</span>      <span class=\"_3PpPJ OrSDI\" data-name=\"F#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">F#m</span>          <span class=\"_3PpPJ OrSDI\" data-name=\"A\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">A</span> <span class=\"_3PpPJ OrSDI\" data-name=\"F\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">F</span> <span class=\"_3PpPJ OrSDI\" data-name=\"C#\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">C#</span>\n</span><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">For I have seen the ending and there is no ascending rise\n</span></span><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\"> \n</span><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">[Chorus 1]\n</span><span class=\"_2jIGi\" style=\"box-sizing: inherit; display: block;\"><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">             <span class=\"_3PpPJ OrSDI\" data-name=\"Bm\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">Bm</span>             <span class=\"_3PpPJ OrSDI\" data-name=\"C#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">C#m</span>         <span class=\"_3PpPJ OrSDI\" data-name=\"F#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">F#m</span>\n</span><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">Oh back when I was younger, was told by other youngsters\n</span></span><span class=\"_2jIGi\" style=\"box-sizing: inherit; display: block;\"><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">            <span class=\"_3PpPJ OrSDI\" data-name=\"Bm\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">Bm</span>              <span class=\"_3PpPJ OrSDI\" data-name=\"C#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">C#m</span>         <span class=\"_3PpPJ OrSDI\" data-name=\"F#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">F#m</span>\n</span><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">That my end will be torture beneath the earth\n</span></span><span class=\"_2jIGi\" style=\"box-sizing: inherit; display: block;\"><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">             <span class=\"_3PpPJ OrSDI\" data-name=\"Bm\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">Bm</span>                     <span class=\"_3PpPJ OrSDI\" data-name=\"C#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">C#m</span>      <span class=\"_3PpPJ OrSDI\" data-name=\"F#m\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">F#m</span>\n</span><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">‘cuz I don’t see what they see when death is staring at me\n</span></span><span class=\"_2jIGi\" style=\"box-sizing: inherit; display: block;\"><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">        <span class=\"_3PpPJ OrSDI\" data-name=\"Bm\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">Bm</span>           <span class=\"_3PpPJ OrSDI\" data-name=\"D\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">D</span>*           <span class=\"_3PpPJ OrSDI\" data-name=\"E\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">E</span>             <span class=\"_3PpPJ OrSDI\" data-name=\"F\" style=\"box-sizing: inherit; font-weight: 700; cursor: default; display: inline-block; padding: 0.1em 0.5em 0.15em; margin: -0.1em -0.5em -0.15em; border-radius: 2px; position: relative; color: rgb(0, 0, 0);\">F</span>\n</span><span class=\"_3rlxz\" style=\"box-sizing: inherit; display: block; min-height: 1.4em;\">I see a window, a limit, to live it, or not at all</span></span></pre>', 'public/uploads/blog/b2aa058d-7d5d-4f2d-8bff-060af8568d1c.jpeg', 1, 'chrod gitar ', 'gitar, chord, rare occasion', 'deskripsi SEO', '2022-01-05', '2022-01-05'),
('d939479e-c8d3-4b0c-a068-d22df2ad475c', 'juduk bkig', 'juduk-bkig', '<p>isi blog</p>', 'public/uploads/blog/d939479e-c8d3-4b0c-a068-d22df2ad475c.jpeg', 1, 'keyword', 'tags, tags, tags', 'deskripsi seo', '2022-01-05', '2022-01-05');

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

--
-- Dumping data for table `carts`
--

INSERT INTO `carts` (`id`, `user_id`, `produk_id`, `variasi_id`, `quantity`) VALUES
(55, 24, '5f33071d-8885-47ad-88bf-a3a736ec9b03', 'a8cb0287-e205-4290-98da-9cf9e8463ca6', 2),
(56, 24, 'e5cfeec1-a682-4236-9270-addb8b031f61', '', 3),
(57, 24, 'd45078c4-b18a-48b4-b82f-ce6a24a09ecf', '', 4);

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
  `created_at` date NOT NULL DEFAULT current_timestamp()
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
('5dbe63ef-15a8-4074-a899-7f636556462d', 'test', 'test', 'public/uploads/kategori/cover-5dbe63ef-15a8-4074-a899-7f636556462d.jpeg', 1, '2021-12-02 04:24:12', '2021-12-02 04:24:12'),
('72b43b5b-3b56-432c-9b1b-942ea708f6bf', 'Rings', 'rings', 'public/uploads/kategori/cover-72b43b5b-3b56-432c-9b1b-942ea708f6bf.jpeg', 0, '2021-12-02 04:11:29', '2021-12-02 04:11:29'),
('a5dfb114-4680-42b5-bd19-9eb6aa63f2e4', 'Necklace', 'necklace', 'public/uploads/kategori/cover-a5dfb114-4680-42b5-bd19-9eb6aa63f2e4.jpeg', 0, '2021-12-03 07:33:30', '2021-12-03 07:33:30'),
('bea1fc4d-6dc9-4e5f-b7a3-504b617a1436', 'Pendant', 'pendant', 'public/uploads/kategori/cover-bea1fc4d-6dc9-4e5f-b7a3-504b617a1436.jpeg', 0, '2021-12-03 07:34:02', '2021-12-03 07:34:02'),
('ec74b172-281b-4dae-8494-d98e73f6f906', 'Earrings', 'earrings', 'public/uploads/kategori/cover-ec74b172-281b-4dae-8494-d98e73f6f906.jpeg', 0, '2021-12-03 07:34:58', '2021-12-03 07:34:58');

-- --------------------------------------------------------

--
-- Table structure for table `mitra`
--

CREATE TABLE `mitra` (
  `id` int(11) NOT NULL,
  `seller_id` varchar(6) COLLATE utf8mb3_unicode_ci NOT NULL,
  `nama` varchar(30) COLLATE utf8mb3_unicode_ci NOT NULL,
  `email` varchar(60) COLLATE utf8mb3_unicode_ci NOT NULL,
  `alamat` varchar(60) COLLATE utf8mb3_unicode_ci NOT NULL,
  `kota` varchar(30) COLLATE utf8mb3_unicode_ci NOT NULL,
  `kota_id` int(11) NOT NULL,
  `kontak` varchar(30) COLLATE utf8mb3_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_unicode_ci;

--
-- Dumping data for table `mitra`
--

INSERT INTO `mitra` (`id`, `seller_id`, `nama`, `email`, `alamat`, `kota`, `kota_id`, `kontak`) VALUES
(18, 'DGM001', 'John Doe', 'johndoe@gmail.com', 'Jl.A No.1', 'Sukabumi', 431, '081-111-111-112'),
(19, 'DGM002', 'Jane Doe', 'janedoe@gmail.com', 'Jl. A Kec.B', 'Bandung Barat', 24, '081-000-000-002'),
(24, 'DGM003', 'johan Doe', 'johandoe@gmail.com', 'Jl.A Kec.B', 'Bandung Barat', 24, '081-xxx-xxx-xxx'),
(25, 'DGM004', 'Jenita Doe', 'jenitadoe@gmail.com', 'Jl.A Gang.B', 'Bandung', 23, '081-xxx-xxx-xxx');

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

--
-- Dumping data for table `order_bayar`
--

INSERT INTO `order_bayar` (`id`, `order_checkout_id`, `jumlah`, `status`, `token`, `redirect_url`, `tgl_bayar`) VALUES
(39, 'e7e4bbcb-1a35-411f-9c5b-fb7ef6cb0f4d', '4060000.00', 'settlement', '8ef9a1a6-7385-4a82-8809-343ea4ff669a', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/8ef9a1a6-7385-4a82-8809-343ea4ff669a', '2022-02-04 13:05:14'),
(40, 'af209daf-eb87-4205-9eda-5a3c6294b8bd', '15168.00', 'settlement', 'a9ed2705-0270-41f9-af5e-8120752a1baf', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/a9ed2705-0270-41f9-af5e-8120752a1baf', '2022-02-08 12:56:31');

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
  `tgl_order` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `order_checkout`
--

INSERT INTO `order_checkout` (`id`, `status`, `user_id`, `alamat_id`, `kurir`, `paket_kurir`, `ongkir`, `no_resi`, `invoice_no`, `tgl_order`) VALUES
('af209daf-eb87-4205-9eda-5a3c6294b8bd', 'dipesan', 27, 35, 'jne', 'REG', 15000, NULL, 'DGM-2556-822022', '2022-02-08 12:55:52'),
('e7e4bbcb-1a35-411f-9c5b-fb7ef6cb0f4d', 'dikirim', 27, 34, 'jne', 'OKE', 10000, '6666', 'DGM-7903-422022', '2022-02-04 12:54:02');

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

--
-- Dumping data for table `order_item`
--

INSERT INTO `order_item` (`id`, `order_checkout_id`, `produk_id`, `variasi_id`, `quantity`) VALUES
(500000038, 'e7e4bbcb-1a35-411f-9c5b-fb7ef6cb0f4d', '7e6cd294-00a7-437c-b687-495097bc95d8', '', 1),
(500000039, 'af209daf-eb87-4205-9eda-5a3c6294b8bd', 'dca07151-a309-4029-9ef2-c7da4b7c0518', '978b591b-6570-41d3-9acc-fc1d25070e47', 2);

-- --------------------------------------------------------

--
-- Table structure for table `password_resets`
--

CREATE TABLE `password_resets` (
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
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
('5f33071d-8885-47ad-88bf-a3a736ec9b03', 'Test Produk 4', 'test-produk-4', '<p>ini adalah testing produk ke empat (dimana bervariasi tapi tidak ada discount)</p>', '[{\"nama\":\"bahan \",\"value\":\"berlian\"}]', 450, 'g', 45, 45, '45', 0, 'TP-004', 0, 0, 0, 1, 0, 0, '72b43b5b-3b56-432c-9b1b-942ea708f6bf', '2022-01-07 12:59:28', '2022-01-07 12:59:28'),
('7e6cd294-00a7-437c-b687-495097bc95d8', 'One Ring ', 'one-ring', '<p>jangan pakai cincin ini kalau Anda tidak mau jadi gollum</p>', '[{\"nama\":\"Bahan\",\"value\":\"Emas\"},{\"nama\":\"Kemurnian \",\"value\":\"30 Carrat\"},{\"nama\":\"Efek \",\"value\":\"Kegilaan\"}]', 25, 'g', 15, 15, '15', 4500000, 'LOTR-001', 4, 0, 0, 0, 0, 10, '72b43b5b-3b56-432c-9b1b-942ea708f6bf', '2022-02-03 11:57:52', '2022-02-03 11:57:52'),
('ab4cae25-4308-4092-bfa2-eea8c105f253', 'test', 'test', '<p>test</p>', '[{\"nama\":\"test\",\"value\":\"test\"},{\"nama\":\"test\",\"value\":\"test\"}]', 560, 'g', 12, 12, '12', 120, 'test', 1, 0, 0, 0, 0, 6, '72b43b5b-3b56-432c-9b1b-942ea708f6bf', '2022-02-08 13:08:27', '2022-02-08 13:08:27'),
('d45078c4-b18a-48b4-b82f-ce6a24a09ecf', 'Test Produk 2', 'test-produk-2', '<p>ini adalah test produk</p>', '[{\"nama\":\"Bahan\",\"value\":\"emas\"}]', 125, 'g', 100, 100, '100', 120, 'TP-002', 12, 0, 0, 0, 0, 0, '72b43b5b-3b56-432c-9b1b-942ea708f6bf', '2022-01-06 12:10:20', '2022-01-06 12:10:20'),
('dca07151-a309-4029-9ef2-c7da4b7c0518', 'Test Produk', 'test-produk', '<p>ini adalah test produk&nbsp;</p>', '[{\"nama\":\"Bahan\",\"value\":\"Emas\"},{\"nama\":\"Kemurnian\",\"value\":\"24 Carat\"}]', 250, 'g', 120, 120, '120', 0, 'TP002', 0, 0, 0, 1, 0, 30, 'a5dfb114-4680-42b5-bd19-9eb6aa63f2e4', '2022-01-06 11:54:03', '2022-01-06 11:54:03'),
('e5cfeec1-a682-4236-9270-addb8b031f61', 'Produk testing 3', 'produk-testing-3', '<p>ini adalah produk testing ke 3</p>', '[{\"nama\":\"Bahan\",\"value\":\"Emas\"}]', 120, 'g', 150, 150, '150', 150, 'TP-003', 15, 0, 0, 0, 0, 40, 'ec74b172-281b-4dae-8494-d98e73f6f906', '2022-01-07 12:57:01', '2022-01-07 12:57:01');

-- --------------------------------------------------------

--
-- Table structure for table `produk_foto`
--

CREATE TABLE `produk_foto` (
  `id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `produk_id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `path` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_utama` tinyint(4) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `produk_foto`
--

INSERT INTO `produk_foto` (`id`, `produk_id`, `path`, `is_utama`, `created_at`, `updated_at`) VALUES
('21eb1cf6-a9fc-4384-9059-dc337352e333', 'e5cfeec1-a682-4236-9270-addb8b031f61', '/uploads/produk/e5cfeec1-a682-4236-9270-addb8b031f61/foto-produk-1.png', 1, NULL, NULL),
('2ac9e12b-9ab5-4bbc-8e01-97a18866c9ae', '5f33071d-8885-47ad-88bf-a3a736ec9b03', '/uploads/produk/5f33071d-8885-47ad-88bf-a3a736ec9b03/foto-produk-1.png', 1, NULL, NULL),
('2d319a39-8c75-4688-9bfb-82a62ab344d8', 'ab4cae25-4308-4092-bfa2-eea8c105f253', '/uploads/produk/ab4cae25-4308-4092-bfa2-eea8c105f253/foto-produk-2.png', 0, NULL, NULL),
('39cde0c9-fb23-4eda-82ec-2c3fd72bb9c8', '7e6cd294-00a7-437c-b687-495097bc95d8', '/uploads/produk/7e6cd294-00a7-437c-b687-495097bc95d8/foto-produk-1.png', 1, NULL, NULL),
('514d9622-e57e-49e0-8255-bd635d9f151d', 'd45078c4-b18a-48b4-b82f-ce6a24a09ecf', '/uploads/produk/d45078c4-b18a-48b4-b82f-ce6a24a09ecf/foto-produk-1.png', 1, NULL, NULL),
('b8c56b9b-e55c-4ef7-bfd3-6e55ccf38956', 'dca07151-a309-4029-9ef2-c7da4b7c0518', '/uploads/produk/dca07151-a309-4029-9ef2-c7da4b7c0518/foto-produk-2.png', 0, NULL, NULL),
('bd457549-b284-41ef-a2cc-7499702978f7', 'ab4cae25-4308-4092-bfa2-eea8c105f253', '/uploads/produk/ab4cae25-4308-4092-bfa2-eea8c105f253/foto-produk-1.png', 1, NULL, NULL),
('f71c8cf5-46b5-4d14-a868-ba70f89c2665', 'dca07151-a309-4029-9ef2-c7da4b7c0518', '/uploads/produk/dca07151-a309-4029-9ef2-c7da4b7c0518/foto-produk-1.png', 1, NULL, NULL);

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

--
-- Dumping data for table `produk_variasi`
--

INSERT INTO `produk_variasi` (`id`, `variant`, `sku`, `harga`, `produk_id`, `stok`, `deleted`, `created_at`, `updated_at`) VALUES
('978b591b-6570-41d3-9acc-fc1d25070e47', 'Variant 2', 'TV-02', '120.00', 'dca07151-a309-4029-9ef2-c7da4b7c0518', 13, 0, '2022-01-06 11:54:03', '2022-01-06 11:54:03'),
('a8cb0287-e205-4290-98da-9cf9e8463ca6', 'Varian 2', 'TP004V2', '650.00', '5f33071d-8885-47ad-88bf-a3a736ec9b03', 19, 0, '2022-01-07 12:59:28', '2022-01-07 12:59:28'),
('d072ce26-e6a0-4a3e-81c2-7299a7198e7c', 'Varian 1', 'TP004V1', '750.00', '5f33071d-8885-47ad-88bf-a3a736ec9b03', 20, 0, '2022-01-07 12:59:28', '2022-01-07 12:59:28'),
('f4381167-1070-4d9e-908b-fd1d3585510f', 'Variant 1', 'TV-01', '120.00', 'dca07151-a309-4029-9ef2-c7da4b7c0518', 15, 0, '2022-01-06 11:54:03', '2022-01-06 11:54:03');

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
  `tgl_mulai` date DEFAULT NULL,
  `tgl_selesai` date DEFAULT NULL,
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Dumping data for table `session`
--

INSERT INTO `session` (`k`, `v`, `e`) VALUES
('03e89595-85e1-4c4d-81fa-bdf88c65b061', 0x0eff81040102ff8200010c01100000fe0116ff82000105746f6b656e06737472696e670cfe010100fffe4265617265722065794a68624763694f694a49557a49314e694973496e523563434936496b705856434a392e65794a6c6257467062434936496e4e31634756795957527461573541595752746157347559323974496977695a586877496a6f694d6a41794d6930774d7930774e3151784f446f304e6a6f774f4334314d5455354d444d334e7a4d724d4463364d4441694c434a705a4349364d53776961584e66646d567961575a705a5751694f6d5a6862484e6c4c434a306558426c496a6f69633356775a584967595752746157346966512e62677279436f6e6b62387a7a42485249785456324e6a783235426c316c69457556775753676f356c566e59, 1644666368),
('2e0c3922-f757-47fd-825f-19cd0960ae46', 0x0eff81040102ff8200010c01100000fe0113ff82000105746f6b656e06737472696e670cffff00fffc4265617265722065794a68624763694f694a49557a49314e694973496e523563434936496b705856434a392e65794a6c6257467062434936496d4a79595752735a586c6f5a584a745a584e765147647459576c734c6d4e7662534973496d563463434936496a49774d6a49744d444d744d4464554d6a49364d5445364d4463754f5445304e7a55304e6a6b344b7a41334f6a417749697769615751694f6a45774c434a70633139325a584a705a6d6c6c5a4349365a6d467363325573496e5235634755694f694a685a47317062694a392e557469557567515f734c504347373576355163644c656573495f686732675173705f655a43537666435051, 1644678667),
('3d91b044-2c43-44a0-a846-d37082d756ce', 0x0eff81040102ff8200010c01100000fe0113ff82000105746f6b656e06737472696e670cffff00fffc4265617265722065794a68624763694f694a49557a49314e694973496e523563434936496b705856434a392e65794a6c6257467062434936496d4a79595752735a586c6f5a584a745a584e765147647459576c734c6d4e7662534973496d563463434936496a49774d6a49744d444d744d5442554d4449364d4455364d4445754e6a51794d4445304d7a457a4b7a41334f6a417749697769615751694f6a45774c434a70633139325a584a705a6d6c6c5a4349365a6d467363325573496e5235634755694f694a685a47317062694a392e4a6d73726234434b4b694c735f73613630315641685049504357426b544c644e47684c7835726b7a6f3267, 1644865501),
('ba102d51-7174-4c34-a0cb-e4454ce924d0', 0x0eff81040102ff8200010c01100000fe0116ff82000105746f6b656e06737472696e670cfe010100fffe4265617265722065794a68624763694f694a49557a49314e694973496e523563434936496b705856434a392e65794a6c6257467062434936496e4e31634756795957527461573541595752746157347559323974496977695a586877496a6f694d6a41794d6930774d7930774e6c51794d446f774e7a6f784d4334324e7a6b344f4455354d5451724d4463364d4441694c434a705a4349364d53776961584e66646d567961575a705a5751694f6d5a6862484e6c4c434a306558426c496a6f69633356775a584967595752746157346966512e7a4e302d576e51706a2d3676635f645563726f344a6974445066686143644a4846782d5f5830594a364234, 1644584830),
('c74a3567-42d4-42b7-b6cb-25ac9b73f2bb', 0x0eff81040102ff8200010c01100000fe0111ff82000105746f6b656e06737472696e670cfffd00fffa4265617265722065794a68624763694f694a49557a49314e694973496e523563434936496b705856434a392e65794a6c6257467062434936496d35685a487068636d313164474678615734305147647459576c734c6d4e7662534973496d563463434936496a49774d6a49744d444d744d445a554d546b364e4449364e4459754e7a6b334e6a4d314e7a41344b7a41334f6a417749697769615751694f6a49334c434a70633139325a584a705a6d6c6c5a43493664484a315a53776964486c775a534936496e567a5a58496966512e61624476413965374c4c4a78627533573741617a4733757a4c59345837455a345951493364543246516c73, 1644583366);

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
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

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
  `tgl_lahir` date NOT NULL,
  `is_verified` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `hp`, `email`, `password`, `tgl_lahir`, `is_verified`, `created_at`, `updated_at`) VALUES
(24, 'user01', '081-555-656-12', 'user01@gmail.com', '$2a$10$SP2Bk/KSnVxK1IRUKnVrmeiQ8fCfj4F2hmqOlba2cGNTumuIPL05O', '1990-10-02', 1, '2021-12-03 04:30:17', '2021-12-03 04:30:17'),
(26, 'user02', '081-991-991-992', 'user02@gmail.com', '$2a$10$IueotRWZe.Tr9zsLizdzXu5jLUnDiGsnIS2xNjRNrh13UOljsHT4W', '2021-12-01', 1, '2021-12-29 01:55:41', '2021-12-29 01:55:41'),
(27, 'Nadzar Mutaqin', '081-111-222-333', 'nadzarmutaqin4@gmail.com', '$2a$10$kBWma06cu0g13MrND1tIz.fPcy5W2ab1y2WRzD40BzoRF0ScxlQNu', '1997-02-14', 1, '2022-02-03 11:46:32', '2022-02-03 11:46:32');

-- --------------------------------------------------------

--
-- Table structure for table `verifikasi`
--

CREATE TABLE `verifikasi` (
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `kode` varchar(200) COLLATE utf8mb3_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_unicode_ci;

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
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `alamat`
--
ALTER TABLE `alamat`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;

--
-- AUTO_INCREMENT for table `alamat_origin`
--
ALTER TABLE `alamat_origin`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `bank`
--
ALTER TABLE `bank`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=139;

--
-- AUTO_INCREMENT for table `carts`
--
ALTER TABLE `carts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=60;

--
-- AUTO_INCREMENT for table `feedback`
--
ALTER TABLE `feedback`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `mitra`
--
ALTER TABLE `mitra`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `order_bayar`
--
ALTER TABLE `order_bayar`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- AUTO_INCREMENT for table `order_item`
--
ALTER TABLE `order_item`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=500000040;

--
-- AUTO_INCREMENT for table `rekening`
--
ALTER TABLE `rekening`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

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
