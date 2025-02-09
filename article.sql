-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 09, 2025 at 07:07 PM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `article`
--

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` int(10) UNSIGNED NOT NULL,
  `title` varchar(200) NOT NULL,
  `content` text NOT NULL,
  `category` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `status` varchar(100) NOT NULL COMMENT 'Publish | Draft | Thrash'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `title`, `content`, `category`, `created_at`, `updated_at`, `status`) VALUES
(2, 'j8DKIFpR6h5LYqbVxXNW', 'SDRcc8z1iKELFr1APC2Dp7wi06yEjoMjIqZX92ZJaoetO3O09AqOYtvNMFSNPIbefSiumrwG8T5yb5Hk4oR7rWzUV6fPUqpdWlDcb0kdAgEMonGlXl7d3NFhIYk48smwbjXmEx0O6g6gqFMs5xap3aL9HnfCdyJxZAGtHyvfQS91TP4LhgCv2CLuGRcIJ2ZYaKmUkj7Y', '9FE', '2025-02-09 09:23:42', '2025-02-09 09:23:42', 'publish'),
(3, 'Exploring the beauty of nature through adventure and discovery in vast landscapes and hidden trails around the world brings joy and peace to the soul every day.', 'The journey begins with a single step into the unknown, where breathtaking landscapes unfold before our eyes. Towering mountains, lush forests, and serene rivers create a mesmerizing tapestry of nature’s finest artistry. As the sun sets, painting the sky in vibrant hues of orange and pink, we find ourselves immersed in the magic of the moment. Every trail we traverse whispers ancient stories of explorers who came before us, leaving behind traces of their dreams and aspirations. The crisp morning air fills our lungs, invigorating our spirits as we continue our adventure. Along the way, we encounter fascinating wildlife, their curious eyes watching our every move. The gentle rustling of leaves and the soothing sounds of flowing water provide a symphony of tranquility. With each passing day, we discover hidden gems—secret waterfalls cascading down rocky cliffs, meadows adorned with vibrant wildflowers, and caves echoing the whispers of history. These moments remind us of the profound connection we share with nature. As we reach the peak of our journey, standing atop a majestic summit, we realize that the true essence of adventure lies not in the destination but in the experiences we gather along the way.', 'adv', '2025-02-09 09:27:41', '2025-02-09 09:27:41', 'publish'),
(5, 'Embracing the wonders of the natural world through exploration, adventure, and discovery brings inspiration, joy, and endless memories to life', 'Nature has always been a source of wonder and inspiration. The vast mountains, deep forests, and endless oceans remind us of the beauty that exists beyond our daily routines. Walking along hidden trails, we uncover breathtaking views and unexpected encounters with wildlife. Birds sing melodious tunes, rivers carve paths through valleys, and the sky transforms with each passing hour. Every sunrise paints a masterpiece, while every sunset tells a different story. Traveling to remote places, we experience the peace that only nature can provide. The whispering wind through towering trees, the rhythmic crashing of waves, and the distant call of an eagle remind us of the world\'s harmony. Adventures take us to places untouched by modern life, where we connect with the earth in its purest form. Climbing rocky cliffs, crossing wooden bridges, and sitting beneath star-filled skies, we embrace the simplicity of existence. Each journey teaches us resilience, patience, and gratitude. The world is vast and filled with wonders waiting to be explored. In the end, it is not just about the places we visit but the experiences we cherish and the memories we carry forever.', 'ntr', '2025-02-09 12:42:50', '2025-02-09 12:42:50', 'draft');

-- --------------------------------------------------------

--
-- Table structure for table `schema_migrations`
--

CREATE TABLE `schema_migrations` (
  `version` bigint(20) NOT NULL,
  `dirty` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `schema_migrations`
--

INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
(20250209073811, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `schema_migrations`
--
ALTER TABLE `schema_migrations`
  ADD PRIMARY KEY (`version`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
