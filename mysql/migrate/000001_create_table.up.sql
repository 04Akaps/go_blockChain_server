CREATE TABLE `evmLaunchpad` (
	`id` bigint PRIMARY KEY AUTO_INCREMENT,
	`eoa_address` varchar(255) NOT NULL,
	`contract_address` varchar(255) NOT NULL,
	`network_chain_id` integer NOT NULL,
	`price` integer NOT NULL,
	`meta_data_uri` varchar(255) NOT NULL,
	`created_at` varchar(255) NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

  CREATE TABLE `aptosLaunchpad` (
	`id` bigint PRIMARY KEY AUTO_INCREMENT,
	`user_address` varchar(255) NOT NULL,
	`module_address` varchar(255) NOT NULL,
	`price` integer NOT NULL,
	`created_at` varchar(255) NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

  CREATE INDEX `evmLaunchpad_index_0` ON `evmLaunchpad` (`eoaAddress`);

  CREATE INDEX `aptosLaunchpad_index_1` ON `aptosLaunchpad` (`userAddress`);
