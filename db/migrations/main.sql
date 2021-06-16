SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

CREATE TABLE `main` (
  `id` int(11) NOT NULL,
  `currency1` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `currency2` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `rate` float NOT NULL,
  `updated_at` text COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `main`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `main`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

